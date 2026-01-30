package trace_instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
tls trace instance can be imported using the id, e.g.
```
$ terraform import volcengine_tls_trace_instance.default instance-1234567890
```

*/

func ResourceVolcengineTlsTraceInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineTlsTraceInstanceCreate,
		Read:   resourceVolcengineTlsTraceInstanceRead,
		Update: resourceVolcengineTlsTraceInstanceUpdate,
		Delete: resourceVolcengineTlsTraceInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the project.",
			},
			"trace_instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the trace instance.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the trace instance.",
			},
			"backend_config": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The backend config of the trace instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ttl": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Total log retention time in days.",
						},
						"enable_hot_ttl": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to enable tiered storage.",
						},
						"hot_ttl": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Standard storage duration in days.",
						},
						"cold_ttl": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Infrequent storage duration in days.",
						},
						"archive_ttl": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Archive storage duration in days.",
						},
						"auto_split": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to enable auto split.",
						},
						"max_split_partitions": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Max split partitions.",
						},
					},
				},
			},
		},
	}
}

func resourceVolcengineTlsTraceInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsTraceInstanceService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Create(service, d, ResourceVolcengineTlsTraceInstance())
}

func resourceVolcengineTlsTraceInstanceRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsTraceInstanceService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Read(service, d, ResourceVolcengineTlsTraceInstance())
}

func resourceVolcengineTlsTraceInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsTraceInstanceService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Update(service, d, ResourceVolcengineTlsTraceInstance())
}

func resourceVolcengineTlsTraceInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsTraceInstanceService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineTlsTraceInstance())
}
