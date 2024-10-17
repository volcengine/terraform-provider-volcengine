package vedb_mysql_endpoint_public_address

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVedbMysqlEndpointPublicAddresss() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVedbMysqlEndpointPublicAddresssRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of IDs.",
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
			// TODO: change this field to the target datasource
			"instances": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

					},
				},
			},
		},
	}
}

func dataSourceVolcengineVedbMysqlEndpointPublicAddresssRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVedbMysqlEndpointPublicAddressService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVedbMysqlEndpointPublicAddresss())
}
