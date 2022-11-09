package receptorPackage

import (
	receptorLog "receptor/trr-jamf/logging"
	"github.com/trustero/api/go/receptor_sdk"
	"github.com/trustero/api/go/receptor_v1"
	"github.com/trustero/jamf-api-client-go/classic/computers"
)

const (
	receptorName = "trr-custom"
	serviceId    = "Custom Service"
	serviceName  = "Custom Service"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	BaseUrl  string `json:"baseurl"`
}



func GetReceptorTypeImpl() string {
	return receptorName
}

func GetKnownServicesImpl() []string {
	return []string{serviceName}
}

func VerifyImpl(baseurl string, username string, password string) (ok bool, err error) {
	receptorLog.Info("Entering VerifyImpl")
	ok = true
	_, err = computers.NewService(baseurl, username, password, nil)
	if err != nil {
		receptorLog.Err(err, "Could not verify, error in Jamf Computers NewService for %s ", username)
		return false, nil
	}
	receptorLog.Info("Leaving VerifyImpl")
	return ok, err
}

func DiscoverImpl(baseUrl string, username string, password string) (svcs []*receptor_v1.ServiceEntity, err error) {
	receptorLog.Info("Entering DiscoverImpl")
	services := receptor_sdk.NewServiceEntities()
	services.AddService(serviceName, serviceId, "ComputerList", "1")
	receptorLog.Info("Leaving DiscoverImpl")
	return services.Entities, err
}

func ReportImpl(baseurl string, username string, password string) (evidences []*receptor_sdk.Evidence, err error) {
	receptorLog.Info("Entering ReportImpl")
	report := receptor_sdk.NewReport()
	var computersService *computers.Service
	computersService, err = computers.NewService(baseurl, username, password, nil)
	if err != nil {
		receptorLog.Err(err, "could not discover, error in %s", serviceName)

	}
	var evidence *receptor_sdk.Evidence
	if evidence, err = getComputerEvidence(computersService); err != nil {
		receptorLog.Err(err, "could not generate evidence, error in getComputerEvidence")
		return
	}
	report.AddEvidence(evidence)
	receptorLog.Info("Leaving ReportImpl")
	return report.Evidences, err
}
