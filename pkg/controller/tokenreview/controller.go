
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


package tokenreview

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/alok87/apik8s/pkg/apis/authentication/v1alpha1"
	"github.com/alok87/apik8s/pkg/controller/sharedinformers"
	listers "github.com/alok87/apik8s/pkg/client/listers_generated/authentication/v1alpha1"
)

// +controller:group=authentication,version=v1alpha1,kind=TokenReview,resource=tokenreviews
type TokenReviewControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about TokenReview
	lister listers.TokenReviewLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *TokenReviewControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing tokenreviews labels
	c.lister = arguments.GetSharedInformers().Factory.Authentication().V1alpha1().TokenReviews().Lister()
}

// Reconcile handles enqueued messages
func (c *TokenReviewControllerImpl) Reconcile(u *v1alpha1.TokenReview) error {
	// Implement controller logic here
	log.Printf("Running reconcile TokenReview for %s\n", u.Name)
	return nil
}

func (c *TokenReviewControllerImpl) Get(namespace, name string) (*v1alpha1.TokenReview, error) {
	return c.lister.TokenReviews(namespace).Get(name)
}
