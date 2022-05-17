package vestack

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/vestack/clb/acl"
	"github.com/volcengine/terraform-provider-vestack/vestack/clb/acl_entry"
	"github.com/volcengine/terraform-provider-vestack/vestack/clb/certificate"
	"github.com/volcengine/terraform-provider-vestack/vestack/clb/clb"
	"github.com/volcengine/terraform-provider-vestack/vestack/clb/listener"
	"github.com/volcengine/terraform-provider-vestack/vestack/clb/rule"
	"github.com/volcengine/terraform-provider-vestack/vestack/clb/server_group"
	"github.com/volcengine/terraform-provider-vestack/vestack/clb/server_group_server"
	"github.com/volcengine/terraform-provider-vestack/vestack/ebs/volume"
	"github.com/volcengine/terraform-provider-vestack/vestack/ebs/volume_attach"
	"github.com/volcengine/terraform-provider-vestack/vestack/ecs/ecs_instance"
	"github.com/volcengine/terraform-provider-vestack/vestack/ecs/ecs_instance_state"
	"github.com/volcengine/terraform-provider-vestack/vestack/ecs/image"
	"github.com/volcengine/terraform-provider-vestack/vestack/ecs/zone"
	"github.com/volcengine/terraform-provider-vestack/vestack/eip/eip_address"
	"github.com/volcengine/terraform-provider-vestack/vestack/eip/eip_associate"
	"github.com/volcengine/terraform-provider-vestack/vestack/nat/nat_gateway"
	"github.com/volcengine/terraform-provider-vestack/vestack/nat/snat_entry"
	"github.com/volcengine/terraform-provider-vestack/vestack/vpc/network_interface"
	"github.com/volcengine/terraform-provider-vestack/vestack/vpc/network_interface_attach"
	"github.com/volcengine/terraform-provider-vestack/vestack/vpc/route_entry"
	"github.com/volcengine/terraform-provider-vestack/vestack/vpc/route_table"
	"github.com/volcengine/terraform-provider-vestack/vestack/vpc/route_table_associate"
	"github.com/volcengine/terraform-provider-vestack/vestack/vpc/security_group"
	"github.com/volcengine/terraform-provider-vestack/vestack/vpc/security_group_rule"
	"github.com/volcengine/terraform-provider-vestack/vestack/vpc/subnet"
	"github.com/volcengine/terraform-provider-vestack/vestack/vpc/vpc"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VESTACK_ACCESS_KEY", nil),
				Description: "The Access Key for Vestack Provider",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VESTACK_SECRET_KEY", nil),
				Description: "The Secret Key for Vestack Provider",
			},
			"session_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VESTACK_SESSION_TOKEN", nil),
				Description: "The Session Token for Vestack Provider",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VESTACK_REGION", nil),
				Description: "The Region for Vestack Provider",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VESTACK_ENDPOINT", nil),
				Description: "The Customer Endpoint for Vestack Provider",
			},
			"disable_ssl": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("vestack_DISABLE_SSL", nil),
				Description: "Disable SSL for Vestack Provider",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"vestack_vpcs":               vpc.DataSourceVestackVpcs(),
			"vestack_subnets":            subnet.DataSourceVestackSubnets(),
			"vestack_route_tables":       route_table.DataSourceVestackRouteTables(),
			"vestack_route_entries":      route_entry.DataSourceVestackRouteEntries(),
			"vestack_security_groups":    security_group.DataSourceVestackSecurityGroups(),
			"vestack_network_interfaces": network_interface.DataSourceVestackNetworkInterfaces(),

			// ================ EIP ================
			"vestack_eip_addresses": eip_address.DataSourceVestackEipAddresses(),

			// ================ CLB ================
			"vestack_acls":                 acl.DataSourceVestackAcls(),
			"vestack_clbs":                 clb.DataSourceVestackClbs(),
			"vestack_listeners":            listener.DataSourceVestackListeners(),
			"vestack_server_groups":        server_group.DataSourceVestackServerGroups(),
			"vestack_certificates":         certificate.DataSourceVestackCertificates(),
			"vestack_clb_rules":            rule.DataSourceVestackRules(),
			"vestack_server_group_servers": server_group_server.DataSourceVestackServerGroupServers(),

			// ================ EBS ================
			"vestack_volumes": volume.DataSourceVestackVolumes(),

			// ================ ECS ================
			"vestack_ecs_instances": ecs_instance.DataSourceVestackEcsInstances(),
			"vestack_images":        image.DataSourceVestackImages(),
			"vestack_zones":         zone.DataSourceVestackZones(),

			// ================ NAT ================
			"vestack_snat_entries": snat_entry.DataSourceVestackSnatEntries(),
			"vestack_nat_gateways": nat_gateway.DataSourceVestackNatGateways(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"vestack_vpc":                      vpc.ResourceVestackVpc(),
			"vestack_subnet":                   subnet.ResourceVestackSubnet(),
			"vestack_route_table":              route_table.ResourceVestackRouteTable(),
			"vestack_route_entry":              route_entry.ResourceVestackRouteEntry(),
			"vestack_route_table_associate":    route_table_associate.ResourceVestackRouteTableAssociate(),
			"vestack_security_group":           security_group.ResourceVestackSecurityGroup(),
			"vestack_network_interface":        network_interface.ResourceVestackNetworkInterface(),
			"vestack_network_interface_attach": network_interface_attach.ResourceVestackNetworkInterfaceAttach(),
			"vestack_security_group_rule":      security_group_rule.ResourceVestackSecurityGroupRule(),

			// ================ EIP ================
			"vestack_eip_address":   eip_address.ResourceVestackEipAddress(),
			"vestack_eip_associate": eip_associate.ResourceVestackEipAssociate(),

			// ================ CLB ================
			"vestack_acl":                 acl.ResourceVestackAcl(),
			"vestack_clb":                 clb.ResourceVestackClb(),
			"vestack_listener":            listener.ResourceVestackListener(),
			"vestack_server_group":        server_group.ResourceVestackServerGroup(),
			"vestack_certificate":         certificate.ResourceVestackCertificate(),
			"vestack_clb_rule":            rule.ResourceVestackRule(),
			"vestack_server_group_server": server_group_server.ResourceVestackServerGroupServer(),
			"vestack_acl_entry":           acl_entry.ResourceVestackAclEntry(),

			// ================ EBS ================
			"vestack_volume":        volume.ResourceVestackVolume(),
			"vestack_volume_attach": volume_attach.ResourceVestackVolumeAttach(),

			// ================ ECS ================
			"vestack_ecs_instance":       ecs_instance.ResourceVestackEcsInstance(),
			"vestack_ecs_instance_state": ecs_instance_state.ResourceVestackEcsInstanceState(),

			// ================ NAT ================
			"vestack_snat_entry":  snat_entry.ResourceVestackSnatEntry(),
			"vestack_nat_gateway": nat_gateway.ResourceVestackNatGateway(),
		},
		ConfigureFunc: ProviderConfigure,
	}
}

func ProviderConfigure(d *schema.ResourceData) (interface{}, error) {
	config := ve.Config{
		AccessKey:    d.Get("access_key").(string),
		SecretKey:    d.Get("secret_key").(string),
		SessionToken: d.Get("session_token").(string),
		Region:       d.Get("region").(string),
		Endpoint:     d.Get("endpoint").(string),
		DisableSSL:   d.Get("disable_ssl").(bool),
	}
	client, err := config.Client()
	return client, err
}
