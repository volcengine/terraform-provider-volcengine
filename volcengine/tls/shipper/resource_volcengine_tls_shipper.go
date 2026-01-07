package shipper

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Shipper can be imported using the id, e.g.
```
$ terraform import volcengine_shipper.default resource_id
```

*/

func ResourceVolcengineShipper() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineShipperCreate,
		Read:   resourceVolcengineShipperRead,
		Update: resourceVolcengineShipperUpdate,
		Delete: resourceVolcengineShipperDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"content_info": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Configuration of the delivery format for log content.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"format": {
							Optional:    true,
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Log content parsing format.",
						},
						"csv_info": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "CSV format log content configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"keys": {
										Type:        schema.TypeSet,
										Required:    true,
										Description: "Configure the fields that need to be delivered.",
										Set:         schema.HashString,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"delimiter": {
										Required: true,
										Type:     schema.TypeString,
										Description: "Delimiters are supported, including commas, " +
											"tabs, vertical bars, semicolons, and Spaces.",
									},
									"escape_char": {
										Required: true,
										Type:     schema.TypeString,
										Description: "When the field content contains a delimiter, " +
											"use an escape character to wrap the field. Currently, only single quotes, " +
											"double quotes, and null characters are supported.",
									},
									"print_header": {
										Required:    true,
										Type:        schema.TypeBool,
										Description: "Whether to print the Key on the first line.",
									},
									"non_field_content": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "Invalid field filling content, with a length ranging from 0 to 128.",
									},
								},
							},
						},
						"json_info": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "JSON format log content configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"keys": {
										Type:     schema.TypeSet,
										Optional: true,
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
										Required:    true,
										Type:        schema.TypeBool,
										Description: "Enable the flag.",
									},
									"escape": {
										Optional:    true,
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
			"kafka_shipper_info": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "JSON format log content configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_time": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "Delivery end time, millisecond timestamp. If not configured, it will keep delivering.",
						},
						"compress": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "Compression formats currently supported include snappy, gzip, lz4, and none.",
						},
						"instance": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "Kafka instance.",
						},
						"start_time": {
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Type:        schema.TypeInt,
							Description: "Delivery start time, millisecond timestamp. If not configured, the default is the current time.",
						},
						"kafka_topic": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "The name of the Kafka Topic.",
						},
					},
				},
			},
			"shipper_end_time": {
				Optional: true,
				Type:     schema.TypeInt,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() != ""
				},
				Description: "Delivery end time, millisecond timestamp. If not configured, it will keep delivering. " +
					"If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"shipper_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Delivery configuration name.",
			},
			"shipper_start_time": {
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() != ""
				},
				Type: schema.TypeInt,
				Description: "Delivery start time, millisecond timestamp. If not configured, it defaults to the current time. " +
					"If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"shipper_type": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "The type of delivery.",
			},
			"role_trn": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The role trn.",
			},
			"topic_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The log topic ID where the log to be delivered is located.",
			},
			"status": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeBool,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() != ""
				},
				Description: "Whether to enable the delivery configuration. The default value is true.",
			},
			"tos_shipper_info": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "Deliver the relevant configuration to the object storage (TOS).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bucket": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "When choosing a TOS bucket, it must be located in the same region as the source log topic.",
						},
						"prefix": {
							Optional: true,
							Computed: true,
							Type:     schema.TypeString,
							Description: "The top-level directory name of the storage bucket." +
								" All log data delivered through this delivery configuration will be delivered to this directory.",
						},
						"max_size": {
							Optional: true,
							Computed: true,
							Type:     schema.TypeInt,
							Description: "The maximum size of the original file that can be delivered to each partition (Shard), " +
								"that is, the size of the uncompressed log file. The unit is MiB, and the value range is 5 to 256.",
						},
						"compress": {
							Optional:    true,
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Compression formats currently supported include snappy, gzip, lz4, and none.",
						},
						"interval": {
							Optional:    true,
							Computed:    true,
							Type:        schema.TypeInt,
							Description: "The delivery time interval, measured in seconds, ranges from 300 to 900.",
						},
						"partition_format": {
							Optional:    true,
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Partition rules for delivering logs.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineShipperCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewShipperService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineShipper())
	if err != nil {
		return fmt.Errorf("error on creating shipper %q, %s", d.Id(), err)
	}
	return resourceVolcengineShipperRead(d, meta)
}

func resourceVolcengineShipperRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewShipperService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineShipper())
	if err != nil {
		return fmt.Errorf("error on reading shipper %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineShipperUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewShipperService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineShipper())
	if err != nil {
		return fmt.Errorf("error on updating shipper %q, %s", d.Id(), err)
	}
	return resourceVolcengineShipperRead(d, meta)
}

func resourceVolcengineShipperDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewShipperService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineShipper())
	if err != nil {
		return fmt.Errorf("error on deleting shipper %q, %s", d.Id(), err)
	}
	return err
}
