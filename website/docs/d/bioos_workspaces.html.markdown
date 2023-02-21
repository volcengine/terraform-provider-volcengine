---
subcategory: "BIOOS"
layout: "volcengine"
page_title: "Volcengine: volcengine_bioos_workspaces"
sidebar_current: "docs-volcengine-datasource-bioos_workspaces"
description: |-
  Use this data source to query detailed information of bioos workspaces
---
# volcengine_bioos_workspaces
Use this data source to query detailed information of bioos workspaces
## Example Usage
```hcl
data "volcengine_bioos_workspaces" "default" {

}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of workspace ids.
* `keyword` - (Optional) Keyword to filter by workspace name or description.
* `output_file` - (Optional) File name where to save data source results.
* `sort_by` - (Optional) Sort Field (Name CreateTime).
* `sort_order` - (Optional) The sort order.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `items` - A list of workspaces.
    * `cover_download_url` - The URL of the cover.
    * `create_time` - The creation time of the workspace.
    * `description` - The description of the workspace.
    * `id` - The id of the workspace.
    * `name` - The name of the workspace.
    * `owner_name` - The name of the owner of the workspace.
    * `role` - The role of the user.
    * `s3_bucket` - S3 bucket address.
    * `update_time` - The update time of the workspace.
* `total_count` - The total count of Workspace query.


