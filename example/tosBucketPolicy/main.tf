resource "volcengine_tos_bucket_policy" "default" {
  bucket_name = "bucket-20230418"
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
          "trn:tos:::bucket-20230418"
        ]
      }
    ]
  })
}