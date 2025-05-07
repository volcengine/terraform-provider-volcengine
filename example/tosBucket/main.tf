// create tos bucket
resource "volcengine_tos_bucket" "foo" {
  bucket_name = "tf-acc-test-bucket"
  #  storage_class        = "IA"
  public_acl           = "private"
  az_redundancy        = "multi-az"
  enable_version       = true
  bucket_acl_delivered = true
  account_acl {
    account_id = "1"
    permission = "READ"
  }
  account_acl {
    account_id = "2001"
    permission = "WRITE_ACP"
  }
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

// create tos bucket policy
resource "volcengine_tos_bucket_policy" "foo" {
  bucket_name = volcengine_tos_bucket.foo.id
  policy = jsonencode({
    Statement = [
      {
        Sid    = "test"
        Effect = "Allow"
        Principal = [
          "AccountId/subUserName"
        ]
        Action = [
          "tos:List*"
        ]
        Resource = [
          "trn:tos:::${volcengine_tos_bucket.foo.id}"
        ]
      }
    ]
  })
}

