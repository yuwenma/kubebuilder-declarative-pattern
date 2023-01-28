#!/bin/bash

# Copyright 2019 The Kubernetes Authors.
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

CI_ROOT=$(dirname "${BASH_SOURCE}")
REPO_ROOT=$(dirname "${BASH_SOURCE}")/../..

source "${CI_ROOT}/fetch_kubebuilder_release_bin.sh"

cd "${REPO_ROOT}"

export GO111MODULE=on

go test github.com/yuwenma/kubebuilder-declarative-pattern/pkg/...

cd examples/guestbook-operator
go test github.com/yuwenma/kubebuilder-declarative-pattern/examples/guestbook-operator/controllers/...
