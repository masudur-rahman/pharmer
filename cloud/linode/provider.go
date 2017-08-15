package linode

import (
	proto "github.com/appscode/api/kubernetes/v1beta1"
	"github.com/appscode/pharmer/api"
	"github.com/appscode/pharmer/cloud"
)

const (
	UID = "linode"
)

func init() {
	cloud.RegisterProvider(UID, new(provider))
}

type provider struct {
}

var _ cloud.Provider = &provider{}

func (p *provider) Create(ctx *api.Cluster, req *proto.ClusterCreateRequest) error {
	return (&clusterManager{ctx: ctx}).create(req)
}

func (p *provider) Scale(ctx *api.Cluster, req *proto.ClusterReconfigureRequest) error {
	return cloud.UnsupportedOperation
}

func (p *provider) Delete(ctx *api.Cluster, req *proto.ClusterDeleteRequest) error {
	return (&clusterManager{ctx: ctx}).delete(req)
}

func (p *provider) SetVersion(ctx *api.Cluster, req *proto.ClusterReconfigureRequest) error {
	return cloud.UnsupportedOperation
}

func (p *provider) UploadStartupConfig(ctx *api.Cluster) error {
	conn, err := NewConnector(ctx)
	if err != nil {
		return err
	}
	cm := &clusterManager{ctx: ctx, conn: conn}
	return cm.UploadStartupConfig()
}

func (p *provider) GetInstance(ctx *api.Cluster, md *api.InstanceMetadata) (*api.KubernetesInstance, error) {
	conn, err := NewConnector(ctx)
	if err != nil {
		return nil, err
	}
	im := &instanceManager{conn: conn}
	return im.GetInstance(md)
}

func (p *provider) MatchInstance(i *api.KubernetesInstance, md *api.InstanceMetadata) bool {
	return i.InternalIP == md.InternalIP
}
