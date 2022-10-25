---
subcategory: "CR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cr_tag"
sidebar_current: "docs-volcengine-resource-cr_tag"
description: |-
  Provides a resource to manage cr tag
---
# volcengine_cr_tag
Provides a resource to manage cr tag
## Example Usage
```hcl
# Tag cannot be created,please import by command `terraform import volcengine_cr_tag.default registry:namespace:repository:tag`
resource "volcengine_cr_tags" "default" {

}
```
## Argument Reference
The following arguments are supported:
* `name` - (Required) The name of OCI product.
* `namespace` - (Required, ForceNew) The target namespace name.
* `registry` - (Required, ForceNew) The CrRegistry name.
* `repository` - (Required, ForceNew) The name of repository.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
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
* `push_time` - The last push time of OCI product.
* `size` - The size of OCI product.
* `type` - The type of OCI product tag.


## Import
CR tags can be imported using the registry:namespace:repository:tag, e.g.
```
$ terraform import volcengine_cr_tag.default cr-basic:namespace-1:repo-1:v1
```

