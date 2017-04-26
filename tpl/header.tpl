{{define "header"}}
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta content="width=device-width,initial-scale=1.0,user-scalable=no" name="viewport" />
    <title>Joker</title>
    <link href="//fonts.googleapis.com/css?family=Lobster" rel="stylesheet">
    <link href="/css/style.css" media="screen" rel="stylesheet" type="text/css">
  </head>
  <body>
    <div class="container">
  <div class="header">
  <div class="isu-title">
    <h1><a href="/">Joker</a></h1>
  </div>
  <ul class="isu-header-menu">
    {{ if .Me }}
    <li><a href="/login">ログイン</a></li>
    <li><a href="/signup">ユーザ登録</a></li>
    {{ else }}
    <li><span class="isu-account-name">{{.Me.account}}</span>さん</li>
    <li><a href="/logout">ログアウト</a></li>
    {{ end }}
  </ul>
  <div class="header__column">
    <a href="/theme/new">テーマ投稿</a>
  </div>
</div>
{{end}}