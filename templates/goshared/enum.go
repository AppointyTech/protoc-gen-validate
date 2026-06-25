package goshared

const enumTpl = `
		{{ $f := .Field }}{{ $r := .Rules }}
		{{ template "const" . }}
		{{ template "in" . }}

		{{ if $r.GetDefinedOnly }}
			if _, ok := {{ trimStar (typ $f).Element }}_name[int32({{ accessor . }})]; !ok {
				return {{ err . "value must be one of the defined enum values" }}
			}
		{{ end }}
`
