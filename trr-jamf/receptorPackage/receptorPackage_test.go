/* TODO: Name package */

package receptorPackage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var demoCredentials = Credentials{
	Username: "demo",
	Password: "tryitout",
	BaseUrl:  "https://dummy.jamfcloud.com",
}

func TestGetReceptorTypeImpl(t *testing.T) {
	assert.Equal(t, "trr-custom", GetReceptorTypeImpl())
}

func TestGetKnownServicesImpl(t *testing.T) {
	svcs := GetKnownServicesImpl()
	assert.Len(t, svcs, 1)
	assert.Equal(t, "Custom Service", svcs[0])
}

func TestVerify(t *testing.T) {
	ok, err := VerifyImpl(demoCredentials.BaseUrl, demoCredentials.Username, demoCredentials.Password)
	assert.False(t, ok)
	assert.Nil(t, err)
}

func TestDiscover(t *testing.T) {
	svcs, err := DiscoverImpl(demoCredentials.BaseUrl, demoCredentials.Username, demoCredentials.Password)
	assert.Len(t, svcs, 1)
	assert.Nil(t, err)
}
func TestReport(t *testing.T) {
	evs, err := ReportImpl(demoCredentials.BaseUrl, demoCredentials.Username, demoCredentials.Password)
	assert.Nil(t, err)
	assert.Len(t, evs, 1)
}
