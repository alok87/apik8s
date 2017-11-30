
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


package mykind

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/alok87/apik8s/pkg/apis/group/v1"
	"github.com/alok87/apik8s/pkg/controller/sharedinformers"
	listers "github.com/alok87/apik8s/pkg/client/listers_generated/group/v1"
)

// +controller:group=group,version=v1,kind=Mykind,resource=mykinds
type MykindControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Mykind
	lister listers.MykindLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *MykindControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing mykinds labels
	c.lister = arguments.GetSharedInformers().Factory.Group().V1().Mykinds().Lister()
}

// Reconcile handles enqueued messages
func (c *MykindControllerImpl) Reconcile(u *v1.Mykind) error {
	// Implement controller logic here
	log.Printf("Running reconcile Mykind for %s\n", u.Name)
	return nil
}

func (c *MykindControllerImpl) Get(namespace, name string) (*v1.Mykind, error) {
	return c.lister.Mykinds(namespace).Get(name)
}
