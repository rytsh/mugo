# Intro

Go template is a powerful tool to generate text output. It is used in many places, such as Kubernetes, Helm, and so on.

To see the details of Go template, please refer to the official document: [https://pkg.go.dev/text/template](https://pkg.go.dev/text/template).

## Use examples with mugo

With `mugo`, values can be passed directly on the command line to render simple templates.

```sh
mugo -s -d '{"name": "mugo"}' - <<< "{{.name}}"
```

Alternatively, enter a template interactively and finish standard input with `Ctrl+D`.

```sh
mugo -d '{"name": "mugo"}' -
```

Template files are a better option for more complex content.

This is a `values.yaml` file:

```yaml
name: mugo
```

This is the template file `name.tpl`:

```tpl
{{ .name }}
```

Render it with `mugo`:

```sh
$ mugo -s -d @values.yaml -t @name.tpl
mugo
```

Log output is written to standard error and rendered template output is written to standard output.

Use shell redirection or `--output` to write the result to a file.

```sh
mugo -d @values.yaml name.tpl -o name.txt
```

```sh
mugo -d @values.yaml name.tpl > name.txt
```
