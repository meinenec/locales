package shi_Tfng_MA

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type shi_Tfng_MA struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            []byte
	group              []byte
	minus              []byte
	percent            []byte
	perMille           []byte
	timeSeparator      []byte
	inifinity          []byte
	currencies         [][]byte // idx = enum of currency code
	monthsAbbreviated  [][]byte
	monthsNarrow       [][]byte
	monthsWide         [][]byte
	daysAbbreviated    [][]byte
	daysNarrow         [][]byte
	daysShort          [][]byte
	daysWide           [][]byte
	periodsAbbreviated [][]byte
	periodsNarrow      [][]byte
	periodsShort       [][]byte
	periodsWide        [][]byte
	erasAbbreviated    [][]byte
	erasNarrow         [][]byte
	erasWide           [][]byte
	timezones          map[string][]byte
}

// New returns a new instance of translator for the 'shi_Tfng_MA' locale
func New() locales.Translator {
	return &shi_Tfng_MA{
		locale:             "shi_Tfng_MA",
		pluralsCardinal:    []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		decimal:            []byte{0x2c},
		group:              []byte{0xc2, 0xa0},
		minus:              []byte{},
		percent:            []byte{},
		perMille:           []byte{},
		timeSeparator:      []byte{0x3a},
		currencies:         [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44, 0x20}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e, 0x20}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c, 0x20}, {0x41, 0x4d, 0x44, 0x20}, {0x41, 0x4e, 0x47, 0x20}, {0x41, 0x4f, 0x41, 0x20}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53, 0x20}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44, 0x20}, {0x41, 0x57, 0x47, 0x20}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e, 0x20}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d, 0x20}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44, 0x20}, {0x42, 0x44, 0x54, 0x20}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e, 0x20}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44, 0x20}, {0x42, 0x49, 0x46, 0x20}, {0x42, 0x4d, 0x44, 0x20}, {0x42, 0x4e, 0x44, 0x20}, {0x42, 0x4f, 0x42, 0x20}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x42, 0x52, 0x4c, 0x20}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44, 0x20}, {0x42, 0x54, 0x4e, 0x20}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50, 0x20}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52, 0x20}, {0x42, 0x5a, 0x44, 0x20}, {0x43, 0x41, 0x44, 0x20}, {0x43, 0x44, 0x46, 0x20}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46, 0x20}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50, 0x20}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0x59, 0x20}, {0x43, 0x4f, 0x50, 0x20}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43, 0x20}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43, 0x20}, {0x43, 0x55, 0x50, 0x20}, {0x43, 0x56, 0x45, 0x20}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b, 0x20}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46, 0x20}, {0x44, 0x4b, 0x4b, 0x20}, {0x44, 0x4f, 0x50, 0x20}, {0x44, 0x5a, 0x44, 0x20}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50, 0x20}, {0x45, 0x52, 0x4e, 0x20}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42, 0x20}, {0x45, 0x55, 0x52, 0x20}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44, 0x20}, {0x46, 0x4b, 0x50, 0x20}, {0x46, 0x52, 0x46, 0x20}, {0x47, 0x42, 0x50, 0x20}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c, 0x20}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53, 0x20}, {0x47, 0x49, 0x50, 0x20}, {0x47, 0x4d, 0x44, 0x20}, {0x47, 0x4e, 0x46, 0x20}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51, 0x20}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44, 0x20}, {0x48, 0x4b, 0x44, 0x20}, {0x48, 0x4e, 0x4c, 0x20}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b, 0x20}, {0x48, 0x54, 0x47, 0x20}, {0x48, 0x55, 0x46, 0x20}, {0x49, 0x44, 0x52, 0x20}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0x49, 0x4c, 0x53, 0x20}, {0x49, 0x4e, 0x52, 0x20}, {0x49, 0x51, 0x44, 0x20}, {0x49, 0x52, 0x52, 0x20}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b, 0x20}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44, 0x20}, {0x4a, 0x4f, 0x44, 0x20}, {0x4a, 0x50, 0x59, 0x20}, {0x4b, 0x45, 0x53, 0x20}, {0x4b, 0x47, 0x53, 0x20}, {0x4b, 0x48, 0x52, 0x20}, {0x4b, 0x4d, 0x46, 0x20}, {0x4b, 0x50, 0x57, 0x20}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0x4b, 0x52, 0x57, 0x20}, {0x4b, 0x57, 0x44, 0x20}, {0x4b, 0x59, 0x44, 0x20}, {0x4b, 0x5a, 0x54, 0x20}, {0x4c, 0x41, 0x4b, 0x20}, {0x4c, 0x42, 0x50, 0x20}, {0x4c, 0x4b, 0x52, 0x20}, {0x4c, 0x52, 0x44, 0x20}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c, 0x20}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c, 0x20}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44, 0x20}, {0x4d, 0x41, 0x44, 0x20}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c, 0x20}, {0x4d, 0x47, 0x41, 0x20}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44, 0x20}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b, 0x20}, {0x4d, 0x4e, 0x54, 0x20}, {0x4d, 0x4f, 0x50, 0x20}, {0x4d, 0x52, 0x4f, 0x20}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52, 0x20}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52, 0x20}, {0x4d, 0x57, 0x4b, 0x20}, {0x4d, 0x58, 0x4e, 0x20}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52, 0x20}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e, 0x20}, {0x4e, 0x41, 0x44, 0x20}, {0x4e, 0x47, 0x4e, 0x20}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f, 0x20}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b, 0x20}, {0x4e, 0x50, 0x52, 0x20}, {0x4e, 0x5a, 0x44, 0x20}, {0x4f, 0x4d, 0x52, 0x20}, {0x50, 0x41, 0x42, 0x20}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e, 0x20}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b, 0x20}, {0x50, 0x48, 0x50, 0x20}, {0x50, 0x4b, 0x52, 0x20}, {0x50, 0x4c, 0x4e, 0x20}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47, 0x20}, {0x51, 0x41, 0x52, 0x20}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e, 0x20}, {0x52, 0x53, 0x44, 0x20}, {0x52, 0x55, 0x42, 0x20}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46, 0x20}, {0x53, 0x41, 0x52, 0x20}, {0x53, 0x42, 0x44, 0x20}, {0x53, 0x43, 0x52, 0x20}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47, 0x20}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b, 0x20}, {0x53, 0x47, 0x44, 0x20}, {0x53, 0x48, 0x50, 0x20}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c, 0x20}, {0x53, 0x4f, 0x53, 0x20}, {0x53, 0x52, 0x44, 0x20}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50, 0x20}, {0x53, 0x54, 0x44, 0x20}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50, 0x20}, {0x53, 0x5a, 0x4c, 0x20}, {0x54, 0x48, 0x42, 0x20}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53, 0x20}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54, 0x20}, {0x54, 0x4e, 0x44, 0x20}, {0x54, 0x4f, 0x50, 0x20}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59, 0x20}, {0x54, 0x54, 0x44, 0x20}, {0x54, 0x57, 0x44, 0x20}, {0x54, 0x5a, 0x53, 0x20}, {0x55, 0x41, 0x48, 0x20}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58, 0x20}, {0x55, 0x53, 0x44, 0x20}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55, 0x20}, {0x55, 0x5a, 0x53, 0x20}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46, 0x20}, {0x56, 0x4e, 0x44, 0x20}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56, 0x20}, {0x57, 0x53, 0x54, 0x20}, {0x58, 0x41, 0x46, 0x20}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x58, 0x43, 0x44, 0x20}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x58, 0x4f, 0x46, 0x20}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46, 0x20}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52, 0x20}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52, 0x20}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57, 0x20}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		monthsAbbreviated:  [][]uint8{[]uint8(nil), {0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x8f, 0xe2, 0xb5, 0x8f}, {0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x95, 0xe2, 0xb4, 0xb0}, {0xe2, 0xb5, 0x8e, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x95}, {0xe2, 0xb5, 0x89, 0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x94}, {0xe2, 0xb5, 0x8e, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa2}, {0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x8f}, {0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x8d}, {0xe2, 0xb5, 0x96, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x9b}, {0xe2, 0xb5, 0x9b, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x9c}, {0xe2, 0xb4, 0xbd, 0xe2, 0xb5, 0x9c, 0xe2, 0xb5, 0x93}, {0xe2, 0xb5, 0x8f, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0xa1}, {0xe2, 0xb4, 0xb7, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x8a}},
		monthsNarrow:       [][]uint8{[]uint8(nil), {0xe2, 0xb5, 0x89}, {0xe2, 0xb4, 0xb1}, {0xe2, 0xb5, 0x8e}, {0xe2, 0xb5, 0x89}, {0xe2, 0xb5, 0x8e}, {0xe2, 0xb5, 0xa2}, {0xe2, 0xb5, 0xa2}, {0xe2, 0xb5, 0x96}, {0xe2, 0xb5, 0x9b}, {0xe2, 0xb4, 0xbd}, {0xe2, 0xb5, 0x8f}, {0xe2, 0xb4, 0xb7}},
		monthsWide:         [][]uint8{[]uint8(nil), {0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x8f, 0xe2, 0xb5, 0x8f, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x94}, {0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x95, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x95}, {0xe2, 0xb5, 0x8e, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x95, 0xe2, 0xb5, 0x9a}, {0xe2, 0xb5, 0x89, 0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x94, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x94}, {0xe2, 0xb5, 0x8e, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x93}, {0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x8f, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x93}, {0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x8d, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0xa3}, {0xe2, 0xb5, 0x96, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x9b, 0xe2, 0xb5, 0x9c}, {0xe2, 0xb5, 0x9b, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x9c, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x8f, 0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x94}, {0xe2, 0xb4, 0xbd, 0xe2, 0xb5, 0x9c, 0xe2, 0xb5, 0x93, 0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x94}, {0xe2, 0xb5, 0x8f, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0xa1, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x8f, 0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x94}, {0xe2, 0xb4, 0xb7, 0xe2, 0xb5, 0x93, 0xe2, 0xb5, 0x8a, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x8f, 0xe2, 0xb4, 0xb1, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x94}},
		daysAbbreviated:    [][]uint8{{0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99, 0xe2, 0xb4, 0xb0}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x8f}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99, 0xe2, 0xb5, 0x89}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb4, 0xbd, 0xe2, 0xb5, 0x95}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb4, 0xbd, 0xe2, 0xb5, 0xa1}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x8e}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99, 0xe2, 0xb5, 0x89, 0xe2, 0xb4, 0xb9}},
		daysWide:           [][]uint8{{0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x8e, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa2, 0xe2, 0xb5, 0x8f, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x8f, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb4, 0xbd, 0xe2, 0xb5, 0x95, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb4, 0xbd, 0xe2, 0xb5, 0xa1, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99}, {0xe2, 0xb5, 0x99, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x8e, 0xe2, 0xb5, 0xa1, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99}, {0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99, 0xe2, 0xb5, 0x89, 0xe2, 0xb4, 0xb9, 0xe2, 0xb5, 0xa2, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x99}},
		periodsAbbreviated: [][]uint8{{0xe2, 0xb5, 0x9c, 0xe2, 0xb5, 0x89, 0xe2, 0xb4, 0xbc, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa1, 0xe2, 0xb5, 0x9c}, {0xe2, 0xb5, 0x9c, 0xe2, 0xb4, 0xb0, 0xe2, 0xb4, 0xb7, 0xe2, 0xb4, 0xb3, 0xe2, 0xb4, 0xb3, 0xe2, 0xb5, 0xaf, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x9c}},
		periodsWide:        [][]uint8{{0xe2, 0xb5, 0x9c, 0xe2, 0xb5, 0x89, 0xe2, 0xb4, 0xbc, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0xa1, 0xe2, 0xb5, 0x9c}, {0xe2, 0xb5, 0x9c, 0xe2, 0xb4, 0xb0, 0xe2, 0xb4, 0xb7, 0xe2, 0xb4, 0xb3, 0xe2, 0xb4, 0xb3, 0xe2, 0xb5, 0xaf, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x9c}},
		erasAbbreviated:    [][]uint8{{0xe2, 0xb4, 0xb7, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x84}, {0xe2, 0xb4, 0xb7, 0xe2, 0xb4, 0xbc, 0xe2, 0xb5, 0x84}},
		erasNarrow:         [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:           [][]uint8{{0xe2, 0xb4, 0xb7, 0xe2, 0xb4, 0xb0, 0xe2, 0xb5, 0x9c, 0x20, 0xe2, 0xb5, 0x8f, 0x20, 0xe2, 0xb5, 0x84, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x99, 0xe2, 0xb4, 0xb0}, {0xe2, 0xb4, 0xb7, 0xe2, 0xb4, 0xbc, 0xe2, 0xb4, 0xbc, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x94, 0x20, 0xe2, 0xb5, 0x8f, 0x20, 0xe2, 0xb5, 0x84, 0xe2, 0xb5, 0x89, 0xe2, 0xb5, 0x99, 0xe2, 0xb4, 0xb0}},
		timezones:          map[string][]uint8{"∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "GFT": {0x47, 0x46, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "HAT": {0x48, 0x41, 0x54}, "COT": {0x43, 0x4f, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "EAT": {0x45, 0x41, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "GYT": {0x47, 0x59, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "WIB": {0x57, 0x49, 0x42}, "UYT": {0x55, 0x59, 0x54}, "SRT": {0x53, 0x52, 0x54}, "MST": {0x4d, 0x53, 0x54}, "ADT": {0x41, 0x44, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "CDT": {0x43, 0x44, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "EST": {0x45, 0x53, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "AST": {0x41, 0x53, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "EDT": {0x45, 0x44, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "PST": {0x50, 0x53, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "CAT": {0x43, 0x41, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "WAT": {0x57, 0x41, 0x54}, "PDT": {0x50, 0x44, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "SGT": {0x53, 0x47, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "WIT": {0x57, 0x49, 0x54}, "VET": {0x56, 0x45, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "ECT": {0x45, 0x43, 0x54}, "CST": {0x43, 0x53, 0x54}, "ART": {0x41, 0x52, 0x54}, "JST": {0x4a, 0x53, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "AWST": {0x41, 0x57, 0x53, 0x54}, "BT": {0x42, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "IST": {0x49, 0x53, 0x54}},
	}
}

// Locale returns the current translators string locale
func (shi *shi_Tfng_MA) Locale() string {
	return shi.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'shi_Tfng_MA'
func (shi *shi_Tfng_MA) PluralsCardinal() []locales.PluralRule {
	return shi.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'shi_Tfng_MA'
func (shi *shi_Tfng_MA) PluralsOrdinal() []locales.PluralRule {
	return shi.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'shi_Tfng_MA'
func (shi *shi_Tfng_MA) PluralsRange() []locales.PluralRule {
	return shi.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'shi_Tfng_MA'
func (shi *shi_Tfng_MA) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	} else if n >= 2 && n <= 10 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'shi_Tfng_MA'
func (shi *shi_Tfng_MA) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'shi_Tfng_MA'
func (shi *shi_Tfng_MA) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (shi *shi_Tfng_MA) MonthAbbreviated(month time.Month) []byte {
	return shi.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (shi *shi_Tfng_MA) MonthsAbbreviated() [][]byte {
	return shi.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (shi *shi_Tfng_MA) MonthNarrow(month time.Month) []byte {
	return shi.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (shi *shi_Tfng_MA) MonthsNarrow() [][]byte {
	return shi.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (shi *shi_Tfng_MA) MonthWide(month time.Month) []byte {
	return shi.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (shi *shi_Tfng_MA) MonthsWide() [][]byte {
	return shi.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (shi *shi_Tfng_MA) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return shi.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (shi *shi_Tfng_MA) WeekdaysAbbreviated() [][]byte {
	return shi.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (shi *shi_Tfng_MA) WeekdayNarrow(weekday time.Weekday) []byte {
	return shi.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (shi *shi_Tfng_MA) WeekdaysNarrow() [][]byte {
	return shi.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (shi *shi_Tfng_MA) WeekdayShort(weekday time.Weekday) []byte {
	return shi.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (shi *shi_Tfng_MA) WeekdaysShort() [][]byte {
	return shi.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (shi *shi_Tfng_MA) WeekdayWide(weekday time.Weekday) []byte {
	return shi.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (shi *shi_Tfng_MA) WeekdaysWide() [][]byte {
	return shi.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'shi_Tfng_MA' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'shi_Tfng_MA' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (shi *shi_Tfng_MA) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := shi.currencies[currency]
	l := len(s) + len(shi.decimal) + len(shi.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, shi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(shi.group) - 1; j >= 0; j-- {
					b = append(b, shi.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(shi.minus) - 1; j >= 0; j-- {
			b = append(b, shi.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, shi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'shi_Tfng_MA'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := shi.currencies[currency]
	l := len(s) + len(shi.decimal) + len(shi.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, shi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(shi.group) - 1; j >= 0; j-- {
					b = append(b, shi.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(shi.minus) - 1; j >= 0; j-- {
			b = append(b, shi.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, shi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, symbol...)
	} else {

		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, shi.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, shi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, shi.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, shi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'shi_Tfng_MA'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (shi *shi_Tfng_MA) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}
