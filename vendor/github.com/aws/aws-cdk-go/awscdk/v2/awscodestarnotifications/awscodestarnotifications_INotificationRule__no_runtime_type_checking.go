//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscodestarnotifications

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_INotificationRule) validateAddTargetParameters(target INotificationRuleTarget) error {
	return nil
}

