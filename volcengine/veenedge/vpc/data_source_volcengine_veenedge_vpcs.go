package vpc

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVpcs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVpcsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of vpc IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Vpc.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Vpc query.",
			},
			"vpc_instances": {
				Description: "The collection of Vpc query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_identity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of VPC.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of VPC.",
						},
						"account_identity": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The id of account.",
						},
						"user_identity": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The id of user.",
						},
						"vpc_ns": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The namespace of VPC.",
						},
						"cluster_vpc_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The cluster vpc id.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of VPC.",
						},
						"is_default": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is default vpc.",
						},
						"desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of VPC.",
						},
						"vpc_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of VPC.",
						},
						"create_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The create time of VPC.",
						},
						"update_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The update time of VPC.",
						},
						"cluster": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The cluster info.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of cluster.",
									},
									"country": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The country of cluster.",
									},
									"region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region of cluster.",
									},
									"province": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The province of cluster.",
									},
									"city": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The city of cluster.",
									},
									"isp": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The isp of cluster.",
									},
									"level": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The level of cluster.",
									},
									"alias": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The alias of cluster.",
									},
								},
							},
						},
						"sub_nets": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The subnets info.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_identity": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The account id.",
									},
									"cidr_mask": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The mask of cidr.",
									},
									"create_time": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The creation time.",
									},
									"update_time": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The update time.",
									},
									"user_identity": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The id of user.",
									},
									"cidr_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ip of cidr.",
									},
									"subnet_identity": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of subnet.",
									},
								},
							},
						},
						"resource_statistic": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The resource statistic info.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"veen_instance_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The count of instance.",
									},
									"veew_lb_instance_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The count of load balancers.",
									},
									"veew_sg_instance_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The count of security groups.",
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

func dataSourceVolcengineVpcsRead(d *schema.ResourceData, meta interface{}) error {
	vpcService := NewVpcService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(vpcService, d, DataSourceVolcengineVpcs())
}
