/*
Copyright 2017 The Kubernetes Authors.

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
package fake

import (
	group_v1 "github.com/alok87/apik8s/pkg/apis/group/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMykinds implements MykindInterface
type FakeMykinds struct {
	Fake *FakeGroupV1
	ns   string
}

var mykindsResource = schema.GroupVersionResource{Group: "group.mydomain", Version: "v1", Resource: "mykinds"}

var mykindsKind = schema.GroupVersionKind{Group: "group.mydomain", Version: "v1", Kind: "Mykind"}

// Get takes name of the mykind, and returns the corresponding mykind object, and an error if there is any.
func (c *FakeMykinds) Get(name string, options v1.GetOptions) (result *group_v1.Mykind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mykindsResource, c.ns, name), &group_v1.Mykind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*group_v1.Mykind), err
}

// List takes label and field selectors, and returns the list of Mykinds that match those selectors.
func (c *FakeMykinds) List(opts v1.ListOptions) (result *group_v1.MykindList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mykindsResource, mykindsKind, c.ns, opts), &group_v1.MykindList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &group_v1.MykindList{}
	for _, item := range obj.(*group_v1.MykindList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mykinds.
func (c *FakeMykinds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mykindsResource, c.ns, opts))

}

// Create takes the representation of a mykind and creates it.  Returns the server's representation of the mykind, and an error, if there is any.
func (c *FakeMykinds) Create(mykind *group_v1.Mykind) (result *group_v1.Mykind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mykindsResource, c.ns, mykind), &group_v1.Mykind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*group_v1.Mykind), err
}

// Update takes the representation of a mykind and updates it. Returns the server's representation of the mykind, and an error, if there is any.
func (c *FakeMykinds) Update(mykind *group_v1.Mykind) (result *group_v1.Mykind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mykindsResource, c.ns, mykind), &group_v1.Mykind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*group_v1.Mykind), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMykinds) UpdateStatus(mykind *group_v1.Mykind) (*group_v1.Mykind, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mykindsResource, "status", c.ns, mykind), &group_v1.Mykind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*group_v1.Mykind), err
}

// Delete takes name of the mykind and deletes it. Returns an error if one occurs.
func (c *FakeMykinds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mykindsResource, c.ns, name), &group_v1.Mykind{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMykinds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mykindsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &group_v1.MykindList{})
	return err
}

// Patch applies the patch and returns the patched mykind.
func (c *FakeMykinds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *group_v1.Mykind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mykindsResource, c.ns, name, data, subresources...), &group_v1.Mykind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*group_v1.Mykind), err
}
