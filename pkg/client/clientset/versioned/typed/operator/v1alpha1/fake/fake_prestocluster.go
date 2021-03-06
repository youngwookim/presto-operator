/*
Copyright 2019 oneonestar.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/oneonestar/presto-operator/pkg/apis/operator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePrestoClusters implements PrestoClusterInterface
type FakePrestoClusters struct {
	Fake *FakeOperatorV1alpha1
	ns   string
}

var prestoclustersResource = schema.GroupVersionResource{Group: "operator.prestosql.io", Version: "v1alpha1", Resource: "prestoclusters"}

var prestoclustersKind = schema.GroupVersionKind{Group: "operator.prestosql.io", Version: "v1alpha1", Kind: "PrestoCluster"}

// Get takes name of the prestoCluster, and returns the corresponding prestoCluster object, and an error if there is any.
func (c *FakePrestoClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.PrestoCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(prestoclustersResource, c.ns, name), &v1alpha1.PrestoCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrestoCluster), err
}

// List takes label and field selectors, and returns the list of PrestoClusters that match those selectors.
func (c *FakePrestoClusters) List(opts v1.ListOptions) (result *v1alpha1.PrestoClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(prestoclustersResource, prestoclustersKind, c.ns, opts), &v1alpha1.PrestoClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PrestoClusterList{ListMeta: obj.(*v1alpha1.PrestoClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.PrestoClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested prestoClusters.
func (c *FakePrestoClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(prestoclustersResource, c.ns, opts))

}

// Create takes the representation of a prestoCluster and creates it.  Returns the server's representation of the prestoCluster, and an error, if there is any.
func (c *FakePrestoClusters) Create(prestoCluster *v1alpha1.PrestoCluster) (result *v1alpha1.PrestoCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(prestoclustersResource, c.ns, prestoCluster), &v1alpha1.PrestoCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrestoCluster), err
}

// Update takes the representation of a prestoCluster and updates it. Returns the server's representation of the prestoCluster, and an error, if there is any.
func (c *FakePrestoClusters) Update(prestoCluster *v1alpha1.PrestoCluster) (result *v1alpha1.PrestoCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(prestoclustersResource, c.ns, prestoCluster), &v1alpha1.PrestoCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrestoCluster), err
}

// Delete takes name of the prestoCluster and deletes it. Returns an error if one occurs.
func (c *FakePrestoClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(prestoclustersResource, c.ns, name), &v1alpha1.PrestoCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrestoClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(prestoclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PrestoClusterList{})
	return err
}

// Patch applies the patch and returns the patched prestoCluster.
func (c *FakePrestoClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PrestoCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(prestoclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.PrestoCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrestoCluster), err
}
