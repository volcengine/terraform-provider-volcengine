package kms_replicate_key

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*
Import
The KmsReplicateKey is not support imported.
*/
func ResourceVolcengineKmsReplicateKey() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKmsReplicateKeyCreate,
		Read:   resourceVolcengineKmsReplicateKeyRead,
		Delete: resourceVolcengineKmsReplicateKeyDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			// 注意：此处的参数在调用 ReplicateKey API 时，description和tags会作用于复制出来的密钥，即副本密钥
			// 由于复制出来的密钥的 keyring_name、key_name 和 key_id 与原密钥相同，只有region的区别；
			// 并且，UniversalClient 的签名地域来自 Provider region ，调用层没有暴露'请求级 region'。
			// 因此，继续对资源进行修改，会作用于原密钥，而不是副本密钥。故而Resource中的参数均设置为ForceNew属性
			"keyring_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The name of the keyring.",
			},
			"key_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				AtLeastOneOf: []string{"key_name", "key_id"},
				Description:  "The name of the key. Note: Only multi-region keys support replication.",
			},
			"key_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				AtLeastOneOf: []string{"key_name", "key_id"},
				Description:  "The id of the key. When key_id is not specified, both keyring_name and key_name must be specified.",
			},
			"replica_region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The target region for replica key.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The description of the replicated regional key.",
			},
			"tags": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Tags.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Key of Tags.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Value of Tags.",
						},
					},
				},
			},
			"replica_key_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the replica key.",
			},
		},
	}
	return resource
}

func resourceVolcengineKmsReplicateKeyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsReplicateKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKmsReplicateKey())
	if err != nil {
		return fmt.Errorf("error on creating kms_replicate_key %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsReplicateKeyRead(d, meta)
}

func resourceVolcengineKmsReplicateKeyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsReplicateKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKmsReplicateKey())
	if err != nil {
		return fmt.Errorf("error on reading kms_replicate_key %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKmsReplicateKeyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsReplicateKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineKmsReplicateKey())
	if err != nil {
		return fmt.Errorf("error on updating kms_replicate_key %q, %s", d.Id(), err)
	}
	return resourceVolcengineKmsReplicateKeyRead(d, meta)
}

func resourceVolcengineKmsReplicateKeyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKmsReplicateKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKmsReplicateKey())
	if err != nil {
		return fmt.Errorf("error on deleting kms_replicate_key %q, %s", d.Id(), err)
	}
	return err
}
