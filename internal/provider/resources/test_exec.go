package resources

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"golang.org/x/crypto/ssh"
)

// 実装が想定されるインターフェースを満たしていることを確認します。
var ( 
    _ resource.Resource = &testExecResource {} 
    _ resource.ResourceWithConfigure = &testExecResource {}
)


// NewtestExecResource は、プロバイダーの実装を簡素化するヘルパー関数です。
func NewtestExecResource () resource.Resource { 
    return &testExecResource {} 
} 

// testExecResource はリソースの実装です。
type testExecResource struct {
	client *ssh.Client
}

type testExecResourceModel struct {
	Id       types.String `tfsdk:"id"`
	FileName types.String `tfsdk:"file_name"`
}

// Metadata はリソース タイプ名を返します。
func (r*testExecResource) Metadata (_ context.Context , req resource.MetadataRequest , resp * resource.MetadataResponse ) { 
    resp.TypeName = req.ProviderTypeName + "_test_exec" 
} 

// Schema はリソースのスキーマを定義します。
func (r *testExecResource) Schema (_ context.Context , _ resource.SchemaRequest , resp * resource.SchemaResponse ) { 
    resp.Schema = schema.Schema {
		Attributes: map[string]schema.Attribute {
			"id": schema.StringAttribute {
				Computed: true,
			},
			"file_name": schema.StringAttribute {
				Optional: true,
				Computed: true,
				Default: stringdefault.StaticString("/test_file"),
			},
		},
	} 
} 

// Create はリソースを作成し、初期の Terraform 状態を設定します。
func (r *testExecResource) Create (ctx context.Context , req resource.CreateRequest , resp * resource.CreateResponse ) {
	var plan testExecResourceModel
	// Plan からユーザーが指定した属性を構造体へ展開
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// ssh接続
	session, err := r.client.NewSession()
	if err != nil {
		resp.Diagnostics.AddError(
			"SSH接続に失敗",
			fmt.Sprintf("SSH接続に失敗: %v", err),
		)
		return
	}
	defer session.Close()

	// リモートでコマンドの実行
	err = session.Run("touch " + plan.FileName.ValueString())

	if err != nil {
		fmt.Printf("Stderr: %v\n", err)
	}
	// stateの作成結果を追加
	plan.Id = types.StringValue(uuid.NewString())

	// stateをセット
    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
} 

// Read は最新のデータで Terraform 状態を更新します。
func (r *testExecResource) Read (ctx context.Context , req resource.ReadRequest , resp * resource.ReadResponse) {
    var state testExecResourceModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *testExecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *testExecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

// 外部クライアントの作成
func (r *testExecResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*ssh.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *ssh.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.client = client
}