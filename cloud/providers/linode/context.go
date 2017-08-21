package linode

import (
	proto "github.com/appscode/api/kubernetes/v1beta1"
	"github.com/appscode/errors"
	"github.com/appscode/go/crypto/rand"
	_env "github.com/appscode/go/env"
	"github.com/appscode/pharmer/api"
	"github.com/appscode/pharmer/cloud"
	"github.com/appscode/pharmer/context"
	"github.com/appscode/pharmer/phid"
)

type ClusterManager struct {
	ctx     context.Context
	cluster *api.Cluster
	ins     *api.ClusterInstances
	conn    *cloudConnector
	namer   namer
}

var _ cloud.ClusterProvider = &ClusterManager{}

const (
	UID = "linode"
)

func init() {
	cloud.RegisterCloudProvider(UID, func(ctx context.Context) (cloud.Interface, error) { return New(ctx), nil })
}

func New(ctx context.Context) cloud.Interface {
	return &ClusterManager{ctx: ctx}
}

func (cm *ClusterManager) Clusters() cloud.ClusterProvider {
	return cm
}

func (cm *ClusterManager) Credentials() cloud.CredentialProvider {
	return cm
}

func (p *ClusterManager) Scale(req *proto.ClusterReconfigureRequest) error {
	return cloud.UnsupportedOperation
}

func (p *ClusterManager) SetVersion(req *proto.ClusterReconfigureRequest) error {
	return cloud.UnsupportedOperation
}

func (c *ClusterManager) GetInstance(md *api.InstanceMetadata) (*api.Instance, error) {
	conn, err := NewConnector(c.ctx, nil)
	if err != nil {
		return nil, err
	}
	im := &instanceManager{conn: conn}
	return im.GetInstance(md)
}

func (p *ClusterManager) MatchInstance(i *api.Instance, md *api.InstanceMetadata) bool {
	return i.Status.InternalIP == md.InternalIP
}

func (cm *ClusterManager) initContext(req *proto.ClusterCreateRequest) error {
	err := cm.LoadDefaultContext()
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.namer = namer{cluster: cm.cluster}

	//cluster.Spec.ctx.Name = req.Name
	//cluster.Spec.ctx.PHID = phid.NewKubeCluster()
	//cluster.Spec.ctx.Provider = req.Provider
	//cluster.Spec.ctx.Zone = req.Zone

	cm.cluster.Spec.Region = cm.cluster.Spec.Zone
	cm.cluster.Spec.DoNotDelete = req.DoNotDelete

	cm.cluster.SetNodeGroups(req.NodeGroups)

	cm.cluster.Spec.KubernetesMasterName = cm.namer.MasterName()
	cm.cluster.Spec.SSHKey, err = api.NewSSHKeyPair()
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.cluster.Spec.SSHKeyExternalID = cm.namer.GenSSHKeyExternalID()
	cm.cluster.Spec.SSHKeyPHID = phid.NewSSHKey()

	cloud.GenClusterTokens(cm.cluster)

	return nil
}

func (cm *ClusterManager) LoadDefaultContext() error {
	err := cloud.LoadDefaultGenericContext(cm.ctx, cm.cluster)
	if err != nil {
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	cm.cluster.Spec.OS = "debian"
	cm.cluster.Spec.InstanceRootPassword = "@dmin123"
	if _env.FromHost().IsPublic() {
		cm.cluster.Spec.InstanceRootPassword = rand.GeneratePassword()
	}
	cm.cluster.Spec.MasterSKU = "2" // plan_id 2 label {"Linode 4096"} cpu 2 ram 4096 disk 48

	cloud.BuildRuntimeConfig(cm.cluster)
	return nil
}

func (cm *ClusterManager) UploadStartupConfig() error {
	if api.UseFirebase() {
		return cloud.UploadStartupConfigInFirebase(cm.ctx, cm.cluster)
	}
	return nil
}
