package tombuildsstuff

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azurerm/internal/sdk"
)

var _ sdk.DataSource = FakeDataSource{}

type FakeDataSource struct {
}

func (f FakeDataSource) Arguments() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func (f FakeDataSource) Attributes() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"location": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func (f FakeDataSource) ModelObject() interface{} {
	return nil
}

func (f FakeDataSource) ResourceType() string {
	return "azurerm_fake"
}

func (f FakeDataSource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			metadata.ResourceData.SetId("tombuildsstuff")
			metadata.ResourceData.Set("location", "Berlin")
			return nil
		},
		Timeout: 10 * time.Minute,
	}
}
