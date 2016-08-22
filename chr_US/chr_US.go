package chr_US

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type chr_US struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
	perMille               []byte
	timeSeparator          []byte
	inifinity              []byte
	currencies             [][]byte // idx = enum of currency code
	currencyNegativePrefix []byte
	currencyNegativeSuffix []byte
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

// New returns a new instance of translator for the 'chr_US' locale
func New() locales.Translator {
	return &chr_US{
		locale:                 "chr_US",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44, 0x20}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e, 0x20}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c, 0x20}, {0x41, 0x4d, 0x44, 0x20}, {0x41, 0x4e, 0x47, 0x20}, {0x41, 0x4f, 0x41, 0x20}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53, 0x20}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44, 0x20}, {0x41, 0x57, 0x47, 0x20}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e, 0x20}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d, 0x20}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44, 0x20}, {0x42, 0x44, 0x54, 0x20}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e, 0x20}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44, 0x20}, {0x42, 0x49, 0x46, 0x20}, {0x42, 0x4d, 0x44, 0x20}, {0x42, 0x4e, 0x44, 0x20}, {0x42, 0x4f, 0x42, 0x20}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x42, 0x52, 0x4c, 0x20}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44, 0x20}, {0x42, 0x54, 0x4e, 0x20}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50, 0x20}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52, 0x20}, {0x42, 0x5a, 0x44, 0x20}, {0x43, 0x41, 0x44, 0x20}, {0x43, 0x44, 0x46, 0x20}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46, 0x20}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50, 0x20}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0x59, 0x20}, {0x43, 0x4f, 0x50, 0x20}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43, 0x20}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43, 0x20}, {0x43, 0x55, 0x50, 0x20}, {0x43, 0x56, 0x45, 0x20}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b, 0x20}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46, 0x20}, {0x44, 0x4b, 0x4b, 0x20}, {0x44, 0x4f, 0x50, 0x20}, {0x44, 0x5a, 0x44, 0x20}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50, 0x20}, {0x45, 0x52, 0x4e, 0x20}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42, 0x20}, {0x45, 0x55, 0x52, 0x20}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44, 0x20}, {0x46, 0x4b, 0x50, 0x20}, {0x46, 0x52, 0x46, 0x20}, {0x47, 0x42, 0x50, 0x20}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c, 0x20}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53, 0x20}, {0x47, 0x49, 0x50, 0x20}, {0x47, 0x4d, 0x44, 0x20}, {0x47, 0x4e, 0x46, 0x20}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51, 0x20}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44, 0x20}, {0x48, 0x4b, 0x44, 0x20}, {0x48, 0x4e, 0x4c, 0x20}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b, 0x20}, {0x48, 0x54, 0x47, 0x20}, {0x48, 0x55, 0x46, 0x20}, {0x49, 0x44, 0x52, 0x20}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0x49, 0x4c, 0x53, 0x20}, {0x49, 0x4e, 0x52, 0x20}, {0x49, 0x51, 0x44, 0x20}, {0x49, 0x52, 0x52, 0x20}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b, 0x20}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44, 0x20}, {0x4a, 0x4f, 0x44, 0x20}, {0x4a, 0x50, 0x59, 0x20}, {0x4b, 0x45, 0x53, 0x20}, {0x4b, 0x47, 0x53, 0x20}, {0x4b, 0x48, 0x52, 0x20}, {0x4b, 0x4d, 0x46, 0x20}, {0x4b, 0x50, 0x57, 0x20}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0x4b, 0x52, 0x57, 0x20}, {0x4b, 0x57, 0x44, 0x20}, {0x4b, 0x59, 0x44, 0x20}, {0x4b, 0x5a, 0x54, 0x20}, {0x4c, 0x41, 0x4b, 0x20}, {0x4c, 0x42, 0x50, 0x20}, {0x4c, 0x4b, 0x52, 0x20}, {0x4c, 0x52, 0x44, 0x20}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c, 0x20}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c, 0x20}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44, 0x20}, {0x4d, 0x41, 0x44, 0x20}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c, 0x20}, {0x4d, 0x47, 0x41, 0x20}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44, 0x20}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b, 0x20}, {0x4d, 0x4e, 0x54, 0x20}, {0x4d, 0x4f, 0x50, 0x20}, {0x4d, 0x52, 0x4f, 0x20}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52, 0x20}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52, 0x20}, {0x4d, 0x57, 0x4b, 0x20}, {0x4d, 0x58, 0x4e, 0x20}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52, 0x20}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e, 0x20}, {0x4e, 0x41, 0x44, 0x20}, {0x4e, 0x47, 0x4e, 0x20}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f, 0x20}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b, 0x20}, {0x4e, 0x50, 0x52, 0x20}, {0x4e, 0x5a, 0x44, 0x20}, {0x4f, 0x4d, 0x52, 0x20}, {0x50, 0x41, 0x42, 0x20}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e, 0x20}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b, 0x20}, {0x50, 0x48, 0x50, 0x20}, {0x50, 0x4b, 0x52, 0x20}, {0x50, 0x4c, 0x4e, 0x20}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47, 0x20}, {0x51, 0x41, 0x52, 0x20}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e, 0x20}, {0x52, 0x53, 0x44, 0x20}, {0x52, 0x55, 0x42, 0x20}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46, 0x20}, {0x53, 0x41, 0x52, 0x20}, {0x53, 0x42, 0x44, 0x20}, {0x53, 0x43, 0x52, 0x20}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47, 0x20}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b, 0x20}, {0x53, 0x47, 0x44, 0x20}, {0x53, 0x48, 0x50, 0x20}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c, 0x20}, {0x53, 0x4f, 0x53, 0x20}, {0x53, 0x52, 0x44, 0x20}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50, 0x20}, {0x53, 0x54, 0x44, 0x20}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50, 0x20}, {0x53, 0x5a, 0x4c, 0x20}, {0x54, 0x48, 0x42, 0x20}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53, 0x20}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54, 0x20}, {0x54, 0x4e, 0x44, 0x20}, {0x54, 0x4f, 0x50, 0x20}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59, 0x20}, {0x54, 0x54, 0x44, 0x20}, {0x54, 0x57, 0x44, 0x20}, {0x54, 0x5a, 0x53, 0x20}, {0x55, 0x41, 0x48, 0x20}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58, 0x20}, {0x55, 0x53, 0x44, 0x20}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55, 0x20}, {0x55, 0x5a, 0x53, 0x20}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46, 0x20}, {0x56, 0x4e, 0x44, 0x20}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56, 0x20}, {0x57, 0x53, 0x54, 0x20}, {0x58, 0x41, 0x46, 0x20}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x58, 0x43, 0x44, 0x20}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x58, 0x4f, 0x46, 0x20}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46, 0x20}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52, 0x20}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52, 0x20}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57, 0x20}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0xe1, 0x8e, 0xa4, 0xe1, 0x8f, 0x83}, {0xe1, 0x8e, 0xa7, 0xe1, 0x8e, 0xa6}, {0xe1, 0x8e, 0xa0, 0xe1, 0x8f, 0x85}, {0xe1, 0x8e, 0xa7, 0xe1, 0x8f, 0xac}, {0xe1, 0x8e, 0xa0, 0xe1, 0x8f, 0x82}, {0xe1, 0x8f, 0x95, 0xe1, 0x8e, 0xad}, {0xe1, 0x8e, 0xab, 0xe1, 0x8f, 0xb0}, {0xe1, 0x8e, 0xa6, 0xe1, 0x8e, 0xb6}, {0xe1, 0x8f, 0x9a, 0xe1, 0x8e, 0xb5}, {0xe1, 0x8f, 0x9a, 0xe1, 0x8f, 0x82}, {0xe1, 0x8f, 0x85, 0xe1, 0x8f, 0x93}, {0xe1, 0x8e, 0xa5, 0xe1, 0x8f, 0x8d}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0xe1, 0x8e, 0xa4}, {0xe1, 0x8e, 0xa7}, {0xe1, 0x8e, 0xa0}, {0xe1, 0x8e, 0xa7}, {0xe1, 0x8e, 0xa0}, {0xe1, 0x8f, 0x95}, {0xe1, 0x8e, 0xab}, {0xe1, 0x8e, 0xa6}, {0xe1, 0x8f, 0x9a}, {0xe1, 0x8f, 0x9a}, {0xe1, 0x8f, 0x85}, {0xe1, 0x8e, 0xa5}},
		monthsWide:             [][]uint8{[]uint8(nil), {0xe1, 0x8e, 0xa4, 0xe1, 0x8f, 0x83, 0xe1, 0x8e, 0xb8, 0xe1, 0x8f, 0x94, 0xe1, 0x8f, 0x85}, {0xe1, 0x8e, 0xa7, 0xe1, 0x8e, 0xa6, 0xe1, 0x8e, 0xb5}, {0xe1, 0x8e, 0xa0, 0xe1, 0x8f, 0x85, 0xe1, 0x8f, 0xb1}, {0xe1, 0x8e, 0xa7, 0xe1, 0x8f, 0xac, 0xe1, 0x8f, 0x82}, {0xe1, 0x8e, 0xa0, 0xe1, 0x8f, 0x82, 0xe1, 0x8f, 0x8d, 0xe1, 0x8e, 0xac, 0xe1, 0x8f, 0x98}, {0xe1, 0x8f, 0x95, 0xe1, 0x8e, 0xad, 0xe1, 0x8e, 0xb7, 0xe1, 0x8f, 0xb1}, {0xe1, 0x8e, 0xab, 0xe1, 0x8f, 0xb0, 0xe1, 0x8f, 0x89, 0xe1, 0x8f, 0x82}, {0xe1, 0x8e, 0xa6, 0xe1, 0x8e, 0xb6, 0xe1, 0x8f, 0x82}, {0xe1, 0x8f, 0x9a, 0xe1, 0x8e, 0xb5, 0xe1, 0x8f, 0x8d, 0xe1, 0x8f, 0x97}, {0xe1, 0x8f, 0x9a, 0xe1, 0x8f, 0x82, 0xe1, 0x8f, 0x85, 0xe1, 0x8f, 0x97}, {0xe1, 0x8f, 0x85, 0xe1, 0x8f, 0x93, 0xe1, 0x8f, 0x95, 0xe1, 0x8f, 0x86}, {0xe1, 0x8e, 0xa5, 0xe1, 0x8f, 0x8d, 0xe1, 0x8e, 0xa9, 0xe1, 0x8f, 0xb1}},
		daysAbbreviated:        [][]uint8{{0xe1, 0x8f, 0x86, 0xe1, 0x8f, 0x8d, 0xe1, 0x8e, 0xac}, {0xe1, 0x8f, 0x89, 0xe1, 0x8f, 0x85, 0xe1, 0x8e, 0xaf}, {0xe1, 0x8f, 0x94, 0xe1, 0x8e, 0xb5, 0xe1, 0x8f, 0x81}, {0xe1, 0x8f, 0xa6, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0x81}, {0xe1, 0x8f, 0x85, 0xe1, 0x8e, 0xa9, 0xe1, 0x8f, 0x81}, {0xe1, 0x8f, 0xa7, 0xe1, 0x8e, 0xbe, 0xe1, 0x8e, 0xa9}, {0xe1, 0x8f, 0x88, 0xe1, 0x8f, 0x95, 0xe1, 0x8e, 0xbe}},
		daysNarrow:             [][]uint8{{0xe1, 0x8f, 0x86}, {0xe1, 0x8f, 0x89}, {0xe1, 0x8f, 0x94}, {0xe1, 0x8f, 0xa6}, {0xe1, 0x8f, 0x85}, {0xe1, 0x8f, 0xa7}, {0xe1, 0x8e, 0xa4}},
		daysWide:               [][]uint8{{0xe1, 0x8e, 0xa4, 0xe1, 0x8e, 0xbe, 0xe1, 0x8f, 0x99, 0xe1, 0x8f, 0x93, 0xe1, 0x8f, 0x86, 0xe1, 0x8f, 0x8d, 0xe1, 0x8e, 0xac}, {0xe1, 0x8e, 0xa4, 0xe1, 0x8e, 0xbe, 0xe1, 0x8f, 0x99, 0xe1, 0x8f, 0x93, 0xe1, 0x8f, 0x89, 0xe1, 0x8f, 0x85, 0xe1, 0x8e, 0xaf}, {0xe1, 0x8f, 0x94, 0xe1, 0x8e, 0xb5, 0xe1, 0x8f, 0x81, 0xe1, 0x8e, 0xa2, 0xe1, 0x8e, 0xa6}, {0xe1, 0x8f, 0xa6, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0x81, 0xe1, 0x8e, 0xa2, 0xe1, 0x8e, 0xa6}, {0xe1, 0x8f, 0x85, 0xe1, 0x8e, 0xa9, 0xe1, 0x8f, 0x81, 0xe1, 0x8e, 0xa2, 0xe1, 0x8e, 0xa6}, {0xe1, 0x8f, 0xa7, 0xe1, 0x8e, 0xbe, 0xe1, 0x8e, 0xa9, 0xe1, 0x8e, 0xb6, 0xe1, 0x8f, 0x8d, 0xe1, 0x8f, 0x97}, {0xe1, 0x8e, 0xa4, 0xe1, 0x8e, 0xbe, 0xe1, 0x8f, 0x99, 0xe1, 0x8f, 0x93, 0xe1, 0x8f, 0x88, 0xe1, 0x8f, 0x95, 0xe1, 0x8e, 0xbe}},
		periodsAbbreviated:     [][]uint8{{0xe1, 0x8f, 0x8c, 0xe1, 0x8e, 0xbe, 0xe1, 0x8e, 0xb4}, {0xe1, 0x8f, 0x92, 0xe1, 0x8e, 0xaf, 0xe1, 0x8f, 0xb1, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0x97, 0xe1, 0x8f, 0xa2}},
		periodsWide:            [][]uint8{{0xe1, 0x8f, 0x8c, 0xe1, 0x8e, 0xbe, 0xe1, 0x8e, 0xb4}, {0xe1, 0x8f, 0x92, 0xe1, 0x8e, 0xaf, 0xe1, 0x8f, 0xb1, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0x97, 0xe1, 0x8f, 0xa2}},
		erasAbbreviated:        [][]uint8{{0xe1, 0x8e, 0xa4, 0xe1, 0x8f, 0x93, 0xe1, 0x8e, 0xb7, 0xe1, 0x8e, 0xb8}, {0xe1, 0x8e, 0xa4, 0xe1, 0x8e, 0xb6, 0xe1, 0x8f, 0x90, 0xe1, 0x8f, 0x85}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0xe1, 0x8f, 0x8f, 0x20, 0xe1, 0x8f, 0xa5, 0xe1, 0x8f, 0x8c, 0x20, 0xe1, 0x8e, 0xbe, 0xe1, 0x8f, 0x95, 0xe1, 0x8e, 0xb2, 0xe1, 0x8f, 0x8d, 0xe1, 0x8e, 0xac, 0xe1, 0x8e, 0xbe}, {0xe1, 0x8e, 0xa0, 0xe1, 0x8e, 0xa9, 0xe1, 0x8f, 0x83, 0xe1, 0x8e, 0xae, 0xe1, 0x8e, 0xb5, 0xe1, 0x8f, 0x93, 0xe1, 0x8f, 0x8d, 0xe1, 0x8f, 0x97, 0xe1, 0x8f, 0xb1, 0x20, 0xe1, 0x8e, 0xa0, 0xe1, 0x8f, 0x95, 0xe1, 0x8f, 0x98, 0xe1, 0x8f, 0xb1, 0xe1, 0x8f, 0x8d, 0xe1, 0x8e, 0xac, 0x20, 0xe1, 0x8f, 0xb1, 0xe1, 0x8e, 0xb0, 0xe1, 0x8f, 0xa9, 0x20, 0xe1, 0x8f, 0xa7, 0xe1, 0x8f, 0x93, 0xe1, 0x8f, 0x82, 0xe1, 0x8e, 0xb8, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0x8d, 0xe1, 0x8f, 0x97}},
		timezones:              map[string][]uint8{"HNT": {0x48, 0x4e, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "WIB": {0x57, 0x49, 0x42}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "WIT": {0x57, 0x49, 0x54}, "MDT": {0xe1, 0x8e, 0xa3, 0xe1, 0x8f, 0x93, 0xe1, 0x8e, 0xb8, 0x20, 0xe1, 0x8e, 0xa2, 0xe1, 0x8e, 0xa6, 0x20, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0xb3, 0xe1, 0x8f, 0xa9, 0xe1, 0x8e, 0xaa, 0xe1, 0x8f, 0x97}, "EST": {0xe1, 0x8e, 0xa7, 0xe1, 0x8e, 0xb8, 0xe1, 0x8e, 0xac, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0x97, 0xe1, 0x8f, 0xa2, 0x20, 0xe1, 0x8f, 0xb0, 0xe1, 0x8e, 0xb5, 0xe1, 0x8f, 0x8a, 0x20, 0xe1, 0x8f, 0x97, 0xe1, 0x8f, 0x99, 0xe1, 0x8e, 0xb3, 0xe1, 0x8e, 0xa9, 0x20, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0xb3, 0xe1, 0x8f, 0xa9, 0xe1, 0x8e, 0xaa, 0xe1, 0x8f, 0x97}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "VET": {0x56, 0x45, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "MST": {0xe1, 0x8e, 0xa3, 0xe1, 0x8f, 0x93, 0xe1, 0x8e, 0xb8, 0x20, 0xe1, 0x8f, 0xb0, 0xe1, 0x8e, 0xb5, 0xe1, 0x8f, 0x8a, 0x20, 0xe1, 0x8f, 0x97, 0xe1, 0x8f, 0x99, 0xe1, 0x8e, 0xb3, 0xe1, 0x8e, 0xa9, 0x20, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0xb3, 0xe1, 0x8f, 0xa9, 0xe1, 0x8e, 0xaa, 0xe1, 0x8f, 0x97}, "TMT": {0x54, 0x4d, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "ECT": {0x45, 0x43, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "COT": {0x43, 0x4f, 0x54}, "CAT": {0x43, 0x41, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "IST": {0x49, 0x53, 0x54}, "AST": {0x41, 0x53, 0x54}, "GFT": {0x47, 0x46, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "ACST": {0x41, 0x43, 0x53, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "EDT": {0xe1, 0x8e, 0xa7, 0xe1, 0x8e, 0xb8, 0xe1, 0x8e, 0xac, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0x97, 0xe1, 0x8f, 0xa2, 0x20, 0xe1, 0x8e, 0xa2, 0xe1, 0x8e, 0xa6, 0x20, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0xb3, 0xe1, 0x8f, 0xa9, 0xe1, 0x8e, 0xaa, 0xe1, 0x8f, 0x97}, "MYT": {0x4d, 0x59, 0x54}, "SRT": {0x53, 0x52, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "ART": {0x41, 0x52, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "GYT": {0x47, 0x59, 0x54}, "BT": {0x42, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "WAT": {0x57, 0x41, 0x54}, "CDT": {0xe1, 0x8e, 0xa0, 0xe1, 0x8f, 0xb0, 0xe1, 0x8e, 0xb5, 0x20, 0xe1, 0x8e, 0xa2, 0xe1, 0x8e, 0xa6, 0x20, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0xb3, 0xe1, 0x8f, 0xa9, 0xe1, 0x8e, 0xaa, 0xe1, 0x8f, 0x97}, "EAT": {0x45, 0x41, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "CLT": {0x43, 0x4c, 0x54}, "UYT": {0x55, 0x59, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "HAT": {0x48, 0x41, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "CST": {0xe1, 0x8e, 0xa0, 0xe1, 0x8f, 0xb0, 0xe1, 0x8e, 0xb5, 0x20, 0xe1, 0x8f, 0xb0, 0xe1, 0x8e, 0xb5, 0xe1, 0x8f, 0x8a, 0x20, 0xe1, 0x8f, 0x97, 0xe1, 0x8f, 0x99, 0xe1, 0x8e, 0xb3, 0xe1, 0x8e, 0xa9, 0x20, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0xb3, 0xe1, 0x8f, 0xa9, 0xe1, 0x8e, 0xaa, 0xe1, 0x8f, 0x97}, "GMT": {0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0xa4, 0x20, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0xb3, 0xe1, 0x8f, 0x8d, 0xe1, 0x8f, 0x97, 0x20, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0xb3, 0xe1, 0x8f, 0xa9, 0xe1, 0x8e, 0xaa, 0xe1, 0x8f, 0x97}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "PST": {0xe1, 0x8f, 0xad, 0xe1, 0x8f, 0x95, 0xe1, 0x8e, 0xb5, 0xe1, 0x8e, 0xac, 0x20, 0xe1, 0x8f, 0xb0, 0xe1, 0x8e, 0xb5, 0xe1, 0x8f, 0x8a, 0x20, 0xe1, 0x8f, 0x97, 0xe1, 0x8f, 0x99, 0xe1, 0x8e, 0xb3, 0xe1, 0x8e, 0xa9, 0x20, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0xb3, 0xe1, 0x8f, 0xa9, 0xe1, 0x8e, 0xaa, 0xe1, 0x8f, 0x97}, "PDT": {0xe1, 0x8f, 0xad, 0xe1, 0x8f, 0x95, 0xe1, 0x8e, 0xb5, 0xe1, 0x8e, 0xac, 0x20, 0xe1, 0x8e, 0xa2, 0xe1, 0x8e, 0xa6, 0x20, 0xe1, 0x8e, 0xa2, 0xe1, 0x8f, 0xb3, 0xe1, 0x8f, 0xa9, 0xe1, 0x8e, 0xaa, 0xe1, 0x8f, 0x97}, "ADT": {0x41, 0x44, 0x54}, "SGT": {0x53, 0x47, 0x54}, "BOT": {0x42, 0x4f, 0x54}},
	}
}

// Locale returns the current translators string locale
func (chr *chr_US) Locale() string {
	return chr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'chr_US'
func (chr *chr_US) PluralsCardinal() []locales.PluralRule {
	return chr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'chr_US'
func (chr *chr_US) PluralsOrdinal() []locales.PluralRule {
	return chr.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'chr_US'
func (chr *chr_US) PluralsRange() []locales.PluralRule {
	return chr.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'chr_US'
func (chr *chr_US) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'chr_US'
func (chr *chr_US) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'chr_US'
func (chr *chr_US) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (chr *chr_US) MonthAbbreviated(month time.Month) []byte {
	return chr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (chr *chr_US) MonthsAbbreviated() [][]byte {
	return chr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (chr *chr_US) MonthNarrow(month time.Month) []byte {
	return chr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (chr *chr_US) MonthsNarrow() [][]byte {
	return chr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (chr *chr_US) MonthWide(month time.Month) []byte {
	return chr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (chr *chr_US) MonthsWide() [][]byte {
	return chr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (chr *chr_US) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return chr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (chr *chr_US) WeekdaysAbbreviated() [][]byte {
	return chr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (chr *chr_US) WeekdayNarrow(weekday time.Weekday) []byte {
	return chr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (chr *chr_US) WeekdaysNarrow() [][]byte {
	return chr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (chr *chr_US) WeekdayShort(weekday time.Weekday) []byte {
	return chr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (chr *chr_US) WeekdaysShort() [][]byte {
	return chr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (chr *chr_US) WeekdayWide(weekday time.Weekday) []byte {
	return chr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (chr *chr_US) WeekdaysWide() [][]byte {
	return chr.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'chr_US' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (chr *chr_US) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(chr.decimal) + len(chr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, chr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, chr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, chr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'chr_US' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (chr *chr_US) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(chr.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, chr.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, chr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, chr.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'chr_US'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (chr *chr_US) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := chr.currencies[currency]
	l := len(s) + len(chr.decimal) + len(chr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, chr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, chr.group[0])
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

	if num < 0 {
		b = append(b, chr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, chr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'chr_US'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (chr *chr_US) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := chr.currencies[currency]
	l := len(s) + len(chr.decimal) + len(chr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, chr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, chr.group[0])
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

		for j := len(chr.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, chr.currencyNegativePrefix[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, chr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, chr.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'chr_US'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (chr *chr_US) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'chr_US'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (chr *chr_US) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, chr.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'chr_US'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (chr *chr_US) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, chr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'chr_US'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (chr *chr_US) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, chr.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, chr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'chr_US'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (chr *chr_US) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, chr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, chr.periodsAbbreviated[0]...)
	} else {
		b = append(b, chr.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'chr_US'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (chr *chr_US) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, chr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, chr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, chr.periodsAbbreviated[0]...)
	} else {
		b = append(b, chr.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'chr_US'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (chr *chr_US) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, chr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, chr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, chr.periodsAbbreviated[0]...)
	} else {
		b = append(b, chr.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'chr_US'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (chr *chr_US) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, chr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, chr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, chr.periodsAbbreviated[0]...)
	} else {
		b = append(b, chr.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := chr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
