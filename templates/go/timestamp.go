package golang

const timestampTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if or $r.Lt $r.Lte $r.Gt $r.Gte $r.LtNow $r.GtNow $r.Within $r.Const }}
		if t := {{ accessor . }}; t != nil {
			if err := t.CheckValid(); err != nil { return {{ errCause . "err" "value is not a valid timestamp" }} }
			ts := t.AsTime()

			{{ template "timestampcmp" . }}
		}
	{{ end }}
`
