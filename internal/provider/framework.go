//go:build framework
// +build framework

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ tfsdk.Provider = Provider{}

func AzureProvider() tfsdk.Provider {
	return Provider{}
}

type Provider struct {
}

// GetSchema returns the schema for this provider's configuration. If
// this provider has no configuration, return an empty schema.Schema.
func (p Provider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"subscription_id": {
				Type:     types.StringType,
				Optional: true,
				// TODO: missing the ability to specify a default value
				// TODO: missing the ability to source values from an environment variable
				Description: "The Subscription ID which should be used.",
			},

			"client_id": {
				Type:     types.StringType,
				Optional: true,
				// TODO: missing the ability to specify a default value
				// TODO: missing the ability to source values from an environment variable
				Description: "The Client ID which should be used.",
			},

			"tenant_id": {
				Type:     types.StringType,
				Optional: true,
				// TODO: missing the ability to specify a default value
				// TODO: missing the ability to source values from an environment variable
				Description: "The Tenant ID which should be used.",
			},
			"environment": {
				Type:     types.StringType,
				Required: true,
				// TODO: missing the ability to specify a default value
				// TODO: missing the ability to source values from an environment variable
				Description: "The Cloud Environment which should be used. Possible values are public, usgovernment, and china. Defaults to public.",
			},

			"auxiliary_tenant_ids": {
				Type: types.ListType{
					ElemType: types.StringType,
				},
				Optional: true,
				//// TODO: can't do minimum items
				//MaxItems: 3,
				// perhaps this can be done via Validators but :shrug:
				//Validators: []tfsdk.AttributeValidator{
				//
				//},
			},

			"metadata_host": {
				Type:     types.StringType,
				Required: true,
				// TODO: missing the ability to specify a default value
				// TODO: missing the ability to source values from an environment variable
				Description: "The Hostname which should be used for the Azure Metadata Service.",
			},

			// Client Certificate specific fields
			"client_certificate_path": {
				Type:     types.StringType,
				Optional: true,
				// TODO: missing the ability to specify a default value
				// TODO: missing the ability to source values from an environment variable
				Description: "The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate.",
			},

			"client_certificate_password": {
				Type:     types.StringType,
				Optional: true,
				// TODO: missing the ability to specify a default value
				// TODO: missing the ability to source values from an environment variable
				Description: "The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate",
			},

			// Client Secret specific fields
			"client_secret": {
				Type:     types.StringType,
				Optional: true,
				// TODO: missing the ability to specify a default value
				// TODO: missing the ability to source values from an environment variable
				Description: "The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.",
			},

			// Managed Service Identity specific fields
			"use_msi": {
				Type:     types.BoolType,
				Optional: true,
				// TODO: missing the ability to specify a default value
				// TODO: missing the ability to source values from an environment variable
				Description: "Allowed Managed Service Identity be used for Authentication.",
			},
			"msi_endpoint": {
				Type:     types.StringType,
				Optional: true,
				// TODO: missing the ability to specify a default value
				// TODO: missing the ability to source values from an environment variable
				Description: "The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected automatically. ",
			},

			// Managed Tracking GUID for User-agent
			"partner_id": {
				Type:     types.StringType,
				Optional: true,
				// TODO: missing the ability to specify a default value
				// TODO: missing the ability to source values from an environment variable
				// TODO: missing a helper to do UUID validation
				Description: "A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.",
			},

			"disable_correlation_request_id": {
				Type:     types.BoolType,
				Optional: true,
				// TODO: missing the ability to specify a default value
				// TODO: missing the ability to source values from an environment variable
				Description: "This will disable the x-ms-correlation-request-id header.",
			},

			"disable_terraform_partner_id": {
				Type:     types.BoolType,
				Optional: true,
				// TODO: missing the ability to specify a default value
				// TODO: missing the ability to source values from an environment variable
				Description: "This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.",
			},

			// TODO: `features` is a block
			//"features": schemaFeatures(supportLegacyTestSuite),

			// Advanced feature flags
			"skip_provider_registration": {
				Type:     types.BoolType,
				Optional: true,
				// TODO: missing the ability to specify a default value
				// TODO: missing the ability to source values from an environment variable
				Description: "Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already registered?",
			},

			"storage_use_azuread": {
				Type:     types.BoolType,
				Optional: true,
				// TODO: missing the ability to specify a default value
				// TODO: missing the ability to source values from an environment variable
				Description: "Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?",
			},
		},
		Blocks:              nil,
		Version:             1,
		DeprecationMessage:  "",
		Description:         "",
		MarkdownDescription: "",
	}, nil
}

// Configure is called at the beginning of the provider lifecycle, when
// Terraform sends to the provider the values the user specified in the
// provider configuration block. These are supplied in the
// ConfigureProviderRequest argument.
// Values from provider configuration are often used to initialise an
// API client, which should be stored on the struct implementing the
// Provider interface.
func (p Provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	// TODO: implement me
}

// GetDataSources returns a map of the data source types this provider
// supports.
func (p Provider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	// TODO: implement me
	return nil, nil
}

// GetResources returns a map of the resource types this provider
// supports.
func (p Provider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	// TODO: implement me
	return nil, nil
}
