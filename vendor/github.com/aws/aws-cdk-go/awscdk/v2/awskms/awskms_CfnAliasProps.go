package awskms


// Properties for defining a `CfnAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAliasProps := &cfnAliasProps{
//   	aliasName: jsii.String("aliasName"),
//   	targetKeyId: jsii.String("targetKeyId"),
//   }
//
type CfnAliasProps struct {
	// Specifies the alias name. This value must begin with `alias/` followed by a name, such as `alias/ExampleAlias` .
	//
	// > If you change the value of a `Replacement` property, such as `AliasName` , the existing alias is deleted and a new alias is created for the specified KMS key. This change can disrupt applications that use the alias. It can also allow or deny access to a KMS key affected by attribute-based access control (ABAC).
	//
	// The alias must be string of 1-256 characters. It can contain only alphanumeric characters, forward slashes (/), underscores (_), and dashes (-). The alias name cannot begin with `alias/aws/` . The `alias/aws/` prefix is reserved for [AWS managed keys](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk) .
	//
	// *Pattern* : `alias/^[a-zA-Z0-9/_-]+$`
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `256`.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	// Associates the alias with the specified [customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk) . The KMS key must be in the same AWS account and Region.
	//
	// A valid key ID is required. If you supply a null or empty string value, this operation returns an error.
	//
	// For help finding the key ID and ARN, see [Finding the key ID and ARN](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html#find-cmk-id-arn) in the *AWS Key Management Service Developer Guide* .
	//
	// Specify the key ID or the key ARN of the KMS key.
	//
	// For example:
	//
	// - Key ID: `1234abcd-12ab-34cd-56ef-1234567890ab`
	// - Key ARN: `arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`
	//
	// To get the key ID and key ARN for a KMS key, use [ListKeys](https://docs.aws.amazon.com/kms/latest/APIReference/API_ListKeys.html) or [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html) .
	TargetKeyId *string `field:"required" json:"targetKeyId" yaml:"targetKeyId"`
}

