/* TODO: Name package */
package receptorPackage

import (
	receptorLog "receptor/trr-jamf/logging"
	//"receptor/trr-jamf/receptorPackage"
	"strconv"

	"github.com/trustero/api/go/receptor_sdk"
	"github.com/trustero/api/go/receptor_v1"
	"github.com/trustero/jamf-api-client-go/classic/computers"
)

const (
	receptorName = "trr-jamf"
	serviceName  = "trustero_jamf"
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

	var computersService *computers.Service
	computersService, err = computers.NewService(base_Url, userName, password, nil)

	if err != nil {
		receptorLog.Err(err, "could not discover, error in %s", serviceName)

	}

	var apicalls []string
	discovered := Discovered{}
	discovered.Computers, apicalls, _ = findComputers(computersService)
	for _, s := range apicalls {
		receptorLog.Info(s)

	}
	receptorLog.Info(">>>>>>>>>")
	for _, c := range discovered.Computers {
		receptorLog.Info("%s | %d ", c.General.Name, c.General.Id)

	}
	receptorLog.Info(">>>>>>>>> count %d : %d", len(discovered.Computers), len(apicalls))

	// Example: Get All Computers
	var comps []computers.ComputerNameId
	comps, _, err = computersService.List()
	if err != nil {
		return
	}
	for comp := range comps {
		receptorLog.Info(strconv.Itoa(comp))
	}
	receptorLog.Info("Leaving DiscoverImpl")
	return

}

func ReportImpl(base_Url string, userName string, password string) (evidences []*receptor_sdk.Evidence, err error) {
	receptorLog.Info("Entering ReportImpl")
	/* TODO: Implement Report logic here */
	receptorLog.Info("Leaving ReportImpl")
	return
}
