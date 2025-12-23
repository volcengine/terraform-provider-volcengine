package rds_postgresql_allowlist

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsPostgresqlAllowlist can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_allowlist.default resource_id
```

*/

func ResourceVolcengineRdsPostgresqlAllowlist() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlAllowlistCreate,
		Read:   resourceVolcengineRdsPostgresqlAllowlistRead,
		Update: resourceVolcengineRdsPostgresqlAllowlistUpdate,
		Delete: resourceVolcengineRdsPostgresqlAllowlistDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_ids": {
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				MinItems:      1,
				MaxItems:      300,
				ConflictsWith: []string{"allow_list", "user_allow_list", "security_group_bind_infos"},
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "IDs of PostgreSQL instances to unify allowlists. When set, creation uses UnifyNewAllowList to merge existing instance allowlists into a new one. " +
					"Supports merging and generating allowlists of up to 300 instances.",
			},
			"allow_list_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the postgresql allow list.",
			},
			"allow_list_desc": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the postgresql allow list.",
			},
			"allow_list_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"IPv4"}, false),
				Description:  "The type of IP address in the whitelist. Currently only `IPv4` addresses are supported.",
			},
			"allow_list": {
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				MinItems:      1,
				MaxItems:      300,
				Set:           schema.HashString,
				ConflictsWith: []string{"user_allow_list"},
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Enter an IP address or a range of IP addresses in CIDR format. This field cannot be used together with the user_allow_list field.",
			},
			"allow_list_category": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"Default", "Ordinary"}, false),
				Description: "The category of the allow list. Valid values: Ordinary, Default. " +
					"When this parameter is used as a request parameter, there is no default value.",
			},
			"security_group_bind_infos": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_list": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "IP addresses in the security group.",
						},
						"bind_mode": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"IngressDirectionIp", "AssociateEcsIp"}, false),
							Description:  "The binding mode of the security group. Valid values: IngressDirectionIp, AssociateEcsIp.",
						},
						"security_group_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ID of the security group.",
						},
						"security_group_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the security group.",
						},
					},
				},
				Description: "The information of security groups to bind with the allow list.",
			},
			"user_allow_list": {
				Type:          schema.TypeSet,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"allow_list"},
				Set:           schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "IP addresses outside security groups to be added to the allowlist. Cannot be used with allow_list.",
			},
			"update_security_group": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to update the security groups bound to the allowlist when modifying.",
			},

			// computed fields
			"associated_instance_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of instances bound under the whitelist.",
			},
			"associated_instances": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of postgresql instances.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the postgresql instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the postgresql instance.",
						},
						"vpc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vpc.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlAllowlistCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlAllowlistService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlAllowlist())
	if err != nil {
		return fmt.Errorf("error on creating rds_postgresql_allowlist %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlAllowlistRead(d, meta)
}

func resourceVolcengineRdsPostgresqlAllowlistRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlAllowlistService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsPostgresqlAllowlist())
	if err != nil {
		return fmt.Errorf("error on reading rds_postgresql_allowlist %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsPostgresqlAllowlistUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlAllowlistService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsPostgresqlAllowlist())
	if err != nil {
		return fmt.Errorf("error on updating rds_postgresql_allowlist %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlAllowlistRead(d, meta)
}

func resourceVolcengineRdsPostgresqlAllowlistDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlAllowlistService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsPostgresqlAllowlist())
	if err != nil {
		return fmt.Errorf("error on deleting rds_postgresql_allowlist %q, %s", d.Id(), err)
	}
	return err
}
