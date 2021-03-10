/*
Copyright 2021 The OpenEBS Authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openebs/lvm-localpv/pkg/apis/openebs.io/lvm/v1alpha1"
	"github.com/openebs/lvm-localpv/pkg/generated/clientset/internalclientset/scheme"
	rest "k8s.io/client-go/rest"
)

type LocalV1alpha1Interface interface {
	RESTClient() rest.Interface
	LVMNodesGetter
	LVMSnapshotsGetter
	LVMVolumesGetter
}

// LocalV1alpha1Client is used to interact with features provided by the local.openebs.io group.
type LocalV1alpha1Client struct {
	restClient rest.Interface
}

func (c *LocalV1alpha1Client) LVMNodes(namespace string) LVMNodeInterface {
	return newLVMNodes(c, namespace)
}

func (c *LocalV1alpha1Client) LVMSnapshots(namespace string) LVMSnapshotInterface {
	return newLVMSnapshots(c, namespace)
}

func (c *LocalV1alpha1Client) LVMVolumes(namespace string) LVMVolumeInterface {
	return newLVMVolumes(c, namespace)
}

// NewForConfig creates a new LocalV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*LocalV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &LocalV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new LocalV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *LocalV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new LocalV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *LocalV1alpha1Client {
	return &LocalV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *LocalV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
