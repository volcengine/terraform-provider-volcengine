package rds_postgresql_instance_price_difference

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlInstancePriceDifferences() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlInstancePriceDifferencesRead,
		Schema: map[string]*schema.Schema{
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
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance ID.",
			},
			"modify_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Usually",
				ValidateFunc: validation.StringInSlice([]string{"Usually", "Temporary"}, false),
				Description: "Spec change type. Usually or Temporary. Default value: Usually. " +
					"This parameter can only take the value Temporary when the billing type of the instance is a yearly/monthly subscription instance.",
			},
			"rollback_time": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Rollback time for Temporary change, UTC format yyyy-MM-ddTHH:mm:ss.sssZ. " +
					"This parameter is required when the modify_type is set to Temporary.",
			},
			"node_info": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 2,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_id": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "The id of the node." +
								"When the modify_type is set to Temporary, this parameter is required.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The AZ of the node.",
						},
						"node_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"Primary", "Secondary", "ReadOnly"}, false),
							Description:  "The type of the node. Valid values: Primary, Secondary, ReadOnly.",
						},
						"node_spec": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The specification of the node.",
						},
						"node_operate_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"Create", "Modify"}, false),
							Description:  "The operate type of the node. Valid values: Create, Modify.",
						},
					},
				},
				Description: "Instance spec nodes. Primary=1, Secondary=1, ReadOnly=0~10.",
			},
			"storage_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"LocalSSD"}, false),
				Description:  "The type of the storage. Valid values: LocalSSD.",
			},
			"storage_space": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The storage space of the instance. Value range: [20, 3000], unit: GB, step 10GB.",
			},
			"charge_info": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"charge_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"PostPaid", "PrePaid"}, false),
							Description:  "The charge type of the instance. Valid values: PostPaid, PrePaid.",
						},
						"auto_renew": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to auto renew the subscription in a pre-paid scenario.",
						},
						"period_unit": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"Month", "Year"}, false),
							Description:  "Purchase cycle in a pre-paid scenario. Valid values: Month, Year.",
						},
						"period": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Subscription duration in a pre-paid scenario.Default value:1.",
						},
						"number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     1,
							Description: "Number of purchased instances. Can be an integer between 1 and 20. Default value:1.",
						},
					},
				},
				Description: "Charge info of the instance.",
			},
			"instances_price": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"currency": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Currency unit.",
						},
						"discount_price": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Instance price after discount.",
						},
						"original_price": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Instance price before discount.",
						},
						"payable_price": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Price payable of instance.",
						},
						"charge_item_prices": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"charge_item_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Billing item name. Values:Primary, Secondary, ReadOnly, Storage.",
									},
									"charge_item_key": {
										Type:     schema.TypeString,
										Computed: true,
										Description: "If charge_item_key is Primary, Secondary, or ReadOnly, this parameter returns the instance specification, such as rds.pg.d1.1c2g. " +
											"If charge_item_key is Storage, this parameter returns the stored key, such as rds.pg.d1.localssd.",
									},
									"charge_item_value": {
										Type:     schema.TypeInt,
										Computed: true,
										Description: "If charge_item_key is Primary, Secondary, or ReadOnly, this parameter returns the number of nodes, with a value of \"1\". " +
											"If charge_item_key is Storage, his parameter returns the storage size in GB.",
									},
									"node_num_per_instance": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of nodes of each instance.",
									},
									"original_price": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Original price of each charge item.",
									},
									"discount_price": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Discount price of each charge item.",
									},
									"payable_price": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Payable price of each charge item.",
									},
								},
							},
						},
					},
				},
				Description: "The collection of query.",
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlInstancePriceDifferencesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlInstancePriceDifferenceService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlInstancePriceDifferences())
}
