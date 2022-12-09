package volcengine

import (
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/kubeconfig"
	"strings"

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
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cr/cr_authorization_token"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cr/cr_endpoint"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cr/cr_namespace"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cr/cr_registry"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cr/cr_registry_state"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cr/cr_repository"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cr/cr_tag"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ebs/volume"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ebs/volume_attach"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_deployment_set"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_deployment_set_associate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_instance_state"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_key_pair"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_key_pair_associate"
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
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_user_policy_attachment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nat/dnat_entry"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nat/nat_gateway"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nat/snat_entry"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds/rds_account"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds/rds_account_privilege"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds/rds_database"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds/rds_instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds/rds_ip_list"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds/rds_parameter_template"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_v2/rds_instance_v2"
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
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpn/vpn_connection"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpn/vpn_gateway"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpn/vpn_gateway_route"
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
			"proxy_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VOLCENGINE_PROXY_URL", nil),
				Description: "PROXY URL for Volcengine Provider",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"volcengine_vpcs":                 vpc.DataSourceVolcengineVpcs(),
			"volcengine_subnets":              subnet.DataSourceVolcengineSubnets(),
			"volcengine_route_tables":         route_table.DataSourceVolcengineRouteTables(),
			"volcengine_route_entries":        route_entry.DataSourceVolcengineRouteEntries(),
			"volcengine_security_groups":      security_group.DataSourceVolcengineSecurityGroups(),
			"volcengine_security_group_rules": security_group_rule.DataSourceVolcengineSecurityGroupRules(),
			"volcengine_network_interfaces":   network_interface.DataSourceVolcengineNetworkInterfaces(),

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
			"volcengine_ecs_instances":       ecs_instance.DataSourceVolcengineEcsInstances(),
			"volcengine_images":              image.DataSourceVolcengineImages(),
			"volcengine_zones":               zone.DataSourceVolcengineZones(),
			"volcengine_ecs_deployment_sets": ecs_deployment_set.DataSourceVolcengineEcsDeploymentSets(),
			"volcengine_ecs_key_pairs":       ecs_key_pair.DataSourceVolcengineEcsKeyPairs(),

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
			//"volcengine_cen_service_route_entries": 	cen_service_route_entry.DataSourceVolcengineCenServiceRouteEntries(),
			"volcengine_cen_route_entries": cen_route_entry.DataSourceVolcengineCenRouteEntries(),

			// ================ VPN ================
			"volcengine_vpn_gateways":       vpn_gateway.DataSourceVolcengineVpnGateways(),
			"volcengine_customer_gateways":  customer_gateway.DataSourceVolcengineCustomerGateways(),
			"volcengine_vpn_connections":    vpn_connection.DataSourceVolcengineVpnConnections(),
			"volcengine_vpn_gateway_routes": vpn_gateway_route.DataSourceVolcengineVpnGatewayRoutes(),

			// ================ VKE ================
			"volcengine_vke_nodes":          node.DataSourceVolcengineVkeNodes(),
			"volcengine_vke_clusters":       cluster.DataSourceVolcengineVkeVkeClusters(),
			"volcengine_vke_node_pools":     node_pool.DataSourceVolcengineNodePools(),
			"volcengine_vke_addons":         addon.DataSourceVolcengineVkeAddons(),
			"volcengine_vke_support_addons": support_addon.DataSourceVolcengineVkeVkeSupportedAddons(),
			"volcengine_vke_kubeconfigs":    kubeconfig.DataSourceVolcengineVkeKubeconfigs(),

			// ================ IAM ================
			"volcengine_iam_policies": iam_policy.DataSourceVolcengineIamPolicies(),
			"volcengine_iam_roles":    iam_role.DataSourceVolcengineIamRoles(),
			"volcengine_iam_users":    iam_user.DataSourceVolcengineIamUsers(),

			// ================ RDS V1 ==============
			"volcengine_rds_instances":           rds_instance.DataSourceVolcengineRdsInstances(),
			"volcengine_rds_databases":           rds_database.DataSourceVolcengineRdsDatabases(),
			"volcengine_rds_accounts":            rds_account.DataSourceVolcengineRdsAccounts(),
			"volcengine_rds_ip_lists":            rds_ip_list.DataSourceVolcengineRdsIpLists(),
			"volcengine_rds_parameter_templates": rds_parameter_template.DataSourceVolcengineRdsParameterTemplates(),

			// ================ RDS V2 ==============
			"volcengine_rds_instances_v2": rds_instance_v2.DataSourceVolcengineRdsInstances(),

			// ================ ESCloud =============
			"volcengine_escloud_instances": instance.DataSourceVolcengineESCloudInstances(),
			"volcengine_escloud_regions":   region.DataSourceVolcengineESCloudRegions(),
			"volcengine_escloud_zones":     esZone.DataSourceVolcengineESCloudZones(),

			// ================ TOS ================
			"volcengine_tos_buckets": bucket.DataSourceVolcengineTosBuckets(),
			"volcengine_tos_objects": object.DataSourceVolcengineTosObjects(),

			// ================ CR ================
			"volcengine_cr_registries":           cr_registry.DataSourceVolcengineCrRegistries(),
			"volcengine_cr_namespaces":           cr_namespace.DataSourceVolcengineCrNamespaces(),
			"volcengine_cr_repositories":         cr_repository.DataSourceVolcengineCrRepositories(),
			"volcengine_cr_tags":                 cr_tag.DataSourceVolcengineCrTags(),
			"volcengine_cr_authorization_tokens": cr_authorization_token.DataSourceVolcengineCrAuthorizationTokens(),
			"volcengine_cr_endpoints":            cr_endpoint.DataSourceVolcengineCrEndpoints(),

			// ================ Veenedge ================
			"volcengine_veenedge_cloud_servers":       cloud_server.DataSourceVolcengineVeenedgeCloudServers(),
			"volcengine_veenedge_instances":           veInstance.DataSourceVolcengineInstances(),
			"volcengine_veenedge_instance_types":      instance_types.DataSourceVolcengineInstanceTypes(),
			"volcengine_veenedge_available_resources": available_resource.DataSourceVolcengineAvailableResources(),
			"volcengine_veenedge_vpcs":                veVpc.DataSourceVolcengineVpcs(),
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
			"volcengine_ecs_instance":                 ecs_instance.ResourceVolcengineEcsInstance(),
			"volcengine_ecs_instance_state":           ecs_instance_state.ResourceVolcengineEcsInstanceState(),
			"volcengine_ecs_deployment_set":           ecs_deployment_set.ResourceVolcengineEcsDeploymentSet(),
			"volcengine_ecs_deployment_set_associate": ecs_deployment_set_associate.ResourceVolcengineEcsDeploymentSetAssociate(),
			"volcengine_ecs_key_pair":                 ecs_key_pair.ResourceVolcengineEcsKeyPair(),
			"volcengine_ecs_key_pair_associate":       ecs_key_pair_associate.ResourceVolcengineEcsKeyPairAssociate(),

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
			//"volcengine_cen_service_route_entry": 			cen_service_route_entry.ResourceVolcengineCenServiceRouteEntry(),
			//"volcengine_cen_route_entry": 					cen_route_entry.ResourceVolcengineCenRouteEntry(),

			// ================ VPN ================
			"volcengine_vpn_gateway":       vpn_gateway.ResourceVolcengineVpnGateway(),
			"volcengine_customer_gateway":  customer_gateway.ResourceVolcengineCustomerGateway(),
			"volcengine_vpn_connection":    vpn_connection.ResourceVolcengineVpnConnection(),
			"volcengine_vpn_gateway_route": vpn_gateway_route.ResourceVolcengineVpnGatewayRoute(),

			// ================ VKE ================
			"volcengine_vke_node":                           node.ResourceVolcengineVkeNode(),
			"volcengine_vke_cluster":                        cluster.ResourceVolcengineVkeCluster(),
			"volcengine_vke_node_pool":                      node_pool.ResourceVolcengineNodePool(),
			"volcengine_vke_addon":                          addon.ResourceVolcengineVkeAddon(),
			"volcengine_vke_default_node_pool":              default_node_pool.ResourceVolcengineDefaultNodePool(),
			"volcengine_vke_default_node_pool_batch_attach": default_node_pool_batch_attach.ResourceVolcengineDefaultNodePoolBatchAttach(),
			"volcengine_vke_kubeconfig":                     kubeconfig.ResourceVolcengineVkeKubeconfig(),

			// ================ IAM ================
			"volcengine_iam_policy":                 iam_policy.ResourceVolcengineIamPolicy(),
			"volcengine_iam_role":                   iam_role.ResourceVolcengineIamRole(),
			"volcengine_iam_role_policy_attachment": iam_role_policy_attachment.ResourceVolcengineIamRolePolicyAttachment(),
			"volcengine_iam_access_key":             iam_access_key.ResourceVolcengineIamAccessKey(),
			"volcengine_iam_user":                   iam_user.ResourceVolcengineIamUser(),
			"volcengine_iam_login_profile":          iam_login_profile.ResourceVolcengineIamLoginProfile(),
			"volcengine_iam_user_policy_attachment": iam_user_policy_attachment.ResourceVolcengineIamUserPolicyAttachment(),

			// ================ RDS V1 ==============
			"volcengine_rds_instance":           rds_instance.ResourceVolcengineRdsInstance(),
			"volcengine_rds_database":           rds_database.ResourceVolcengineRdsDatabase(),
			"volcengine_rds_account":            rds_account.ResourceVolcengineRdsAccount(),
			"volcengine_rds_ip_list":            rds_ip_list.ResourceVolcengineRdsIpList(),
			"volcengine_rds_account_privilege":  rds_account_privilege.ResourceVolcengineRdsAccountPrivilege(),
			"volcengine_rds_parameter_template": rds_parameter_template.ResourceVolcengineRdsParameterTemplate(),

			// ================ RDS V2 ==============
			"volcengine_rds_instance_v2": rds_instance_v2.ResourceVolcengineRdsInstance(),

			// ================ ESCloud ================
			"volcengine_escloud_instance": instance.ResourceVolcengineESCloudInstance(),

			//================= TOS =================
			"volcengine_tos_bucket": bucket.ResourceVolcengineTosBucket(),
			"volcengine_tos_object": object.ResourceVolcengineTosObject(),

			// ================ CR ================
			"volcengine_cr_registry":       cr_registry.ResourceVolcengineCrRegistry(),
			"volcengine_cr_registry_state": cr_registry_state.ResourceVolcengineCrRegistryState(),
			"volcengine_cr_namespace":      cr_namespace.ResourceVolcengineCrNamespace(),
			"volcengine_cr_repository":     cr_repository.ResourceVolcengineCrRepository(),
			"volcengine_cr_tag":            cr_tag.ResourceVolcengineCrTag(),
			"volcengine_cr_endpoint":       cr_endpoint.ResourceVolcengineCrEndpoint(),

			// ================ Veenedge ================
			"volcengine_veenedge_cloud_server": cloud_server.ResourceVolcengineCloudServer(),
			"volcengine_veenedge_instance":     veInstance.ResourceVolcengineInstance(),
			"volcengine_veenedge_vpc":          veVpc.ResourceVolcengineVpc(),
		},
		ConfigureFunc: ProviderConfigure,
	}
}

func ProviderConfigure(d *schema.ResourceData) (interface{}, error) {
	config := ve.Config{
		AccessKey:       d.Get("access_key").(string),
		SecretKey:       d.Get("secret_key").(string),
		SessionToken:    d.Get("session_token").(string),
		Region:          d.Get("region").(string),
		Endpoint:        d.Get("endpoint").(string),
		DisableSSL:      d.Get("disable_ssl").(bool),
		CustomerHeaders: map[string]string{},
		ProxyUrl:        d.Get("proxy_url").(string),
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

	client, err := config.Client()
	return client, err
}
