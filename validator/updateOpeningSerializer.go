package validator

import "fmt"

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role == "" || r.Company == "" || r.Location == "" || r.Link == "" || r.Remote == nil || r.Salary <= 0 {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}
