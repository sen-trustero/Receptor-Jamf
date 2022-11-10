/* TODO: Name package */

package receptorPackage_test

import (
	"github.com/stretchr/testify/assert"
	"receptor/trr-jamf/receptorPackage"
	"testing"
)

var demoCredentials = receptorPackage.Credentials{
	Username: "demo",
	Password: "tryitout",
	BaseUrl:  "https://dummy.jamfcloud.com",
}

func TestGetReceptorTypeImpl(t *testing.T) {
	/* TODO: Write tests */
	assert.Equal(t, "trr-custom", receptorPackage.GetReceptorTypeImpl())
}

func TestGetKnownServicesImpl(t *testing.T) {
	/* TODO: Write tests */
	svcs := receptorPackage.GetKnownServicesImpl()
	assert.Len(t, svcs, 1)
	assert.Equal(t, "Custom Service", svcs[0])
}

func TestVerify(t *testing.T) {
	/* TODO: Write tests */
	ok, err := receptorPackage.VerifyImpl(demoCredentials.BaseUrl, demoCredentials.Username, demoCredentials.Password)
	assert.False(t, ok)
	assert.Nil(t, err)
}

func TestDiscover(t *testing.T) {
	/* TODO: Write tests */
	svcs, err := receptorPackage.DiscoverImpl(demoCredentials.BaseUrl, demoCredentials.Username, demoCredentials.Password)
	assert.Len(t, svcs, 1)
	assert.Nil(t, err)
}
func TestReport(t *testing.T) {
	/* TODO: Write tests */
	evs, err := receptorPackage.ReportImpl(demoCredentials.BaseUrl, demoCredentials.Username, demoCredentials.Password)
	assert.Len(t, evs, 0)
	assert.Nil(t, err)
}
