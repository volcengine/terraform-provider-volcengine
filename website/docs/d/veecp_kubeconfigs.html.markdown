---
subcategory: "VEECP"
layout: "volcengine"
page_title: "Volcengine: volcengine_veecp_kubeconfigs"
sidebar_current: "docs-volcengine-datasource-veecp_kubeconfigs"
description: |-
  Use this data source to query detailed information of veecp kubeconfigs
---
# volcengine_veecp_kubeconfigs
Use this data source to query detailed information of veecp kubeconfigs
## Example Usage
```hcl
data "volcengine_veecp_kubeconfigs" "foo" {
  cluster_ids = []
  ids         = []
  page_number = 1
  page_size   = 1
  role_ids    = []
  types       = []
}
```
## Argument Reference
The following arguments are supported:
* `cluster_ids` - (Optional) A list of Cluster IDs.
* `ids` - (Optional) A list of Kubeconfig IDs.
* `name_regex` - (Optional) A Name Regex of Kubeconfig.
* `output_file` - (Optional) File name where to save data source results.
* `page_number` - (Optional) The page number of Kubeconfigs query.
* `page_size` - (Optional) The page size of Kubeconfigs query.
* `role_ids` - (Optional) A list of Role IDs.
* `types` - (Optional) The type of Kubeconfigs query.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `kubeconfigs` - The collection of VkeKubeconfig query.
    * `cluster_id` - The Cluster ID of the Kubeconfig.
    * `create_time` - The create time of the Kubeconfig.
    * `expire_time` - The expire time of the Kubeconfig.
    * `id` - The ID of the Kubeconfig.
    * `kubeconfig_id` - The ID of the Kubeconfig.
    * `kubeconfig` - Kubeconfig data with public/private network access, returned in BASE64 encoding.
    * `type` - The type of the Kubeconfig.
    * `user_id` - The account ID of the Kubeconfig.
* `total_count` - The total count of Kubeconfig query.


