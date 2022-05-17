package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/volcengine/terraform-provider-vestack/vestack"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vestack.Provider,
	})
}
