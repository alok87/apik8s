
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


package token

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/alok87/apik8s/pkg/apis/authentication/v1"
	"github.com/alok87/apik8s/pkg/controller/sharedinformers"
	listers "github.com/alok87/apik8s/pkg/client/listers_generated/authentication/v1"
)

// +controller:group=authentication,version=v1,kind=Token,resource=tokens
type TokenControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Token
	lister listers.TokenLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *TokenControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing tokens labels
	c.lister = arguments.GetSharedInformers().Factory.Authentication().V1().Tokens().Lister()
}

// Reconcile handles enqueued messages
func (c *TokenControllerImpl) Reconcile(u *v1.Token) error {
	// Implement controller logic here
	log.Printf("Running reconcile Token for %s\n", u.Name)
	return nil
}

func (c *TokenControllerImpl) Get(namespace, name string) (*v1.Token, error) {
	return c.lister.Tokens(namespace).Get(name)
}
