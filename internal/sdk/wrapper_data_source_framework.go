//go:build framework
// +build framework

package sdk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
)

type DataSourceBuilderWrapper interface {
	// GetSchema returns the schema for this resource.
	GetSchema(context.Context) (tfsdk.Schema, diag.Diagnostics)

	// NewDataSource instantiates a new DataSource of this DataSourceType.
	NewDataSource() func(context.Context, *clients.Client) (tfsdk.DataSource, diag.Diagnostics)
}

var _ DataSourceBuilderWrapper = dataSourceBuilder{}

type dataSourceBuilder struct {
	typedDataSource DataSource
}

func NewDataSourceBuilder(typedDataSource DataSource) dataSourceBuilder {
	return dataSourceBuilder{
		typedDataSource: typedDataSource,
	}
}

func (d dataSourceBuilder) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	attributes := make(map[string]tfsdk.Attribute, 0)

	for k, v := range d.typedDataSource.Attributes() {
		attr, err := frameworkAttributeFromPluginSdkType(v)
		if err != nil {
			e := diag.Diagnostics{}
			e = append(e, diag.NewErrorDiagnostic("internal-error", err.Error()))
			return tfsdk.Schema{}, e
		}

		attributes[k] = *attr
	}
	for k, v := range d.typedDataSource.Arguments() {
		attr, err := frameworkAttributeFromPluginSdkType(v)
		if err != nil {
			e := diag.Diagnostics{}
			e = append(e, diag.NewErrorDiagnostic("internal-error", err.Error()))
			return tfsdk.Schema{}, e
		}

		attributes[k] = *attr
	}

	return tfsdk.Schema{
		Attributes: attributes,
	}, nil
}

func (d dataSourceBuilder) NewDataSource() func(context.Context, *clients.Client) (tfsdk.DataSource, diag.Diagnostics) {
	return func(ctx context.Context, client *clients.Client) (tfsdk.DataSource, diag.Diagnostics) {
		return dataSourceWrapper{
			client:          client,
			typedDataSource: d.typedDataSource,
		}, nil
	}
}

var _ tfsdk.DataSource = dataSourceWrapper{}

type dataSourceWrapper struct {
	client          *clients.Client
	typedDataSource DataSource
}

func (d dataSourceWrapper) Read(ctx context.Context, request tfsdk.ReadDataSourceRequest, response *tfsdk.ReadDataSourceResponse) {
	f := d.typedDataSource.Read()

	resourceData := NewFrameworkResourceData(ctx, request.Config.Schema, &response.State)
	err := f.Func(ctx, ResourceMetaData{
		Client:                   d.client,
		Logger:                   NullLogger{},
		ResourceData:             resourceData,
		ResourceDiff:             nil,
		serializationDebugLogger: nil,
	})
	if err != nil {
		response.Diagnostics.AddError("peforming read", err.Error())
		return
	}

	response.State = *resourceData.state
}
