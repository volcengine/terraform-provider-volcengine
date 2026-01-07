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
				Required:    true,
				Description: "The project id of tls topic.",
			},
			"topic_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "The id of tls topic. This field supports fuzzy queries. It is not supported to specify both TopicName and TopicId at the same time.",
				ConflictsWith: []string{"topic_name"},
			},
			"topic_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "The name of tls topic. This field supports fuzzy queries. It is not supported to specify both TopicName and TopicId at the same time.",
				ConflictsWith: []string{"topic_id"},
			},
			"tags": ve.TagsSchema(),
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
						"log_public_ip": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable the function of recording public IP.",
						},
						"enable_hot_ttl": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable tiered storage.",
						},
						"hot_ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Standard storage duration, valid when enable_hot_ttl is true.",
						},
						"cold_ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Infrequent storage duration, valid when enable_hot_ttl is true.",
						},
						"archive_ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Archive storage duration, valid when enable_hot_ttl is true.",
						},
						"encrypt_conf": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "Data encryption configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable data encryption.",
									},
									"encrypt_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The encryption type.",
									},
									"user_cmk_info": {
										Type:        schema.TypeList,
										Computed:    true,
										MaxItems:    1,
										Description: "The user custom key.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user_cmk_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The key id.",
												},
												"region_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The key region.",
												},
												"trn": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The key trn.",
												},
											},
										},
									},
								},
							},
						},
						"tags": ve.TagsSchemaComputed(),
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
