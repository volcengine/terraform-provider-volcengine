package sweep

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

func SharedClientForRegionWithResourceId(region string) (interface{}, error) {
	var (
		accessKey  string
		secretKey  string
		endpoint   string
		disableSSL bool
	)

	if accessKey = os.Getenv("VOLCENGINE_ACCESS_KEY"); accessKey == "" {
		return nil, fmt.Errorf("%s can not be empty", "VOLCENGINE_ACCESS_KEY")
	}
	if secretKey = os.Getenv("VOLCENGINE_SECRET_KEY"); secretKey == "" {
		return nil, fmt.Errorf("%s can not be empty", "VOLCENGINE_SECRET_KEY")
	}
	if endpoint = os.Getenv("VOLCENGINE_ENDPOINT"); endpoint == "" {
		return nil, fmt.Errorf("%s can not be empty", "VOLCENGINE_ENDPOINT")
	}
	disableSSL, _ = strconv.ParseBool(os.Getenv("VOLCENGINE_DISABLE_SSL"))

	config := ve.Config{
		AccessKey:         accessKey,
		SecretKey:         secretKey,
		Region:            region,
		Endpoint:          endpoint,
		DisableSSL:        disableSSL,
		SessionToken:      os.Getenv("VOLCENGINE_SESSION_TOKEN"),
		ProxyUrl:          os.Getenv("VOLCENGINE_PROXY_URL"),
		CustomerHeaders:   map[string]string{},
		CustomerEndpoints: defaultCustomerEndPoints(),
	}

	headers := os.Getenv("VOLCENGINE_CUSTOMER_HEADERS")
	if headers != "" {
		hs1 := strings.Split(headers, ",")
		for _, hh := range hs1 {
			hs2 := strings.Split(hh, ":")
			if len(hs2) == 2 {
				config.CustomerHeaders[hs2[0]] = hs2[1]
			}
		}
	}

	endpoints := os.Getenv("VOLCENGINE_CUSTOMER_ENDPOINTS")
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
	if err != nil {
		return nil, err
	}

	return client, nil
}

func defaultCustomerEndPoints() map[string]string {
	return map[string]string{
		"veenedge": "veenedge.volcengineapi.com",
	}
}

type SweeperInstance interface {
	GetId() string
	Delete() error
}

func SweeperScheduler(sweepInstances []SweeperInstance) error {
	var (
		wg       sync.WaitGroup
		syncMap  sync.Map
		errorStr string
	)

	if len(sweepInstances) == 0 {
		return nil
	}
	wg.Add(len(sweepInstances))

	for _, value := range sweepInstances {
		sweepInstance := value
		go func() {
			defer func() {
				if err := recover(); err != nil {
					logger.DebugInfo(" Sweep Resource Panic, resource: ", sweepInstance.GetId(), "error: ", err)
				}
				wg.Done()
			}()

			err := sweepInstance.Delete()
			if err != nil {
				syncMap.Store(sweepInstance.GetId(), err)
			}
		}()
	}

	wg.Wait()
	for _, sweepInstance := range sweepInstances {
		if v, exist := syncMap.Load(sweepInstance.GetId()); exist {
			if err, ok := v.(error); ok {
				errorStr = errorStr + "Sweep Resource " + sweepInstance.GetId() + " error: " + err.Error() + ";\n"
			}
		}
	}
	if len(errorStr) > 0 {
		return fmt.Errorf(errorStr)
	}
	return nil
}
