# Example: TOS Bucket Website Configuration

resource "volcengine_tos_bucket_website" "example" {
  bucket_name = "tflyb7"

  index_document {
    suffix = "index.html"
    support_sub_dir = false  # ForbiddenSubDir = "false" means support_sub_dir = false
  }

  error_document {
    key = "error1.html"
  }

  routing_rules {
    condition {
      http_error_code_returned_equals = "404"
      key_prefix_equals = "red/"
    }
    redirect {
      host_name = "example.com"
      http_redirect_code = "301"
      protocol = "http"
      replace_key_prefix_with = "redirect2/"
    }
  }
}
