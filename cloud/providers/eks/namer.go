package eks

import (
	"fmt"

	"github.com/appscode/go/crypto/rand"
	api "github.com/pharmer/pharmer/apis/v1alpha1"
)

type namer struct {
	cluster *api.Cluster
}

func (n namer) AdminUsername() string {
	return "pharmer"
}

func (n namer) GenSSHKeyExternalID() string {
	return n.cluster.Name + "-" + rand.Characters(6)
}

func (n namer) GetStackServiceRole() string {
	return fmt.Sprintf("EKS-%v-ServiceRole", n.cluster.Name)
}

func (n namer) GetClusterVPC() string {
	return fmt.Sprintf("EKS-%v-VPC", n.cluster.Name)
}
