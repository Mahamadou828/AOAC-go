//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func validateArchitecture_CustomParameters(name *string) error {
	return nil
}

