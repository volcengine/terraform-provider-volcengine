---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_website"
sidebar_current: "docs-volcengine-resource-tos_bucket_website"
description: |-
  Provides a resource to manage tos bucket website
---
# volcengine_tos_bucket_website
Provides a resource to manage tos bucket website
## Example Usage
```hcl
# Example: TOS Bucket Website Configuration

resource "volcengine_tos_bucket_website" "example" {
  bucket_name = "tflyb7"

  index_document {
    suffix          = "index.html"
    support_sub_dir = false # ForbiddenSubDir = "false" means support_sub_dir = false
  }

  error_document {
    key = "error1.html"
  }

  routing_rules {
    condition {
      http_error_code_returned_equals = "404"
      key_prefix_equals               = "red/"
    }
    redirect {
      host_name               = "example.com"
      http_redirect_code      = "301"
      protocol                = "http"
      replace_key_prefix_with = "redirect2/"
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the TOS bucket.
* `index_document` - (Required) The index document configuration for the website.
* `error_document` - (Optional) The error document configuration for the website.
* `redirect_all_requests_to` - (Optional) The redirect configuration for all requests.
* `routing_rules` - (Optional) The routing rules for the website.

The `condition` object supports the following:

* `http_error_code_returned_equals` - (Optional) The HTTP error code that must match for the rule to apply, e.g., 404.
* `key_prefix_equals` - (Optional) The key prefix that must match for the rule to apply.

The `error_document` object supports the following:

* `key` - (Optional) The key of the error document object, e.g., error.html.

The `index_document` object supports the following:

* `suffix` - (Required) The suffix of the index document, e.g., index.html.
* `support_sub_dir` - (Optional) Whether to support subdirectory indexing. Default is false.

The `redirect_all_requests_to` object supports the following:

* `host_name` - (Optional) The target host name for redirect.
* `protocol` - (Optional) The protocol for redirect. Valid values: http, https.

The `redirect` object supports the following:

* `host_name` - (Optional) The host name to redirect to.
* `http_redirect_code` - (Optional) The HTTP redirect code to use, e.g., 301, 302.
* `protocol` - (Optional) The protocol to use for the redirect. Valid values: http, https.
* `replace_key_prefix_with` - (Optional) The key prefix to replace the original key prefix with.
* `replace_key_with` - (Optional) The key to replace the original key with.

The `routing_rules` object supports the following:

* `condition` - (Required) The condition for the routing rule.
* `redirect` - (Required) The redirect configuration for the routing rule.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TosBucketWebsite can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_website.default bucket_name
```

