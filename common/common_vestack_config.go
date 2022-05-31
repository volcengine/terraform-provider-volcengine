package common

import (
	"context"
	"fmt"
	"net/http"

	"github.com/volcengine/volcstack-go-sdk/service/autoscaling"
	"github.com/volcengine/volcstack-go-sdk/service/clb"
	"github.com/volcengine/volcstack-go-sdk/service/ecs"
	"github.com/volcengine/volcstack-go-sdk/service/natgateway"
	"github.com/volcengine/volcstack-go-sdk/service/rdsmysql"
	"github.com/volcengine/volcstack-go-sdk/service/rdsmysqlv2"
	"github.com/volcengine/volcstack-go-sdk/service/storageebs"
	"github.com/volcengine/volcstack-go-sdk/service/vke"
	"github.com/volcengine/volcstack-go-sdk/service/vpc"
	"github.com/volcengine/volcstack-go-sdk/service/vpn"
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/credentials"
	"github.com/volcengine/volcstack-go-sdk/volcstack/session"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

type Config struct {
	AccessKey       string
	SecretKey       string
	SessionToken    string
	Region          string
	Endpoint        string
	DisableSSL      bool
	CustomerHeaders map[string]string
}

func (c *Config) Client() (*SdkClient, error) {
	var client SdkClient
	version := fmt.Sprintf("%s/%s", TerraformProviderName, TerraformProviderVersion)

	config := volcstack.NewConfig().
		WithRegion(c.Region).
		WithExtraUserAgent(volcstack.String(version)).
		WithCredentials(credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, c.SessionToken)).
		WithDisableSSL(c.DisableSSL).
		WithExtendHttpRequest(func(ctx context.Context, request *http.Request) {
			if len(c.CustomerHeaders) > 0 {
				for k, v := range c.CustomerHeaders {
					request.Header.Add(k, v)
				}
			}
		}).
		WithEndpoint(volcstackutil.NewEndpoint().WithCustomerEndpoint(c.Endpoint).GetEndpoint())

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, fmt.Errorf("session init error %w", err)
	}

	client.Region = c.Region
	client.VpcClient = vpc.New(sess)
	client.ClbClient = clb.New(sess)
	client.EcsClient = ecs.New(sess)
	client.EbsClient = storageebs.New(sess)
	client.VpnClient = vpn.New(sess)
	client.NatClient = natgateway.New(sess)
	client.AutoScalingClient = autoscaling.New(sess)
	client.RdsClient = rdsmysql.New(sess)
	client.RdsClientV2 = rdsmysqlv2.New(sess)
	client.VkeClient = vke.New(sess)
	client.UniversalClient = NewUniversalClient(sess)

	InitLocks()
	InitSyncLimit()
	return &client, nil
}
