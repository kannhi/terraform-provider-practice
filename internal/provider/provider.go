package provider

import (
	"context"
	"os"
    "net"
    "strconv"
	"time"
    "fmt"
	// "go/types"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"golang.org/x/crypto/ssh"

	"terraform-provider-rke2/internal/provider/resources"
)

var (
    _ provider.Provider = &rke2Provider{}
)

func New(version string) func() provider.Provider {
    return func() provider.Provider {
        return &rke2Provider{
            version: version,
        }
    }
}

type rke2Provider struct {
    version string
}

func (p *rke2Provider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
    resp.TypeName = "rke2"
    resp.Version = p.version
}

func (p *rke2Provider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
    resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"user": schema.StringAttribute{
				Required: true,
			},
			"privatekey": schema.StringAttribute{
				Required: true,
				Sensitive: true,
			},
			"host": schema.StringAttribute{
				Required: true,
			},
			"timeout": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

// 「ユーザーの設定」→ tfsdkタグで指定されたGoの構造体のフィールドへ→ スキーマの型に合った**types.付きの型**（例: types.String）で値を受け取る
type rke2ProviderModel struct {
    User types.String `tfsdk:"user"`
	Privatekey types.String `tfsdk:"privatekey"`
	Host types.String `tfsdk:"host"`
	Timeout types.String `tfsdk:"timeout"`
}

func (p *rke2Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
    var config rke2ProviderModel
    diags := req.Config.Get(ctx, &config)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    // If practitioner provided a configuration value for any of the
    // attributes, it must be a known value.
	// terraformの設定的にreference等で設定値を引っ張ってくるときには`Unknown`という状態がある
	// ここはprovider設定なのでIsUnknownを許容しない
    if config.User.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
			// path.Root は、「エラーが設定ファイルの『どの項目』で起きたか」を指し示すための住所（パス）を作るメソッド
			// terraformのプラグインフレームワークのpathは今処理しているブロックの中からhostを取り出してくれる
			// blockがネストしていたら`path.Root("advanced_settings").AtName("host")`のように書く
            path.Root("user"),
            "Unknown user",
            "The provider cannot create the ssh user as there is an unknown configuration value for the ssh user. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the RKE2_USER environment variable.",
        )
    }

    if config.Privatekey.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
            path.Root("privatekey"),
            "Unknown private key",
            "The provider cannot create the ssh private key as there is an unknown configuration value for the ssh private key. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the RKE2_PRIVATE_KEY environment variable.",
        )
    }

    if config.Host.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
            path.Root("host"),
            "Unknown host",
            "The provider cannot create the ssh connection as there is an unknown configuration value for the ssh host. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the RKE2_HOST environment variable.",
        )
    }

    if config.Timeout.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
            path.Root("timeout"),
            "Unknown timeout",
            "The provider cannot create the ssh connection as there is an unknown configuration value for the ssh timeout. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the RKE2_TIMEOUT environment variable.",
        )
    }

    if resp.Diagnostics.HasError() {
        return
    }

    // Default values to environment variables, but override
    // with Terraform configuration value if set.
	// 環境変数から値を抽出する
    user := os.Getenv("RKE2_USER")
    privatekey := os.Getenv("RKE2_PRIVATE_KEY")
    host := os.Getenv("RKE2_HOST")
    timeout := os.Getenv("RKE2_TIMEOUT")

    if !config.User.IsNull() {
        user = config.User.ValueString()
    }

    if !config.Privatekey.IsNull() {
        privatekey = config.Privatekey.ValueString()
    }

    if !config.Host.IsNull() {
        host = config.Host.ValueString()
    }

    if !config.Timeout.IsNull() {
        timeout = config.Timeout.ValueString()
    }

    // If any of the expected configurations are missing, return
    // errors with provider-specific guidance.
    if user == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("user"),
            "Missing user",
            "The provider cannot create the ssh connection as there is a missing or empty value for the ssh user. "+
                "Set the user value in the configuration or use the RKE2_USER environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if privatekey == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("privatekey"),
            "Missing private key",
            "The provider cannot create the ssh connection as there is a missing or empty value for the ssh private key. "+
                "Set the private key value in the configuration or use the RKE2_PRIVATE_KEY environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if host == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("host"),
            "Missing host",
            "The provider cannot create the ssh connection as there is a missing or empty value for the ssh host. "+
                "Set the host value in the configuration or use the RKE2_HOST environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if timeout == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("timeout"),
            "Missing timeout",
            "The provider cannot create the ssh connection as there is a missing or empty value for the ssh timeout. "+
                "Set the timeout value in the configuration or use the RKE2_TIMEOUT environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if resp.Diagnostics.HasError() {
        return
    }

    // Create a new HashiCups client using the configuration values
	// APIでクライアントを作成(認証する)
    client, err := ssh_make_client(user, privatekey, host, timeout)
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create SSH client",
            fmt.Sprintf("Failed to create SSH client: %s", err.Error()),
        )
        return
    }

    // Make the HashiCups client available during DataSource and Resource
    // type Configure methods.
    resp.DataSourceData = client
    resp.ResourceData = client
}

func ssh_make_client(user string, private_key string, host string, timeout string) (*ssh.Client, error) {
    // targetの情報
	port := "22" // のちに変数化
    timeout_sec, _ := strconv.Atoi(timeout)
    ssh_signer, err := ssh.ParsePrivateKey([]byte(private_key))
    if err != nil {
        return nil, fmt.Errorf("failed to parse private key: %v", err)
    }
	// sshClientConfigの作成(target)
	targetSshConfig := &ssh.ClientConfig{
        User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(ssh_signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        Timeout:         time.Duration(timeout_sec),
	}
    
    // ↓ssh.Dialで事足りる可能性があったのでコメントアウト
    // // targetへ直接tcp接続
    // direct_conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Duration(timeout_sec))
    // if err != nil {
    //     return nil, fmt.Errorf("TCP接続に失敗（タイムアウト含む）: %v", err)
    // }
    
	// // TargetへのsshClientを作成
	// pConnect, pChans, pReqs, err := ssh.NewClientConn(direct_conn, net.JoinHostPort(host, port), targetSshConfig)
	// if err != nil {
    //     return nil, fmt.Errorf("SSH接続に失敗: %v", err)
	// }
	// client := ssh.NewClient(pConnect, pChans, pReqs)

    client, err := ssh.Dial("tcp", net.JoinHostPort(host, port), targetSshConfig)
    if err != nil {
        return nil, fmt.Errorf("SSH接続に失敗: %v", err)
    }

	return client, nil
}

// DataSources defines the data sources implemented in the provider.
func (p *rke2Provider) DataSources(_ context.Context) []func() datasource.DataSource {
    return nil
    // return []func() datasource.DataSource {
    //     Newrke2ServerNodeDataSource,
    // }
}

// Resources defines the resources implemented in the provider.
func (p *rke2Provider) Resources(_ context.Context) []func() resource.Resource {
    return []func() resource.Resource {
        // resources.NewRke2ServerNodeResource,
        resources.NewtestExecResource,
    }
}