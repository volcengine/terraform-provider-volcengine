package cloud_server

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CloudServer can be imported using the id, e.g.
```
$ terraform import volcengine_veenedge_cloud_server.default cloudserver-n769ewmjjqyqh5dv
```

After the veenedge cloud server is created, a default edge instance will be created, we recommend managing this default instance as follows
```
resource "volcengine_veenedge_instance" "foo1" {
  instance_id = volcengine_veenedge_cloud_server.foo.default_instance_id
}
```
*/

func ResourceVolcengineCloudServer() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCloudServerCreate,
		Read:   resourceVolcengineCloudServerRead,
		Delete: resourceVolcengineCloudServerDelete,
		Update: resourceVolcengineCloudServerUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cloudserver_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of cloud server.",
			},
			"image_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The image id of cloud server.",
			},
			"spec_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The spec name of cloud server.",
			},
			"server_area_level": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"region", "city"}, false),
				Description:  "The server area level. The value can be `region` or `city`.",
			},
			"storage_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				ForceNew:    true,
				Description: "The config of the storage.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"system_disk": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							MaxItems:    1,
							Description: "The disk info of system.",
							Elem:        diskResSpec,
						},
						"data_disk_list": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: "The disk list info of data.",
							Elem:        diskResSpec,
						},
					},
				},
			},
			"network_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				ForceNew:    true,
				Description: "The config of the network.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bandwidth_peak": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The peak of bandwidth.",
						},
						"internal_bandwidth_peak": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "The internal peak of bandwidth.",
						},
						"enable_ipv6": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "Whether enable ipv6.",
						},
						"custom_internal_interface_name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The name of custom internal interface.",
						},
						"custom_external_interface_name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The name of custom external interface.",
						},
					},
				},
			},
			"secret_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"KeyPair", "Password"}, false),
				Description:  "The type of secret. The value can be `KeyPair` or `Password`.",
			},
			"secret_data": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The data of secret. The value can be Password or KeyPair ID.",
			},
			"default_area_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of default area.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() != ""
				},
			},
			"default_isp": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The default isp info.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() != ""
				},
			},
			"default_cluster_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of default cluster.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() != ""
				},
			},
			"schedule_strategy": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				ForceNew:    true,
				Description: "The schedule strategy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"schedule_strategy": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"dispersion", "concentration"}, false),
							Description:  "The type of schedule strategy. The value can be `dispersion` or `concentration`.",
						},
						"price_strategy": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"high_priority", "low_priority"}, false),
							Description:  "The price strategy. The value can be `high_priority` or `low_priority`.",
						},
						"network_strategy": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The network strategy.",
						},
					},
				},
			},
			"billing_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "The config of the billing.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"computing_billing_method": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"MonthlyPeak", "DailyPeak"}, false),
							Description:  "The method of computing billing. The value can be `MonthlyPeak` or `DailyPeak`.",
						},
						"bandwidth_billing_method": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"MonthlyP95", "DailyPeak"}, false),
							Description:  "The method of bandwidth billing. The value can be `MonthlyP95` or `DailyPeak`.",
						},
					},
				},
			},
			"custom_data": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The custom data.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Sensitive:   true,
							Description: "The custom data info.",
						},
					},
				},
			},
			"default_instance_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default instance id generate by cloud server.",
			},
		},
	}

	return resource
}

var diskResSpec = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"storage_type": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"CloudBlockHDD", "CloudBlockSSD"}, false),
			Description:  "The type of storage. The value can be `CloudBlockHDD` or `CloudBlockSSD`.",
		},
		"capacity": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The capacity of storage.",
		},
	},
}

func resourceVolcengineCloudServerCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudServerService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineCloudServer())
	if err != nil {
		return fmt.Errorf(" Error on creating cloud server %q,%s", d.Id(), err)
	}
	return resourceVolcengineCloudServerRead(d, meta)
}

func resourceVolcengineCloudServerDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudServerService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineCloudServer())
	if err != nil {
		return fmt.Errorf("error on deleting cloud server %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudServerRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudServerService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineCloudServer())
	if err != nil {
		return fmt.Errorf("error on reading cloud server %q,%s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudServerUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudServerService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineCloudServer())
	if err != nil {
		return fmt.Errorf("error on updating CloudServer  %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudServerRead(d, meta)
}