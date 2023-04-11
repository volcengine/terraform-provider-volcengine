package vpc_endpoint_service

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcenginePrivatelinkVpcEndpointServices() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcenginePrivatelinkVpcEndpointServiceRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The IDs of vpc endpoint service.",
			},
			"service_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of vpc endpoint service.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of vpc endpoint service.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of vpc endpoint service.",
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
			"services": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of service.",
						},
						"service_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of service.",
						},
						"service_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of service.",
						},
						"service_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain of service.",
						},
						"service_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of service.",
						},
						"service_resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type of service.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of service.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of service.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of service.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of service.",
						},
						"auto_accept_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether auto accept node connect.",
						},
						"zone_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The IDs of zones.",
						},
						"resources": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The resources info.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of resource.",
									},
									"resource_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of resource.",
									},
									"zone_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The zone id of resource.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcenginePrivatelinkVpcEndpointServiceRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcenginePrivatelinkVpcEndpointServices())
}
