//go:build sweep
// +build sweep

package vpc

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/sweep"
)

func init() {
	resource.AddTestSweepers("volcengine_vpc", &resource.Sweeper{
		Name: "volcengine_vpc",
		F:    testSweepVolcengineVpcResource,
	})
}

func testSweepVolcengineVpcResource(region string) error {
	var (
		results   []interface{}
		err       error
		skipSweep bool
	)

	prefixes := []string{
		"acc-test",
	}

	sharedClient, err := sweep.SharedClientForRegionWithResourceId(region)
	if err != nil {
		return fmt.Errorf("getting volcengine client error: %s", err.Error())
	}

	client := sharedClient.(*ve.SdkClient)
	service := NewVpcService(client)

	sweepResources := make([]sweep.SweeperInstance, 0)
	results, err = service.ReadResources(map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("vpc ReadResources error: %s", err.Error())
	}

	for _, value := range results {
		resource, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("result is not map")
		}
		resourceId := resource["VpcId"]
		resourceName := resource["VpcName"]
		skipSweep = true
		for _, prefix := range prefixes {
			if strings.HasPrefix(resourceName.(string), prefix) {
				skipSweep = false
				break
			}
		}
		if skipSweep {
			logger.DebugInfo(" Skip sweep vpc: %s (%s)", resourceId, resourceName)
			continue
		}

		r := ResourceVolcengineVpc()
		d := r.Data(nil)
		d.SetId(resourceId.(string))

		sweepResources = append(sweepResources, sweep.NewSweepResource(r, d, client))
	}

	err = sweep.SweeperScheduler(sweepResources)
	if err != nil {
		return fmt.Errorf(" Sweep vpc Error, region: %s, err: %s", region, err)
	}

	return nil
}
