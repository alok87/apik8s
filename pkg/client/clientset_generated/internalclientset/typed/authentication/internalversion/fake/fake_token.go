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
	authentication "github.com/alok87/apik8s/pkg/apis/authentication"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTokens implements TokenInterface
type FakeTokens struct {
	Fake *FakeAuthentication
	ns   string
}

var tokensResource = schema.GroupVersionResource{Group: "authentication", Version: "", Resource: "tokens"}

var tokensKind = schema.GroupVersionKind{Group: "authentication", Version: "", Kind: "Token"}

// Get takes name of the token, and returns the corresponding token object, and an error if there is any.
func (c *FakeTokens) Get(name string, options v1.GetOptions) (result *authentication.Token, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tokensResource, c.ns, name), &authentication.Token{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authentication.Token), err
}

// List takes label and field selectors, and returns the list of Tokens that match those selectors.
func (c *FakeTokens) List(opts v1.ListOptions) (result *authentication.TokenList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tokensResource, tokensKind, c.ns, opts), &authentication.TokenList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &authentication.TokenList{}
	for _, item := range obj.(*authentication.TokenList).Items {
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
func (c *FakeTokens) Create(token *authentication.Token) (result *authentication.Token, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tokensResource, c.ns, token), &authentication.Token{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authentication.Token), err
}

// Update takes the representation of a token and updates it. Returns the server's representation of the token, and an error, if there is any.
func (c *FakeTokens) Update(token *authentication.Token) (result *authentication.Token, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tokensResource, c.ns, token), &authentication.Token{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authentication.Token), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTokens) UpdateStatus(token *authentication.Token) (*authentication.Token, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tokensResource, "status", c.ns, token), &authentication.Token{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authentication.Token), err
}

// Delete takes name of the token and deletes it. Returns an error if one occurs.
func (c *FakeTokens) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tokensResource, c.ns, name), &authentication.Token{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTokens) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tokensResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &authentication.TokenList{})
	return err
}

// Patch applies the patch and returns the patched token.
func (c *FakeTokens) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *authentication.Token, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tokensResource, c.ns, name, data, subresources...), &authentication.Token{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authentication.Token), err
}
