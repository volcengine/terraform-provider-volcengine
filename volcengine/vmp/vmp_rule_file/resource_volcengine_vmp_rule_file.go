package vmp_rule_file

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VMP Rule File can be imported using the workspace_id:rule_file_id, e.g.
(We can only get rule file by WorkspaceId and RuleFileId)
```
$ terraform import volcengine_vmp_rule_file.default 60dde3ca-951c-4c05-8777-e5a7caa07ad6:d6f72bd9-674e-4651-b98c-3797657d9614
```

*/

func ResourceVolcengineVmpRuleFile() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVmpRuleFileCreate,
		Read:   resourceVolcengineVmpRuleFileRead,
		Update: resourceVolcengineVmpRuleFileUpdate,
		Delete: resourceVolcengineVmpRuleFileDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				var err error
				items := strings.Split(d.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{d}, fmt.Errorf("import id must be of the form workspace_id:rule_file_id")
				}
				err = d.Set("workspace_id", items[0])
				if err != nil {
					return []*schema.ResourceData{d}, err
				}
				err = d.Set("rule_file_id", items[1])
				if err != nil {
					return []*schema.ResourceData{d}, err
				}

				return []*schema.ResourceData{d}, nil
			},
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
				Description: "The name of the rule file.",
			},
			"workspace_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the workspace.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the rule file.",
			},
			"content": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The content of the rule file.",
			},
			"rule_file_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of rule file.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of workspace.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of workspace.",
			},
			"last_update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The last update time of rule file.",
			},
			"rule_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The rule count number of rule file.",
			},
		},
	}
	return resource
}

func resourceVolcengineVmpRuleFileCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineVmpRuleFile())
	if err != nil {
		return fmt.Errorf("error on creating rule file %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpRuleFileRead(d, meta)
}

func resourceVolcengineVmpRuleFileRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineVmpRuleFile())
	if err != nil {
		return fmt.Errorf("error on reading rule file %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVmpRuleFileUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineVmpRuleFile())
	if err != nil {
		return fmt.Errorf("error on updating rule file %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpRuleFileRead(d, meta)
}

func resourceVolcengineVmpRuleFileDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineVmpRuleFile())
	if err != nil {
		return fmt.Errorf("error on deleting rule file %q, %s", d.Id(), err)
	}
	return err
}
