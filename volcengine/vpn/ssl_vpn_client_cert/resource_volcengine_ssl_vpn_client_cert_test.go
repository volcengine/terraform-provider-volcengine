package ssl_vpn_client_cert_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpn/ssl_vpn_client_cert"
)

const testAccVolcengineSslVpnClientCertCreateConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block = "172.16.0.0/24"
  zone_id = data.volcengine_zones.foo.zones[0].id
  vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_vpn_gateway" "foo" {
  vpc_id = volcengine_vpc.foo.id
  subnet_id = volcengine_subnet.foo.id
  bandwidth = 5
  vpn_gateway_name = "acc-test1"
  description = "acc-test1"
  period = 7
  project_name = "default"
  ssl_enabled = true
  ssl_max_connections = 5
}

resource "volcengine_ssl_vpn_server" "foo" {
  vpn_gateway_id = volcengine_vpn_gateway.foo.id
  local_subnets = [volcengine_subnet.foo.cidr_block]
  client_ip_pool = "172.16.2.0/24"
  ssl_vpn_server_name = "acc-test-ssl"
  description = "acc-test"
  protocol = "UDP"
  cipher = "AES-128-CBC"
  auth = "SHA1"
  compress = true
}

resource "volcengine_ssl_vpn_client_cert" "foo" {
  ssl_vpn_server_id = volcengine_ssl_vpn_server.foo.id
  ssl_vpn_client_cert_name = "acc-test-client-cert"
  description = "acc-test"
}
`

func TestAccVolcengineSslVpnClientCertResource_Basic(t *testing.T) {
	resourceName := "volcengine_ssl_vpn_client_cert.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return ssl_vpn_client_cert.NewSslVpnClientCertService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineSslVpnClientCertCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "ssl_vpn_client_cert_name", "acc-test-client-cert"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
					resource.TestCheckResourceAttr(acc.ResourceId, "certificate_status", "Available"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "open_vpn_client_config"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "ssl_vpn_server_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "ca_certificate"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "client_certificate"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "client_key"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "creation_time"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "update_time"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "expired_time"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccVolcengineSslVpnClientCertUpdateConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block = "172.16.0.0/24"
  zone_id = data.volcengine_zones.foo.zones[0].id
  vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_vpn_gateway" "foo" {
  vpc_id = volcengine_vpc.foo.id
  subnet_id = volcengine_subnet.foo.id
  bandwidth = 5
  vpn_gateway_name = "acc-test1"
  description = "acc-test1"
  period = 7
  project_name = "default"
  ssl_enabled = true
  ssl_max_connections = 5
}

resource "volcengine_ssl_vpn_server" "foo" {
  vpn_gateway_id = volcengine_vpn_gateway.foo.id
  local_subnets = [volcengine_subnet.foo.cidr_block]
  client_ip_pool = "172.16.2.0/24"
  ssl_vpn_server_name = "acc-test-ssl"
  description = "acc-test"
  protocol = "UDP"
  cipher = "AES-128-CBC"
  auth = "SHA1"
  compress = true
}

resource "volcengine_ssl_vpn_client_cert" "foo" {
  ssl_vpn_server_id = volcengine_ssl_vpn_server.foo.id
  ssl_vpn_client_cert_name = "acc-test-client-cert-new"
  description = "acc-test-new"
}
`

func TestAccVolcengineSslVpnClientCertResource_Update(t *testing.T) {
	resourceName := "volcengine_ssl_vpn_client_cert.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return ssl_vpn_client_cert.NewSslVpnClientCertService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineSslVpnClientCertCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "ssl_vpn_client_cert_name", "acc-test-client-cert"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
					resource.TestCheckResourceAttr(acc.ResourceId, "certificate_status", "Available"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "open_vpn_client_config"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "ssl_vpn_server_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "ca_certificate"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "client_certificate"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "client_key"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "creation_time"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "update_time"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "expired_time"),
				),
			},
			{
				Config: testAccVolcengineSslVpnClientCertUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "ssl_vpn_client_cert_name", "acc-test-client-cert-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
					resource.TestCheckResourceAttr(acc.ResourceId, "certificate_status", "Available"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "open_vpn_client_config"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "ssl_vpn_server_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "ca_certificate"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "client_certificate"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "client_key"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "creation_time"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "update_time"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "expired_time"),
				),
			},
			{
				Config:             testAccVolcengineSslVpnClientCertUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
