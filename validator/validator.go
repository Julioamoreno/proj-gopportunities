package validator

import (
	"fmt"

	"github.com/Julioamoreno/proj-gopportunities/config"
)

var (
	logger *config.Logger
)

func Init() {
	logger = config.GetLogger("Validator")
}

func IsRequiredError(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}
