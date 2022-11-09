/* TODO: Name package */
package receptorPackage

import (
	receptorLog "receptor/trr-jamf/logging"

	"github.com/trustero/api/go/receptor_sdk"
	"github.com/trustero/api/go/receptor_v1"
	"github.com/trustero/jamf-api-client-go/classic/computers"
)

const (
	receptorName = "trr-custom"
	serviceId    = "trustero_jamf"
	serviceName  = "Custom Name"
)

func GetReceptorTypeImpl() string {
	return receptorName
}

func GetKnownServicesImpl() []string {
	return []string{serviceName}
}

func VerifyImpl(base_Url string, userName string, password string) (ok bool, err error) {
	receptorLog.Info("Entering VerifyImpl")
	ok = true
	_, err = computers.NewService(base_Url, userName, password, nil)
	if err != nil {
		receptorLog.Err(err, "Could not verify, error in Jamf Computers NewService for %s ", userName)
		return false, nil
	}
	receptorLog.Info("Leaving VerifyImpl")
	return ok, err
}

func DiscoverImpl(base_Url string, userName string, password string) (svcs []*receptor_v1.ServiceEntity, err error) {
	receptorLog.Info("Entering DiscoverImpl")
	services := receptor_sdk.NewServiceEntities()
	services.ServiceProviderAccount = "jamf"
	services.AddService(serviceName, serviceId, "Computer List", "1")
	svcs = services.Entities
	receptorLog.Info("Leaving DiscoverImpl")
	return

}

func ReportImpl(base_Url string, userName string, password string) (evidences []*receptor_sdk.Evidence, err error) {
	receptorLog.Info("Entering ReportImpl")
	report := receptor_sdk.NewReport()

	var computersService *computers.Service
	computersService, err = computers.NewService(base_Url, userName, password, nil)

	if err != nil {
		receptorLog.Err(err, "could not discover, error in %s", serviceName)

	}
	var evidence *receptor_sdk.Evidence
	if evidence, err = getComputerEvidence(computersService); err != nil {
		receptorLog.Err(err, "could not discover, error in %s", serviceName)
		return

	}
	report.AddEvidence(evidence)
	receptorLog.Info("Leaving ReportImpl")
	return report.Evidences, err
}
