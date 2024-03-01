package cdn_shared_config

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CdnSharedConfig can be imported using the id, e.g.
```
$ terraform import volcengine_cdn_shared_config.default resource_id
```

*/

func ResourceVolcengineCdnSharedConfig() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCdnSharedConfigCreate,
		Read:   resourceVolcengineCdnSharedConfigRead,
		Update: resourceVolcengineCdnSharedConfigUpdate,
		Delete: resourceVolcengineCdnSharedConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"config_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The name of the shared config. " +
					"The name cannot be the same as the name of an existing global configuration under the main account.",
			},
			"config_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The type of the shared config. " +
					"The type of global configuration. " +
					"The parameter can have the following values: " +
					"`deny_ip_access_rule`: represents IP blacklist. " +
					"`allow_ip_access_rule`: represents IP whitelist. " +
					"`deny_referer_access_rule`: represents Referer blacklist. " +
					"`allow_referer_access_rule`: represents Referer whitelist. " +
					"`common_match_list`: represents common list.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The ProjectName of the cdn shared config.",
			},
			"allow_ip_access_rule": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "The configuration for IP whitelist corresponds to ConfigType allow_ip_access_rule.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("config_type").(string) != "allow_ip_access_rule"
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rules": {
							Type:     schema.TypeSet,
							Set:      schema.HashString,
							Required: true,
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
				MaxItems:    1,
				Optional:    true,
				Description: "The configuration for IP blacklist is denoted by ConfigType deny_ip_access_rule.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("config_type").(string) != "deny_ip_access_rule"
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rules": {
							Type:     schema.TypeSet,
							Set:      schema.HashString,
							Required: true,
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
				MaxItems:    1,
				Optional:    true,
				Description: "The configuration for the Referer whitelist corresponds to ConfigType allow_referer_access_rule.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("config_type").(string) != "allow_referer_access_rule"
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_empty": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
							Description: "Indicates whether an empty Referer header, or a request without a Referer header, is not allowed. " +
								"Default is false.",
						},
						"common_type": {
							Type:        schema.TypeList,
							Required:    true,
							MaxItems:    1,
							Description: "The content indicating the Referer whitelist.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ignore_case": {
										Type:        schema.TypeBool,
										Default:     true,
										Optional:    true,
										Description: "This list is case-sensitive when matching requests. Default is true.",
									},
									"rules": {
										Type:     schema.TypeSet,
										Set:      schema.HashString,
										Required: true,
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
				MaxItems:    1,
				Optional:    true,
				Description: "The configuration for the Referer blacklist corresponds to ConfigType deny_referer_access_rule.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("config_type").(string) != "deny_referer_access_rule"
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_empty": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
							Description: "Indicates whether an empty Referer header, or a request without a Referer header, is not allowed. " +
								"Default is false.",
						},
						"common_type": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Required:    true,
							Description: "The content indicating the Referer blacklist.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ignore_case": {
										Type:        schema.TypeBool,
										Default:     true,
										Optional:    true,
										Description: "This list is case-sensitive when matching requests. Default is true.",
									},
									"rules": {
										Type:     schema.TypeSet,
										Set:      schema.HashString,
										Required: true,
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
				MaxItems:    1,
				Optional:    true,
				Description: "The configuration for a common list is represented by ConfigType common_match_list.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("config_type").(string) != "common_match_list"
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"common_type": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Required:    true,
							Description: "The content indicating the Referer blacklist.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ignore_case": {
										Type:        schema.TypeBool,
										Default:     true,
										Optional:    true,
										Description: "This list is case-sensitive when matching requests. Default is true.",
									},
									"rules": {
										Type:     schema.TypeSet,
										Set:      schema.HashString,
										Required: true,
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
	}
	return resource
}

func resourceVolcengineCdnSharedConfigCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCdnSharedConfigService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCdnSharedConfig())
	if err != nil {
		return fmt.Errorf("error on creating cdn_shared_config %q, %s", d.Id(), err)
	}
	return resourceVolcengineCdnSharedConfigRead(d, meta)
}

func resourceVolcengineCdnSharedConfigRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCdnSharedConfigService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCdnSharedConfig())
	if err != nil {
		return fmt.Errorf("error on reading cdn_shared_config %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCdnSharedConfigUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCdnSharedConfigService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineCdnSharedConfig())
	if err != nil {
		return fmt.Errorf("error on updating cdn_shared_config %q, %s", d.Id(), err)
	}
	return resourceVolcengineCdnSharedConfigRead(d, meta)
}

func resourceVolcengineCdnSharedConfigDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCdnSharedConfigService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCdnSharedConfig())
	if err != nil {
		return fmt.Errorf("error on deleting cdn_shared_config %q, %s", d.Id(), err)
	}
	return err
}
