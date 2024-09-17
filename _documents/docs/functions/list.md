# List

List function result of `mugo -l` command:

```go
abbrev(int, string) string
abbrevboth(int, int, string) string
add(...interface {}) int64
add1(interface {}) int64
add1f(interface {}) float64
addf(...interface {}) float64
adler32sum(string) string
ago(interface {}) string
all(...interface {}) bool
any(...interface {}) bool
append(interface {}, interface {}) []interface {}
atoi(string) int
b32dec(string) string
b32enc(string) string
b64dec(string) string
b64enc(string) string
base(string) string
bcrypt(string) string
biggest(interface {}, ...interface {}) int64
buildCustomCert(string, string) (sprig.certificate, error)
camelcase(string) string
cast() cast.Cast
├─ StringToDate(string) (time.Time, error)
├─ StringToDateInDefaultLocation(string, *time.Location) (time.Time, error)
├─ ToBool(interface {}) bool
├─ ToBoolE(interface {}) (bool, error)
├─ ToBoolSlice(interface {}) []bool
├─ ToBoolSliceE(interface {}) ([]bool, error)
├─ ToDuration(interface {}) time.Duration
├─ ToDurationE(interface {}) (time.Duration, error)
├─ ToDurationSlice(interface {}) []time.Duration
├─ ToDurationSliceE(interface {}) ([]time.Duration, error)
├─ ToFloat32(interface {}) float32
├─ ToFloat32E(interface {}) (float32, error)
├─ ToFloat64(interface {}) float64
├─ ToFloat64E(interface {}) (float64, error)
├─ ToInt(interface {}) int
├─ ToInt16(interface {}) int16
├─ ToInt16E(interface {}) (int16, error)
├─ ToInt32(interface {}) int32
├─ ToInt32E(interface {}) (int32, error)
├─ ToInt64(interface {}) int64
├─ ToInt64E(interface {}) (int64, error)
├─ ToInt8(interface {}) int8
├─ ToInt8E(interface {}) (int8, error)
├─ ToIntE(interface {}) (int, error)
├─ ToIntSlice(interface {}) []int
├─ ToIntSliceE(interface {}) ([]int, error)
├─ ToSlice(interface {}) []interface {}
├─ ToSliceE(interface {}) ([]interface {}, error)
├─ ToString(interface {}) string
├─ ToStringE(interface {}) (string, error)
├─ ToStringMap(interface {}) map[string]interface {}
├─ ToStringMapBool(interface {}) map[string]bool
├─ ToStringMapBoolE(interface {}) (map[string]bool, error)
├─ ToStringMapE(interface {}) (map[string]interface {}, error)
├─ ToStringMapInt(interface {}) map[string]int
├─ ToStringMapInt64(interface {}) map[string]int64
├─ ToStringMapInt64E(interface {}) (map[string]int64, error)
├─ ToStringMapIntE(interface {}) (map[string]int, error)
├─ ToStringMapString(interface {}) map[string]string
├─ ToStringMapStringE(interface {}) (map[string]string, error)
├─ ToStringMapStringSlice(interface {}) map[string][]string
├─ ToStringMapStringSliceE(interface {}) (map[string][]string, error)
├─ ToStringSlice(interface {}) []string
├─ ToStringSliceE(interface {}) ([]string, error)
├─ ToTime(interface {}) time.Time
├─ ToTimeE(interface {}) (time.Time, error)
├─ ToTimeInDefaultLocation(interface {}, *time.Location) time.Time
├─ ToTimeInDefaultLocationE(interface {}, *time.Location) (time.Time, error)
├─ ToUint(interface {}) uint
├─ ToUint16(interface {}) uint16
├─ ToUint16E(interface {}) (uint16, error)
├─ ToUint32(interface {}) uint32
├─ ToUint32E(interface {}) (uint32, error)
├─ ToUint64(interface {}) uint64
├─ ToUint64E(interface {}) (uint64, error)
├─ ToUint8(interface {}) uint8
├─ ToUint8E(interface {}) (uint8, error)
└─ ToUintE(interface {}) (uint, error)
cat(...interface {}) string
ceil(interface {}) float64
chunk(int, interface {}) [][]interface {}
clean(string) string
coalesce(...interface {}) interface {}
codec() codec.Codec
├─ ByteToString([]uint8) string
├─ IndentByte(int, []uint8) []uint8
├─ JsonDecode([]uint8) (interface {}, error)
├─ JsonEncode(interface {}, bool) ([]uint8, error)
├─ Markdown([]uint8) []uint8
├─ StringToByte(string) []uint8
├─ TomlDecode([]uint8) (interface {}, error)
├─ TomlEncode(interface {}) ([]uint8, error)
├─ YamlDecode([]uint8) (interface {}, error)
└─ YamlEncode(interface {}) ([]uint8, error)
compact(interface {}) []interface {}
concat(...interface {}) interface {}
contains(string, string) bool
crypto() crypto.Crypto
├─ Base64(interface {}) (string, error)
├─ Base64B([]uint8) string
├─ Base64Decode(interface {}) ([]uint8, error)
├─ FNV32a(interface {}) (int, error)
├─ HMAC(interface {}, interface {}, interface {}) (string, error)
├─ JwtParseUnverified(string) (map[string]interface {}, error)
├─ MD5(interface {}) (string, error)
├─ MD5B([]uint8) []uint8
├─ SHA1(interface {}) (string, error)
├─ SHA1B([]uint8) []uint8
├─ SHA256(interface {}) (string, error)
└─ SHA256B([]uint8) []uint8
date(string, interface {}) string
dateInZone(string, interface {}, string) string
dateModify(string, time.Time) time.Time
date_in_zone(string, interface {}, string) string
date_modify(string, time.Time) time.Time
decryptAES(string, string) (string, error)
deepCopy(interface {}) interface {}
deepEqual(interface {}, interface {}) bool
default(interface {}, ...interface {}) interface {}
derivePassword(uint32, string, string, string, string) string
dict(...interface {}) map[string]interface {}
dig(...interface {}) (interface {}, error)
dir(string) string
div(interface {}, interface {}) int64
divf(interface {}, ...interface {}) float64
duration(interface {}) string
durationRound(interface {}) string
empty(interface {}) bool
encryptAES(string, string) (string, error)
env(string) string
exec(string) (map[string]interface {}, error)
execTemplate(string, interface {}) (string, error)
expandenv(string) string
ext(string) string
fail(string) (string, error)
faker() faker.Faker
├┐ Address() faker.Address
│├─ Address() string
│├─ BuildingNumber() string
│├─ City() string
│├─ CityPrefix() string
│├─ CitySuffix() string
│├─ Country() string
│├─ CountryAbbr() string
│├─ CountryCode() string
│├─ Latitude() float64
│├─ Longitude() float64
│├─ PostCode() string
│├─ SecondaryAddress() string
│├─ State() string
│├─ StateAbbr() string
│├─ StreetAddress() string
│├─ StreetName() string
│└─ StreetSuffix() string
├┐ App() faker.App
│├─ Name() string
│└─ Version() string
├─ Asciify(string) string
├┐ Beer() faker.Beer
│├─ Alcohol() string
│├─ Blg() string
│├─ Hop() string
│├─ Ibu() string
│├─ Malt() string
│├─ Name() string
│└─ Style() string
├┐ BinaryString() faker.BinaryString
│└─ BinaryString(int) string
├┐ Blood() faker.Blood
│└─ Name() string
├─ Bool() bool
├─ BoolWithChance(int) bool
├┐ Boolean() faker.Boolean
│├─ Bool() bool
│├─ BoolInt() int
│├─ BoolString(string, string) string
│└─ BoolWithChance(int) bool
├─ Bothify(string) string
├┐ Car() faker.Car
│├─ Category() string
│├─ FuelType() string
│├─ Maker() string
│├─ Model() string
│├─ Plate() string
│└─ TransmissionGear() string
├┐ Color() faker.Color
│├─ CSS() string
│├─ ColorName() string
│├─ Hex() string
│├─ RGB() string
│├─ RGBAsArray() [3]string
│└─ SafeColorName() string
├┐ Company() faker.Company
│├─ BS() string
│├─ CatchPhrase() string
│├─ EIN() int
│├─ JobTitle() string
│├─ Name() string
│└─ Suffix() string
├┐ Crypto() faker.Crypto
│├─ Bech32Address() string
│├─ Bech32AddressWithLength(int) string
│├─ BitcoinAddress() string
│├─ EtheriumAddress() string
│├─ P2PKHAddress() string
│├─ P2PKHAddressWithLength(int) string
│├─ P2SHAddress() string
│└─ P2SHAddressWithLength(int) string
├┐ Currency() faker.Currency
│├─ Code() string
│├─ Country() string
│├─ Currency() string
│├─ CurrencyAndCode() (string, string)
│└─ Number() int
├┐ Directory() faker.Directory
│├─ Directory(int) string
│├─ DriveLetter() string
│├─ UnixDirectory(int) string
│└─ WindowsDirectory(int) string
├┐ Emoji() faker.Emoji
│├─ Emoji() string
│└─ EmojiCode() string
├┐ File() faker.File
│├─ AbsoluteFilePath(int) string
│├─ AbsoluteFilePathForUnix(int) string
│├─ AbsoluteFilePathForWindows(int) string
│├─ Extension() string
│└─ FilenameWithExtension() string
├─ Float(int, int, int) float64
├─ Float32(int, int, int) float32
├─ Float64(int, int, int) float64
├┐ Food() faker.Food
│├─ Fruit() string
│└─ Vegetable() string
├┐ Gamer() faker.Gamer
│└─ Tag() string
├┐ Gender() faker.Gender
│├─ Abbr() string
│└─ Name() string
├┐ Genre() faker.Genre
│├─ Name() string
│└─ NameWithDescription() (string, string)
├┐ Hash() faker.Hash
│├─ MD5() string
│├─ SHA256() string
│└─ SHA512() string
├┐ Image() faker.Image
│└─ Image(int, int) *os.File
├─ Int() int
├─ Int16() int16
├─ Int16Between(int16, int16) int16
├─ Int32() int32
├─ Int32Between(int32, int32) int32
├─ Int64() int64
├─ Int64Between(int64, int64) int64
├─ Int8() int8
├─ Int8Between(int8, int8) int8
├─ IntBetween(int, int) int
├┐ Internet() faker.Internet
│├─ CompanyEmail() string
│├─ Domain() string
│├─ Email() string
│├─ FreeEmail() string
│├─ FreeEmailDomain() string
│├─ HTTPMethod() string
│├─ Ipv4() string
│├─ Ipv6() string
│├─ LocalIpv4() string
│├─ MacAddress() string
│├─ Password() string
│├─ Query() string
│├─ SafeEmail() string
│├─ SafeEmailDomain() string
│├─ Slug() string
│├─ StatusCode() int
│├─ StatusCodeMessage() string
│├─ StatusCodeWithMessage() string
│├─ TLD() string
│├─ URL() string
│└─ User() string
├─ Json() faker.Json
├┐ Language() faker.Language
│├─ Language() string
│├─ LanguageAbbr() string
│└─ ProgrammingLanguage() string
├─ Letter() string
├─ Lexify(string) string
├┐ Lorem() faker.Lorem
│├─ Bytes(int) []uint8
│├─ Paragraph(int) string
│├─ Paragraphs(int) []string
│├─ Sentence(int) string
│├─ Sentences(int) []string
│├─ Text(int) string
│├─ Word() string
│└─ Words(int) []string
├┐ LoremFlickr() faker.LoremFlickr
│└─ Image(int, int, []string, string, bool) *os.File
├─ Map() map[string]interface {}
├┐ MimeType() faker.MimeType
│└─ MimeType() string
├┐ Music() faker.Music
│├─ Author() string
│├─ Genre() string
│├─ Length() time.Duration
│└─ Name() string
├─ Numerify(string) string
├┐ Payment() faker.Payment
│├─ CreditCardExpirationDateString() string
│├─ CreditCardNumber() string
│└─ CreditCardType() string
├┐ Person() faker.Person
│├─ Contact() faker.ContactInfo
│├─ FirstName() string
│├─ FirstNameFemale() string
│├─ FirstNameMale() string
│├─ Gender() string
│├─ GenderFemale() string
│├─ GenderMale() string
│├─ Image() *os.File
│├─ LastName() string
│├─ Name() string
│├─ NameAndGender() (string, string)
│├─ NameFemale() string
│├─ NameMale() string
│├─ SSN() string
│├─ Suffix() string
│├─ Title() string
│├─ TitleFemale() string
│└─ TitleMale() string
├┐ Pet() faker.Pet
│├─ Cat() string
│├─ Dog() string
│└─ Name() string
├┐ Phone() faker.Phone
│├─ AreaCode() string
│├─ E164Number() string
│├─ ExchangeCode() string
│├─ Number() string
│├─ TollFreeAreaCode() string
│└─ ToolFreeNumber() string
├┐ ProfileImage() faker.ProfileImage
│└─ Image() *os.File
├─ RandomDigit() int
├─ RandomDigitNot(...int) int
├─ RandomDigitNotNull() int
├─ RandomFloat(int, int, int) float64
├─ RandomIntElement([]int) int
├─ RandomLetter() string
├─ RandomNumber(int) int
├─ RandomStringElement([]string) string
├─ RandomStringMapKey(map[string]string) string
├─ RandomStringMapValue(map[string]string) string
├─ RandomStringWithLength(int) string
├─ ShuffleString(string) string
├┐ Struct() faker.Struct
│└─ Fill(interface {})
├┐ Time() faker.Time
│├─ ANSIC(time.Time) string
│├─ AmPm() string
│├─ Century() string
│├─ DayOfMonth() int
│├─ DayOfWeek() time.Weekday
│├─ ISO8601(time.Time) string
│├─ Kitchen(time.Time) string
│├─ Month() time.Month
│├─ MonthName() string
│├─ RFC1123(time.Time) string
│├─ RFC1123Z(time.Time) string
│├─ RFC3339(time.Time) string
│├─ RFC3339Nano(time.Time) string
│├─ RFC822(time.Time) string
│├─ RFC822Z(time.Time) string
│├─ RFC850(time.Time) string
│├─ RubyDate(time.Time) string
│├─ Time(time.Time) time.Time
│├─ TimeBetween(time.Time, time.Time) time.Time
│├─ Timezone() string
│├─ Unix(time.Time) int64
│├─ UnixDate(time.Time) string
│└─ Year() int
├─ UInt() uint
├─ UInt16() uint16
├─ UInt16Between(uint16, uint16) uint16
├─ UInt32() uint32
├─ UInt32Between(uint32, uint32) uint32
├─ UInt64() uint64
├─ UInt64Between(uint64, uint64) uint64
├─ UInt8() uint8
├─ UInt8Between(uint8, uint8) uint8
├─ UIntBetween(uint, uint) uint
├┐ UUID() faker.UUID
│└─ V4() string
├┐ UserAgent() faker.UserAgent
│├─ Chrome() string
│├─ Firefox() string
│├─ InternetExplorer() string
│├─ Opera() string
│├─ Safari() string
│└─ UserAgent() string
└┐ YouTube() faker.YouTube
 ├─ GenerateEmbededURL() string
 ├─ GenerateFullURL() string
 ├─ GenerateShareURL() string
 └─ GenerateVideoID() string
file() *file.File
├─ Read(string) ([]uint8, error)
├─ Save(string, []uint8) (bool, error)
└─ Write(string, []uint8) (bool, error)
first(interface {}) interface {}
float64(interface {}) float64
floor(interface {}) float64
fromJson(string) interface {}
genCA(string, int) (sprig.certificate, error)
genCAWithKey(string, int, string) (sprig.certificate, error)
genPrivateKey(string) string
genSelfSignedCert(string, []interface {}, []interface {}, int) (sprig.certificate, error)
genSelfSignedCertWithKey(string, []interface {}, []interface {}, int, string) (sprig.certificate, error)
genSignedCert(string, []interface {}, []interface {}, int, sprig.certificate) (sprig.certificate, error)
genSignedCertWithKey(string, []interface {}, []interface {}, int, sprig.certificate, string) (sprig.certificate, error)
get(map[string]interface {}, string) interface {}
getHostByName(string) string
has(interface {}, interface {}) bool
hasKey(map[string]interface {}, string) bool
hasPrefix(string, string) bool
hasSuffix(string, string) bool
hello() string
html2() html2.HTML2
├─ EscapeString(string) string
└─ UnescapeString(string) string
htmlDate(interface {}) string
htmlDateInZone(interface {}, string) string
htpasswd(string, string) string
humanize() humanize.Humanize
├─ BigBytes(*big.Int) string
├─ BigComma(*big.Int) string
├─ BigCommaf(*big.Float) string
├─ BigIBytes(*big.Int) string
├─ Bytes(uint64) string
├─ Comma(int64) string
├─ Commaf(float64) string
├─ CommafWithDigits(float64, int) string
├─ ComputeSI(float64) (float64, string)
├─ CustomRelTime(time.Time, time.Time, string, string, []humanize.RelTimeMagnitude) string
├─ FormatFloat(string, float64) string
├─ FormatInteger(string, int) string
├─ Ftoa(float64) string
├─ FtoaWithDigits(float64, int) string
├─ IBytes(uint64) string
├─ Ordinal(int) string
├─ ParseBigBytes(string) (*big.Int, error)
├─ ParseBytes(string) (uint64, error)
├─ RelTime(time.Time, time.Time, string, string) string
├─ SI(float64, string) string
└─ Time(time.Time) string
indent(int, string) string
initial(interface {}) []interface {}
initials(string) string
int(interface {}) int
int64(interface {}) int64
isAbs(string) bool
join(string, interface {}) string
kebabcase(string) string
keys(...map[string]interface {}) []string
kindIs(string, interface {}) bool
kindOf(interface {}) string
last(interface {}) interface {}
list(...interface {}) []interface {}
log() log.Log
├─ Debug(interface {}) interface {}
├─ Error(interface {}) interface {}
├─ Fatal(interface {}) interface {}
├─ Info(interface {}) interface {}
├─ Panic(interface {}) interface {}
└─ Warn(interface {}) interface {}
lower(string) string
map() *maps.Map
├─ Get(string, map[string]interface {}) interface {}
└─ Set(string, interface {}) map[string]interface {}
math() math.Math
└─ RoundDecimal(int, float64) float64
max(interface {}, ...interface {}) int64
maxf(interface {}, ...interface {}) float64
merge(map[string]interface {}, ...map[string]interface {}) interface {}
mergeOverwrite(map[string]interface {}, ...map[string]interface {}) interface {}
min(interface {}, ...interface {}) int64
minf(interface {}, ...interface {}) float64
minify() minify.Minify
mod(interface {}, interface {}) int64
mul(interface {}, ...interface {}) int64
mulf(interface {}, ...interface {}) float64
mustAppend(interface {}, interface {}) ([]interface {}, error)
mustChunk(int, interface {}) ([][]interface {}, error)
mustCompact(interface {}) ([]interface {}, error)
mustDateModify(string, time.Time) (time.Time, error)
mustDeepCopy(interface {}) (interface {}, error)
mustFirst(interface {}) (interface {}, error)
mustFromJson(string) (interface {}, error)
mustHas(interface {}, interface {}) (bool, error)
mustInitial(interface {}) ([]interface {}, error)
mustLast(interface {}) (interface {}, error)
mustMerge(map[string]interface {}, ...map[string]interface {}) (interface {}, error)
mustMergeOverwrite(map[string]interface {}, ...map[string]interface {}) (interface {}, error)
mustPrepend(interface {}, interface {}) ([]interface {}, error)
mustPush(interface {}, interface {}) ([]interface {}, error)
mustRegexFind(string, string) (string, error)
mustRegexFindAll(string, string, int) ([]string, error)
mustRegexMatch(string, string) (bool, error)
mustRegexReplaceAll(string, string, string) (string, error)
mustRegexReplaceAllLiteral(string, string, string) (string, error)
mustRegexSplit(string, string, int) ([]string, error)
mustRest(interface {}) ([]interface {}, error)
mustReverse(interface {}) ([]interface {}, error)
mustSlice(interface {}, ...interface {}) (interface {}, error)
mustToDate(string, string) (time.Time, error)
mustToJson(interface {}) (string, error)
mustToPrettyJson(interface {}) (string, error)
mustToRawJson(interface {}) (string, error)
mustUniq(interface {}) ([]interface {}, error)
mustWithout(interface {}, ...interface {}) ([]interface {}, error)
must_date_modify(string, time.Time) (time.Time, error)
nindent(int, string) string
nospace(string) string
nothing(...interface {}) string
now() time.Time
omit(map[string]interface {}, ...string) map[string]interface {}
os() *os.Os
├─ FileExists(interface {}) (bool, error)
├─ ReadDir(interface {}) ([]fs.FileInfo, error)
├─ ReadFile(interface {}) (string, error)
└─ Stat(interface {}) (fs.FileInfo, error)
osBase(string) string
osClean(string) string
osDir(string) string
osExt(string) string
osIsAbs(string) bool
pick(map[string]interface {}, ...string) map[string]interface {}
pluck(string, ...map[string]interface {}) []interface {}
plural(string, string, int) string
prepend(interface {}, interface {}) []interface {}
push(interface {}, interface {}) []interface {}
quote(...interface {}) string
randAlpha(int) string
randAlphaNum(int) string
randAscii(int) string
randBytes(int) (string, error)
randInt(int, int) int
randNumeric(int) string
random() random.Random
├─ Alpha(int) string
├─ AlphaNum(int) string
├─ Ascii(int) string
├─ Float(float64, float64) float64
├─ Intn(int, int) int
└─ Numeric(int) string
regexFind(string, string) string
regexFindAll(string, string, int) []string
regexMatch(string, string) bool
regexQuoteMeta(string) string
regexReplaceAll(string, string, string) string
regexReplaceAllLiteral(string, string, string) string
regexSplit(string, string, int) []string
repeat(int, string) string
replace(string, string, string) string
rest(interface {}) []interface {}
reverse(interface {}) []interface {}
round(interface {}, int, ...float64) float64
semver(string) (*semver.Version, error)
├─ Compare(*semver.Version) int
├─ Equal(*semver.Version) bool
├─ GreaterThan(*semver.Version) bool
├─ GreaterThanEqual(*semver.Version) bool
├─ IncMajor() semver.Version
├─ IncMinor() semver.Version
├─ IncPatch() semver.Version
├─ LessThan(*semver.Version) bool
├─ LessThanEqual(*semver.Version) bool
├─ Major() uint64
├─ MarshalJSON() ([]uint8, error)
├─ MarshalText() ([]uint8, error)
├─ Metadata() string
├─ Minor() uint64
├─ Original() string
├─ Patch() uint64
├─ Prerelease() string
├─ Scan(interface {}) error
├─ SetMetadata(string) (semver.Version, error)
├─ SetPrerelease(string) (semver.Version, error)
├─ String() string
├─ UnmarshalJSON([]uint8) error
├─ UnmarshalText([]uint8) error
└─ Value() (driver.Value, error)
semverCompare(string, string) (bool, error)
seq(...int) string
set(map[string]interface {}, string, interface {}) map[string]interface {}
sha1sum(string) string
sha256sum(string) string
sha512sum(string) string
shuffle(string) string
slice(interface {}, ...interface {}) interface {}
snakecase(string) string
sortAlpha(interface {}) []string
split(string, string) map[string]string
splitList(string, string) []string
splitn(string, int, string) map[string]string
squote(...interface {}) string
sub(interface {}, interface {}) int64
subf(interface {}, ...interface {}) float64
substr(int, int, string) string
swapcase(string) string
ternary(interface {}, interface {}, bool) interface {}
time() time.Time
├─ Add(time.Duration) time.Time
├─ AddDate(int, int, int) time.Time
├─ After(time.Time) bool
├─ AppendFormat([]uint8, string) []uint8
├─ Before(time.Time) bool
├─ Clock() (int, int, int)
├─ Compare(time.Time) int
├─ Date() (int, time.Month, int)
├─ Day() int
├─ Equal(time.Time) bool
├─ Format(string) string
├─ GoString() string
├─ GobEncode() ([]uint8, error)
├─ Hour() int
├─ ISOWeek() (int, int)
├─ In(*time.Location) time.Time
├─ IsDST() bool
├─ IsZero() bool
├─ Local() time.Time
├┐ Location() *time.Location
│└─ String() string
├─ MarshalBinary() ([]uint8, error)
├─ MarshalJSON() ([]uint8, error)
├─ MarshalText() ([]uint8, error)
├─ Minute() int
├─ Month() time.Month
├─ Nanosecond() int
├─ Round(time.Duration) time.Time
├─ Second() int
├─ String() string
├─ Sub(time.Time) time.Duration
├─ Truncate(time.Duration) time.Time
├─ UTC() time.Time
├─ Unix() int64
├─ UnixMicro() int64
├─ UnixMilli() int64
├─ UnixNano() int64
├─ Weekday() time.Weekday
├─ Year() int
├─ YearDay() int
├─ Zone() (string, int)
└─ ZoneBounds() (time.Time, time.Time)
title(string) string
toDate(string, string) time.Time
toDecimal(interface {}) int64
toJson(interface {}) string
toPrettyJson(interface {}) string
toRawJson(interface {}) string
toString(interface {}) string
toStrings(interface {}) []string
trim(string) string
trimAll(string, string) string
trimPrefix(string, string) string
trimSuffix(string, string) string
trimall(string, string) string
trunc(int, string) string
tuple(...interface {}) []interface {}
typeIs(string, interface {}) bool
typeIsLike(string, interface {}) bool
typeOf(interface {}) string
uniq(interface {}) []interface {}
unixEpoch(time.Time) string
unset(map[string]interface {}, string) map[string]interface {}
until(int) []int
untilStep(int, int, int) []int
untitle(string) string
upper(string) string
urlJoin(map[string]interface {}) string
urlParse(string) map[string]interface {}
uuidv4() string
values(map[string]interface {}) []interface {}
without(interface {}, ...interface {}) []interface {}
wrap(int, string) string
wrapWith(int, string, string) string
```
