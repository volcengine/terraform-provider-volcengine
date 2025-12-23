resource "volcengine_tos_bucket_lifecycle" "foo" {
  bucket_name = "tflybtest5"
  rules {
    id     = "rule1"
    status = "Enabled"
    prefix = "documents/"

    expiration {
      days = 122
    }

    tags {
        key   = "example1"
        value = "example-value1"
    }

    tags {
        key   = "example2"
        value = "example-value2"
    }

    filter {
      object_size_greater_than = 1024
      object_size_less_than    = 10485760
      greater_than_include_equal = "Enabled"
      less_than_include_equal = "Disabled"
    }
    non_current_version_expiration {
      non_current_days = 90
    }
    non_current_version_transitions {
      non_current_days = 30
      storage_class    = "IA"
    }
    non_current_version_transitions {
      non_current_days = 31
      storage_class    = "ARCHIVE"
    }


    transitions {
      days          = 7
      storage_class = "IA"
    }
    
    transitions {
      days          = 30
      storage_class = "ARCHIVE"
    }
  }
  
  rules {
    id     = "rule2"
    status = "Enabled"
    prefix = "logs/"
    
    expiration {
      days = 90
    }
    
    non_current_version_expiration {
      non_current_days = 30
    }
    
    non_current_version_transitions {
      non_current_days = 7
      storage_class    = "IA"
    }
  }
  
  rules {
    id     = "rule3"
    status = "Disabled"
    prefix = "temp/"
    
    expiration {
      date = "2025-12-31T00:00:00.000Z"
    }
    
    abort_incomplete_multipart_upload {
      days_after_initiation = 1
    }
  }
}