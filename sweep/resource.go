package sweep

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

type sweepResource struct {
	client       *ve.SdkClient
	resource     *schema.Resource
	resourceData *schema.ResourceData
}

func NewSweepResource(resource *schema.Resource, resourceData *schema.ResourceData, client *ve.SdkClient) *sweepResource {
	return &sweepResource{
		client:       client,
		resource:     resource,
		resourceData: resourceData,
	}
}

func (s *sweepResource) Delete() error {
	return s.resource.Delete(s.resourceData, s.client)
}

func (s *sweepResource) GetId() string {
	return s.resourceData.Id()
}
