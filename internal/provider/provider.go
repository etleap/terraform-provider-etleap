// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/etleap/terraform-provider-etleap/internal/sdk"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &EtleapProvider{}

type EtleapProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// EtleapProviderModel describes the provider data model.
type EtleapProviderModel struct {
	ServerURL types.String `tfsdk:"server_url"`
	Username  types.String `tfsdk:"username"`
	Password  types.String `tfsdk:"password"`
}

func (p *EtleapProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "etleap"
	resp.Version = p.version
}

func (p *EtleapProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `Etleap API v2: Etleap API v2`,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to https://api.etleap.com/api/v2)",
				Optional:            true,
				Required:            false,
			},
			"username": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *EtleapProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data EtleapProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "https://api.etleap.com/api/v2"
	}

	username := data.Username.ValueString()
	password := data.Password.ValueString()
	security := shared.Security{
		Username: username,
		Password: password,
	}

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *EtleapProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewConnectionACTIVECAMPAIGNResource,
		NewConnectionBIGQUERYResource,
		NewConnectionBINGADSResource,
		NewConnectionBLACKLINEResource,
		NewConnectionBRAINTREEResource,
		NewConnectionCONFLUENTCLOUDResource,
		NewConnectionCOUPAResource,
		NewConnectionCRITEOResource,
		NewConnectionDb2Resource,
		NewConnectionDb2SHARDEDResource,
		NewConnectionDELTALAKEResource,
		NewConnectionEGNYTEResource,
		NewConnectionELASTICSEARCHResource,
		NewConnectionELLUMINATEResource,
		NewConnectionELOQUAResource,
		NewConnectionERPXResource,
		NewConnectionFACEBOOKADSResource,
		NewConnectionFIFTEENFIVEResource,
		NewConnectionFRESHCHATResource,
		NewConnectionFRESHSALESResource,
		NewConnectionFRESHWORKSResource,
		NewConnectionFTPResource,
		NewConnectionGONGResource,
		NewConnectionGOOGLEADSResource,
		NewConnectionGOOGLEANALYTICSGa4Resource,
		NewConnectionGOOGLECLOUDSTORAGEResource,
		NewConnectionGOOGLESHEETSResource,
		NewConnectionHUBSPOTResource,
		NewConnectionICEBERGResource,
		NewConnectionIMPACTRADIUSResource,
		NewConnectionINTERCOMResource,
		NewConnectionJIRAALIGNResource,
		NewConnectionJIRACLOUDResource,
		NewConnectionKAFKAResource,
		NewConnectionKUSTOMERResource,
		NewConnectionLDAPResource,
		NewConnectionLDAPVIRTUALLISTVIEWResource,
		NewConnectionLINKEDINADSResource,
		NewConnectionMARKETOResource,
		NewConnectionMICROSOFTENTRAIDResource,
		NewConnectionMIXPANELResource,
		NewConnectionMONGODBResource,
		NewConnectionMYSQLResource,
		NewConnectionMYSQLSHARDEDResource,
		NewConnectionNETSUITEResource,
		NewConnectionNETSUITEV2Resource,
		NewConnectionORACLEResource,
		NewConnectionORACLESHARDEDResource,
		NewConnectionOUTLOOKResource,
		NewConnectionOUTREACHResource,
		NewConnectionPINTERESTADSResource,
		NewConnectionPOSTGRESResource,
		NewConnectionPOSTGRESSHARDEDResource,
		NewConnectionQUORAADSResource,
		NewConnectionRAVEMEDIDATAResource,
		NewConnectionRECURLYResource,
		NewConnectionREDSHIFTResource,
		NewConnectionREDSHIFTSHARDEDResource,
		NewConnectionS3DATALAKEResource,
		NewConnectionS3INPUTResource,
		NewConnectionSALESFORCEResource,
		NewConnectionSALESFORCEMARKETINGCLOUDResource,
		NewConnectionSAPCONCURResource,
		NewConnectionSAPHANAResource,
		NewConnectionSAPHANASHARDEDResource,
		NewConnectionSEISMICResource,
		NewConnectionSERVICENOWResource,
		NewConnectionSFTPResource,
		NewConnectionSHOPIFYResource,
		NewConnectionSKYWARDResource,
		NewConnectionSNAPCHATADSResource,
		NewConnectionSNOWFLAKEResource,
		NewConnectionSNOWFLAKESHARDEDResource,
		NewConnectionSQLSERVERResource,
		NewConnectionSQLSERVERSHARDEDResource,
		NewConnectionSQSResource,
		NewConnectionSQUAREResource,
		NewConnectionSTRIPEResource,
		NewConnectionSUMTOTALResource,
		NewConnectionTHETRADEDESKResource,
		NewConnectionTIKTOKADSResource,
		NewConnectionTWILIOResource,
		NewConnectionTWITTERADSResource,
		NewConnectionUSERDEFINEDAPIResource,
		NewConnectionUSERVOICEResource,
		NewConnectionVEEVAResource,
		NewConnectionVERIZONMEDIADSPResource,
		NewConnectionWORKDAYREPORTResource,
		NewConnectionWORKFRONTResource,
		NewConnectionZENDESKResource,
		NewConnectionZOOMPHONEResource,
		NewConnectionZUORAResource,
		NewDbtScheduleResource,
		NewModelResource,
		NewPipelineResource,
		NewTeamResource,
	}
}

func (p *EtleapProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewConnectionACTIVECAMPAIGNDataSource,
		NewConnectionBIGQUERYDataSource,
		NewConnectionBINGADSDataSource,
		NewConnectionBLACKLINEDataSource,
		NewConnectionBRAINTREEDataSource,
		NewConnectionCONFLUENTCLOUDDataSource,
		NewConnectionCOUPADataSource,
		NewConnectionCRITEODataSource,
		NewConnectionDb2DataSource,
		NewConnectionDb2SHARDEDDataSource,
		NewConnectionDELTALAKEDataSource,
		NewConnectionEGNYTEDataSource,
		NewConnectionELASTICSEARCHDataSource,
		NewConnectionELLUMINATEDataSource,
		NewConnectionELOQUADataSource,
		NewConnectionERPXDataSource,
		NewConnectionFACEBOOKADSDataSource,
		NewConnectionFIFTEENFIVEDataSource,
		NewConnectionFRESHCHATDataSource,
		NewConnectionFRESHSALESDataSource,
		NewConnectionFRESHWORKSDataSource,
		NewConnectionFTPDataSource,
		NewConnectionGONGDataSource,
		NewConnectionGOOGLEADSDataSource,
		NewConnectionGOOGLEANALYTICSGa4DataSource,
		NewConnectionGOOGLECLOUDSTORAGEDataSource,
		NewConnectionGOOGLESHEETSDataSource,
		NewConnectionHUBSPOTDataSource,
		NewConnectionICEBERGDataSource,
		NewConnectionIMPACTRADIUSDataSource,
		NewConnectionINTERCOMDataSource,
		NewConnectionJIRAALIGNDataSource,
		NewConnectionJIRACLOUDDataSource,
		NewConnectionKAFKADataSource,
		NewConnectionKUSTOMERDataSource,
		NewConnectionLDAPDataSource,
		NewConnectionLDAPVIRTUALLISTVIEWDataSource,
		NewConnectionLINKEDINADSDataSource,
		NewConnectionMARKETODataSource,
		NewConnectionMICROSOFTENTRAIDDataSource,
		NewConnectionMIXPANELDataSource,
		NewConnectionMONGODBDataSource,
		NewConnectionMYSQLDataSource,
		NewConnectionMYSQLSHARDEDDataSource,
		NewConnectionNETSUITEDataSource,
		NewConnectionNETSUITEV2DataSource,
		NewConnectionORACLEDataSource,
		NewConnectionORACLESHARDEDDataSource,
		NewConnectionOUTLOOKDataSource,
		NewConnectionOUTREACHDataSource,
		NewConnectionPINTERESTADSDataSource,
		NewConnectionPOSTGRESDataSource,
		NewConnectionPOSTGRESSHARDEDDataSource,
		NewConnectionQUORAADSDataSource,
		NewConnectionRAVEMEDIDATADataSource,
		NewConnectionRECURLYDataSource,
		NewConnectionREDSHIFTDataSource,
		NewConnectionREDSHIFTSHARDEDDataSource,
		NewConnectionS3DATALAKEDataSource,
		NewConnectionS3INPUTDataSource,
		NewConnectionSALESFORCEDataSource,
		NewConnectionSALESFORCEMARKETINGCLOUDDataSource,
		NewConnectionSAPCONCURDataSource,
		NewConnectionSAPHANADataSource,
		NewConnectionSAPHANASHARDEDDataSource,
		NewConnectionSEISMICDataSource,
		NewConnectionSERVICENOWDataSource,
		NewConnectionSFTPDataSource,
		NewConnectionSHOPIFYDataSource,
		NewConnectionSKYWARDDataSource,
		NewConnectionSNAPCHATADSDataSource,
		NewConnectionSNOWFLAKEDataSource,
		NewConnectionSNOWFLAKESHARDEDDataSource,
		NewConnectionSQLSERVERDataSource,
		NewConnectionSQLSERVERSHARDEDDataSource,
		NewConnectionSQSDataSource,
		NewConnectionSQUAREDataSource,
		NewConnectionSTRIPEDataSource,
		NewConnectionSUMTOTALDataSource,
		NewConnectionTHETRADEDESKDataSource,
		NewConnectionTIKTOKADSDataSource,
		NewConnectionTWILIODataSource,
		NewConnectionTWITTERADSDataSource,
		NewConnectionUSERDEFINEDAPIDataSource,
		NewConnectionUSERVOICEDataSource,
		NewConnectionVEEVADataSource,
		NewConnectionVERIZONMEDIADSPDataSource,
		NewConnectionWORKDAYREPORTDataSource,
		NewConnectionWORKFRONTDataSource,
		NewConnectionZENDESKDataSource,
		NewConnectionZOOMPHONEDataSource,
		NewConnectionZUORADataSource,
		NewDbtScheduleDataSource,
		NewModelDataSource,
		NewPipelineDataSource,
		NewTeamDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &EtleapProvider{
			version: version,
		}
	}
}
