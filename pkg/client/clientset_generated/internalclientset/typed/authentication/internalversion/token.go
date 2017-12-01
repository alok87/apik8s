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
	authentication "github.com/alok87/apik8s/pkg/apis/authentication"
	scheme "github.com/alok87/apik8s/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Create(*authentication.Token) (*authentication.Token, error)
	Update(*authentication.Token) (*authentication.Token, error)
	UpdateStatus(*authentication.Token) (*authentication.Token, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*authentication.Token, error)
	List(opts v1.ListOptions) (*authentication.TokenList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authentication.Token, err error)
	TokenExpansion
}

// tokens implements TokenInterface
type tokens struct {
	client rest.Interface
	ns     string
}

// newTokens returns a Tokens
func newTokens(c *AuthenticationClient, namespace string) *tokens {
	return &tokens{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the token, and returns the corresponding token object, and an error if there is any.
func (c *tokens) Get(name string, options v1.GetOptions) (result *authentication.Token, err error) {
	result = &authentication.Token{}
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
func (c *tokens) List(opts v1.ListOptions) (result *authentication.TokenList, err error) {
	result = &authentication.TokenList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tokens.
func (c *tokens) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a token and creates it.  Returns the server's representation of the token, and an error, if there is any.
func (c *tokens) Create(token *authentication.Token) (result *authentication.Token, err error) {
	result = &authentication.Token{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tokens").
		Body(token).
		Do().
		Into(result)
	return
}

// Update takes the representation of a token and updates it. Returns the server's representation of the token, and an error, if there is any.
func (c *tokens) Update(token *authentication.Token) (result *authentication.Token, err error) {
	result = &authentication.Token{}
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

func (c *tokens) UpdateStatus(token *authentication.Token) (result *authentication.Token, err error) {
	result = &authentication.Token{}
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
func (c *tokens) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tokens").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tokens) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tokens").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched token.
func (c *tokens) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authentication.Token, err error) {
	result = &authentication.Token{}
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
