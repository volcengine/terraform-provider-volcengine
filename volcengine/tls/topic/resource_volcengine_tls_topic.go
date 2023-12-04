package topic

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.",
			},
			"shard_count": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The count of shards in the tls topic. Valid value range: 1-10.",
			},
			"auto_split": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				Description: "Whether to enable automatic partition splitting function of the tls topic.\n" +
					"true: (default) When the amount of data written exceeds the capacity of existing partitions for 5 consecutive minutes, " +
					"Log Service will automatically split partitions based on the data volume to meet business needs. " +
					"However, the number of partitions after splitting cannot exceed the maximum number of partitions. " +
					"Newly split partitions within the last 15 minutes will not be automatically split again.\n" +
					"false: Disables automatic partition splitting.",
			},
			"max_split_shard": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return !d.Get("auto_split").(bool)
				},
				Description: "The maximum number of partitions, which is the maximum number of partitions after partition splitting. " +
					"The value range is 1 to 10, with a default of 10.",
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
