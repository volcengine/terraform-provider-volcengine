package registry

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCrRegistrys() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCrRegistrysRead,
		Schema: map[string]*schema.Schema{
			"names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The list of instance names.",
			},
			"types": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"Basic",
						"Enterprise",
					}, false),
				},
				Set:         schema.HashString,
				Description: "The list of instance types.",
			},
			"statuses": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The list of instance statuses.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"phase": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The phase of status.",
							ValidateFunc: validation.StringInSlice([]string{
								"Creating", "Running", "Stopped", "Starting", "Deleting", "Failed",
							}, false),
						},
						"conditions": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The condition of instance.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									"Ok", "Progressing", "Degraded", "Balance", "Released", "Unknown",
								}, false),
							},
						},
					},
				},
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of instance query.",
			},
			"registries": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of instance query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"registry": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of instance.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of instance.",
						},
						"charge_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge type of instance.",
						},
						"status": {
							Type:        schema.TypeSet,
							MaxItems:    1,
							Computed:    true,
							Description: "The status of instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"phase": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The phase status of instance.",
									},
									"conditions": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The condition of instance.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of instance.",
						},
					},
				},
			},
			"user": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "The collection of user.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The username of cr instance.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of user.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCrRegistrysRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCrRegistryService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCrRegistrys())
}
