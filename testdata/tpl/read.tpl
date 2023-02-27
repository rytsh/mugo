{
    {{$dir := os.ReadDir .}}
    {{- range $i, $v := $dir }}
    "{{ $v.Name }}": {
        "isDir": "{{ $v.IsDir }}",
        "size": "{{ $v.Size | cast.ToUint64 | humanize.Bytes }}",
        "modTime": "{{ time.Format time.RFC3339 $v.ModTime }}"
        {{- if and $v.IsDir (not (hasPrefix "." $v.Name)) }},
        "files": {{template "read.tpl" (printf "%s/%s" $ $v.Name)}}
        {{ end }}
    }
    {{- if lt (add $i 1) (len $dir) }},{{ end }}
    {{- end }}
}
