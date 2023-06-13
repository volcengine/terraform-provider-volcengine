package host_group

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Tls Host Group can be imported using the id, e.g.
```
$ terraform import volcengine_tls_host_group.default edf052s21s*******dc15
```

*/

func ResourceVolcengineTlsHostGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTlsHostGroupCreate,
		Read:   resourceVolcengineTlsHostGroupRead,
		Update: resourceVolcengineTlsHostGroupUpdate,
		Delete: resourceVolcengineTlsHostGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"host_group_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of host group.",
			},
			"host_group_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of host group. The value can be IP or Label.",
			},
			"host_ip_list": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Set:         schema.HashString,
				Description: "The ip list of host group.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("host_group_type").(string) == "IP" {
						return false
					}
					return true
				},
			},
			"host_identifier": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The identifier of host.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("host_group_type").(string) == "Label" {
						return false
					}
					return true
				},
			},
			"auto_update": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether enable auto update.",
			},
			"update_start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return !d.Get("auto_update").(bool)
				},
				Description: "The update start time of log collector.",
			},
			"update_end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return !d.Get("auto_update").(bool)
				},
				Description: "The update end time of log collector.",
			},
			"service_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether enable service logging.",
			},
			"iam_project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The project name of iam.",
			},

			// computed data
			"host_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The count of host.",
			},
			"rule_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The rule count of host.",
			},
			"normal_heartbeat_status_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The normal heartbeat status count of host.",
			},
			"abnormal_heartbeat_status_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The abnormal heartbeat status count of host.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of host group.",
			},
			"modify_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The modify time of host group.",
			},
			"agent_latest_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The latest version of log collector.",
			},
		},
	}
	return resource
}

func resourceVolcengineTlsHostGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineTlsHostGroup())
	if err != nil {
		return fmt.Errorf("error on creating tls host group %q, %s", d.Id(), err)
	}
	return resourceVolcengineTlsHostGroupRead(d, meta)
}

func resourceVolcengineTlsHostGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineTlsHostGroup())
	if err != nil {
		return fmt.Errorf("error on updating tls host group %q, %s", d.Id(), err)
	}
	return resourceVolcengineTlsHostGroupRead(d, meta)
}

func resourceVolcengineTlsHostGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineTlsHostGroup())
	if err != nil {
		return fmt.Errorf("error on reading tls host group %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTlsHostGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineTlsHostGroup())
	if err != nil {
		return fmt.Errorf("error on deleting tls host group %q, %s", d.Id(), err)
	}
	return err
}
