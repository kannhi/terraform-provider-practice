|    | Flag     | Description          | Default                       | Environment Variable   |
|---:|:---------|:---------------------|:------------------------------|:-----------------------|
|  0 | config   | Path to config file  | /etc/rancher/rke2/config.yaml | RKE2_CONFIG_FILE       |
|  1 | debug    | Turn on debug logs   | nan                           | RKE2_DEBUG             |
|  2 | data_dir | Folder to hold state | "/var/lib/rancher/rke2"       | nan                    |


|    | Flag              | Description                                                                                                                                                                  | Default                  |
|---:|:------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:-------------------------|
|  0 | bind_address      | rke2 bind address                                                                                                                                                            | 0.0.0.0                  |
|  1 | advertise_address | IPv4/IPv6 address that apiserver uses to advertise to members of the cluster                                                                                                 | node-external-ip/node-ip |
|  2 | tls_san           | Add additional hostnames or IPv4/IPv6 addresses as Subject Alternative Names on the server TLS cert                                                                          | nan                      |
|  3 | tls_san_security  | Protect the server TLS cert by refusing to add Subject Alternative Names not associated with the kubernetes apiserver service, server nodes, or values of the tls-san option | true                     |


|    | Flag                    | Description                                                                                                                           | Default         | Environment Variable   |
|---:|:------------------------|:--------------------------------------------------------------------------------------------------------------------------------------|:----------------|:-----------------------|
|  0 | cluster_cidr            | IPv4/IPv6 network CIDRs to use for pod IPs                                                                                            | 10.42.0.0/16    | nan                    |
|  1 | service_cidr            | IPv4/IPv6 network CIDRs to use for service IPs                                                                                        | 10.43.0.0/16    | nan                    |
|  2 | service_node_port_range | Port range to reserve for services with NodePort visibility                                                                           | "30000-32767"   | nan                    |
|  3 | cluster_dns             | IPv4 Cluster IP for coredns service. Should be in your service-cidr range                                                             | 10.43.0.10      | nan                    |
|  4 | cluster_domain          | Cluster Domain                                                                                                                        | "cluster.local" | nan                    |
|  5 | egress_selector_mode    | One of 'agent', 'cluster', 'pod', 'disabled'                                                                                          | "agent"         | nan                    |
|  6 | servicelb_namespace     | Namespace of the pods for the servicelb component                                                                                     | "kube-system"   | nan                    |
|  7 | cni                     | CNI Plugins to deploy, one of none, calico, canal, cilium; optionally with multus as the first value to enable the multus meta-plugin | canal           | RKE2_CNI               |


|    | Flag                  | Description                                    | Environment Variable   |
|---:|:----------------------|:-----------------------------------------------|:-----------------------|
|  0 | write_kubeconfig      | Write kubeconfig for admin client to this file | RKE2_KUBECONFIG_OUTPUT |
|  1 | write_kubeconfig_mode | Write kubeconfig with this mode                | RKE2_KUBECONFIG_MODE   |


|    | Flag           | Description                        |
|---:|:---------------|:-----------------------------------|
|  0 | helm_job_image | Default image to use for helm jobs |


|    | Flag             | Description                                                       | Environment Variable   |
|---:|:-----------------|:------------------------------------------------------------------|:-----------------------|
|  0 | token            | Shared secret used to join a server or agent to a cluster         | RKE2_TOKEN             |
|  1 | token_file       | File containing the token                                         | RKE2_TOKEN_FILE        |
|  2 | agent_token      | Shared secret used to join agents to the cluster, but not servers | RKE2_AGENT_TOKEN       |
|  3 | agent_token_file | File containing the agent secret                                  | RKE2_AGENT_TOKEN_FILE  |
|  4 | server           | Server to connect to, used to join a cluster                      | RKE2_URL               |
|  5 | cluster_reset    | Forget all peers and become sole member of a new cluster          | RKE2_CLUSTER_RESET     |


|    | Flag                        | Description                                                          | Default                         | Environment Variable   |
|---:|:----------------------------|:---------------------------------------------------------------------|:--------------------------------|:-----------------------|
|  0 | cluster_reset_restore_path  | Path to snapshot file to be restored                                 | nan                             | nan                    |
|  1 | etcd_expose_metrics         | Expose etcd metrics to client interface.                             | false                           | nan                    |
|  2 | etcd_disable_snapshots      | Disable automatic etcd snapshots                                     | nan                             | nan                    |
|  3 | etcd_snapshot_name          | Set the base name of etcd snapshots                                  | etcd-snapshot-<unix-timestamp>) | nan                    |
|  4 | etcd_snapshot_schedule_cron | Snapshot interval time in cron spec. eg. every 5 hours '0 */5 * * *' | "0 */12 * * *"                  | nan                    |
|  5 | etcd_snapshot_retention     | Number of snapshots to retain                                        | 5                               | nan                    |
|  6 | etcd_snapshot_dir           | Directory to save db snapshots.                                      | ${data-dir}/db/snapshots        | nan                    |
|  7 | etcd_snapshot_compress      | Compress etcd snapshot                                               | nan                             | nan                    |
|  8 | etcd_s3                     | Enable backup to S3                                                  | nan                             | nan                    |
|  9 | etcd_s3_endpoint            | S3 endpoint url                                                      | "s3.amazonaws.com"              | nan                    |
| 10 | etcd_s3_endpoint_ca         | S3 custom CA cert to connect to S3 endpoint                          | nan                             | nan                    |
| 11 | etcd_s3_skip_ssl_verify     | Disables S3 SSL certificate validation                               | nan                             | nan                    |
| 12 | etcd_s3_access_key          | S3 access key                                                        | nan                             | AWS_ACCESS_KEY_ID      |
| 13 | etcd_s3_secret_key          | S3 secret key                                                        | nan                             | AWS_SECRET_ACCESS_KEY  |
| 14 | etcd_s3_bucket              | S3 bucket name                                                       | nan                             | nan                    |
| 15 | etcd_s3_region              | S3 region / bucket location (optional)                               | "us-east-1"                     | nan                    |
| 16 | etcd_s3_folder              | S3 folder                                                            | nan                             | nan                    |
| 17 | etcd_s3_insecure            | Disables S3 over HTTPS                                               | nan                             | nan                    |
| 18 | etcd_s3_timeout             | S3 timeout                                                           | 5m0s                            | nan                    |


|    | Flag              | Description                         |   Default |
|---:|:------------------|:------------------------------------|----------:|
|  0 | etcd_s3_retention | Number of snapshots in S3 to retain |         5 |


|    | Flag                              | Description                                               |
|---:|:----------------------------------|:----------------------------------------------------------|
|  0 | kube_apiserver_arg                | Customized flag for kube-apiserver process                |
|  1 | etcd_arg                          | Customized flag for etcd process                          |
|  2 | kube_controller_manager_arg       | Customized flag for kube-controller-manager process       |
|  3 | kube_scheduler_arg                | Customized flag for kube-scheduler process                |
|  4 | kube_cloud_controller_manager_arg | Customized flag for kube-cloud-controller-manager process |


|    | Flag                                 | Description                                                                                                                               | Environment Variable                      |
|---:|:-------------------------------------|:------------------------------------------------------------------------------------------------------------------------------------------|:------------------------------------------|
|  0 | disable                              | Do not deploy packaged components and delete any deployed components (valid items: rke2-coredns, rke2-ingress-nginx, rke2-metrics-server) | nan                                       |
|  1 | disable_scheduler                    | Disable Kubernetes default scheduler                                                                                                      | nan                                       |
|  2 | disable_cloud_controller             | Disable rke2 default cloud controller manager                                                                                             | nan                                       |
|  3 | disable_kube_proxy                   | Disable running kube-proxy                                                                                                                | nan                                       |
|  4 | enable_servicelb                     | Enable rke2 default cloud controller manager's service controller                                                                         | RKE2_ENABLE_SERVICELB                     |
|  5 | control_plane_resource_requests      | Control Plane resource requests                                                                                                           | RKE2_CONTROL_PLANE_RESOURCE_REQUESTS      |
|  6 | control_plane_resource_limits        | Control Plane resource limits                                                                                                             | RKE2_CONTROL_PLANE_RESOURCE_LIMITS        |
|  7 | control_plane_probe_configuration    | Control Plane Probe configuration                                                                                                         | RKE2_CONTROL_PLANE_PROBE_CONFIGURATION    |
|  8 | kube_apiserver_extra_mount           | kube-apiserver extra volume mounts                                                                                                        | RKE2_KUBE_APISERVER_EXTRA_MOUNT           |
|  9 | kube_scheduler_extra_mount           | kube-scheduler extra volume mounts                                                                                                        | RKE2_KUBE_SCHEDULER_EXTRA_MOUNT           |
| 10 | kube_controller_manager_extra_mount  | kube-controller-manager extra volume mounts                                                                                               | RKE2_KUBE_CONTROLLER_MANAGER_EXTRA_MOUNT  |
| 11 | kube_proxy_extra_mount               | kube-proxy extra volume mounts                                                                                                            | RKE2_KUBE_PROXY_EXTRA_MOUNT               |
| 12 | etcd_extra_mount                     | etcd extra volume mounts                                                                                                                  | RKE2_ETCD_EXTRA_MOUNT                     |
| 13 | cloud_controller_manager_extra_mount | cloud-controller-manager extra volume mounts                                                                                              | RKE2_CLOUD_CONTROLLER_MANAGER_EXTRA_MOUNT |
| 14 | kube_apiserver_extra_env             | kube-apiserver extra environment variables                                                                                                | RKE2_KUBE_APISERVER_EXTRA_ENV             |
| 15 | kube_scheduler_extra_env             | kube-scheduler extra environment variables                                                                                                | RKE2_KUBE_SCHEDULER_EXTRA_ENV             |
| 16 | kube_controller_manager_extra_env    | kube-controller-manager extra environment variables                                                                                       | RKE2_KUBE_CONTROLLER_MANAGER_EXTRA_ENV    |
| 17 | kube_proxy_extra_env                 | kube-proxy extra environment variables                                                                                                    | RKE2_KUBE_PROXY_EXTRA_ENV                 |
| 18 | etcd_extra_env                       | etcd extra environment variables                                                                                                          | RKE2_ETCD_EXTRA_ENV                       |
| 19 | cloud_controller_manager_extra_env   | cloud-controller-manager extra environment variables                                                                                      | RKE2_CLOUD_CONTROLLER_MANAGER_EXTRA_ENV   |
| 20 | ingress_controller                   | Ingress Controller to deploy one of, none, ingress-nginx, traefik                                                                         | nan                                       |


|    | Flag                           | Description                                                                   | Environment Variable                |
|---:|:-------------------------------|:------------------------------------------------------------------------------|:------------------------------------|
|  0 | kube_apiserver_image           | Override image to use for kube-apiserver                                      | RKE2_KUBE_APISERVER_IMAGE           |
|  1 | kube_controller_manager_image  | Override image to use for kube-controller-manager                             | RKE2_KUBE_CONTROLLER_MANAGER_IMAGE  |
|  2 | cloud_controller_manager_image | Override image to use for cloud-controller-manager                            | RKE2_CLOUD_CONTROLLER_MANAGER_IMAGE |
|  3 | kube_proxy_image               | Override image to use for kube-proxy                                          | RKE2_KUBE_PROXY_IMAGE               |
|  4 | kube_scheduler_image           | Override image to use for kube-scheduler                                      | RKE2_KUBE_SCHEDULER_IMAGE           |
|  5 | pause_image                    | Override image to use for pause                                               | RKE2_PAUSE_IMAGE                    |
|  6 | runtime_image                  | Override image to use for runtime binaries (containerd, kubectl, crictl, etc) | RKE2_RUNTIME_IMAGE                  |
|  7 | etcd_image                     | Override image to use for etcd                                                | RKE2_ETCD_IMAGE                     |


|    | Flag                  | Description                            | Environment Variable       |
|---:|:----------------------|:---------------------------------------|:---------------------------|
|  0 | cloud_provider_name   | Cloud provider name                    | RKE2_CLOUD_PROVIDER_NAME   |
|  1 | cloud_provider_config | Cloud provider configuration file path | RKE2_CLOUD_PROVIDER_CONFIG |


|    | Flag                               | Description                                                                                            | Environment Variable                    |
|---:|:-----------------------------------|:-------------------------------------------------------------------------------------------------------|:----------------------------------------|
|  0 | profile                            | Validate system configuration against the selected benchmark (valid items: cis, cis-1.23 (deprecated)) | RKE2_CIS_PROFILE                        |
|  1 | audit_policy_file                  | Path to the file that defines the audit policy configuration                                           | RKE2_AUDIT_POLICY_FILE                  |
|  2 | pod_security_admission_config_file | Path to the file that defines Pod Security Admission configuration                                     | RKE2_POD_SECURITY_ADMISSION_CONFIG_FILE |
|  3 | secrets_encryption_provider        | Encryption provider to use                                                                             | nan                                     |


|    | Flag              | Description                                                                         | Environment Variable   |
|---:|:------------------|:------------------------------------------------------------------------------------|:-----------------------|
|  0 | embedded_registry | Enable embedded distributed container registry; requires use of embedded containerd | nan                    |
|  1 | enable_pprof      | Enable pprof endpoint on supervisor port                                            | nan                    |
|  2 | kubelet_path      | Override kubelet binary path                                                        | RKE2_KUBELET_PATH      |


|    | Flag                              | Description                                                                                                                                                                                          | Default                                           | Environment Variable   |
|---:|:----------------------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:--------------------------------------------------|:-----------------------|
|  0 | node_name                         | Node name                                                                                                                                                                                            | nan                                               | RKE2_NODE_NAME         |
|  1 | with_node_id                      | Append id to node name                                                                                                                                                                               | nan                                               | nan                    |
|  2 | node_label                        | Registering and starting kubelet with set of labels                                                                                                                                                  | nan                                               | nan                    |
|  3 | node_taint                        | Registering kubelet with set of taints                                                                                                                                                               | nan                                               | nan                    |
|  4 | image_credential_provider_bin_dir | The path to the directory where credential provider plugin binaries are located                                                                                                                      | "/var/lib/rancher/credentialprovider/bin"         | nan                    |
|  5 | image_credential_provider_config  | The path to the credential provider plugin config file                                                                                                                                               | "/var/lib/rancher/credentialprovider/config.yaml" | nan                    |
|  6 | protect_kernel_defaults           | Kernel tuning behavior. If set, error if kernel tunables are different than kubelet defaults.                                                                                                        | nan                                               | nan                    |
|  7 | selinux                           | Enable SELinux in containerd                                                                                                                                                                         | nan                                               | RKE2_SELINUX           |
|  8 | lb_server_port                    | Local port for supervisor client load-balancer. If the supervisor and apiserver are not colocated an additional port 1 less than this port will also be used for the apiserver client load-balancer. | 6444                                              | RKE2_LB_SERVER_PORT    |


|    | Flag                       | Description                                                                                                                    | Default                             | Environment Variable         |
|---:|:---------------------------|:-------------------------------------------------------------------------------------------------------------------------------|:------------------------------------|:-----------------------------|
|  0 | container_runtime_endpoint | Disable embedded containerd and use the CRI socket at the given path; when used with --docker this sets the docker socket path | nan                                 | nan                          |
|  1 | default_runtime            | Set the default runtime in containerd                                                                                          | nan                                 | nan                          |
|  2 | snapshotter                | Override default containerd snapshotter                                                                                        | "overlayfs"                         | nan                          |
|  3 | private_registry           | Private registry configuration file                                                                                            | "/etc/rancher/rke2/registries.yaml" | nan                          |
|  4 | system_default_registry    | Private registry to be used for all system images                                                                              | nan                                 | RKE2_SYSTEM_DEFAULT_REGISTRY |


|    | Flag                              | Description                                                                                            |
|---:|:----------------------------------|:-------------------------------------------------------------------------------------------------------|
|  0 | disable_default_registry_endpoint | Disables containerd's fallback default registry endpoint when a mirror is configured for that registry |


|    | Flag             | Description                                           | Environment Variable   |
|---:|:-----------------|:------------------------------------------------------|:-----------------------|
|  0 | node_ip          | IPv4/IPv6 addresses to advertise for node             | nan                    |
|  1 | node_external_ip | IPv4/IPv6 external IP addresses to advertise for node | nan                    |
|  2 | resolv_conf      | Kubelet resolv.conf file                              | RKE2_RESOLV_CONF       |


|    | Flag           | Description                            |
|---:|:---------------|:---------------------------------------|
|  0 | kubelet_arg    | Customized flag for kubelet process    |
|  1 | kube_proxy_arg | Customized flag for kube-proxy process |

---

Disable
Disable_scheduler
Disable_cloud_controller
Disable_kube_proxy
Enable_servicelb
Control_plane_resource_requests
Control_plane_resource_limits
Control_plane_probe_configuration
Kube_apiserver_extra_mount
Kube_scheduler_extra_mount
Kube_controller_manager_extra_mount
Kube_proxy_extra_mount
Etcd_extra_mount
Cloud_controller_manager_extra_mount
Kube_apiserver_extra_env
Kube_scheduler_extra_env
Kube_controller_manager_extra_env
Kube_proxy_extra_env
Etcd_extra_env
Cloud_controller_manager_extra_env
Ingress_controller
Kube_apiserver_image
Kube_controller_manager_image
Cloud_controller_manager_image
Kube_proxy_image
Kube_scheduler_image
Pause_image
Runtime_image
Etcd_image
Cloud_provider_name
Cloud_provider_config
Profile
Audit_policy_file
Pod_security_admission_config_file
Secrets_encryption_provider
Embedded_registry
Enable_pprof
Kubelet_path
Node_name
With_node_id
Node_label
Node_taint
Image_credential_provider_bin_dir
Image_credential_provider_config
Protect_kernel_defaults
Selinux
Lb_server_port
Container_runtime_endpoint
Default_runtime
Snapshotter
Private_registry
System_default_registry
Disable_default_registry_endpoint
Node_ip
Node_external_ip
Resolv_conf
Kubelet_arg
Kube_proxy_arg
