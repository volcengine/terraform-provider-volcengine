---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_project"
sidebar_current: "docs-volcengine-resource-tls_project"
description: |-
  Provides a resource to manage tls project
---
# volcengine_tls_project
Provides a resource to manage tls project
## Example Usage
```hcl
resource "volcengine_tls_project" "foo" {
  project_name     = "tf-test"
  description      = "tf-desc"
  iam_project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `project_name` - (Required) The name of the tls project.
* `description` - (Optional) The description of the tls project.
* `iam_project_name` - (Optional) The IAM project name of the tls project.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The create time of the tls project.
* `inner_net_domain` - The inner net domain of the tls project.
* `topic_count` - The count of topics in the tls project.


## Import
Tls Project can be imported using the id, e.g.
```
$ terraform import volcengine_tls_project.default e020c978-4f05-40e1-9167-0113d3ef****
```

