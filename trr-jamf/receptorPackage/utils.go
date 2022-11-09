/* TODO: Name package */
package receptorPackage

import (
	"fmt"
	receptorLog "receptor/trr-jamf/logging"

	"github.com/trustero/api/go/receptor_sdk"
	computers "github.com/trustero/jamf-api-client-go/classic/computers"
)

type computerResponse struct {
	computer *computers.Computer
	//apiCalls []string
}

type JamfComputerInfo struct {
	Username        string `trustero:"display:Username;order:1"`
	Email           string `trustero:"display:Email;order:2"`
	ComputerId      string `trustero:"display:Computer Id;order:3"`
	OsVersion       string `trustero:"display:OS Version;order:4"`
	MacAddress      string `trustero:"display:MAC Address;order:5"`
	Users           int    `trustero:"display:Users;order:6"`
	FilevaultUsers  int    `trustero:"display:Filevault Users;order:7"`
	XprotectVersion string `trustero:"display:Xprotect Version;order:8"`
}

func getComputerEvidence(computerService *computers.Service) (evidence *receptor_sdk.Evidence, err error) {
	computerList, resp, err := computerService.List()
	if err != nil {
		receptorLog.Err(err, "could not generate evidence, error in getComputerEvidence")
		return
	}

	evidence = receptor_sdk.NewEvidence(serviceName, "JamfComputers", "ComputerList", "List of Computers in Jamf")
	evidence.AddSource(resp.Request.URL.String(), computerList)
	for _, cmp := range computerList {
		result := &computerResponse{}
		if result.computer, resp, err = computerService.GetById(cmp.Id); err != nil {
			receptorLog.Error(err.Error())
			return
		}

		item := JamfComputerInfo{
			Username:        result.computer.UserLocation.RealName,
			Email:           result.computer.UserLocation.EmailAddress,
			ComputerId:      result.computer.General.UDID,
			OsVersion:       fmt.Sprintf("%s %s", result.computer.Hardware.OSName, result.computer.Hardware.OSVersion),
			MacAddress:      result.computer.General.MACAddress,
			Users:           len(result.computer.Groups.Memberships),
			FilevaultUsers:  len(result.computer.Hardware.FilevaultUsers),
			XprotectVersion: result.computer.Hardware.XProtectVersion,
		}

		evidence.AddRow(item)
		evidence.AddSource(resp.Request.URL.String(), result.computer)

	}

	return
}
