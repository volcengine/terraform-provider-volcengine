data "volcengine_iam_user_group_attachments" "default" {
  user_group_name = "xRqElT"
}

provider "volcengine" {
  enable_standard_endpoint = false
}