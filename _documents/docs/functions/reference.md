# Reference

## ungrouped functions

### exec(string) (map[string]interface {}, error)

Trust required function.

```tpl
{{ exec "echo 'Hello World'" }}
```

Output: `map[stdout:Hello World]`

```
stdout []byte
stderr []byte
status int
```

### execTemplate(string, interface {}) (string, error)

Execute defined template and return.

Input: template-name, data

### nothing([]interface {}) string

Not output anything, just return empty string. Useful for empty processed values.

## sprig (included directly)

`sprig` functions directly available not need to _sprig._ prefix!

### abbrev(int, string) string

```tpl
{{ abbrev 10 "Hello World" }}
```

Output: `Hello W...`

### abbrevboth(int, int, string) string

```tpl
{{ abbrevboth 5 10 "1234 5678 9123" }}
```

Output: `...5678...`

### add([]interface {}) int64

```tpl
{{ add 1 2 3 4 5 }}
```

Output: `15`

### add1(interface {}) int64

```tpl
{{ add1 1 }}
```

Output: `2`

### add1f(interface {}) float64

```tpl
{{ add1f 1.1 }}
```

Output: `2.1`

### addf([]interface {}) float64

```tpl
{{ addf 1.1 2.2 3.3 4.4 5.5 }}
```

Output: `16.5`

### adler32sum(string) string

```tpl
{{ adler32sum "Hello World" }}
```

Output: `403375133`

### ago(interface {}) string

Timestamp to human readable time ago.

```tpl
{{ ago 1687989601 }}
```

Output: `11s`

### all([]interface {}) bool

```tpl
{{ all true true false }}
```

Output: `false`

### any([]interface {}) bool

```tpl
{{ any true true false }}
```

Output: `true`

### append(interface {}, interface {}) []interface {}

First argument must be a slice.

```tpl
{{ append . "add" }}
```

Output: `[1 add]`

### atoi(string) int

```tpl
{{ atoi "123" }}
```

Output: `123`

### b32dec(string) string

```tpl
{{ b32dec "MZXW6YTBOI======" }}
```

Output: `foobar`

### b32enc(string) string

```tpl
{{ b32enc "foobar" }}
```

Output: `MZXW6YTBOI======`

### b64dec(string) string

```tpl
{{ b64dec "SGVsbG8gV29ybGQ=" }}
```

Output: `Hello World`

### b64enc(string) string

```tpl
{{ b64enc "Hello World" }}
```

Output: `SGVsbG8gV29ybGQ=`

### base(string) string

`path.Base` function.

```tpl
{{ base "/foo/bar/baz.js" }}
```

Output: `baz.js`

### bcrypt(string) string

`golang.org/x/crypto/bcrypt` package's `GenerateFromPassword` function.

```tpl
{{ bcrypt "Hello World" }}
```

Output: `$2a$10$dTu/HetKKYglHR1bs9mBFeGkOFgIMAabTdC0TNPh5ucVJLQLTfVYG`

### biggest(interface {}, []interface {}) int64

```tpl
{{ biggest 1 2 3 4 5 }}
```

Output: `5`

### buildCustomCert(string, string) (sprig.certificate, error)
### camelcase(string) string

```tpl
{{ camelcase "hello_world" }}
```

Output: `HelloWorld`

### cat([]interface {}) string

```tpl
{{ cat "Hello" "World" }}
```

Output: `Hello World`

### ceil(interface {}) float64

```tpl
{{ ceil 1.1 }}
```

Output: `2`

### chunk(int, interface {}) [][]interface {}

```tpl
{{ chunk 2 (list 1 2 3 4 5) }}
```

Output: `[[1 2] [3 4] [5]]`

### clean(string) string

`path.Clean` function.

```tpl
{{ clean "/foo/bar/.." }}
```

Output: `/foo`

### coalesce([]interface {}) interface {}

Returns the first non-empty value.

```tpl
{{ coalesce "" "Hello World" }}
```

Output: `Hello World`

### compact(interface {}) []interface {}
### concat([]interface {}) interface {}
### contains(string, string) bool

`strings.Contains` function but arguments are reversed.

```tpl
{{ contains "World" "Hello World" }}
```

Output: `true`

### date(string, interface {}) string

Date can be a `time.Time` or an `int, int32, int64`.

```tpl
{{ date "2006-01-02" 1690151142 }}
```

Output: `2023-07-24`

### dateInZone(string, interface {}, string) string
### dateModify(string, time.Time) time.Time
### date_in_zone(string, interface {}, string) string
### date_modify(string, time.Time) time.Time
### decryptAES(string, string) (string, error)
### deepCopy(interface {}) interface {}
### deepEqual(interface {}, interface {}) bool
### default(interface {}, []interface {}) interface {}
### derivePassword(uint32, string, string, string, string) string
### dict([]interface {}) map[string]interface {}

```tpl
{{ dict "a" 1 "b" 2 }}
```

Output: `map[a:1 b:2]`

### dig([]interface {}) (interface {}, error)
### dir(string) string

`path.Dir` function.

```tpl
{{ dir "/foo/bar/baz.js" }}
```

Output: `/foo/bar`

### div(interface {}, interface {}) int64
### divf(interface {}, []interface {}) float64
### duration(interface {}) string

```tpl
{{ duration "3600" }}
```

Output: `1h0m0s`

### durationRound(interface {}) string
### empty(interface {}) bool

```tpl
{{ empty "" }}
```

Output: `true`

### encryptAES(string, string) (string, error)
### env(string) string

```tpl
{{ env "HOME" }}
```

Output: `/home/rytsh`

### expandenv(string) string

```tpl
{{ expandenv "$HOME" }}
{{ expandenv "${USER}" }}
```

Output:

```
/home/rytsh
rytsh
```

### ext(string) string

`path.Ext` function.

```tpl
{{ ext "/foo/bar/baz.js" }}
```

Output: `.js`

### fail(string) (string, error)

```tpl
{{ fail "FAILED PROGRAM" }}
```

Output: program fail, exit code 1

### first(interface {}) interface {}
### float64(interface {}) float64
### floor(interface {}) float64
### fromJson(string) interface {}
### genCA(string, int) (sprig.certificate, error)
### genCAWithKey(string, int, string) (sprig.certificate, error)
### genPrivateKey(string) string
### genSelfSignedCert(string, []interface {}, []interface {}, int) (sprig.certificate, error)
### genSelfSignedCertWithKey(string, []interface {}, []interface {}, int, string) (sprig.certificate, error)
### genSignedCert(string, []interface {}, []interface {}, int, sprig.certificate) (sprig.certificate, error)
### genSignedCertWithKey(string, []interface {}, []interface {}, int, sprig.certificate, string) (sprig.### certificate, error)
### get(map[string]interface {}, string) interface {}
### getHostByName(string) string
### has(interface {}, interface {}) bool
### hasKey(map[string]interface {}, string) bool
### hasPrefix(string, string) bool
### hasSuffix(string, string) bool
### hello() string
### htmlDate(interface {}) string
### htmlDateInZone(interface {}, string) string
### htpasswd(string, string) string
### indent(int, string) string
### initial(interface {}) []interface {}
### initials(string) string
### int(interface {}) int
### int64(interface {}) int64
### isAbs(string) bool
### join(string, interface {}) string
### kebabcase(string) string
### keys([]map[string]interface {}) []string
### kindIs(string, interface {}) bool
### kindOf(interface {}) string
### last(interface {}) interface {}
### list([]interface {}) []interface {}
### lower(string) string
### max(interface {}, []interface {}) int64
### maxf(interface {}, []interface {}) float64
### merge(map[string]interface {}, []map[string]interface {}) interface {}
### mergeOverwrite(map[string]interface {}, []map[string]interface {}) interface {}
### min(interface {}, []interface {}) int64
### minf(interface {}, []interface {}) float64
### mod(interface {}, interface {}) int64
### mul(interface {}, []interface {}) int64
### mulf(interface {}, []interface {}) float64
### mustAppend(interface {}, interface {}) ([]interface {}, error)
### mustChunk(int, interface {}) ([][]interface {}, error)
### mustCompact(interface {}) ([]interface {}, error)
### mustDateModify(string, time.Time) (time.Time, error)
### mustDeepCopy(interface {}) (interface {}, error)
### mustFirst(interface {}) (interface {}, error)
### mustFromJson(string) (interface {}, error)
### mustHas(interface {}, interface {}) (bool, error)
### mustInitial(interface {}) ([]interface {}, error)
### mustLast(interface {}) (interface {}, error)
### mustMerge(map[string]interface {}, []map[string]interface {}) (interface {}, error)
### mustMergeOverwrite(map[string]interface {}, []map[string]interface {}) (interface {}, error)
### mustPrepend(interface {}, interface {}) ([]interface {}, error)
### mustPush(interface {}, interface {}) ([]interface {}, error)
### mustRegexFind(string, string) (string, error)
### mustRegexFindAll(string, string, int) ([]string, error)
### mustRegexMatch(string, string) (bool, error)
### mustRegexReplaceAll(string, string, string) (string, error)
### mustRegexReplaceAllLiteral(string, string, string) (string, error)
### mustRegexSplit(string, string, int) ([]string, error)
### mustRest(interface {}) ([]interface {}, error)
### mustReverse(interface {}) ([]interface {}, error)
### mustSlice(interface {}, []interface {}) (interface {}, error)
### mustToDate(string, string) (time.Time, error)
### mustToJson(interface {}) (string, error)
### mustToPrettyJson(interface {}) (string, error)
### mustToRawJson(interface {}) (string, error)
### mustUniq(interface {}) ([]interface {}, error)
### mustWithout(interface {}, []interface {}) ([]interface {}, error)
### must_date_modify(string, time.Time) (time.Time, error)
### nindent(int, string) string
### nospace(string) string
### nothing([]interface {}) string
### now() time.Time
### omit(map[string]interface {}, []string) map[string]interface {}
### osBase(string) string
### osClean(string) string
### osDir(string) string
### osExt(string) string
### osIsAbs(string) bool
### pick(map[string]interface {}, []string) map[string]interface {}
### pluck(string, []map[string]interface {}) []interface {}
### plural(string, string, int) string
### prepend(interface {}, interface {}) []interface {}
### push(interface {}, interface {}) []interface {}
### quote([]interface {}) string
### randAlpha(int) string
### randAlphaNum(int) string
### randAscii(int) string
### randBytes(int) (string, error)
### randInt(int, int) int
### randNumeric(int) string
### regexFind(string, string) string
### regexFindAll(string, string, int) []string
### regexMatch(string, string) bool
### regexQuoteMeta(string) string
### regexReplaceAll(string, string, string) string
### regexReplaceAllLiteral(string, string, string) string
### regexSplit(string, string, int) []string
### repeat(int, string) string
### replace(string, string, string) string
### rest(interface {}) []interface {}
### reverse(interface {}) []interface {}
### round(interface {}, int, []float64) float64
### semver(string) (*semver.Version, error)
### semverCompare(string, string) (bool, error)
### seq([]int) string
### set(map[string]interface {}, string, interface {}) map[string]interface {}
### sha1sum(string) string
### sha256sum(string) string
### shuffle(string) string
### slice(interface {}, []interface {}) interface {}
### snakecase(string) string
### sortAlpha(interface {}) []string
### split(string, string) map[string]string
### splitList(string, string) []string
### splitn(string, int, string) map[string]string
### squote([]interface {}) string
### sub(interface {}, interface {}) int64
### subf(interface {}, []interface {}) float64
### substr(int, int, string) string
### swapcase(string) string
### ternary(interface {}, interface {}, bool) interface {}
### title(string) string
### toDate(string, string) time.Time
### toDecimal(interface {}) int64
### toJson(interface {}) string
### toPrettyJson(interface {}) string
### toRawJson(interface {}) string
### toString(interface {}) string
### toStrings(interface {}) []string
### trim(string) string
### trimAll(string, string) string
### trimPrefix(string, string) string
### trimSuffix(string, string) string
### trimall(string, string) string
### trunc(int, string) string
### tuple([]interface {}) []interface {}
### typeIs(string, interface {}) bool
### typeIsLike(string, interface {}) bool
### typeOf(interface {}) string
### uniq(interface {}) []interface {}
### unixEpoch(time.Time) string
### unset(map[string]interface {}, string) map[string]interface {}
### until(int) []int
### untilStep(int, int, int) []int
### untitle(string) string
### upper(string) string
### urlJoin(map[string]interface {}) string
### urlParse(string) map[string]interface {}
### uuidv4() string
### values(map[string]interface {}) []interface {}
### without(interface {}, []interface {}) []interface {}
### wrap(int, string) string
### wrapWith(int, string, string) string

## cast

Functions of [github.com/spf13/cast](https://github.com/spf13/cast) package.

### cast.StringToDate(s string) (time.Time, error)
### cast.StringToDateInDefaultLocation(s string, location *time.Location) (time.Time, error)
### cast.ToBool(i interface{}) bool
### cast.ToBoolE(i interface{}) (bool, error)
### cast.ToBoolSlice(i interface{}) []bool
### cast.ToBoolSliceE(i interface{}) ([]bool, error)
### cast.ToDuration(i interface{}) time.Duration
### cast.ToDurationE(i interface{}) (d time.Duration, err error)
### cast.ToDurationSlice(i interface{}) []time.Duration
### cast.ToDurationSliceE(i interface{}) ([]time.Duration, error)
### cast.ToFloat32(i interface{}) float32
### cast.ToFloat32E(i interface{}) (float32, error)
### cast.ToFloat64(i interface{}) float64
### cast.ToFloat64E(i interface{}) (float64, error)
### cast.ToInt(i interface{}) int
### cast.ToInt16(i interface{}) int16
### cast.ToInt16E(i interface{}) (int16, error)
### cast.ToInt32(i interface{}) int32
### cast.ToInt32E(i interface{}) (int32, error)
### cast.ToInt64(i interface{}) int64
### cast.ToInt64E(i interface{}) (int64, error)
### cast.ToInt8(i interface{}) int8
### cast.ToInt8E(i interface{}) (int8, error)
### cast.ToIntE(i interface{}) (int, error)
### cast.ToIntSlice(i interface{}) []int
### cast.ToIntSliceE(i interface{}) ([]int, error)
### cast.ToSlice(i interface{}) []interface{}
### cast.ToSliceE(i interface{}) ([]interface{}, error)
### cast.ToString(i interface{}) string
### cast.ToStringE(i interface{}) (string, error)
### cast.ToStringMap(i interface{}) map[string]interface{}
### cast.ToStringMapBool(i interface{}) map[string]bool
### cast.ToStringMapBoolE(i interface{}) (map[string]bool, error)
### cast.ToStringMapE(i interface{}) (map[string]interface{}, error)
### cast.ToStringMapInt(i interface{}) map[string]int
### cast.ToStringMapInt64(i interface{}) map[string]int64
### cast.ToStringMapInt64E(i interface{}) (map[string]int64, error)
### cast.ToStringMapIntE(i interface{}) (map[string]int, error)
### cast.ToStringMapString(i interface{}) map[string]string
### cast.ToStringMapStringE(i interface{}) (map[string]string, error)
### cast.ToStringMapStringSlice(i interface{}) map[string][]string
### cast.ToStringMapStringSliceE(i interface{}) (map[string][]string, error)
### cast.ToStringSlice(i interface{}) []string
### cast.ToStringSliceE(i interface{}) ([]string, error)
### cast.ToTime(i interface{}) time.Time
### cast.ToTimeE(i interface{}) (tim time.Time, err error)
### cast.ToTimeInDefaultLocation(i interface{}, location *time.Location) time.Time
### cast.ToTimeInDefaultLocationE(i interface{}, location *time.Location) (tim time.Time, err error)
### cast.ToUint(i interface{}) uint
### cast.ToUint16(i interface{}) uint16
### cast.ToUint16E(i interface{}) (uint16, error)
### cast.ToUint32(i interface{}) uint32
### cast.ToUint32E(i interface{}) (uint32, error)
### cast.ToUint64(i interface{}) uint64
### cast.ToUint64E(i interface{}) (uint64, error)
### cast.ToUint8(i interface{}) uint8
### cast.ToUint8E(i interface{}) (uint8, error)
### cast.ToUintE(i interface{}) (uint, error)

## codec

### codec.JsonDecode(v []byte) (any, error)
### codec.JsonEncode(v any, pretty bool) ([]byte, error)
### codec.YamlDecode(v []byte) (any, error)
### codec.YamlEncode(v any) ([]byte, error)
### codec.TomlDecode(v []byte) (any, error)
### codec.TomlEncode(v any) ([]byte, error)
### codec.Markdown(data []byte) []byte
### codec.ByteToString(b []byte) string
### codec.StringToByte(s string) []byte
### codec.IndentByte(i int, data []byte) []byte

## crypto

### crypto.Base64(v any) (string, error)
### crypto.Base64B(v []byte) string
### crypto.Base64Decode(v any) ([]byte, error)
### crypto.MD5(v any) (string, error)
### crypto.MD5B(v []byte) []byte
### crypto.SHA1(v any) (string, error)
### crypto.SHA1B(v []byte) []byte
### crypto.SHA256(v any) (string, error)
### crypto.SHA256B(v []byte) []byte
### crypto.FNV32a(v any) (int, error)
### crypto.HMAC(h interface{}, k interface{}, m interface{}) (string, error)
### crypto.JwtParseUnverified(token string) (map[string]interface{}, error)

## faker

Functions of [github.com/jaswdr/faker](https://github.com/jaswdr/faker) package.

## file

Trust required!

### file.Save(fileName string, data []byte) (bool, error)
### file.Write(fileName string, data []byte) (bool, error)
### file.Read(fileName string) ([]byte, error)

## html2

### html2.EscapeString(v string) string
### html2.UnescapeString(v string) string

## humanize

Functions of [github.com/dustin/go-humanize](https://github.com/dustin/go-humanize) package.

### humanize.BigBytes(s *big.Int) string
### humanize.BigComma(b *big.Int) string
### humanize.BigCommaf(v *big.Float) string
### humanize.BigIBytes(s *big.Int) string
### humanize.Bytes(v uint64) string
### humanize.Comma(v int64) string
### humanize.Commaf(v float64) string
### humanize.CommafWithDigits(f float64, decimals int) string
### humanize.ComputeSI(input float64) (float64, string)
### humanize.CustomRelTime(a time.Time, b time.Time, albl string, blbl string, magnitudes []humanize.RelTimeMagnitude) string
### humanize.FormatFloat(format string, n float64) string
### humanize.FormatInteger(format string, n int) string
### humanize.Ftoa(num float64) string
### humanize.FtoaWithDigits(num float64, digits int) string
### humanize.IBytes(v uint64) string
### humanize.Ordinal(x int) string
### humanize.ParseBigBytes(s string) (*big.Int, error)
### humanize.ParseBytes(s string) (uint64, error)
### humanize.RelTime(a time.Time, b time.Time, albl string, blbl string) string
### humanize.SI(input float64, unit string) string
### humanize.Time(then time.Time) string

## log

Before to use log functions, you need to give a logger when initialize `fstore` functions.

Example of zerolog usage:

```go
fstore.FuncMap(
	fstore.WithLog(logz.AdapterKV{Log: log.Logger}),
)
```

Arguments should be support `Adapter` interface.

```go
type Adapter interface {
	Error(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Debug(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
}
```

Log functions return the same value (...interface{} part) as input.

### log.Debug(string, ...interface {}) interface {}
### log.Error(string, ...interface {}) interface {}
### log.Info(string, ...interface {}) interface {}
### log.Warn(string, ...interface {}) interface {}

## map

### map.Set(key string, value interface{}) map[string]interface{}
### map.Get(key string, data map[string]interface{}) interface{}

## os

### os.ReadDir(i any) ([]os.FileInfo, error)
### os.ReadFile(i any) (string, error)
### os.FileExists(i any) (bool, error)
### os.Stat(i any) (os.FileInfo, error)

## time

### time.Now() time.Time
### time.RFC3339() string
### time.Format(format string, t time.Time) string
### time.UTC(t time.Time) time.Time
### time.AddDuration(t time.Time, d time.Duration) time.Time
### time.Duration(d string) (time.Duration, error)
### time.AddDate(t time.Time, years, months, days int) time.Time

## random

### random.Intn(min, max int) int
### random.Alpha(n int) string
### random.AlphaNum(n int) string
### random.Ascii(n int) string
### random.Numeric(n int) string
### random.Float(min float64, max float64) float64

## math

_decimal.Decimal_ functions, accepts any type of number with _decimal.Decimal_, _string_, _json.Number_.

### Abs(interface {}) (json.Number, error)
### Add(interface {}, ...interface {}) (json.Number, error)
### Ceil(interface {}) (json.Number, error)
### Div(interface {}, ...interface {}) (json.Number, error)
### Equal(interface {}, interface {}) (bool, error)
### Floor(interface {}) (json.Number, error)
### GreaterThan(interface {}, interface {}) (bool, error)
### GreaterThanOrEqual(interface {}, interface {}) (bool, error)
### LessThan(interface {}, interface {}) (bool, error)
### LessThanOrEqual(interface {}, interface {}) (bool, error)
### Max(interface {}, ...interface {}) (json.Number, error)
### Min(interface {}, ...interface {}) (json.Number, error)
### Mod(interface {}, interface {}) (json.Number, error)
### Mul(interface {}, ...interface {}) (json.Number, error)
### Pow(interface {}, interface {}) (json.Number, error)
### Round(interface {}, interface {}) (json.Number, error)
### RoundBankers(interface {}, interface {}) (json.Number, error)
### RoundCash(interface {}, interface {}) (json.Number, error)
### RoundCeil(interface {}, interface {}) (json.Number, error)
### RoundDown(interface {}, interface {}) (json.Number, error)
### RoundFloor(interface {}, interface {}) (json.Number, error)
### RoundUp(interface {}, interface {}) (json.Number, error)
### Sign(interface {}) (int, error)
### Sub(interface {}, ...interface {}) (json.Number, error)
### Truncate(interface {}, interface {}) (json.Number, error)
