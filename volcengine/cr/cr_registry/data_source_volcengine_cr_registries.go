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
				Description: "The list of registry names to query.",
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
				Description: "The list of registry types to query.",
			},
			"projects": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The list of project names to query.",
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
			"resource_tags": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The tags of cr registry.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Key of Tags.",
						},
						"values": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The Value of Tags.",
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
						"project": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ProjectName of the cr registry.",
						},
						"resource_tags": ve.TagsSchemaComputed(),
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
						"proxy_cache_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable proxy cache.",
						},
						"proxy_cache": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The proxy cache of registry. This field is valid when proxy_cache_enabled is true.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of proxy cache. Valid values: `DockerHub`, `DockerRegistry`.",
									},
									"endpoint": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The endpoint of proxy cache.",
									},
									"username": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The username of proxy cache.",
									},
									"skip_ssl_verify": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to skip ssl verify.",
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
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineCrRegistries())
}
