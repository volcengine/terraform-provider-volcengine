---
subcategory: "CR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cr_tags"
sidebar_current: "docs-volcengine-datasource-cr_tags"
description: |-
  Use this data source to query detailed information of cr tags
---
# volcengine_cr_tags
Use this data source to query detailed information of cr tags
## Example Usage
```hcl
data "volcengine_cr_tags" "foo" {
  registry   = "enterprise-1"
  namespace  = "test"
  repository = "repo"
  types      = ["Image"]
}
```
## Argument Reference
The following arguments are supported:
* `namespace` - (Required) The CR namespace.
* `registry` - (Required) The CR instance name.
* `repository` - (Required) The repository name.
* `names` - (Optional) The list of instance names.
* `output_file` - (Optional) File name where to save data source results.
* `types` - (Optional) The list of OCI product tag type.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `tags` - The collection of repository query.
    * `chart_attribute` - The chart attribute,valid when tag type is Chart.
        * `api_version` - The Helm version.
        * `name` - The Helm Chart name.
        * `version` - The Helm Chart version.
    * `digest` - The digest of OCI product.
    * `image_attributes` - The list of image attributes,valid when tag type is Image.
        * `architecture` - The image architecture.
        * `author` - The image author.
        * `digest` - The digest of image.
        * `os` - The iamge os.
    * `name` - The name of OCI product tag.
    * `push_time` - The last push time of OCI product.
    * `size` - The size of OCI product.
    * `type` - The type of OCI product tag.
* `total_count` - The total count of tag query.


