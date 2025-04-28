package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	_ "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	_ "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = (*zip2b64Provider)(nil)

//var _ provider.ProviderWithMetadata = (*zip2b64Provider)(nil)

type zip2b64Provider struct{}

func New() func() provider.Provider {
	return func() provider.Provider {
		return &zip2b64Provider{}
	}
}

func (p *zip2b64Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *zip2b64Provider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "thisisbud"
}

func (p *zip2b64Provider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		//NewDataSource,
	}
}

func (p *zip2b64Provider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewResource,
	}
}

func (p *zip2b64Provider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
}
