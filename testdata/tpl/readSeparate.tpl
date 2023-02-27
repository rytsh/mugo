{
    {{$dir := os.ReadDir .}}
    {{- range $i, $v := $dir }}
    "{{ $v.Name }}": {
        "isDir": "{{ $v.IsDir }}",
        "size": "{{ $v.Size | cast.ToUint64 | humanize.Bytes }}",
        "modTime": "{{ time.Format time.RFC3339 $v.ModTime }}"
        {{- if and $v.IsDir (not (hasPrefix "." $v.Name)) }}
            {{execTemplate "readSeparate.tpl" (printf "%s/%s" $ $v.Name) | mustFromJson | mustToPrettyJson | codec.StringToByte | file.Save (printf "output/%s/%s/info.json" $ $v.Name) | nothing }}
        {{ end }}
    }
    {{- if lt (add $i 1) (len $dir) }},{{ end }}
    {{- end }}
}
