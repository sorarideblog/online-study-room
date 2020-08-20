# デプロイの仕方
deploy.shを実行して。内容はトリガーの種類ごとに書き換えて。

# ハマったこと
GoLandの不具合なのかわからないけど、プロジェクト内外に関わらず（GOPATH以下はやったことないが）、
`go get`しても読み込んでくれないし、go mod init firebase.google.com/go でモジュールインストールすると、
firebase.google.com/goそのもののモジュールはself importエラーとなってしまう
(自動追加でインストールされるcloud.google.com/goなどはちゃんと読み込めるのに)。

原因としては、go mod のファイル名は適当でいいのに、インストールしたいパッケージの名前で設定してしまったこと。
よって、名前空間がかぶってしまって`self import`エラーが出たというわけ。

## 解決法
プロジェクト内で`go mod init test.module`を実行する。
go.modファイルの名前（ここではtest.module）は適当で良いが、必ずドットで区切ること。ここがハマった原因。