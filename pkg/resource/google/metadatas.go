package google

import "github.com/cloudskiff/driftctl/pkg/resource"

func InitResourcesMetadata(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	initGoogleStorageBucketMetadata(resourceSchemaRepository)
}
