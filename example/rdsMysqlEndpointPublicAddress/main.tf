resource "volcengine_rds_mysql_endpoint_public_address" "foo" {
    eip_id = "eip-rrq618fo9c00v0x58s4r6ky"
    endpoint_id = "mysql-38c3d4f05f6e-custom-01b0"
    instance_id = "mysql-38c3d4f05f6e"
    domain = "mysql-38c3d4f05f6e-test-01b0-public.rds.volces.com"
}
