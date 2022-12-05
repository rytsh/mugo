package humanize

import (
	"github.com/dustin/go-humanize"
)

func FuncMap() map[string]interface{} {
	fMap := map[string]interface{}{
		"bigBytes":         humanize.BigBytes,
		"bigComma":         humanize.BigComma,
		"bigCommaf":        humanize.BigCommaf,
		"bigIBytes":        humanize.BigIBytes,
		"bytes":            humanize.Bytes,
		"comma":            humanize.Comma,
		"commaf":           humanize.Commaf,
		"commafWithDigits": humanize.CommafWithDigits,
		// "computeSI":        humanize.ComputeSI,
		"customRelTime":  humanize.CustomRelTime,
		"formatFloat":    humanize.FormatFloat,
		"formatInteger":  humanize.FormatInteger,
		"ftoa":           humanize.Ftoa,
		"ftoaWithDigits": humanize.FtoaWithDigits,
		"iBytes":         humanize.IBytes,
		"ordinal":        humanize.Ordinal,
		"parseBigBytes":  humanize.ParseBigBytes,
		"parseBytes":     humanize.ParseBytes,
		// "parseSI":        humanize.ParseSI,
		"relTime":      humanize.RelTime,
		"sI":           humanize.SI,
		"sIWithDigits": humanize.SIWithDigits,
		"time":         humanize.Time,
	}

	return fMap
}
