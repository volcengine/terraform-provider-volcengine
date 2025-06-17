package waf_domain

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
WafDomain can be imported using the id, e.g.
```
$ terraform import volcengine_waf_domain.default resource_id
```

*/

func ResourceVolcengineWafDomain() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineWafDomainCreate,
		Read:   resourceVolcengineWafDomainRead,
		Update: resourceVolcengineWafDomainUpdate,
		Delete: resourceVolcengineWafDomainDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"access_mode": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Access mode.",
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "List of domain names that need to be protected by WAF.",
			},
			"client_ip_location": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The method of obtaining the client IP.",
			},
			"cloud_access_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Access port information.If AccessMode is Alb/CLB, this field is required.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// waf 不生效 暂时忽略
						//"defence_mode": {
						//	Type:        schema.TypeInt,
						//	Optional:    true,
						//	Description: "The protection mode of the instance. Works only on modified scenes.",
						//	DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
						//		// 创建时不存在这个参数，修改时存在这个参数
						//		return d.Id() == ""
						//	},
						//},
						"instance_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The name of instance. Works only on modified scenes.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								// 创建时不存在这个参数，修改时存在这个参数
								return d.Id() == ""
							},
						},
						"lost_association_from_alb": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Whether the instance is unbound from the alb and is unbound on the ALB side. Works only on modified scenes.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								// 创建时不存在这个参数，修改时存在这个参数
								return d.Id() == ""
							},
						},
						"instance_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ID of instance.",
						},
						"listener_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The ID of listener.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The type of Listener protocol.",
						},
						"port": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The port number corresponding to the listener.",
						},
						"access_protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The access protocol needs to be consistent with the monitoring protocol.",
						},
					},
				},
			},
			"protocols": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Access protocol types.",
			},
			"protocol_ports": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "Access port information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
							Set:         schema.HashInt,
							Description: "Ports supported by the HTTP protocol.",
						},
						"https": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
							Set:         schema.HashInt,
							Description: "Ports supported by the HTTPs protocol.",
						},
					},
				},
			},
			"enable_http2": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable HTTP 2.0.",
			},
			"protocol_follow": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable protocol following.",
			},
			"enable_ipv6": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether it supports protecting IPv6 requests.",
			},
			"certificate_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "When the protocol type is HTTPS, the bound certificate ID needs to be entered.",
			},
			"tls_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the log service.",
			},
			"proxy_config": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable proxy configuration.",
			},
			"ssl_protocols": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "TLS protocol version.",
			},
			"ssl_ciphers": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Encryption kit.",
			},
			"keep_alive_time_out": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Long connection retention time.",
			},
			"keep_alive_request": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The number of long connection multiplexes.",
			},
			"client_max_body_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The client requests the maximum value of body.",
			},
			"lb_algorithm": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The types of load balancing algorithms.",
			},
			"public_real_server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Connect to the source return mode.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ID of vpc.",
			},
			"backend_groups": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "The configuration of source station.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_port": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Set:      schema.HashInt,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
							Description: "Access port number.",
						},
						"backends": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "The details of the source station group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The agreement of Source Station.",
									},
									"ip": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Source station IP address.",
									},
									"port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Source station port number.",
									},
									"weight": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "The weight of the source station rules.",
									},
								},
							},
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Source station group name.",
						},
					},
				},
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of project. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"redirect_https": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "When only the HTTPS protocol is enabled, whether to redirect HTTP requests to HTTPS. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"proxy_write_time_out": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The timeout period during which the WAF transmits the request to the backend server.",
			},
			"proxy_read_time_out": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The timeout period during which WAF reads the response from the backend server.",
			},
			"proxy_retry": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The number of retries for WAF back to source.",
			},
			"proxy_keep_alive_time_out": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Idle long connection timeout period.",
			},
			"proxy_keep_alive": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The number of reusable WAF origin long connections.",
			},
			"custom_header": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Custom Header.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"proxy_connect_time_out": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The timeout period for establishing a connection between the WAF and the backend server.",
			},
			"volc_certificate_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "When the protocol type is HTTPS, the bound certificate ID needs to be entered. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"certificate_platform": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Certificate custody platform. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"custom_sni": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Custom SNI needs to be configured when EnableSNI=1. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"enable_custom_redirect": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Whether to enable user-defined redirection. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"enable_sni": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Whether to enable the SNI configuration. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"llm_available": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Is LLM available. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"tls_fields_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Details of log field configuration. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"headers_config": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The configuration of Headers. Works only on modified scenes.",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								// 创建时不存在这个参数，修改时存在这个参数
								return d.Id() == ""
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:     schema.TypeInt,
										Optional: true,
										DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
											// 创建时不存在这个参数，修改时存在这个参数
											return d.Id() == ""
										},
										Description: "Whether the log contains this field. Works only on modified scenes.",
									},
									"excluded_key_list": {
										Type:     schema.TypeSet,
										Optional: true,
										Set:      schema.HashString,
										DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
											// 创建时不存在这个参数，修改时存在这个参数
											return d.Id() == ""
										},
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "For the use of composite fields, " +
											"exclude the fields in the keyword list from the JSON of the fields. Works only on modified scenes.",
									},
									"statistical_key_list": {
										Type:     schema.TypeSet,
										Optional: true,
										Set:      schema.HashString,
										DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
											// 创建时不存在这个参数，修改时存在这个参数
											return d.Id() == ""
										},
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "Create statistical indexes for the fields of the list. Works only on modified scenes.",
									},
								},
							},
						},
					},
				},
			},
			// 防护网站开关相关参数
			"bot_repeat_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Whether to enable the bot frequency limit policy. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"bot_dytoken_enable": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Whether to enable the bot dynamic token. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"auto_cc_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the intelligent CC protection strategy. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"bot_sequence_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Whether to enable the bot behavior map. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"bot_sequence_default_action": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Set the default actions of the bot behavior map strategy. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"bot_frequency_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Whether to enable the bot frequency limit policy. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"waf_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the vulnerability protection strategy. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"cc_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the CC protection policy. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"white_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the access list policy. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"black_ip_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the access ban list policy. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"black_lct_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the geographical location access control policy. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"waf_white_req_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the whitening strategy for vulnerability protection requests. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"white_field_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the whitening strategy for vulnerability protection fields. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"custom_rsp_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the custom response interception policy. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"system_bot_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the managed Bot classification strategy. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"custom_bot_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the custom Bot classification strategy. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"api_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the API protection policy. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"tamper_proof_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the page tamper-proof policy. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			"dlp_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to activate the strategy for preventing the leakage of sensitive information. Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
			},
			// 更新域名防护模式
			"extra_defence_mode_lb_instance": {
				Type:     schema.TypeList,
				Optional: true,
				Description: "The protection mode of the exception instance. " +
					"It takes effect when the access mode is accessed through an application load balancing (ALB) instance (AccessMode=20)." +
					" Works only on modified scenes.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"defence_mode": {
							Type:     schema.TypeInt,
							Optional: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								// 创建时不存在这个参数，修改时存在这个参数
								return d.Id() == ""
							},
							Description: "Set the protection mode for exceptional ALB instances. Works only on modified scenes.",
						},
						"instance_id": {
							Type:     schema.TypeString,
							Optional: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								// 创建时不存在这个参数，修改时存在这个参数
								return d.Id() == ""
							},
							Description: "The Id of ALB instance. Works only on modified scenes.",
						},
					},
				},
			},
			"defence_mode": {
				Type:     schema.TypeInt,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
				Description: "The protection mode of the instance. Works only on modified scenes.",
			},
			"defence_mode_computed": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The protection mode of the instance.",
			},
			"cname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CNAME value generated by the WAF instance.",
			},
			"server_ips": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP of the WAF protection instance.",
			},
			"advanced_defense_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "High-defense instance IP.",
			},
			"advanced_defense_ipv6": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "High-defense instance IPv6.",
			},
			"certificate_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the certificate.",
			},
			"attack_status": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The status of the attack.",
			},
			"status": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The status of access.",
			},
			"src_ips": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "WAF source IP.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time.",
			},
		},
	}
	return resource
}

func resourceVolcengineWafDomainCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafDomainService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineWafDomain())
	if err != nil {
		return fmt.Errorf("error on creating waf_domain %q, %s", d.Id(), err)
	}
	return resourceVolcengineWafDomainRead(d, meta)
}

func resourceVolcengineWafDomainRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafDomainService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineWafDomain())
	if err != nil {
		return fmt.Errorf("error on reading waf_domain %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineWafDomainUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafDomainService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineWafDomain())
	if err != nil {
		return fmt.Errorf("error on updating waf_domain %q, %s", d.Id(), err)
	}
	return resourceVolcengineWafDomainRead(d, meta)
}

func resourceVolcengineWafDomainDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafDomainService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineWafDomain())
	if err != nil {
		return fmt.Errorf("error on deleting waf_domain %q, %s", d.Id(), err)
	}
	return err
}
