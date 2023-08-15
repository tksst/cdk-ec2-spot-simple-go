// CDK construct library to create EC2 Spot Instances simply.
package cdkec2spotsimple

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-ec2-spot-simple.SpotInstance",
		reflect.TypeOf((*SpotInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addSecurityGroup", GoMethod: "AddSecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addUserData", GoMethod: "AddUserData"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "instance", GoGetter: "Instance"},
			_jsii_.MemberProperty{JsiiProperty: "instanceAvailabilityZone", GoGetter: "InstanceAvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "instanceId", GoGetter: "InstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "instancePrivateDnsName", GoGetter: "InstancePrivateDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "instancePrivateIp", GoGetter: "InstancePrivateIp"},
			_jsii_.MemberProperty{JsiiProperty: "instancePublicDnsName", GoGetter: "InstancePublicDnsName"},
			_jsii_.MemberProperty{JsiiProperty: "instancePublicIp", GoGetter: "InstancePublicIp"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "osType", GoGetter: "OsType"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userData", GoGetter: "UserData"},
		},
		func() interface{} {
			j := jsiiProxy_SpotInstance{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2Instance)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-ec2-spot-simple.SpotInstanceProps",
		reflect.TypeOf((*SpotInstanceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-ec2-spot-simple.SpotReqCancelerProps",
		reflect.TypeOf((*SpotReqCancelerProps)(nil)).Elem(),
	)
}
