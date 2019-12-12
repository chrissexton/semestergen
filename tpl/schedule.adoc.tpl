{{- range $key, $value := .Days}}
{{- range $value.Links}}
:{{.Slug}}: link:{{.URL}}
{{- end}}
{{- end}}

[%header,format=psv,cols="^5h,47d,47a"]
|===
| Day | Topics                                                          | Notes
{{range $key, $value := .Days}}
|  {{$value.Num}}   | {{$value.Title}}   | {{range $idx, $l := $value.Links}}{{if $idx}}; {{end}}{{$l.Slug | printf "{%s}"}}[{{$l.Title}}]{{end}}
{{end}}
|===
