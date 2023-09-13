---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_index"
sidebar_current: "docs-volcengine-resource-tls_index"
description: |-
  Provides a resource to manage tls index
---
# volcengine_tls_index
Provides a resource to manage tls index
## Example Usage
```hcl
resource "volcengine_tls_index" "foo" {
  topic_id = "7ce12237-6670-44a7-9d79-2e36961586e6"

  #  full_text {
  #    case_sensitive = true
  #    delimiter = "!"
  #    include_chinese = false
  #  }

  key_value {
    key             = "k1"
    value_type      = "json"
    case_sensitive  = true
    delimiter       = "!"
    include_chinese = false
    sql_flag        = false
    json_keys {
      key        = "class"
      value_type = "text"
    }
    json_keys {
      key        = "age"
      value_type = "long"
    }
  }

  key_value {
    key             = "k5"
    value_type      = "text"
    case_sensitive  = true
    delimiter       = "!"
    include_chinese = false
    sql_flag        = false
  }

  user_inner_key_value {
    key             = "__content__"
    value_type      = "json"
    delimiter       = ",:-/ "
    case_sensitive  = false
    include_chinese = false
    sql_flag        = false
    json_keys {
      key        = "age"
      value_type = "long"
    }
    json_keys {
      key        = "name"
      value_type = "long"
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `topic_id` - (Required, ForceNew) The topic id of the tls index.
* `full_text` - (Optional) The full text info of the tls index.
* `key_value` - (Optional) The key value info of the tls index.
* `user_inner_key_value` - (Optional) The reserved field index configuration of the tls index.

The `full_text` object supports the following:

* `case_sensitive` - (Required) Whether the FullTextInfo is case sensitive.
* `delimiter` - (Optional) The delimiter of the FullTextInfo.
* `include_chinese` - (Optional) Whether the FullTextInfo include chinese.

The `json_keys` object supports the following:

* `key` - (Required) The key of the subfield key value index.
* `value_type` - (Required) The type of value. Valid values: `long`, `double`, `text`.

The `key_value` object supports the following:

* `key` - (Required) The key of the KeyValueInfo.
* `value_type` - (Required) The type of value. Valid values: `long`, `double`, `text`, `json`.
* `case_sensitive` - (Optional) Whether the value is case sensitive.
* `delimiter` - (Optional) The delimiter of the value.
* `include_chinese` - (Optional) Whether the value include chinese.
* `json_keys` - (Optional) The JSON subfield key value index.
* `sql_flag` - (Optional) Whether the filed is enabled for analysis.

The `user_inner_key_value` object supports the following:

* `key` - (Required) The key of the KeyValueInfo.
* `value_type` - (Required) The type of value. Valid values: `long`, `double`, `text`, `json`.
* `case_sensitive` - (Optional) Whether the value is case sensitive.
* `delimiter` - (Optional) The delimiter of the value.
* `include_chinese` - (Optional) Whether the value include chinese.
* `json_keys` - (Optional) The JSON subfield key value index.
* `sql_flag` - (Optional) Whether the filed is enabled for analysis.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of the tls index.
* `modify_time` - The modify time of the tls index.


## Import
Tls Index can be imported using the topic id, e.g.
```
$ terraform import volcengine_tls_index.default index:edf051ed-3c46-49ba-9339-bea628fe****
```

