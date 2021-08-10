package resource

type ResourceType string

var supportedTypes = map[string]struct{}{
	"aws_ami":                               {},
	"aws_cloudfront_distribution":           {},
	"aws_db_instance":                       {},
	"aws_db_subnet_group":                   {},
	"aws_default_route_table":               {},
	"aws_default_security_group":            {},
	"aws_default_subnet":                    {},
	"aws_default_vpc":                       {},
	"aws_dynamodb_table":                    {},
	"aws_ebs_snapshot":                      {},
	"aws_ebs_volume":                        {},
	"aws_ecr_repository":                    {},
	"aws_eip":                               {},
	"aws_eip_association":                   {},
	"aws_iam_access_key":                    {},
	"aws_iam_policy":                        {},
	"aws_iam_policy_attachment":             {},
	"aws_iam_role":                          {},
	"aws_iam_role_policy":                   {},
	"aws_iam_role_policy_attachment":        {},
	"aws_iam_user":                          {},
	"aws_iam_user_policy":                   {},
	"aws_iam_user_policy_attachment":        {},
	"aws_instance":                          {},
	"aws_internet_gateway":                  {},
	"aws_key_pair":                          {},
	"aws_kms_alias":                         {},
	"aws_kms_key":                           {},
	"aws_lambda_event_source_mapping":       {},
	"aws_lambda_function":                   {},
	"aws_nat_gateway":                       {},
	"aws_route":                             {},
	"aws_route53_health_check":              {},
	"aws_route53_record":                    {},
	"aws_route53_zone":                      {},
	"aws_route_table":                       {},
	"aws_route_table_association":           {},
	"aws_s3_bucket":                         {},
	"aws_s3_bucket_analytics_configuration": {},
	"aws_s3_bucket_inventory":               {},
	"aws_s3_bucket_metric":                  {},
	"aws_s3_bucket_notification":            {},
	"aws_s3_bucket_policy":                  {},
	"aws_security_group":                    {},
	"aws_security_group_rule":               {},
	"aws_sns_topic":                         {},
	"aws_sns_topic_policy":                  {},
	"aws_sns_topic_subscription":            {},
	"aws_sqs_queue":                         {},
	"aws_sqs_queue_policy":                  {},
	"aws_subnet":                            {},
	"aws_vpc":                               {},

	"github_branch_protection": {},
	"github_membership":        {},
	"github_repository":        {},
	"github_team":              {},
	"github_team_membership":   {},

	"google_storage_bucket": {},
}

func IsResourceTypeSupported(ty string) bool {
	_, exist := supportedTypes[ty]
	return exist
}

func (ty ResourceType) String() string {
	return string(ty)
}
