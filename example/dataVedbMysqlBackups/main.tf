data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
    vpc_name   = "acc-test-vpc"
    cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
    subnet_name = "acc-test-subnet"
    cidr_block = "172.16.0.0/24"
    zone_id = data.volcengine_zones.foo.zones[2].id
    vpc_id = volcengine_vpc.foo.id
}


resource "volcengine_vedb_mysql_instance" "foo" {
    charge_type = "PostPaid"
    storage_charge_type = "PostPaid"
    db_engine_version = "MySQL_8_0"
    db_minor_version = "3.0"
    node_number = 2
    node_spec = "vedb.mysql.x4.large"
    subnet_id = volcengine_subnet.foo.id
    instance_name = "tf-test"
    project_name = "testA"
    tags {
        key = "tftest"
        value = "tftest"
    }
    tags {
        key = "tftest2"
        value = "tftest2"
    }
}

resource "volcengine_vedb_mysql_backup" "foo" {
    instance_id = volcengine_vedb_mysql_instance.foo.id
    backup_policy {
        backup_time = "18:00Z-20:00Z"
        full_backup_period = "Monday,Tuesday,Wednesday"
        backup_retention_period = 8
    }
}

data "volcengine_vedb_mysql_backups" "foo"{
    instance_id = volcengine_vedb_mysql_instance.foo.id
}