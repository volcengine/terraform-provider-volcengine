---
subcategory: "TOS(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_tos_bucket_mirror_back"
sidebar_current: "docs-volcengine-resource-tos_bucket_mirror_back"
description: |-
  Provides a resource to manage tos bucket mirror back
---
# volcengine_tos_bucket_mirror_back
Provides a resource to manage tos bucket mirror back
## Example Usage
```hcl
resource "volcengine_tos_bucket_mirror_back" "foo" {
  bucket_name = "tflyb7"

  rules {
    id = "1"

    condition {
      http_code   = 404
      key_prefix  = "object-key-prefix"
      key_suffix  = "object-key-suffix"
      allow_host  = ["example1.volcengine.com"]
      http_method = ["GET", "HEAD"]
    }

    redirect {
      redirect_type            = "Mirror"
      fetch_source_on_redirect = false
      pass_query               = true
      follow_redirect          = true

      mirror_header {
        pass_all = true
        pass     = ["aaa", "bbb"]
        remove   = ["xxx", "yyy"]
      }

      public_source {
        source_endpoint {
          primary  = ["http://abc.123/"]
          follower = ["http://abc.456/"]
        }
      }

      transform {
        with_key_prefix = "addtional-key-prefix"
        with_key_suffix = "addtional-key-suffix"

        replace_key_prefix {
          key_prefix   = "key-prefix"
          replace_with = "replace-with"
        }
      }
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `bucket_name` - (Required, ForceNew) The name of the TOS bucket.
* `rules` - (Required) The mirror_back rules of the bucket.

The `condition` object supports the following:

* `http_code` - (Required) Error code for triggering the source re-fetch function.
* `allow_host` - (Optional) Only when a specific domain name is supported will the origin retrieval be triggered.
* `http_method` - (Optional) The type of request that triggers the re-sourcing process.
* `key_prefix` - (Optional) The prefix of the object name that matches the source object.
* `key_suffix` - (Optional) The suffix of the object name that matches the source object.

The `credential_provider` object supports the following:

* `role` - (Optional) The role.

The `fetch_header_to_meta_data_rules` object supports the following:

* `meta_data_suffix` - (Required) The metadata suffix.
* `source_header` - (Required) The source header.

The `follower` object supports the following:

* `bucket_name` - (Optional) The bucket name.
* `credential_provider` - (Optional) The credential provider.
* `endpoint` - (Optional) The endpoint.

The `mirror_header` object supports the following:

* `pass_all` - (Optional) Whether to pass all headers.
* `pass` - (Optional) The headers to pass.
* `remove` - (Optional) The headers to remove.
* `set` - (Optional) The mirror header configuration.

The `primary` object supports the following:

* `bucket_name` - (Optional) The bucket name.
* `credential_provider` - (Optional) The credential provider.
* `endpoint` - (Optional) The endpoint.

The `private_source` object supports the following:

* `source_endpoint` - (Optional) The source endpoint.

The `public_source` object supports the following:

* `fixed_endpoint` - (Optional) Whether the endpoint is fixed.
* `source_endpoint` - (Optional) The source endpoint.

The `redirect` object supports the following:

* `fetch_header_to_meta_data_rules` - (Optional) The fetch header to metadata rules.
* `fetch_source_on_redirect_with_query` - (Optional) Whether to fetch source on redirect with query.
* `fetch_source_on_redirect` - (Optional) Whether to fetch source on redirect.
* `follow_redirect` - (Optional) Whether to follow redirects.
* `mirror_header` - (Optional) The mirror header configuration.
* `pass_query` - (Optional) Whether to pass query parameters.
* `private_source` - (Optional) The private source configuration.
* `public_source` - (Optional) The public source configuration.
* `redirect_type` - (Optional) The type of redirect.
* `transform` - (Optional) The transform configuration.

The `replace_key_prefix` object supports the following:

* `key_prefix` - (Optional) The key prefix to replace.
* `replace_with` - (Optional) The value to replace with.

The `rules` object supports the following:

* `condition` - (Optional) The condition of the mirror_back rule.
* `id` - (Optional) The ID of the mirror_back rule.
* `redirect` - (Optional) The redirect configuration of the mirror_back rule.

The `set` object supports the following:

* `key` - (Optional) The key of the header.
* `value` - (Optional) The value of the header.

The `source_endpoint` object supports the following:

* `follower` - (Optional) The follower endpoints.
* `primary` - (Optional) The primary endpoints.

The `transform` object supports the following:

* `replace_key_prefix` - (Optional) The replace key prefix configuration.
* `with_key_prefix` - (Optional) The key prefix to add.
* `with_key_suffix` - (Optional) The key suffix to add.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
TosBucketMirrorBack can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_mirror_back.default bucket_name
```

