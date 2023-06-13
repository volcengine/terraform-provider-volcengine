---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_indexes"
sidebar_current: "docs-volcengine-datasource-tls_indexes"
description: |-
  Use this data source to query detailed information of tls indexes
---
# volcengine_tls_indexes
Use this data source to query detailed information of tls indexes
## Example Usage
```hcl
data "volcengine_tls_indexes" "default" {
  ids = ["65d67d34-c5b4-4ec8-b3a9-175d3366****"]
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Required) The list of topic id of tls index.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `tls_indexes` - The collection of tls index query.
    * `create_time` - The create time of the tls index.
    * `full_text` - The FullText index of the tls topic.
        * `case_sensitive` - Whether the FullText index is case sensitive.
        * `delimiter` - The delimiter of the FullText index.
        * `include_chinese` - Whether the FullText index include chinese.
    * `id` - The topic id of the tls index.
    * `key_value` - The KeyValue index of the tls topic.
        * `case_sensitive` - Whether the value is case sensitive.
        * `delimiter` - The delimiter of the value.
        * `include_chinese` - Whether the value include chinese.
        * `json_keys` - The JSON subfield key value index.
            * `case_sensitive` - Whether the value is case sensitive.
            * `delimiter` - The delimiter of the value.
            * `include_chinese` - Whether the value include chinese.
            * `key` - The key of the subfield key value index.
            * `sql_flag` - Whether the filed is enabled for analysis.
            * `value_type` - The type of value.
        * `key` - The key of the KeyValue index.
        * `sql_flag` - Whether the filed is enabled for analysis.
        * `value_type` - The type of value.
    * `modify_time` - The modify time of the tls index.
    * `topic_id` - The topic id of the tls index.
* `total_count` - The total count of tls index query.


