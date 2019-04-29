package jamf_api

import (
	"fmt"
	"sort"
)

// Computer info
type ComputersComputer struct {
	// They are default attributes.
	// Note `,omitempty` is added because attributes may internally changed by Jamf.
	ID              uint64  `json:"id,omitempty"`
	Name            string  `json:"name,omitempty"`
	Managed         bool    `json:"managed,omitempty"`
	Username        string  `json:"username,omitempty"`
	Model           string  `json:"model,omitempty"`
	Department      string  `json:"department,omitempty"`
	Building        string  `json:"building,omitempty"`
	MacAddress      string  `json:"mac_address,omitempty"`
	Udid            string  `json:"udid,omitempty"`
	SerialNumber    string  `json:"serial_number,omitempty"`
	ReportDateUtc   string  `json:"report_date_utc,omitempty"`
	ReportDateEpoch float64 `json:"report_date_epoch,omitempty"`
}

// Computers describe
type Computers struct {
	Computers []ComputersComputer `json:"computers"`
}

// SortByComputerName sorts computers by Computer Name
func SortByComputerName(computers []ComputersComputer) {
	swapRule := func(c1 *ComputersComputer, c2 *ComputersComputer) bool { return c1.Name < c2.Name }

	sorter := &computerSorter{
		computers: computers,
		by:        swapRule,
	}

	sort.Sort(sorter)
}

// SortByUserName sorts computers by UserName
func SortByUserName(computers []ComputersComputer) {
	swapRule := func(c1 *ComputersComputer, c2 *ComputersComputer) bool { return c1.Username < c2.Username }

	sorter := &computerSorter{
		computers: computers,
		by:        swapRule,
	}

	sort.Sort(sorter)
}

type computerSorter struct {
	computers []ComputersComputer
	by        func(c1, c2 *ComputersComputer) bool
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

type ComputersService struct {
	Service
	Client *Client
	data   interface{}
}

func (s *ComputersService) GetComputers() ([]ComputersComputer, error) {
	res, err := s.Client.DoGetRequest("/computers/subset/basic")

	if err != nil {
		return nil, fmt.Errorf("failed requesting: %v", err)
	}

	var computers *Computers
	if err := JSONBodyDecoder(res, &computers); err != nil {
		return nil, fmt.Errorf("failed decoding response body: %v", err)
	}

	return computers.Computers, err
}
