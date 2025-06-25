package shipper

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineShippers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineShippersRead,
		Schema: map[string]*schema.Schema{
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the log item ID for querying the data delivery configuration under the specified log item.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the name of the log item for querying the data delivery configuration under the specified log item. Support fuzzy matching.",
			},
			"iam_project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the IAM project name for querying the data delivery configuration under the specified IAM project.",
			},
			"shipper_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Delivery configuration name.",
			},
			"shipper_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Delivery configuration ID.",
			},
			"topic_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Specify the name of the log topic for querying the data delivery configuration related to this log topic. Support fuzzy matching.",
			},
			"topic_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Specify the log topic ID for querying the data delivery configuration related to this log topic.",
			},
			"shipper_type": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Specify the delivery type for querying the delivery configuration related to that delivery type.",
			},
			"shippers": {
				Description: "Submit the relevant information of the configuration.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Computed:    true,
							Type:        schema.TypeBool,
							Description: "Whether to enable the delivery configuration.",
						},
						"topic_id": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The log topic ID where the log to be delivered is located.",
						},
						"project_id": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The log project ID where the log to be delivered is located.",
						},
						"shipper_id": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Deliver configuration ID.",
						},
						"topic_name": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The name of the log topic where the log to be delivered is located.",
						},
						"create_time": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Processing task creation time.",
						},
						"modify_time": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The most recent modification time of the processing task.",
						},
						"content_info": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The content format configuration of the delivery log.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"format": {
										Computed:    true,
										Type:        schema.TypeString,
										Description: "Log content parsing format.",
									},
									"csv_info": {
										Type:        schema.TypeList,
										Computed:    true,
										MaxItems:    1,
										Description: "CSV format log content configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"keys": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Configure the fields that need to be delivered.",
													Set:         schema.HashString,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"delimiter": {
													Computed: true,
													Type:     schema.TypeString,
													Description: "Delimiters are supported, including commas, " +
														"tabs, vertical bars, semicolons, and Spaces.",
												},
												"escape_char": {
													Computed: true,
													Type:     schema.TypeString,
													Description: "When the field content contains a delimiter, " +
														"use an escape character to wrap the field. Currently, only single quotes, " +
														"double quotes, and null characters are supported.",
												},
												"print_header": {
													Computed:    true,
													Type:        schema.TypeBool,
													Description: "Whether to print the Key on the first line.",
												},
												"non_field_content": {
													Computed:    true,
													Type:        schema.TypeString,
													Description: "Invalid field filling content, with a length ranging from 0 to 128.",
												},
											},
										},
									},
									"json_info": {
										Type:        schema.TypeList,
										Computed:    true,
										MaxItems:    1,
										Description: "JSON format log content configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"keys": {
													Type:     schema.TypeSet,
													Computed: true,
													Description: "When delivering in JSON format, if this parameter is not configured, " +
														"it indicates that all fields have been delivered." +
														" Including __content__ (choice), __source__, __path__, __time__, __image_name__," +
														" __container_name__, __pod_name__, __pod_uid__, namespace, __tag____client_ip__, __tag____receive_time__.",
													Set: schema.HashString,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"enable": {
													Computed:    true,
													Type:        schema.TypeBool,
													Description: "Enable the flag.",
												},
												"escape": {
													Computed:    true,
													Type:        schema.TypeBool,
													Description: "Whether to escape or not. It must be configured as true.",
												},
											},
										},
									},
								},
							},
						},
						"dashboard_id": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The default built-in dashboard ID for delivery.",
						},
						"project_name": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The name of the log item where the log to be delivered is located.",
						},
						"shipper_name": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Delivery configuration name.",
						},
						"shipper_start_time": {
							Computed:    true,
							Type:        schema.TypeInt,
							Description: "Delivery start time, millisecond timestamp. If not configured, it defaults to the current time.",
						},
						"shipper_type": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The type of delivery.",
						},
						"shipper_end_time": {
							Computed:    true,
							Type:        schema.TypeInt,
							Description: "Delivery end time, millisecond timestamp. If not configured, it will keep delivering.",
						},
						"tos_shipper_info": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "Deliver the relevant configuration to the object storage (TOS).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bucket": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "When choosing a TOS bucket, it must be located in the same region as the source log topic.",
									},
									"prefix": {
										Computed: true,
										Type:     schema.TypeString,
										Description: "The top-level directory name of the storage bucket." +
											" All log data delivered through this delivery configuration will be delivered to this directory.",
									},
									"max_size": {
										Computed: true,
										Type:     schema.TypeInt,
										Description: "The maximum size of the original file that can be delivered to each partition (Shard), " +
											"that is, the size of the uncompressed log file. The unit is MiB, and the value range is 5 to 256.",
									},
									"compress": {
										Computed:    true,
										Type:        schema.TypeString,
										Description: "Compression formats currently supported include snappy, gzip, lz4, and none.",
									},
									"interval": {
										Computed:    true,
										Type:        schema.TypeInt,
										Description: "The delivery time interval, measured in seconds, ranges from 300 to 900.",
									},
									"partition_format": {
										Computed:    true,
										Type:        schema.TypeString,
										Description: "Partition rules for delivering logs.",
									},
								},
							},
						},
						"kafka_shipper_info": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "JSON format log content configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_time": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Delivery end time, millisecond timestamp. If not configured, it will keep delivering.",
									},
									"compress": {
										Computed:    true,
										Type:        schema.TypeString,
										Description: "Compression formats currently supported include snappy, gzip, lz4, and none.",
									},
									"instance": {
										Computed:    true,
										Type:        schema.TypeString,
										Description: "Kafka instance.",
									},
									"start_time": {
										Computed:    true,
										Type:        schema.TypeInt,
										Description: "Delivery start time, millisecond timestamp. If not configured, the default is the current time.",
									},
									"kafka_topic": {
										Computed:    true,
										Type:        schema.TypeString,
										Description: "The name of the Kafka Topic.",
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

func dataSourceVolcengineShippersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewShipperService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineShippers())
}
