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

These events are in a public calendar you may add to your own calendar software: https://p69-caldav.icloud.com/published/2/AAAAAAAAAAAAAAAAAAAAANFRXqg59DGi6z91QTJ-xG1Ope_sr0CiqvA6uMlifzjd3vdL7pulndzvh4lIbs7jcjVkXXeLS7hOJi_yt4J96Y8[Published iCal Link]
