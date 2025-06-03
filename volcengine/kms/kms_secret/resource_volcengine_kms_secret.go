package kms_secret

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KmsSecret can be imported using the id, e.g.
```
$ terraform import volcengine_kms_secret.default resource_id
```

*/

func ResourceVolcengineKmsSecret() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsSecretCreate,
		Read:   resourceVolcengineKmsSecretRead,
		Update: resourceVolcengineKmsSecretUpdate,
		Delete: resourceVolcengineKmsSecretDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the secret.",
			},
			"secret_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the secret.",
			},
			"secret_value": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The value of the secret.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the secret.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the secret.",
			},
			"encryption_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The TRN of the KMS key used to encrypt the secret value.",
			},
			"extended_config": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The extended configurations of the secret.",
			},
			"automatic_rotation": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "The rotation state of the secret.",
			},
			"rotation_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The interval at which automatic rotation is performed.",
			},
			"creation_date": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The date when the secret was created.",
			},
			"update_date": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The date when the secret was updated.",
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The tenant ID of the secret.",
			},
			"trn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The information about the tenant resource name (TRN).",
			},
			"managed": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether the secret is hosted.",
			},
			"rotation_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The rotation state of the secret.",
			},
			"rotation_interval_second": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Rotation interval second.",
			},
			"last_rotation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The last time the secret was rotated.",
			},
			"schedule_rotation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The next time the secret will be rotated.",
			},
			"schedule_delete_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time when the secret will be deleted.",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of secret.",
			},
			"uuid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of secret.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsSecretCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsSecret())
	if err != nil {
		return fmt.Errorf("error on creating kms_secret %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsSecretRead(d, meta)
}

func resourceVolcengineKmsSecretRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsSecret())
	if err != nil {
		return fmt.Errorf("error on reading kms_secret %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsSecretUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineKmsSecret())
	if err != nil {
		return fmt.Errorf("error on updating kms_secret %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsSecretRead(d, meta)
}

func resourceVolcengineKmsSecretDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsSecretService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsSecret())
	if err != nil {
		return fmt.Errorf("error on deleting kms_secret %q, %s", d.Id(), err)
	}
	return err
}
