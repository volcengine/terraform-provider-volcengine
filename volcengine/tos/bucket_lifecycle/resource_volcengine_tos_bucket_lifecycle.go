package tos_bucket_lifecycle

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TosBucketLifecycle can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_lifecycle.default bucket_name
```

*/

func ResourceVolcengineTosBucketLifecycle() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTosBucketLifecycleCreate,
		Read:   resourceVolcengineTosBucketLifecycleRead,
		Update: resourceVolcengineTosBucketLifecycleUpdate,
		Delete: resourceVolcengineTosBucketLifecycleDelete,
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
			"rules": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The lifecycle rules of the bucket.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The ID of the lifecycle rule.",
						},
						"prefix": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The prefix of the lifecycle rule.",
						},
						"status": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The status of the lifecycle rule. Valid values: Enabled, Disabled.",
						},
						"expiration": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The expiration configuration of the lifecycle rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"days": {
										Type:          schema.TypeInt,
										Optional:      true,
										ConflictsWith: []string{"rules.expiration.date"},
										Description:   "The number of days after object creation when the rule takes effect.",
									},
									"date": {
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"rules.expiration.days"},
										Description:   "The date when the rule takes effect. Format: 2023-01-01T00:00:00.000Z.",
									},
								},
							},
						},
						"tags": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The tag filters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The key of the tag.",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The value of the tag.",
									},
								},
							},
						},
						"transitions": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The transition configuration of the lifecycle rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"days": {
										Type:          schema.TypeInt,
										Optional:      true,
										ConflictsWith: []string{"rules.transitions.date"},
										Description:   "The number of days after object creation when the transition takes effect.",
									},
									"date": {
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"rules.transitions.days"},
										Description:   "The date when the transition takes effect. Format: 2023-01-01T00:00:00.000Z.",
									},
									"storage_class": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The storage class to transition to. Valid values: IA, ARCHIVE, COLD_ARCHIVE.",
									},
								},
							},
						},
						"non_current_version_expiration": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The non-current version expiration configuration of the lifecycle rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"non_current_days": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The number of days after object creation when the non-current version expiration takes effect.",
									},
								},
							},
						},
						"non_current_version_transitions": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The non-current version transition configuration of the lifecycle rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"non_current_days": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The number of days after object creation when the non-current version transition takes effect.",
									},
									"storage_class": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The storage class to transition to. Valid values: IA, ARCHIVE, COLD_ARCHIVE.",
									},
								},
							},
						},
						"filter": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The filter configuration of the lifecycle rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"object_size_greater_than": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The minimum object size in bytes for the rule to apply.",
									},
									"greater_than_include_equal": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Whether to enable equal conditions. The value can only be "Enabled" or "Disabled". If not configured, it will default to "Disabled".`,
									},
									"object_size_less_than": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum object size in bytes for the rule to apply.",
									},
									"less_than_include_equal": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `Whether to enable equal conditions. The value can only be "Enabled" or "Disabled". If not configured, it will default to "Disabled".`,
									},
								},
							},
						},
						"abort_incomplete_multipart_upload": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The abort incomplete multipart upload configuration of the lifecycle rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"days_after_initiation": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The number of days after initiation when the incomplete multipart upload should be aborted.",
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

func resourceVolcengineTosBucketLifecycleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketLifecycleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTosBucketLifecycle())
	if err != nil {
		return fmt.Errorf("error on creating tos_bucket_lifecycle %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketLifecycleRead(d, meta)
}

func resourceVolcengineTosBucketLifecycleRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketLifecycleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTosBucketLifecycle())
	if err != nil {
		return fmt.Errorf("error on reading tos_bucket_lifecycle %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTosBucketLifecycleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketLifecycleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTosBucketLifecycle())
	if err != nil {
		return fmt.Errorf("error on updating tos_bucket_lifecycle %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketLifecycleRead(d, meta)
}

func resourceVolcengineTosBucketLifecycleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketLifecycleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTosBucketLifecycle())
	if err != nil {
		return fmt.Errorf("error on deleting tos_bucket_lifecycle %q, %s", d.Id(), err)
	}
	return nil
}
