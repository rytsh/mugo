{
    {{$dir := readDir .}}
    {{- range $i, $v := $dir }}
    "{{ $v.Name }}": {
        "isDir": "{{ $v.IsDir }}",
        "size": "{{ $v.Size | uint64 | bytes }}",
        "modTime": "{{ rfc3339 $v.ModTime }}"
        {{- if and $v.IsDir (not (hasPrefix "." $v.Name)) }}
            {{execTemplate "readSeparate.tpl" (printf "%s/%s" $ $v.Name) | mustFromJson | mustToPrettyJson | saveFile (printf "output/%s/%s/info.json" $ $v.Name) | nothing }}
        {{ end }}
    }
    {{- if lt (add $i 1) (len $dir) }},{{ end }}
    {{- end }}
}
