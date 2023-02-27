{{- $body := (file.Read "testdata/data/index.md" | codec.Markdown | codec.ByteToString) -}}
{{- define "layout" -}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
{{ . | indent 4 }}
</body>
</html>
{{- end -}}

{{- execTemplate "layout" $body }}
