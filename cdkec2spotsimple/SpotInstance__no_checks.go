//go:build no_runtime_type_checking

// CDK construct library to create EC2 Spot Instances simply.
package cdkec2spotsimple

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SpotInstance) validateAddSecurityGroupParameters(securityGroup awsec2.ISecurityGroup) error {
	return nil
}

func (s *jsiiProxy_SpotInstance) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_SpotInstance) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SpotInstance) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SpotInstance) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSpotInstance_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSpotInstance_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSpotInstanceParameters(scope constructs.Construct, id *string, props *SpotInstanceProps) error {
	return nil
}

