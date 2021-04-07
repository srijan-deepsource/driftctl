// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/cloudskiff/driftctl/pkg/dctlcty"
)

const AwsDefaultSubnetResourceType = "aws_default_subnet"

type AwsDefaultSubnet struct {
	Arn                         *string           `cty:"arn" computed:"true"`
	AssignIpv6AddressOnCreation *bool             `cty:"assign_ipv6_address_on_creation" computed:"true"`
	AvailabilityZone            *string           `cty:"availability_zone"`
	AvailabilityZoneId          *string           `cty:"availability_zone_id" computed:"true"`
	CidrBlock                   *string           `cty:"cidr_block" computed:"true"`
	Id                          string            `cty:"id" computed:"true"`
	Ipv6CidrBlock               *string           `cty:"ipv6_cidr_block" computed:"true"`
	Ipv6CidrBlockAssociationId  *string           `cty:"ipv6_cidr_block_association_id" computed:"true"`
	MapPublicIpOnLaunch         *bool             `cty:"map_public_ip_on_launch" computed:"true"`
	OutpostArn                  *string           `cty:"outpost_arn"`
	OwnerId                     *string           `cty:"owner_id" computed:"true"`
	Tags                        map[string]string `cty:"tags"`
	VpcId                       *string           `cty:"vpc_id" computed:"true"`
	Timeouts                    *struct {
		Create *string `cty:"create"`
		Delete *string `cty:"delete"`
	} `cty:"timeouts" diff:"-"`
	CtyVal *cty.Value `diff:"-"`
}

func (r *AwsDefaultSubnet) TerraformId() string {
	return r.Id
}

func (r *AwsDefaultSubnet) TerraformType() string {
	return AwsDefaultSubnetResourceType
}

func (r *AwsDefaultSubnet) CtyValue() *cty.Value {
	return r.CtyVal
}

var awsDefaultSubnetTags = map[string]string{}

func awsDefaultSubnetNormalizer(val *dctlcty.CtyAttributes) {
	val.SafeDelete([]string{"timeouts"})
}
