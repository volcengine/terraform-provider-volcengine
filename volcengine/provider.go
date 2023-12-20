package volcengine

import (
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_acl"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_certificate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_server_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_server_group_server"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_zone"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cen/cen_service_route_entry"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloud_monitor/cloud_monitor_contact"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloud_monitor/cloud_monitor_contact_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloudfs/cloudfs_access"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloudfs/cloudfs_file_system"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloudfs/cloudfs_namespace"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloudfs/cloudfs_ns_quota"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloudfs/cloudfs_quota"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_iam_role_attachment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_user_group_policy_attachment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nas/nas_file_system"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nas/nas_region"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nas/nas_snapshot"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nas/nas_zone"
	trEntry "github.com/volcengine/terraform-provider-volcengine/volcengine/transit_router/route_entry"
	trTable "github.com/volcengine/terraform-provider-volcengine/volcengine/transit_router/route_table"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/transit_router/route_table_association"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/transit_router/route_table_propagation"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/transit_router/shared_transit_router_state"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/transit_router/transit_router"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/transit_router/transit_router_bandwidth_package"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/transit_router/transit_router_direct_connect_gateway_attachment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/transit_router/transit_router_grant_rule"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/transit_router/transit_router_peer_attachment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/transit_router/transit_router_vpc_attachment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/transit_router/transit_router_vpn_attachment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/support_resource_types"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/ha_vip"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/ha_vip_associate"
	"strings"

	regions "github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/region"

	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds/rds_parameter_template"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/redis/instance_state"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/redis/pitr_time_period"

	"github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/ssl_state"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/tls/alarm"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/tls/alarm_notify_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/tls/host"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/tls/host_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/tls/kafka_consumer"
	tlsRule "github.com/volcengine/terraform-provider-volcengine/volcengine/tls/rule"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/tls/rule_applier"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/tls/shard"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/tos/bucket_policy"

	plSecurityGroup "github.com/volcengine/terraform-provider-volcengine/volcengine/privatelink/security_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/privatelink/vpc_endpoint"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/privatelink/vpc_endpoint_connection"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/privatelink/vpc_endpoint_service"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/privatelink/vpc_endpoint_service_permission"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/privatelink/vpc_endpoint_service_resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/privatelink/vpc_endpoint_zone"

	"github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/spec"

	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mysql/rds_mysql_account"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mysql/rds_mysql_database"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/kubeconfig"

	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/prefix_list"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/autoscaling/scaling_activity"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/autoscaling/scaling_configuration"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/autoscaling/scaling_configuration_attachment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/autoscaling/scaling_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/autoscaling/scaling_group_enabler"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/autoscaling/scaling_instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/autoscaling/scaling_instance_attachment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/autoscaling/scaling_lifecycle_hook"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/autoscaling/scaling_policy"
	bioosCluster "github.com/volcengine/terraform-provider-volcengine/volcengine/bioos/cluster"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/bioos/cluster_bind"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/bioos/workspace"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cen/cen"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cen/cen_attach_instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cen/cen_bandwidth_package"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cen/cen_bandwidth_package_associate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cen/cen_grant_instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cen/cen_inter_region_bandwidth"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cen/cen_route_entry"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/acl"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/acl_entry"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/certificate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/clb"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/listener"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/rule"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/server_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/server_group_server"
	clbZone "github.com/volcengine/terraform-provider-volcengine/volcengine/clb/zone"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cr/cr_authorization_token"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cr/cr_endpoint"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cr/cr_namespace"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cr/cr_registry"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cr/cr_registry_state"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cr/cr_repository"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cr/cr_tag"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cr/cr_vpc_endpoint"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/direct_connect/direct_connect_bgp_peer"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/direct_connect/direct_connect_connection"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/direct_connect/direct_connect_gateway"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/direct_connect/direct_connect_gateway_route"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/direct_connect/direct_connect_virtual_interface"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ebs/volume"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ebs/volume_attach"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_available_resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_command"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_deployment_set"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_deployment_set_associate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_instance_state"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_instance_type"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_invocation"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_invocation_result"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_key_pair"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_key_pair_associate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_launch_template"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/image"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/zone"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/eip/eip_address"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/eip/eip_associate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/escloud/instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/escloud/region"
	esZone "github.com/volcengine/terraform-provider-volcengine/volcengine/escloud/zone"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_access_key"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_login_profile"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_policy"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_role"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_role_policy_attachment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_user"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_user_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_user_group_attachment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_user_policy_attachment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/account"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/allow_list"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/allow_list_associate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/endpoint"
	mongodbInstance "github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/instance_parameter"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/instance_parameter_log"
	mongodbRegion "github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/region"
	mongodbZone "github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/zone"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nat/dnat_entry"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nat/nat_gateway"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nat/snat_entry"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds/rds_account"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds/rds_account_privilege"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds/rds_database"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds/rds_instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds/rds_ip_list"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mysql/allowlist"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mysql/allowlist_associate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mysql/rds_mysql_instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mysql/rds_mysql_instance_readonly_node"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_v2/rds_instance_v2"

	tlsIndex "github.com/volcengine/terraform-provider-volcengine/volcengine/tls/index"
	tlsProject "github.com/volcengine/terraform-provider-volcengine/volcengine/tls/project"
	tlsTopic "github.com/volcengine/terraform-provider-volcengine/volcengine/tls/topic"

	redisAccount "github.com/volcengine/terraform-provider-volcengine/volcengine/redis/account"
	redis_allow_list "github.com/volcengine/terraform-provider-volcengine/volcengine/redis/allow_list"
	redis_allow_list_associate "github.com/volcengine/terraform-provider-volcengine/volcengine/redis/allow_list_associate"
	redis_backup "github.com/volcengine/terraform-provider-volcengine/volcengine/redis/backup"
	redis_backup_restore "github.com/volcengine/terraform-provider-volcengine/volcengine/redis/backup_restore"
	redisContinuousBackup "github.com/volcengine/terraform-provider-volcengine/volcengine/redis/continuous_backup"
	redis_endpoint "github.com/volcengine/terraform-provider-volcengine/volcengine/redis/endpoint"
	redisInstance "github.com/volcengine/terraform-provider-volcengine/volcengine/redis/instance"
	redisRegion "github.com/volcengine/terraform-provider-volcengine/volcengine/redis/region"
	redisZone "github.com/volcengine/terraform-provider-volcengine/volcengine/redis/zone"

	"github.com/volcengine/terraform-provider-volcengine/volcengine/tos/bucket"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/tos/object"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/veenedge/available_resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/veenedge/cloud_server"
	veInstance "github.com/volcengine/terraform-provider-volcengine/volcengine/veenedge/instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/veenedge/instance_types"
	veVpc "github.com/volcengine/terraform-provider-volcengine/volcengine/veenedge/vpc"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/addon"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/cluster"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/default_node_pool"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/default_node_pool_batch_attach"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/node"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/node_pool"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/support_addon"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/ipv6_address"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/ipv6_address_bandwidth"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/ipv6_gateway"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/network_acl"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/network_acl_associate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/network_interface"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/network_interface_attach"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/route_entry"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/route_table"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/route_table_associate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/security_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/security_group_rule"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/subnet"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/vpc"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpn/customer_gateway"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpn/ssl_vpn_client_cert"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpn/ssl_vpn_server"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpn/vpn_connection"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpn/vpn_gateway"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpn/vpn_gateway_route"

	"github.com/volcengine/terraform-provider-volcengine/volcengine/nas/nas_mount_point"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nas/nas_permission_group"

	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_ca_certificate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_customized_cfg"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_health_check_template"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_listener"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_listener_domain_extension"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_rule"

	"github.com/volcengine/terraform-provider-volcengine/volcengine/bandwidth_package/bandwidth_package"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/bandwidth_package/bandwidth_package_attachment"
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
			"customer_headers": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLCENGINE_CUSTOMER_HEADERS", nil),
				Description: "CUSTOMER HEADERS for Volcengine Provider",
			},
			"customer_endpoints": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLCENGINE_CUSTOMER_ENDPOINTS", nil),
				Description: "CUSTOMER ENDPOINTS for Volcengine Provider",
			},
			"proxy_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLCENGINE_PROXY_URL", nil),
				Description: "PROXY URL for Volcengine Provider",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"volcengine_vpcs":                        vpc.DataSourceVolcengineVpcs(),
			"volcengine_subnets":                     subnet.DataSourceVolcengineSubnets(),
			"volcengine_route_tables":                route_table.DataSourceVolcengineRouteTables(),
			"volcengine_route_entries":               route_entry.DataSourceVolcengineRouteEntries(),
			"volcengine_security_groups":             security_group.DataSourceVolcengineSecurityGroups(),
			"volcengine_security_group_rules":        security_group_rule.DataSourceVolcengineSecurityGroupRules(),
			"volcengine_network_interfaces":          network_interface.DataSourceVolcengineNetworkInterfaces(),
			"volcengine_network_acls":                network_acl.DataSourceVolcengineNetworkAcls(),
			"volcengine_vpc_ipv6_gateways":           ipv6_gateway.DataSourceVolcengineIpv6Gateways(),
			"volcengine_vpc_ipv6_address_bandwidths": ipv6_address_bandwidth.DataSourceVolcengineIpv6AddressBandwidths(),
			"volcengine_vpc_ipv6_addresses":          ipv6_address.DataSourceVolcengineIpv6Addresses(),
			"volcengine_vpc_prefix_lists":            prefix_list.DataSourceVolcengineVpcPrefixLists(),
			"volcengine_ha_vips":                     ha_vip.DataSourceVolcengineHaVips(),

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
			"volcengine_clb_zones":            clbZone.DataSourceVolcengineClbZones(),

			// ================ EBS ================
			"volcengine_volumes": volume.DataSourceVolcengineVolumes(),

			// ================ ECS ================
			"volcengine_ecs_instances":           ecs_instance.DataSourceVolcengineEcsInstances(),
			"volcengine_images":                  image.DataSourceVolcengineImages(),
			"volcengine_zones":                   zone.DataSourceVolcengineZones(),
			"volcengine_regions":                 regions.DataSourceVolcengineRegions(),
			"volcengine_ecs_deployment_sets":     ecs_deployment_set.DataSourceVolcengineEcsDeploymentSets(),
			"volcengine_ecs_key_pairs":           ecs_key_pair.DataSourceVolcengineEcsKeyPairs(),
			"volcengine_ecs_launch_templates":    ecs_launch_template.DataSourceVolcengineEcsLaunchTemplates(),
			"volcengine_ecs_commands":            ecs_command.DataSourceVolcengineEcsCommands(),
			"volcengine_ecs_invocations":         ecs_invocation.DataSourceVolcengineEcsInvocations(),
			"volcengine_ecs_invocation_results":  ecs_invocation_result.DataSourceVolcengineEcsInvocationResults(),
			"volcengine_ecs_available_resources": ecs_available_resource.DataSourceVolcengineEcsAvailableResources(),
			"volcengine_ecs_instance_types":      ecs_instance_type.DataSourceVolcengineEcsInstanceTypes(),

			// ================ NAT ================
			"volcengine_snat_entries": snat_entry.DataSourceVolcengineSnatEntries(),
			"volcengine_nat_gateways": nat_gateway.DataSourceVolcengineNatGateways(),
			"volcengine_dnat_entries": dnat_entry.DataSourceVolcengineDnatEntries(),

			// ================ AutoScaling ================
			"volcengine_scaling_groups":          scaling_group.DataSourceVolcengineScalingGroups(),
			"volcengine_scaling_configurations":  scaling_configuration.DataSourceVolcengineScalingConfigurations(),
			"volcengine_scaling_policies":        scaling_policy.DataSourceVolcengineScalingPolicies(),
			"volcengine_scaling_activities":      scaling_activity.DataSourceVolcengineScalingActivities(),
			"volcengine_scaling_lifecycle_hooks": scaling_lifecycle_hook.DataSourceVolcengineScalingLifecycleHooks(),
			"volcengine_scaling_instances":       scaling_instance.DataSourceVolcengineScalingInstances(),

			// ================ Cen ================
			"volcengine_cens":                        cen.DataSourceVolcengineCens(),
			"volcengine_cen_attach_instances":        cen_attach_instance.DataSourceVolcengineCenAttachInstances(),
			"volcengine_cen_bandwidth_packages":      cen_bandwidth_package.DataSourceVolcengineCenBandwidthPackages(),
			"volcengine_cen_inter_region_bandwidths": cen_inter_region_bandwidth.DataSourceVolcengineCenInterRegionBandwidths(),
			"volcengine_cen_service_route_entries":   cen_service_route_entry.DataSourceVolcengineCenServiceRouteEntries(),
			"volcengine_cen_route_entries":           cen_route_entry.DataSourceVolcengineCenRouteEntries(),

			// ================ VPN ================
			"volcengine_vpn_gateways":         vpn_gateway.DataSourceVolcengineVpnGateways(),
			"volcengine_customer_gateways":    customer_gateway.DataSourceVolcengineCustomerGateways(),
			"volcengine_vpn_connections":      vpn_connection.DataSourceVolcengineVpnConnections(),
			"volcengine_vpn_gateway_routes":   vpn_gateway_route.DataSourceVolcengineVpnGatewayRoutes(),
			"volcengine_ssl_vpn_servers":      ssl_vpn_server.DataSourceVolcengineSslVpnServers(),
			"volcengine_ssl_vpn_client_certs": ssl_vpn_client_cert.DataSourceVolcengineSslVpnClientCerts(),

			// ================ VKE ================
			"volcengine_vke_nodes":                  node.DataSourceVolcengineVkeNodes(),
			"volcengine_vke_clusters":               cluster.DataSourceVolcengineVkeVkeClusters(),
			"volcengine_vke_node_pools":             node_pool.DataSourceVolcengineNodePools(),
			"volcengine_vke_addons":                 addon.DataSourceVolcengineVkeAddons(),
			"volcengine_vke_support_addons":         support_addon.DataSourceVolcengineVkeVkeSupportedAddons(),
			"volcengine_vke_kubeconfigs":            kubeconfig.DataSourceVolcengineVkeKubeconfigs(),
			"volcengine_vke_support_resource_types": support_resource_types.DataSourceVolcengineVkeVkeSupportResourceTypes(),

			// ================ IAM ================
			"volcengine_iam_policies":                      iam_policy.DataSourceVolcengineIamPolicies(),
			"volcengine_iam_roles":                         iam_role.DataSourceVolcengineIamRoles(),
			"volcengine_iam_users":                         iam_user.DataSourceVolcengineIamUsers(),
			"volcengine_iam_user_groups":                   iam_user_group.DataSourceVolcengineIamUserGroups(),
			"volcengine_iam_user_group_policy_attachments": iam_user_group_policy_attachment.DataSourceVolcengineIamUserGroupPolicyAttachments(),

			// ================ RDS V1 ==============
			"volcengine_rds_instances":           rds_instance.DataSourceVolcengineRdsInstances(),
			"volcengine_rds_databases":           rds_database.DataSourceVolcengineRdsDatabases(),
			"volcengine_rds_accounts":            rds_account.DataSourceVolcengineRdsAccounts(),
			"volcengine_rds_ip_lists":            rds_ip_list.DataSourceVolcengineRdsIpLists(),
			"volcengine_rds_parameter_templates": rds_parameter_template.DataSourceVolcengineRdsParameterTemplates(),

			// ================ ESCloud =============
			"volcengine_escloud_instances": instance.DataSourceVolcengineESCloudInstances(),
			"volcengine_escloud_regions":   region.DataSourceVolcengineESCloudRegions(),
			"volcengine_escloud_zones":     esZone.DataSourceVolcengineESCloudZones(),

			// ================ TOS ================
			"volcengine_tos_buckets": bucket.DataSourceVolcengineTosBuckets(),
			"volcengine_tos_objects": object.DataSourceVolcengineTosObjects(),

			// ================ Redis =============
			"volcengine_redis_allow_lists":       redis_allow_list.DataSourceVolcengineRedisAllowLists(),
			"volcengine_redis_backups":           redis_backup.DataSourceVolcengineRedisBackups(),
			"volcengine_redis_regions":           redisRegion.DataSourceVolcengineRedisRegions(),
			"volcengine_redis_zones":             redisZone.DataSourceVolcengineRedisZones(),
			"volcengine_redis_accounts":          redisAccount.DataSourceVolcengineRedisAccounts(),
			"volcengine_redis_instances":         redisInstance.DataSourceVolcengineRedisDbInstances(),
			"volcengine_redis_pitr_time_windows": pitr_time_period.DataSourceVolcengineRedisPitrTimeWindows(),

			// ================ CR ================
			"volcengine_cr_registries":           cr_registry.DataSourceVolcengineCrRegistries(),
			"volcengine_cr_namespaces":           cr_namespace.DataSourceVolcengineCrNamespaces(),
			"volcengine_cr_repositories":         cr_repository.DataSourceVolcengineCrRepositories(),
			"volcengine_cr_tags":                 cr_tag.DataSourceVolcengineCrTags(),
			"volcengine_cr_authorization_tokens": cr_authorization_token.DataSourceVolcengineCrAuthorizationTokens(),
			"volcengine_cr_endpoints":            cr_endpoint.DataSourceVolcengineCrEndpoints(),
			"volcengine_cr_vpc_endpoints":        cr_vpc_endpoint.DataSourceVolcengineCrVpcEndpoints(),

			// ================ Veenedge ================
			"volcengine_veenedge_cloud_servers":       cloud_server.DataSourceVolcengineVeenedgeCloudServers(),
			"volcengine_veenedge_instances":           veInstance.DataSourceVolcengineInstances(),
			"volcengine_veenedge_instance_types":      instance_types.DataSourceVolcengineInstanceTypes(),
			"volcengine_veenedge_available_resources": available_resource.DataSourceVolcengineAvailableResources(),
			"volcengine_veenedge_vpcs":                veVpc.DataSourceVolcengineVpcs(),

			// ================ MongoDB =============
			"volcengine_mongodb_instances":               mongodbInstance.DataSourceVolcengineMongoDBInstances(),
			"volcengine_mongodb_endpoints":               endpoint.DataSourceVolcengineMongoDBEndpoints(),
			"volcengine_mongodb_allow_lists":             allow_list.DataSourceVolcengineMongoDBAllowLists(),
			"volcengine_mongodb_instance_parameters":     instance_parameter.DataSourceVolcengineMongoDBInstanceParameters(),
			"volcengine_mongodb_instance_parameter_logs": instance_parameter_log.DataSourceVolcengineMongoDBInstanceParameterLogs(),
			"volcengine_mongodb_regions":                 mongodbRegion.DataSourceVolcengineMongoDBRegions(),
			"volcengine_mongodb_zones":                   mongodbZone.DataSourceVolcengineMongoDBZones(),
			"volcengine_mongodb_accounts":                account.DataSourceVolcengineMongoDBAccounts(),
			"volcengine_mongodb_specs":                   spec.DataSourceVolcengineMongoDBSpecs(),
			"volcengine_mongodb_ssl_states":              ssl_state.DataSourceVolcengineMongoDBSSLStates(),

			// ================ Bioos ==================
			"volcengine_bioos_clusters":   bioosCluster.DataSourceVolcengineBioosClusters(),
			"volcengine_bioos_workspaces": workspace.DataSourceVolcengineBioosWorkspaces(),

			// ================ PrivateLink ==================
			"volcengine_privatelink_vpc_endpoints":                    vpc_endpoint.DataSourceVolcenginePrivatelinkVpcEndpoints(),
			"volcengine_privatelink_vpc_endpoint_services":            vpc_endpoint_service.DataSourceVolcenginePrivatelinkVpcEndpointServices(),
			"volcengine_privatelink_vpc_endpoint_service_permissions": vpc_endpoint_service_permission.DataSourceVolcenginePrivatelinkVpcEndpointServicePermissions(),
			"volcengine_privatelink_vpc_endpoint_connections":         vpc_endpoint_connection.DataSourceVolcenginePrivatelinkVpcEndpointConnections(),
			"volcengine_privatelink_vpc_endpoint_zones":               vpc_endpoint_zone.DataSourceVolcenginePrivatelinkVpcEndpointZones(),

			// ================ RDS V2 ==============
			"volcengine_rds_instances_v2": rds_instance_v2.DataSourceVolcengineRdsInstances(),

			// ================ RdsMysql ================
			"volcengine_rds_mysql_instances":  rds_mysql_instance.DataSourceVolcengineRdsMysqlInstances(),
			"volcengine_rds_mysql_accounts":   rds_mysql_account.DataSourceVolcengineRdsMysqlAccounts(),
			"volcengine_rds_mysql_databases":  rds_mysql_database.DataSourceVolcengineRdsMysqlDatabases(),
			"volcengine_rds_mysql_allowlists": allowlist.DataSourceVolcengineRdsMysqlAllowLists(),

			// ================ TLS ================
			"volcengine_tls_rules":               tlsRule.DataSourceVolcengineTlsRules(),
			"volcengine_tls_alarms":              alarm.DataSourceVolcengineTlsAlarms(),
			"volcengine_tls_alarm_notify_groups": alarm_notify_group.DataSourceVolcengineTlsAlarmNotifyGroups(),
			"volcengine_tls_rule_appliers":       rule_applier.DataSourceVolcengineTlsRuleAppliers(),
			"volcengine_tls_shards":              shard.DataSourceVolcengineTlsShards(),
			"volcengine_tls_kafka_consumers":     kafka_consumer.DataSourceVolcengineTlsKafkaConsumers(),
			"volcengine_tls_host_groups":         host_group.DataSourceVolcengineTlsHostGroups(),
			"volcengine_tls_hosts":               host.DataSourceVolcengineTlsHosts(),
			"volcengine_tls_projects":            tlsProject.DataSourceVolcengineTlsProjects(),
			"volcengine_tls_topics":              tlsTopic.DataSourceVolcengineTlsTopics(),
			"volcengine_tls_indexes":             tlsIndex.DataSourceVolcengineTlsIndexes(),

			// ================ Cloudfs ================
			"volcengine_cloudfs_quotas":       cloudfs_quota.DataSourceVolcengineCloudfsQuotas(),
			"volcengine_cloudfs_file_systems": cloudfs_file_system.DataSourceVolcengineCloudfsFileSystems(),
			"volcengine_cloudfs_accesses":     cloudfs_access.DataSourceVolcengineCloudfsAccesses(),
			"volcengine_cloudfs_ns_quotas":    cloudfs_ns_quota.DataSourceVolcengineCloudfsNsQuotas(),
			"volcengine_cloudfs_namespaces":   cloudfs_namespace.DataSourceVolcengineCloudfsNamespaces(),

			// ================ NAS ================
			"volcengine_nas_file_systems":      nas_file_system.DataSourceVolcengineNasFileSystems(),
			"volcengine_nas_snapshots":         nas_snapshot.DataSourceVolcengineNasSnapshots(),
			"volcengine_nas_zones":             nas_zone.DataSourceVolcengineNasZones(),
			"volcengine_nas_regions":           nas_region.DataSourceVolcengineNasRegions(),
			"volcengine_nas_mount_points":      nas_mount_point.DataSourceVolcengineNasMountPoints(),
			"volcengine_nas_permission_groups": nas_permission_group.DataSourceVolcengineNasPermissionGroups(),

			// ================ TransitRouter =============
			"volcengine_transit_routers":                                   transit_router.DataSourceVolcengineTransitRouters(),
			"volcengine_transit_router_vpc_attachments":                    transit_router_vpc_attachment.DataSourceVolcengineTransitRouterVpcAttachments(),
			"volcengine_transit_router_vpn_attachments":                    transit_router_vpn_attachment.DataSourceVolcengineTransitRouterVpnAttachments(),
			"volcengine_transit_router_route_tables":                       trTable.DataSourceVolcengineTransitRouterRouteTables(),
			"volcengine_transit_router_route_entries":                      trEntry.DataSourceVolcengineTransitRouterRouteEntries(),
			"volcengine_transit_router_route_table_associations":           route_table_association.DataSourceVolcengineTransitRouterRouteTableAssociations(),
			"volcengine_transit_router_route_table_propagations":           route_table_propagation.DataSourceVolcengineTransitRouterRouteTablePropagations(),
			"volcengine_transit_router_bandwidth_packages":                 transit_router_bandwidth_package.DataSourceVolcengineTransitRouterBandwidthPackages(),
			"volcengine_transit_router_grant_rules":                        transit_router_grant_rule.DataSourceVolcengineTransitRouterGrantRules(),
			"volcengine_transit_router_direct_connect_gateway_attachments": transit_router_direct_connect_gateway_attachment.DataSourceVolcengineTransitRouterDirectConnectGatewayAttachments(),
			"volcengine_transit_router_peer_attachments":                   transit_router_peer_attachment.DataSourceVolcengineTransitRouterPeerAttachments(),

			// ================ DirectConnect ================
			"volcengine_direct_connect_connections":        direct_connect_connection.DataSourceVolcengineDirectConnectConnections(),
			"volcengine_direct_connect_gateways":           direct_connect_gateway.DataSourceVolcengineDirectConnectGateways(),
			"volcengine_direct_connect_virtual_interfaces": direct_connect_virtual_interface.DataSourceVolcengineDirectConnectVirtualInterfaces(),
			"volcengine_direct_connect_bgp_peers":          direct_connect_bgp_peer.DataSourceVolcengineDirectConnectBgpPeers(),
			"volcengine_direct_connect_gateway_routes":     direct_connect_gateway_route.DataSourceVolcengineDirectConnectGatewayRoutes(),

			// ================ ALB ================
			"volcengine_alb_zones":                      alb_zone.DataSourceVolcengineAlbZones(),
			"volcengine_alb_acls":                       alb_acl.DataSourceVolcengineAlbAcls(),
			"volcengine_alb_listeners":                  alb_listener.DataSourceVolcengineListeners(),
			"volcengine_alb_customized_cfgs":            alb_customized_cfg.DataSourceVolcengineAlbCustomizedCfgs(),
			"volcengine_alb_health_check_templates":     alb_health_check_template.DataSourceVolcengineAlbHealthCheckTemplates(),
			"volcengine_alb_listener_domain_extensions": alb_listener_domain_extension.DataSourceVolcengineListenerDomainExtensions(),
			"volcengine_alb_server_group_servers":       alb_server_group_server.DataSourceVolcengineAlbServerGroupServers(),
			"volcengine_alb_certificates":               alb_certificate.DataSourceVolcengineAlbCertificates(),
			"volcengine_alb_rules":                      alb_rule.DataSourceVolcengineAlbRules(),
			"volcengine_alb_ca_certificates":            alb_ca_certificate.DataSourceVolcengineAlbCaCertificates(),
			"volcengine_albs":                           alb.DataSourceVolcengineAlbs(),
			"volcengine_alb_server_groups":              alb_server_group.DataSourceVolcengineAlbServerGroups(),

			// ============= Bandwidth Package =============
			"volcengine_bandwidth_packages": bandwidth_package.DataSourceVolcengineBandwidthPackages(),

			// ============= Cloud Monitor =============
			"volcengine_cloud_monitor_contacts":       cloud_monitor_contact.DataSourceVolcengineCloudMonitorContacts(),
			"volcengine_cloud_monitor_contact_groups": cloud_monitor_contact_group.DataSourceVolcengineCloudMonitorContactGroups(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"volcengine_vpc":                        vpc.ResourceVolcengineVpc(),
			"volcengine_subnet":                     subnet.ResourceVolcengineSubnet(),
			"volcengine_route_table":                route_table.ResourceVolcengineRouteTable(),
			"volcengine_route_entry":                route_entry.ResourceVolcengineRouteEntry(),
			"volcengine_route_table_associate":      route_table_associate.ResourceVolcengineRouteTableAssociate(),
			"volcengine_security_group":             security_group.ResourceVolcengineSecurityGroup(),
			"volcengine_network_interface":          network_interface.ResourceVolcengineNetworkInterface(),
			"volcengine_network_interface_attach":   network_interface_attach.ResourceVolcengineNetworkInterfaceAttach(),
			"volcengine_security_group_rule":        security_group_rule.ResourceVolcengineSecurityGroupRule(),
			"volcengine_network_acl":                network_acl.ResourceVolcengineNetworkAcl(),
			"volcengine_network_acl_associate":      network_acl_associate.ResourceVolcengineNetworkAclAssociate(),
			"volcengine_vpc_ipv6_gateway":           ipv6_gateway.ResourceVolcengineIpv6Gateway(),
			"volcengine_vpc_ipv6_address_bandwidth": ipv6_address_bandwidth.ResourceVolcengineIpv6AddressBandwidth(),
			"volcengine_vpc_prefix_list":            prefix_list.ResourceVolcengineVpcPrefixList(),
			"volcengine_ha_vip":                     ha_vip.ResourceVolcengineHaVip(),
			"volcengine_ha_vip_associate":           ha_vip_associate.ResourceVolcengineHaVipAssociate(),

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
			"volcengine_ecs_instance":                 ecs_instance.ResourceVolcengineEcsInstance(),
			"volcengine_ecs_instance_state":           ecs_instance_state.ResourceVolcengineEcsInstanceState(),
			"volcengine_ecs_deployment_set":           ecs_deployment_set.ResourceVolcengineEcsDeploymentSet(),
			"volcengine_ecs_deployment_set_associate": ecs_deployment_set_associate.ResourceVolcengineEcsDeploymentSetAssociate(),
			"volcengine_ecs_key_pair":                 ecs_key_pair.ResourceVolcengineEcsKeyPair(),
			"volcengine_ecs_key_pair_associate":       ecs_key_pair_associate.ResourceVolcengineEcsKeyPairAssociate(),
			"volcengine_ecs_launch_template":          ecs_launch_template.ResourceVolcengineEcsLaunchTemplate(),
			"volcengine_ecs_command":                  ecs_command.ResourceVolcengineEcsCommand(),
			"volcengine_ecs_invocation":               ecs_invocation.ResourceVolcengineEcsInvocation(),
			"volcengine_iam_role_attachment":          ecs_iam_role_attachment.ResourceVolcengineIamRoleAttachment(),

			// ================ NAT ================
			"volcengine_snat_entry":  snat_entry.ResourceVolcengineSnatEntry(),
			"volcengine_nat_gateway": nat_gateway.ResourceVolcengineNatGateway(),
			"volcengine_dnat_entry":  dnat_entry.ResourceVolcengineDnatEntry(),

			// ================ AutoScaling ================
			"volcengine_scaling_group":                    scaling_group.ResourceVolcengineScalingGroup(),
			"volcengine_scaling_configuration":            scaling_configuration.ResourceVolcengineScalingConfiguration(),
			"volcengine_scaling_policy":                   scaling_policy.ResourceVolcengineScalingPolicy(),
			"volcengine_scaling_instance_attachment":      scaling_instance_attachment.ResourceVolcengineScalingInstanceAttachment(),
			"volcengine_scaling_lifecycle_hook":           scaling_lifecycle_hook.ResourceVolcengineScalingLifecycleHook(),
			"volcengine_scaling_group_enabler":            scaling_group_enabler.ResourceVolcengineScalingGroupEnabler(),
			"volcengine_scaling_configuration_attachment": scaling_configuration_attachment.ResourceVolcengineScalingConfigurationAttachment(),

			// ================ Cen ================
			"volcengine_cen":                             cen.ResourceVolcengineCen(),
			"volcengine_cen_attach_instance":             cen_attach_instance.ResourceVolcengineCenAttachInstance(),
			"volcengine_cen_grant_instance":              cen_grant_instance.ResourceVolcengineCenGrantInstance(),
			"volcengine_cen_bandwidth_package":           cen_bandwidth_package.ResourceVolcengineCenBandwidthPackage(),
			"volcengine_cen_bandwidth_package_associate": cen_bandwidth_package_associate.ResourceVolcengineCenBandwidthPackageAssociate(),
			"volcengine_cen_inter_region_bandwidth":      cen_inter_region_bandwidth.ResourceVolcengineCenInterRegionBandwidth(),
			"volcengine_cen_service_route_entry":         cen_service_route_entry.ResourceVolcengineCenServiceRouteEntry(),
			"volcengine_cen_route_entry":                 cen_route_entry.ResourceVolcengineCenRouteEntry(),

			// ================ VPN ================
			"volcengine_vpn_gateway":         vpn_gateway.ResourceVolcengineVpnGateway(),
			"volcengine_customer_gateway":    customer_gateway.ResourceVolcengineCustomerGateway(),
			"volcengine_vpn_connection":      vpn_connection.ResourceVolcengineVpnConnection(),
			"volcengine_vpn_gateway_route":   vpn_gateway_route.ResourceVolcengineVpnGatewayRoute(),
			"volcengine_ssl_vpn_server":      ssl_vpn_server.ResourceVolcengineSslVpnServer(),
			"volcengine_ssl_vpn_client_cert": ssl_vpn_client_cert.ResourceVolcengineSslClientCertServer(),

			// ================ VKE ================
			"volcengine_vke_node":                           node.ResourceVolcengineVkeNode(),
			"volcengine_vke_cluster":                        cluster.ResourceVolcengineVkeCluster(),
			"volcengine_vke_node_pool":                      node_pool.ResourceVolcengineNodePool(),
			"volcengine_vke_addon":                          addon.ResourceVolcengineVkeAddon(),
			"volcengine_vke_default_node_pool":              default_node_pool.ResourceVolcengineDefaultNodePool(),
			"volcengine_vke_default_node_pool_batch_attach": default_node_pool_batch_attach.ResourceVolcengineDefaultNodePoolBatchAttach(),
			"volcengine_vke_kubeconfig":                     kubeconfig.ResourceVolcengineVkeKubeconfig(),

			// ================ IAM ================
			"volcengine_iam_policy":                       iam_policy.ResourceVolcengineIamPolicy(),
			"volcengine_iam_role":                         iam_role.ResourceVolcengineIamRole(),
			"volcengine_iam_role_policy_attachment":       iam_role_policy_attachment.ResourceVolcengineIamRolePolicyAttachment(),
			"volcengine_iam_access_key":                   iam_access_key.ResourceVolcengineIamAccessKey(),
			"volcengine_iam_user":                         iam_user.ResourceVolcengineIamUser(),
			"volcengine_iam_login_profile":                iam_login_profile.ResourceVolcengineIamLoginProfile(),
			"volcengine_iam_user_policy_attachment":       iam_user_policy_attachment.ResourceVolcengineIamUserPolicyAttachment(),
			"volcengine_iam_user_group":                   iam_user_group.ResourceVolcengineIamUserGroup(),
			"volcengine_iam_user_group_attachment":        iam_user_group_attachment.ResourceVolcengineIamUserGroupAttachment(),
			"volcengine_iam_user_group_policy_attachment": iam_user_group_policy_attachment.ResourceVolcengineIamUserGroupPolicyAttachment(),

			// ================ RDS V1 ==============
			"volcengine_rds_instance":           rds_instance.ResourceVolcengineRdsInstance(),
			"volcengine_rds_database":           rds_database.ResourceVolcengineRdsDatabase(),
			"volcengine_rds_account":            rds_account.ResourceVolcengineRdsAccount(),
			"volcengine_rds_ip_list":            rds_ip_list.ResourceVolcengineRdsIpList(),
			"volcengine_rds_account_privilege":  rds_account_privilege.ResourceVolcengineRdsAccountPrivilege(),
			"volcengine_rds_parameter_template": rds_parameter_template.ResourceVolcengineRdsParameterTemplate(),

			// ================ ESCloud ================
			"volcengine_escloud_instance": instance.ResourceVolcengineESCloudInstance(),

			//================= TOS =================
			"volcengine_tos_bucket":        bucket.ResourceVolcengineTosBucket(),
			"volcengine_tos_object":        object.ResourceVolcengineTosObject(),
			"volcengine_tos_bucket_policy": bucket_policy.ResourceVolcengineTosBucketPolicy(),

			// ================ Redis ==============
			"volcengine_redis_allow_list":           redis_allow_list.ResourceVolcengineRedisAllowList(),
			"volcengine_redis_endpoint":             redis_endpoint.ResourceVolcengineRedisEndpoint(),
			"volcengine_redis_allow_list_associate": redis_allow_list_associate.ResourceVolcengineRedisAllowListAssociate(),
			"volcengine_redis_backup":               redis_backup.ResourceVolcengineRedisBackup(),
			"volcengine_redis_backup_restore":       redis_backup_restore.ResourceVolcengineRedisBackupRestore(),
			"volcengine_redis_account":              redisAccount.ResourceVolcengineRedisAccount(),
			"volcengine_redis_instance":             redisInstance.ResourceVolcengineRedisDbInstance(),
			"volcengine_redis_instance_state":       instance_state.ResourceVolcengineRedisInstanceState(),
			"volcengine_redis_continuous_backup":    redisContinuousBackup.ResourceVolcengineRedisContinuousBackup(),

			// ================ CR ================
			"volcengine_cr_registry":       cr_registry.ResourceVolcengineCrRegistry(),
			"volcengine_cr_registry_state": cr_registry_state.ResourceVolcengineCrRegistryState(),
			"volcengine_cr_namespace":      cr_namespace.ResourceVolcengineCrNamespace(),
			"volcengine_cr_repository":     cr_repository.ResourceVolcengineCrRepository(),
			"volcengine_cr_tag":            cr_tag.ResourceVolcengineCrTag(),
			"volcengine_cr_endpoint":       cr_endpoint.ResourceVolcengineCrEndpoint(),
			"volcengine_cr_vpc_endpoint":   cr_vpc_endpoint.ResourceVolcengineCrVpcEndpoint(),

			// ================ Veenedge ================
			"volcengine_veenedge_cloud_server": cloud_server.ResourceVolcengineCloudServer(),
			"volcengine_veenedge_instance":     veInstance.ResourceVolcengineInstance(),
			"volcengine_veenedge_vpc":          veVpc.ResourceVolcengineVpc(),

			// ================ MongoDB ================
			"volcengine_mongodb_instance":             mongodbInstance.ResourceVolcengineMongoDBInstance(),
			"volcengine_mongodb_endpoint":             endpoint.ResourceVolcengineMongoDBEndpoint(),
			"volcengine_mongodb_allow_list":           allow_list.ResourceVolcengineMongoDBAllowList(),
			"volcengine_mongodb_instance_parameter":   instance_parameter.ResourceVolcengineMongoDBInstanceParameter(),
			"volcengine_mongodb_allow_list_associate": allow_list_associate.ResourceVolcengineMongodbAllowListAssociate(),
			"volcengine_mongodb_ssl_state":            ssl_state.ResourceVolcengineMongoDBSSLState(),

			// ================ Bioos ================
			"volcengine_bioos_cluster":      bioosCluster.ResourceVolcengineBioosCluster(),
			"volcengine_bioos_workspace":    workspace.ResourceVolcengineBioosWorkspace(),
			"volcengine_bioos_cluster_bind": cluster_bind.ResourceVolcengineBioosClusterBind(),

			// ================ PrivateLink ==================
			"volcengine_privatelink_vpc_endpoint":                    vpc_endpoint.ResourceVolcenginePrivatelinkVpcEndpoint(),
			"volcengine_privatelink_vpc_endpoint_service":            vpc_endpoint_service.ResourceVolcenginePrivatelinkVpcEndpointService(),
			"volcengine_privatelink_vpc_endpoint_service_resource":   vpc_endpoint_service_resource.ResourceVolcenginePrivatelinkVpcEndpointServiceResource(),
			"volcengine_privatelink_vpc_endpoint_service_permission": vpc_endpoint_service_permission.ResourceVolcenginePrivatelinkVpcEndpointServicePermission(),
			"volcengine_privatelink_security_group":                  plSecurityGroup.ResourceVolcenginePrivatelinkSecurityGroupService(),
			"volcengine_privatelink_vpc_endpoint_connection":         vpc_endpoint_connection.ResourceVolcenginePrivatelinkVpcEndpointConnectionService(),
			"volcengine_privatelink_vpc_endpoint_zone":               vpc_endpoint_zone.ResourceVolcenginePrivatelinkVpcEndpointZone(),

			// ================ RDS V2 ==============
			"volcengine_rds_instance_v2": rds_instance_v2.ResourceVolcengineRdsInstance(),

			// ================ RdsMysql ================
			"volcengine_rds_mysql_instance":               rds_mysql_instance.ResourceVolcengineRdsMysqlInstance(),
			"volcengine_rds_mysql_instance_readonly_node": rds_mysql_instance_readonly_node.ResourceVolcengineRdsMysqlInstanceReadonlyNode(),
			"volcengine_rds_mysql_account":                rds_mysql_account.ResourceVolcengineRdsMysqlAccount(),
			"volcengine_rds_mysql_database":               rds_mysql_database.ResourceVolcengineRdsMysqlDatabase(),
			"volcengine_rds_mysql_allowlist":              allowlist.ResourceVolcengineRdsMysqlAllowlist(),
			"volcengine_rds_mysql_allowlist_associate":    allowlist_associate.ResourceVolcengineRdsMysqlAllowlistAssociate(),

			// ================ TLS ================
			"volcengine_tls_kafka_consumer":     kafka_consumer.ResourceVolcengineTlsKafkaConsumer(),
			"volcengine_tls_host_group":         host_group.ResourceVolcengineTlsHostGroup(),
			"volcengine_tls_rule":               tlsRule.ResourceVolcengineTlsRule(),
			"volcengine_tls_rule_applier":       rule_applier.ResourceVolcengineTlsRuleApplier(),
			"volcengine_tls_alarm":              alarm.ResourceVolcengineTlsAlarm(),
			"volcengine_tls_alarm_notify_group": alarm_notify_group.ResourceVolcengineTlsAlarmNotifyGroup(),
			"volcengine_tls_host":               host.ResourceVolcengineTlsHost(),
			"volcengine_tls_project":            tlsProject.ResourceVolcengineTlsProject(),
			"volcengine_tls_topic":              tlsTopic.ResourceVolcengineTlsTopic(),
			"volcengine_tls_index":              tlsIndex.ResourceVolcengineTlsIndex(),

			// ================ Cloudfs ================
			"volcengine_cloudfs_file_system": cloudfs_file_system.ResourceVolcengineCloudfsFileSystem(),
			"volcengine_cloudfs_access":      cloudfs_access.ResourceVolcengineCloudfsAccess(),
			"volcengine_cloudfs_namespace":   cloudfs_namespace.ResourceVolcengineCloudfsNamespace(),

			// ================ NAS ================
			"volcengine_nas_file_system":      nas_file_system.ResourceVolcengineNasFileSystem(),
			"volcengine_nas_snapshot":         nas_snapshot.ResourceVolcengineNasSnapshot(),
			"volcengine_nas_mount_point":      nas_mount_point.ResourceVolcengineNasMountPoint(),
			"volcengine_nas_permission_group": nas_permission_group.ResourceVolcengineNasPermissionGroup(),

			// ================ TransitRouter =============
			"volcengine_transit_router":                                   transit_router.ResourceVolcengineTransitRouter(),
			"volcengine_transit_router_vpc_attachment":                    transit_router_vpc_attachment.ResourceVolcengineTransitRouterVpcAttachment(),
			"volcengine_transit_router_vpn_attachment":                    transit_router_vpn_attachment.ResourceVolcengineTransitRouterVpnAttachment(),
			"volcengine_transit_router_route_table":                       trTable.ResourceVolcengineTransitRouterRouteTable(),
			"volcengine_transit_router_route_entry":                       trEntry.ResourceVolcengineTransitRouterRouteEntry(),
			"volcengine_transit_router_route_table_association":           route_table_association.ResourceVolcengineTransitRouterRouteTableAssociation(),
			"volcengine_transit_router_route_table_propagation":           route_table_propagation.ResourceVolcengineTransitRouterRouteTablePropagation(),
			"volcengine_transit_router_bandwidth_package":                 transit_router_bandwidth_package.ResourceVolcengineTransitRouterBandwidthPackage(),
			"volcengine_transit_router_grant_rule":                        transit_router_grant_rule.ResourceVolcengineTransitRouterGrantRule(),
			"volcengine_transit_router_direct_connect_gateway_attachment": transit_router_direct_connect_gateway_attachment.ResourceVolcengineTransitRouterDirectConnectGatewayAttachment(),
			"volcengine_transit_router_shared_transit_router_state":       shared_transit_router_state.ResourceVolcengineSharedTransitRouterState(),
			"volcengine_transit_router_peer_attachment":                   transit_router_peer_attachment.ResourceVolcengineTransitRouterPeerAttachment(),

			// ================ DirectConnect ================
			"volcengine_direct_connect_connection":        direct_connect_connection.ResourceVolcengineDirectConnectConnection(),
			"volcengine_direct_connect_gateway":           direct_connect_gateway.ResourceVolcengineDirectConnectGateway(),
			"volcengine_direct_connect_virtual_interface": direct_connect_virtual_interface.ResourceVolcengineDirectConnectVirtualInterface(),
			"volcengine_direct_connect_bgp_peer":          direct_connect_bgp_peer.ResourceVolcengineDirectConnectBgpPeer(),
			"volcengine_direct_connect_gateway_route":     direct_connect_gateway_route.ResourceVolcengineDirectConnectGatewayRoute(),

			// ================ ALB ================
			"volcengine_alb_acl":                       alb_acl.ResourceVolcengineAlbAcl(),
			"volcengine_alb_listener":                  alb_listener.ResourceVolcengineAlbListener(),
			"volcengine_alb_customized_cfg":            alb_customized_cfg.ResourceVolcengineAlbCustomizedCfg(),
			"volcengine_alb_health_check_template":     alb_health_check_template.ResourceVolcengineAlbHealthCheckTemplate(),
			"volcengine_alb_listener_domain_extension": alb_listener_domain_extension.ResourceVolcengineAlbListenerDomainExtension(),
			"volcengine_alb_server_group_server":       alb_server_group_server.ResourceVolcengineAlbServerGroupServer(),
			"volcengine_alb_certificate":               alb_certificate.ResourceVolcengineAlbCertificate(),
			"volcengine_alb_rule":                      alb_rule.ResourceVolcengineAlbRule(),
			"volcengine_alb_ca_certificate":            alb_ca_certificate.ResourceVolcengineAlbCaCertificate(),
			"volcengine_alb":                           alb.ResourceVolcengineAlb(),
			"volcengine_alb_server_group":              alb_server_group.ResourceVolcengineAlbServerGroup(),

			// ============= Bandwidth Package =============
			"volcengine_bandwidth_package":            bandwidth_package.ResourceVolcengineBandwidthPackage(),
			"volcengine_bandwidth_package_attachment": bandwidth_package_attachment.ResourceVolcengineBandwidthPackageAttachment(),

			// ============= Cloud Monitor =============
			"volcengine_cloud_monitor_contact":       cloud_monitor_contact.ResourceVolcengineCloudMonitorContact(),
			"volcengine_cloud_monitor_contact_group": cloud_monitor_contact_group.ResourceVolcengineCloudMonitorContactGroup(),
		},
		ConfigureFunc: ProviderConfigure,
	}
}

func ProviderConfigure(d *schema.ResourceData) (interface{}, error) {
	config := ve.Config{
		AccessKey:         d.Get("access_key").(string),
		SecretKey:         d.Get("secret_key").(string),
		SessionToken:      d.Get("session_token").(string),
		Region:            d.Get("region").(string),
		Endpoint:          d.Get("endpoint").(string),
		DisableSSL:        d.Get("disable_ssl").(bool),
		CustomerHeaders:   map[string]string{},
		CustomerEndpoints: defaultCustomerEndPoints(),
		ProxyUrl:          d.Get("proxy_url").(string),
	}

	headers := d.Get("customer_headers").(string)
	if headers != "" {
		hs1 := strings.Split(headers, ",")
		for _, hh := range hs1 {
			hs2 := strings.Split(hh, ":")
			if len(hs2) == 2 {
				config.CustomerHeaders[hs2[0]] = hs2[1]
			}
		}
	}

	endpoints := d.Get("customer_endpoints").(string)
	if endpoints != "" {
		ends := strings.Split(endpoints, ",")
		for _, end := range ends {
			point := strings.Split(end, ":")
			if len(point) == 2 {
				config.CustomerEndpoints[point[0]] = point[1]
			}
		}
	}

	client, err := config.Client()
	return client, err
}

func defaultCustomerEndPoints() map[string]string {
	return map[string]string{
		"veenedge": "veenedge.volcengineapi.com",
	}
}
