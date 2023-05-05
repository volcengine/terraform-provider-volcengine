package topic

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsTopics() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsTopicsRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project id of tls topic.",
			},
			"topic_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of tls topic. This field supports fuzzy queries. It is not supported to specify both TopicName and TopicId at the same time.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of tls topic. This field supports fuzzy queries. It is not supported to specify both TopicName and TopicId at the same time.",
			},
			"is_full_name": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to match accurately when filtering based on TopicName.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("topic_name").(string) == ""
				},
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of tls topic.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of tls topic query.",
			},
			"tls_topics": {
				Description: "The collection of tls topic query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the tls topic.",
						},
						"topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the tls topic.",
						},
						"topic_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the tls topic.",
						},
						"project_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project id of the tls topic.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the tls topic.",
						},
						"ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The data storage time of the tls topic. Unit: Day.",
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
						"shard_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count of shards in the tls topic.",
						},
						"auto_split": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable automatic partition splitting function of the tls topic.",
						},
						"max_split_shard": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The max count of shards in the tls topic.",
						},
						"enable_tracking": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable WebTracking function of the tls topic.",
						},
						"time_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the time field.",
						},
						"time_format": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The format of the time field.",
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
						"full_text": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The FullText index of the tls topic.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"case_sensitive": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the FullText index is case sensitive.",
									},
									"delimiter": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The delimiter of the FullText index.",
									},
									"include_chinese": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the FullText index include chinese.",
									},
								},
							},
						},
						"key_value": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The KeyValue index of the tls topic.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The key of the KeyValue index.",
									},
									"value": {
										Type:        schema.TypeList,
										Computed:    true,
										MaxItems:    1,
										Description: "The value info of the KeyValue index",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The type of value. Valid values: `long`, `double`, `text`, `json`.",
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
												"json_keys": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The JSON subfield key value index.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The key of the subfield key value index.",
															},
															"value": {
																Type:        schema.TypeList,
																Computed:    true,
																MaxItems:    1,
																Description: "The value info of the subfield key value index.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"value_type": {
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The type of value.",
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
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsTopicsRead(d *schema.ResourceData, meta interface{}) error {
	tlsProjectService := NewTlsTopicService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(tlsProjectService, d, DataSourceVolcengineTlsTopics())
}
