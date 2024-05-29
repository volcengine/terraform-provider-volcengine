resource "volcengine_tos_bucket_policy" "default" {
  bucket_name = "tf-acc-test-bucket"
  policy = jsonencode({
    Statement = [
      {
        Sid = "test"
        Effect = "Allow"
        Principal = [
          "AccountId/subUserName"
        ]
        Action = [
          "tos:List*"
        ]
        Resource = [
          "trn:tos:::tf-acc-test-bucket"
        ]
      }
    ]
  })
}