package cdn_shared_config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCdnSharedConfigs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCdnSharedConfigsRead,
		Schema: map[string]*schema.Schema{
			"config_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the shared config.",
			},
			"config_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the shared config.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the project.",
			},
			"config_type_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The config type list. The parameter value can be a combination of available values for ConfigType. " +
					"ConfigType and ConfigTypeList cannot be specified at the same time.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},
			"config_data": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the config.",
						},
						"config_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the config.",
						},
						"domain_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of domains.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the project.",
						},
						"update_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The update time of the shared config.",
						},
						"allow_ip_access_rule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The configuration for IP whitelist corresponds to ConfigType allow_ip_access_rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rules": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The entries in this list are an array of IP addresses and CIDR network segments. " +
											"The total number of entries cannot exceed 3,000. " +
											"The IP addresses and segments can be in IPv4 and IPv6 format. " +
											"Duplicate entries in the list will be removed and will not count towards the limit.",
									},
								},
							},
						},
						"deny_ip_access_rule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The configuration for IP blacklist is denoted by ConfigType deny_ip_access_rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rules": {
										Computed: true,
										Type:     schema.TypeList,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The entries in this list are an array of IP addresses and CIDR network segments. " +
											"The total number of entries cannot exceed 3,000. " +
											"The IP addresses and segments can be in IPv4 and IPv6 format. " +
											"Duplicate entries in the list will be removed and will not count towards the limit.",
									},
								},
							},
						},
						"allow_referer_access_rule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The configuration for the Referer whitelist corresponds to ConfigType allow_referer_access_rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow_empty": {
										Type:     schema.TypeBool,
										Computed: true,
										Description: "Indicates whether an empty Referer header, or a request without a Referer header, is not allowed. " +
											"Default is false.",
									},
									"common_type": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The content indicating the Referer whitelist.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ignore_case": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "This list is case-sensitive when matching requests. Default is true.",
												},
												"rules": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
													Description: "The entries in this list are an array of IP addresses and CIDR network segments. " +
														"The total number of entries cannot exceed 3,000. " +
														"The IP addresses and segments can be in IPv4 and IPv6 format. " +
														"Duplicate entries in the list will be removed and will not count towards the limit.",
												},
											},
										},
									},
								},
							},
						},
						"deny_referer_access_rule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The configuration for the Referer blacklist corresponds to ConfigType deny_referer_access_rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow_empty": {
										Type:     schema.TypeBool,
										Computed: true,
										Description: "Indicates whether an empty Referer header, or a request without a Referer header, is not allowed. " +
											"Default is false.",
									},
									"common_type": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The content indicating the Referer blacklist.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ignore_case": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "This list is case-sensitive when matching requests. Default is true.",
												},
												"rules": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
													Description: "The entries in this list are an array of IP addresses and CIDR network segments. " +
														"The total number of entries cannot exceed 3,000. " +
														"The IP addresses and segments can be in IPv4 and IPv6 format. " +
														"Duplicate entries in the list will be removed and will not count towards the limit.",
												},
											},
										},
									},
								},
							},
						},
						"common_match_list": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The configuration for a common list is represented by ConfigType common_match_list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"common_type": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The content indicating the Referer blacklist.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ignore_case": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "This list is case-sensitive when matching requests. Default is true.",
												},
												"rules": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
													Description: "The entries in this list are an array of IP addresses and CIDR network segments. " +
														"The total number of entries cannot exceed 3,000. " +
														"The IP addresses and segments can be in IPv4 and IPv6 format. " +
														"Duplicate entries in the list will be removed and will not count towards the limit.",
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

func dataSourceVolcengineCdnSharedConfigsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCdnSharedConfigService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCdnSharedConfigs())
}
