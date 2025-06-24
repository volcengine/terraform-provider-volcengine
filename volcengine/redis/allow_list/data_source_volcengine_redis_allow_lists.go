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
			"query_default": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Filter whether to query only the default whitelist based on the type of whitelist.",
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Filter out the whitelist that meets the conditions based on the IP address. " +
					"When using IPAddress query, it will precisely match this IP address and filter the IP address segments containing this IP address.",
			},
			"ip_segment": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Screen out the whitelist that meets the conditions based on the IP address segment. " +
					"When using IPSegment queries, the IP address segment will be precisely matched for filtering.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the project to which the white list belongs.",
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
						"allow_list_category": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the whitelist.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the project to which the white list belongs.",
						},
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
						"security_group_bind_infos": {
							Description: "The current whitelist is the list of security group information that has been associated.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bind_mode": {
										Type:     schema.TypeString,
										Computed: true,
										Description: "Security group association mode. The value range is as follows: " +
											"IngressDirectionIp: The input direction IP, which is the IP involved in the TCP protocol and ALL protocol in the source address of the secure group input direction to access the database. " +
											"If the source address is configured as a secure group, it will be ignored. " +
											"AssociateEcsIp: Associate ECS IP, which allows cloud servers within the security group to access the database." +
											" Currently, only the IP information of the main network card is supported for import.",
									},
									"security_group_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The associated security group ID.",
									},
									"security_group_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the associated security group.",
									},
									"ip_list": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The list of ips in the associated security group has been linked.",
									},
								},
							},
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
