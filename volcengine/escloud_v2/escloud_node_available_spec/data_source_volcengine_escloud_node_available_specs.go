package escloud_node_available_spec

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineEscloudNodeAvailableSpecs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineEscloudNodeAvailableSpecssRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the instance.",
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
			"node_specs": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"configuration_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The configuration code.",
						},
						"az_available_specs_sold_out": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The available specs sold out.",
						},
						"network_specs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The network specs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"spec_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The spec name.",
									},
									"network_role": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The network role.",
									},
								},
							},
						},
						"node_available_specs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The node available specs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of node.",
									},
									"resource_spec_names": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The resource spec names of node.",
									},
									"storage_spec_names": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The storage spec names of node.",
									},
								},
							},
						},
						"resource_specs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The resource specs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of resource spec.",
									},
									"display_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The display name of resource spec.",
									},
									"memory": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The memory of resource spec. Unit: GiB.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of resource spec.",
									},
									"cpu": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The cpu of resource spec. Unit: Core.",
									},
								},
							},
						},
						"storage_specs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The storage specs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of storage spec.",
									},
									"display_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The display name of storage spec.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of storage spec.",
									},
									"max_size": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The max size of storage spec. Unit: GiB.",
									},
									"min_size": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The min size of storage spec. Unit: GiB.",
									},
									"size": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The size of storage spec.",
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

func dataSourceVolcengineEscloudNodeAvailableSpecssRead(d *schema.ResourceData, meta interface{}) error {
	service := NewEscloudNodeAvailableSpecsService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineEscloudNodeAvailableSpecs())
}
