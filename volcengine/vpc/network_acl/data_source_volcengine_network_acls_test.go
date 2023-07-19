package network_acl_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/network_acl"
)

const testAccNetworkAclConfig = `
	data "volcengine_zones" "foo"{
	}

	resource "volcengine_vpc" "foo" {
	  vpc_name   = "acc-test-vpc"
	  cidr_block = "172.16.0.0/16"
	}

	resource "volcengine_network_acl" "foo" {
	  vpc_id = "${volcengine_vpc.foo.id}"
	  network_acl_name = "acc-test-acl"
	  ingress_acl_entries {
		network_acl_entry_name = "acc-ingress1"
		policy = "accept"
		protocol = "all"
		source_cidr_ip = "192.168.0.0/24"
	  }
	  egress_acl_entries {
		network_acl_entry_name = "acc-egress2"
		policy = "accept"
		protocol = "all"
		destination_cidr_ip = "192.168.0.0/16"
	  }
	}
	data "volcengine_network_acls" "foo"{
  	network_acl_name = "acc-test-acl"
	}
`

func TestAccVolcengineNetworkAclDataSource_Basic(t *testing.T) {
	resourceName := "data.volcengine_network_acls.foo"
	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &network_acl.VolcengineNetworkAclService{},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkAclConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "network_acls.#", "1"),
				),
			},
		},
	})
}
