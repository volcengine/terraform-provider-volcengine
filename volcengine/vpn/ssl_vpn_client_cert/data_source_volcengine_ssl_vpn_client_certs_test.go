package ssl_vpn_client_cert_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpn/ssl_vpn_client_cert"
)

const testAccVolcengineSslVpnClientCertsDatasourceConfig = `
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block = "172.16.0.0/24"
  zone_id = "cn-guilin-a"
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
  ssl_vpn_client_cert_name = "acc-test-client-cert-${count.index}"
  description = "acc-test"
  count = 5
}

data "volcengine_ssl_vpn_client_certs" "foo" {
  ids = volcengine_ssl_vpn_client_cert.foo[*].id
}
`

func TestAccVolcengineSslVpnClientCertsDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_ssl_vpn_client_certs.foo"

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
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineSslVpnClientCertsDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "ssl_vpn_client_certs.#", "5"),
				),
			},
		},
	})
}
