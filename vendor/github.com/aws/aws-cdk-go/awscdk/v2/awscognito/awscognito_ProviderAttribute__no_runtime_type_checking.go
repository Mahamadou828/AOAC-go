//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func validateProviderAttribute_OtherParameters(attributeName *string) error {
	return nil
}
