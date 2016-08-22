package yi

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type yi struct {
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
	currencyPositivePrefix []byte
	currencyPositiveSuffix []byte
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

// New returns a new instance of translator for the 'yi' locale
func New() locales.Translator {
	return &yi{
		locale:                 "yi",
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
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44, 0x20}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e, 0x20}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c, 0x20}, {0x41, 0x4d, 0x44, 0x20}, {0x41, 0x4e, 0x47, 0x20}, {0x41, 0x4f, 0x41, 0x20}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53, 0x20}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44, 0x20}, {0x41, 0x57, 0x47, 0x20}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e, 0x20}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d, 0x20}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44, 0x20}, {0x42, 0x44, 0x54, 0x20}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e, 0x20}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44, 0x20}, {0x42, 0x49, 0x46, 0x20}, {0x42, 0x4d, 0x44, 0x20}, {0x42, 0x4e, 0x44, 0x20}, {0x42, 0x4f, 0x42, 0x20}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x42, 0x52, 0x4c, 0x20}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44, 0x20}, {0x42, 0x54, 0x4e, 0x20}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50, 0x20}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52, 0x20}, {0x42, 0x5a, 0x44, 0x20}, {0x43, 0x41, 0x44, 0x20}, {0x43, 0x44, 0x46, 0x20}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46, 0x20}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50, 0x20}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0x59, 0x20}, {0x43, 0x4f, 0x50, 0x20}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43, 0x20}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43, 0x20}, {0x43, 0x55, 0x50, 0x20}, {0x43, 0x56, 0x45, 0x20}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b, 0x20}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46, 0x20}, {0x44, 0x4b, 0x4b, 0x20}, {0x44, 0x4f, 0x50, 0x20}, {0x44, 0x5a, 0x44, 0x20}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50, 0x20}, {0x45, 0x52, 0x4e, 0x20}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42, 0x20}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44, 0x20}, {0x46, 0x4b, 0x50, 0x20}, {0x46, 0x52, 0x46, 0x20}, {0xc2, 0xa3}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c, 0x20}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53, 0x20}, {0x47, 0x49, 0x50, 0x20}, {0x47, 0x4d, 0x44, 0x20}, {0x47, 0x4e, 0x46, 0x20}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51, 0x20}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44, 0x20}, {0x48, 0x4b, 0x44, 0x20}, {0x48, 0x4e, 0x4c, 0x20}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b, 0x20}, {0x48, 0x54, 0x47, 0x20}, {0x48, 0x55, 0x46, 0x20}, {0x49, 0x44, 0x52, 0x20}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0x49, 0x4c, 0x53, 0x20}, {0xe2, 0x82, 0xb9}, {0x49, 0x51, 0x44, 0x20}, {0x49, 0x52, 0x52, 0x20}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b, 0x20}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44, 0x20}, {0x4a, 0x4f, 0x44, 0x20}, {0x4a, 0x50, 0xc2, 0xa5}, {0x4b, 0x45, 0x53, 0x20}, {0x4b, 0x47, 0x53, 0x20}, {0x4b, 0x48, 0x52, 0x20}, {0x4b, 0x4d, 0x46, 0x20}, {0x4b, 0x50, 0x57, 0x20}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0x4b, 0x52, 0x57, 0x20}, {0x4b, 0x57, 0x44, 0x20}, {0x4b, 0x59, 0x44, 0x20}, {0x4b, 0x5a, 0x54, 0x20}, {0x4c, 0x41, 0x4b, 0x20}, {0x4c, 0x42, 0x50, 0x20}, {0x4c, 0x4b, 0x52, 0x20}, {0x4c, 0x52, 0x44, 0x20}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c, 0x20}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c, 0x20}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44, 0x20}, {0x4d, 0x41, 0x44, 0x20}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c, 0x20}, {0x4d, 0x47, 0x41, 0x20}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44, 0x20}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b, 0x20}, {0x4d, 0x4e, 0x54, 0x20}, {0x4d, 0x4f, 0x50, 0x20}, {0x4d, 0x52, 0x4f, 0x20}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52, 0x20}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52, 0x20}, {0x4d, 0x57, 0x4b, 0x20}, {0x4d, 0x58, 0x4e, 0x20}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52, 0x20}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e, 0x20}, {0x4e, 0x41, 0x44, 0x20}, {0x4e, 0x47, 0x4e, 0x20}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f, 0x20}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b, 0x20}, {0x4e, 0x50, 0x52, 0x20}, {0x4e, 0x5a, 0x44, 0x20}, {0x4f, 0x4d, 0x52, 0x20}, {0x50, 0x41, 0x42, 0x20}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e, 0x20}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b, 0x20}, {0x50, 0x48, 0x50, 0x20}, {0x50, 0x4b, 0x52, 0x20}, {0x50, 0x4c, 0x4e, 0x20}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47, 0x20}, {0x51, 0x41, 0x52, 0x20}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e, 0x20}, {0x52, 0x53, 0x44, 0x20}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46, 0x20}, {0x53, 0x41, 0x52, 0x20}, {0x53, 0x42, 0x44, 0x20}, {0x53, 0x43, 0x52, 0x20}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47, 0x20}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b, 0x20}, {0x53, 0x47, 0x44, 0x20}, {0x53, 0x48, 0x50, 0x20}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c, 0x20}, {0x53, 0x4f, 0x53, 0x20}, {0x53, 0x52, 0x44, 0x20}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50, 0x20}, {0x53, 0x54, 0x44, 0x20}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50, 0x20}, {0x53, 0x5a, 0x4c, 0x20}, {0x54, 0x48, 0x42, 0x20}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53, 0x20}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54, 0x20}, {0x54, 0x4e, 0x44, 0x20}, {0x54, 0x4f, 0x50, 0x20}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59, 0x20}, {0x54, 0x54, 0x44, 0x20}, {0x54, 0x57, 0x44, 0x20}, {0x54, 0x5a, 0x53, 0x20}, {0x55, 0x41, 0x48, 0x20}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58, 0x20}, {0x24}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55, 0x20}, {0x55, 0x5a, 0x53, 0x20}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46, 0x20}, {0x56, 0x4e, 0x44, 0x20}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56, 0x20}, {0x57, 0x53, 0x54, 0x20}, {0x58, 0x41, 0x46, 0x20}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x58, 0x43, 0x44, 0x20}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x58, 0x4f, 0x46, 0x20}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46, 0x20}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52, 0x20}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52, 0x20}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57, 0x20}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0x4b},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0x4b},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0xd7, 0x99, 0xd7, 0x90, 0xd6, 0xb7, 0xd7, 0xa0, 0xd7, 0x95, 0xd7, 0x90, 0xd6, 0xb7, 0xd7, 0xa8}, {0xd7, 0xa4, 0xd6, 0xbf, 0xd7, 0xa2, 0xd7, 0x91, 0xd7, 0xa8, 0xd7, 0x95, 0xd7, 0x90, 0xd6, 0xb7, 0xd7, 0xa8}, {0xd7, 0x9e, 0xd7, 0xa2, 0xd7, 0xa8, 0xd7, 0xa5}, {0xd7, 0x90, 0xd6, 0xb7, 0xd7, 0xa4, 0xd6, 0xbc, 0xd7, 0xa8, 0xd7, 0x99, 0xd7, 0x9c}, {0xd7, 0x9e, 0xd7, 0x99, 0xd7, 0x99}, {0xd7, 0x99, 0xd7, 0x95, 0xd7, 0xa0, 0xd7, 0x99}, {0xd7, 0x99, 0xd7, 0x95, 0xd7, 0x9c, 0xd7, 0x99}, {0xd7, 0x90, 0xd7, 0x95, 0xd7, 0x99, 0xd7, 0x92, 0xd7, 0x95, 0xd7, 0xa1, 0xd7, 0x98}, {0xd7, 0xa1, 0xd7, 0xa2, 0xd7, 0xa4, 0xd6, 0xbc, 0xd7, 0x98, 0xd7, 0xa2, 0xd7, 0x9e, 0xd7, 0x91, 0xd7, 0xa2, 0xd7, 0xa8}, {0xd7, 0x90, 0xd7, 0xa7, 0xd7, 0x98, 0xd7, 0x90, 0xd7, 0x91, 0xd7, 0xa2, 0xd7, 0xa8}, {0xd7, 0xa0, 0xd7, 0x90, 0xd7, 0x95, 0xd7, 0x95, 0xd7, 0xa2, 0xd7, 0x9e, 0xd7, 0x91, 0xd7, 0xa2, 0xd7, 0xa8}, {0xd7, 0x93, 0xd7, 0xa2, 0xd7, 0xa6, 0xd7, 0xa2, 0xd7, 0x9e, 0xd7, 0x91, 0xd7, 0xa2, 0xd7, 0xa8}},
		monthsWide:             [][]uint8{[]uint8(nil), {0xd7, 0x99, 0xd7, 0x90, 0xd6, 0xb7, 0xd7, 0xa0, 0xd7, 0x95, 0xd7, 0x90, 0xd6, 0xb7, 0xd7, 0xa8}, {0xd7, 0xa4, 0xd6, 0xbf, 0xd7, 0xa2, 0xd7, 0x91, 0xd7, 0xa8, 0xd7, 0x95, 0xd7, 0x90, 0xd6, 0xb7, 0xd7, 0xa8}, {0xd7, 0x9e, 0xd7, 0xa2, 0xd7, 0xa8, 0xd7, 0xa5}, {0xd7, 0x90, 0xd6, 0xb7, 0xd7, 0xa4, 0xd6, 0xbc, 0xd7, 0xa8, 0xd7, 0x99, 0xd7, 0x9c}, {0xd7, 0x9e, 0xd7, 0x99, 0xd7, 0x99}, {0xd7, 0x99, 0xd7, 0x95, 0xd7, 0xa0, 0xd7, 0x99}, {0xd7, 0x99, 0xd7, 0x95, 0xd7, 0x9c, 0xd7, 0x99}, {0xd7, 0x90, 0xd7, 0x95, 0xd7, 0x99, 0xd7, 0x92, 0xd7, 0x95, 0xd7, 0xa1, 0xd7, 0x98}, {0xd7, 0xa1, 0xd7, 0xa2, 0xd7, 0xa4, 0xd6, 0xbc, 0xd7, 0x98, 0xd7, 0xa2, 0xd7, 0x9e, 0xd7, 0x91, 0xd7, 0xa2, 0xd7, 0xa8}, {0xd7, 0x90, 0xd7, 0xa7, 0xd7, 0x98, 0xd7, 0x90, 0xd7, 0x91, 0xd7, 0xa2, 0xd7, 0xa8}, {0xd7, 0xa0, 0xd7, 0x90, 0xd7, 0x95, 0xd7, 0x95, 0xd7, 0xa2, 0xd7, 0x9e, 0xd7, 0x91, 0xd7, 0xa2, 0xd7, 0xa8}, {0xd7, 0x93, 0xd7, 0xa2, 0xd7, 0xa6, 0xd7, 0xa2, 0xd7, 0x9e, 0xd7, 0x91, 0xd7, 0xa2, 0xd7, 0xa8}},
		daysAbbreviated:        [][]uint8{{0xd7, 0x96, 0xd7, 0x95, 0xd7, 0xa0, 0xd7, 0x98, 0xd7, 0x99, 0xd7, 0xa7}, {0xd7, 0x9e, 0xd7, 0x90, 0xd6, 0xb8, 0xd7, 0xa0, 0xd7, 0x98, 0xd7, 0x99, 0xd7, 0xa7}, {0xd7, 0x93, 0xd7, 0x99, 0xd7, 0xa0, 0xd7, 0xa1, 0xd7, 0x98, 0xd7, 0x99, 0xd7, 0xa7}, {0xd7, 0x9e, 0xd7, 0x99, 0xd7, 0x98, 0xd7, 0x95, 0xd7, 0x95, 0xd7, 0x90, 0xd7, 0x9a}, {0xd7, 0x93, 0xd7, 0x90, 0xd7, 0xa0, 0xd7, 0xa2, 0xd7, 0xa8, 0xd7, 0xa9, 0xd7, 0x98, 0xd7, 0x99, 0xd7, 0xa7}, {0xd7, 0xa4, 0xd6, 0xbf, 0xd7, 0xa8, 0xd7, 0xb2, 0xd6, 0xb7, 0xd7, 0x98, 0xd7, 0x99, 0xd7, 0xa7}, {0xd7, 0xa9, 0xd7, 0x91, 0xd7, 0xaa}},
		daysShort:              [][]uint8{{0xd7, 0x96, 0xd7, 0x95, 0xd7, 0xa0, 0xd7, 0x98, 0xd7, 0x99, 0xd7, 0xa7}, {0xd7, 0x9e, 0xd7, 0x90, 0xd6, 0xb8, 0xd7, 0xa0, 0xd7, 0x98, 0xd7, 0x99, 0xd7, 0xa7}, {0xd7, 0x93, 0xd7, 0x99, 0xd7, 0xa0, 0xd7, 0xa1, 0xd7, 0x98, 0xd7, 0x99, 0xd7, 0xa7}, {0xd7, 0x9e, 0xd7, 0x99, 0xd7, 0x98, 0xd7, 0x95, 0xd7, 0x95, 0xd7, 0x90, 0xd7, 0x9a}, {0xd7, 0x93, 0xd7, 0x90, 0xd7, 0xa0, 0xd7, 0xa2, 0xd7, 0xa8, 0xd7, 0xa9, 0xd7, 0x98, 0xd7, 0x99, 0xd7, 0xa7}, {0xd7, 0xa4, 0xd6, 0xbf, 0xd7, 0xa8, 0xd7, 0xb2, 0xd6, 0xb7, 0xd7, 0x98, 0xd7, 0x99, 0xd7, 0xa7}, {0xd7, 0xa9, 0xd7, 0x91, 0xd7, 0xaa}},
		daysWide:               [][]uint8{{0xd7, 0x96, 0xd7, 0x95, 0xd7, 0xa0, 0xd7, 0x98, 0xd7, 0x99, 0xd7, 0xa7}, {0xd7, 0x9e, 0xd7, 0x90, 0xd6, 0xb8, 0xd7, 0xa0, 0xd7, 0x98, 0xd7, 0x99, 0xd7, 0xa7}, {0xd7, 0x93, 0xd7, 0x99, 0xd7, 0xa0, 0xd7, 0xa1, 0xd7, 0x98, 0xd7, 0x99, 0xd7, 0xa7}, {0xd7, 0x9e, 0xd7, 0x99, 0xd7, 0x98, 0xd7, 0x95, 0xd7, 0x95, 0xd7, 0x90, 0xd7, 0x9a}, {0xd7, 0x93, 0xd7, 0x90, 0xd7, 0xa0, 0xd7, 0xa2, 0xd7, 0xa8, 0xd7, 0xa9, 0xd7, 0x98, 0xd7, 0x99, 0xd7, 0xa7}, {0xd7, 0xa4, 0xd6, 0xbf, 0xd7, 0xa8, 0xd7, 0xb2, 0xd6, 0xb7, 0xd7, 0x98, 0xd7, 0x99, 0xd7, 0xa7}, {0xd7, 0xa9, 0xd7, 0x91, 0xd7, 0xaa}},
		periodsAbbreviated:     [][]uint8{{0xd7, 0xa4, 0xd6, 0xbf, 0xd7, 0x90, 0xd6, 0xb7, 0xd7, 0xa8, 0xd7, 0x9e, 0xd7, 0x99, 0xd7, 0x98, 0xd7, 0x90, 0xd6, 0xb8, 0xd7, 0x92}, {0xd7, 0xa0, 0xd7, 0x90, 0xd6, 0xb8, 0xd7, 0x9b, 0xd7, 0x9e, 0xd7, 0x99, 0xd7, 0x98, 0xd7, 0x90, 0xd6, 0xb8, 0xd7, 0x92}},
		periodsWide:            [][]uint8{{0xd7, 0xa4, 0xd6, 0xbf, 0xd7, 0x90, 0xd6, 0xb7, 0xd7, 0xa8, 0xd7, 0x9e, 0xd7, 0x99, 0xd7, 0x98, 0xd7, 0x90, 0xd6, 0xb8, 0xd7, 0x92}, {0xd7, 0xa0, 0xd7, 0x90, 0xd6, 0xb8, 0xd7, 0x9b, 0xd7, 0x9e, 0xd7, 0x99, 0xd7, 0x98, 0xd7, 0x90, 0xd6, 0xb8, 0xd7, 0x92}},
		timezones:              map[string][]uint8{"GYT": {0x47, 0x59, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "EST": {0x45, 0x53, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "WAT": {0x57, 0x41, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}, "EDT": {0x45, 0x44, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "BT": {0x42, 0x54}, "WIB": {0x57, 0x49, 0x42}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "IST": {0x49, 0x53, 0x54}, "VET": {0x56, 0x45, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "OEZ": {0x4f, 0x45, 0x5a}, "GFT": {0x47, 0x46, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "SRT": {0x53, 0x52, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "MYT": {0x4d, 0x59, 0x54}, "CAT": {0x43, 0x41, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "PDT": {0x50, 0x44, 0x54}, "SGT": {0x53, 0x47, 0x54}, "UYT": {0x55, 0x59, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "COT": {0x43, 0x4f, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "CST": {0x43, 0x53, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "HNT": {0x48, 0x4e, 0x54}, "HAT": {0x48, 0x41, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "ECT": {0x45, 0x43, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "AST": {0x41, 0x53, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "WIT": {0x57, 0x49, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "PST": {0x50, 0x53, 0x54}, "ART": {0x41, 0x52, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "EAT": {0x45, 0x41, 0x54}, "CDT": {0x43, 0x44, 0x54}, "MST": {0x4d, 0x53, 0x54}, "ADT": {0x41, 0x44, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}},
	}
}

// Locale returns the current translators string locale
func (yi *yi) Locale() string {
	return yi.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'yi'
func (yi *yi) PluralsCardinal() []locales.PluralRule {
	return yi.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'yi'
func (yi *yi) PluralsOrdinal() []locales.PluralRule {
	return yi.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'yi'
func (yi *yi) PluralsRange() []locales.PluralRule {
	return yi.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'yi'
func (yi *yi) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'yi'
func (yi *yi) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'yi'
func (yi *yi) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (yi *yi) MonthAbbreviated(month time.Month) []byte {
	return yi.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (yi *yi) MonthsAbbreviated() [][]byte {
	return yi.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (yi *yi) MonthNarrow(month time.Month) []byte {
	return yi.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (yi *yi) MonthsNarrow() [][]byte {
	return yi.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (yi *yi) MonthWide(month time.Month) []byte {
	return yi.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (yi *yi) MonthsWide() [][]byte {
	return yi.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (yi *yi) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return yi.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (yi *yi) WeekdaysAbbreviated() [][]byte {
	return yi.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (yi *yi) WeekdayNarrow(weekday time.Weekday) []byte {
	return yi.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (yi *yi) WeekdaysNarrow() [][]byte {
	return yi.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (yi *yi) WeekdayShort(weekday time.Weekday) []byte {
	return yi.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (yi *yi) WeekdaysShort() [][]byte {
	return yi.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (yi *yi) WeekdayWide(weekday time.Weekday) []byte {
	return yi.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (yi *yi) WeekdaysWide() [][]byte {
	return yi.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'yi' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yi *yi) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'yi' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (yi *yi) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'yi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yi *yi) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := yi.currencies[currency]
	l := len(s) + len(yi.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yi.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(yi.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, yi.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, yi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, yi.currencyPositiveSuffix...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'yi'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yi *yi) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := yi.currencies[currency]
	l := len(s) + len(yi.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yi.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(yi.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, yi.currencyNegativePrefix[j])
		}

		b = append(b, yi.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(yi.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, yi.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, yi.currencyNegativeSuffix...)
	} else {

		b = append(b, yi.currencyPositiveSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'yi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yi *yi) FmtDateShort(t time.Time) []byte {

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

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'yi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yi *yi) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xd7, 0x98, 0xd7, 0x9f, 0x20}...)
	b = append(b, yi.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'yi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yi *yi) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xd7, 0x98, 0xd7, 0x9f, 0x20}...)
	b = append(b, yi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'yi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yi *yi) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, yi.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xd7, 0x98, 0xd7, 0x9f, 0x20}...)
	b = append(b, yi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'yi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yi *yi) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'yi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yi *yi) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'yi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yi *yi) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'yi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (yi *yi) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := yi.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
