package tos_bucket_inventory

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTosBucketInventories() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTosBucketInventoriesRead,
		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name the TOS bucket.",
			},
			"inventory_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id the TOS bucket inventory.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of TOS bucket inventory.",
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
			"inventory_configurations": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bucket_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the bucket.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the bucket inventory.",
						},
						"is_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable the bucket inventory.",
						},
						"included_object_versions": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The export version of object. Valid values: `All`, `Current`.",
						},
						"schedule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The export schedule of the bucket inventory.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"frequency": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The export schedule of the bucket inventory. Valid values: `Daily`, `Weekly`.",
									},
								},
							},
						},
						"filter": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The filter of the bucket inventory.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The prefix matching information of the exported object. If not set, a list of all objects in the bucket will be generated by default.",
									},
								},
							},
						},
						"optional_fields": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The information exported from the bucket inventory.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"field": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The information exported from the bucket inventory. Valid values: `Size`, `LastModifiedDate`, `ETag`, `StorageClass`, `IsMultipartUploaded`, `EncryptionStatus`, `CRC64`, `ReplicationStatus`.",
									},
								},
							},
						},
						"destination": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The destination information of the bucket inventory.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tos_bucket_destination": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The destination tos bucket information of the bucket inventory.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bucket": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of the destination tos bucket.",
												},
												"account_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The account id of the destination tos bucket.",
												},
												"role": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The role name used to grant object storage access to read all files from the source bucket and write files to the destination bucket.",
												},
												"format": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The format of the bucket inventory. Valid values: `CSV`.",
												},
												"prefix": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The storage path prefix of the bucket inventory in destination tos bucket.",
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

func dataSourceVolcengineTosBucketInventoriesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTosBucketInventoryService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineTosBucketInventories())
}
