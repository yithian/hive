// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/hive/pkg/apis/hive/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MachinePoolNameLeaseLister helps list MachinePoolNameLeases.
type MachinePoolNameLeaseLister interface {
	// List lists all MachinePoolNameLeases in the indexer.
	List(selector labels.Selector) (ret []*v1.MachinePoolNameLease, err error)
	// MachinePoolNameLeases returns an object that can list and get MachinePoolNameLeases.
	MachinePoolNameLeases(namespace string) MachinePoolNameLeaseNamespaceLister
	MachinePoolNameLeaseListerExpansion
}

// machinePoolNameLeaseLister implements the MachinePoolNameLeaseLister interface.
type machinePoolNameLeaseLister struct {
	indexer cache.Indexer
}

// NewMachinePoolNameLeaseLister returns a new MachinePoolNameLeaseLister.
func NewMachinePoolNameLeaseLister(indexer cache.Indexer) MachinePoolNameLeaseLister {
	return &machinePoolNameLeaseLister{indexer: indexer}
}

// List lists all MachinePoolNameLeases in the indexer.
func (s *machinePoolNameLeaseLister) List(selector labels.Selector) (ret []*v1.MachinePoolNameLease, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MachinePoolNameLease))
	})
	return ret, err
}

// MachinePoolNameLeases returns an object that can list and get MachinePoolNameLeases.
func (s *machinePoolNameLeaseLister) MachinePoolNameLeases(namespace string) MachinePoolNameLeaseNamespaceLister {
	return machinePoolNameLeaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MachinePoolNameLeaseNamespaceLister helps list and get MachinePoolNameLeases.
type MachinePoolNameLeaseNamespaceLister interface {
	// List lists all MachinePoolNameLeases in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.MachinePoolNameLease, err error)
	// Get retrieves the MachinePoolNameLease from the indexer for a given namespace and name.
	Get(name string) (*v1.MachinePoolNameLease, error)
	MachinePoolNameLeaseNamespaceListerExpansion
}

// machinePoolNameLeaseNamespaceLister implements the MachinePoolNameLeaseNamespaceLister
// interface.
type machinePoolNameLeaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MachinePoolNameLeases in the indexer for a given namespace.
func (s machinePoolNameLeaseNamespaceLister) List(selector labels.Selector) (ret []*v1.MachinePoolNameLease, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MachinePoolNameLease))
	})
	return ret, err
}

// Get retrieves the MachinePoolNameLease from the indexer for a given namespace and name.
func (s machinePoolNameLeaseNamespaceLister) Get(name string) (*v1.MachinePoolNameLease, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("machinepoolnamelease"), name)
	}
	return obj.(*v1.MachinePoolNameLease), nil
}