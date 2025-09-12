package math

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/spf13/cast"

	"github.com/rytsh/mugo/fstore"
)

func init() {
	fstore.AddStruct("math", Math{})
}

type Math struct{}

func (Math) Add(a any, v ...any) (json.Number, error) {
	return calcList(a, v, func(d1, d2 decimal.Decimal) decimal.Decimal {
		return d1.Add(d2)
	})
}

func (Math) Sub(a any, v ...any) (json.Number, error) {
	return calcList(a, v, func(d1, d2 decimal.Decimal) decimal.Decimal {
		return d1.Sub(d2)
	})
}

func (Math) Mul(a any, v ...any) (json.Number, error) {
	return calcList(a, v, func(d1, d2 decimal.Decimal) decimal.Decimal {
		return d1.Mul(d2)
	})
}

func (Math) Div(a any, v ...any) (json.Number, error) {
	return calcList(a, v, func(d1, d2 decimal.Decimal) decimal.Decimal {
		return d1.Div(d2)
	})
}

func (Math) Mod(a, b any) (json.Number, error) {
	return calc(a, b, func(d1, d2 decimal.Decimal) decimal.Decimal {
		return d1.Mod(d2)
	})
}

func (Math) Max(a any, v ...any) (json.Number, error) {
	return calcList(a, v, func(d1, d2 decimal.Decimal) decimal.Decimal {
		if d2.GreaterThan(d1) {
			return d2
		}
		return d1
	})
}

func (Math) Min(a any, v ...any) (json.Number, error) {
	return calcList(a, v, func(d1, d2 decimal.Decimal) decimal.Decimal {
		if d2.LessThan(d1) {
			return d2
		}
		return d1
	})
}

func (Math) Abs(a any) (json.Number, error) {
	aDec, err := convertDecimal(a)
	if err != nil {
		return "", err
	}

	return json.Number(aDec.Abs().String()), nil
}

func (Math) Ceil(a any) (json.Number, error) {
	aDec, err := convertDecimal(a)
	if err != nil {
		return "", err
	}

	return json.Number(aDec.Ceil().String()), nil
}

func (Math) Floor(a any) (json.Number, error) {
	aDec, err := convertDecimal(a)
	if err != nil {
		return "", err
	}

	return json.Number(aDec.Floor().String()), nil
}

func (Math) Pow(a, b any) (json.Number, error) {
	return calc(a, b, func(d1, d2 decimal.Decimal) decimal.Decimal {
		return d1.Pow(d2)
	})
}

func (Math) Equal(a, b any) (bool, error) {
	aDec, bDec, err := getValues(a, b)
	if err != nil {
		return false, err
	}

	return aDec.Equal(bDec), nil
}

func (Math) GreaterThan(a, b any) (bool, error) {
	aDec, bDec, err := getValues(a, b)
	if err != nil {
		return false, err
	}

	return aDec.GreaterThan(bDec), nil
}

func (Math) GreaterThanOrEqual(a, b any) (bool, error) {
	aDec, bDec, err := getValues(a, b)
	if err != nil {
		return false, err
	}

	return aDec.GreaterThanOrEqual(bDec), nil
}

func (Math) LessThan(a, b any) (bool, error) {
	aDec, bDec, err := getValues(a, b)
	if err != nil {
		return false, err
	}

	return aDec.LessThan(bDec), nil
}

func (Math) LessThanOrEqual(a, b any) (bool, error) {
	aDec, bDec, err := getValues(a, b)
	if err != nil {
		return false, err
	}

	return aDec.LessThanOrEqual(bDec), nil
}

func (Math) Sign(a any) (int, error) {
	aDec, err := convertDecimal(a)
	if err != nil {
		return 0, err
	}

	return aDec.Sign(), nil
}

func (Math) Round(a any, precision any) (json.Number, error) {
	aDec, err := convertDecimal(a)
	if err != nil {
		return "", err
	}

	precisionInt, err := convertInt32(precision)
	if err != nil {
		return "", err
	}

	return json.Number(aDec.Round(precisionInt).String()), nil
}

func (Math) RoundBankers(a any, precision any) (json.Number, error) {
	aDec, err := convertDecimal(a)
	if err != nil {
		return "", err
	}

	precisionInt, err := convertInt32(precision)
	if err != nil {
		return "", err
	}

	return json.Number(aDec.RoundBank(precisionInt).String()), nil
}

func (Math) RoundCash(a any, precision any) (json.Number, error) {
	aDec, err := convertDecimal(a)
	if err != nil {
		return "", err
	}

	precisionUint8, err := convertUint8(precision)
	if err != nil {
		return "", err
	}

	return json.Number(aDec.RoundCash(precisionUint8).String()), nil
}

func (Math) RoundCeil(a any, precision any) (json.Number, error) {
	aDec, err := convertDecimal(a)
	if err != nil {
		return "", err
	}

	precisionInt, err := convertInt32(precision)
	if err != nil {
		return "", err
	}

	return json.Number(aDec.RoundCeil(precisionInt).String()), nil
}

func (Math) RoundFloor(a any, precision any) (json.Number, error) {
	aDec, err := convertDecimal(a)
	if err != nil {
		return "", err
	}

	precisionInt, err := convertInt32(precision)
	if err != nil {
		return "", err
	}

	return json.Number(aDec.RoundFloor(precisionInt).String()), nil
}

func (Math) RoundUp(a any, precision any) (json.Number, error) {
	aDec, err := convertDecimal(a)
	if err != nil {
		return "", err
	}

	precisionInt, err := convertInt32(precision)
	if err != nil {
		return "", err
	}

	return json.Number(aDec.RoundUp(precisionInt).String()), nil
}

func (Math) RoundDown(a any, precision any) (json.Number, error) {
	aDec, err := convertDecimal(a)
	if err != nil {
		return "", err
	}

	precisionInt, err := convertInt32(precision)
	if err != nil {
		return "", err
	}

	return json.Number(aDec.RoundDown(precisionInt).String()), nil
}

func (Math) Truncate(a any, precision any) (json.Number, error) {
	aDec, err := convertDecimal(a)
	if err != nil {
		return "", err
	}

	precisionInt, err := convertInt32(precision)
	if err != nil {
		return "", err
	}

	return json.Number(aDec.Truncate(precisionInt).String()), nil
}

func getValues(a, b any) (decimal.Decimal, decimal.Decimal, error) {
	aDec, err := convertDecimal(a)
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, err
	}

	bDec, err := convertDecimal(b)
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, err
	}

	return aDec, bDec, nil
}

func calcList(a any, v []any, op func(decimal.Decimal, decimal.Decimal) decimal.Decimal) (json.Number, error) {
	aDec, err := convertDecimal(a)
	if err != nil {
		return "", err
	}

	for _, b := range v {
		bDec, err := convertDecimal(b)
		if err != nil {
			return "", err
		}

		aDec = op(aDec, bDec)
	}

	return json.Number(aDec.String()), nil
}

func calc(a, b any, op func(decimal.Decimal, decimal.Decimal) decimal.Decimal) (json.Number, error) {
	aDec, err := convertDecimal(a)
	if err != nil {
		return "", err
	}

	bDec, err := convertDecimal(b)
	if err != nil {
		return "", err
	}

	return json.Number(op(aDec, bDec).String()), nil
}

func convertDecimal(a any) (decimal.Decimal, error) {
	switch v := a.(type) {
	case int:
		return decimal.NewFromInt(int64(v)), nil
	case int8:
		return decimal.NewFromInt(int64(v)), nil
	case int32:
		return decimal.NewFromInt(int64(v)), nil
	case int64:
		return decimal.NewFromInt(v), nil
	case uint:
		return decimal.NewFromUint64(uint64(v)), nil
	case uint8:
		return decimal.NewFromUint64(uint64(v)), nil
	case uint32:
		return decimal.NewFromUint64(uint64(v)), nil
	case uint64:
		return decimal.NewFromUint64(v), nil
	case float32:
		return decimal.NewFromFloat(float64(v)), nil
	case float64:
		return decimal.NewFromFloat(v), nil
	case decimal.Decimal:
		return v, nil
	case string:
		return decimal.NewFromString(v)
	case json.Number:
		return decimal.NewFromString(string(v))
	default:
		return decimal.NewFromInt(0), fmt.Errorf("unsupported type %T", a)
	}
}

func convertInt32(a any) (int32, error) {
	switch v := a.(type) {
	case decimal.Decimal:
		return int32(v.IntPart()), nil
	default:
		return cast.ToInt32E(v)
	}
}

func convertUint8(a any) (uint8, error) {
	switch v := a.(type) {
	case decimal.Decimal:
		return uint8(v.IntPart()), nil
	default:
		return cast.ToUint8E(v)
	}
}
