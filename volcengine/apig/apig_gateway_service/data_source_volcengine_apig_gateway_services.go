package apig_gateway_service

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineApigGatewayServices() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineApigGatewayServicesRead,
		Schema: map[string]*schema.Schema{
			"gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The gateway id of api gateway service.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of api gateway service. This field support fuzzy query.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of api gateway service.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},
			"gateway_services": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of the api gateway service.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the api gateway service.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the api gateway service.",
						},
						"gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The gateway id of the api gateway service.",
						},
						"gateway_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The gateway name of the api gateway service.",
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The error message of the api gateway service.",
						},
						"comments": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The comments of the api gateway service.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the api gateway service.",
						},
						"protocol": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The protocol of the api gateway service.",
						},
						"auth_spec": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The auth spec of the api gateway service.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the api gateway service enable auth.",
									},
								},
							},
						},
						"domains": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The domains of the api gateway service.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The domain of the api gateway service.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the domain.",
									},
								},
							},
						},
						"custom_domains": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The custom domains of the api gateway service.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the custom domain.",
									},
									"domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The custom domain of the api gateway service.",
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

func dataSourceVolcengineApigGatewayServicesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewApigGatewayServiceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineApigGatewayServices())
}
