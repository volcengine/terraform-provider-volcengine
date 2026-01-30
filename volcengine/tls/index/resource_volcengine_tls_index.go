package index

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Tls Index can be imported using the topic id, e.g.
```
$ terraform import volcengine_tls_index.default index:edf051ed-3c46-49ba-9339-bea628fe****
```

*/

func ResourceVolcengineTlsIndex() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTlsIndexCreate,
		Read:   resourceVolcengineTlsIndexRead,
		Update: resourceVolcengineTlsIndexUpdate,
		Delete: resourceVolcengineTlsIndexDelete,
		Importer: &schema.ResourceImporter{
			State: tlsIndexImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The topic id of the tls index.",
			},
			"full_text": {
				Type:         schema.TypeList,
				Optional:     true,
				MaxItems:     1,
				AtLeastOneOf: []string{"full_text", "key_value"},
				Description:  "The full text info of the tls index.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"case_sensitive": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether the FullTextInfo is case sensitive.",
						},
						"delimiter": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The delimiter of the FullTextInfo.",
						},
						"include_chinese": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether the FullTextInfo include chinese.",
						},
					},
				},
			},
			"key_value": {
				Type:         schema.TypeSet,
				Optional:     true,
				AtLeastOneOf: []string{"full_text", "key_value"},
				Description:  "The key value info of the tls index.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The key of the KeyValueInfo.",
						},
						"value_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of value. Valid values: `long`, `double`, `text`, `json`.",
						},
						"case_sensitive": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether the value is case sensitive.",
						},
						"delimiter": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The delimiter of the value.",
						},
						"include_chinese": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether the value include chinese.",
						},
						"sql_flag": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether the filed is enabled for analysis.",
						},
						"index_all": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether to create indexes for all fields in JSON fields with text values. This field is valid when the `value_type` is `json`.",
						},
						"json_keys": {
							Type:     schema.TypeSet,
							Optional: true,
							//DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
							//	logger.DebugInfo("testValueType1", k)
							//	items := strings.Split(k, ".")
							//	if len(items) > 2 {
							//		key := strings.Join(items[:2], ".")
							//		logger.DebugInfo("testValueType2", key, d.Get(key+".value_type"))
							//		if valueType := d.Get(key + ".value_type"); valueType != nil && valueType.(string) != "json" {
							//			return true
							//		}
							//	}
							//	return false
							//},
							Description: "The JSON subfield key value index.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The key of the subfield key value index.",
									},
									"value_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The type of value. Valid values: `long`, `double`, `text`.",
									},
								},
							},
						},
					},
				},
			},
			"user_inner_key_value": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The reserved field index configuration of the tls index.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The key of the KeyValueInfo.",
						},
						"value_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of value. Valid values: `long`, `double`, `text`, `json`.",
						},
						"case_sensitive": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether the value is case sensitive.",
						},
						"delimiter": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The delimiter of the value.",
						},
						"include_chinese": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether the value include chinese.",
						},
						"sql_flag": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether the filed is enabled for analysis.",
						},
						"json_keys": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The JSON subfield key value index.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The key of the subfield key value index.",
									},
									"value_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The type of value. Valid values: `long`, `double`, `text`.",
									},
								},
							},
						},
					},
				},
			},
			"max_text_len": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The max text length of the tls index.",
			},
			"enable_auto_index": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to enable auto index.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the tls index.",
			},
			"modify_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The modify time of the tls index.",
			},
		},
	}
	return resource
}

func resourceVolcengineTlsIndexCreate(d *schema.ResourceData, meta interface{}) (err error) {
	tlsIndexService := NewTlsIndexService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(tlsIndexService, d, ResourceVolcengineTlsIndex())
	if err != nil {
		return fmt.Errorf("error on creating TlsIndex %q, %w", d.Id(), err)
	}
	return resourceVolcengineTlsIndexRead(d, meta)
}

func resourceVolcengineTlsIndexRead(d *schema.ResourceData, meta interface{}) (err error) {
	tlsIndexService := NewTlsIndexService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(tlsIndexService, d, ResourceVolcengineTlsIndex())
	if err != nil {
		return fmt.Errorf("error on reading TlsIndex %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineTlsIndexUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	tlsIndexService := NewTlsIndexService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(tlsIndexService, d, ResourceVolcengineTlsIndex())
	if err != nil {
		return fmt.Errorf("error on updating TlsIndex %q, %w", d.Id(), err)
	}
	return resourceVolcengineTlsIndexRead(d, meta)
}

func resourceVolcengineTlsIndexDelete(d *schema.ResourceData, meta interface{}) (err error) {
	tlsIndexService := NewTlsIndexService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(tlsIndexService, d, ResourceVolcengineTlsIndex())
	if err != nil {
		return fmt.Errorf("error on deleting TlsIndex %q, %w", d.Id(), err)
	}
	return err
}
