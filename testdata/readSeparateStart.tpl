{{- execTemplate "readSeparate.tpl" . | mustFromJson | mustToPrettyJson | stringToByte | saveFile "output/info.json" | nothing -}}
