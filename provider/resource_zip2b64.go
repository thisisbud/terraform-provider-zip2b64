package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/thisisbud/terraform-provider-zip2b64/client"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*zip2b64Resource)(nil)

type zip2b64Resource struct {
	provider zip2b64Provider //nolint:unused
}

func NewResource() resource.Resource {
	return &zip2b64Resource{}
}

func (e *zip2b64Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "zip2b64"
}

func (e *zip2b64Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"filename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"base64file": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"filecontents_base64": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

type zip2b64ResourceData struct {
	Id                 types.String `tfsdk:"id"`
	Filename           types.String `tfsdk:"filename"`
	Base64File         types.String `tfsdk:"base64file"`
	FileContentsBase64 types.String `tfsdk:"filecontents_base64"`
}

func (e *zip2b64Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data zip2b64ResourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	base64file := data.Base64File.String()
	filenameToExtract := data.Filename.String()

	filecontentsBase64, err := client.ZipExtract(base64file, filenameToExtract)
	if err != nil {
		//return fmt.Errorf("error Getting resource '%v'", err)
		resp.Diagnostics.AddError(
			"Error extracting zip",
			fmt.Sprintf("... details ... %s", err),
		)
		return
	}

	data.Id = types.StringValue(filenameToExtract)
	data.Filename = types.StringValue(filenameToExtract)
	data.Base64File = types.StringValue(base64file)
	data.FileContentsBase64 = types.StringValue(filecontentsBase64)

	tflog.Trace(ctx, "created a resource")

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (e *zip2b64Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data zip2b64ResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	base64file := data.Base64File.String()
	filenameToExtract := data.Filename.String()

	filecontentsBase64, err := client.ZipExtract(base64file, filenameToExtract)
	if err != nil {
		//return fmt.Errorf("error Getting resource '%v'", err)
		resp.Diagnostics.AddError(
			"Error extracting zip",
			fmt.Sprintf("... details ... %s", err),
		)
		return
	}

	data.Id = types.StringValue(filenameToExtract)
	data.Filename = types.StringValue(filenameToExtract)
	data.Base64File = types.StringValue(base64file)
	data.FileContentsBase64 = types.StringValue(filecontentsBase64)

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (e *zip2b64Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data zip2b64ResourceData

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	base64file := data.Base64File.String()
	filenameToExtract := data.Filename.String()

	filecontentsBase64, err := client.ZipExtract(base64file, filenameToExtract)
	if err != nil {
		//return fmt.Errorf("error Getting resource '%v'", err)
		resp.Diagnostics.AddError(
			"Error extracting zip",
			fmt.Sprintf("... details ... %s", err),
		)
		return
	}

	data.Id = types.StringValue(filenameToExtract)
	data.Filename = types.StringValue(filenameToExtract)
	data.Base64File = types.StringValue(base64file)
	data.FileContentsBase64 = types.StringValue(filecontentsBase64)

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (e *zip2b64Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data zip2b64ResourceData

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete resource using 3rd party API.
}
