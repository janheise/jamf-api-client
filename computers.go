package jamf_api

import (
	//"net/http"
	"fmt"
)

// Computer info
type Computer struct {
	// They are default attributes.
	// Note `,omitempty` is added because attributes may internally changed by Jamf.
	ComputerName                 string `json:"Computer_Name,omitempty"`
	FullName                     string `json:"Full_Name,omitempty"`
	LastReportedIPAddress        string `json:"Last_Reported_IP_Address,omitempty"`
	LastInventoryUpdate          string `json:"Last_Inventory_Update,omitempty"`
	LastICloudBackup             string `json:"Last_iCloud_Backup,omitempty"`
	LastEnrollment               string `json:"Last_Enrollment,omitempty"`
	LastCheckIn                  string `json:"Last_Check_in,omitempty"`
	AssetTag                     string `json:"Asset_Tag,omitempty"`
}

// Computers describe
type Computers struct {
	Computers []Computer `json:"computer_reports"`
}

type ComputerService struct {
	Service
	Client *Client
	data	interface{}
}

func (s *ComputerService) GetComputers() ([]Computer, error) {
	res, err := s.Client.DoGetRequest("/computerreports/id/1")

	if err != nil {
		return nil, fmt.Errorf("failed requesting: %v", err)
	}

	var computers *Computers
	if err := JSONBodyDecoder(res, &computers); err != nil {
		return nil, fmt.Errorf("failed decoding response body: %v", err)
	}

	return computers.Computers, err
}