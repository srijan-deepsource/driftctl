// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/cloudskiff/driftctl/pkg/dctlcty"
)

const AwsInstanceResourceType = "aws_instance"

type AwsInstance struct {
	Ami                               *string           `cty:"ami"`
	Arn                               *string           `cty:"arn" computed:"true"`
	AssociatePublicIpAddress          *bool             `cty:"associate_public_ip_address" computed:"true"`
	AvailabilityZone                  *string           `cty:"availability_zone" computed:"true"`
	CpuCoreCount                      *int              `cty:"cpu_core_count" computed:"true"`
	CpuThreadsPerCore                 *int              `cty:"cpu_threads_per_core" computed:"true"`
	DisableApiTermination             *bool             `cty:"disable_api_termination"`
	EbsOptimized                      *bool             `cty:"ebs_optimized"`
	GetPasswordData                   *bool             `cty:"get_password_data"`
	Hibernation                       *bool             `cty:"hibernation"`
	HostId                            *string           `cty:"host_id" computed:"true"`
	IamInstanceProfile                *string           `cty:"iam_instance_profile"`
	Id                                string            `cty:"id" computed:"true"`
	InstanceInitiatedShutdownBehavior *string           `cty:"instance_initiated_shutdown_behavior"`
	InstanceState                     *string           `cty:"instance_state" computed:"true"`
	InstanceType                      *string           `cty:"instance_type"`
	Ipv6AddressCount                  *int              `cty:"ipv6_address_count" computed:"true"`
	Ipv6Addresses                     []string          `cty:"ipv6_addresses" computed:"true"`
	KeyName                           *string           `cty:"key_name" computed:"true"`
	Monitoring                        *bool             `cty:"monitoring"`
	OutpostArn                        *string           `cty:"outpost_arn" computed:"true"`
	PasswordData                      *string           `cty:"password_data" computed:"true"`
	PlacementGroup                    *string           `cty:"placement_group" computed:"true"`
	PrimaryNetworkInterfaceId         *string           `cty:"primary_network_interface_id" computed:"true"`
	PrivateDns                        *string           `cty:"private_dns" computed:"true"`
	PrivateIp                         *string           `cty:"private_ip" computed:"true"`
	PublicDns                         *string           `cty:"public_dns" computed:"true"`
	PublicIp                          *string           `cty:"public_ip" computed:"true"`
	SecondaryPrivateIps               []string          `cty:"secondary_private_ips" computed:"true"`
	SecurityGroups                    []string          `cty:"security_groups" computed:"true"`
	SourceDestCheck                   *bool             `cty:"source_dest_check"`
	SubnetId                          *string           `cty:"subnet_id" computed:"true"`
	Tags                              map[string]string `cty:"tags"`
	Tenancy                           *string           `cty:"tenancy" computed:"true"`
	UserData                          *string           `cty:"user_data"`
	UserDataBase64                    *string           `cty:"user_data_base64"`
	VolumeTags                        map[string]string `cty:"volume_tags" diff:"-" computed:"true"`
	VpcSecurityGroupIds               []string          `cty:"vpc_security_group_ids" computed:"true"`
	CreditSpecification               *[]struct {
		CpuCredits *string `cty:"cpu_credits"`
	} `cty:"credit_specification"`
	EbsBlockDevice *[]struct {
		DeleteOnTermination *bool   `cty:"delete_on_termination"`
		DeviceName          *string `cty:"device_name"`
		Encrypted           *bool   `cty:"encrypted" computed:"true"`
		Iops                *int    `cty:"iops" computed:"true"`
		KmsKeyId            *string `cty:"kms_key_id" computed:"true"`
		SnapshotId          *string `cty:"snapshot_id" computed:"true"`
		VolumeId            *string `cty:"volume_id" computed:"true"`
		VolumeSize          *int    `cty:"volume_size" computed:"true"`
		VolumeType          *string `cty:"volume_type" computed:"true"`
	} `cty:"ebs_block_device"`
	EphemeralBlockDevice *[]struct {
		DeviceName  *string `cty:"device_name"`
		NoDevice    *bool   `cty:"no_device"`
		VirtualName *string `cty:"virtual_name"`
	} `cty:"ephemeral_block_device"`
	MetadataOptions *[]struct {
		HttpEndpoint            *string `cty:"http_endpoint" computed:"true"`
		HttpPutResponseHopLimit *int    `cty:"http_put_response_hop_limit" computed:"true"`
		HttpTokens              *string `cty:"http_tokens" computed:"true"`
	} `cty:"metadata_options"`
	NetworkInterface *[]struct {
		DeleteOnTermination *bool   `cty:"delete_on_termination"`
		DeviceIndex         *int    `cty:"device_index"`
		NetworkInterfaceId  *string `cty:"network_interface_id"`
	} `cty:"network_interface"`
	RootBlockDevice *[]struct {
		DeleteOnTermination *bool   `cty:"delete_on_termination"`
		DeviceName          *string `cty:"device_name" computed:"true"`
		Encrypted           *bool   `cty:"encrypted" computed:"true"`
		Iops                *int    `cty:"iops" computed:"true"`
		KmsKeyId            *string `cty:"kms_key_id" computed:"true"`
		VolumeId            *string `cty:"volume_id" computed:"true"`
		VolumeSize          *int    `cty:"volume_size" computed:"true"`
		VolumeType          *string `cty:"volume_type" computed:"true"`
	} `cty:"root_block_device"`
	Timeouts *struct {
		Create *string `cty:"create"`
		Delete *string `cty:"delete"`
		Update *string `cty:"update"`
	} `cty:"timeouts" diff:"-"`
	CtyVal *cty.Value `diff:"-"`
}

func (r *AwsInstance) TerraformId() string {
	return r.Id
}

func (r *AwsInstance) TerraformType() string {
	return AwsInstanceResourceType
}

func (r *AwsInstance) CtyValue() *cty.Value {
	return r.CtyVal
}

var awsInstanceTags = map[string]string{}

func awsInstanceNormalizer(val *dctlcty.CtyAttributes) {
	val.SafeDelete([]string{"volume_tags"})
	val.SafeDelete([]string{"timeouts"})
}
