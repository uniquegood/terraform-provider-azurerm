//go:build framework
// +build framework

package sdk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
)

type ResourceBuilderWrapper interface {
	// GetSchema returns the schema for this resource.
	GetSchema(context.Context) (tfsdk.Schema, diag.Diagnostics)

	// NewResource instantiates a new Resource of this ResourceType.
	NewResource() func(context.Context, *clients.Client) (tfsdk.Resource, diag.Diagnostics)
}

var _ ResourceBuilderWrapper = resourceBuilder{}

var _ tfsdk.Resource = resourceWrapper{}

type resourceBuilder struct {
	typedResource Resource
}

func NewResourceBuilder(typedResource Resource) resourceBuilder {
	return resourceBuilder{
		typedResource: typedResource,
	}
}

func (r resourceBuilder) NewResource() func(context.Context, *clients.Client) (tfsdk.Resource, diag.Diagnostics) {
	return func(ctx context.Context, client *clients.Client) (tfsdk.Resource, diag.Diagnostics) {
		return resourceWrapper{
			typedResource: r.typedResource,
			client:        client,
		}, nil
	}
}

func (r resourceBuilder) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	attributes := make(map[string]tfsdk.Attribute, 0)

	for k, v := range r.typedResource.Attributes() {
		attr, err := frameworkAttributeFromPluginSdkType(v)
		if err != nil {
			e := diag.Diagnostics{}
			e = append(e, diag.NewErrorDiagnostic("internal-error", err.Error()))
			return tfsdk.Schema{}, e
		}

		attributes[k] = *attr
	}
	for k, v := range r.typedResource.Arguments() {
		attr, err := frameworkAttributeFromPluginSdkType(v)
		if err != nil {
			e := diag.Diagnostics{}
			e = append(e, diag.NewErrorDiagnostic("internal-error", err.Error()))
			return tfsdk.Schema{}, e
		}

		attributes[k] = *attr
	}

	version := 0
	if v, ok := r.typedResource.(ResourceWithStateMigration); ok {
		// TODO support state migrations
		version = v.StateUpgraders().SchemaVersion
	}

	return tfsdk.Schema{
		Attributes: attributes,
		Version:    int64(version),
	}, nil
}

type resourceWrapper struct {
	typedResource Resource
	client        *clients.Client
}

func (r resourceWrapper) Create(ctx context.Context, request tfsdk.CreateResourceRequest, response *tfsdk.CreateResourceResponse) {
	f := r.typedResource.Create()

	resourceData := NewFrameworkResourceData(ctx, &response.State)
	resourceData.WithConfig(request.Config)
	resourceData.WithPlan(request.Plan)
	err := f.Func(ctx, ResourceMetaData{
		Client:                   r.client,
		Logger:                   NullLogger{},
		ResourceData:             resourceData,
		ResourceDiff:             nil,
		serializationDebugLogger: nil,
	})
	if err != nil {
		response.Diagnostics.AddError("performing create", err.Error())
		return
	}

	response.State = *resourceData.state
}

func (r resourceWrapper) Read(ctx context.Context, request tfsdk.ReadResourceRequest, response *tfsdk.ReadResourceResponse) {
	f := r.typedResource.Read()

	resourceData := NewFrameworkResourceData(ctx, &response.State)
	resourceData.WithExistingState(request.State)
	err := f.Func(ctx, ResourceMetaData{
		Client:                   r.client,
		Logger:                   NullLogger{},
		ResourceData:             resourceData,
		ResourceDiff:             nil,
		serializationDebugLogger: nil,
	})
	if err != nil {
		response.Diagnostics.AddError("performing read", err.Error())
		return
	}

	response.State = *resourceData.state
}

func (r resourceWrapper) Update(ctx context.Context, request tfsdk.UpdateResourceRequest, response *tfsdk.UpdateResourceResponse) {
	rwu, ok := r.typedResource.(ResourceWithUpdate)
	if !ok {
		response.Diagnostics.AddError("doesn't support update", "this resource does not support update")
		return
	}

	f := rwu.Update()

	resourceData := NewFrameworkResourceData(ctx, &response.State)
	resourceData.WithConfig(request.Config)
	resourceData.WithExistingState(request.State)
	resourceData.WithPlan(request.Plan)
	err := f.Func(ctx, ResourceMetaData{
		Client:                   r.client,
		Logger:                   NullLogger{},
		ResourceData:             resourceData,
		ResourceDiff:             nil,
		serializationDebugLogger: nil,
	})
	if err != nil {
		response.Diagnostics.AddError("performing update", err.Error())
		return
	}

	response.State = *resourceData.state
}

func (r resourceWrapper) Delete(ctx context.Context, request tfsdk.DeleteResourceRequest, response *tfsdk.DeleteResourceResponse) {
	f := r.typedResource.Delete()

	// TODO: note the request doesn't define the config for the schema here, but it should
	resourceData := NewFrameworkResourceData(ctx, &response.State)
	resourceData.WithExistingState(request.State)
	err := f.Func(ctx, ResourceMetaData{
		Client:                   r.client,
		Logger:                   NullLogger{},
		ResourceData:             resourceData,
		ResourceDiff:             nil,
		serializationDebugLogger: nil,
	})
	if err != nil {
		response.Diagnostics.AddError("performing delete", err.Error())
		return
	}

	response.State = *resourceData.state
}

func (r resourceWrapper) ImportState(ctx context.Context, request tfsdk.ImportResourceStateRequest, response *tfsdk.ImportResourceStateResponse) {
	rwi, ok := r.typedResource.(ResourceWithCustomImporter)
	if !ok {
		// TODO: should this be ResourceImportStateNotImplemented?
		response.Diagnostics.AddError("doesn't support import", "this resource doesn't support import")
		return
	}

	f := rwi.CustomImporter()

	resourceData := NewFrameworkResourceData(ctx, &response.State)
	resourceData.WithExistingID(request.ID)
	err := f(ctx, ResourceMetaData{
		Client:                   r.client,
		Logger:                   NullLogger{},
		ResourceData:             resourceData,
		ResourceDiff:             nil,
		serializationDebugLogger: nil,
	})
	if err != nil {
		response.Diagnostics.AddError("performing delete", err.Error())
		return
	}

	response.State = *resourceData.state
}
