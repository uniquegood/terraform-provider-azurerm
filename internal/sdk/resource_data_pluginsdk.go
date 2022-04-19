//go:build !framework
// +build !framework

package sdk

import (
	"time"

	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
)

var _ ResourceData = &PluginSdkResourceData{}

type PluginSdkResourceData struct {
	resourceData *pluginsdk.ResourceData
}

func NewPluginSdkResourceData(d *pluginsdk.ResourceData) *PluginSdkResourceData {
	return &PluginSdkResourceData{
		resourceData: d,
	}
}

// Get returns a value from either the config/state depending on where this is called
// in Create and Update functions this will return from the config
// in Read, Exists and Import functions this will return from the state
// NOTE: this should not be called from Delete functions.
func (p *PluginSdkResourceData) Get(key string) interface{} {
	return p.resourceData.Get(key)
}

func (p *PluginSdkResourceData) GetChange(key string) (interface{}, interface{}) {
	return p.resourceData.GetChange(key)
}

func (p *PluginSdkResourceData) GetRawValue(key string) (interface{}, bool) {
	return p.resourceData.GetOkExists(key)
}

func (p *PluginSdkResourceData) GetValue(key string) (interface{}, bool) {
	return p.resourceData.GetOk(key)
}

func (p *PluginSdkResourceData) HasChange(key string) bool {
	return p.resourceData.HasChange(key)
}

func (p *PluginSdkResourceData) HasChanges(keys ...string) bool {
	return p.resourceData.HasChanges(keys...)
}

func (p *PluginSdkResourceData) Id() string {
	return p.resourceData.Id()
}

func (p *PluginSdkResourceData) IsNewResource() bool {
	return p.resourceData.IsNewResource()
}

func (p *PluginSdkResourceData) Set(key string, value interface{}) error {
	return p.resourceData.Set(key, value)
}

func (p *PluginSdkResourceData) SetConnInfo(input map[string]string) {
	p.resourceData.SetConnInfo(input)
}

func (p *PluginSdkResourceData) SetId(id string) {
	p.resourceData.SetId(id)
}

func (p PluginSdkResourceData) Timeout(key string) time.Duration {
	return p.resourceData.Timeout(key)
}

// TODO: remove below here - these are just for compatibility whilst we migrate across to the wrapper

func (p *PluginSdkResourceData) GetOk(key string) (interface{}, bool) {
	return p.resourceData.GetOk(key)
}

func (p *PluginSdkResourceData) GetOkExists(key string) (interface{}, bool) {
	return p.resourceData.GetOkExists(key)
}
