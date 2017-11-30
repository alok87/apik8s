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
package internalversion

import (
	group "github.com/alok87/apik8s/pkg/apis/group"
	scheme "github.com/alok87/apik8s/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MykindsGetter has a method to return a MykindInterface.
// A group's client should implement this interface.
type MykindsGetter interface {
	Mykinds(namespace string) MykindInterface
}

// MykindInterface has methods to work with Mykind resources.
type MykindInterface interface {
	Create(*group.Mykind) (*group.Mykind, error)
	Update(*group.Mykind) (*group.Mykind, error)
	UpdateStatus(*group.Mykind) (*group.Mykind, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*group.Mykind, error)
	List(opts v1.ListOptions) (*group.MykindList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *group.Mykind, err error)
	MykindExpansion
}

// mykinds implements MykindInterface
type mykinds struct {
	client rest.Interface
	ns     string
}

// newMykinds returns a Mykinds
func newMykinds(c *GroupClient, namespace string) *mykinds {
	return &mykinds{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mykind, and returns the corresponding mykind object, and an error if there is any.
func (c *mykinds) Get(name string, options v1.GetOptions) (result *group.Mykind, err error) {
	result = &group.Mykind{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mykinds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Mykinds that match those selectors.
func (c *mykinds) List(opts v1.ListOptions) (result *group.MykindList, err error) {
	result = &group.MykindList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mykinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mykinds.
func (c *mykinds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mykinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a mykind and creates it.  Returns the server's representation of the mykind, and an error, if there is any.
func (c *mykinds) Create(mykind *group.Mykind) (result *group.Mykind, err error) {
	result = &group.Mykind{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mykinds").
		Body(mykind).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mykind and updates it. Returns the server's representation of the mykind, and an error, if there is any.
func (c *mykinds) Update(mykind *group.Mykind) (result *group.Mykind, err error) {
	result = &group.Mykind{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mykinds").
		Name(mykind.Name).
		Body(mykind).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *mykinds) UpdateStatus(mykind *group.Mykind) (result *group.Mykind, err error) {
	result = &group.Mykind{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mykinds").
		Name(mykind.Name).
		SubResource("status").
		Body(mykind).
		Do().
		Into(result)
	return
}

// Delete takes name of the mykind and deletes it. Returns an error if one occurs.
func (c *mykinds) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mykinds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mykinds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mykinds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mykind.
func (c *mykinds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *group.Mykind, err error) {
	result = &group.Mykind{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mykinds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
