/* TODO: Name package */

package receptorPackage_test

import (
	"receptor/trr-jamf/receptorPackage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReceptorTypeImpl(t *testing.T) {
	/* TODO: Write tests */
	assert.Equal(t, "trr-jamf", receptorPackage.GetReceptorTypeImpl())
}

func TestGetKnownServicesImpl(t *testing.T) {
	/* TODO: Write tests */
	svcs := receptorPackage.GetKnownServicesImpl()
	assert.Len(t, svcs, 1)
	assert.Equal(t, "Jamf Service", svcs[0])
}

func TestVerify(t *testing.T) {
	/* TODO: Write tests */
	ok, err := receptorPackage.VerifyImpl(url, userName, password)
	assert.False(t, ok)
	assert.Nil(t, err)
}

func TestDiscover(t *testing.T) {
	/* TODO: Write tests */
	svcs, err := receptorPackage.DiscoverImpl(url, userName, password)
	assert.Len(t, svcs, 0)
	assert.Nil(t, err)
}
func TestReport(t *testing.T) {
	/* TODO: Write tests */
	evs, err := receptorPackage.ReportImpl(url, userName, password)
	assert.Len(t, evs, 0)
	assert.Nil(t, err)
}
