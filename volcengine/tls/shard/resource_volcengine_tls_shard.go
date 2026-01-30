package shard

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TlsShard can not be imported using the id

*/

func ResourceVolcengineTlsShard() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineTlsShardCreate,
		Read:   resourceVolcengineTlsShardRead,
		Update: resourceVolcengineTlsShardUpdate,
		Delete: resourceVolcengineTlsShardDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the topic.",
			},
			"shard_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the shard to split.",
			},
			"number": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The number of splits. Must be a non-zero even number, such as 2, 4, 8, or 16.",
			},
			"shards": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of shards after split.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the topic.",
						},
						"shard_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the shard.",
						},
						"inclusive_begin_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The inclusive begin key of the shard.",
						},
						"exclusive_end_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The exclusive end key of the shard.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the shard.",
						},
						"modify_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modification time of the shard.",
						},
						"stop_write_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The stop write time of the shard.",
						},
					},
				},
			},
		},
	}
}

func resourceVolcengineTlsShardCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Create(service, d, ResourceVolcengineTlsShard())
}

func resourceVolcengineTlsShardRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Read(service, d, ResourceVolcengineTlsShard())
}

func resourceVolcengineTlsShardUpdate(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Update(service, d, ResourceVolcengineTlsShard())
}

func resourceVolcengineTlsShardDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineTlsShard())
}
