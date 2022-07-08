package commands

import (
	"fmt"
	"github.com/Mahamadou828/AOAC/business/sys/aws"
)

func SSMCreateSecret(cl *aws.Client) error {
	fmt.Println("Starting creating secret")

	for {
		var name, val string
		fmt.Printf("enter the secret name: ")
		if _, err := fmt.Scan(&name); err != nil {
			return fmt.Errorf("invalid secret name: %v", err)
		}
		fmt.Printf("enter the secret value: ")
		if _, err := fmt.Scan(&val); err != nil {
			return fmt.Errorf("invalid secret value: %v", err)
		}
		fmt.Println("creating new secret")
		if err := cl.SSM.CreateSecret(name, val); err != nil {
			return fmt.Errorf("can't create secret: %v", err)
		}

		fmt.Println("secret created")
		var choice string
		fmt.Printf("Would you like to continue (y|n): ")
		if _, err := fmt.Scan(&choice); err != nil {
			return fmt.Errorf("failed to continue: %v", err)
		}
		if choice == "n" {
			break
		}

	}
	return nil
}

func SSMCreatePool(cl *aws.Client) error {
	if err := cl.SSM.CreatePool(); err != nil {
		return fmt.Errorf("can't create pool: %v", err)
	}
	return nil
}
