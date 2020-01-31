{{- range $key, $value := .Assignments}}
{{- range $value.Links}}
:{{.Slug}}: link:{{.URL}}
{{- end}}
{{- end}}

[%header,format=psv,cols="^20h,^20h,<60d"]
|===
| Given  | Due Date    | Assignment
{{range $key, $value := .Assignments}}
|  {{getDate $value.Assigned}}   | {{getDate $value.Due}}   | {{$value.Title}}: {{range $idx, $el := $value.Links}}{{if $idx}}; {{end}}{{$el.Slug | printf "{%s}"}}[{{$el.Title}}]{{end}}
{{end}}
|===

