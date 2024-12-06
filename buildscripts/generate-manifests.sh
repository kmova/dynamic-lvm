#!/usr/bin/env bash

# Copyright 2020 The OpenEBS Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

## find or download controller-gen
CONTROLLER_GEN=$(which controller-gen)

if [ "$CONTROLLER_GEN" = "" ]
then
  echo "ERROR: failed to get controller-gen, Please run make bootstrap to install it";
  exit 1;
fi

$CONTROLLER_GEN crd:trivialVersions=false,preserveUnknownFields=false paths=./pkg/apis/... output:crd:artifacts:config=deploy/yamls

## create the the crd yamls

echo '

##############################################
###########                       ############
###########   LVMVolume CRD       ############
###########                       ############
##############################################

# LVMVolume CRD is autogenerated via `make manifests` command.
# Do the modification in the code and run the `make manifests` command
# to generate the CRD definition' > deploy/yamls/lvmvolume-crd.yaml

cat deploy/yamls/local.openebs.io_lvmvolumes.yaml >> deploy/yamls/lvmvolume-crd.yaml
rm deploy/yamls/local.openebs.io_lvmvolumes.yaml

echo '

##############################################
###########                       ############
###########   LVMSnapshot CRD     ############
###########                       ############
##############################################

# LVMSnapshot CRD is autogenerated via `make manifests` command.
# Do the modification in the code and run the `make manifests` command
# to generate the CRD definition' > deploy/yamls/lvmsnapshot-crd.yaml

cat deploy/yamls/local.openebs.io_lvmsnapshots.yaml >> deploy/yamls/lvmsnapshot-crd.yaml
rm deploy/yamls/local.openebs.io_lvmsnapshots.yaml

echo '

##############################################
###########                       ############
###########     LVMNode CRD       ############
###########                       ############
##############################################

# LVMNode CRD is autogenerated via `make manifests` command.
# Do the modification in the code and run the `make manifests` command
# to generate the CRD definition' > deploy/yamls/lvmnode-crd.yaml

cat deploy/yamls/local.openebs.io_lvmnodes.yaml >> deploy/yamls/lvmnode-crd.yaml
rm deploy/yamls/local.openebs.io_lvmnodes.yaml

## create the operator file using all the yamls

echo '# This manifest is autogenerated via `make manifests` command
# Do the modification to the lvm-driver.yaml in directory deploy/yamls/
# and then run `make manifests` command

# This manifest deploys the OpenEBS LVM control plane components,
# with associated CRs & RBAC rules.
' > deploy/lvm-operator.yaml

# Add namespace creation to the Operator yaml
cat deploy/yamls/namespace.yaml >> deploy/lvm-operator.yaml

# Add LVMVolume v1alpha1 CRDs to the Operator yaml
cat deploy/yamls/lvmvolume-crd.yaml >> deploy/lvm-operator.yaml

# Add LVMSnapshot v1alpha1 CRDs to the Operator yaml
cat deploy/yamls/lvmsnapshot-crd.yaml >> deploy/lvm-operator.yaml

# Add LVMNode v1alpha1 CRDs to the Operator yaml
cat deploy/yamls/lvmnode-crd.yaml >> deploy/lvm-operator.yaml

# Add the driver deployment to the Operator yaml
cat deploy/yamls/lvm-driver.yaml >> deploy/lvm-operator.yaml

# To use your own boilerplate text use:
#   --go-header-file ${SCRIPT_ROOT}/hack/custom-boilerplate.go.txt
