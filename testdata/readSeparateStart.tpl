{{- execTemplate "readSeparate.tpl" . | mustFromJson | mustToPrettyJson | codec.StringToByte | file.Save "output/info.json" | nothing -}}
