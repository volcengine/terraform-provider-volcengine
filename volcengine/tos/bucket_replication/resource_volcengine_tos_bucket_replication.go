package tos_bucket_replication

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TosBucketReplication can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_replication.default bucket_name
```

*/

func ResourceVolcengineTosBucketReplication() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTosBucketReplicationCreate,
		Read:   resourceVolcengineTosBucketReplicationRead,
		Update: resourceVolcengineTosBucketReplicationUpdate,
		Delete: resourceVolcengineTosBucketReplicationDelete,
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
			"role": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IAM role for replication.",
			},
			"rules": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The replication rules of the bucket.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID of the replication rule.",
						},
						"status": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The status of the replication rule. Valid values: Enabled, Disabled.",
						},
						"prefix_set": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The prefix set for the replication rule.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"destination": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Required:    true,
							Description: "The destination configuration of the replication rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bucket": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The destination bucket name.",
									},
									"location": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The destination bucket location.",
									},
									"storage_class": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The storage class for the destination bucket. Valid values: STANDARD, IA, ARCHIVE, COLD_ARCHIVE.",
									},
									"storage_class_inherit_directive": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The storage class inherit directive. Valid values: COPY, OVERRIDE.",
									},
								},
							},
						},
						"historical_object_replication": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "Enabled",
							Description: "Whether to replicate historical objects. Valid values: Enabled, Disabled.",
						},
						"transfer_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "internal",
							Description: "Specify the data transmission link to be used for cross-regional replication. Valid values: internal, tos_acc.",
						},
						"access_control_translation": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Required:    true,
							Description: "The access control translation configuration of the replication rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"owner": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The owner of the destination object.",
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

func resourceVolcengineTosBucketReplicationCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketReplicationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTosBucketReplication())
	if err != nil {
		return fmt.Errorf("error on creating tos_bucket_replication %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketReplicationRead(d, meta)
}

func resourceVolcengineTosBucketReplicationRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketReplicationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTosBucketReplication())
	if err != nil {
		return fmt.Errorf("error on reading tos_bucket_replication %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTosBucketReplicationUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketReplicationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTosBucketReplication())
	if err != nil {
		return fmt.Errorf("error on updating tos_bucket_replication %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketReplicationRead(d, meta)
}

func resourceVolcengineTosBucketReplicationDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketReplicationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTosBucketReplication())
	if err != nil {
		return fmt.Errorf("error on deleting tos_bucket_replication %q, %s", d.Id(), err)
	}
	return nil
}
