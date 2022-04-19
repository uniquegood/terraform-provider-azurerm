//go:build framework
// +build framework

package sdk

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func frameworkAttributeFromPluginSdkType(input *schema.Schema) (*tfsdk.Attribute, error) {
	simpleTypes := map[schema.ValueType]attr.Type{
		schema.TypeBool:   types.BoolType,
		schema.TypeInt:    types.Int64Type,
		schema.TypeFloat:  types.Float64Type,
		schema.TypeString: types.StringType,
	}
	if frameworkType, ok := simpleTypes[input.Type]; ok {
		return mapAttribute(tfsdk.Attribute{
			Type: frameworkType,
		}, input), nil
	}

	if input.Type == schema.TypeMap {
		if input.Elem == nil {
			return nil, fmt.Errorf("the Elem was nil for the Map type")
		}
		v, ok := input.Elem.(*schema.Schema)
		if !ok {
			return nil, fmt.Errorf("expected an Elem for the Map item")
		}
		if v.Type != schema.TypeString {
			return nil, fmt.Errorf("a Map type must be of TypeString in Plugin SDKv2")
		}

		return mapAttribute(tfsdk.Attribute{
			Type: types.MapType{
				// @tombuildsstuff opinionated, btu all maps in Plugin SDK v2 are Maps of Strings
				ElemType: types.StringType,
			},
		}, input), nil
	}

	if input.Type == schema.TypeList {
		if input.Elem == nil {
			return nil, fmt.Errorf("the Elem was nil for the List type")
		}

		// either it's a List of a Simple Type
		elem, ok := input.Elem.(*schema.Schema)
		if ok {
			nestedElemType, err := frameworkAttributeFromPluginSdkType(elem)
			if err != nil {
				return nil, fmt.Errorf("parsing nested object for list %+v", err)
			}

			attribute := tfsdk.Attribute{
				Type: types.ListType{
					ElemType: nestedElemType.Type,
				},
				// TODO: PlanModifiers to do Min/Max Items
			}
			return mapAttribute(attribute, input), nil
		}

		// or it's actually a List of an Object, either a Singular or Multiple objects
		resource, ok := input.Elem.(*schema.Resource)
		if !ok {
			return nil, fmt.Errorf("the List Elem was a not *schema.Resource or *schema.Schema - got: %+v", input.Elem)
		}
		nestedAttributes := make(map[string]tfsdk.Attribute)
		for k, v := range resource.Schema {
			nestedAttr, err := frameworkAttributeFromPluginSdkType(v)
			if err != nil {
				return nil, fmt.Errorf("converting list nested schema item %q to a Framework type: %+v", err)
			}

			nestedAttributes[k] = *nestedAttr
		}

		attribute := tfsdk.Attribute{
			Attributes: tfsdk.ListNestedAttributes(nestedAttributes, tfsdk.ListNestedAttributesOptions{
				// this is just boilerplate today
			}),
			// TODO: PlanModifiers to do Min/Max Items
		}
		if input.MaxItems == 1 {
			attribute.Attributes = tfsdk.SingleNestedAttributes(nestedAttributes)
		}
		return mapAttribute(attribute, input), nil
	}

	if input.Type == schema.TypeSet {
		if input.Elem == nil {
			return nil, fmt.Errorf("the Elem was nil for the Set type")
		}

		// either it's a List of a Simple Type
		elem, ok := input.Elem.(*schema.Schema)
		if ok {
			nestedElemType, err := frameworkAttributeFromPluginSdkType(elem)
			if err != nil {
				return nil, fmt.Errorf("parsing nested object for Set %+v", err)
			}

			attribute := tfsdk.Attribute{
				Type: types.SetType{
					ElemType: nestedElemType.Type,
				},
				// TODO: PlanModifiers to do Min/Max Items
			}
			return mapAttribute(attribute, input), nil
		}

		// or it's actually a List of an Object, either a Singular or Multiple objects
		resource, ok := input.Elem.(*schema.Resource)
		if !ok {
			return nil, fmt.Errorf("the List Elem was a not *schema.Resource or *schema.Schema - got: %+v", input.Elem)
		}
		nestedAttributes := make(map[string]tfsdk.Attribute)
		for k, v := range resource.Schema {
			nestedAttr, err := frameworkAttributeFromPluginSdkType(v)
			if err != nil {
				return nil, fmt.Errorf("converting list nested schema item %q to a Framework type: %+v", err)
			}

			nestedAttributes[k] = *nestedAttr
		}

		attribute := tfsdk.Attribute{
			Attributes: tfsdk.SetNestedAttributes(nestedAttributes, tfsdk.SetNestedAttributesOptions{
				// this is just boilerplate today
			}),
			// TODO: PlanModifiers to do Min/Max Items
		}
		if input.MaxItems == 1 {
			attribute.Attributes = tfsdk.SingleNestedAttributes(nestedAttributes)
		}
		return mapAttribute(attribute, input), nil
	}

	panic(fmt.Sprintf("unsupported plugin sdk type: %+v", input.Type))
}

func mapAttribute(input tfsdk.Attribute, old *schema.Schema) *tfsdk.Attribute {
	if old.Required {
		input.Required = true
	}
	if old.Optional {
		input.Optional = true
	}
	if old.Computed {
		input.Computed = true
	}
	if old.Sensitive {
		input.Sensitive = true
	}
	if old.Description != "" {
		// NOTE: DescriptionMarkdown is supported, but right now all of our fields are
		// non-Markdown, so this doesn't matter
		input.Description = old.Description
	}
	if old.Deprecated != "" {
		input.DeprecationMessage = old.Deprecated
	}

	input.PlanModifiers = planModifiersFromPluginSdkAttribute(input.PlanModifiers, old)
	input.Validators = validatorsFromPluginSdkAttribute(input.Validators, old)
	return &input
}

func planModifiersFromPluginSdkAttribute(input tfsdk.AttributePlanModifiers, old *schema.Schema) tfsdk.AttributePlanModifiers {
	// TODO: map any old ones across
	return input
}

func validatorsFromPluginSdkAttribute(input []tfsdk.AttributeValidator, old *schema.Schema) []tfsdk.AttributeValidator {
	// TODO: map any old ones across
	return input
}
