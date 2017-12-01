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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package authentication

import (
	"fmt"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
)

var (
	InternalToken = builders.NewInternalResource(
		"tokens",
		func() runtime.Object { return &Token{} },
		func() runtime.Object { return &TokenList{} },
	)
	InternalTokenStatus = builders.NewInternalResourceStatus(
		"tokens",
		func() runtime.Object { return &Token{} },
		func() runtime.Object { return &TokenList{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("authentication.practodev.com").WithKinds(
		InternalToken,
		InternalTokenStatus,
	)

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Token struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   TokenSpec
	Status TokenStatus
}

type TokenSpec struct {
}

type TokenStatus struct {
}

//
// Token Functions and Structs
//
// +k8s:deepcopy-gen=false
type TokenStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type TokenStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TokenList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Token
}

func (Token) NewStatus() interface{} {
	return TokenStatus{}
}

func (pc *Token) GetStatus() interface{} {
	return pc.Status
}

func (pc *Token) SetStatus(s interface{}) {
	pc.Status = s.(TokenStatus)
}

func (pc *Token) GetSpec() interface{} {
	return pc.Spec
}

func (pc *Token) SetSpec(s interface{}) {
	pc.Spec = s.(TokenSpec)
}

func (pc *Token) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *Token) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc Token) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store Token.
// +k8s:deepcopy-gen=false
type TokenRegistry interface {
	ListTokens(ctx request.Context, options *internalversion.ListOptions) (*TokenList, error)
	GetToken(ctx request.Context, id string, options *metav1.GetOptions) (*Token, error)
	CreateToken(ctx request.Context, id *Token) (*Token, error)
	UpdateToken(ctx request.Context, id *Token) (*Token, error)
	DeleteToken(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewTokenRegistry(sp builders.StandardStorageProvider) TokenRegistry {
	return &storageToken{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageToken struct {
	builders.StandardStorageProvider
}

func (s *storageToken) ListTokens(ctx request.Context, options *internalversion.ListOptions) (*TokenList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*TokenList), err
}

func (s *storageToken) GetToken(ctx request.Context, id string, options *metav1.GetOptions) (*Token, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*Token), nil
}

func (s *storageToken) CreateToken(ctx request.Context, object *Token) (*Token, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, false)
	if err != nil {
		return nil, err
	}
	return obj.(*Token), nil
}

func (s *storageToken) UpdateToken(ctx request.Context, object *Token) (*Token, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object, builders.Scheme))
	if err != nil {
		return nil, err
	}
	return obj.(*Token), nil
}

func (s *storageToken) DeleteToken(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}
