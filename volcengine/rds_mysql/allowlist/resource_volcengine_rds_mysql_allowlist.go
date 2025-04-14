package allowlist

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RDS AllowList can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mysql_allowlist.default acl-d1fd76693bd54e658912e7337d5b****
```

*/

func ResourceVolcengineRdsMysqlAllowlist() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRdsMysqlAllowlistCreate,
		Read:   resourceVolcengineRdsMysqlAllowlistRead,
		Update: resourceVolcengineRdsMysqlAllowlistUpdate,
		Delete: resourceVolcengineRdsMysqlAllowlistDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"allow_list_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the allow list.",
			},
			"allow_list_desc": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the allow list.",
			},
			"allow_list_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The type of IP address in the whitelist. Currently only IPv4 addresses are supported.",
			},
			"allow_list": {
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				AtLeastOneOf:  []string{"security_group_ids", "user_allow_list", "security_group_bind_infos"},
				ConflictsWith: []string{"security_group_ids", "security_group_bind_infos"},
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
				// 不可与安全组相关参数一起使用，否则会有漂移
				Description: "Enter an IP address or a range of IP addresses in CIDR format. Please note that if you want to use security group - related parameters, do not use this field. Instead, use the user_allow_list.",
			},
			"security_group_ids": {
				Type:          schema.TypeSet,
				Optional:      true,
				AtLeastOneOf:  []string{"allow_list", "user_allow_list", "security_group_bind_infos"},
				ConflictsWith: []string{"security_group_bind_infos", "allow_list"},
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The security group ids of the allow list.",
			},
			"security_group_bind_infos": {
				Type:     schema.TypeSet,
				Optional: true,
				//Computed:      true,
				ConflictsWith: []string{"security_group_ids", "allow_list"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bind_mode": {
							Type:     schema.TypeString,
							Required: true,
							Description: "The schema for the associated security group." +
								"\n IngressDirectionIp: Incoming Direction IP. \n AssociateEcsIp: Associate ECSIP. " +
								"\nexplain: In the CreateAllowList interface, SecurityGroupBindInfoObject BindMode and SecurityGroupId fields are required.",
						},
						"security_group_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The security group id of the allow list.",
						},
					},
				},
				Description: "Whitelist information for the associated security group.",
			},
			"allow_list_category": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "White list category. " +
					"Values:\nOrdinary: Ordinary white list.\n" +
					"Default: Default white list.\n " +
					"Description: When this parameter is used as a request parameter, the default value is Ordinary.",
			},
			"user_allow_list": {
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"allow_list"},
				AtLeastOneOf:  []string{"security_group_ids", "allow_list", "security_group_bind_infos"},
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "IP addresses outside the security group that need to be added to the whitelist." +
					" IP addresses or IP address segments in CIDR format can be entered. " +
					"Note: This field cannot be used simultaneously with AllowList.",
			},
			"allow_list_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the allow list.",
			},
		},
	}
}

func resourceVolcengineRdsMysqlAllowlistCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlAllowListService(meta.(*volc.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsMysqlAllowlist())
	if err != nil {
		return fmt.Errorf("error creating RDS Mysql Allowlist service: %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlAllowlistRead(d, meta)
}

func resourceVolcengineRdsMysqlAllowlistRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlAllowListService(meta.(*volc.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsMysqlAllowlist())
	if err != nil {
		return fmt.Errorf("error reading RDS Mysql Allowlist service: %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsMysqlAllowlistUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlAllowListService(meta.(*volc.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsMysqlAllowlist())
	if err != nil {
		return fmt.Errorf("error updating RDS Mysql Allowlist service: %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlAllowlistRead(d, meta)
}

func resourceVolcengineRdsMysqlAllowlistDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlAllowListService(meta.(*volc.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsMysqlAllowlist())
	if err != nil {
		return fmt.Errorf("error deleting RDS Mysql Allowlist service: %q, %w", d.Id(), err)
	}
	return err
}
