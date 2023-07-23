# Intro

Go template is a powerful tool to generate text output. It is used in many places, such as Kubernetes, Helm, and so on.

To see the details of Go template, please refer to the official document: [https://pkg.go.dev/text/template](https://pkg.go.dev/text/template).

## Use examples with mugo

With `mugo` we can use directly giveing values on commandline to run simple templates.

```sh
mugo -s -d '{"name": "mugo"}' - <<< "{{.name}}"
```

Or we write template and stop stdin with `ctrl+d`.

```sh
mugo -d '{"name": "mugo"}' -
```

But for more complex templates files are better option for now.

This is a `values.yaml` file

```yaml
name: mugo
```

And this si our template file `name.tpl`

```tpl
{{ .name }}
```

To run with `mugo`

```sh
mugo -s -d @values.yaml name.tpl
#mugo
```

Log output is move to stderr and template output is moved to stdout.

When we redirect to file, we use redirection or `-o` flag.

```sh
mugo -d @values.yaml name.tpl -o name.txt
```

```sh
mugo -d @values.yaml name.tpl > name.txt
```
