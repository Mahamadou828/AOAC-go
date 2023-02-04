//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsec2

import (
	"fmt"
)

func validateNewGatewayVpcEndpointAwsServiceParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

