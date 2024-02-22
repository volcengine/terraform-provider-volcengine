package cdn_domain

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCdnDomains() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCdnDomainsRead,
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search by specifying domain name keywords, with fuzzy matching.",
			},
			"service_type": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The business type of the domain name is indicated by this parameter. " +
					"The possible values are: `download`: for file downloads. `web`: for web pages. " +
					"`video`: for audio and video on demand.",
			},
			"resource_tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Filter by specified domain name tags, up to 10 tags can be specified. " +
					"Each tag is entered as a string in the format of key:value.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of the domain.",
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of the domain.",
			},
			"origin_protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Configure the origin protocol for the accelerated domain.",
			},
			"ipv6": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: "Specify IPv6 configuration to filter accelerated domain names. " +
					"The optional values for this parameter are as follows: " +
					"`true`: Indicates that the accelerated domain name supports requests using IPv6 addresses." +
					"`false`: Indicates that the accelerated domain name does not support requests using IPv6 addresses.",
			},
			"https": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: "Specify HTTPS configuration to filter accelerated domains. " +
					"The optional values for this parameter are as follows: " +
					"`true`: Indicates that the accelerated domain has enabled HTTPS function." +
					"`false`: Indicates that the accelerated domain has not enabled HTTPS function.",
			},
			"primary_origin": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify a primary origin server for filtering accelerated domains.",
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
			"domains": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Search by specifying domain name keywords, with fuzzy matching.",
						},
						"service_type": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The business type of the domain name is indicated by this parameter. " +
								"The possible values are: `download`: for file downloads. `web`: for web pages. " +
								"`video`: for audio and video on demand.",
						},
						"resource_tags": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Indicate the tags you have set for this domain name. You can set up to 10 tags.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The key of the tag.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value of the tag.",
									},
								},
							},
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the domain.",
						},
						"project": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the domain.",
						},
						"origin_protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Configure the origin protocol for the accelerated domain.",
						},
						"cname": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CNAME address of the domain is automatically assigned when adding the domain.",
						},
						"service_region": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "Indicates the acceleration area. The parameter can take the following values: " +
								"`chinese_mainland`: Indicates mainland China. `global`: Indicates global." +
								" `outside_chinese_mainland`: Indicates global (excluding mainland China).",
						},
						"ipv6": {
							Type:     schema.TypeBool,
							Computed: true,
							Description: "Specify IPv6 configuration to filter accelerated domain names. " +
								"The optional values for this parameter are as follows: " +
								"`true`: Indicates that the accelerated domain name supports requests using IPv6 addresses." +
								"`false`: Indicates that the accelerated domain name does not support requests using IPv6 addresses.",
						},
						"https": {
							Type:     schema.TypeBool,
							Computed: true,
							Description: "Specify HTTPS configuration to filter accelerated domains. " +
								"The optional values for this parameter are as follows: " +
								"`true`: Indicates that the accelerated domain has enabled HTTPS function." +
								"`false`: Indicates that the accelerated domain has not enabled HTTPS function.",
						},
						"create_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The creation time of the domain.",
						},
						"update_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The update time of the domain.",
						},
						"primary_origin": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "List of primary source servers to accelerate the domain name.",
						},
						"backup_origin": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
							Description: "The list of backup origin servers for accelerating this domain name. " +
								"If no backup origin server is configured for this acceleration domain name, " +
								"the parameter value is null.",
						},
						"cache_shared": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "Indicates the role of the accelerated domain in the shared cache configuration. " +
								"This parameter can take the following values: " +
								"`target_host`: Indicates that there is a shared cache configuration where the role of the accelerated domain is the target domain." +
								"`cache_shared_on`: Indicates that there is a shared cache configuration where the role of the accelerated domain is the configured domain.`\"\"`: " +
								"This parameter value is empty, indicating that the accelerated domain does not exist in any shared cache configuration.",
						},
						"cache_shared_target_host": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "If CacheShared is cache_shared_on, " +
								"it means the target domain name that shares cache with the accelerated domain name. " +
								"If CacheShared is target_host or an empty value, the parameter value is empty.",
						},
						"is_conflict_domain": {
							Type:     schema.TypeBool,
							Computed: true,
							Description: "Indicates whether the accelerated domain name is a conflicting domain name. " +
								"By default, each accelerated domain name is unique in the content delivery network. " +
								"If you need to add an accelerated domain name that already exists in the content delivery network, " +
								"you need to submit a ticket. If the domain name is added successfully, " +
								"it becomes a conflicting domain name.",
						},
						"domain_lock": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"remark": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "If the Status is on, this parameter value records the reason for the lock.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Indicates whether the domain name is locked.",
									},
								},
							},
							Description: "Indicates the locked status of the accelerated domain.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCdnDomainsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCdnDomainService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCdnDomains())
}
