# terraform-provider-practice

## 目的

* rke2のリソースを定義するためのプロバイダーを実装する
* (実際はgoを書くための練習の気持ちもある)

## 構成

```
├── design.md
├── go.mod
├── go.sum
├── internal
│   └── provider
│       ├── datasources
│       │   └── rke2_server_node.go
│       ├── provider.go
│       └── resources
│           ├── rke2_server_node.go
│           └── test_exec.go
├── main.go
└── terraform-provider-rke2
```

* design.md
    * rke2の設定値を書き出したもの
* internal/provider
    * rke2のプロバイダーの実装
    * provider.go
        * rke2を導入したいサーバーにsshで接続するためのクライアントを作成する
    * datasources
        * rke2_server_node.go(未実装)
    * resources
        * rke2_server_node.go(未実装)
        * test_exec.go
            * sshしたサーバーでコマンドを実行する流れを確認する
* main.go
    * メインファイル
* terraform-provider-rke2
    * バイナリファイル

## 実行手法

1. このリポジトリをクローン
2. 以下コマンドを順に実行

    ```bash
    go mod tidy
    go build -o .
    ```

3. `~/.terraformrc`の書き換え

    ```hcl
    provider_installation { 

    dev_overrides { 
        "local/xxx/rke2" = "<PATH>/terraform-provider-practice" 
    } 

    direct { } 
    }
    ```

4. 以下のような設定でterraformを実行

    ```hcl
    terraform {
    required_providers {
        rke2 = {
        source = "local/xxx/rke2"
        }
    }
    }

    provider "rke2" {
    user       = "user"
    privatekey = file("~/.ssh/id_rsa")
    host       = "raspberrypi.local"
    timeout    = "10s"
    }

    resource "rke2_test_exec" "test_exec" {
    file_name = "~/test_file_2"
    }

    output "rke2_test_exec_output" {
    value = rke2_test_exec.test_exec.file_name
    }
    ```