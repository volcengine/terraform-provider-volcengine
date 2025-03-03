package address_book

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineAddressBooks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAddressBooksRead,
		Schema: map[string]*schema.Schema{
			"group_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The group type of address book. Valid values: `ip`, `port`, `domain`.",
			},
			"group_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"description", "address"},
				Description:   "The group name of address book. This field support fuzzy query.",
			},
			"description": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"group_name", "address"},
				Description:   "The group type of address book. This field support fuzzy query.",
			},
			"address": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"group_name", "description"},
				Description:   "The group type of address book. This field support fuzzy query.",
			},
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

			"address_books": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The uuid of the address book.",
						},
						"group_uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The uuid of the address book.",
						},
						"group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the address book.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the address book.",
						},
						"group_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the address book.",
						},
						"ref_cnt": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The reference count of the address book.",
						},
						"address_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The address list of the address book.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineAddressBooksRead(d *schema.ResourceData, meta interface{}) error {
	service := NewAddressBookService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineAddressBooks())
}
