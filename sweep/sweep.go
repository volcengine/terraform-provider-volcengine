package sweep

import (
	"fmt"
	"os"
	"sync"

	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

func SharedClientForRegionWithResourceId(region string) (interface{}, error) {
	var accessKey string
	if accessKey = os.Getenv("VOLCENGINE_ACCESS_KEY"); accessKey == "" {
		return nil, fmt.Errorf("%s can not be empty", "VOLCENGINE_ACCESS_KEY")
	}

	var secretKey string
	if secretKey = os.Getenv("VOLCENGINE_SECRET_KEY"); secretKey == "" {
		return nil, fmt.Errorf("%s can not be empty", "VOLCENGINE_SECRET_KEY")
	}

	var endpoint string
	if endpoint = os.Getenv("VOLCENGINE_ENDPOINT"); endpoint == "" {
		return nil, fmt.Errorf("%s can not be empty", "VOLCENGINE_ENDPOINT")
	}

	config := ve.Config{
		AccessKey: accessKey,
		SecretKey: secretKey,
		Region:    region,
		Endpoint:  endpoint,
	}

	client, err := config.Client()
	if err != nil {
		return nil, err
	}

	return client, nil
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
