package v1

import (
	"google.golang.org/grpc/codes"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LogicalVolumeSpec defines the desired state of LogicalVolume
type LogicalVolumeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Name        string            `json:"name"`
	NodeName    string            `json:"nodeName"`
	Size        resource.Quantity `json:"size"`
	DeviceClass string            `json:"deviceClass,omitempty"`

	// 'type' specifies the logical volume type.
	// Can be 'thin-snapshot' or left empty in case of a regular volume.
	// +kubebuilder:validation:Optional
	Type string `json:"type"`

	// 'sourceID' specifies the volumeID of the parent logical volume; if present.
	// +kubebuilder:validation:Optional
	SourceID string `json:"sourceID"`

	//'accessType' specifies how the user intends to consume the snapshot logical volume.
	// Allowed values are: "ro" (read-only) and "rw" (read-write).
	// Since the kubernetes snapshots are Read-Only, set accessType as 'ro' to activate thin-snapshots as read-only volumes.
	// On the other hand, set accessType as 'rw' to activate thin-snapshots as read-write volumes.
	// +kubebuilder:validation:Optional
	AccessType string `json:"accessType"`
}

// LogicalVolumeStatus defines the observed state of LogicalVolume
type LogicalVolumeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	VolumeID    string             `json:"volumeID,omitempty"`
	Code        codes.Code         `json:"code,omitempty"`
	Message     string             `json:"message,omitempty"`
	CurrentSize *resource.Quantity `json:"currentSize,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// LogicalVolume is the Schema for the logicalvolumes API
type LogicalVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LogicalVolumeSpec   `json:"spec,omitempty"`
	Status LogicalVolumeStatus `json:"status,omitempty"`
}

// IsCompatibleWith returns true if the LogicalVolume is compatible.
func (lv *LogicalVolume) IsCompatibleWith(lv2 *LogicalVolume) bool {
	if lv.Spec.Name != lv2.Spec.Name {
		return false
	}
	if lv.Spec.SourceID != lv2.Spec.SourceID {
		return false
	}
	if lv.Spec.Size.Cmp(lv2.Spec.Size) != 0 {
		return false
	}
	return true
}

//+kubebuilder:object:root=true

// LogicalVolumeList contains a list of LogicalVolume
type LogicalVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogicalVolume `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LogicalVolume{}, &LogicalVolumeList{})
}
