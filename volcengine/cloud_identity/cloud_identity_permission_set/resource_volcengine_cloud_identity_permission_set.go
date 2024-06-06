package cloud_identity_permission_set

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CloudIdentityPermissionSet can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_identity_permission_set.default resource_id
```

*/

func ResourceVolcengineCloudIdentityPermissionSet() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCloudIdentityPermissionSetCreate,
		Read:   resourceVolcengineCloudIdentityPermissionSetRead,
		Update: resourceVolcengineCloudIdentityPermissionSetUpdate,
		Delete: resourceVolcengineCloudIdentityPermissionSetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the cloud identity permission set.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the cloud identity permission set.",
			},
			"relay_state": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The relay state of the cloud identity permission set.",
			},
			"session_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The session duration of the cloud identity permission set. Unit: second. Valid value range in 3600~43200.",
			},
			"permission_policies": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The policies of the cloud identity permission set.",
				Set:         PermissionPolicyHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"permission_policy_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"System", "Inline"}, false),
							Description:  "The type of the cloud identity permission set policy. Valid values: `System`, `Inline`.",
						},
						"permission_policy_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The name of the cloud identity permission set system policy. When the `permission_policy_type` is `System`, this field must be specified.",
						},
						"inline_policy_document": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The document of the cloud identity permission set inline policy. When the `permission_policy_type` is `Inline`, this field must be specified.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineCloudIdentityPermissionSetCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityPermissionSetService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCloudIdentityPermissionSet())
	if err != nil {
		return fmt.Errorf("error on creating cloud_identity_permission_set %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudIdentityPermissionSetRead(d, meta)
}

func resourceVolcengineCloudIdentityPermissionSetRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityPermissionSetService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCloudIdentityPermissionSet())
	if err != nil {
		return fmt.Errorf("error on reading cloud_identity_permission_set %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudIdentityPermissionSetUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityPermissionSetService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineCloudIdentityPermissionSet())
	if err != nil {
		return fmt.Errorf("error on updating cloud_identity_permission_set %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudIdentityPermissionSetRead(d, meta)
}

func resourceVolcengineCloudIdentityPermissionSetDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityPermissionSetService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCloudIdentityPermissionSet())
	if err != nil {
		return fmt.Errorf("error on deleting cloud_identity_permission_set %q, %s", d.Id(), err)
	}
	return err
}

func permissionPolicyHashBase(m map[string]interface{}) (buf bytes.Buffer) {
	policyType := strings.ToLower(m["permission_policy_type"].(string))
	//buf.WriteString(fmt.Sprintf("%s-", policyType))
	if policyType == "system" {
		policyName := strings.ToLower(m["permission_policy_name"].(string))
		buf.WriteString(fmt.Sprintf("%v#%v", policyType, policyName))
	} else if policyType == "inline" {
		policyDocument := strings.ToLower(m["inline_policy_document"].(string))
		buf.WriteString(fmt.Sprintf("%v#%v", policyType, policyDocument))
	}
	return buf
}

func PermissionPolicyHash(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	buf := permissionPolicyHashBase(m)
	return hashcode.String(buf.String())
}
