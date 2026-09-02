# Function List

The function registry depends on the imported registry packages and the enabled or disabled function options. The CLI is the canonical source for the functions available in the installed version:

```sh
mugo --list --no-stdin --silience
```

The output is sorted by function name and includes signatures when they can be determined at runtime.

## Filtering

Use the same selection flags with `--list` to inspect a restricted function set:

```sh
mugo --list --enable-group sprig --no-stdin --silience
mugo --list --disable-func exec --disable-func file --no-stdin --silience
```

Available selectors:

- `--enable-group`: enable only named direct function groups, such as `sprig`.
- `--enable-func`: enable only named functions or structured function objects.
- `--disable-group`: remove a direct function group.
- `--disable-func`: remove a function or structured function object.

Structured entries such as `codec`, `crypto`, `file`, `map`, `math`, `os`, `random`, and `time` expose methods in templates. For example:

```tpl
{{ codec.ByteToString (codec.StringToByte "mugo") }}
{{ math.Add 1 2 }}
{{ time.RFC3339 }}
```

Functions that execute commands or write files reject those operations unless trust is explicitly enabled with `--trust` or `fstore.WithTrust(true)`.
