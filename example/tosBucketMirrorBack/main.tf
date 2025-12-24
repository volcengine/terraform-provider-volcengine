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
        with_key_prefix  = "addtional-key-prefix"
        with_key_suffix  = "addtional-key-suffix"
        
        replace_key_prefix {
          key_prefix   = "key-prefix"
          replace_with = "replace-with"
        }
      }
    }
  }
}