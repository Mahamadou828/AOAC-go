//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsiam

import (
	"fmt"
)

func (i *jsiiProxy_IResourceWithPolicy) validateAddToResourcePolicyParameters(statement PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

