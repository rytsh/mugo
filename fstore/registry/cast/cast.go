package cast

import (
	"time"

	"github.com/spf13/cast"

	"github.com/rytsh/mugo/fstore"
)

func init() {
	fstore.AddStruct("cast", Cast{})
}

// Cast is a collection of cast github.com/spf13/cast functions.
type Cast struct{}

func (Cast) StringToDate(s string) (time.Time, error) {
	return cast.StringToDate(s)
}
func (Cast) StringToDateInDefaultLocation(s string, location *time.Location) (time.Time, error) {
	return cast.StringToDateInDefaultLocation(s, location)
}
func (Cast) ToBool(i any) bool {
	return cast.ToBool(i)
}
func (Cast) ToBoolE(i any) (bool, error) {
	return cast.ToBoolE(i)
}
func (Cast) ToBoolSlice(i any) []bool {
	return cast.ToBoolSlice(i)
}
func (Cast) ToBoolSliceE(i any) ([]bool, error) {
	return cast.ToBoolSliceE(i)
}
func (Cast) ToDuration(i any) time.Duration {
	return cast.ToDuration(i)
}
func (Cast) ToDurationE(i any) (d time.Duration, err error) {
	return cast.ToDurationE(i)
}
func (Cast) ToDurationSlice(i any) []time.Duration {
	return cast.ToDurationSlice(i)
}
func (Cast) ToDurationSliceE(i any) ([]time.Duration, error) {
	return cast.ToDurationSliceE(i)
}
func (Cast) ToFloat32(i any) float32 {
	return cast.ToFloat32(i)
}
func (Cast) ToFloat32E(i any) (float32, error) {
	return cast.ToFloat32E(i)
}
func (Cast) ToFloat64(i any) float64 {
	return cast.ToFloat64(i)
}
func (Cast) ToFloat64E(i any) (float64, error) {
	return cast.ToFloat64E(i)
}
func (Cast) ToInt(i any) int {
	return cast.ToInt(i)
}
func (Cast) ToInt16(i any) int16 {
	return cast.ToInt16(i)
}
func (Cast) ToInt16E(i any) (int16, error) {
	return cast.ToInt16E(i)
}
func (Cast) ToInt32(i any) int32 {
	return cast.ToInt32(i)
}
func (Cast) ToInt32E(i any) (int32, error) {
	return cast.ToInt32E(i)
}
func (Cast) ToInt64(i any) int64 {
	return cast.ToInt64(i)
}
func (Cast) ToInt64E(i any) (int64, error) {
	return cast.ToInt64E(i)
}
func (Cast) ToInt8(i any) int8 {
	return cast.ToInt8(i)
}
func (Cast) ToInt8E(i any) (int8, error) {
	return cast.ToInt8E(i)
}
func (Cast) ToIntE(i any) (int, error) {
	return cast.ToIntE(i)
}
func (Cast) ToIntSlice(i any) []int {
	return cast.ToIntSlice(i)
}
func (Cast) ToIntSliceE(i any) ([]int, error) {
	return cast.ToIntSliceE(i)
}
func (Cast) ToSlice(i any) []any {
	return cast.ToSlice(i)
}
func (Cast) ToSliceE(i any) ([]any, error) {
	return cast.ToSliceE(i)
}
func (Cast) ToString(i any) string {
	return cast.ToString(i)
}
func (Cast) ToStringE(i any) (string, error) {
	return cast.ToStringE(i)
}
func (Cast) ToStringMap(i any) map[string]any {
	return cast.ToStringMap(i)
}
func (Cast) ToStringMapBool(i any) map[string]bool {
	return cast.ToStringMapBool(i)
}
func (Cast) ToStringMapBoolE(i any) (map[string]bool, error) {
	return cast.ToStringMapBoolE(i)
}
func (Cast) ToStringMapE(i any) (map[string]any, error) {
	return cast.ToStringMapE(i)
}
func (Cast) ToStringMapInt(i any) map[string]int {
	return cast.ToStringMapInt(i)
}
func (Cast) ToStringMapInt64(i any) map[string]int64 {
	return cast.ToStringMapInt64(i)
}
func (Cast) ToStringMapInt64E(i any) (map[string]int64, error) {
	return cast.ToStringMapInt64E(i)
}
func (Cast) ToStringMapIntE(i any) (map[string]int, error) {
	return cast.ToStringMapIntE(i)
}
func (Cast) ToStringMapString(i any) map[string]string {
	return cast.ToStringMapString(i)
}
func (Cast) ToStringMapStringE(i any) (map[string]string, error) {
	return cast.ToStringMapStringE(i)
}
func (Cast) ToStringMapStringSlice(i any) map[string][]string {
	return cast.ToStringMapStringSlice(i)
}
func (Cast) ToStringMapStringSliceE(i any) (map[string][]string, error) {
	return cast.ToStringMapStringSliceE(i)
}
func (Cast) ToStringSlice(i any) []string {
	return cast.ToStringSlice(i)
}
func (Cast) ToStringSliceE(i any) ([]string, error) {
	return cast.ToStringSliceE(i)
}
func (Cast) ToTime(i any) time.Time {
	return cast.ToTime(i)
}
func (Cast) ToTimeE(i any) (tim time.Time, err error) {
	return cast.ToTimeE(i)
}
func (Cast) ToTimeInDefaultLocation(i any, location *time.Location) time.Time {
	return cast.ToTimeInDefaultLocation(i, location)
}
func (Cast) ToTimeInDefaultLocationE(i any, location *time.Location) (tim time.Time, err error) {
	return cast.ToTimeInDefaultLocationE(i, location)
}
func (Cast) ToUint(i any) uint {
	return cast.ToUint(i)
}
func (Cast) ToUint16(i any) uint16 {
	return cast.ToUint16(i)
}
func (Cast) ToUint16E(i any) (uint16, error) {
	return cast.ToUint16E(i)
}
func (Cast) ToUint32(i any) uint32 {
	return cast.ToUint32(i)
}
func (Cast) ToUint32E(i any) (uint32, error) {
	return cast.ToUint32E(i)
}
func (Cast) ToUint64(i any) uint64 {
	return cast.ToUint64(i)
}
func (Cast) ToUint64E(i any) (uint64, error) {
	return cast.ToUint64E(i)
}
func (Cast) ToUint8(i any) uint8 {
	return cast.ToUint8(i)
}
func (Cast) ToUint8E(i any) (uint8, error) {
	return cast.ToUint8E(i)
}
func (Cast) ToUintE(i any) (uint, error) {
	return cast.ToUintE(i)
}
