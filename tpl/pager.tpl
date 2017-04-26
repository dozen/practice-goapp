{{define "pager"}}
<div class="pager">
{{if .Prev}}<a href="{{ .Prev }}">prev</a>&nbsp;{{end}}
{{if .Next}}<a href="{{ .Next }}">next</a>{{end}}
</div>
{{end}}