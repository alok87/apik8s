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
package v1

import (
	v1 "github.com/alok87/apik8s/pkg/apis/authentication/v1"
	scheme "github.com/alok87/apik8s/pkg/client/clientset_generated/clientset/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TokensGetter has a method to return a TokenInterface.
// A group's client should implement this interface.
type TokensGetter interface {
	Tokens(namespace string) TokenInterface
}

// TokenInterface has methods to work with Token resources.
type TokenInterface interface {
	Create(*v1.Token) (*v1.Token, error)
	Update(*v1.Token) (*v1.Token, error)
	UpdateStatus(*v1.Token) (*v1.Token, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Token, error)
	List(opts meta_v1.ListOptions) (*v1.TokenList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Token, err error)
	TokenExpansion
}

// tokens implements TokenInterface
type tokens struct {
	client rest.Interface
	ns     string
}

// newTokens returns a Tokens
func newTokens(c *AuthenticationV1Client, namespace string) *tokens {
	return &tokens{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the token, and returns the corresponding token object, and an error if there is any.
func (c *tokens) Get(name string, options meta_v1.GetOptions) (result *v1.Token, err error) {
	result = &v1.Token{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tokens").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Tokens that match those selectors.
func (c *tokens) List(opts meta_v1.ListOptions) (result *v1.TokenList, err error) {
	result = &v1.TokenList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tokens.
func (c *tokens) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a token and creates it.  Returns the server's representation of the token, and an error, if there is any.
func (c *tokens) Create(token *v1.Token) (result *v1.Token, err error) {
	result = &v1.Token{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tokens").
		Body(token).
		Do().
		Into(result)
	return
}

// Update takes the representation of a token and updates it. Returns the server's representation of the token, and an error, if there is any.
func (c *tokens) Update(token *v1.Token) (result *v1.Token, err error) {
	result = &v1.Token{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tokens").
		Name(token.Name).
		Body(token).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *tokens) UpdateStatus(token *v1.Token) (result *v1.Token, err error) {
	result = &v1.Token{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tokens").
		Name(token.Name).
		SubResource("status").
		Body(token).
		Do().
		Into(result)
	return
}

// Delete takes name of the token and deletes it. Returns an error if one occurs.
func (c *tokens) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tokens").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tokens) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tokens").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched token.
func (c *tokens) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Token, err error) {
	result = &v1.Token{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tokens").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
