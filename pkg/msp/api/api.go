/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: [Default license](LICENSE)
*/

package api

import (
	"errors"
)

var (
	// ErrCARegistrarNotFound indicates the CA registrar was not found
	ErrCARegistrarNotFound = errors.New("CA registrar not found")
)

// CAClient provides management of identities in a Fabric network
type CAClient interface {
	Enroll(request *EnrollmentRequest) error
	Reenroll(request *ReenrollmentRequest) error
	Register(request *RegistrationRequest) (string, error)
	Revoke(request *RevocationRequest) (*RevocationResponse, error)
	GetCAInfo() (*GetCAInfoResponse, error)
	CreateIdentity(request *IdentityRequest) (*IdentityResponse, error)
	GetIdentity(id, caname string) (*IdentityResponse, error)
	ModifyIdentity(request *IdentityRequest) (*IdentityResponse, error)
	RemoveIdentity(request *RemoveIdentityRequest) (*IdentityResponse, error)
	GetAllIdentities(caname string) ([]*IdentityResponse, error)
	GetAffiliation(affiliation, caname string) (*AffiliationResponse, error)
	GetAllAffiliations(caname string) (*AffiliationResponse, error)
	AddAffiliation(request *AffiliationRequest) (*AffiliationResponse, error)
	ModifyAffiliation(request *ModifyAffiliationRequest) (*AffiliationResponse, error)
	RemoveAffiliation(request *AffiliationRequest) (*AffiliationResponse, error)
}

// AttributeRequest is a request for an attribute.
type AttributeRequest struct {
	Name     string
	Optional bool
}

// RegistrationRequest defines the attributes required to register a user with the CA
type RegistrationRequest struct {
	// Name is the unique name of the identity
	Name string
	// Type of identity being registered (e.g. "peer, app, user")
	Type string
	// MaxEnrollments is the number of times the secret can  be reused to enroll.
	// if omitted, this defaults to max_enrollments configured on the server
	MaxEnrollments int
	// The identity's affiliation e.g. org1.department1
	Affiliation string
	// Optional attributes associated with this identity
	Attributes []Attribute
	// CAName is the name of the CA to connect to
	CAName string
	// Secret is an optional password.  If not specified,
	// a random secret is generated.  In both cases, the secret
	// is returned from registration.
	Secret string
}

// EnrollmentRequest is a request to enroll an identity
type EnrollmentRequest struct {
	// The identity name to enroll
	Name string
	// The secret returned via Register
	Secret string
	// CAName is the name of the CA to connect to
	CAName string
	// AttrReqs are requests for attributes to add to the certificate.
	// Each attribute is added only if the requestor owns the attribute.
	AttrReqs []*AttributeRequest
	// Profile is the name of the signing profile to use in issuing the X509 certificate
	Profile string
	// Label is the label to use in HSM operations
	Label string
	// The type of the enrollment request: x509 or idemix
	// The default is a request for an X509 enrollment certificate
	Type string
	// CSR is Certificate Signing Request info
	CSR *CSRInfo
}

// CSRInfo is Certificate Signing Request (CSR) Information
type CSRInfo struct {
	CN    string
	Hosts []string
}

// ReenrollmentRequest is a request to reenroll an identity.
// This is useful to renew a certificate before it has expired.
type ReenrollmentRequest struct {
	// The identity name to enroll
	Name string
	// Profile is the name of the signing profile to use in issuing the certificate
	Profile string
	// Label is the label to use in HSM operations
	Label string
	// CAName is the name of the CA to connect to
	CAName string
	// AttrReqs are requests for attributes to add to the certificate.
	// Each attribute is added only if the requestor owns the attribute.
	AttrReqs []*AttributeRequest
	// CSR is Certificate Signing Request info
	CSR *CSRInfo
}

// Attribute defines additional attributes that may be passed along during registration
type Attribute struct {
	Name  string
	Value string
	ECert bool
}

// RevocationRequest defines the attributes required to revoke credentials with the CA
type RevocationRequest struct {
	// Name of the identity whose certificates should be revoked
	// If this field is omitted, then Serial and AKI must be specified.
	Name string
	// Serial number of the certificate to be revoked
	// If this is omitted, then Name must be specified
	Serial string
	// AKI (Authority Key Identifier) of the certificate to be revoked
	AKI string
	// Reason is the reason for revocation. See https://godoc.org/golang.org/x/crypto/ocsp
	// for valid values. The default value is 0 (ocsp.Unspecified).
	Reason string
	// CAName is the name of the CA to connect to
	CAName string
	// GenCRL specifies whether to generate a CRL
	GenCRL bool
}

// RevocationResponse represents response from the server for a revocation request
type RevocationResponse struct {
	// RevokedCerts is an array of certificates that were revoked
	RevokedCerts []RevokedCert
	// CRL is PEM-encoded certificate revocation list (CRL) that contains all unexpired revoked certificates
	CRL []byte
}

// RevokedCert represents a revoked certificate
type RevokedCert struct {
	// Serial number of the revoked certificate
	Serial string
	// AKI of the revoked certificate
	AKI string
}

// IdentityRequest represents the request to add/update identity to the fabric-ca-server
type IdentityRequest struct {

	// The enrollment ID which uniquely identifies an identity (required)
	ID string

	// The identity's affiliation
	Affiliation string

	// Array of attributes to assign to the user
	Attributes []Attribute

	// Type of identity being registered (e.g. 'peer, app, user'). Default is 'user'.
	Type string

	// The maximum number of times the secret can be reused to enroll (default CA's Max Enrollment)
	MaxEnrollments int

	// The enrollment secret. If not provided, a random secret is generated.
	Secret string

	// Name of the CA to send the request to within the Fabric CA server (optional)
	CAName string
}

// RemoveIdentityRequest represents the request to remove an existing identity from the
// fabric-ca-server
type RemoveIdentityRequest struct {

	// The enrollment ID which uniquely identifies an identity
	ID string

	// Force delete
	Force bool

	// Name of the CA
	CAName string
}

// IdentityResponse is the response from the any read/add/modify/remove identity call
type IdentityResponse struct {

	// The enrollment ID which uniquely identifies an identity
	ID string

	// The identity's affiliation
	Affiliation string

	// Array of attributes assigned to the user
	Attributes []Attribute

	// Type of identity (e.g. 'peer, app, user')
	Type string

	// The maximum number of times the secret can be reused to enroll
	MaxEnrollments int

	// The enrollment secret
	Secret string

	// Name of the CA
	CAName string
}

// AffiliationRequest represents the request to add/remove affiliation to the fabric-ca-server
type AffiliationRequest struct {
	// Name of the affiliation
	Name string

	// Creates parent affiliations if they do not exist
	Force bool

	// Name of the CA
	CAName string
}

// ModifyAffiliationRequest represents the request to modify an existing affiliation on the
// fabric-ca-server
type ModifyAffiliationRequest struct {
	AffiliationRequest
	// New name of the affiliation
	NewName string
}

// AffiliationResponse contains the response for get, add, modify, and remove an affiliation
type AffiliationResponse struct {
	AffiliationInfo
	CAName string
}

// AffiliationInfo contains the affiliation name, child affiliation info, and identities
// associated with this affiliation.
type AffiliationInfo struct {
	Name         string
	Affiliations []AffiliationInfo
	Identities   []IdentityInfo
}

// IdentityInfo contains information about an identity
type IdentityInfo struct {
	ID             string
	Type           string
	Affiliation    string
	Attributes     []Attribute
	MaxEnrollments int
}

// GetCAInfoResponse is the response from the GetCAInfo call
type GetCAInfoResponse struct {
	// CAName is the name of the CA
	CAName string
	// CAChain is the PEM-encoded bytes of the fabric-ca-server's CA chain.
	// The 1st element of the chain is the root CA cert
	CAChain []byte
	// Idemix issuer public key of the CA
	IssuerPublicKey []byte
	// Idemix issuer revocation public key of the CA
	IssuerRevocationPublicKey []byte
	// Version of the server
	Version string
}
