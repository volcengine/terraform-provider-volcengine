resource "volcengine_organization" "foo" {

}

data "volcengine_organization_units" "foo" {
    depends_on = [volcengine_organization.foo]
}

resource "volcengine_organization_unit" "foo" {
    name = "tf-test-unit"
    parent_id = [for unit in data.volcengine_organization_units.foo.units : unit.id if unit.parent_id == "0"][0]
    description = "tf-test"
}