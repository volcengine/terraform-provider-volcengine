package ecs_key_pair

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineEcsKeyPairs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineEcsKeyPairsRead,
		Schema: map[string]*schema.Schema{
			"finger_print": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The finger print info.",
			},
			"key_pair_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of key pair.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of ECS key pairs.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of ECS key pair query.",
			},
			"key_pair_names": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Key pair names info.",
			},
			"key_pair_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Ids of key pair.",
			},
			"key_pairs": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The target query key pairs info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"finger_print": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The finger print info.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of key pair.",
						},
						"key_pair_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of key pair.",
						},
						"key_pair_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of key pair.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of key pair.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of key pair.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of key pair.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineEcsKeyPairsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewEcsKeyPairService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineEcsKeyPairs())
}