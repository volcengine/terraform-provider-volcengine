resource "vestack_access_key" "foo" {
  user_name = "tf-test1"
  status = "active"
  secret_file = "./sk"
#  pgp_key = "keybase:some_person_that_exists"
}