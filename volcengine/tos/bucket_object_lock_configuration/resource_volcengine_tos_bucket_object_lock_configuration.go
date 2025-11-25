package tos_bucket_object_lock_configuration

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TosBucketObjectLockConfiguration can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_object_lock_configuration.default bucket_name
```

*/

func ResourceVolcengineTosBucketObjectLockConfiguration() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTosBucketObjectLockConfigurationCreate,
		Read:   resourceVolcengineTosBucketObjectLockConfigurationRead,
		Update: resourceVolcengineTosBucketObjectLockConfigurationUpdate,
		Delete: resourceVolcengineTosBucketObjectLockConfigurationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the TOS bucket.",
			},
			"rule": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "The object lock rule configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_retention": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Required:    true,
							Description: "The default retention configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The default retention mode. Valid values: COMPLIANCE, GOVERNANCE.",
									},
									"days": {
										Type:          schema.TypeInt,
										Optional:      true,
										ConflictsWith: []string{"rule.default_retention.years"},
										Description:   "The number of days for the default retention period.",
									},
									"years": {
										Type:          schema.TypeInt,
										Optional:      true,
										ConflictsWith: []string{"rule.default_retention.days"},
										Description:   "The number of years for the default retention period.",
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

func resourceVolcengineTosBucketObjectLockConfigurationCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketObjectLockConfigurationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTosBucketObjectLockConfiguration())
	if err != nil {
		return fmt.Errorf("error on creating tos_bucket_object_lock_configuration %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketObjectLockConfigurationRead(d, meta)
}

func resourceVolcengineTosBucketObjectLockConfigurationRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketObjectLockConfigurationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTosBucketObjectLockConfiguration())
	if err != nil {
		return fmt.Errorf("error on reading tos_bucket_object_lock_configuration %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTosBucketObjectLockConfigurationUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketObjectLockConfigurationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTosBucketObjectLockConfiguration())
	if err != nil {
		return fmt.Errorf("error on updating tos_bucket_object_lock_configuration %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketObjectLockConfigurationRead(d, meta)
}

func resourceVolcengineTosBucketObjectLockConfigurationDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketObjectLockConfigurationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTosBucketObjectLockConfiguration())
	if err != nil {
		return fmt.Errorf("error on deleting tos_bucket_object_lock_configuration %q, %s", d.Id(), err)
	}
	return nil
}
