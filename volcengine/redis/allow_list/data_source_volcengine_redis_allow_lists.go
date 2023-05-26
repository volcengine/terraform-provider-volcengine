package allow_list

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRedisAllowLists() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRedisAllowListRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Id of instance.",
			},
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Id of region.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Allow List.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of allow list query.",
			},
			"allow_lists": {
				Description: "Information of list of allow list.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allow_list_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of allow list.",
						},
						"allow_list_ip_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The IP number of allow list.",
						},
						"allow_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "Ip list of allow list.",
						},
						"allow_list_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Id of allow list.",
						},
						"allow_list_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of allow list.",
						},
						"allow_list_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of allow list.",
						},
						"associated_instance_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of instance that associated to allow list.",
						},
						"associated_instances": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Instances associated by this allow list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Id of instance.",
									},
									"instance_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of instance.",
									},
									"vpc": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Id of virtual private cloud.",
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

func dataSourceVolcengineRedisAllowListRead(d *schema.ResourceData, meta interface{}) error {
	redisAllowListService := NewRedisAllowListService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Data(redisAllowListService, d, DataSourceVolcengineRedisAllowLists())
	return err
}
