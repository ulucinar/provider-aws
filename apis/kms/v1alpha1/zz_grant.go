/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GrantParameters defines the desired state of Grant
type GrantParameters struct {
	// Region is which region the Grant will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// Specifies a grant constraint.
	//
	// Do not include confidential or sensitive information in this field. This
	// field may be displayed in plaintext in CloudTrail logs and other output.
	//
	// KMS supports the EncryptionContextEquals and EncryptionContextSubset grant
	// constraints, which allow the permissions in the grant only when the encryption
	// context in the request matches (EncryptionContextEquals) or includes (EncryptionContextSubset)
	// the encryption context specified in the constraint.
	//
	// The encryption context grant constraints are supported only on grant operations
	// (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-grant-operations)
	// that include an EncryptionContext parameter, such as cryptographic operations
	// on symmetric encryption KMS keys. Grants with grant constraints can include
	// the DescribeKey and RetireGrant operations, but the constraint doesn't apply
	// to these operations. If a grant with a grant constraint includes the CreateGrant
	// operation, the constraint requires that any grants created with the CreateGrant
	// permission have an equally strict or stricter encryption context constraint.
	//
	// You cannot use an encryption context grant constraint for cryptographic operations
	// with asymmetric KMS keys or HMAC KMS keys. Operations with these keys don't
	// support an encryption context.
	//
	// Each constraint value can include up to 8 encryption context pairs. The encryption
	// context value in each constraint cannot exceed 384 characters. For information
	// about grant constraints, see Using grant constraints (https://docs.aws.amazon.com/kms/latest/developerguide/create-grant-overview.html#grant-constraints)
	// in the Key Management Service Developer Guide. For more information about
	// encryption context, see Encryption context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the Key Management Service Developer Guide .
	Constraints *GrantConstraints `json:"constraints,omitempty"`
	// A list of grant tokens.
	//
	// Use a grant token when your permission to call this operation comes from
	// a new grant that has not yet achieved eventual consistency. For more information,
	// see Grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token)
	// and Using a grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token)
	// in the Key Management Service Developer Guide.
	GrantTokens []*string `json:"grantTokens,omitempty"`
	// The identity that gets the permissions specified in the grant.
	//
	// To specify the grantee principal, use the Amazon Resource Name (ARN) of an
	// Amazon Web Services principal. Valid principals include Amazon Web Services
	// accounts, IAM users, IAM roles, federated users, and assumed role users.
	// For help with the ARN syntax for a principal, see IAM ARNs (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns)
	// in the Identity and Access Management User Guide .
	// +kubebuilder:validation:Required
	GranteePrincipal *string `json:"granteePrincipal"`
	// A friendly name for the grant. Use this value to prevent the unintended creation
	// of duplicate grants when retrying this request.
	//
	// Do not include confidential or sensitive information in this field. This
	// field may be displayed in plaintext in CloudTrail logs and other output.
	//
	// When this value is absent, all CreateGrant requests result in a new grant
	// with a unique GrantId even if all the supplied parameters are identical.
	// This can result in unintended duplicates when you retry the CreateGrant request.
	//
	// When this value is present, you can retry a CreateGrant request with identical
	// parameters; if the grant already exists, the original GrantId is returned
	// without creating a new grant. Note that the returned grant token is unique
	// with every CreateGrant request, even when a duplicate GrantId is returned.
	// All grant tokens for the same grant ID can be used interchangeably.
	Name *string `json:"name,omitempty"`
	// A list of operations that the grant permits.
	//
	// This list must include only operations that are permitted in a grant. Also,
	// the operation must be supported on the KMS key. For example, you cannot create
	// a grant for a symmetric encryption KMS key that allows the Sign operation,
	// or a grant for an asymmetric KMS key that allows the GenerateDataKey operation.
	// If you try, KMS returns a ValidationError exception. For details, see Grant
	// operations (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-grant-operations)
	// in the Key Management Service Developer Guide.
	// +kubebuilder:validation:Required
	Operations []*string `json:"operations"`
	// The principal that has permission to use the RetireGrant operation to retire
	// the grant.
	//
	// To specify the principal, use the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of an Amazon Web Services principal. Valid principals include Amazon Web
	// Services accounts, IAM users, IAM roles, federated users, and assumed role
	// users. For help with the ARN syntax for a principal, see IAM ARNs (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns)
	// in the Identity and Access Management User Guide .
	//
	// The grant determines the retiring principal. Other principals might have
	// permission to retire the grant or revoke the grant. For details, see RevokeGrant
	// and Retiring and revoking grants (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#grant-delete)
	// in the Key Management Service Developer Guide.
	RetiringPrincipal     *string `json:"retiringPrincipal,omitempty"`
	CustomGrantParameters `json:",inline"`
}

// GrantSpec defines the desired state of Grant
type GrantSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GrantParameters `json:"forProvider"`
}

// GrantObservation defines the observed state of Grant
type GrantObservation struct {
	// The unique identifier for the grant.
	//
	// You can use the GrantId in a ListGrants, RetireGrant, or RevokeGrant operation.
	GrantID *string `json:"grantID,omitempty"`
	// The grant token.
	//
	// Use a grant token when your permission to call this operation comes from
	// a new grant that has not yet achieved eventual consistency. For more information,
	// see Grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token)
	// and Using a grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token)
	// in the Key Management Service Developer Guide.
	GrantToken *string `json:"grantToken,omitempty"`
}

// GrantStatus defines the observed state of Grant.
type GrantStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GrantObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Grant is the Schema for the Grants API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Grant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GrantSpec   `json:"spec"`
	Status            GrantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GrantList contains a list of Grants
type GrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Grant `json:"items"`
}

// Repository type metadata.
var (
	GrantKind             = "Grant"
	GrantGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GrantKind}.String()
	GrantKindAPIVersion   = GrantKind + "." + GroupVersion.String()
	GrantGroupVersionKind = GroupVersion.WithKind(GrantKind)
)

func init() {
	SchemeBuilder.Register(&Grant{}, &GrantList{})
}