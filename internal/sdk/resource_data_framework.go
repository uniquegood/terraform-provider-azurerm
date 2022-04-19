//go:build framework
// +build framework

package sdk

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ ResourceData = &FrameworkResourceData{}

type FrameworkResourceData struct {
	ctx    context.Context
	schema tfsdk.Schema
	state  *tfsdk.State
}

func NewFrameworkResourceData(ctx context.Context, schema tfsdk.Schema, state *tfsdk.State) *FrameworkResourceData {
	return &FrameworkResourceData{
		ctx:    ctx,
		state:  state,
		schema: schema,
	}
}

// resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("blah"), "somevalue")
// data.Id = types.String{Value: "example-id"}

// resp.State.RemoveResource(ctx)

func (f *FrameworkResourceData) Get(key string) interface{} {
	var out interface{}
	f.state.GetAttribute(f.ctx, tftypes.NewAttributePath().WithAttributeName(key), out)
	return out
}

func (f *FrameworkResourceData) GetChange(key string) (original interface{}, updated interface{}) {
	//TODO implement me
	panic("implement me")
}

func (f *FrameworkResourceData) GetValue(key string) (value interface{}, isSet bool) {
	//TODO implement me
	panic("implement me")
}

func (f *FrameworkResourceData) GetRawValue(key string) (value interface{}, isSet bool) {
	//TODO implement me
	panic("implement me")
}

func (f *FrameworkResourceData) HasChange(key string) bool {
	//TODO implement me
	panic("implement me")
}

func (f *FrameworkResourceData) HasChanges(keys ...string) bool {
	//TODO implement me
	panic("implement me")
}

func (f *FrameworkResourceData) Id() string {
	//TODO implement me
	panic("implement me")
}

func (f *FrameworkResourceData) IsNewResource() bool {
	// TODO: implement me
	return false
}

func (f *FrameworkResourceData) Set(key string, value interface{}) error {
	d := f.state.SetAttribute(f.ctx, tftypes.NewAttributePath().WithAttributeName(key), value)
	if d.HasError() {
		// TODO: until Error() is implemented
		s := make([]string, 0)
		for _, e := range d {
			s = append(s, fmt.Sprintf("%s: %s", e.Summary(), e.Detail()))
		}

		return fmt.Errorf("setting attribute %q:\n\n%s", strings.Join(s, "\n\n"))
	}
	return nil
}

func (f *FrameworkResourceData) SetConnInfo(v map[string]string) {
	//TODO implement me
	panic("implement me")
}

func (f *FrameworkResourceData) SetId(id string) {
	if id == "" {
		f.state.RemoveResource(context.TODO())
	} else {
		f.Set("id", id)
	}
}

func (f *FrameworkResourceData) Timeout(key string) time.Duration {
	//TODO implement me
	panic("implement me")
}

func (f *FrameworkResourceData) GetOk(key string) (interface{}, bool) {
	//TODO implement me
	panic("implement me")
}

func (f *FrameworkResourceData) GetOkExists(key string) (interface{}, bool) {
	//TODO implement me
	panic("implement me")
}
