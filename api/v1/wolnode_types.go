/*
Copyright 2022.

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

package v1

import (
	apiv1 "github.com/aaletov/k8s-wol/api/generated/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// WolNode is the Schema for the wolnodes API
// +kubebuilder:printcolumn:name="MAC",type="string",JSONPath=".spec.MAC",description="Node MAC"
// +kubebuilder:printcolumn:name="IsAlive",type="string",JSONPath=".spec.IsAlive",description="Is node alive"
type WolNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec apiv1.WolNode `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// WolNodeList contains a list of WolNode
type WolNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WolNode `json:"items"`
}

func (in *WolNode) DeepCopyInto(out *WolNode) {
	// We must avoid copying WolNode because it contains MessageState
	// which contains mutex
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec.MAC = in.Spec.MAC
	out.Spec.IsAlive = in.Spec.IsAlive
}

func init() {
	SchemeBuilder.Register(&WolNode{}, &WolNodeList{})
}
