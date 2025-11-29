package datasources 

import ( 
  "context"
  "fmt"

  "github.com/hashicorp/terraform-plugin-framework/datasource" 
  "github.com/hashicorp/terraform-plugin-framework/datasource/schema"

  "golang.org/x/crypto/ssh"
) 

// 実装が期待されるインターフェースを満たしていることを確認します。
var (
  // メタデータやスキーマ、リードが使える
  _ datasource.DataSource = &rke2ServerNodeDataSource{}
  // 設定値のメソッドを使える
  _ datasource.DataSourceWithConfigure = &rke2ServerNodeDataSource{}
) 

// Newrke2ServerNodeDataSource は、プロバイダーの実装を簡素化するヘルパー関数です。
func NewRke2ServerNodeDataSource () datasource.DataSource { 
  return & rke2ServerNodeDataSource {} 
} 

// rke2ServerNodeDataSource はデータ ソースの実装です。
type rke2ServerNodeDataSource struct {
  client *ssh.Client
}

// メタデータはデータ ソース タイプ名を返します。
func (d *rke2ServerNodeDataSource ) Metadata (_ context.Context , req datasource.MetadataRequest , resp * datasource.MetadataResponse ) { 
  resp.TypeName = req.ProviderTypeName + "_server_node" 
} 

// Schema はデータ ソースのスキーマを定義します。
func (d *rke2ServerNodeDataSource ) Schema (_ context.Context , _ datasource.SchemaRequest , resp * datasource.SchemaResponse ) { 
  resp.Schema = schema.Schema {
    Attributes: map[string]schema.Attribute {

    },
  }
} 

// Read は最新のデータで Terraform の状態を更新します。
func (d *rke2ServerNodeDataSource ) Read (ctx context.Context , req datasource.ReadRequest , resp * datasource.ReadResponse ) { 
}

// sshクライアントを取得
func (d *rke2ServerNodeDataSource ) Configure (ctx context.Context , req datasource.ConfigureRequest , resp * datasource.ConfigureResponse ) {
  if req.ProviderData == nil {
    return
  }

  client, ok := req.ProviderData.(*ssh.Client)
  if !ok {
    resp.Diagnostics.AddError(
      "Unexpected Data Source Configure Type",
      fmt.Sprintf("Expected *ssh.Client, got %T", req.ProviderData),
    )
    return
  }
  d.client = client
}