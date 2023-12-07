package nas_permission_group

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Nas Permission Group can be imported using the id, e.g.
```
$ terraform import volcengine_nas_permission_group.default pgroup-1f85db2c****
```

*/

func ResourceVolcengineNasPermissionGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineNasPermissionGroupCreate,
		Read:   resourceVolcengineNasPermissionGroupRead,
		Update: resourceVolcengineNasPermissionGroupUpdate,
		Delete: resourceVolcengineNasPermissionGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"permission_group_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the permission group.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the permission group.",
			},
			"permission_rules": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The list of permissions rules.",
				Set:         permissionRuleHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr_ip": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Client IP addresses that are allowed access.",
						},
						"rw_mode": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"RW", "RO"}, false),
							Description:  "Permission group read and write rules. The value description is as follows:\n`RW`: Allows reading and writing.\n`RO`: read-only mode.",
						},
						"use_mode": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								"All_squash",
								"No_all_squash",
								"Root_squash",
								"No_root_squash",
							}, false),
							Description: "Permission group user permissions. The value description is as follows:\n`All_squash`: All access users are mapped to anonymous users or user groups.\n`No_all_squash`: The access user is first matched with the local user, and then mapped to an anonymous user or user group after the match fails.\n`Root_squash`: Map the Root user as an anonymous user or user group.\n`No_root_squash`: The Root user maintains the Root account authority.",
						},
					},
				},
			},
			"permission_group_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the permission group.",
			},
		},
	}
}

func resourceVolcengineNasPermissionGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineNasPermissionGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineNasPermissionGroup())
	if err != nil {
		return fmt.Errorf("error on creating volcengine Nas Permission Group: %q, %w", d.Id(), err)
	}
	return resourceVolcengineNasPermissionGroupRead(d, meta)
}

func resourceVolcengineNasPermissionGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineNasPermissionGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineNasPermissionGroup())
	if err != nil {
		return fmt.Errorf("error on reading volcengine Nas Permission Group: %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineNasPermissionGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineNasPermissionGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineNasPermissionGroup())
	if err != nil {
		return fmt.Errorf("error on updating volcengine Nas Permission Group: %q, %w", d.Id(), err)
	}
	return resourceVolcengineNasPermissionGroupRead(d, meta)
}

func resourceVolcengineNasPermissionGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVolcengineNasPermissionGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineNasPermissionGroup())
	if err != nil {
		return fmt.Errorf("error on deleting volcengine Nas Permission Group: %q, %w", d.Id(), err)
	}
	return nil
}

func permissionRuleHash(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v:", m["cidr_ip"]))
	buf.WriteString(fmt.Sprintf("%v:", m["rw_mode"]))
	buf.WriteString(fmt.Sprintf("%v:", m["use_mode"]))
	return hashcode.String(buf.String())
}
