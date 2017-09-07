package cloud

import (
	"errors"

	"github.com/appscode/pharmer/api"
)

var InstanceNotFound = errors.New("Instance not found")
var UnsupportedOperation = errors.New("Unsupported operation")

type ClusterManager interface {
	DefaultSpec(cluster *api.Cluster) (*api.Cluster, error)
	CreateMasterNodeSet(cluster *api.Cluster) (*api.NodeSet, error)
	Apply(cluster *api.Cluster, dryRun bool) error
	IsValid(cluster *api.Cluster) (bool, error)

	// IsValid(cluster *api.Cluster) (bool, error)
	// Delete(req *proto.ClusterDeleteRequest) error
	// SetVersion(req *proto.ClusterReconfigureRequest) error
	// Scale(req *proto.ClusterReconfigureRequest) error
	// GetInstance(md *api.InstanceStatus) (*api.Instance, error)
}
