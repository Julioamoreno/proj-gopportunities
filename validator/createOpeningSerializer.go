package validator

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("malformed request body")
	}
	if r.Role == "" {
		return IsRequiredError("role", "string")
	}
	if r.Company == "" {
		return IsRequiredError("company", "string")
	}
	if r.Location == "" {
		return IsRequiredError("location", "string")
	}
	if r.Link == "" {
		return IsRequiredError("link", "string")
	}
	if r.Remote == nil {
		return IsRequiredError("remote", "boolean")
	}
	if r.Salary <= 0 {
		return IsRequiredError("salary", "int64")
	}

	return nil
}
