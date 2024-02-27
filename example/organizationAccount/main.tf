resource "volcengine_organization_unit" "foo" {
  name        = "acc-test-org-unit"
  parent_id   = "730671013833632****"
  description = "acc-test"
}

resource "volcengine_organization_account" "foo" {
  account_name = "acc-test-account"
  show_name    = "acc-test-account"
  description  = "acc-test"
  org_unit_id  = volcengine_organization_unit.foo.id

  tags {
    key   = "k1"
    value = "v1"
  }
}