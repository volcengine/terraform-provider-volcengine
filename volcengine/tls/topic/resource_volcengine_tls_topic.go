package topic

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

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
				Description:  "The max count of shards in the tls topic.",
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
			"full_text": {
				Type:     schema.TypeList,
				Optional: true,
				//Computed:    true,
				MaxItems:    1,
				Description: "The FullText index of the tls topic.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"case_sensitive": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether the FullText index is case sensitive.",
						},
						"delimiter": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The delimiter of the FullText index.",
						},
						"include_chinese": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether the FullText index include chinese.",
						},
					},
				},
			},
			"key_value": {
				Type:     schema.TypeList,
				Optional: true,
				//Computed:    true,
				Description: "The KeyValue index of the tls topic.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The key of the KeyValue index.",
						},
						"value": {
							Type:        schema.TypeList,
							Required:    true,
							MaxItems:    1,
							Description: "The value info of the KeyValue index",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value_type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"long", "double", "text", "json"}, false),
										Description:  "The type of value. Valid values: `long`, `double`, `text`, `json`.",
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
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
											items := strings.Split(k, ".")
											if len(items) > 4 {
												key := strings.Join(items[:4], ".")
												if valueType := d.Get(key + ".value_type"); valueType != nil && valueType.(string) != "json" {
													return true
												}
											}
											return false
										},
										Description: "The JSON subfield key value index.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The key of the subfield key value index.",
												},
												"value": {
													Type:        schema.TypeList,
													Required:    true,
													MaxItems:    1,
													Description: "The value info of the subfield key value index.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"value_type": {
																Type:         schema.TypeString,
																Required:     true,
																ValidateFunc: validation.StringInSlice([]string{"long", "double", "text"}, false),
																Description:  "The type of value. Valid values: `long`, `double`, `text`.",
															},
															"case_sensitive": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Whether the value is case sensitive.",
															},
															"delimiter": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The delimiter of the value.",
															},
															"include_chinese": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Whether the value include chinese.",
															},
															"sql_flag": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Whether the filed is enabled for analysis.",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},

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
			"index_create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of index in the tls topic.",
			},
			"index_modify_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The modify time of index in the tls topic.",
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
