---
subcategory: "APIG"
layout: "volcengine"
page_title: "Volcengine: volcengine_apig_gateways"
sidebar_current: "docs-volcengine-datasource-apig_gateways"
description: |-
  Use this data source to query detailed information of apig gateways
---
# volcengine_apig_gateways
Use this data source to query detailed information of apig gateways
## Example Usage
```hcl
data "volcengine_apig_gateways" "foo" {
  ids = ["gd13d8c6eq1emkiunq6p0", "gd07fq7pte3scmnaj1b1g"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of api gateway IDs.
* `name_regex` - (Optional) A Name Regex of Resource.
* `name` - (Optional) The name of api gateway. This field support fuzzy query.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The project name of api gateway.
* `status` - (Optional) The status of api gateway.
* `tags` - (Optional) Tags.
* `type` - (Optional) The type of api gateway.
* `vpc_ids` - (Optional) A list of vpc IDs.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `gateways` - The collection of query.
    * `backend_spec` - The backend spec of the api gateway.
        * `is_vke_with_flannel_cni_supported` - Whether the api gateway support vke flannel cni.
        * `vke_pod_cidr` - The vke pod cidr of the api gateway.
    * `comments` - The comments of the api gateway.
    * `create_time` - The create time of the api gateway.
    * `id` - The Id of the api gateway.
    * `log_spec` - The log spec of the api gateway.
        * `enable` - Whether the api gateway enable tls log.
        * `project_id` - The project id of the tls.
        * `topic_id` - The topic id of the tls.
    * `message` - The error message of the api gateway.
    * `monitor_spec` - The monitor spec of the api gateway.
        * `enable` - Whether the api gateway enable monitor.
        * `workspace_id` - The workspace id of the monitor.
    * `name` - The name of the api gateway.
    * `network_spec` - The network spec of the api gateway.
        * `subnet_ids` - The subnet ids of the api gateway.
        * `vpc_id` - The vpc id of the api gateway.
    * `project_name` - The project name of the api gateway.
    * `region` - The region of the api gateway.
    * `resource_spec` - The resource spec of the api gateway.
        * `clb_spec_code` - The clb spec code of the resource spec.
        * `instance_spec_code` - The instance spec code of the resource spec.
        * `network_type` - The network type of the api gateway.
            * `enable_private_network` - Whether the api gateway enable private network.
            * `enable_public_network` - Whether the api gateway enable public network.
        * `public_network_bandwidth` - The public network bandwidth of the resource spec.
        * `public_network_billing_type` - The public network billing type of the resource spec.
        * `replicas` - The replicas of the resource spec.
    * `status` - The status of the api gateway.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `type` - The type of the api gateway.
    * `version` - The version of the api gateway.
* `total_count` - The total count of query.


