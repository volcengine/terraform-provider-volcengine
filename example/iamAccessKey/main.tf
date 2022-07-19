resource "volcengine_iam_access_key" "foo" {
  user_name = ""
  status = "active"
  secret_file = "./sk"
#  pgp_key = "keybase:some_person_that_exists"
}