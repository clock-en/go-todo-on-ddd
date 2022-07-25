# TODO API by golang

## 使用技術
- Golang (v1.18.2)
  - GraphQL (gqlgen v0.17.12)
- MySQL (v8.0)
- Docker

## Getting Started

### 起動前の初期設定 (Git Clone後にまずやること)

1. `.devcontainer/.env.example` を複製して `.devcontainer/.env.development` を作成する (※ リネームではなく必ず複製してください)
2. 作成した `.env.development` の `MYSQL_ROOT_PASSWORD` と `MYSQL_PASSWORD` に任意のパスワードを設定する (アプリケーションがMySQLにアクセスするために必要となります)

### Dockerの起動 (VScode devcontainer を使用)

1. VScodeに Remote Development をインストールする (ms-vscode-remote.vscode-remote-extensionpack)
2. 「起動前の初期設定 (Git Clone後にまずやること)」に記載されている手順で設定を行う
3. cloneしたディレクトリに移動してVScodeで開く
4. VScode画面左下にある 緑色の><マークをクリック
5. 上部に表示されたパレットから Open Folder in Container... を選択
6. 新しく立ち上がったVScodeのローディングが完了するのを待つ (以降はVScode上の操作が Docker コンテナ内の操作となります)

### アプリケーションの起動
1. 起動後のVScodeターミナル上で `go run main.go` を実行してください
2. ブラウザで http://localhost:8080 にアクセス

## アーキテクチャ

### ディレクトリ構造
本アプリケーションはクリーンアーキテクチャを採用しています。
下記方針でディレクトリは構成されています。

- domain : Enterprise Business Rules
  - entity : DDD におけるエンティティを実装
  - vo (Value Object) : DDD における値オブジェクトを実装
  - service (Domain Service) : DDD におけるドメインサービスを実装
- usecase : Application Business Rules
  - 各ユースケースファイル : 機能別のユースケースを実装
  - dto : 各ユースケースの InputData と OutputData を定義
- adapter : Interface Adapters
  - controller : usecase と graph の中間処理を実装
  - repository : usecase と dao の中間処理を実装
- infrastructure : Frameworks & Drivers
  - graph (GraphQL Server by gqlgen) : Webインフラとの接続処理を実装
  - dao (Data Access Object) : DBインフラとの接続処理を実装

基本的に依存方向は 同じレイヤー内 または 下位レイヤー -> 上位レイヤーとなるように実装をする。