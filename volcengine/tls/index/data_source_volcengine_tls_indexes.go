package index

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsIndexes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsTopicsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of topic id of tls index.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of tls index query.",
			},
			"tls_indexes": {
				Description: "The collection of tls index query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The topic id of the tls index.",
						},
						"topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The topic id of the tls index.",
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
						"max_text_len": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The max text length of the tls index.",
						},
						"enable_auto_index": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable auto index.",
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
									"index_all": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to create indexes for all fields in JSON fields with text values.",
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
						"user_inner_key_value": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The reserved field index configuration of the tls topic.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The key of the KeyValue index.",
									},
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
	}
}

func dataSourceVolcengineTlsTopicsRead(d *schema.ResourceData, meta interface{}) error {
	tlsIndexService := NewTlsIndexService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(tlsIndexService, d, DataSourceVolcengineTlsIndexes())
}
