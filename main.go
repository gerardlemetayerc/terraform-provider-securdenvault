package main

import (
	"github.com/gerardlemetayerc/terraform-provider-securdenvault/vault"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: device42.Provider,
	})
}