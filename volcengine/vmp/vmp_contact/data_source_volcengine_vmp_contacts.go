package vmp_contact

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVmpContacts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVmpContactsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of contact ids.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of contact.",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The email of contact.",
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
			"contacts": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of contact.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of contact.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of contact.",
						},
						"email": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The email of contact.",
						},
						"email_active": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the email of contact active.",
						},
						"contact_group_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "A list of contact group ids.",
						},
						"webhook": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The webhook of contact.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The address of webhook.",
									},
									"token": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The token of webhook.",
									},
								},
							},
						},
						"lark_bot_webhook": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The lark bot webhook of contact.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The address of webhook.",
									},
									"secret_key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The secret key of webhook.",
									},
								},
							},
						},
						"phone_number": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The phone number of contact.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"country_code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The country code of phone number.",
									},
									"number": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The number of phone number.",
									},
								},
							},
						},
						"phone_number_active": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether phone number is active.",
						},
						"ding_talk_bot_webhook": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The ding talk bot webhook of contact.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The address of webhook.",
									},
									"secret_key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The secret key of webhook.",
									},
									"at_mobiles": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The mobiles of user.",
									},
									"at_user_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The ids of user.",
									},
								},
							},
						},
						"we_com_bot_webhook": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The we com bot webhook of contact.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The address of webhook.",
									},
									"at_user_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The ids of user.",
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

func dataSourceVolcengineVmpContactsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineVmpContacts())
}
