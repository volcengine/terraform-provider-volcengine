---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_key_pairs"
sidebar_current: "docs-volcengine-datasource-ecs_key_pairs"
description: |-
  Use this data source to query detailed information of ecs key pairs
---
# volcengine_ecs_key_pairs
Use this data source to query detailed information of ecs key pairs
## Example Usage
```hcl
data "volcengine_ecs_key_pairs" "default" {
  key_pair_ids = ["kp-l8u16bn69r8ny0fz01oc", "kp-l8u1wc12n3a82ep2s1px"]
}
```
## Argument Reference
The following arguments are supported:
* `finger_print` - (Optional) The finger print info.
* `key_pair_ids` - (Optional) Ids of key pair.
* `key_pair_name` - (Optional) Name of key pair.
* `key_pair_names` - (Optional) Key pair names info.
* `name_regex` - (Optional) A Name Regex of ECS key pairs.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `key_pairs` - The target query key pairs info.
    * `created_at` - The creation time of key pair.
    * `description` - The description of key pair.
    * `finger_print` - The finger print info.
    * `id` - The id of key pair.
    * `key_pair_id` - The id of key pair.
    * `key_pair_name` - The name of key pair.
    * `updated_at` - The update time of key pair.
* `total_count` - The total count of ECS key pair query.


