package cluster

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineBioosClusters() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineBioosClustersRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of cluster ids.",
			},
			"status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The status of the clusters.",
			},
			"type": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The type of the clusters.",
			},
			"public": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "whether it is a public cluster.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Vpc query.",
			},
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of cluster.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the cluster.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the cluster.",
						},
						"vke_config_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vke cluster id.",
						},
						"vke_config_storage_class": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the StorageClass that the vke cluster has installed.",
						},
						"external_config_wes_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The WES endpoint.",
						},
						"external_config_jupyterhub_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The endpoint of jupyterhub.",
						},
						"external_config_jupyterhub_jwt_secret": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The jupyterhub jwt secret.",
						},
						"external_config_resource_scheduler": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "External Resource Scheduler.",
						},
						"external_config_filesystem": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Workflow computing engine file system (currently supports tos, local).",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the bioos cluster.",
						},
						"start_time": {
							Computed:    true,
							Type:        schema.TypeInt,
							Description: "The start time of the cluster.",
						},
						"stopped_time": {
							Computed:    true,
							Type:        schema.TypeInt,
							Description: "The end time of the cluster.",
						},
						"bound": {
							Computed:    true,
							Type:        schema.TypeBool,
							Description: "Whether there is a bound workspace.",
						},
						"public": {
							Computed:    true,
							Type:        schema.TypeBool,
							Description: "whether it is a public cluster.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineBioosClustersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVolcengineBioosClusterService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineBioosClusters())
}
