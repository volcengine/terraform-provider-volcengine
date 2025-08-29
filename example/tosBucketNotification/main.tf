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
