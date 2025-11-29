package resources

// import (
// 	"context"
// 	"fmt"

// 	"github.com/hashicorp/terraform-plugin-framework/resource"
// 	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
// 	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
// 	"github.com/hashicorp/terraform-plugin-framework/types"
// 	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
// 	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
// 	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
// 	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"

// 	"golang.org/x/crypto/ssh"
// )

// // Ensure the implementation satisfies the expected interfaces.
// var (
//     _ resource.Resource = &rke2ServerNodeResource{}
//     _ resource.ResourceWithConfigure = &rke2ServerNodeResource{}
// )

// // Newrke2ServerNodeResource is a helper function to simplify the provider implementation.
// func NewRke2ServerNodeResource() resource.Resource {
//     return &rke2ServerNodeResource{}
// }

// // rke2ServerNodeResource is the resource implementation.
// type rke2ServerNodeResource struct{
// 	client *ssh.Client
// }

// type rke2ServerNodeResourceModel struct {
// 	Config                            types.String `tfsdk:"config"`
// 	Debug                             types.Bool   `tfsdk:"debug"`
// 	DataDir                           types.String `tfsdk:"data_dir"`
// 	BindAddress                       types.String `tfsdk:"bind_address"`
// 	AdvertiseAddress                  types.String `tfsdk:"advertise_address"`
// 	TlsSan                            types.Set    `tfsdk:"tls_san"`
// 	TlsSanSecurity                    types.Bool   `tfsdk:"tls_san_security"`
// 	ClusterCidr                       types.String `tfsdk:"cluster_cidr"`
// 	ServiceCidr                       types.String `tfsdk:"service_cidr"`
// 	ServiceNodePortRange              types.String `tfsdk:"service_node_port_range"`
// 	ClusterDns                        types.String `tfsdk:"cluster_dns"`
// 	ClusterDomain                     types.String `tfsdk:"cluster_domain"`
// 	EgressSelectorMode                types.String `tfsdk:"egress_selector_mode"`
// 	ServiceLbNamespace                types.String `tfsdk:"service_lb_namespace"`
// 	Cni                               types.String `tfsdk:"cni"`
// 	WriteKubeconfig                   types.String `tfsdk:"write_kubeconfig"`
// 	WriteKubeconfigMode               types.String `tfsdk:"write_kubeconfig_mode"`
// 	HelmJobImage                      types.List   `tfsdk:"helm_job_image"`
// 	Token                             types.String `tfsdk:"token"`
// 	TokenFile                         types.String `tfsdk:"token_file"`
// 	AgentToken                        types.String `tfsdk:"agent_token"`
// 	AgentTokenFile                    types.String `tfsdk:"agent_token_file"`
// 	Server                            types.Bool   `tfsdk:"server"`
// 	ClusterReset                      types.Bool   `tfsdk:"cluster_reset"`
// 	ClusterResetRestorePath           types.String `tfsdk:"cluster_reset_restore_path"`
// 	EtcdExposeMetrics                 types.Bool   `tfsdk:"etcd_expose_metrics"`
// 	EtcdDisableSnapshots              types.Bool   `tfsdk:"etcd_disable_snapshots"`
// 	EtcdSnapshotName                  types.String `tfsdk:"etcd_snapshot_name"`
// 	EtcdSnapshotScheduleCron          types.String `tfsdk:"etcd_snapshot_schedule_cron"`
// 	EtcdSnapshotRetention             types.Int64  `tfsdk:"etcd_snapshot_retention"`
// 	EtcdSnapshotDir                   types.String `tfsdk:"etcd_snapshot_dir"`
// 	EtcdSnapshotCompress              types.Bool   `tfsdk:"etcd_snapshot_compress"`
// 	EtcdS3                            types.Bool   `tfsdk:"etcd_s3"`
// 	EtcdS3Endpoint                    types.String `tfsdk:"etcd_s3_endpoint"`
// 	EtcdS3EndpointCa                  types.String `tfsdk:"etcd_s3_endpoint_ca"`
// 	EtcdS3SkipSslVerify               types.String `tfsdk:"etcd_s3_skip_ssl_verify"`
// 	EtcdS3AccessKey                   types.String `tfsdk:"etcd_s3_access_key"`
// 	EtcdS3SecretKey                   types.String `tfsdk:"etcd_s3_secret_key"`
// 	EtcdS3Bucket                      types.String `tfsdk:"etcd_s3_bucket"`
// 	EtcdS3Region                      types.String `tfsdk:"etcd_s3_region"`
// 	EtcdS3Folder                      types.String `tfsdk:"etcd_s3_folder"`
// 	EtcdS3Insecure                    types.Bool   `tfsdk:"etcd_s3_insecure"`
// 	EtcdS3Timeout                     types.String `tfsdk:"etcd_s3_timeout"`
// 	EtcdS3Retention                   types.Number `tfsdk:"etcd_s3_retention"`
// 	KubeApiserverImage                types.String `tfsdk:"kube_apiserver_image"`
// 	KubeControllerManagerImage        types.String `tfsdk:"kube_controller_manager_image"`
// 	KubeSchedulerImage                types.String `tfsdk:"kube_scheduler_image"`
// 	KubeProxyImage                    types.String `tfsdk:"kube_proxy_image"`
// 	PauseImage                        types.String `tfsdk:"pause_image"`
// 	EtcdImage                         types.String `tfsdk:"etcd_image"`
// 	RuntimeImage                      types.String `tfsdk:"runtime_image"`
// 	CloudControllerManagerImage       types.String `tfsdk:"cloud_controller_manager_image"`
// 	KubeApiserverArg                  types.List   `tfsdk:"kube_apiserver_arg"`
// 	EtcdArg                           types.List   `tfsdk:"etcd_arg"`
// 	KubeControllerManagerArg          types.List   `tfsdk:"kube_controller_manager_arg"`
// 	KubeSchedulerArg                  types.List   `tfsdk:"kube_scheduler_arg"`
// 	KubeCloudControllerManagerArg     types.List   `tfsdk:"kube_cloud_controller_manager_arg"`
// 	Disable                           types.List   `tfsdk:"disable"`
// 	DisableScheduler                  types.Bool   `tfsdk:"disable_scheduler"`
// 	DisableCloudController            types.Bool   `tfsdk:"disable_cloud_controller"`
// 	DisableKubeProxy                  types.Bool   `tfsdk:"disable_kube_proxy"`
// 	EnableServicelb                   types.Bool   `tfsdk:"enable_servicelb"`
// 	IngressController                 types.String `tfsdk:"ingress_controller"`
// 	ControlPlaneResourceRequests      types.List   `tfsdk:"control_plane_resource_requests"`
// 	ControlPlaneResourceLimits        types.List   `tfsdk:"control_plane_resource_limits"`
// 	ControlPlaneProbeConfiguration    types.List   `tfsdk:"control_plane_probe_configuration"`
// 	KubeApiserverExtraMount           types.List   `tfsdk:"kube_apiserver_extra_mount"`
// 	KubeSchedulerExtraMount           types.List   `tfsdk:"kube_scheduler_extra_mount"`
// 	KubeControllerManagerExtraMount   types.List   `tfsdk:"kube_controller_manager_extra_mount"`
// 	KubeProxyExtraMount               types.List   `tfsdk:"kube_proxy_extra_mount"`
// 	EtcdExtraMount                    types.List   `tfsdk:"etcd_extra_mount"`
// 	CloudControllerManagerExtraMount  types.List   `tfsdk:"cloud_controller_manager_extra_mount"`
// 	KubeApiserverExtraEnv             types.List   `tfsdk:"kube_apiserver_extra_env"`
// 	KubeSchedulerExtraEnv             types.List   `tfsdk:"kube_scheduler_extra_env"`
// 	KubeControllerManagerExtraEnv     types.List   `tfsdk:"kube_controller_manager_extra_env"`
// 	KubeProxyExtraEnv                 types.List   `tfsdk:"kube_proxy_extra_env"`
// 	EtcdExtraEnv                      types.List   `tfsdk:"etcd_extra_env"`
// 	CloudControllerManagerExtraEnv    types.List   `tfsdk:"cloud_controller_manager_extra_env"`
// 	CloudProviderName                 types.String `tfsdk:"cloud_provider_name"`
// 	CloudProviderConfig               types.String `tfsdk:"cloud_provider_config"`
// 	Profile                           types.String `tfsdk:"profile"`
// 	AuditPolicyFile                   types.String `tfsdk:"audit_policy_file"`
// 	PodSecurityAdmissionConfigFile    types.String `tfsdk:"pod_security_admission_config_file"`
// 	SecretsEncryptionProvider         types.String `tfsdk:"secrets_encryption_provider"`
// 	NodeName                          types.String `tfsdk:"node_name"`
// 	WithNodeId                        types.Bool   `tfsdk:"with_node_id"`
// 	NodeLabel                         types.List   `tfsdk:"node_label"`
// 	NodeTaint                         types.List   `tfsdk:"node_taint"`
// 	ImageCredentialProviderBinDir     types.String `tfsdk:"image_credential_provider_bin_dir"`
// 	ImageCredentialProviderConfig     types.String `tfsdk:"image_credential_provider_config"`
// 	ProtectKernelDefaults             types.Bool   `tfsdk:"protect_kernel_defaults"`
// 	Selinux                           types.Bool   `tfsdk:"selinux"`
// 	LbServerPort                      types.Int64  `tfsdk:"lb_server_port"`
// 	EmbeddedRegistry                  types.Bool   `tfsdk:"embedded_registry"`
// 	EnablePprof                       types.Bool   `tfsdk:"enable_pprof"`
// 	KubeletPath                       types.String `tfsdk:"kubelet_path"`
// 	ContainerRuntimeEndpoint          types.String `tfsdk:"container_runtime_endpoint"`
// 	DefaultRuntime                    types.String `tfsdk:"default_runtime"`
// 	Snapshotter                       types.String `tfsdk:"snapshotter"`
// 	PrivateRegistry                   types.String `tfsdk:"private_registry"`
// 	SystemDefaultRegistry             types.String `tfsdk:"system_default_registry"`
// 	DisableDefaultRegistryEndpoint    types.Bool   `tfsdk:"disable_default_registry_endpoint"`
// 	NodeIp                            types.List   `tfsdk:"node_ip"`
// 	NodeExternalIp                    types.List   `tfsdk:"node_external_ip"`
// 	ResolvConf                        types.String `tfsdk:"resolv_conf"`
// 	KubeletArg                        types.List   `tfsdk:"kubelet_arg"`
// 	KubeProxyArg                      types.List   `tfsdk:"kube_proxy_arg"`
// }

// // Metadata returns the resource type name.
// func (r *rke2ServerNodeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
//     resp.TypeName = req.ProviderTypeName + "_order"
// }

// // Schema defines the schema for the resource.
// // 使用上設定値はアルファベットに読み込まれ、最後に読み込まれたやつがその値を取る
// // ラベルやテイントなどのリストは`node-label+: sss`のように書くことで既存のリストに追加されるようになる
// // 参考: https://docs.rke2.io/reference/server_config
// func (r *rke2ServerNodeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
// 	resp.Schema = schema.Schema{
// 		Attributes: map[string]schema.Attribute{

// 			////////////////////////// common //////////////////////////
// 			// config fileのパス
// 			"config": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "RKE2設定ファイルのパス (例: /etc/rancher/rke2/config.yaml)。",
// 				Default: stringdefault.StaticString("/etc/rancher/rke2/config.yaml"),
// 			},
// 			// debug logを吐くか
// 			"debug": schema.BoolAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "デバッグロギングを有効にします。",
// 			},
// 			// etcdや様々な実行ファイルやキャッシュが保存されている箇所
// 			"data_dir": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "RKE2が状態、etcdデータ、マニフェストを保存するディレクトリ (デフォルト: /var/lib/rancher/rke2)。",
// 				Default: stringdefault.StaticString("/var/lib/rancher/rke2"	),
// 			},

// 			////////////////////////// Listener //////////////////////////
// 			// 特に設定しない場合は0.0.0.0として全てのネットワークインターフェースの接続を待ち受けることになる
// 			"bind_address": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "RKE2サーバーが内部通信のためにバインドするIPアドレス (デフォルト: 0.0.0.0)。",
// 				Default: stringdefault.StaticString("0.0.0.0"),
// 			},
// 			// 他のノードへの通知に使用されるアドレス
// 			"advertise_address": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "クラスター内の他のノードにRKE2サーバーを通知するために使用されるIPアドレス。",
// 				Default: stringdefault.StaticString("node-external-ip/node-ip"),
// 			},
// 			// 追加するホスト名
// 			"tls_san": schema.SetAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "サーバー証明書を生成するための追加のホスト名またはIPアドレス。",
// 			},
// 			// k8s apiserverに関連づいていないsanを拒否することでサーバー証明書を守るかどうか
// 			"tls_san_security": schema.BoolAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "Kubernetes APIサーバーに関連付けられていないTLS SANを拒否します。",
// 				Default: booldefault.StaticBool(true),
// 			},

// 			////////////////////////// Networking //////////////////////////
// 			"cluster_cidr": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "Podに割り当てるCIDR範囲 (デフォルト: 10.42.0.0/16)。",
// 				Default: stringdefault.StaticString("10.42.0.0/16"),
// 			},
// 			"service_cidr": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "Serviceに割り当てるCIDR範囲 (デフォルト: 10.43.0.0/16)。",
// 				Default: stringdefault.StaticString("10.43.0.0/16"),
// 			},
// 			"service_node_port_range": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "NodePortサービスのポート範囲 (デフォルト: 30000-32767)。",
// 				Default: stringdefault.StaticString("30000-32767"),
// 			},
// 			// corednsのip
// 			"cluster_dns": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "クラスターDNSサービス (CoreDNS) のIPアドレス。",
// 				Default: stringdefault.StaticString("10.43.0.10"),
// 			},
// 			"cluster_domain": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "クラスタードメイン (デフォルト: cluster.local)。",
// 				Default: stringdefault.StaticString("cluster.local"),
// 			},
// 			// apiサーバーからpodへの接続がエージェントノード経由か、直接ネットワーク接続をするか
// 			// agent, cluster, pod, disabled
// 			"egress_selector_mode": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "APIサーバーからPodへのトラフィック (kubectl exec/logsなど) のルーティング方法 (agent, cluster, pod, disabled)。",
// 				Validators: []validator.String{
// 					stringvalidator.OneOf("agent", "cluster", "pod", "disabled"),
// 				},
// 				Default: stringdefault.StaticString("agent"),
// 			},
// 			"service_lb_namespace": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "サービスロードバランサーをデプロイする名前空間。",
// 				Default: stringdefault.StaticString("kube-system"),
// 			},
// 			// calico, canal, cillium
// 			"cni": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "使用するCNIプラグイン (例: canal, cilium, none)。",
// 				Validators: []validator.String{
// 					stringvalidator.OneOf("calico", "canal", "cilium", "none"),
// 				},
// 				Default: stringdefault.StaticString("canal"),
// 			},

// 			////////////////////////// Client //////////////////////////
// 			// kubeconfigファイルの書き込みモード
// 			"write_kubeconfig": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "kubeconfigファイルのアウトプット先",
// 			},
// 			"write_kubeconfig_mode": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "kubeconfigファイルを書き込む際のパーミッション (例: 0644)。",
// 			},

// 			////////////////////////// Helm //////////////////////////
// 			"helm_job_image": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Computed:    true,
// 			},

// 			////////////////////////// Cluster //////////////////////////
// 			// サーバーノードがクラスターに参加するためのトークン
// 			"token": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "サーバーノードがクラスターに参加するための共有シークレット/トークン。",
// 				Sensitive:   true,
// 			},
// 			"token_file": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "サーバーノードがクラスターに参加するための共有シークレット/トークンが保存されているファイル",
// 			},
// 			"agent_token": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "エージェントノードがクラスターに参加するための共有シークレット/トークン。",
// 				Sensitive:   true,
// 			},
// 			"agent_token_file": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "エージェントノードがクラスターに参加するための共有シークレット/トークンが保存されているファイル",
// 			},
// 			"server": schema.BoolAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "このノードをサーバーノードとして動作させるかどうかを定義します。",
// 			},
// 			"cluster_reset": schema.BoolAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "すべてのクラスター接続を怪し、新しいクラスターにジョインするか",
// 			},

// 			////////////////////////// Database //////////////////////////
// 			"cluster_reset_restore_path": schema.StringAttribute{
// 				Optional: true,
// 				Computed: true,
// 			},
// 			"etcd_expose_metrics": schema.BoolAttribute{
// 				Optional: true,
// 				Computed: true,
// 				Default: booldefault.StaticBool(false),
// 			},
// 			"etcd_disable_snapshots": schema.BoolAttribute{
// 				Optional: true,
// 				Computed: true,
// 			},
// 			"etcd_snapshot_name": schema.StringAttribute{
// 				Optional: true,
// 				Computed: true,
// 			},
// 			"etcd_snapshot_schedule_cron": schema.StringAttribute{
// 				Optional: true,
// 				Computed: true,
// 				Default: stringdefault.StaticString("0 */12 * * *"),
// 			},
// 			// etcdスナップショットの保持世代数
// 			"etcd_snapshot_retention": schema.Int64Attribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "etcdスナップショットの保持世代数。",
// 				Default: int64default.StaticInt64(5),
// 			},
// 			// etcdスナップショットを保存するディレクトリ
// 			"etcd_snapshot_dir": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "etcdスナップショットを保存するディレクトリのパス。",
// 				// 本来的には${data-dir}/db/snapshotsがデフォルト
// 			},
// 			// etcdスナップショットを圧縮するかどうか
// 			"etcd_snapshot_compress": schema.BoolAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "etcdスナップショットをgzipで圧縮するかどうか。",
// 			},
// 			"etcd_s3": schema.BoolAttribute{
// 				Optional: true,
// 				Computed: true,
// 			},
// 			// etcdのピア通信に使用するIPアドレス (advertise-addressと連動)
// 			"etcd_s3_endpoint": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "etcdのS3バックアップに使用するカスタムS3互換エンドポイント。",
// 				Default: stringdefault.StaticString("s3.amazonaws.com"),
// 			},
// 			"etcd_s3_endpoint_ca": schema.StringAttribute{
// 				Optional: true,
// 				Computed: true,
// 			},
// 			"etcd_s3_skip_ssl_verify": schema.StringAttribute{
// 				Optional: true,
// 				Computed: true,
// 			},
// 			// S3バックアップ用のアクセスキー
// 			"etcd_s3_access_key": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "S3バケットへのバックアップに使用するアクセスキー。",
// 			},
// 			// S3バックアップ用のシークレットキー
// 			"etcd_s3_secret_key": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "S3バケットへのバックアップに使用するシークレットキー。",
// 			},
// 			// S3バケット名
// 			"etcd_s3_bucket": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "S3バックアップを保存するバケット名。",
// 			},
// 			// S3バケットのリージョン
// 			"etcd_s3_region": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "S3バケットが配置されているリージョン。",
// 				Default: stringdefault.StaticString("us-east-1"),
// 			},
// 			// S3バックアップのパスプレフィックス
// 			"etcd_s3_folder": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "S3バケット内のパスプレフィックス。",
// 			},
// 			"etcd_s3_insecure": schema.BoolAttribute{
// 				Optional: true,
// 				Computed: true,
// 			},
// 			// S3アップロードのタイムアウト
// 			"etcd_s3_timeout": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "S3アップロードのタイムアウト時間。",
// 				Default: stringdefault.StaticString("5m0s"),
// 			},
// 			"etcd_s3_retention": schema.Int64Attribute{
// 				Optional: true,
// 				Computed: true,
// 				Default: int64default.StaticInt64(5),
// 			},

// 			////////////////////////// Image　//////////////////////////
// 			"kube_apiserver_image": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "Kubernetes APIサーバーのコンテナイメージを上書きします。",
// 			},

// 			// `--kube-controller-manager-image`
// 			"kube_controller_manager_image": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "Kubernetes Controller Managerのコンテナイメージを上書きします。",
// 			},

// 			// `--kube-scheduler-image`
// 			"kube_scheduler_image": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "Kubernetes Schedulerのコンテナイメージを上書きします。",
// 			},

// 			// `--kube-proxy-image`
// 			"kube_proxy_image": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "Kube-proxyのコンテナイメージを上書きします。",
// 			},

// 			// `--pause-image`
// 			"pause_image": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "Pod Infra Container (pause) のイメージを上書きします。",
// 			},

// 			// `--etcd-image`
// 			"etcd_image": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "組み込みetcdのコンテナイメージを上書きします。",
// 			},
// 			"runtime_image": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "ランタイムバイナリ (containerd, kubectl, crictlなど) のイメージを上書きします。",
// 			},
// 			"cloud_controller_manager_image": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "cloud-controller-managerのイメージを上書きします。",
// 			},

// 			////////////////////////// Components //////////////////////////
// 			"kube_apiserver_arg": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kube-apiserverプロセスのカスタマイズフラグ。",
// 			},
// 			"etcd_arg": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "etcdプロセスのカスタマイズフラグ。",
// 			},
// 			"kube_controller_manager_arg": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kube-controller-managerプロセスのカスタマイズフラグ。",
// 			},
// 			"kube_scheduler_arg": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kube-schedulerプロセスのカスタマイズフラグ。",
// 			},
// 			"kube_cloud_controller_manager_arg": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kube-cloud-controller-managerプロセスのカスタマイズフラグ。",
// 			},
// 			"disable": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "無効にするコンポーネント (rke2-coredns, rke2-ingress-nginx, rke2-metrics-server)。",
// 			},
// 			"disable_scheduler": schema.BoolAttribute{
// 				Optional:    true,
// 				Description: "Kubernetesのデフォルトスケジューラーを無効にします。",
// 			},
// 			"disable_cloud_controller": schema.BoolAttribute{
// 				Optional:    true,
// 				Description: "RKE2のデフォルトクラウドコントローラーマネージャーを無効にします。",
// 			},
// 			"disable_kube_proxy": schema.BoolAttribute{
// 				Optional:    true,
// 				Description: "kube-proxyの実行を無効にします。",
// 			},
// 			"enable_servicelb": schema.BoolAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "RKE2のデフォルトクラウドコントローラーマネージャーのサービスコントローラーを有効にします。",
// 			},
// 			"ingress_controller": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "デプロイするIngressコントローラー (none, ingress-nginx, traefik)。",
// 				Validators: []validator.String{
// 					stringvalidator.OneOf("none", "ingress-nginx", "traefik"),
// 				},
// 			},
// 			"control_plane_resource_requests": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "コントロールプレーンのリソースリクエスト。",
// 			},
// 			"control_plane_resource_limits": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "コントロールプレーンのリソース制限。",
// 			},
// 			"control_plane_probe_configuration": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "コントロールプレーンのプローブ設定。",
// 			},
// 			"kube_apiserver_extra_mount": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kube-apiserverの追加ボリュームマウント。",
// 			},
// 			"kube_scheduler_extra_mount": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kube-schedulerの追加ボリュームマウント。",
// 			},
// 			"kube_controller_manager_extra_mount": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kube-controller-managerの追加ボリュームマウント。",
// 			},
// 			"kube_proxy_extra_mount": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kube-proxyの追加ボリュームマウント。",
// 			},
// 			"etcd_extra_mount": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "etcdの追加ボリュームマウント。",
// 			},
// 			"cloud_controller_manager_extra_mount": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "cloud-controller-managerの追加ボリュームマウント。",
// 			},
// 			"kube_apiserver_extra_env": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kube-apiserverの追加環境変数。",
// 			},
// 			"kube_scheduler_extra_env": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kube-schedulerの追加環境変数。",
// 			},
// 			"kube_controller_manager_extra_env": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kube-controller-managerの追加環境変数。",
// 			},
// 			"kube_proxy_extra_env": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kube-proxyの追加環境変数。",
// 			},
// 			"etcd_extra_env": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "etcdの追加環境変数。",
// 			},
// 			"cloud_controller_manager_extra_env": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "cloud-controller-managerの追加環境変数。",
// 			},

// 			////////////////////////// Cloud Provider //////////////////////////
// 			"cloud_provider_name": schema.StringAttribute{
// 				Optional:    true,
// 				Description: "クラウドプロバイダー名。",
// 			},
// 			"cloud_provider_config": schema.StringAttribute{
// 				Optional:    true,
// 				Description: "クラウドプロバイダー設定ファイルのパス。",
// 			},

// 			////////////////////////// Security //////////////////////////
// 			"profile": schema.StringAttribute{
// 				Optional:    true,
// 				Description: "システム設定を検証するベンチマーク (cis, cis-1.23)。",
// 			},
// 			"audit_policy_file": schema.StringAttribute{
// 				Optional:    true,
// 				Description: "監査ポリシー設定を定義するファイルのパス。",
// 			},
// 			"pod_security_admission_config_file": schema.StringAttribute{
// 				Optional:    true,
// 				Description: "Pod Security Admission設定を定義するファイルのパス。",
// 			},
// 			"secrets_encryption_provider": schema.StringAttribute{
// 				Optional:    true,
// 				Description: "使用する暗号化プロバイダー。",
// 			},

// 			////////////////////////// Node / Agent //////////////////////////
// 			"node_name": schema.StringAttribute{
// 				Optional:    true,
// 				Description: "ノード名。",
// 			},
// 			"with_node_id": schema.BoolAttribute{
// 				Optional:    true,
// 				Description: "ノード名にIDを追加します。",
// 			},
// 			"node_label": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kubeletに登録するラベルのセット。",
// 			},
// 			"node_taint": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kubeletに登録するテイントのセット。",
// 			},
// 			"image_credential_provider_bin_dir": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "クレデンシャルプロバイダープラグインバイナリが配置されているディレクトリのパス。",
// 				Default: stringdefault.StaticString("/var/lib/rancher/credentialprovider/bin"),
// 			},
// 			"image_credential_provider_config": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "クレデンシャルプロバイダープラグイン設定ファイルのパス。",
// 				Default: stringdefault.StaticString("/var/lib/rancher/credentialprovider/config.yaml"),
// 			},
// 			"protect_kernel_defaults": schema.BoolAttribute{
// 				Optional:    true,
// 				Description: "カーネルチューニングの動作。設定されている場合、カーネルの調整値がkubeletのデフォルトと異なるとエラーになります。",
// 			},
// 			"selinux": schema.BoolAttribute{
// 				Optional:    true,
// 				Description: "containerdでSELinuxを有効にします。",
// 			},
// 			"lb_server_port": schema.Int64Attribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "スーパーバイザークライアントロードバランサーのローカルポート (デフォルト: 6444)。",
// 				Default: int64default.StaticInt64(6444),
// 			},
// 			"embedded_registry": schema.BoolAttribute{
// 				Optional:    true,
// 				Description: "組み込み分散コンテナレジストリを有効にします。",
// 			},
// 			"enable_pprof": schema.BoolAttribute{
// 				Optional:    true,
// 				Description: "スーパーバイザーポートでpprofエンドポイントを有効にします。",
// 			},
// 			"kubelet_path": schema.StringAttribute{
// 				Optional:    true,
// 				Description: "kubeletバイナリのパスを上書きします。",
// 			},

// 			////////////////////////// Container Runtime //////////////////////////
// 			"container_runtime_endpoint": schema.StringAttribute{
// 				Optional:    true,
// 				Description: "組み込みcontainerdを無効にし、指定されたパスのCRIソケットを使用します。",
// 			},
// 			"default_runtime": schema.StringAttribute{
// 				Optional:    true,
// 				Description: "containerdのデフォルトランタイムを設定します。",
// 			},
// 			"snapshotter": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "デフォルトのcontainerdスナップショッターを上書きします (デフォルト: overlayfs)。",
// 				Default: stringdefault.StaticString("overlayfs"),
// 			},
// 			"private_registry": schema.StringAttribute{
// 				Optional:    true,
// 				Computed:    true,
// 				Description: "プライベートレジストリ設定ファイル (デフォルト: /etc/rancher/rke2/registries.yaml)。",
// 				Default: stringdefault.StaticString("/etc/rancher/rke2/registries.yaml"),
// 			},
// 			"system_default_registry": schema.StringAttribute{
// 				Optional:    true,
// 				Description: "すべてのシステムイメージに使用されるプライベートレジストリ。",
// 			},
// 			"disable_default_registry_endpoint": schema.BoolAttribute{
// 				Optional:    true,
// 				Description: "レジストリにミラーが設定されている場合、containerdのフォールバックデフォルトレジストリエンドポイントを無効にします。",
// 			},

// 			////////////////////////// Agent Networking //////////////////////////
// 			"node_ip": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "ノードのアドバタイズ用IPv4/IPv6アドレス。",
// 			},
// 			"node_external_ip": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "ノードのアドバタイズ用外部IPv4/IPv6アドレス。",
// 			},
// 			"resolv_conf": schema.StringAttribute{
// 				Optional:    true,
// 				Description: "Kubeletのresolv.confファイル。",
// 			},

// 			////////////////////////// Agent Flags //////////////////////////
// 			"kubelet_arg": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kubeletプロセスのカスタマイズフラグ。",
// 			},
// 			"kube_proxy_arg": schema.ListAttribute{
// 				ElementType: types.StringType,
// 				Optional:    true,
// 				Description: "kube-proxyプロセスのカスタマイズフラグ。",
// 			},

// 		},
// 	}
// }

// // Create creates the resource and sets the initial Terraform state.
// func (r *rke2ServerNodeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
// 	var plan rke2ServerNodeResourceModel
// 	diags := req.Plan.Get(ctx, &plan)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	var items []
// }

// // Read refreshes the Terraform state with the latest data.
// func (r *rke2ServerNodeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
// }

// // Update updates the resource and sets the updated Terraform state on success.
// func (r *rke2ServerNodeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
// }

// // Delete deletes the resource and removes the Terraform state on success.
// func (r *rke2ServerNodeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
// }

// func (r *rke2ServerNodeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
// 	if req.ProviderData == nil {
// 		return
// 	}

// 	client, ok := req.ProviderData.(*ssh.Client)

// 	if !ok {
// 		resp.Diagnostics.AddError(
// 			"Unexpected Data Source Configure Type",
// 			fmt.Sprintf("Expected *ssh.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
// 		)
// 		return
// 	}
// 	r.client = client
// }