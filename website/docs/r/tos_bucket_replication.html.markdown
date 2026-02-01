---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_replication"
sidebar_current: "docs-volcengine-resource-tos_bucket_replication"
description: |-
  Provides a resource to manage tos bucket replication
---
# volcengine_tos_bucket_replication
Provides a resource to manage tos bucket replication
## Example Usage
```hcl
resource "volcengine_tos_bucket_replication" "foo" {
  bucket_name = "tflyb78"
  role        = "ServiceRoleforReplicationAccessTOS"

  rules {
    id     = "rule3"
    status = "Enabled"

    prefix_set = ["documents/", "images/"]

    destination {
      bucket                          = "tflyb7-replica1"
      location                        = "cn-beijing"
      storage_class                   = "STANDARD"
      storage_class_inherit_directive = "SOURCE_OBJECT"
    }
    transfer_type                 = "internal"
    historical_object_replication = "Enabled"
    access_control_translation {
      owner = "BucketOwnerEntrusted"
    }
  }

  rules {
    id     = "rule2"
    status = "Disabled"

    destination {
      bucket                          = "tflyb7-replica2"
      location                        = "cn-beijing"
      storage_class                   = "IA"
      storage_class_inherit_directive = "DESTINATION_BUCKET"
    }
    access_control_translation {
      owner = "BucketOwnerEntrusted"
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the TOS bucket.
* `role` - (Required) The IAM role for replication.
* `rules` - (Required) The replication rules of the bucket.

The `access_control_translation` object supports the following:

* `owner` - (Optional) The owner of the destination object.

The `destination` object supports the following:

* `bucket` - (Required) The destination bucket name.
* `location` - (Required) The destination bucket location.
* `storage_class_inherit_directive` - (Optional) The storage class inherit directive. Valid values: COPY, OVERRIDE.
* `storage_class` - (Optional) The storage class for the destination bucket. Valid values: STANDARD, IA, ARCHIVE, COLD_ARCHIVE.

The `rules` object supports the following:

* `access_control_translation` - (Required) The access control translation configuration of the replication rule.
* `destination` - (Required) The destination configuration of the replication rule.
* `status` - (Required) The status of the replication rule. Valid values: Enabled, Disabled.
* `historical_object_replication` - (Optional) Whether to replicate historical objects. Valid values: Enabled, Disabled.
* `id` - (Optional) The ID of the replication rule.
* `prefix_set` - (Optional) The prefix set for the replication rule.
* `transfer_type` - (Optional) Specify the data transmission link to be used for cross-regional replication. Valid values: internal, tos_acc.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TosBucketReplication can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_replication.default bucket_name
```

