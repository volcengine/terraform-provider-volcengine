package volcengine

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/volcengine/terraform-provider-volcengine/volcengine/vepfs/vepfs_file_system"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vepfs/vepfs_fileset"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vepfs/vepfs_mount_service"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vepfs/vepfs_mount_service_attachment"

	"github.com/volcengine/terraform-provider-volcengine/volcengine/private_zone/private_zone"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/private_zone/private_zone_record"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/private_zone/private_zone_record_set"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/private_zone/private_zone_record_weight_enabler"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/private_zone/private_zone_resolver_endpoint"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/private_zone/private_zone_resolver_rule"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/private_zone/private_zone_user_vpc_authorization"

	"github.com/volcengine/terraform-provider-volcengine/volcengine/kafka/kafka_consumed_partition"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/kafka/kafka_consumed_topic"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/kafka/kafka_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/kafka/kafka_instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/kafka/kafka_public_address"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/kafka/kafka_region"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/kafka/kafka_sasl_user"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/kafka/kafka_topic"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/kafka/kafka_topic_partition"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/kafka/kafka_zone"

	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloud_identity/cloud_identity_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloud_identity/cloud_identity_permission_set"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloud_identity/cloud_identity_permission_set_assignment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloud_identity/cloud_identity_permission_set_provisioning"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloud_identity/cloud_identity_user"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloud_identity/cloud_identity_user_attachment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloud_identity/cloud_identity_user_provisioning"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/financial_relation/financial_relation"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/organization/organization"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/organization/organization_account"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/organization/organization_service_control_policy"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/organization/organization_service_control_policy_attachment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/organization/organization_service_control_policy_enabler"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/organization/organization_unit"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"

	"github.com/volcengine/terraform-provider-volcengine/volcengine/cdn/cdn_certificate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cdn/cdn_config"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cdn/cdn_domain"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cdn/cdn_shared_config"

	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_acl"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_certificate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_server_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_server_group_server"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_zone"

	"github.com/volcengine/terraform-provider-volcengine/volcengine/cen/cen_service_route_entry"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloud_monitor/cloud_monitor_contact"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloud_monitor/cloud_monitor_contact_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloud_monitor/cloud_monitor_event_rule"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cloud_monitor/cloud_monitor_rule"
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
	mssqlBackup "github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mssql/rds_mssql_backup"
	mssqlInstance "github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mssql/rds_mssql_instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mssql/rds_mssql_region"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mssql/rds_mssql_zone"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_account"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_allowlist"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_allowlist_associate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_database"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_instance_readonly_node"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_schema"
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
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ebs/ebs_auto_snapshot_policy"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ebs/ebs_auto_snapshot_policy_attachment"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ebs/ebs_snapshot"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ebs/ebs_snapshot_group"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ebs/volume"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ebs/volume_attach"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_available_resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_command"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_deployment_set"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_deployment_set_associate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_hpc_cluster"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_instance"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_instance_state"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_instance_type"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_invocation"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_invocation_result"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_key_pair"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_key_pair_associate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_launch_template"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/image"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/image_import"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/image_share_permission"
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
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_saml_provider"
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
			"assume_role": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The ASSUME ROLE block for Volcengine Provider. If provided, terraform will attempt to assume this role using the supplied credentials.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assume_role_trn": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("VOLCENGINE_ASSUME_ROLE_TRN", nil),
							Description: "The TRN of the role to assume.",
						},
						"assume_role_session_name": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("VOLCENGINE_ASSUME_ROLE_SESSION_NAME", nil),
							Description: "The session name to use when making the AssumeRole call.",
						},
						"duration_seconds": {
							Type:     schema.TypeInt,
							Required: true,
							DefaultFunc: func() (interface{}, error) {
								if v := os.Getenv("VOLCENGINE_ASSUME_ROLE_DURATION_SECONDS"); v != "" {
									return strconv.Atoi(v)
								}
								return 3600, nil
							},
							ValidateFunc: validation.IntBetween(900, 43200),
							Description:  "The duration of the session when making the AssumeRole call. Its value ranges from 900 to 43200(seconds), and default is 3600 seconds.",
						},
						"policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A more restrictive policy when making the AssumeRole call.",
						},
					},
				},
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
			"volcengine_volumes":                    volume.DataSourceVolcengineVolumes(),
			"volcengine_ebs_snapshots":              ebs_snapshot.DataSourceVolcengineEbsSnapshots(),
			"volcengine_ebs_snapshot_groups":        ebs_snapshot_group.DataSourceVolcengineEbsSnapshotGroups(),
			"volcengine_ebs_auto_snapshot_policies": ebs_auto_snapshot_policy.DataSourceVolcengineEbsAutoSnapshotPolicies(),

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
			"volcengine_image_share_permissions": image_share_permission.DataSourceVolcengineImageSharePermissions(),
			"volcengine_ecs_hpc_clusters":        ecs_hpc_cluster.DataSourceVolcengineEcsHpcClusters(),

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
			"volcengine_iam_saml_providers":                iam_saml_provider.DataSourceVolcengineIamSamlProviders(),
			"volcengine_iam_access_keys":                   iam_access_key.DataSourceVolcengineIamAccessKeys(),

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
			"volcengine_cloud_monitor_event_rules":    cloud_monitor_event_rule.DataSourceVolcengineCloudMonitorEventRules(),
			"volcengine_cloud_monitor_rules":          cloud_monitor_rule.DataSourceVolcengineCloudMonitorRules(),

			// ================ RdsMssql ================
			"volcengine_rds_mssql_regions":   rds_mssql_region.DataSourceVolcengineRdsMssqlRegions(),
			"volcengine_rds_mssql_zones":     rds_mssql_zone.DataSourceVolcengineRdsMssqlZones(),
			"volcengine_rds_mssql_instances": mssqlInstance.DataSourceVolcengineRdsMssqlInstances(),
			"volcengine_rds_mssql_backups":   mssqlBackup.DataSourceVolcengineRdsMssqlBackups(),

			// ================ Postgresql ================
			"volcengine_rds_postgresql_databases":  rds_postgresql_database.DataSourceVolcengineRdsPostgresqlDatabases(),
			"volcengine_rds_postgresql_accounts":   rds_postgresql_account.DataSourceVolcengineRdsPostgresqlAccounts(),
			"volcengine_rds_postgresql_instances":  rds_postgresql_instance.DataSourceVolcengineRdsPostgresqlInstances(),
			"volcengine_rds_postgresql_allowlists": rds_postgresql_allowlist.DataSourceVolcengineRdsPostgresqlAllowlists(),

			// ================ Organization ================
			"volcengine_organization_units":                    organization_unit.DataSourceVolcengineOrganizationUnits(),
			"volcengine_organization_service_control_policies": organization_service_control_policy.DataSourceVolcengineServiceControlPolicies(),
			"volcengine_organization_accounts":                 organization_account.DataSourceVolcengineOrganizationAccounts(),
			"volcengine_organizations":                         organization.DataSourceVolcengineOrganizations(),
			"volcengine_rds_postgresql_schemas":                rds_postgresql_schema.DataSourceVolcengineRdsPostgresqlSchemas(),

			// ================ CDN ================
			"volcengine_cdn_certificates":   cdn_certificate.DataSourceVolcengineCdnCertificates(),
			"volcengine_cdn_domains":        cdn_domain.DataSourceVolcengineCdnDomains(),
			"volcengine_cdn_configs":        cdn_config.DataSourceVolcengineCdnConfigs(),
			"volcengine_cdn_shared_configs": cdn_shared_config.DataSourceVolcengineCdnSharedConfigs(),

			// ================ Financial Relation ================
			"volcengine_financial_relations": financial_relation.DataSourceVolcengineFinancialRelations(),

			// ================ Cloud Identity ================
			"volcengine_cloud_identity_users":                        cloud_identity_user.DataSourceVolcengineCloudIdentityUsers(),
			"volcengine_cloud_identity_groups":                       cloud_identity_group.DataSourceVolcengineCloudIdentityGroups(),
			"volcengine_cloud_identity_user_provisionings":           cloud_identity_user_provisioning.DataSourceVolcengineCloudIdentityUserProvisionings(),
			"volcengine_cloud_identity_permission_sets":              cloud_identity_permission_set.DataSourceVolcengineCloudIdentityPermissionSets(),
			"volcengine_cloud_identity_permission_set_assignments":   cloud_identity_permission_set_assignment.DataSourceVolcengineCloudIdentityPermissionSetAssignments(),
			"volcengine_cloud_identity_permission_set_provisionings": cloud_identity_permission_set_provisioning.DataSourceVolcengineCloudIdentityPermissionSetProvisionings(),

			// ================ Kafka ================
			"volcengine_kafka_sasl_users":          kafka_sasl_user.DataSourceVolcengineKafkaSaslUsers(),
			"volcengine_kafka_topic_partitions":    kafka_topic_partition.DataSourceVolcengineKafkaTopicPartitions(),
			"volcengine_kafka_groups":              kafka_group.DataSourceVolcengineKafkaGroups(),
			"volcengine_kafka_topics":              kafka_topic.DataSourceVolcengineKafkaTopics(),
			"volcengine_kafka_instances":           kafka_instance.DataSourceVolcengineKafkaInstances(),
			"volcengine_kafka_regions":             kafka_region.DataSourceVolcengineRegions(),
			"volcengine_kafka_zones":               kafka_zone.DataSourceVolcengineZones(),
			"volcengine_kafka_consumed_topics":     kafka_consumed_topic.DataSourceVolcengineKafkaConsumedTopics(),
			"volcengine_kafka_consumed_partitions": kafka_consumed_partition.DataSourceVolcengineKafkaConsumedPartitions(),

			// ================ PrivateZone ================
			"volcengine_private_zones":                   private_zone.DataSourceVolcenginePrivateZones(),
			"volcengine_private_zone_resolver_endpoints": private_zone_resolver_endpoint.DataSourceVolcenginePrivateZoneResolverEndpoints(),
			"volcengine_private_zone_resolver_rules":     private_zone_resolver_rule.DataSourceVolcenginePrivateZoneResolverRules(),
			"volcengine_private_zone_records":            private_zone_record.DataSourceVolcenginePrivateZoneRecords(),
			"volcengine_private_zone_record_sets":        private_zone_record_set.DataSourceVolcenginePrivateZoneRecordSets(),

			// ================ Vepfs ================
			"volcengine_vepfs_file_systems":   vepfs_file_system.DataSourceVolcengineVepfsFileSystems(),
			"volcengine_vepfs_mount_services": vepfs_mount_service.DataSourceVolcengineVepfsMountServices(),
			"volcengine_vepfs_filesets":       vepfs_fileset.DataSourceVolcengineVepfsFilesets(),
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
			"volcengine_volume":                              volume.ResourceVolcengineVolume(),
			"volcengine_volume_attach":                       volume_attach.ResourceVolcengineVolumeAttach(),
			"volcengine_ebs_snapshot":                        ebs_snapshot.ResourceVolcengineEbsSnapshot(),
			"volcengine_ebs_snapshot_group":                  ebs_snapshot_group.ResourceVolcengineEbsSnapshotGroup(),
			"volcengine_ebs_auto_snapshot_policy":            ebs_auto_snapshot_policy.ResourceVolcengineEbsAutoSnapshotPolicy(),
			"volcengine_ebs_auto_snapshot_policy_attachment": ebs_auto_snapshot_policy_attachment.ResourceVolcengineEbsAutoSnapshotPolicyAttachment(),

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
			"volcengine_image":                        image.ResourceVolcengineImage(),
			"volcengine_image_import":                 image_import.ResourceVolcengineImageImport(),
			"volcengine_image_share_permission":       image_share_permission.ResourceVolcengineImageSharePermission(),
			"volcengine_ecs_hpc_cluster":              ecs_hpc_cluster.ResourceVolcengineEcsHpcCluster(),

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
			"volcengine_iam_saml_provider":                iam_saml_provider.ResourceVolcengineIamSamlProvider(),

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
			"volcengine_cloud_monitor_event_rule":    cloud_monitor_event_rule.ResourceVolcengineCloudMonitorEventRule(),
			"volcengine_cloud_monitor_rule":          cloud_monitor_rule.ResourceVolcengineCloudMonitorRule(),

			// ================ RdsMssql ================
			"volcengine_rds_mssql_instance": mssqlInstance.ResourceVolcengineRdsMssqlInstance(),
			"volcengine_rds_mssql_backup":   mssqlBackup.ResourceVolcengineRdsMssqlBackup(),

			// ================ Postgresql ================
			"volcengine_rds_postgresql_database":               rds_postgresql_database.ResourceVolcengineRdsPostgresqlDatabase(),
			"volcengine_rds_postgresql_account":                rds_postgresql_account.ResourceVolcengineRdsPostgresqlAccount(),
			"volcengine_rds_postgresql_instance":               rds_postgresql_instance.ResourceVolcengineRdsPostgresqlInstance(),
			"volcengine_rds_postgresql_instance_readonly_node": rds_postgresql_instance_readonly_node.ResourceVolcengineRdsPostgresqlInstanceReadonlyNode(),
			"volcengine_rds_postgresql_allowlist":              rds_postgresql_allowlist.ResourceVolcengineRdsPostgresqlAllowlist(),
			"volcengine_rds_postgresql_allowlist_associate":    rds_postgresql_allowlist_associate.ResourceVolcengineRdsPostgresqlAllowlistAssociate(),

			// ================ Organization ================
			"volcengine_organization_unit":                              organization_unit.ResourceVolcengineOrganizationUnit(),
			"volcengine_organization_service_control_policy_enabler":    organization_service_control_policy_enabler.ResourceVolcengineOrganizationServiceControlPolicyEnabler(),
			"volcengine_organization_service_control_policy":            organization_service_control_policy.ResourceVolcengineServiceControlPolicy(),
			"volcengine_organization_service_control_policy_attachment": organization_service_control_policy_attachment.ResourceVolcengineServiceControlPolicyAttachment(),
			"volcengine_organization_account":                           organization_account.ResourceVolcengineOrganizationAccount(),
			"volcengine_organization":                                   organization.ResourceVolcengineOrganization(),
			"volcengine_rds_postgresql_schema":                          rds_postgresql_schema.ResourceVolcengineRdsPostgresqlSchema(),

			// ================ CDN ================
			"volcengine_cdn_certificate":   cdn_certificate.ResourceVolcengineCdnCertificate(),
			"volcengine_cdn_domain":        cdn_domain.ResourceVolcengineCdnDomain(),
			"volcengine_cdn_shared_config": cdn_shared_config.ResourceVolcengineCdnSharedConfig(),

			// ================ Financial Relation ================
			"volcengine_financial_relation": financial_relation.ResourceVolcengineFinancialRelation(),

			// ================ Cloud Identity ================
			"volcengine_cloud_identity_user":                        cloud_identity_user.ResourceVolcengineCloudIdentityUser(),
			"volcengine_cloud_identity_group":                       cloud_identity_group.ResourceVolcengineCloudIdentityGroup(),
			"volcengine_cloud_identity_user_attachment":             cloud_identity_user_attachment.ResourceVolcengineCloudIdentityUserAttachment(),
			"volcengine_cloud_identity_user_provisioning":           cloud_identity_user_provisioning.ResourceVolcengineCloudIdentityUserProvisioning(),
			"volcengine_cloud_identity_permission_set":              cloud_identity_permission_set.ResourceVolcengineCloudIdentityPermissionSet(),
			"volcengine_cloud_identity_permission_set_assignment":   cloud_identity_permission_set_assignment.ResourceVolcengineCloudIdentityPermissionSetAssignment(),
			"volcengine_cloud_identity_permission_set_provisioning": cloud_identity_permission_set_provisioning.ResourceVolcengineCloudIdentityPermissionSetProvisioning(),

			// ================ Kafka ================
			"volcengine_kafka_sasl_user":      kafka_sasl_user.ResourceVolcengineKafkaSaslUser(),
			"volcengine_kafka_group":          kafka_group.ResourceVolcengineKafkaGroup(),
			"volcengine_kafka_topic":          kafka_topic.ResourceVolcengineKafkaTopic(),
			"volcengine_kafka_instance":       kafka_instance.ResourceVolcengineKafkaInstance(),
			"volcengine_kafka_public_address": kafka_public_address.ResourceVolcengineKafkaPublicAddress(),

			// ================ PrivateZone ================
			"volcengine_private_zone":                        private_zone.ResourceVolcenginePrivateZone(),
			"volcengine_private_zone_resolver_endpoint":      private_zone_resolver_endpoint.ResourceVolcenginePrivateZoneResolverEndpoint(),
			"volcengine_private_zone_resolver_rule":          private_zone_resolver_rule.ResourceVolcenginePrivateZoneResolverRule(),
			"volcengine_private_zone_record":                 private_zone_record.ResourceVolcenginePrivateZoneRecord(),
			"volcengine_private_zone_record_weight_enabler":  private_zone_record_weight_enabler.ResourceVolcenginePrivateZoneRecordWeightEnabler(),
			"volcengine_private_zone_user_vpc_authorization": private_zone_user_vpc_authorization.ResourceVolcenginePrivateZoneUserVpcAuthorization(),

			// ================ Vepfs ================
			"volcengine_vepfs_file_system":              vepfs_file_system.ResourceVolcengineVepfsFileSystem(),
			"volcengine_vepfs_mount_service":            vepfs_mount_service.ResourceVolcengineVepfsMountService(),
			"volcengine_vepfs_mount_service_attachment": vepfs_mount_service_attachment.ResourceVolcengineVepfsMountServiceAttachment(),
			"volcengine_vepfs_fileset":                  vepfs_fileset.ResourceVolcengineVepfsFileset(),
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

	// get assume role
	var (
		arTrn             string
		arSessionName     string
		arPolicy          string
		arDurationSeconds int
	)

	if v, ok := d.GetOk("assume_role"); ok {
		assumeRoleList, ok := v.([]interface{})
		if !ok {
			return nil, fmt.Errorf("the assume_role is not slice ")
		}
		if len(assumeRoleList) == 1 {
			assumeRoleMap, ok := assumeRoleList[0].(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("the value of the assume_role is not map ")
			}
			arTrn = assumeRoleMap["assume_role_trn"].(string)
			arSessionName = assumeRoleMap["assume_role_session_name"].(string)
			arDurationSeconds = assumeRoleMap["duration_seconds"].(int)
			arPolicy = assumeRoleMap["policy"].(string)
		}
	} else {
		arTrn = os.Getenv("VOLCENGINE_ASSUME_ROLE_TRN")
		arSessionName = os.Getenv("VOLCENGINE_ASSUME_ROLE_SESSION_NAME")
		duration := os.Getenv("VOLCENGINE_ASSUME_ROLE_DURATION_SECONDS")
		if duration != "" {
			durationSeconds, err := strconv.Atoi(duration)
			if err != nil {
				return nil, err
			}
			arDurationSeconds = durationSeconds
		} else {
			arDurationSeconds = 3600
		}
	}

	if arTrn != "" && arSessionName != "" {
		cred, err := assumeRole(config, arTrn, arSessionName, arPolicy, arDurationSeconds)
		if err != nil {
			return nil, err
		}
		config.AccessKey = cred["AccessKeyId"].(string)
		config.SecretKey = cred["SecretAccessKey"].(string)
		config.SessionToken = cred["SessionToken"].(string)
	}

	client, err := config.Client()
	return client, err
}

func assumeRole(c ve.Config, arTrn, arSessionName, arPolicy string, arDurationSeconds int) (map[string]interface{}, error) {
	version := fmt.Sprintf("%s/%s", ve.TerraformProviderName, ve.TerraformProviderVersion)
	conf := volcengine.NewConfig().
		WithRegion(c.Region).
		WithExtraUserAgent(volcengine.String(version)).
		WithCredentials(credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, c.SessionToken)).
		WithDisableSSL(c.DisableSSL).
		WithExtendHttpRequest(func(ctx context.Context, request *http.Request) {
			if len(c.CustomerHeaders) > 0 {
				for k, v := range c.CustomerHeaders {
					request.Header.Add(k, v)
				}
			}
		}).
		WithEndpoint(volcengineutil.NewEndpoint().WithCustomerEndpoint(c.Endpoint).GetEndpoint())

	if c.ProxyUrl != "" {
		u, _ := url.Parse(c.ProxyUrl)
		t := &http.Transport{
			Proxy: http.ProxyURL(u),
		}
		httpClient := http.DefaultClient
		httpClient.Transport = t
		httpClient.Timeout = time.Duration(30000) * time.Millisecond
	}

	sess, err := session.NewSession(conf)
	if err != nil {
		return nil, err
	}

	universalClient := ve.NewUniversalClient(sess, c.CustomerEndpoints)

	action := "AssumeRole"
	req := map[string]interface{}{
		"RoleTrn":         arTrn,
		"RoleSessionName": arSessionName,
		"DurationSeconds": arDurationSeconds,
		"Policy":          arPolicy,
	}
	resp, err := universalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return nil, fmt.Errorf("AssumeRole failed, error: %s", err.Error())
	}
	results, err := ve.ObtainSdkValue("Result.Credentials", *resp)
	if err != nil {
		return nil, err
	}
	cred, ok := results.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("AssumeRole Result.Credentials is not Map")
	}
	return cred, nil
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "sts",
		Version:     "2018-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

func defaultCustomerEndPoints() map[string]string {
	return map[string]string{
		"veenedge": "veenedge.volcengineapi.com",
	}
}
