package kms_key

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKmsKeys() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKmsKeysRead,
		Schema: map[string]*schema.Schema{
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
			"keyring_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Query the Key ring that meets the specified conditions, which is composed of key-value pairs.",
			},
			"keyring_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Query the Key ring that meets the specified conditions, which is composed of key-value pairs.",
			},
			"filters": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Query the Key ring that meets the specified conditions, which is composed of key-value pairs.",
			},
			"tags": ve.TagsSchema(),
			"keys": {
				Description: "Master key list information.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID of the key.",
						},
						"creation_date": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The date when the keyring was created.",
						},
						"update_date": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The date when the keyring was updated.",
						},
						"key_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the key.",
						},
						"key_spec": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The algorithm used in the key.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the key.",
						},
						"key_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state of the key.",
						},
						"key_usage": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The usage of the key.",
						},
						"protection_level": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The protection level of the key.",
						},
						"schedule_delete_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when the key will be deleted.",
						},
						"rotation_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The rotation configuration of the key.",
						},
						"last_rotation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last time the key was rotated.",
						},
						"schedule_rotation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The next time the key will be rotated.",
						},
						"trn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the resource.",
						},
						"origin": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The origin of the key.",
						},
						"key_material_expire_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when the key material will expire.",
						},
						"tags": ve.TagsSchemaComputed(),
						"multi_region": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether it is the master key of the Multi-region type.",
						},
						"multi_region_configuration": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "The configuration of Multi-region key.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"multi_region_key_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the multi-region key.",
									},
									"primary_key": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "Trn and region id of the primary multi-region key",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"trn": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The trn of multi-region key.",
												},
												"region": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The region id of multi-region key.",
												},
											},
										},
									},
									"replica_keys": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Trn and region id of replica multi-region keys.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"trn": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The trn of multi-region key.",
												},
												"region": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The region id of multi-region key.",
												},
											},
										},
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

func dataSourceVolcengineKmsKeysRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKmsKeyService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKmsKeys())
}
