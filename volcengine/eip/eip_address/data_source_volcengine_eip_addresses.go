package eip_address

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineEipAddresses() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineEipAddressesRead,
		Schema: map[string]*schema.Schema{
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "A status of EIP, the value can be `Attaching` or `Detaching` or `Attached` or `Available`.",
				ValidateFunc: validation.StringInSlice([]string{"Attaching", "Detaching", "Attached", "Available"}, false),
			},
			"eip_addresses": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of EIP ip address that you want to query.",
			},
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of EIP allocation ids.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A name of EIP.",
			},
			"isp": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "An ISP of EIP Address, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom`.",
				ValidateFunc: validation.StringInSlice([]string{"BGP", "ChinaMobile", "ChinaUnicom", "ChinaTelecom"}, false),
			},
			"associated_instance_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "A type of associated instance, the value can be `Nat`, `NetworkInterface`, `ClbInstance` or `EcsInstance`.",
				ValidateFunc: validation.StringInSlice([]string{"Nat", "NetworkInterface", "ClbInstance", "EcsInstance"}, false),
			},
			"associated_instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An id of associated instance.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of EIP addresses query.",
			},
			"addresses": {
				Description: "The collection of EIP addresses.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the EIP address.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last update time of the EIP.",
						},
						"allocation_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the EIP address.",
						},
						"allocation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The allocation time of the EIP.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The peek bandwidth of the EIP.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the EIP.",
						},
						"isp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ISP of EIP Address.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance id which be associated to the EIP.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the associated instance.",
						},
						"eip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The EIP ip address of the EIP.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the EIP.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the EIP.",
						},
						"business_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The business status of the EIP.",
						},
						"lock_reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The lock reason of the EIP.",
						},
						"billing_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The billing type of the EIP.",
						},
						"overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The overdue time of the EIP.",
						},
						"deleted_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The deleted time of the EIP.",
						},
						"expired_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expired time of the EIP.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineEipAddressesRead(d *schema.ResourceData, meta interface{}) error {
	eipAddressService := NewEipAddressService(meta.(*ve.SdkClient))
	return eipAddressService.Dispatcher.Data(eipAddressService, d, DataSourceVolcengineEipAddresses())
}
