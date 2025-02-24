package veecp_addon

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVeecpAddons() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVeecpAddonsRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the addon.",
			},
			"pod_network_modes": {
				Type:     schema.TypeSet,
				Set:      schema.HashString,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Description: "The container network model, the value is `Flannel` or `VpcCniShared`. " +
					"Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. " +
					"VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.",
			},
			"deploy_modes": {
				Type:        schema.TypeSet,
				Set:         schema.HashString,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The deploy model, the value is `Managed` or `Unmanaged`.",
			},
			"deploy_node_types": {
				Type:        schema.TypeSet,
				Set:         schema.HashString,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The deploy node types, the value is `Node` or `VirtualNode`. Only effected when deploy_mode is `Unmanaged`.",
			},
			"necessaries": {
				Type:        schema.TypeSet,
				Set:         schema.HashString,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The necessaries of addons, the value is `Required` or `Recommended` or `OnDemand`.",
			},
			"categories": {
				Type:        schema.TypeSet,
				Set:         schema.HashString,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "The categories of addons, the value is `Storage` or `Network` or `Monitor` or `Scheduler` or `Dns` or `Security` or `Gpu` or `Image`.",
			},
			"kubernetes_versions": {
				Type:        schema.TypeSet,
				Set:         schema.HashString,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "A list of Kubernetes Versions.",
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
			"addons": {
				Description: "The collection of addons query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of addon.",
						},
						"versions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The version info of addon.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"version": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The basic version info.",
									},
									"compatible_versions": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The compatible version list.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"compatibilities": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The compatible version list.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"kubernetes_version": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Kubernetes Version of addon.",
												},
											},
										},
									},
								},
							},
						},
						"pod_network_modes": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "The network modes of pod.",
						},
						"deploy_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deploy model.",
						},
						"deploy_node_types": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "The deploy node types.",
						},
						"necessary": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The necessary of addon.",
						},
						"categories": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional:    true,
							Description: "The categories of addon.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVeecpAddonsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVeecpAddonService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVeecpAddons())
}
