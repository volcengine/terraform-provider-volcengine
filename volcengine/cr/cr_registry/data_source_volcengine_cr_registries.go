package cr_registry

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCrRegistries() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCrRegistriesRead,
		Schema: map[string]*schema.Schema{
			"names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The list of registry names.",
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
				Description: "The list of registry types.",
			},
			"statuses": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of registry statuses.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"phase": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "The phase of status.",
							ValidateFunc: validation.StringInSlice([]string{"Creating", "Running", "Stopped", "Starting", "Deleting", "Failed"}, false),
						},
						"condition": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "The condition of registry.",
							ValidateFunc: validation.StringInSlice([]string{"Ok", "Progressing", "Degraded", "Balance", "Released", "Unknown"}, false),
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
				Description: "The total count of registry query.",
			},
			"registries": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of registry query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of registry.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of registry.",
						},
						"charge_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The charge type of registry.",
						},
						"status": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "The status of registry.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"phase": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The phase status of registry.",
									},
									"conditions": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The condition of registry.",
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
							Description: "The creation time of registry.",
						},
						"username": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The username of cr instance.",
						},
						"user_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of user.",
						},
						"domains": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The domain of registry.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The domain of registry.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The domain type of registry.",
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

func dataSourceVolcengineCrRegistriesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCrRegistryService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCrRegistries())
}
