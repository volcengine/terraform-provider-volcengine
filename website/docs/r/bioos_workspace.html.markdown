---
subcategory: "BIOOS"
layout: "volcengine"
page_title: "Volcengine: volcengine_bioos_workspace"
sidebar_current: "docs-volcengine-resource-bioos_workspace"
description: |-
  Provides a resource to manage bioos workspace
---
# volcengine_bioos_workspace
Provides a resource to manage bioos workspace
## Example Usage
```hcl
resource "volcengine_bioos_workspace" "foo" {
  name        = "test-workspace2"         //必填
  description = "test-description23"      //必填
  cover_path  = "template-cover/pic5.png" //选填
}
```
## Argument Reference
The following arguments are supported:
* `description` - (Required) The description of the workspace.
* `name` - (Required) The name of the workspace.
* `cover_path` - (Optional) Cover path (relative path in tos bucket).
* `s3_bucket` - (Optional) s3 bucket address.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `updated` - Whether the update complete.
* `workspace_id` - The id of the workspace.


## Import
Workspace can be imported using the id, e.g.
```
$ terraform import volcengine_bioos_workspace.default *****
```

