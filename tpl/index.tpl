{{define "index"}}
{{template "header" .}}
{{ range .Themes }}
        <div class="isu-submit">

          <div class="isu-form">
            {{ joke_count .id }} 個のジョーク
          </div>

          <div class="isu-post-image">
            <a href="/joke/new?id={{.id}}"><img class="isu-image" src="/uploads/{{.image}}"></a>
          </div>
          <div class="isu-form">
            <img src="/img/photo_16x16.png"><span title="photo by">{{.iu_account}}</span>
            <img src="/img/theme_16x16.png"><span title="theme by">{{.tu_account}}</span>
          </div>
          {{ if .content }}
            <div class="joke-content">{{ .content }}</div>
          {{ end }}
        </div>

{{ end }}
{{ template "pager" . }}
{{end}}