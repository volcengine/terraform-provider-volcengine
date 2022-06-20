package common

import (
	"github.com/volcengine/volcstack-go-sdk/service/autoscaling"
	"github.com/volcengine/volcstack-go-sdk/service/clb"
	"github.com/volcengine/volcstack-go-sdk/service/ecs"
	"github.com/volcengine/volcstack-go-sdk/service/natgateway"
	"github.com/volcengine/volcstack-go-sdk/service/rdsmysql"
	"github.com/volcengine/volcstack-go-sdk/service/rdsmysqlv2"
	"github.com/volcengine/volcstack-go-sdk/service/storageebs"
	"github.com/volcengine/volcstack-go-sdk/service/vpc"
	"github.com/volcengine/volcstack-go-sdk/service/vpn"
)

type SdkClient struct {
	Region            string
	VpcClient         *vpc.VPC
	ClbClient         *clb.CLB
	EcsClient         *ecs.ECS
	EbsClient         *storageebs.STORAGEEBS
	NatClient         *natgateway.NATGATEWAY
	VpnClient         *vpn.VPN
	AutoScalingClient *autoscaling.AUTOSCALING
	RdsClient         *rdsmysql.RDSMYSQL
	RdsClientV2       *rdsmysqlv2.RDSMYSQLV2
	UniversalClient   *Universal
}
