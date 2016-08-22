package qu_EC

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type qu_EC struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
	percentSuffix          []byte
	perMille               []byte
	timeSeparator          []byte
	inifinity              []byte
	currencies             [][]byte // idx = enum of currency code
	currencyPositivePrefix []byte
	currencyNegativePrefix []byte
	monthsAbbreviated      [][]byte
	monthsNarrow           [][]byte
	monthsWide             [][]byte
	daysAbbreviated        [][]byte
	daysNarrow             [][]byte
	daysShort              [][]byte
	daysWide               [][]byte
	periodsAbbreviated     [][]byte
	periodsNarrow          [][]byte
	periodsShort           [][]byte
	periodsWide            [][]byte
	erasAbbreviated        [][]byte
	erasNarrow             [][]byte
	erasWide               [][]byte
	timezones              map[string][]byte
}

// New returns a new instance of translator for the 'qu_EC' locale
func New() locales.Translator {
	return &qu_EC{
		locale:                 "qu_EC",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44, 0x20}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e, 0x20}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c, 0x20}, {0x41, 0x4d, 0x44, 0x20}, {0x41, 0x4e, 0x47, 0x20}, {0x41, 0x4f, 0x41, 0x20}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53, 0x20}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44, 0x20}, {0x41, 0x57, 0x47, 0x20}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e, 0x20}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d, 0x20}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44, 0x20}, {0x42, 0x44, 0x54, 0x20}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e, 0x20}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44, 0x20}, {0x42, 0x49, 0x46, 0x20}, {0x42, 0x4d, 0x44, 0x20}, {0x42, 0x4e, 0x44, 0x20}, {0x42, 0x4f, 0x42, 0x20}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x42, 0x52, 0x4c, 0x20}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44, 0x20}, {0x42, 0x54, 0x4e, 0x20}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50, 0x20}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52, 0x20}, {0x42, 0x5a, 0x44, 0x20}, {0x43, 0x41, 0x44, 0x20}, {0x43, 0x44, 0x46, 0x20}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46, 0x20}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50, 0x20}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0x59, 0x20}, {0x43, 0x4f, 0x50, 0x20}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43, 0x20}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43, 0x20}, {0x43, 0x55, 0x50, 0x20}, {0x43, 0x56, 0x45, 0x20}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b, 0x20}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46, 0x20}, {0x44, 0x4b, 0x4b, 0x20}, {0x44, 0x4f, 0x50, 0x20}, {0x44, 0x5a, 0x44, 0x20}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50, 0x20}, {0x45, 0x52, 0x4e, 0x20}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42, 0x20}, {0x45, 0x55, 0x52, 0x20}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44, 0x20}, {0x46, 0x4b, 0x50, 0x20}, {0x46, 0x52, 0x46, 0x20}, {0x47, 0x42, 0x50, 0x20}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c, 0x20}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53, 0x20}, {0x47, 0x49, 0x50, 0x20}, {0x47, 0x4d, 0x44, 0x20}, {0x47, 0x4e, 0x46, 0x20}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51, 0x20}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44, 0x20}, {0x48, 0x4b, 0x44, 0x20}, {0x48, 0x4e, 0x4c, 0x20}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b, 0x20}, {0x48, 0x54, 0x47, 0x20}, {0x48, 0x55, 0x46, 0x20}, {0x49, 0x44, 0x52, 0x20}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0x49, 0x4c, 0x53, 0x20}, {0x49, 0x4e, 0x52, 0x20}, {0x49, 0x51, 0x44, 0x20}, {0x49, 0x52, 0x52, 0x20}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b, 0x20}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44, 0x20}, {0x4a, 0x4f, 0x44, 0x20}, {0x4a, 0x50, 0x59, 0x20}, {0x4b, 0x45, 0x53, 0x20}, {0x4b, 0x47, 0x53, 0x20}, {0x4b, 0x48, 0x52, 0x20}, {0x4b, 0x4d, 0x46, 0x20}, {0x4b, 0x50, 0x57, 0x20}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0x4b, 0x52, 0x57, 0x20}, {0x4b, 0x57, 0x44, 0x20}, {0x4b, 0x59, 0x44, 0x20}, {0x4b, 0x5a, 0x54, 0x20}, {0x4c, 0x41, 0x4b, 0x20}, {0x4c, 0x42, 0x50, 0x20}, {0x4c, 0x4b, 0x52, 0x20}, {0x4c, 0x52, 0x44, 0x20}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c, 0x20}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c, 0x20}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44, 0x20}, {0x4d, 0x41, 0x44, 0x20}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c, 0x20}, {0x4d, 0x47, 0x41, 0x20}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44, 0x20}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b, 0x20}, {0x4d, 0x4e, 0x54, 0x20}, {0x4d, 0x4f, 0x50, 0x20}, {0x4d, 0x52, 0x4f, 0x20}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52, 0x20}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52, 0x20}, {0x4d, 0x57, 0x4b, 0x20}, {0x4d, 0x58, 0x4e, 0x20}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52, 0x20}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e, 0x20}, {0x4e, 0x41, 0x44, 0x20}, {0x4e, 0x47, 0x4e, 0x20}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f, 0x20}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b, 0x20}, {0x4e, 0x50, 0x52, 0x20}, {0x4e, 0x5a, 0x44, 0x20}, {0x4f, 0x4d, 0x52, 0x20}, {0x50, 0x41, 0x42, 0x20}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b, 0x20}, {0x50, 0x48, 0x50, 0x20}, {0x50, 0x4b, 0x52, 0x20}, {0x50, 0x4c, 0x4e, 0x20}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47, 0x20}, {0x51, 0x41, 0x52, 0x20}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e, 0x20}, {0x52, 0x53, 0x44, 0x20}, {0x52, 0x55, 0x42, 0x20}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46, 0x20}, {0x53, 0x41, 0x52, 0x20}, {0x53, 0x42, 0x44, 0x20}, {0x53, 0x43, 0x52, 0x20}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47, 0x20}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b, 0x20}, {0x53, 0x47, 0x44, 0x20}, {0x53, 0x48, 0x50, 0x20}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c, 0x20}, {0x53, 0x4f, 0x53, 0x20}, {0x53, 0x52, 0x44, 0x20}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50, 0x20}, {0x53, 0x54, 0x44, 0x20}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50, 0x20}, {0x53, 0x5a, 0x4c, 0x20}, {0x54, 0x48, 0x42, 0x20}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53, 0x20}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54, 0x20}, {0x54, 0x4e, 0x44, 0x20}, {0x54, 0x4f, 0x50, 0x20}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59, 0x20}, {0x54, 0x54, 0x44, 0x20}, {0x54, 0x57, 0x44, 0x20}, {0x54, 0x5a, 0x53, 0x20}, {0x55, 0x41, 0x48, 0x20}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58, 0x20}, {0x24}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55, 0x20}, {0x55, 0x5a, 0x53, 0x20}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46, 0x20}, {0x56, 0x4e, 0x44, 0x20}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56, 0x20}, {0x57, 0x53, 0x54, 0x20}, {0x58, 0x41, 0x46, 0x20}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x58, 0x43, 0x44, 0x20}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x58, 0x4f, 0x46, 0x20}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46, 0x20}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52, 0x20}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52, 0x20}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57, 0x20}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		percentSuffix:          []byte{0xc2, 0xa0},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x51, 0x75, 0x6c}, {0x48, 0x61, 0x74}, {0x50, 0x61, 0x75}, {0x41, 0x79, 0x72}, {0x41, 0x79, 0x6d}, {0x49, 0x6e, 0x74}, {0x41, 0x6e, 0x74}, {0x51, 0x68, 0x61}, {0x55, 0x6d, 0x61}, {0x4b, 0x61, 0x6e}, {0x41, 0x79, 0x61}, {0x4b, 0x61, 0x70}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x51, 0x75, 0x6c, 0x6c, 0x61, 0x20, 0x70, 0x75, 0x71, 0x75, 0x79}, {0x48, 0x61, 0x74, 0x75, 0x6e, 0x20, 0x70, 0x75, 0x71, 0x75, 0x79}, {0x50, 0x61, 0x75, 0x71, 0x61, 0x72, 0x20, 0x77, 0x61, 0x72, 0x61, 0x79}, {0x41, 0x79, 0x72, 0x69, 0x77, 0x61}, {0x41, 0x79, 0x6d, 0x75, 0x72, 0x61, 0x79}, {0x49, 0x6e, 0x74, 0x69, 0x20, 0x72, 0x61, 0x79, 0x6d, 0x69}, {0x41, 0x6e, 0x74, 0x61, 0x20, 0x53, 0x69, 0x74, 0x77, 0x61}, {0x51, 0x68, 0x61, 0x70, 0x61, 0x71, 0x20, 0x53, 0x69, 0x74, 0x77, 0x61}, {0x55, 0x6d, 0x61, 0x20, 0x72, 0x61, 0x79, 0x6d, 0x69}, {0x4b, 0x61, 0x6e, 0x74, 0x61, 0x72, 0x61, 0x79}, {0x41, 0x79, 0x61, 0x6d, 0x61, 0x72, 0x71, 0xca, 0xbc, 0x61}, {0x4b, 0x61, 0x70, 0x61, 0x71, 0x20, 0x52, 0x61, 0x79, 0x6d, 0x69}},
		daysAbbreviated:        [][]uint8{{0x44, 0x6f, 0x6d}, {0x4c, 0x75, 0x6e}, {0x4d, 0x61, 0x72}, {0x4d, 0x69, 0xc3, 0xa9}, {0x4a, 0x75, 0x65}, {0x56, 0x69, 0x65}, {0x53, 0x61, 0x62}},
		daysNarrow:             [][]uint8{{0x44}, {0x4c}, {0x4d}, {0x58}, {0x4a}, {0x56}, {0x53}},
		daysWide:               [][]uint8{{0x44, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x6f}, {0x4c, 0x75, 0x6e, 0x65, 0x73}, {0x4d, 0x61, 0x72, 0x74, 0x65, 0x73}, {0x4d, 0x69, 0xc3, 0xa9, 0x72, 0x63, 0x6f, 0x6c, 0x65, 0x73}, {0x4a, 0x75, 0x65, 0x76, 0x65, 0x73}, {0x56, 0x69, 0x65, 0x72, 0x6e, 0x65, 0x73}, {0x53, 0xc3, 0xa1, 0x62, 0x61, 0x64, 0x6f}},
		periodsAbbreviated:     [][]uint8{{0x61, 0x2e, 0x6d, 0x2e}, {0x70, 0x2e, 0x6d, 0x2e}},
		periodsWide:            [][]uint8{{0x61, 0x2e, 0x6d, 0x2e}, {0x70, 0x2e, 0x6d, 0x2e}},
		erasAbbreviated:        [][]uint8{[]uint8(nil), []uint8(nil)},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{[]uint8(nil), []uint8(nil)},
		timezones:              map[string][]uint8{"∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "ART": {0x41, 0x52, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "VET": {0x56, 0x45, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "ECT": {0x45, 0x43, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "EDT": {0x45, 0x44, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "PDT": {0x50, 0x44, 0x54}, "GYT": {0x47, 0x59, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "HAT": {0x48, 0x41, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "CDT": {0x43, 0x44, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "GFT": {0x47, 0x46, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "COT": {0x43, 0x4f, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "EST": {0x45, 0x53, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "WIB": {0x57, 0x49, 0x42}, "CST": {0x43, 0x53, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "SRT": {0x53, 0x52, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "CAT": {0x43, 0x41, 0x54}, "MST": {0x4d, 0x53, 0x54}, "AST": {0x41, 0x53, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "WIT": {0x57, 0x49, 0x54}, "BT": {0x42, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "IST": {0x49, 0x53, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "UYT": {0x55, 0x59, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "WAT": {0x57, 0x41, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "EAT": {0x45, 0x41, 0x54}, "SGT": {0x53, 0x47, 0x54}, "ADT": {0x41, 0x44, 0x54}, "JST": {0x4a, 0x53, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "PST": {0x50, 0x53, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}},
	}
}

// Locale returns the current translators string locale
func (qu *qu_EC) Locale() string {
	return qu.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'qu_EC'
func (qu *qu_EC) PluralsCardinal() []locales.PluralRule {
	return qu.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'qu_EC'
func (qu *qu_EC) PluralsOrdinal() []locales.PluralRule {
	return qu.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'qu_EC'
func (qu *qu_EC) PluralsRange() []locales.PluralRule {
	return qu.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'qu_EC'
func (qu *qu_EC) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'qu_EC'
func (qu *qu_EC) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'qu_EC'
func (qu *qu_EC) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (qu *qu_EC) MonthAbbreviated(month time.Month) []byte {
	return qu.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (qu *qu_EC) MonthsAbbreviated() [][]byte {
	return qu.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (qu *qu_EC) MonthNarrow(month time.Month) []byte {
	return qu.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (qu *qu_EC) MonthsNarrow() [][]byte {
	return qu.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (qu *qu_EC) MonthWide(month time.Month) []byte {
	return qu.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (qu *qu_EC) MonthsWide() [][]byte {
	return qu.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (qu *qu_EC) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return qu.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (qu *qu_EC) WeekdaysAbbreviated() [][]byte {
	return qu.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (qu *qu_EC) WeekdayNarrow(weekday time.Weekday) []byte {
	return qu.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (qu *qu_EC) WeekdaysNarrow() [][]byte {
	return qu.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (qu *qu_EC) WeekdayShort(weekday time.Weekday) []byte {
	return qu.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (qu *qu_EC) WeekdaysShort() [][]byte {
	return qu.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (qu *qu_EC) WeekdayWide(weekday time.Weekday) []byte {
	return qu.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (qu *qu_EC) WeekdaysWide() [][]byte {
	return qu.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'qu_EC' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (qu *qu_EC) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(qu.decimal) + len(qu.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, qu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, qu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, qu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'qu_EC' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (qu *qu_EC) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(qu.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, qu.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, qu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, qu.percentSuffix...)

	b = append(b, qu.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'qu_EC'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (qu *qu_EC) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := qu.currencies[currency]
	l := len(s) + len(qu.decimal) + len(qu.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, qu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, qu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(qu.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, qu.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, qu.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, qu.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'qu_EC'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (qu *qu_EC) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := qu.currencies[currency]
	l := len(s) + len(qu.decimal) + len(qu.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, qu.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, qu.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(qu.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, qu.currencyNegativePrefix[j])
		}

		b = append(b, qu.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(qu.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, qu.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, qu.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'qu_EC'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (qu *qu_EC) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'qu_EC'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (qu *qu_EC) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'qu_EC'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (qu *qu_EC) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'qu_EC'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (qu *qu_EC) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, qu.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, qu.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'qu_EC'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (qu *qu_EC) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, qu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'qu_EC'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (qu *qu_EC) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, qu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, qu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'qu_EC'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (qu *qu_EC) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, qu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, qu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'qu_EC'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (qu *qu_EC) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, qu.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, qu.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := qu.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
