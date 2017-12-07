package jamf_api

import (
	"sort"
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

type By func(c1, c2 *Computer) bool

func (by By) Sort(computers []Computer) {
	cp := &computerSorter{
		computers:computers,
		by:by,
	}
	sort.Sort(cp)
}

type computerSorter struct {
	computers []Computer
	by func(c1, c2 *Computer) bool
}

func (s *computerSorter) Len() int {
	return len(s.computers)
}

func (s *computerSorter) Swap(i, j int) {
	s.computers[i], s.computers[j] = s.computers[j], s.computers[i]
}

func (s *computerSorter) Less(i, j int) bool {
	return s.by(&s.computers[i], &s.computers[j])
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