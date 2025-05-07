resource "volcengine_financial_relation" "foo" {
  sub_account_id = 2100260000
  relation       = 4
  account_alias  = "acc-test-financial"
  auth_list      = [1, 2, 3]
}
