package apig_gateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineApigGateways() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineApigGatewaysRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of api gateway IDs.",
			},
			"vpc_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of vpc IDs.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of api gateway. This field support fuzzy query.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of api gateway.",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of api gateway.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of api gateway.",
			},
			"tags": ve.TagsSchema(),
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
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
			"gateways": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of the api gateway.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the api gateway.",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region of the api gateway.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the api gateway.",
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The error message of the api gateway.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the api gateway.",
						},
						"comments": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The comments of the api gateway.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of the api gateway.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the api gateway.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the api gateway.",
						},
						"tags": ve.TagsSchemaComputed(),
						"network_spec": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The network spec of the api gateway.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpc_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The vpc id of the api gateway.",
									},
									"subnet_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The subnet ids of the api gateway.",
									},
								},
							},
						},
						"backend_spec": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The backend spec of the api gateway.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_vke_with_flannel_cni_supported": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the api gateway support vke flannel cni.",
									},
									"vke_pod_cidr": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The vke pod cidr of the api gateway.",
									},
								},
							},
						},
						"monitor_spec": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The monitor spec of the api gateway.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the api gateway enable monitor.",
									},
									"workspace_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The workspace id of the monitor.",
									},
								},
							},
						},
						"log_spec": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The log spec of the api gateway.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the api gateway enable tls log.",
									},
									"project_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The project id of the tls.",
									},
									"topic_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The topic id of the tls.",
									},
								},
							},
						},
						"resource_spec": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The resource spec of the api gateway.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"replicas": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The replicas of the resource spec.",
									},
									"instance_spec_code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The instance spec code of the resource spec.",
									},
									"clb_spec_code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The clb spec code of the resource spec.",
									},
									"public_network_billing_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The public network billing type of the resource spec.",
									},
									"public_network_bandwidth": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The public network bandwidth of the resource spec.",
									},
									"network_type": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The network type of the api gateway.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enable_public_network": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether the api gateway enable public network.",
												},
												"enable_private_network": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether the api gateway enable private network.",
												},
											},
										},
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

func dataSourceVolcengineApigGatewaysRead(d *schema.ResourceData, meta interface{}) error {
	service := NewApigGatewayService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineApigGateways())
}
