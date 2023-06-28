package funcs

import (
	"time"

	"github.com/spf13/cast"

	"github.com/rytsh/mugo/pkg/fstore/registry"
)

func init() {
	registry.AddGroup("cast", registry.ReturnWithFn(Cast{}))
}

// Cast is a collection of cast github.com/spf13/cast functions.
type Cast struct{}

func (Cast) StringToDate(s string) (time.Time, error) {
	return cast.StringToDate(s)
}
func (Cast) StringToDateInDefaultLocation(s string, location *time.Location) (time.Time, error) {
	return cast.StringToDateInDefaultLocation(s, location)
}
func (Cast) ToBool(i interface{}) bool {
	return cast.ToBool(i)
}
func (Cast) ToBoolE(i interface{}) (bool, error) {
	return cast.ToBoolE(i)
}
func (Cast) ToBoolSlice(i interface{}) []bool {
	return cast.ToBoolSlice(i)
}
func (Cast) ToBoolSliceE(i interface{}) ([]bool, error) {
	return cast.ToBoolSliceE(i)
}
func (Cast) ToDuration(i interface{}) time.Duration {
	return cast.ToDuration(i)
}
func (Cast) ToDurationE(i interface{}) (d time.Duration, err error) {
	return cast.ToDurationE(i)
}
func (Cast) ToDurationSlice(i interface{}) []time.Duration {
	return cast.ToDurationSlice(i)
}
func (Cast) ToDurationSliceE(i interface{}) ([]time.Duration, error) {
	return cast.ToDurationSliceE(i)
}
func (Cast) ToFloat32(i interface{}) float32 {
	return cast.ToFloat32(i)
}
func (Cast) ToFloat32E(i interface{}) (float32, error) {
	return cast.ToFloat32E(i)
}
func (Cast) ToFloat64(i interface{}) float64 {
	return cast.ToFloat64(i)
}
func (Cast) ToFloat64E(i interface{}) (float64, error) {
	return cast.ToFloat64E(i)
}
func (Cast) ToInt(i interface{}) int {
	return cast.ToInt(i)
}
func (Cast) ToInt16(i interface{}) int16 {
	return cast.ToInt16(i)
}
func (Cast) ToInt16E(i interface{}) (int16, error) {
	return cast.ToInt16E(i)
}
func (Cast) ToInt32(i interface{}) int32 {
	return cast.ToInt32(i)
}
func (Cast) ToInt32E(i interface{}) (int32, error) {
	return cast.ToInt32E(i)
}
func (Cast) ToInt64(i interface{}) int64 {
	return cast.ToInt64(i)
}
func (Cast) ToInt64E(i interface{}) (int64, error) {
	return cast.ToInt64E(i)
}
func (Cast) ToInt8(i interface{}) int8 {
	return cast.ToInt8(i)
}
func (Cast) ToInt8E(i interface{}) (int8, error) {
	return cast.ToInt8E(i)
}
func (Cast) ToIntE(i interface{}) (int, error) {
	return cast.ToIntE(i)
}
func (Cast) ToIntSlice(i interface{}) []int {
	return cast.ToIntSlice(i)
}
func (Cast) ToIntSliceE(i interface{}) ([]int, error) {
	return cast.ToIntSliceE(i)
}
func (Cast) ToSlice(i interface{}) []interface{} {
	return cast.ToSlice(i)
}
func (Cast) ToSliceE(i interface{}) ([]interface{}, error) {
	return cast.ToSliceE(i)
}
func (Cast) ToString(i interface{}) string {
	return cast.ToString(i)
}
func (Cast) ToStringE(i interface{}) (string, error) {
	return cast.ToStringE(i)
}
func (Cast) ToStringMap(i interface{}) map[string]interface{} {
	return cast.ToStringMap(i)
}
func (Cast) ToStringMapBool(i interface{}) map[string]bool {
	return cast.ToStringMapBool(i)
}
func (Cast) ToStringMapBoolE(i interface{}) (map[string]bool, error) {
	return cast.ToStringMapBoolE(i)
}
func (Cast) ToStringMapE(i interface{}) (map[string]interface{}, error) {
	return cast.ToStringMapE(i)
}
func (Cast) ToStringMapInt(i interface{}) map[string]int {
	return cast.ToStringMapInt(i)
}
func (Cast) ToStringMapInt64(i interface{}) map[string]int64 {
	return cast.ToStringMapInt64(i)
}
func (Cast) ToStringMapInt64E(i interface{}) (map[string]int64, error) {
	return cast.ToStringMapInt64E(i)
}
func (Cast) ToStringMapIntE(i interface{}) (map[string]int, error) {
	return cast.ToStringMapIntE(i)
}
func (Cast) ToStringMapString(i interface{}) map[string]string {
	return cast.ToStringMapString(i)
}
func (Cast) ToStringMapStringE(i interface{}) (map[string]string, error) {
	return cast.ToStringMapStringE(i)
}
func (Cast) ToStringMapStringSlice(i interface{}) map[string][]string {
	return cast.ToStringMapStringSlice(i)
}
func (Cast) ToStringMapStringSliceE(i interface{}) (map[string][]string, error) {
	return cast.ToStringMapStringSliceE(i)
}
func (Cast) ToStringSlice(i interface{}) []string {
	return cast.ToStringSlice(i)
}
func (Cast) ToStringSliceE(i interface{}) ([]string, error) {
	return cast.ToStringSliceE(i)
}
func (Cast) ToTime(i interface{}) time.Time {
	return cast.ToTime(i)
}
func (Cast) ToTimeE(i interface{}) (tim time.Time, err error) {
	return cast.ToTimeE(i)
}
func (Cast) ToTimeInDefaultLocation(i interface{}, location *time.Location) time.Time {
	return cast.ToTimeInDefaultLocation(i, location)
}
func (Cast) ToTimeInDefaultLocationE(i interface{}, location *time.Location) (tim time.Time, err error) {
	return cast.ToTimeInDefaultLocationE(i, location)
}
func (Cast) ToUint(i interface{}) uint {
	return cast.ToUint(i)
}
func (Cast) ToUint16(i interface{}) uint16 {
	return cast.ToUint16(i)
}
func (Cast) ToUint16E(i interface{}) (uint16, error) {
	return cast.ToUint16E(i)
}
func (Cast) ToUint32(i interface{}) uint32 {
	return cast.ToUint32(i)
}
func (Cast) ToUint32E(i interface{}) (uint32, error) {
	return cast.ToUint32E(i)
}
func (Cast) ToUint64(i interface{}) uint64 {
	return cast.ToUint64(i)
}
func (Cast) ToUint64E(i interface{}) (uint64, error) {
	return cast.ToUint64E(i)
}
func (Cast) ToUint8(i interface{}) uint8 {
	return cast.ToUint8(i)
}
func (Cast) ToUint8E(i interface{}) (uint8, error) {
	return cast.ToUint8E(i)
}
func (Cast) ToUintE(i interface{}) (uint, error) {
	return cast.ToUintE(i)
}
