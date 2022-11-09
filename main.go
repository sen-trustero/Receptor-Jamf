package main

import (
	"github.com/trustero/api/go/receptor_sdk"
	"github.com/trustero/api/go/receptor_sdk/cmd"
	"github.com/trustero/api/go/receptor_v1"
	receptorPackage "receptor/trr-jamf/receptorPackage"
)

type Receptor struct {
	UserName string `trustero:"display:Jamf Username;placeholder:userName"`
	Password string `trustero:"display:Jamf Password;placeholder:password"`
	BaseUrl  string `trustero:"display:Jamf Base Url;placeholder:base_Url"`
}

func (r *Receptor) GetReceptorType() string {
	return receptorPackage.GetReceptorTypeImpl()
}

func (r *Receptor) GetKnownServices() []string {
	return receptorPackage.GetKnownServicesImpl()
}

// This will return Receptor struct defined above when the receptor is asked to
// identify itself
func (r *Receptor) GetCredentialObj() (credentialObj interface{}) {
	return r
}

// This function will call into the service provider API with the provided
// credentials and confirm that the credentials are valid. Usually a simple
// API call like GET org name. If the credentials are not valid,
// return a relevant error message
func (r *Receptor) Verify(credentials interface{}) (ok bool, err error) {
	c := credentials.(*Receptor)
	/* TODO: Change crediential field names */
	return receptorPackage.VerifyImpl(c.BaseUrl, c.UserName, c.Password)
}

// The Discover function returns a list of Service Entities. This function
// makes any relevant API calls to the Service Provider to gather information
// about how many Service Entity Instances are in use. If at any point this
// function runs into an error, log that error and continue
func (r *Receptor) Discover(credentials interface{}) (svcs []*receptor_v1.ServiceEntity, err error) {
	c := credentials.(*Receptor)
	/* TODO: Change crediential field names */
	return receptorPackage.DiscoverImpl(c.BaseUrl, c.UserName, c.Password)
}

// Report will often make the same API calls made in the Discover call, but it
// will additionally create evidences with the data returned from the API calls
func (r *Receptor) Report(credentials interface{}) (evidences []*receptor_sdk.Evidence, err error) {
	c := credentials.(*Receptor)
	/* TODO: Change crediential field names */
	return receptorPackage.ReportImpl(c.BaseUrl, c.UserName, c.Password)
}

func main() {

	cmd.Execute(&Receptor{})
}
