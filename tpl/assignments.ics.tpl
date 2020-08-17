BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//Chris Sexton//NONSGML v1.0//EN
{{- range $key, $value := .Assignments}}
BEGIN:VEVENT
UID:{{$key}}-{{projectSlug}}-cwsexton@iu.edu
DTSTAMP:{{getDTStamp}}
DTSTART:{{getDTStart $value.Due $value.DueDate}}
SUMMARY:{{$value.Title}} Due
END:VEVENT
{{- end}}
END:VCALENDAR
