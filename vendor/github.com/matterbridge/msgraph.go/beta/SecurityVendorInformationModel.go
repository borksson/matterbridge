// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SecurityVendorInformation undocumented
type SecurityVendorInformation struct {
	// Object is the base model of SecurityVendorInformation
	Object
	// Provider undocumented
	Provider *string `json:"provider,omitempty"`
	// ProviderVersion undocumented
	ProviderVersion *string `json:"providerVersion,omitempty"`
	// SubProvider undocumented
	SubProvider *string `json:"subProvider,omitempty"`
	// Vendor undocumented
	Vendor *string `json:"vendor,omitempty"`
}