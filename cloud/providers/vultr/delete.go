package vultr

import (
	"fmt"
	"strings"

	proto "github.com/appscode/api/kubernetes/v1beta1"
	"github.com/appscode/errors"
	"github.com/appscode/pharmer/api"
	"github.com/appscode/pharmer/cloud"
	"github.com/cenkalti/backoff"
)

func (cm *ClusterManager) Delete(req *proto.ClusterDeleteRequest) error {
	defer cm.cluster.Delete()

	if cm.cluster.Status.Phase == api.ClusterPhasePending {
		cm.cluster.Status.Phase = api.ClusterPhaseFailing
	} else if cm.cluster.Status.Phase == api.ClusterPhaseReady {
		cm.cluster.Status.Phase = api.ClusterPhaseDeleting
	}
	// cm.ctx.Store().UpdateKubernetesStatus(cm.ctx.PHID, cm.ctx.Status)

	var err error
	if cm.conn == nil {
		cm.conn, err = NewConnector(cm.ctx, cm.cluster)
		if err != nil {
			cm.cluster.Status.Reason = err.Error()
			return errors.FromErr(err).WithContext(cm.ctx).Err()
		}
	}
	cm.namer = namer{cluster: cm.cluster}
	cm.ins, err = cloud.NewInstances(cm.ctx, cm.cluster)
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.ins.Instances, err = cm.ctx.Store().Instances().LoadInstances(cm.cluster.Name)
	if err != nil {
		cm.cluster.Status.Reason = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	var errs []string
	if cm.cluster.Status.Reason != "" {
		errs = append(errs, cm.cluster.Status.Reason)
	}

	for _, i := range cm.ins.Instances {
		backoff.Retry(func() error {
			err := cm.conn.client.DeleteServer(i.Status.ExternalID)
			if err != nil {
				return err
			}

			return nil
		}, backoff.NewExponentialBackOff())
		cm.ctx.Logger().Infof("Instance %v with id %v for clutser is deleted", i.Name, i.Status.ExternalID, cm.cluster.Name)
	}

	if req.ReleaseReservedIp && cm.cluster.Spec.MasterReservedIP != "" {
		backoff.Retry(func() error {
			return cm.releaseReservedIP(cm.cluster.Spec.MasterReservedIP)
		}, backoff.NewExponentialBackOff())
		cm.ctx.Logger().Infof("Reserved ip for cluster %v", cm.cluster.Name)
	}

	cm.ctx.Logger().Infof("Deleting startup scripts for cluster %v", cm.cluster.Name)
	backoff.Retry(cm.deleteStartupScript, backoff.NewExponentialBackOff())

	// Delete SSH key from DB
	if err := cm.deleteSSHKey(); err != nil {
		errs = append(errs, err.Error())
	}

	if err := cloud.DeleteARecords(cm.ctx, cm.cluster); err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		// Preserve statusCause for failed cluster
		if cm.cluster.Status.Phase == api.ClusterPhaseDeleting {
			cm.cluster.Status.Reason = strings.Join(errs, "\n")
		}
		return fmt.Errorf(strings.Join(errs, "\n"))
	}

	cm.ctx.Logger().Infof("Cluster %v is deleted successfully", cm.cluster.Name)
	return nil
}

func (cm *ClusterManager) releaseReservedIP(ip string) error {
	cm.ctx.Logger().Debugln("Deleting Floating IP", ip)
	err := cm.conn.client.DestroyReservedIP(ip)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	return nil
}

func (cm *ClusterManager) deleteStartupScript() error {
	scripts, err := cm.conn.client.GetStartupScripts()
	if err != nil {
		return err
	}
	for _, script := range scripts {
		if strings.HasPrefix(script.Name, cm.cluster.Name+"-") {
			err := cm.conn.client.DeleteStartupScript(script.ID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (cm *ClusterManager) deleteSSHKey() (err error) {
	cm.ctx.Logger().Infof("Deleting ssh key for cluster %v", cm.cluster.Spec.MasterDiskId)

	if cm.cluster.Spec.SSHKey != nil {
		backoff.Retry(func() error {
			return cm.conn.client.DeleteSSHKey(cm.cluster.Spec.SSHKeyExternalID)
		}, backoff.NewExponentialBackOff())
		cm.ctx.Logger().Infof("SSH key for cluster %v is deleted", cm.cluster.Name)
	}

	if cm.cluster.Spec.SSHKeyPHID != "" {
		//updates := &storage.SSHKey{IsDeleted: 1}
		//cond := &storage.SSHKey{PHID: cm.ctx.SSHKeyPHID}
		//_, err = cm.ctx.Store().Engine.Update(updates, cond)
	}
	return
}
