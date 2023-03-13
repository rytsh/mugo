{{ define "readSeparate" }}
{
    {{$dir := os.ReadDir (print .dir)}}
    {{- range $i, $v := $dir }}
    "{{ $v.Name }}": {
        "url": "{{$.url}}/{{ $v.Name }}{{ ternary "/info.json" "" $v.IsDir }}",
        "isDir": {{ $v.IsDir }},
        "size": "{{ $v.Size | cast.ToUint64 | humanize.Bytes }}",
        "modTime": "{{ time.Format time.RFC3339 $v.ModTime }}"
        {{- if and $v.IsDir (not (hasPrefix "." $v.Name)) }}
            {{execTemplate "readSeparate" (printf `{"dir": "%[1]s/%[3]s", "url": "%[2]s/%[3]s"}` $.dir $.url $v.Name | codec.StringToByte | codec.JsonDecode) | mustFromJson | mustToPrettyJson | codec.StringToByte | file.Save (printf "output/%s/%s/info.json" $.dir $v.Name) | nothing }}
        {{ end }}
    }
    {{- if lt (add $i 1) (len $dir) }},{{ end }}
    {{- end }}
}
{{ end }}

{{- execTemplate "readSeparate" . | mustFromJson | mustToPrettyJson | codec.StringToByte | file.Save "output/info.json" | nothing -}}
