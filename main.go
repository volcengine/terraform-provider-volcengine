package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: volcengine.Provider,
	})
}
