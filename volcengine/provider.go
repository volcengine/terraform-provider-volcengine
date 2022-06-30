package volcengine

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/acl"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/acl_entry"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/certificate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/clb"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/listener"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/rule"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/server_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/server_group_server"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ebs/volume"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ebs/volume_attach"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_instance_state"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/image"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/zone"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/eip/eip_address"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/eip/eip_associate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nat/nat_gateway"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nat/snat_entry"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/tos/bucket"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/tos/object"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/network_interface"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/network_interface_attach"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/route_entry"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/route_table"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/route_table_associate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/security_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/security_group_rule"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/subnet"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/vpc"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLCENGINE_ACCESS_KEY", nil),
				Description: "The Access Key for Volcengine Provider",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLCENGINE_SECRET_KEY", nil),
				Description: "The Secret Key for Volcengine Provider",
			},
			"session_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLCENGINE_SESSION_TOKEN", nil),
				Description: "The Session Token for Volcengine Provider",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLCENGINE_REGION", nil),
				Description: "The Region for Volcengine Provider",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLCENGINE_ENDPOINT", nil),
				Description: "The Customer Endpoint for Volcengine Provider",
			},
			"disable_ssl": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLCENGINE_DISABLE_SSL", nil),
				Description: "Disable SSL for Volcengine Provider",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"volcengine_vpcs":               vpc.DataSourceVolcengineVpcs(),
			"volcengine_subnets":            subnet.DataSourceVolcengineSubnets(),
			"volcengine_route_tables":       route_table.DataSourceVolcengineRouteTables(),
			"volcengine_route_entries":      route_entry.DataSourceVolcengineRouteEntries(),
			"volcengine_security_groups":    security_group.DataSourceVolcengineSecurityGroups(),
			"volcengine_network_interfaces": network_interface.DataSourceVolcengineNetworkInterfaces(),

			// ================ EIP ================
			"volcengine_eip_addresses": eip_address.DataSourceVolcengineEipAddresses(),

			// ================ CLB ================
			"volcengine_acls":                 acl.DataSourceVolcengineAcls(),
			"volcengine_clbs":                 clb.DataSourceVolcengineClbs(),
			"volcengine_listeners":            listener.DataSourceVolcengineListeners(),
			"volcengine_server_groups":        server_group.DataSourceVolcengineServerGroups(),
			"volcengine_certificates":         certificate.DataSourceVolcengineCertificates(),
			"volcengine_clb_rules":            rule.DataSourceVolcengineRules(),
			"volcengine_server_group_servers": server_group_server.DataSourceVolcengineServerGroupServers(),

			// ================ EBS ================
			"volcengine_volumes": volume.DataSourceVolcengineVolumes(),

			// ================ ECS ================
			"volcengine_ecs_instances": ecs_instance.DataSourceVolcengineEcsInstances(),
			"volcengine_images":        image.DataSourceVolcengineImages(),
			"volcengine_zones":         zone.DataSourceVolcengineZones(),

			// ================ NAT ================
			"volcengine_snat_entries": snat_entry.DataSourceVolcengineSnatEntries(),
			"volcengine_nat_gateways": nat_gateway.DataSourceVolcengineNatGateways(),

			// ================ TOS ================
			"volcengine_tos_buckets": bucket.DataSourceVolcengineTosBuckets(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"volcengine_vpc":                      vpc.ResourceVolcengineVpc(),
			"volcengine_subnet":                   subnet.ResourceVolcengineSubnet(),
			"volcengine_route_table":              route_table.ResourceVolcengineRouteTable(),
			"volcengine_route_entry":              route_entry.ResourceVolcengineRouteEntry(),
			"volcengine_route_table_associate":    route_table_associate.ResourceVolcengineRouteTableAssociate(),
			"volcengine_security_group":           security_group.ResourceVolcengineSecurityGroup(),
			"volcengine_network_interface":        network_interface.ResourceVolcengineNetworkInterface(),
			"volcengine_network_interface_attach": network_interface_attach.ResourceVolcengineNetworkInterfaceAttach(),
			"volcengine_security_group_rule":      security_group_rule.ResourceVolcengineSecurityGroupRule(),

			// ================ EIP ================
			"volcengine_eip_address":   eip_address.ResourceVolcengineEipAddress(),
			"volcengine_eip_associate": eip_associate.ResourceVolcengineEipAssociate(),

			// ================ CLB ================
			"volcengine_acl":                 acl.ResourceVolcengineAcl(),
			"volcengine_clb":                 clb.ResourceVolcengineClb(),
			"volcengine_listener":            listener.ResourceVolcengineListener(),
			"volcengine_server_group":        server_group.ResourceVolcengineServerGroup(),
			"volcengine_certificate":         certificate.ResourceVolcengineCertificate(),
			"volcengine_clb_rule":            rule.ResourceVolcengineRule(),
			"volcengine_server_group_server": server_group_server.ResourceVolcengineServerGroupServer(),
			"volcengine_acl_entry":           acl_entry.ResourceVolcengineAclEntry(),

			// ================ EBS ================
			"volcengine_volume":        volume.ResourceVolcengineVolume(),
			"volcengine_volume_attach": volume_attach.ResourceVolcengineVolumeAttach(),

			// ================ ECS ================
			"volcengine_ecs_instance":       ecs_instance.ResourceVolcengineEcsInstance(),
			"volcengine_ecs_instance_state": ecs_instance_state.ResourceVolcengineEcsInstanceState(),

			// ================ NAT ================
			"volcengine_snat_entry":  snat_entry.ResourceVolcengineSnatEntry(),
			"volcengine_nat_gateway": nat_gateway.ResourceVolcengineNatGateway(),

			//================= TOS =================
			"volcengine_tos_bucket": bucket.ResourceVolcengineTosBucket(),
			"volcengine_tos_object": object.ResourceVolcengineTosObject(),
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
