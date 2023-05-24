package vpc_endpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcenginePrivatelinkVpcEndpoints() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcenginePrivatelinkVpcEndpointsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The IDs of vpc endpoint.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The vpc id of vpc endpoint.",
			},
			"endpoint_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of vpc endpoint.",
			},
			"service_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of vpc endpoint service.",
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Creating", "Pending", "Available", "Deleting", "Inactive"}, false),
				Description:  "The status of vpc endpoint. Valid values: `Creating`, `Pending`, `Available`, `Deleting`, `Inactive`.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of vpc endpoint.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Returns the total amount of the data list.",
			},
			"vpc_endpoints": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of vpc endpoint.",
						},
						"endpoint_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of vpc endpoint.",
						},
						"endpoint_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of vpc endpoint.",
						},
						"endpoint_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of vpc endpoint.",
						},
						"endpoint_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain of vpc endpoint.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of vpc endpoint.",
						},
						"business_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether the vpc endpoint is locked.",
						},
						"connection_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The connection  status of vpc endpoint.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of vpc endpoint.",
						},
						"service_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of vpc endpoint service.",
						},
						"service_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of vpc endpoint service.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc id of vpc endpoint.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of vpc endpoint.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of vpc endpoint.",
						},
						"deleted_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The delete time of vpc endpoint.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcenginePrivatelinkVpcEndpointsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVpcEndpointService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcenginePrivatelinkVpcEndpoints())
}
