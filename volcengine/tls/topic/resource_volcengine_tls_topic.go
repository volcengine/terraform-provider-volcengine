package topic

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Tls Topic can be imported using the id, e.g.
```
$ terraform import volcengine_tls_topic.default edf051ed-3c46-49ba-9339-bea628fe****
```

*/

func ResourceVolcengineTlsTopic() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTlsTopicCreate,
		Read:   resourceVolcengineTlsTopicRead,
		Update: resourceVolcengineTlsTopicUpdate,
		Delete: resourceVolcengineTlsTopicDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The project id of the tls topic.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the tls topic.",
			},
			"ttl": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 3650),
				Description:  "The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.",
			},
			"shard_count": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(1, 10),
				Description:  "The count of shards in the tls topic. Valid value range: 1-10.",
			},
			"auto_split": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable automatic partition splitting function of the tls topic.",
			},
			"max_split_shard": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(1, 10),
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return !d.Get("auto_split").(bool)
				},
				Description: "The max count of shards in the tls topic.",
			},
			"enable_tracking": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable WebTracking function of the tls topic.",
			},
			"time_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the time field.",
			},
			"time_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The format of the time field.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the tls project.",
			},
			"tags": ve.TagsSchema(),

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the tls topic.",
			},
			"modify_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The modify time of the tls topic.",
			},
		},
	}
	return resource
}

func resourceVolcengineTlsTopicCreate(d *schema.ResourceData, meta interface{}) (err error) {
	tlsTopicService := NewTlsTopicService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(tlsTopicService, d, ResourceVolcengineTlsTopic())
	if err != nil {
		return fmt.Errorf("error on creating TlsTopic %q, %w", d.Id(), err)
	}
	return resourceVolcengineTlsTopicRead(d, meta)
}

func resourceVolcengineTlsTopicRead(d *schema.ResourceData, meta interface{}) (err error) {
	tlsTopicService := NewTlsTopicService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(tlsTopicService, d, ResourceVolcengineTlsTopic())
	if err != nil {
		return fmt.Errorf("error on reading TlsTopic %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineTlsTopicUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	tlsTopicService := NewTlsTopicService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(tlsTopicService, d, ResourceVolcengineTlsTopic())
	if err != nil {
		return fmt.Errorf("error on updating TlsTopic %q, %w", d.Id(), err)
	}
	return resourceVolcengineTlsTopicRead(d, meta)
}

func resourceVolcengineTlsTopicDelete(d *schema.ResourceData, meta interface{}) (err error) {
	tlsTopicService := NewTlsTopicService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(tlsTopicService, d, ResourceVolcengineTlsTopic())
	if err != nil {
		return fmt.Errorf("error on deleting TlsTopic %q, %w", d.Id(), err)
	}
	return err
}
