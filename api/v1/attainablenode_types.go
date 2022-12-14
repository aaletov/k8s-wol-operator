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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AttainableNodeSpec defines the desired state of AttainableNode
// Name of this object should contain name of the host so it could be
// obtained using WakeNodeUpRequest information
type AttainableNodeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Should be label so AttainableNodeReonciler could list all it's nodes
	// NodeId string `protobuf:"bytes,1,opt,name=NodeId,proto3" json:"NodeId,omitempty"` // Node which can wake up this node
	MAC string `protobuf:"bytes,2,opt,name=MAC,proto3" json:"MAC,omitempty"`
}

// Strongly AFTER Spec
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AttainableNode is the Schema for the attainablenodes API
type AttainableNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AttainableNodeSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// AttainableNodeList contains a list of AttainableNode
type AttainableNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AttainableNode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AttainableNode{}, &AttainableNodeList{})
}
