package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/knockoutez/terraform-provider-crudcrud/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}