package cdkec2spotsimple

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties of `SpotInstance`.
type SpotInstanceProps struct {
	// Type of instance to launch.
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// AMI to launch.
	MachineImage awsec2.IMachineImage `field:"required" json:"machineImage" yaml:"machineImage"`
	// VPC to launch the instance in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Whether the instance could initiate connections to anywhere by default.
	//
	// This property is only used when you do not provide a security group.
	// Default: true.
	//
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// In which AZ to place the instance within the VPC.
	// Default: - Random zone.
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.
	//
	// Each instance that is launched has an associated root device volume,
	// either an Amazon EBS volume or an instance store volume.
	// You can use block device mappings to specify additional EBS volumes or
	// instance store volumes to attach to an instance when it is launched.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html
	//
	// Default: - Uses the block device mapping of the AMI.
	//
	BlockDevices *[]*awsec2.BlockDevice `field:"optional" json:"blockDevices" yaml:"blockDevices"`
	// Whether "Detailed Monitoring" is enabled for this instance Keep in mind that Detailed Monitoring results in extra charges.
	// See: http://aws.amazon.com/cloudwatch/pricing/
	//
	// Default: - false.
	//
	DetailedMonitoring *bool `field:"optional" json:"detailedMonitoring" yaml:"detailedMonitoring"`
	// Apply the given CloudFormation Init configuration to the instance at startup.
	// Default: - no CloudFormation init.
	//
	Init awsec2.CloudFormationInit `field:"optional" json:"init" yaml:"init"`
	// Use the given options for applying CloudFormation Init.
	//
	// Describes the configsets to use and the timeout to wait.
	// Default: - default options.
	//
	InitOptions *awsec2.ApplyCloudFormationInitOptions `field:"optional" json:"initOptions" yaml:"initOptions"`
	// The name of the instance.
	// Default: - CDK generated name.
	//
	InstanceName *string `field:"optional" json:"instanceName" yaml:"instanceName"`
	// Name of SSH keypair to grant access to instance.
	// Default: - No SSH access will be possible.
	//
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Defines a private IP address to associate with an instance.
	//
	// Private IP should be available within the VPC that the instance is build within.
	// Default: - no association.
	//
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Propagate the EC2 instance tags to the EBS volumes.
	// Default: - false.
	//
	PropagateTagsToVolumeOnCreation *bool `field:"optional" json:"propagateTagsToVolumeOnCreation" yaml:"propagateTagsToVolumeOnCreation"`
	// Whether IMDSv2 should be required on this instance.
	// Default: - false.
	//
	RequireImdsv2 *bool `field:"optional" json:"requireImdsv2" yaml:"requireImdsv2"`
	// The length of time to wait for the resourceSignalCount.
	//
	// The maximum value is 43200 (12 hours).
	// Default: Duration.minutes(5)
	//
	ResourceSignalTimeout awscdk.Duration `field:"optional" json:"resourceSignalTimeout" yaml:"resourceSignalTimeout"`
	// An IAM role to associate with the instance profile assigned to this Auto Scaling Group.
	//
	// The role must be assumable by the service principal `ec2.amazonaws.com`:
	//
	// Example:
	//   const role = new iam.Role(this, 'MyRole', {
	//     assumedBy: new iam.ServicePrincipal('ec2.amazonaws.com')
	//   });
	//
	// Default: - A role will automatically be created, it can be accessed via the `role` property.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security Group to assign to this instance.
	// Default: - create new security group.
	//
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Specifies whether to enable an instance launched in a VPC to perform NAT.
	//
	// This controls whether source/destination checking is enabled on the instance.
	// A value of true means that checking is enabled, and false means that checking is disabled.
	// The value must be false for the instance to perform NAT.
	// Default: true.
	//
	SourceDestCheck *bool `field:"optional" json:"sourceDestCheck" yaml:"sourceDestCheck"`
	// Specific UserData to use.
	//
	// The UserData may still be mutated after creation.
	// Default: - A UserData object appropriate for the MachineImage's
	// Operating System is created.
	//
	UserData awsec2.UserData `field:"optional" json:"userData" yaml:"userData"`
	// Changes to the UserData force replacement.
	//
	// Depending the EC2 instance type, changing UserData either
	// restarts the instance or replaces the instance.
	//
	// - Instance store-backed instances are replaced.
	// - EBS-backed instances are restarted.
	//
	// By default, restarting does not execute the new UserData so you
	// will need a different mechanism to ensure the instance is restarted.
	//
	// Setting this to `true` will make the instance's Logical ID depend on the
	// UserData, which will cause CloudFormation to replace it if the UserData
	// changes.
	// Default: - true iff `initOptions` is specified, false otherwise.
	//
	UserDataCausesReplacement *bool `field:"optional" json:"userDataCausesReplacement" yaml:"userDataCausesReplacement"`
	// Where to place the instance within the VPC.
	// Default: - Private subnets.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// The options for the Spot instances.
	// Default: - Use the Launch Template's default InstanceMarketOptions.
	//
	SpotOptions *awsec2.LaunchTemplateSpotOptions `field:"optional" json:"spotOptions" yaml:"spotOptions"`
	// Options related to Lambda functions to cancel spot requests.
	SpotReqCancelerOptions *SpotReqCancelerProps `field:"optional" json:"spotReqCancelerOptions" yaml:"spotReqCancelerOptions"`
}

