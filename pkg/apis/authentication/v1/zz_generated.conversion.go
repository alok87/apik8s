// +build !ignore_autogenerated

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	authentication "github.com/alok87/apik8s/pkg/apis/authentication"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_Token_To_authentication_Token,
		Convert_authentication_Token_To_v1_Token,
		Convert_v1_TokenList_To_authentication_TokenList,
		Convert_authentication_TokenList_To_v1_TokenList,
		Convert_v1_TokenSpec_To_authentication_TokenSpec,
		Convert_authentication_TokenSpec_To_v1_TokenSpec,
		Convert_v1_TokenStatus_To_authentication_TokenStatus,
		Convert_authentication_TokenStatus_To_v1_TokenStatus,
		Convert_v1_TokenStatusStrategy_To_authentication_TokenStatusStrategy,
		Convert_authentication_TokenStatusStrategy_To_v1_TokenStatusStrategy,
		Convert_v1_TokenStrategy_To_authentication_TokenStrategy,
		Convert_authentication_TokenStrategy_To_v1_TokenStrategy,
	)
}

func autoConvert_v1_Token_To_authentication_Token(in *Token, out *authentication.Token, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_TokenSpec_To_authentication_TokenSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_TokenStatus_To_authentication_TokenStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Token_To_authentication_Token is an autogenerated conversion function.
func Convert_v1_Token_To_authentication_Token(in *Token, out *authentication.Token, s conversion.Scope) error {
	return autoConvert_v1_Token_To_authentication_Token(in, out, s)
}

func autoConvert_authentication_Token_To_v1_Token(in *authentication.Token, out *Token, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authentication_TokenSpec_To_v1_TokenSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authentication_TokenStatus_To_v1_TokenStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_authentication_Token_To_v1_Token is an autogenerated conversion function.
func Convert_authentication_Token_To_v1_Token(in *authentication.Token, out *Token, s conversion.Scope) error {
	return autoConvert_authentication_Token_To_v1_Token(in, out, s)
}

func autoConvert_v1_TokenList_To_authentication_TokenList(in *TokenList, out *authentication.TokenList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]authentication.Token)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_TokenList_To_authentication_TokenList is an autogenerated conversion function.
func Convert_v1_TokenList_To_authentication_TokenList(in *TokenList, out *authentication.TokenList, s conversion.Scope) error {
	return autoConvert_v1_TokenList_To_authentication_TokenList(in, out, s)
}

func autoConvert_authentication_TokenList_To_v1_TokenList(in *authentication.TokenList, out *TokenList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Token)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_authentication_TokenList_To_v1_TokenList is an autogenerated conversion function.
func Convert_authentication_TokenList_To_v1_TokenList(in *authentication.TokenList, out *TokenList, s conversion.Scope) error {
	return autoConvert_authentication_TokenList_To_v1_TokenList(in, out, s)
}

func autoConvert_v1_TokenSpec_To_authentication_TokenSpec(in *TokenSpec, out *authentication.TokenSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1_TokenSpec_To_authentication_TokenSpec is an autogenerated conversion function.
func Convert_v1_TokenSpec_To_authentication_TokenSpec(in *TokenSpec, out *authentication.TokenSpec, s conversion.Scope) error {
	return autoConvert_v1_TokenSpec_To_authentication_TokenSpec(in, out, s)
}

func autoConvert_authentication_TokenSpec_To_v1_TokenSpec(in *authentication.TokenSpec, out *TokenSpec, s conversion.Scope) error {
	return nil
}

// Convert_authentication_TokenSpec_To_v1_TokenSpec is an autogenerated conversion function.
func Convert_authentication_TokenSpec_To_v1_TokenSpec(in *authentication.TokenSpec, out *TokenSpec, s conversion.Scope) error {
	return autoConvert_authentication_TokenSpec_To_v1_TokenSpec(in, out, s)
}

func autoConvert_v1_TokenStatus_To_authentication_TokenStatus(in *TokenStatus, out *authentication.TokenStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1_TokenStatus_To_authentication_TokenStatus is an autogenerated conversion function.
func Convert_v1_TokenStatus_To_authentication_TokenStatus(in *TokenStatus, out *authentication.TokenStatus, s conversion.Scope) error {
	return autoConvert_v1_TokenStatus_To_authentication_TokenStatus(in, out, s)
}

func autoConvert_authentication_TokenStatus_To_v1_TokenStatus(in *authentication.TokenStatus, out *TokenStatus, s conversion.Scope) error {
	return nil
}

// Convert_authentication_TokenStatus_To_v1_TokenStatus is an autogenerated conversion function.
func Convert_authentication_TokenStatus_To_v1_TokenStatus(in *authentication.TokenStatus, out *TokenStatus, s conversion.Scope) error {
	return autoConvert_authentication_TokenStatus_To_v1_TokenStatus(in, out, s)
}

func autoConvert_v1_TokenStatusStrategy_To_authentication_TokenStatusStrategy(in *TokenStatusStrategy, out *authentication.TokenStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1_TokenStatusStrategy_To_authentication_TokenStatusStrategy is an autogenerated conversion function.
func Convert_v1_TokenStatusStrategy_To_authentication_TokenStatusStrategy(in *TokenStatusStrategy, out *authentication.TokenStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1_TokenStatusStrategy_To_authentication_TokenStatusStrategy(in, out, s)
}

func autoConvert_authentication_TokenStatusStrategy_To_v1_TokenStatusStrategy(in *authentication.TokenStatusStrategy, out *TokenStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_authentication_TokenStatusStrategy_To_v1_TokenStatusStrategy is an autogenerated conversion function.
func Convert_authentication_TokenStatusStrategy_To_v1_TokenStatusStrategy(in *authentication.TokenStatusStrategy, out *TokenStatusStrategy, s conversion.Scope) error {
	return autoConvert_authentication_TokenStatusStrategy_To_v1_TokenStatusStrategy(in, out, s)
}

func autoConvert_v1_TokenStrategy_To_authentication_TokenStrategy(in *TokenStrategy, out *authentication.TokenStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1_TokenStrategy_To_authentication_TokenStrategy is an autogenerated conversion function.
func Convert_v1_TokenStrategy_To_authentication_TokenStrategy(in *TokenStrategy, out *authentication.TokenStrategy, s conversion.Scope) error {
	return autoConvert_v1_TokenStrategy_To_authentication_TokenStrategy(in, out, s)
}

func autoConvert_authentication_TokenStrategy_To_v1_TokenStrategy(in *authentication.TokenStrategy, out *TokenStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_authentication_TokenStrategy_To_v1_TokenStrategy is an autogenerated conversion function.
func Convert_authentication_TokenStrategy_To_v1_TokenStrategy(in *authentication.TokenStrategy, out *TokenStrategy, s conversion.Scope) error {
	return autoConvert_authentication_TokenStrategy_To_v1_TokenStrategy(in, out, s)
}
