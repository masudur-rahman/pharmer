package store

import (
	"crypto/rsa"
	"crypto/x509"
	"errors"

	"github.com/appscode/pharmer/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NotImplemented = errors.New("Not implemented")

type Interface interface {
	Credentials() CredentialStore

	Clusters() ClusterStore
	NodeSets(cluster string) NodeSetStore
	Instances(cluster string) InstanceStore
	Certificates(cluster string) CertificateStore
	SSHKeys(cluster string) SSHKeyStore
}

type CredentialStore interface {
	List(opts metav1.ListOptions) ([]*api.Credential, error)
	Get(name string) (*api.Credential, error)
	Create(obj *api.Credential) (*api.Credential, error)
	Update(obj *api.Credential) (*api.Credential, error)
	Delete(name string) error
}

type ClusterStore interface {
	List(opts metav1.ListOptions) ([]*api.Cluster, error)
	Get(name string) (*api.Cluster, error)
	Create(obj *api.Cluster) (*api.Cluster, error)
	Update(obj *api.Cluster) (*api.Cluster, error)
	Delete(name string) error
	UpdateStatus(obj *api.Cluster) (*api.Cluster, error)
}

type NodeSetStore interface {
	List(opts metav1.ListOptions) ([]*api.NodeSet, error)
	Get(name string) (*api.NodeSet, error)
	Create(obj *api.NodeSet) (*api.NodeSet, error)
	Update(obj *api.NodeSet) (*api.NodeSet, error)
	Delete(name string) error
	UpdateStatus(obj *api.NodeSet) (*api.NodeSet, error)
}

type InstanceStore interface {
	List(opts metav1.ListOptions) ([]*api.Node, error)
	Get(name string) (*api.Node, error)
	Create(obj *api.Node) (*api.Node, error)
	Update(obj *api.Node) (*api.Node, error)
	Delete(name string) error
	UpdateStatus(obj *api.Node) (*api.Node, error)
}

type CertificateStore interface {
	Get(name string) (*x509.Certificate, *rsa.PrivateKey, error)
	Create(name string, crt *x509.Certificate, key *rsa.PrivateKey) error
	Delete(name string) error
}

type SSHKeyStore interface {
	Get(name string) (pubKey, privKey []byte, err error)
	Create(name string, pubKey, privKey []byte) error
	Delete(name string) error
}
