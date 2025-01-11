package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-postmark/internal/provider/datasource_server"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/mrz1836/postmark"
)

var _ datasource.DataSource = (*serverDataSource)(nil)

func NewServerDataSource() datasource.DataSource {
	return &serverDataSource{}
}

type serverDataSource struct {
	client *postmark.Client
}

func (d *serverDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server"
}

func (d *serverDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring the Server datasource")

	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*postmark.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *serverDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_server.ServerDataSourceSchema(ctx)
}

func (d *serverDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data datasource_server.ServerModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	resp.Diagnostics.Append(d.readFromAPI(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *serverDataSource) readFromAPI(ctx context.Context, server *datasource_server.ServerModel) diag.Diagnostics {
	res, err := d.client.GetServer(ctx, server.Id.ValueString())
	if err != nil {
		clientDiag := diag.NewErrorDiagnostic("Client Error", fmt.Sprintf("Unable to read server, got error: %s", err))
		return diag.Diagnostics{clientDiag}
	}

	server.Id = types.StringValue(strconv.FormatInt(res.ID, 10))
	server.Name = types.StringValue(res.Name)
	server.Color = types.StringValue(res.Color)

	apiTokenDiags := server.ApiTokens.ElementsAs(ctx, &res.APITokens, false)

	if apiTokenDiags.HasError() {
		return apiTokenDiags
	}

	server.ApiTokens, apiTokenDiags = types.ListValueFrom(ctx, server.ApiTokens.ElementType(ctx), res.APITokens)

	if apiTokenDiags.HasError() {
		return apiTokenDiags
	}

	server.DeliveryType = types.StringValue(res.DeliveryType)
	server.InboundAddress = types.StringValue(res.InboundAddress)
	server.InboundDomain = types.StringValue(res.InboundDomain)
	server.InboundHash = types.StringValue(res.InboundHash)
	server.InboundHookUrl = types.StringValue(res.InboundHookURL)
	server.InboundSpamThreshold = types.Int64Value(res.InboundSpamThreshold)
	server.PostFirstOpenOnly = types.BoolValue(res.PostFirstOpenOnly)
	server.RawEmailEnabled = types.BoolValue(res.RawEmailEnabled)
	server.ServerLink = types.StringValue(res.ServerLink)
	server.SmtpApiActivated = types.BoolValue(res.SMTPAPIActivated)
	server.IncludeBounceContentInHook = types.BoolValue(res.IncludeBounceContentInHook)
	server.EnableSmtpApiErrorHooks = types.BoolValue(res.EnableSMTPAPIErrorHooks)

	return nil
}
