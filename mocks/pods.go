// Copyright 2016 Mirantis
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mocks

import (
	"strings"

	"k8s.io/client-go/pkg/api/v1"
)

func MakePod(name string) *v1.Pod {
	status := strings.Split(name, "-")[0]
	pod := &v1.Pod{}
	pod.Name = name
	pod.Namespace = "testing"
	if status == "ready" {
		pod.Status.Phase = "Running"
		pod.Status.Conditions = append(
			pod.Status.Conditions,
			v1.PodCondition{Type: "Ready", Status: "True"},
		)
	} else {
		pod.Status.Phase = "Pending"
	}

	return pod
}
