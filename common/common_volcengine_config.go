package common

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/volcengine/volcengine-go-sdk/service/autoscaling"
	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/service/iam"
	"github.com/volcengine/volcengine-go-sdk/service/natgateway"
	"github.com/volcengine/volcengine-go-sdk/service/rdsmysql"
	"github.com/volcengine/volcengine-go-sdk/service/rdsmysqlv2"
	"github.com/volcengine/volcengine-go-sdk/service/storageebs"
	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/service/vpn"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

type Config struct {
	AccessKey              string
	SecretKey              string
	SessionToken           string
	Region                 string
	Endpoint               string
	DisableSSL             bool
	EnableStandardEndpoint bool
	CustomerHeaders        map[string]string
	CustomerEndpoints      map[string]string
	CustomerEndpointSuffix map[string]string
	ProxyUrl               string
}

func (c *Config) Client() (*SdkClient, error) {
	var client SdkClient
	version := fmt.Sprintf("%s/%s", TerraformProviderName, TerraformProviderVersion)

	config := volcengine.NewConfig().
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
	}

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
	client.IamClient = iam.New(sess)
	client.UniversalClient = NewUniversalClient(sess, c.CustomerEndpoints, c.EnableStandardEndpoint)
	client.BypassSvcClient = NewBypassClient(sess, c.CustomerEndpointSuffix)

	//InitLocks()
	//InitSyncLimit()
	return &client, nil
}

func init() {
	InitLocks()
	//InitSyncLimit()
}
