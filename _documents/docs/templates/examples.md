# Core Functionality

Before to check examples, we need to know about comments and whitespace.

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

Inside of template, we can use variables. Variables are defined by __$variableName := value__ syntax.

```tpl
{{ $x := "Merhaba" -}}
{{ $x }} dünya!
```

Output:

```txt
Merhaba dünya!
```


## For loop

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

If we want to for loop with a map, in that time $index is a key of map and $element is a value of map.

We can also use `range` to count up to a number. But we need to use `until` function and it's a part of `sprig` functions. Mugo has a `sprig` function set by default so we can use it.

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

Binary check functions `eq`, `ne`, `lt`, `le`, `gt`, `ge` but arguments should comparable types and result is boolean.

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

We can use define to define a template and use it later with `template` function and giving variables.

```tpl
{{- define "hello" -}}
Hello {{ .name }}
{{- end -}}

{{ template "hello" . }}
```

Output:

```txt
Merhaba dünya!
```

## With

Use with to limit the scope of a variable inside of a template.

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

`addf` function is a part of `sprig` functions. It uses decimal library to calculate floating point numbers.

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
