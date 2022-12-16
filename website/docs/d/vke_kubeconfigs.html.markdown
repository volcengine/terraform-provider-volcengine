---
subcategory: "VKE"
layout: "volcengine"
page_title: "Volcengine: volcengine_vke_kubeconfigs"
sidebar_current: "docs-volcengine-datasource-vke_kubeconfigs"
description: |-
  Use this data source to query detailed information of vke kubeconfigs
---
# volcengine_vke_kubeconfigs
Use this data source to query detailed information of vke kubeconfigs
## Example Usage
```hcl
data "volcengine_vke_kubeconfigs" "default" {
  cluster_ids = ["cce7hb97qtofmj1oi4udg"]
  types       = ["Private", "Public"]
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


