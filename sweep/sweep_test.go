package sweep_test

import (
	"testing"

	_ "github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/vpc"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}
