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

resource "volcengine_tos_bucket_cors" "foo" {
  bucket_name = volcengine_tos_bucket.foo.id
  cors_rules {
    allowed_origins = ["*"]
    allowed_methods = ["GET", "POST"]
    allowed_headers = ["Authorization"]
    expose_headers  = ["x-tos-request-id"]
    max_age_seconds = 1500
  }
  cors_rules {
    allowed_origins = ["*", "https://www.volcengine.com"]
    allowed_methods = ["POST", "PUT", "DELETE"]
    allowed_headers = ["Authorization"]
    expose_headers  = ["x-tos-request-id"]
    max_age_seconds = 2000
    response_vary   = true
  }
}