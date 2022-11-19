package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Lambda that will only ever be added to a stack once.
//
// This construct is a way to guarantee that the lambda function will be guaranteed to be part of the stack,
// once and only once, irrespective of how many times the construct is declared to be part of the stack.
// This is guaranteed as long as the `uuid` property and the optional `lambdaPurpose` property stay the same
// whenever they're declared into the stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var architecture architecture
//   var code code
//   var codeSigningConfig codeSigningConfig
//   var destination iDestination
//   var eventSource iEventSource
//   var fileSystem fileSystem
//   var key key
//   var lambdaInsightsVersion lambdaInsightsVersion
//   var layerVersion layerVersion
//   var policyStatement policyStatement
//   var profilingGroup profilingGroup
//   var queue queue
//   var role role
//   var runtime runtime
//   var securityGroup securityGroup
//   var size size
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var topic topic
//   var vpc vpc
//
//   singletonFunction := awscdk.Aws_lambda.NewSingletonFunction(this, jsii.String("MySingletonFunction"), &singletonFunctionProps{
//   	code: code,
//   	handler: jsii.String("handler"),
//   	runtime: runtime,
//   	uuid: jsii.String("uuid"),
//
//   	// the properties below are optional
//   	allowAllOutbound: jsii.Boolean(false),
//   	allowPublicSubnet: jsii.Boolean(false),
//   	architecture: architecture,
//   	codeSigningConfig: codeSigningConfig,
//   	currentVersionOptions: &versionOptions{
//   		codeSha256: jsii.String("codeSha256"),
//   		description: jsii.String("description"),
//   		maxEventAge: cdk.duration.minutes(jsii.Number(30)),
//   		onFailure: destination,
//   		onSuccess: destination,
//   		provisionedConcurrentExecutions: jsii.Number(123),
//   		removalPolicy: cdk.removalPolicy_DESTROY,
//   		retryAttempts: jsii.Number(123),
//   	},
//   	deadLetterQueue: queue,
//   	deadLetterQueueEnabled: jsii.Boolean(false),
//   	deadLetterTopic: topic,
//   	description: jsii.String("description"),
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	environmentEncryption: key,
//   	ephemeralStorageSize: size,
//   	events: []*iEventSource{
//   		eventSource,
//   	},
//   	filesystem: fileSystem,
//   	functionName: jsii.String("functionName"),
//   	initialPolicy: []*policyStatement{
//   		policyStatement,
//   	},
//   	insightsVersion: lambdaInsightsVersion,
//   	lambdaPurpose: jsii.String("lambdaPurpose"),
//   	layers: []iLayerVersion{
//   		layerVersion,
//   	},
//   	logRetention: awscdk.Aws_logs.retentionDays_ONE_DAY,
//   	logRetentionRetryOptions: &logRetentionRetryOptions{
//   		base: cdk.*duration.minutes(jsii.Number(30)),
//   		maxRetries: jsii.Number(123),
//   	},
//   	logRetentionRole: role,
//   	maxEventAge: cdk.*duration.minutes(jsii.Number(30)),
//   	memorySize: jsii.Number(123),
//   	onFailure: destination,
//   	onSuccess: destination,
//   	profiling: jsii.Boolean(false),
//   	profilingGroup: profilingGroup,
//   	reservedConcurrentExecutions: jsii.Number(123),
//   	retryAttempts: jsii.Number(123),
//   	role: role,
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	timeout: cdk.*duration.minutes(jsii.Number(30)),
//   	tracing: awscdk.*Aws_lambda.tracing_ACTIVE,
//   	vpc: vpc,
//   	vpcSubnets: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//   	},
//   })
//
type SingletonFunction interface {
	FunctionBase
	// The architecture of this Lambda Function.
	Architecture() Architecture
	// Whether the addPermission() call adds any permissions.
	//
	// True for new Lambdas, false for version $LATEST and imported Lambdas
	// from different accounts.
	CanCreatePermissions() *bool
	// Access the Connections object.
	//
	// Will fail if not a VPC-enabled Lambda Function.
	Connections() awsec2.Connections
	// Returns a `lambda.Version` which represents the current version of this singleton Lambda function. A new version will be created every time the function's configuration changes.
	//
	// You can specify options for this version using the `currentVersionOptions`
	// prop when initializing the `lambda.SingletonFunction`.
	CurrentVersion() Version
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The ARN fo the function.
	FunctionArn() *string
	// The name of the function.
	FunctionName() *string
	// The principal this Lambda Function is running as.
	GrantPrincipal() awsiam.IPrincipal
	// Whether or not this Lambda function was bound to a VPC.
	//
	// If this is is `false`, trying to access the `connections` object will fail.
	IsBoundToVpc() *bool
	// The `$LATEST` version of this function.
	//
	// Note that this is reference to a non-specific AWS Lambda version, which
	// means the function this version refers to can return different results in
	// different invocations.
	//
	// To obtain a reference to an explicit version which references the current
	// function configuration, use `lambdaFunction.currentVersion` instead.
	LatestVersion() IVersion
	// The LogGroup where the Lambda function's logs are made available.
	//
	// If either `logRetention` is set or this property is called, a CloudFormation custom resource is added to the stack that
	// pre-creates the log group as part of the stack deployment, if it already doesn't exist, and sets the correct log retention
	// period (never expire, by default).
	//
	// Further, if the log group already exists and the `logRetention` is not set, the custom resource will reset the log retention
	// to never expire even if it was configured with a different value.
	LogGroup() awslogs.ILogGroup
	// The tree node.
	Node() constructs.Node
	// The construct node where permissions are attached.
	PermissionsNode() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The ARN(s) to put into the resource field of the generated IAM policy for grantInvoke().
	ResourceArnsForGrantInvoke() *[]*string
	// The IAM role associated with this function.
	//
	// Undefined if the function was imported without a role.
	Role() awsiam.IRole
	// The runtime environment for the Lambda function.
	Runtime() Runtime
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Using node.addDependency() does not work on this method as the underlying lambda function is modeled as a singleton across the stack. Use this method instead to declare dependencies.
	AddDependency(up ...constructs.IDependable)
	// Adds an environment variable to this Lambda function.
	//
	// If this is a ref to a Lambda function, this operation results in a no-op.
	AddEnvironment(key *string, value *string, options *EnvironmentOptions) Function
	// Adds an event source to this function.
	//
	// Event sources are implemented in the @aws-cdk/aws-lambda-event-sources module.
	//
	// The following example adds an SQS Queue as an event source:
	// ```
	// import { SqsEventSource } from '@aws-cdk/aws-lambda-event-sources';
	// myFunction.addEventSource(new SqsEventSource(myQueue));
	// ```.
	AddEventSource(source IEventSource)
	// Adds an event source that maps to this AWS Lambda function.
	AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping
	// Adds a url to this lambda function.
	AddFunctionUrl(options *FunctionUrlOptions) FunctionUrl
	// Adds one or more Lambda Layers to this Lambda function.
	AddLayers(layers ...ILayerVersion)
	// Adds a permission to the Lambda resource policy.
	AddPermission(name *string, permission *Permission)
	// Adds a statement to the IAM role assumed by the instance.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Configures options for asynchronous invocation.
	ConfigureAsyncInvoke(options *EventInvokeConfigOptions)
	// A warning will be added to functions under the following conditions: - permissions that include `lambda:InvokeFunction` are added to the unqualified function.
	//
	// - function.currentVersion is invoked before or after the permission is created.
	//
	// This applies only to permissions on Lambda functions, not versions or aliases.
	// This function is overridden as a noOp for QualifiedFunctionBase.
	ConsiderWarningOnInvokeFunctionPermissions(scope constructs.Construct, action *string)
	// The SingletonFunction construct cannot be added as a dependency of another construct using node.addDependency(). Use this method instead to declare this as a dependency of another construct.
	DependOn(down constructs.IConstruct)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grant the given identity permissions to invoke this Lambda.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to invoke this Lambda Function URL.
	GrantInvokeUrl(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Function.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How long execution of this Lambda takes.
	//
	// Average over 5 minutes.
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How many invocations of this Lambda fail.
	//
	// Sum over 5 minutes.
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How often this Lambda is invoked.
	//
	// Sum over 5 minutes.
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How often this Lambda is throttled.
	//
	// Sum over 5 minutes.
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
	WarnInvokeFunctionPermissions(scope constructs.Construct)
}

// The jsii proxy struct for SingletonFunction
type jsiiProxy_SingletonFunction struct {
	jsiiProxy_FunctionBase
}

func (j *jsiiProxy_SingletonFunction) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) CurrentVersion() Version {
	var returns Version
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) LatestVersion() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) ResourceArnsForGrantInvoke() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnsForGrantInvoke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) Runtime() Runtime {
	var returns Runtime
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewSingletonFunction(scope constructs.Construct, id *string, props *SingletonFunctionProps) SingletonFunction {
	_init_.Initialize()

	if err := validateNewSingletonFunctionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SingletonFunction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.SingletonFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSingletonFunction_Override(s SingletonFunction, scope constructs.Construct, id *string, props *SingletonFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.SingletonFunction",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SingletonFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSingletonFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.SingletonFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func SingletonFunction_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateSingletonFunction_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.SingletonFunction",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func SingletonFunction_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateSingletonFunction_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.SingletonFunction",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingletonFunction) AddDependency(up ...constructs.IDependable) {
	args := []interface{}{}
	for _, a := range up {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addDependency",
		args,
	)
}

func (s *jsiiProxy_SingletonFunction) AddEnvironment(key *string, value *string, options *EnvironmentOptions) Function {
	if err := s.validateAddEnvironmentParameters(key, value, options); err != nil {
		panic(err)
	}
	var returns Function

	_jsii_.Invoke(
		s,
		"addEnvironment",
		[]interface{}{key, value, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingletonFunction) AddEventSource(source IEventSource) {
	if err := s.validateAddEventSourceParameters(source); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addEventSource",
		[]interface{}{source},
	)
}

func (s *jsiiProxy_SingletonFunction) AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping {
	if err := s.validateAddEventSourceMappingParameters(id, options); err != nil {
		panic(err)
	}
	var returns EventSourceMapping

	_jsii_.Invoke(
		s,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingletonFunction) AddFunctionUrl(options *FunctionUrlOptions) FunctionUrl {
	if err := s.validateAddFunctionUrlParameters(options); err != nil {
		panic(err)
	}
	var returns FunctionUrl

	_jsii_.Invoke(
		s,
		"addFunctionUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingletonFunction) AddLayers(layers ...ILayerVersion) {
	args := []interface{}{}
	for _, a := range layers {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addLayers",
		args,
	)
}

func (s *jsiiProxy_SingletonFunction) AddPermission(name *string, permission *Permission) {
	if err := s.validateAddPermissionParameters(name, permission); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addPermission",
		[]interface{}{name, permission},
	)
}

func (s *jsiiProxy_SingletonFunction) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := s.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (s *jsiiProxy_SingletonFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_SingletonFunction) ConfigureAsyncInvoke(options *EventInvokeConfigOptions) {
	if err := s.validateConfigureAsyncInvokeParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (s *jsiiProxy_SingletonFunction) ConsiderWarningOnInvokeFunctionPermissions(scope constructs.Construct, action *string) {
	if err := s.validateConsiderWarningOnInvokeFunctionPermissionsParameters(scope, action); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"considerWarningOnInvokeFunctionPermissions",
		[]interface{}{scope, action},
	)
}

func (s *jsiiProxy_SingletonFunction) DependOn(down constructs.IConstruct) {
	if err := s.validateDependOnParameters(down); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"dependOn",
		[]interface{}{down},
	)
}

func (s *jsiiProxy_SingletonFunction) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingletonFunction) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := s.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingletonFunction) GetResourceNameAttribute(nameAttr *string) *string {
	if err := s.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingletonFunction) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingletonFunction) GrantInvokeUrl(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateGrantInvokeUrlParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantInvokeUrl",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingletonFunction) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingletonFunction) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingletonFunction) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingletonFunction) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingletonFunction) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingletonFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingletonFunction) WarnInvokeFunctionPermissions(scope constructs.Construct) {
	if err := s.validateWarnInvokeFunctionPermissionsParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"warnInvokeFunctionPermissions",
		[]interface{}{scope},
	)
}
