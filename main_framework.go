//go:build framework
// +build framework

package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/provider"
)

func launchProvider(debugMode bool) {
	opts := tfsdk.ServeOpts{
		Name:  "registry.terraform.io/hashicorp/azurerm",
		Debug: debugMode,
	}

	err := tfsdk.Serve(context.Background(), provider.AzureProvider, opts)
	if err != nil {
		log.Println(err.Error())
	}
}
