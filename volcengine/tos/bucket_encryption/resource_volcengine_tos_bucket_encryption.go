package tos_bucket_encryption

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TosBucketEncryption can be imported using the id, e.g.
```
$ terraform import volcengine_tos_bucket_encryption.default resource_id
```

*/

func ResourceVolcengineTosBucketEncryption() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTosBucketEncryptionCreate,
		Read:   resourceVolcengineTosBucketEncryptionRead,
		Update: resourceVolcengineTosBucketEncryptionUpdate,
		Delete: resourceVolcengineTosBucketEncryptionDelete,
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
				Description: "The name of the bucket.",
			},
			"rule": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "The rule of the bucket encryption.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"apply_server_side_encryption_by_default": {
							Type:        schema.TypeList,
							Required:    true,
							MaxItems:    1,
							Description: "The server side encryption configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sse_algorithm": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The server side encryption algorithm. Valid values: `kms`, `AES256`, `SM4`.",
									},
									"kms_data_encryption": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
											return d.Get("rule.0.apply_server_side_encryption_by_default.0.sse_algorithm") != "kms"
										},
										Description: "The kms data encryption. Valid values: `AES256`, `SM4`. Default is `AES256`.",
									},
									"kms_master_key_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
											return d.Get("rule.0.apply_server_side_encryption_by_default.0.sse_algorithm") != "kms"
										},
										Description: "The kms master key id. This field is required when `sse_algorithm` is `kms`. The format is `trn:kms:<region>:<accountID>:keyrings/<keyring>/keys/<key>`.",
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

func resourceVolcengineTosBucketEncryptionCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketEncryptionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTosBucketEncryption())
	if err != nil {
		return fmt.Errorf("error on creating tos_bucket_encryption %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketEncryptionRead(d, meta)
}

func resourceVolcengineTosBucketEncryptionRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketEncryptionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTosBucketEncryption())
	if err != nil {
		return fmt.Errorf("error on reading tos_bucket_encryption %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTosBucketEncryptionUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketEncryptionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTosBucketEncryption())
	if err != nil {
		return fmt.Errorf("error on updating tos_bucket_encryption %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketEncryptionRead(d, meta)
}

func resourceVolcengineTosBucketEncryptionDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketEncryptionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTosBucketEncryption())
	if err != nil {
		return fmt.Errorf("error on deleting tos_bucket_encryption %q, %s", d.Id(), err)
	}
	return err
}
