package jamf_api

import (
	"fmt"
)

// General info
type General struct {
	// They are default attributes.
	// Note `,omitempty` is added because attributes may internally changed by Jamf.
	ID               uint64           `json:"id,omitempty"`
	Name             string           `json:"name,omitempty"`
	SerialNumber     string           `json:"serial_number,omitempty"`
	Udid             string           `json:"udid,omitempty"`
	Model            string           `json:"model,omitempty"`
	MacAddress       string           `json:"mac_address,omitempty"`
	ReportDateUtc    string           `json:"report_date_utc,omitempty"`
	ReportDateEpoch  float64          `json:"report_date_epoch,omitempty"`
	ManagementStatus ManagementStatus `json:"management_status,omitempty"`
}

// ManagementStatus info
type ManagementStatus struct {
	EnrolledViaDep         bool `json:"enrolled_via_dep,omitempty"`
	UserApprovedEnrollment bool `json:"user_approved_enrollment,omitempty"`
	UserApprovedMdm        bool `json:"user_approved_mdm,omitempty"`
}

// Hardware info
type Hardware struct {
	ModelIdentifier string `json:"model_identifier,omitempty"`
	OsName          string `json:"os_name,omitempty"`
	OsVersion       string `json:"os_version,omitempty"`
}

// Application info
type Application struct {
	Name    string `json:"name,omitempty"`
	Path    string `json:"path,omitempty"`
	Version string `json:"version,omitempty"`
}

// Software info
type Software struct {
	Applications []Application `json:"applications,omitempty"`
}

// Computer info
type Computer struct {
	General  General  `json:"general,omitempty"`
	Hardware Hardware `json:"hardware,omitempty"`
	Software Software `json:"software,omitempty"`
}

// ComputerWrapper Wrapper for the toplevel element
type ComputerWrapper struct {
	Computer Computer `json:"computer"`
}

// ComputerService Service to call
type ComputerService struct {
	Service
	Client *Client
	data   interface{}
}

func (s *ComputerService) GetComputer(id uint64) (*Computer, error) {
	res, err := s.Client.DoGetRequest(fmt.Sprintf("/computers/id/%d", id))

	if err != nil {
		return nil, fmt.Errorf("failed requesting: %v", err)
	}

	var computer *ComputerWrapper
	if err := JSONBodyDecoder(res, &computer); err != nil {
		return nil, fmt.Errorf("failed decoding response body: %v", err)
	}

	return &computer.Computer, err
}
