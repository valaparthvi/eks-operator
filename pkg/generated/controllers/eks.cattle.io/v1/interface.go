/*
Copyright 2019 Wrangler Sample Controller Authors

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	v1 "github.com/rancher/eks-operator/pkg/apis/eks.cattle.io/v1"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	v1.AddToScheme(schemes.All)
}

type Interface interface {
	EKSClusterConfig() EKSClusterConfigController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (c *version) EKSClusterConfig() EKSClusterConfigController {
	return NewEKSClusterConfigController(schema.GroupVersionKind{Group: "eks.cattle.io", Version: "v1", Kind: "EKSClusterConfig"}, "eksclusterconfigs", true, c.controllerFactory)
}
