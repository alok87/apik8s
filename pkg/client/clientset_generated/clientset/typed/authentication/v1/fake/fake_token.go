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
	authentication_v1 "github.com/alok87/apik8s/pkg/apis/authentication/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTokens implements TokenInterface
type FakeTokens struct {
	Fake *FakeAuthenticationV1
	ns   string
}

var tokensResource = schema.GroupVersionResource{Group: "authentication.practodev.com", Version: "v1", Resource: "tokens"}

var tokensKind = schema.GroupVersionKind{Group: "authentication.practodev.com", Version: "v1", Kind: "Token"}

// Get takes name of the token, and returns the corresponding token object, and an error if there is any.
func (c *FakeTokens) Get(name string, options v1.GetOptions) (result *authentication_v1.Token, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tokensResource, c.ns, name), &authentication_v1.Token{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authentication_v1.Token), err
}

// List takes label and field selectors, and returns the list of Tokens that match those selectors.
func (c *FakeTokens) List(opts v1.ListOptions) (result *authentication_v1.TokenList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tokensResource, tokensKind, c.ns, opts), &authentication_v1.TokenList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &authentication_v1.TokenList{}
	for _, item := range obj.(*authentication_v1.TokenList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tokens.
func (c *FakeTokens) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tokensResource, c.ns, opts))

}

// Create takes the representation of a token and creates it.  Returns the server's representation of the token, and an error, if there is any.
func (c *FakeTokens) Create(token *authentication_v1.Token) (result *authentication_v1.Token, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tokensResource, c.ns, token), &authentication_v1.Token{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authentication_v1.Token), err
}

// Update takes the representation of a token and updates it. Returns the server's representation of the token, and an error, if there is any.
func (c *FakeTokens) Update(token *authentication_v1.Token) (result *authentication_v1.Token, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tokensResource, c.ns, token), &authentication_v1.Token{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authentication_v1.Token), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTokens) UpdateStatus(token *authentication_v1.Token) (*authentication_v1.Token, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tokensResource, "status", c.ns, token), &authentication_v1.Token{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authentication_v1.Token), err
}

// Delete takes name of the token and deletes it. Returns an error if one occurs.
func (c *FakeTokens) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tokensResource, c.ns, name), &authentication_v1.Token{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTokens) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tokensResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &authentication_v1.TokenList{})
	return err
}

// Patch applies the patch and returns the patched token.
func (c *FakeTokens) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authentication_v1.Token, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tokensResource, c.ns, name, data, subresources...), &authentication_v1.Token{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authentication_v1.Token), err
}
