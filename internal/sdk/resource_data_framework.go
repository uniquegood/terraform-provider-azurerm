//go:build framework
// +build framework

package sdk

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ ResourceData = &FrameworkResourceData{}

type FrameworkResourceData struct {
	ctx   context.Context
	state *tfsdk.State

	// config is the user-specified config, which isn't guaranteed to be available
	config *tfsdk.Config

	// plan is the difference between the old state and the new state
	plan *tfsdk.Plan
}

func NewFrameworkResourceData(ctx context.Context, state *tfsdk.State) *FrameworkResourceData {
	return &FrameworkResourceData{
		ctx:   ctx,
		state: state,
	}
}

// WithConfig adds the user-provided config to the ResourceData
func (f *FrameworkResourceData) WithConfig(config tfsdk.Config) {
	f.config = &config
}

// WithExistingID sets an existing known Resource ID into the state
func (f *FrameworkResourceData) WithExistingID(id string) {
	// TODO: should this be setting a local variable rather than setting it into the state?
	f.SetId(id)
}

// WithExistingState ...
func (f *FrameworkResourceData) WithExistingState(state tfsdk.State) {
	// TODO: implement me
}

// WithPlan sets an existing known Plan
func (f *FrameworkResourceData) WithPlan(plan tfsdk.Plan) {
	f.plan = &plan
}

func (f *FrameworkResourceData) Get(key string) interface{} {
	var out interface{}
	path := flatMapToAttributePath(key)
	f.state.GetAttribute(f.ctx, path, out)
	return out
}

//
//func (f *FrameworkResourceData) GetChange(key string) (original interface{}, updated interface{}) {
//	path := flatMapToAttributePath(key)
//	if f.plan != nil {
//		var oldVal interface{}
//		diag := f.plan.GetAttribute(f.ctx, path, &oldVal)
//		if diag == nil {
//			original = oldVal
//		}
//	} else if f.state != nil {
//		var oldVal interface{}
//		diag := f.state.GetAttribute(f.ctx, path, &oldVal)
//		if diag == nil {
//			original = oldVal
//		}
//	}
//
//	var newVal interface{}
//	diag := f.config.GetAttribute(f.ctx, path, &newVal)
//	if diag == nil {
//		updated = newVal
//	}
//	return
//}

func (f *FrameworkResourceData) GetFromConfig(key string) interface{} {
	//TODO implement me
	panic("implement me")
}

func (f *FrameworkResourceData) GetFromState(key string) interface{} {
	//TODO implement me
	panic("implement me")
}

func (f *FrameworkResourceData) HasChange(key string) bool {
	//TODO implement me
	panic("implement me")
}

func (f *FrameworkResourceData) HasChanges(keys []string) bool {
	for _, k := range keys {
		if f.HasChange(k) {
			return true
		}
	}

	return false
}

func (f *FrameworkResourceData) Id() string {
	return f.Get("id").(string)
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
		f.state.RemoveResource(f.ctx)
	} else {
		f.Set("id", id)
	}
}

func flatMapToAttributePath(key string) *tftypes.AttributePath {
	// TODO: implement this properly
	return tftypes.NewAttributePath().WithAttributeName(key)
}
