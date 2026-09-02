# Core Functionality

Before using the examples, it helps to understand comments and whitespace handling.

```tpl
{{/* a comment */}}
{{- /* a comment with white space trimmed from preceding and following text */ -}}
	A comment; discarded. May contain newlines.
	Comments do not nest and must start and end at the
	delimiters, as shown here.
```

Minus sign (-) is a special character that represents white space trimming.

```tpl
Merhaba

 	{{- " " -}}
  	       dünya!
```

Output:

```txt
Merhaba dünya!
```

## Variables

Templates can define variables with the `$variableName := value` syntax.

```tpl
{{ $x := "Merhaba" -}}
{{ $x }} dünya!
```

Output:

```txt
Merhaba dünya!
```


## Range

If you have a list of items, you can iterate over them using the `range` function.

```yaml
list:
- "item-1"
- "item-2"
```

```tpl
Range of list
{{ range $index, $element := .list -}}
{{ $index }}: {{ $element }}
{{ end }}
```

Output:

```txt
Range of list
0: item-1
1: item-2

```

When ranging over a map, `$index` is the key and `$element` is the value.

The Sprig `until` function can generate an integer slice for a numeric range. Sprig functions are enabled by default.

```tpl
Count up to 5
{{ range $index, $element := until 5 -}}
{{ $index }}: {{ $element }}
{{ end }}
```

Output:

```txt
Count up to 5
0: 0
1: 1
2: 2
3: 3
4: 4

```

## If statement

The comparison functions `eq`, `ne`, `lt`, `le`, `gt`, and `ge` return booleans. Their arguments must have compatible types.

```yaml
result: 10
```

```tpl
{{if eq .result 10 -}}
Result is 10
{{- else if eq .result 0 -}}
Result is 0
{{- else -}}
unknown
{{- end}}
```

Output:

```txt
Result is 10
```

## Index

We can use `index` function to get an item from a list.

```yaml
list:
 - "item-1"
 - "item-2"
```

```tpl
First item of list: {{ index .list 0 }}
```

Output:

```txt
First item of list: item-1
```

## Define

Use `define` to create a named template and `template` to execute it with data.

```yaml
name: mugo
```

```tpl
{{- define "hello" -}}
Hello {{ .name }}
{{- end -}}

{{ template "hello" . }}
```

Output:

```txt
Hello mugo
```

## With

Use `with` to limit the scope of a value inside a template.

If `name` value is empty, it will not print anything.

```tpl
{{ with .name -}}
Merhaba {{ . }}
{{- end }}
```

If we want to reach outer scope, we can use `$.` syntax.

```tpl
{{ with .name -}}
Item {{ . }} value is {{ $.value }}
{{- end }}
```

<hr>

# Examples

Some useful examples with templates.

## Sum of variables

`addf` is provided by Sprig and adds floating-point values.

```yaml
values:
 - 6.1
 - 5.81
 - 7.9
```

```tpl
{{ $sum := 0 -}}
{{ range $index, $element := .values -}}
{{ $sum = addf $sum $element -}}
{{ end -}}
Total value {{ $sum }}
```

Output:

```txt
Total value 19.81
```
