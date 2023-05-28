/*
Copyright 2020 The OpenEBS Authors.

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

package usage

import (
	k8sapi "github.com/openebs/lib-csi/pkg/client/k8s"
	"github.com/openebs/lib-csi/pkg/common/env"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"

	openebsversion "github.com/openebs/lvm-localpv/pkg/version"
)

var (
	clusterUUID    = "OPENEBS_IO_USAGE_UUID"
	clusterVersion = "OPENEBS_IO_K8S_VERSION"
	clusterArch    = "OPENEBS_IO_K8S_ARCH"
	openEBSversion = "OPENEBS_IO_VERSION_TAG"
	nodeType       = "OPENEBS_IO_NODE_TYPE"
	installerType  = "OPENEBS_IO_INSTALLER_TYPE"
)

// VersionSet is a struct which stores (sort of) fixed information about a
// k8s environment
type VersionSet struct {
	id             string // OPENEBS_IO_USAGE_UUID
	k8sVersion     string // OPENEBS_IO_K8S_VERSION
	k8sArch        string // OPENEBS_IO_K8S_ARCH
	openebsVersion string // OPENEBS_IO_VERSION_TAG
	nodeType       string // OPENEBS_IO_NODE_TYPE
	installerType  string // OPENEBS_IO_INSTALLER_TYPE
}

// NewVersion returns a new versionSet struct
func NewVersion() *VersionSet {
	return &VersionSet{}
}

// fetchAndSetVersion consumes the Kubernetes API to get environment constants
// and returns a versionSet struct
func (v *VersionSet) fetchAndSetVersion() error {
	var err error
	v.id, err = getUUIDbyNS("default")
	if err != nil {
		return err
	}
	_ = env.Set(clusterUUID, v.id)

	k8s, err := k8sapi.GetServerVersion()
	if err != nil {
		return err
	}
	// eg. linux/amd64
	v.k8sArch = k8s.Platform
	v.k8sVersion = k8s.GitVersion
	// Explicitly informing linters that we intended to avoid errors(errcheck)
	_ = env.Set(clusterArch, v.k8sArch)
	_ = env.Set(clusterVersion, v.k8sVersion)
	v.nodeType, err = k8sapi.GetOSAndKernelVersion()
	_ = env.Set(nodeType, v.nodeType)
	if err != nil {
		return err
	}
	v.openebsVersion = openebsversion.GetVersionDetails()
	_ = env.Set(openEBSversion, v.openebsVersion)
	return nil
}

// getVersion is a wrapper over fetchAndSetVersion
func (v *VersionSet) getVersion(override bool) error {
	// If ENVs aren't set or the override is true, fetch the required
	// values from the K8s APIserver
	if _, present := env.Lookup(openEBSversion); !present || override {
		if err := v.fetchAndSetVersion(); err != nil {
			klog.Error(err.Error())
			return err
		}
	}
	// Fetch data from ENV
	v.id = env.Get(clusterUUID)
	v.k8sArch = env.Get(clusterArch)
	v.k8sVersion = env.Get(clusterVersion)
	v.nodeType = env.Get(nodeType)
	v.openebsVersion = env.Get(openEBSversion)
	v.installerType = env.Get(installerType)
	return nil
}

// getUUIDbyNS returns the metadata.object.uid of a namespace in Kubernetes
func getUUIDbyNS(namespace string) (string, error) {
	ns := k8sapi.Namespace()
	NSstruct, err := ns.Get(namespace, metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	if NSstruct != nil {
		return string(NSstruct.GetObjectMeta().GetUID()), nil
	}
	return "", nil
}
