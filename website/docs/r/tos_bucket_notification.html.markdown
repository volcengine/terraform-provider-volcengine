---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_notification"
sidebar_current: "docs-volcengine-resource-tos_bucket_notification"
description: |-
  Provides a resource to manage tos bucket notification
---
# volcengine_tos_bucket_notification
Provides a resource to manage tos bucket notification
## Example Usage
```hcl
resource "volcengine_tos_bucket" "foo" {
  bucket_name   = "tf-acc-test-bucket"
  public_acl    = "private"
  az_redundancy = "multi-az"
  project_name  = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_tos_bucket_notification" "foo" {
  bucket_name = volcengine_tos_bucket.foo.id
  rules {
    rule_id = "acc-test-rule"
    events  = ["tos:ObjectCreated:Put", "tos:ObjectCreated:Post"]
    destination {
      ve_faas {
        function_id = "80w95pns"
      }
      ve_faas {
        function_id = "crnrfajj"
      }
    }
    filter {
      tos_key {
        filter_rules {
          name  = "prefix"
          value = "a"
        }
        filter_rules {
          name  = "suffix"
          value = "b"
        }
      }
    }
  }
}

resource "volcengine_tos_bucket_notification" "foo1" {
  bucket_name = volcengine_tos_bucket.foo.id
  rules {
    rule_id = "acc-test-rule-1"
    events  = ["tos:ObjectRemoved:Delete", "tos:ObjectRemoved:DeleteMarkerCreated"]
    destination {
      ve_faas {
        function_id = "80w95pns"
      }
      ve_faas {
        function_id = "crnrfajj"
      }
    }
    filter {
      tos_key {
        filter_rules {
          name  = "prefix"
          value = "aaa"
        }
        filter_rules {
          name  = "suffix"
          value = "bbb"
        }
      }
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the bucket.
* `rules` - (Required) The notification rule of the bucket.

The `destination` object supports the following:

* `ve_faas` - (Optional) The VeFaas info of the destination.

The `filter_rules` object supports the following:

* `name` - (Optional) The name of the filter rule. Valid values: `prefix`, `suffix`.
* `value` - (Optional) The value of the filter rule.

The `filter` object supports the following:

* `tos_key` - (Optional) The tos filter of the notification.

The `rules` object supports the following:

* `destination` - (Required) The destination info of the notification.
* `events` - (Required) The event type of the notification.
* `rule_id` - (Required, ForceNew) The rule name of the notification.
* `filter` - (Optional) The filter of the notification.

The `tos_key` object supports the following:

* `filter_rules` - (Optional) The filter rules of the notification.

The `ve_faas` object supports the following:

* `function_id` - (Required) The function id of the destination.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `version` - The version of the notification.


## Import
TosBucketNotification can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_notification.default resource_id
```

