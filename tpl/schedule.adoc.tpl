:nofooter:
{{- range $key, $value := .Days}}
{{- range $value.Links}}
:{{.Slug}}: link:{{.URL}}
{{- end}}
{{- end}}

[%header,format=psv,cols="^5h,47d,47a"]
|===
| Day | Topics                                                          | Notes
{{range $key, $value := .Days}}
// {{$value.Date}}
|  {{$value.Num}}   a| {{$value.Title}}   a| {{range $idx, $l := $value.Links}}* {{$l.Slug | printf "{%s}"}}[{{$l.Title}}]
{{end}}
{{end}}
|===
