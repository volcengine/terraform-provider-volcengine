---
subcategory: "ROCKETMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rocketmq_access_keys"
sidebar_current: "docs-volcengine-datasource-rocketmq_access_keys"
description: |-
  Use this data source to query detailed information of rocketmq access keys
---
# volcengine_rocketmq_access_keys
Use this data source to query detailed information of rocketmq access keys
## Example Usage
```hcl
data "volcengine_rocketmq_access_keys" "foo" {
  instance_id = "rocketmq-cnoeea6b32118fc2"
}
```
## Argument Reference
The following arguments are supported:
* `instance_id` - (Required) The id of rocketmq instance.
* `access_key` - (Optional) The access key id of the rocketmq key.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `access_keys` - The collection of query.
    * `access_key` - The access key id of the rocketmq key.
    * `acl_config_json` - The acl config of the rocketmq key.
    * `actived` - The active status of the rocketmq key.
    * `all_authority` - The default authority of the rocketmq key.
    * `create_time` - The create time of the rocketmq key.
    * `description` - The description of the rocketmq key.
    * `instance_id` - The id of rocketmq instance.
    * `secret_key` - The secret key of the rocketmq key.
    * `topic_permissions` - The custom authority of the rocketmq key.
        * `permission` - The custom authority for the topic.
        * `topic_name` - The name of the rocketmq topic.
* `total_count` - The total count of query.


