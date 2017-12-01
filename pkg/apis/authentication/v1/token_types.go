
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
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/alok87/apik8s/pkg/apis/authentication"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Token
// +k8s:openapi-gen=true
// +resource:path=tokens,strategy=TokenStrategy
type Token struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TokenSpec   `json:"spec,omitempty"`
	Status TokenStatus `json:"status,omitempty"`
}

// TokenSpec defines the desired state of Token
type TokenSpec struct {
}

// TokenStatus defines the observed state of Token
type TokenStatus struct {
}

// Validate checks that an instance of Token is well formed
func (TokenStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*authentication.Token)
	log.Printf("Validating fields for Token %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default Token field values
func (TokenSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*Token)
	// set default field values here
	log.Printf("Defaulting fields for Token %s\n", obj.Name)
}
