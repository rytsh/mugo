{{- execTemplate "readSeparate.tpl" . | mustFromJson | mustToPrettyJson | saveFile "output/info.json" | nothing -}}
