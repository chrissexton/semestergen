- {{.Project}} @parallel(true) @autodone(false) @context(Place : IUS) @tags(Place : IUS) @defer(2019-12-20 00:00)
	- Assignments @parallel(false) @autodone(false) @context(Place : IUS) @tags(Place : IUS)
{{- range $key, $value := .Assignments}}
        - {{$value.Title}} @context(Place : IUS) @tags(Place : IUS) @due({{getDate $value.Assigned}} {{dueTime}})
{{- end}}
	- Grading @parallel(false) @autodone(false) @context(Place : IUS) @tags(Place : IUS)
{{- range $key, $value := .Assignments}}
        - {{$value.Title}} @context(Place : IUS) @tags(Place : IUS) @defer({{getDate $value.Due}} {{dueTime}})
{{- end}}
	- Lectures @parallel(false) @autodone(false) @context(Place : IUS) @tags(Place : IUS)
{{- range $key, $value := .Days}}
		- {{$value.Title}} @context(Place : IUS) @tags(Place : IUS) @due({{getDate $value.Num}} {{dueTime}})
{{- end}}
