package ko_KP

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ko_KP struct {
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

// New returns a new instance of translator for the 'ko_KP' locale
func New() locales.Translator {
	return &ko_KP{
		locale:                 "ko_KP",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
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
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x31, 0xec, 0x9b, 0x94}, {0x32, 0xec, 0x9b, 0x94}, {0x33, 0xec, 0x9b, 0x94}, {0x34, 0xec, 0x9b, 0x94}, {0x35, 0xec, 0x9b, 0x94}, {0x36, 0xec, 0x9b, 0x94}, {0x37, 0xec, 0x9b, 0x94}, {0x38, 0xec, 0x9b, 0x94}, {0x39, 0xec, 0x9b, 0x94}, {0x31, 0x30, 0xec, 0x9b, 0x94}, {0x31, 0x31, 0xec, 0x9b, 0x94}, {0x31, 0x32, 0xec, 0x9b, 0x94}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x31, 0xec, 0x9b, 0x94}, {0x32, 0xec, 0x9b, 0x94}, {0x33, 0xec, 0x9b, 0x94}, {0x34, 0xec, 0x9b, 0x94}, {0x35, 0xec, 0x9b, 0x94}, {0x36, 0xec, 0x9b, 0x94}, {0x37, 0xec, 0x9b, 0x94}, {0x38, 0xec, 0x9b, 0x94}, {0x39, 0xec, 0x9b, 0x94}, {0x31, 0x30, 0xec, 0x9b, 0x94}, {0x31, 0x31, 0xec, 0x9b, 0x94}, {0x31, 0x32, 0xec, 0x9b, 0x94}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x31, 0xec, 0x9b, 0x94}, {0x32, 0xec, 0x9b, 0x94}, {0x33, 0xec, 0x9b, 0x94}, {0x34, 0xec, 0x9b, 0x94}, {0x35, 0xec, 0x9b, 0x94}, {0x36, 0xec, 0x9b, 0x94}, {0x37, 0xec, 0x9b, 0x94}, {0x38, 0xec, 0x9b, 0x94}, {0x39, 0xec, 0x9b, 0x94}, {0x31, 0x30, 0xec, 0x9b, 0x94}, {0x31, 0x31, 0xec, 0x9b, 0x94}, {0x31, 0x32, 0xec, 0x9b, 0x94}},
		daysAbbreviated:        [][]uint8{{0xec, 0x9d, 0xbc}, {0xec, 0x9b, 0x94}, {0xed, 0x99, 0x94}, {0xec, 0x88, 0x98}, {0xeb, 0xaa, 0xa9}, {0xea, 0xb8, 0x88}, {0xed, 0x86, 0xa0}},
		daysNarrow:             [][]uint8{{0xec, 0x9d, 0xbc}, {0xec, 0x9b, 0x94}, {0xed, 0x99, 0x94}, {0xec, 0x88, 0x98}, {0xeb, 0xaa, 0xa9}, {0xea, 0xb8, 0x88}, {0xed, 0x86, 0xa0}},
		daysShort:              [][]uint8{{0xec, 0x9d, 0xbc}, {0xec, 0x9b, 0x94}, {0xed, 0x99, 0x94}, {0xec, 0x88, 0x98}, {0xeb, 0xaa, 0xa9}, {0xea, 0xb8, 0x88}, {0xed, 0x86, 0xa0}},
		daysWide:               [][]uint8{{0xec, 0x9d, 0xbc, 0xec, 0x9a, 0x94, 0xec, 0x9d, 0xbc}, {0xec, 0x9b, 0x94, 0xec, 0x9a, 0x94, 0xec, 0x9d, 0xbc}, {0xed, 0x99, 0x94, 0xec, 0x9a, 0x94, 0xec, 0x9d, 0xbc}, {0xec, 0x88, 0x98, 0xec, 0x9a, 0x94, 0xec, 0x9d, 0xbc}, {0xeb, 0xaa, 0xa9, 0xec, 0x9a, 0x94, 0xec, 0x9d, 0xbc}, {0xea, 0xb8, 0x88, 0xec, 0x9a, 0x94, 0xec, 0x9d, 0xbc}, {0xed, 0x86, 0xa0, 0xec, 0x9a, 0x94, 0xec, 0x9d, 0xbc}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsWide:            [][]uint8{{0xec, 0x98, 0xa4, 0xec, 0xa0, 0x84}, {0xec, 0x98, 0xa4, 0xed, 0x9b, 0x84}},
		erasAbbreviated:        [][]uint8{{0x42, 0x43}, {0x41, 0x44}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0xea, 0xb8, 0xb0, 0xec, 0x9b, 0x90, 0xec, 0xa0, 0x84}, {0xec, 0x84, 0x9c, 0xea, 0xb8, 0xb0}},
		timezones:              map[string][]uint8{"WARST": {0xec, 0x95, 0x84, 0xeb, 0xa5, 0xb4, 0xed, 0x97, 0xa8, 0xed, 0x8b, 0xb0, 0xeb, 0x82, 0x98, 0x20, 0xec, 0x84, 0x9c, 0xeb, 0xb6, 0x80, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "ACST": {0xec, 0x98, 0xa4, 0xec, 0x8a, 0xa4, 0xed, 0x8a, 0xb8, 0xeb, 0xa0, 0x88, 0xec, 0x9d, 0xbc, 0xeb, 0xa6, 0xac, 0xec, 0x95, 0x84, 0x20, 0xec, 0xa4, 0x91, 0xeb, 0xb6, 0x80, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "AEST": {0xec, 0x98, 0xa4, 0xec, 0x8a, 0xa4, 0xed, 0x8a, 0xb8, 0xeb, 0xa0, 0x88, 0xec, 0x9d, 0xbc, 0xeb, 0xa6, 0xac, 0xec, 0x95, 0x84, 0x20, 0xeb, 0x8f, 0x99, 0xeb, 0xb6, 0x80, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "ART": {0xec, 0x95, 0x84, 0xeb, 0xa5, 0xb4, 0xed, 0x97, 0xa8, 0xed, 0x8b, 0xb0, 0xeb, 0x82, 0x98, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "UYST": {0xec, 0x9a, 0xb0, 0xeb, 0xa3, 0xa8, 0xea, 0xb3, 0xbc, 0xec, 0x9d, 0xb4, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "ACWDT": {0xec, 0x98, 0xa4, 0xec, 0x8a, 0xa4, 0xed, 0x8a, 0xb8, 0xeb, 0xa0, 0x88, 0xec, 0x9d, 0xbc, 0xeb, 0xa6, 0xac, 0xec, 0x95, 0x84, 0x20, 0xec, 0xa4, 0x91, 0xec, 0x84, 0x9c, 0xeb, 0xb6, 0x80, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "MDT": {0xeb, 0xaf, 0xb8, 0x20, 0xec, 0x82, 0xb0, 0xec, 0xa7, 0x80, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "ARST": {0xec, 0x95, 0x84, 0xeb, 0xa5, 0xb4, 0xed, 0x97, 0xa8, 0xed, 0x8b, 0xb0, 0xeb, 0x82, 0x98, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "JDT": {0xec, 0x9d, 0xbc, 0xeb, 0xb3, 0xb8, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "UYT": {0xec, 0x9a, 0xb0, 0xeb, 0xa3, 0xa8, 0xea, 0xb3, 0xbc, 0xec, 0x9d, 0xb4, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "HAT": {0xeb, 0x89, 0xb4, 0xed, 0x8e, 0x80, 0xeb, 0x93, 0xa4, 0xeb, 0x9e, 0x9c, 0xeb, 0x93, 0x9c, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "BT": {0xeb, 0xb6, 0x80, 0xed, 0x83, 0x84, 0x20, 0xec, 0x8b, 0x9c, 0xea, 0xb0, 0x84}, "GMT": {0xea, 0xb7, 0xb8, 0xeb, 0xa6, 0xac, 0xeb, 0x8b, 0x88, 0xec, 0xb9, 0x98, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "IST": {0xec, 0x9d, 0xb8, 0xeb, 0x8f, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "WEZ": {0xec, 0x84, 0x9c, 0xec, 0x9c, 0xa0, 0xeb, 0x9f, 0xbd, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "ACWST": {0xec, 0x98, 0xa4, 0xec, 0x8a, 0xa4, 0xed, 0x8a, 0xb8, 0xeb, 0xa0, 0x88, 0xec, 0x9d, 0xbc, 0xeb, 0xa6, 0xac, 0xec, 0x95, 0x84, 0x20, 0xec, 0xa4, 0x91, 0xec, 0x84, 0x9c, 0xeb, 0xb6, 0x80, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "WIB": {0xec, 0x84, 0x9c, 0xeb, 0xb6, 0x80, 0x20, 0xec, 0x9d, 0xb8, 0xeb, 0x8f, 0x84, 0xeb, 0x84, 0xa4, 0xec, 0x8b, 0x9c, 0xec, 0x95, 0x84, 0x20, 0xec, 0x8b, 0x9c, 0xea, 0xb0, 0x84}, "CHAST": {0xec, 0xb1, 0x84, 0xed, 0x85, 0x80, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "MESZ": {0xec, 0xa4, 0x91, 0xeb, 0xb6, 0x80, 0x20, 0xec, 0x9c, 0xa0, 0xeb, 0x9f, 0xbd, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "ADT": {0xeb, 0xaf, 0xb8, 0x20, 0xeb, 0x8c, 0x80, 0xec, 0x84, 0x9c, 0xec, 0x96, 0x91, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "COT": {0xec, 0xbd, 0x9c, 0xeb, 0xa1, 0xac, 0xeb, 0xb9, 0x84, 0xec, 0x95, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "VET": {0xeb, 0xb2, 0xa0, 0xeb, 0x84, 0xa4, 0xec, 0x88, 0x98, 0xec, 0x97, 0x98, 0xeb, 0x9d, 0xbc, 0x20, 0xec, 0x8b, 0x9c, 0xea, 0xb0, 0x84}, "CHADT": {0xec, 0xb1, 0x84, 0xed, 0x85, 0x80, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "JST": {0xec, 0x9d, 0xbc, 0xeb, 0xb3, 0xb8, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "HKST": {0xed, 0x99, 0x8d, 0xec, 0xbd, 0xa9, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "ACDT": {0xec, 0x98, 0xa4, 0xec, 0x8a, 0xa4, 0xed, 0x8a, 0xb8, 0xeb, 0xa0, 0x88, 0xec, 0x9d, 0xbc, 0xeb, 0xa6, 0xac, 0xec, 0x95, 0x84, 0x20, 0xec, 0xa4, 0x91, 0xeb, 0xb6, 0x80, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "MST": {0xeb, 0xaf, 0xb8, 0x20, 0xec, 0x82, 0xb0, 0xec, 0x95, 0x85, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "NZDT": {0xeb, 0x89, 0xb4, 0xec, 0xa7, 0x88, 0xeb, 0x9e, 0x9c, 0xeb, 0x93, 0x9c, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "HKT": {0xed, 0x99, 0x8d, 0xec, 0xbd, 0xa9, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "AWDT": {0xec, 0x98, 0xa4, 0xec, 0x8a, 0xa4, 0xed, 0x8a, 0xb8, 0xeb, 0xa0, 0x88, 0xec, 0x9d, 0xbc, 0xeb, 0xa6, 0xac, 0xec, 0x95, 0x84, 0x20, 0xec, 0x84, 0x9c, 0xeb, 0xb6, 0x80, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "CST": {0xeb, 0xaf, 0xb8, 0x20, 0xec, 0xa4, 0x91, 0xeb, 0xb6, 0x80, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "HAST": {0xed, 0x95, 0x98, 0xec, 0x99, 0x80, 0xec, 0x9d, 0xb4, 0x20, 0xec, 0x95, 0x8c, 0xeb, 0xa5, 0x98, 0xec, 0x83, 0xa8, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "PDT": {0xeb, 0xaf, 0xb8, 0x20, 0xed, 0x83, 0x9c, 0xed, 0x8f, 0x89, 0xec, 0x96, 0x91, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "EST": {0xeb, 0xaf, 0xb8, 0x20, 0xeb, 0x8f, 0x99, 0xeb, 0xb6, 0x80, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "WESZ": {0xec, 0x84, 0x9c, 0xec, 0x9c, 0xa0, 0xeb, 0x9f, 0xbd, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "WAT": {0xec, 0x84, 0x9c, 0xec, 0x95, 0x84, 0xed, 0x94, 0x84, 0xeb, 0xa6, 0xac, 0xec, 0xb9, 0xb4, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "OEZ": {0xeb, 0x8f, 0x99, 0xec, 0x9c, 0xa0, 0xeb, 0x9f, 0xbd, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "NZST": {0xeb, 0x89, 0xb4, 0xec, 0xa7, 0x88, 0xeb, 0x9e, 0x9c, 0xeb, 0x93, 0x9c, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "ECT": {0xec, 0x97, 0x90, 0xec, 0xbd, 0xb0, 0xeb, 0x8f, 0x84, 0xeb, 0xa5, 0xb4, 0x20, 0xec, 0x8b, 0x9c, 0xea, 0xb0, 0x84}, "CAT": {0xec, 0xa4, 0x91, 0xec, 0x95, 0x99, 0xec, 0x95, 0x84, 0xed, 0x94, 0x84, 0xeb, 0xa6, 0xac, 0xec, 0xb9, 0xb4, 0x20, 0xec, 0x8b, 0x9c, 0xea, 0xb0, 0x84}, "TMT": {0xed, 0x88, 0xac, 0xeb, 0xa5, 0xb4, 0xed, 0x81, 0xac, 0xeb, 0xa9, 0x94, 0xeb, 0x8b, 0x88, 0xec, 0x8a, 0xa4, 0xed, 0x83, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "TMST": {0xed, 0x88, 0xac, 0xeb, 0xa5, 0xb4, 0xed, 0x81, 0xac, 0xeb, 0xa9, 0x94, 0xeb, 0x8b, 0x88, 0xec, 0x8a, 0xa4, 0xed, 0x83, 0x84, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "OESZ": {0xeb, 0x8f, 0x99, 0xec, 0x9c, 0xa0, 0xeb, 0x9f, 0xbd, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "SGT": {0xec, 0x8b, 0xb1, 0xea, 0xb0, 0x80, 0xed, 0x8f, 0xac, 0xeb, 0xa5, 0xb4, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "HNT": {0xeb, 0x89, 0xb4, 0xed, 0x8e, 0x80, 0xeb, 0x93, 0xa4, 0xeb, 0x9e, 0x9c, 0xeb, 0x93, 0x9c, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "AWST": {0xec, 0x98, 0xa4, 0xec, 0x8a, 0xa4, 0xed, 0x8a, 0xb8, 0xeb, 0xa0, 0x88, 0xec, 0x9d, 0xbc, 0xeb, 0xa6, 0xac, 0xec, 0x95, 0x84, 0x20, 0xec, 0x84, 0x9c, 0xeb, 0xb6, 0x80, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "WITA": {0xec, 0xa4, 0x91, 0xeb, 0xb6, 0x80, 0x20, 0xec, 0x9d, 0xb8, 0xeb, 0x8f, 0x84, 0xeb, 0x84, 0xa4, 0xec, 0x8b, 0x9c, 0xec, 0x95, 0x84, 0x20, 0xec, 0x8b, 0x9c, 0xea, 0xb0, 0x84}, "MEZ": {0xec, 0xa4, 0x91, 0xeb, 0xb6, 0x80, 0x20, 0xec, 0x9c, 0xa0, 0xeb, 0x9f, 0xbd, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "EDT": {0xeb, 0xaf, 0xb8, 0x20, 0xeb, 0x8f, 0x99, 0xeb, 0xb6, 0x80, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "∅∅∅": {0xec, 0x95, 0x84, 0xeb, 0xa7, 0x88, 0xec, 0xa1, 0xb4, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "WART": {0xec, 0x95, 0x84, 0xeb, 0xa5, 0xb4, 0xed, 0x97, 0xa8, 0xed, 0x8b, 0xb0, 0xeb, 0x82, 0x98, 0x20, 0xec, 0x84, 0x9c, 0xeb, 0xb6, 0x80, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "AEDT": {0xec, 0x98, 0xa4, 0xec, 0x8a, 0xa4, 0xed, 0x8a, 0xb8, 0xeb, 0xa0, 0x88, 0xec, 0x9d, 0xbc, 0xeb, 0xa6, 0xac, 0xec, 0x95, 0x84, 0x20, 0xeb, 0x8f, 0x99, 0xeb, 0xb6, 0x80, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "AKST": {0xec, 0x95, 0x8c, 0xeb, 0x9e, 0x98, 0xec, 0x8a, 0xa4, 0xec, 0xb9, 0xb4, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "MYT": {0xeb, 0xa7, 0x90, 0xeb, 0xa0, 0x88, 0xec, 0x9d, 0xb4, 0xec, 0x8b, 0x9c, 0xec, 0x95, 0x84, 0x20, 0xec, 0x8b, 0x9c, 0xea, 0xb0, 0x84}, "CDT": {0xeb, 0xaf, 0xb8, 0x20, 0xec, 0xa4, 0x91, 0xeb, 0xb6, 0x80, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "HADT": {0xed, 0x95, 0x98, 0xec, 0x99, 0x80, 0xec, 0x9d, 0xb4, 0x20, 0xec, 0x95, 0x8c, 0xeb, 0xa5, 0x98, 0xec, 0x83, 0xa8, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "AKDT": {0xec, 0x95, 0x8c, 0xeb, 0x9e, 0x98, 0xec, 0x8a, 0xa4, 0xec, 0xb9, 0xb4, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "COST": {0xec, 0xbd, 0x9c, 0xeb, 0xa1, 0xac, 0xeb, 0xb9, 0x84, 0xec, 0x95, 0x84, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "SAST": {0xeb, 0x82, 0xa8, 0xec, 0x95, 0x84, 0xed, 0x94, 0x84, 0xeb, 0xa6, 0xac, 0xec, 0xb9, 0xb4, 0x20, 0xec, 0x8b, 0x9c, 0xea, 0xb0, 0x84}, "PST": {0xeb, 0xaf, 0xb8, 0x20, 0xed, 0x83, 0x9c, 0xed, 0x8f, 0x89, 0xec, 0x96, 0x91, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "AST": {0xeb, 0x8c, 0x80, 0xec, 0x84, 0x9c, 0xec, 0x96, 0x91, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "BOT": {0xeb, 0xb3, 0xbc, 0xeb, 0xa6, 0xac, 0xeb, 0xb9, 0x84, 0xec, 0x95, 0x84, 0x20, 0xec, 0x8b, 0x9c, 0xea, 0xb0, 0x84}, "CLST": {0xec, 0xb9, 0xa0, 0xeb, 0xa0, 0x88, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "GFT": {0xed, 0x94, 0x84, 0xeb, 0x9e, 0x91, 0xec, 0x8a, 0xa4, 0xeb, 0xa0, 0xb9, 0x20, 0xea, 0xb0, 0x80, 0xec, 0x9d, 0xb4, 0xec, 0x95, 0x84, 0xeb, 0x82, 0x98, 0x20, 0xec, 0x8b, 0x9c, 0xea, 0xb0, 0x84}, "WIT": {0xeb, 0x8f, 0x99, 0xeb, 0xb6, 0x80, 0x20, 0xec, 0x9d, 0xb8, 0xeb, 0x8f, 0x84, 0xeb, 0x84, 0xa4, 0xec, 0x8b, 0x9c, 0xec, 0x95, 0x84, 0x20, 0xec, 0x8b, 0x9c, 0xea, 0xb0, 0x84}, "SRT": {0xec, 0x88, 0x98, 0xeb, 0xa6, 0xac, 0xeb, 0x82, 0xa8, 0x20, 0xec, 0x8b, 0x9c, 0xea, 0xb0, 0x84}, "EAT": {0xeb, 0x8f, 0x99, 0xec, 0x95, 0x84, 0xed, 0x94, 0x84, 0xeb, 0xa6, 0xac, 0xec, 0xb9, 0xb4, 0x20, 0xec, 0x8b, 0x9c, 0xea, 0xb0, 0x84}, "LHDT": {0xeb, 0xa1, 0x9c, 0xeb, 0x93, 0x9c, 0x20, 0xed, 0x95, 0x98, 0xec, 0x9a, 0xb0, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "GYT": {0xea, 0xb0, 0x80, 0xec, 0x9d, 0xb4, 0xec, 0x95, 0x84, 0xeb, 0x82, 0x98, 0x20, 0xec, 0x8b, 0x9c, 0xea, 0xb0, 0x84}, "WAST": {0xec, 0x84, 0x9c, 0xec, 0x95, 0x84, 0xed, 0x94, 0x84, 0xeb, 0xa6, 0xac, 0xec, 0xb9, 0xb4, 0x20, 0xed, 0x95, 0x98, 0xea, 0xb3, 0x84, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "CLT": {0xec, 0xb9, 0xa0, 0xeb, 0xa0, 0x88, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}, "ChST": {0xec, 0xb0, 0xa8, 0xeb, 0xaa, 0xa8, 0xeb, 0xa1, 0x9c, 0x20, 0xec, 0x8b, 0x9c, 0xea, 0xb0, 0x84}, "LHST": {0xeb, 0xa1, 0x9c, 0xeb, 0x93, 0x9c, 0x20, 0xed, 0x95, 0x98, 0xec, 0x9a, 0xb0, 0x20, 0xed, 0x91, 0x9c, 0xec, 0xa4, 0x80, 0xec, 0x8b, 0x9c}},
	}
}

// Locale returns the current translators string locale
func (ko *ko_KP) Locale() string {
	return ko.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ko_KP'
func (ko *ko_KP) PluralsCardinal() []locales.PluralRule {
	return ko.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ko_KP'
func (ko *ko_KP) PluralsOrdinal() []locales.PluralRule {
	return ko.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ko_KP'
func (ko *ko_KP) PluralsRange() []locales.PluralRule {
	return ko.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ko_KP'
func (ko *ko_KP) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ko_KP'
func (ko *ko_KP) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ko_KP'
func (ko *ko_KP) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ko *ko_KP) MonthAbbreviated(month time.Month) []byte {
	return ko.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ko *ko_KP) MonthsAbbreviated() [][]byte {
	return ko.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ko *ko_KP) MonthNarrow(month time.Month) []byte {
	return ko.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ko *ko_KP) MonthsNarrow() [][]byte {
	return ko.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ko *ko_KP) MonthWide(month time.Month) []byte {
	return ko.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ko *ko_KP) MonthsWide() [][]byte {
	return ko.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ko *ko_KP) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return ko.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ko *ko_KP) WeekdaysAbbreviated() [][]byte {
	return ko.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ko *ko_KP) WeekdayNarrow(weekday time.Weekday) []byte {
	return ko.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ko *ko_KP) WeekdaysNarrow() [][]byte {
	return ko.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ko *ko_KP) WeekdayShort(weekday time.Weekday) []byte {
	return ko.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ko *ko_KP) WeekdaysShort() [][]byte {
	return ko.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ko *ko_KP) WeekdayWide(weekday time.Weekday) []byte {
	return ko.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ko *ko_KP) WeekdaysWide() [][]byte {
	return ko.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ko_KP' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ko *ko_KP) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ko.decimal) + len(ko.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ko.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ko.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ko.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ko_KP' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ko *ko_KP) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ko.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ko.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ko.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ko.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ko_KP'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ko *ko_KP) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ko.currencies[currency]
	l := len(s) + len(ko.decimal) + len(ko.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ko.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ko.group[0])
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
		b = append(b, ko.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ko.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ko_KP'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ko *ko_KP) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ko.currencies[currency]
	l := len(s) + len(ko.decimal) + len(ko.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ko.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ko.group[0])
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

		for j := len(ko.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ko.currencyNegativePrefix[j])
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
			b = append(b, ko.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ko.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'ko_KP'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ko *ko_KP) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'ko_KP'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ko *ko_KP) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'ko_KP'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ko *ko_KP) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0xeb, 0x85, 0x84, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0xec, 0x9b, 0x94, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xec, 0x9d, 0xbc}...)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'ko_KP'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ko *ko_KP) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0xeb, 0x85, 0x84, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0xec, 0x9b, 0x94, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xec, 0x9d, 0xbc, 0x20}...)
	b = append(b, ko.daysWide[t.Weekday()]...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'ko_KP'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ko *ko_KP) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ko.periodsAbbreviated[0]...)
	} else {
		b = append(b, ko.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ko.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'ko_KP'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ko *ko_KP) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ko.periodsAbbreviated[0]...)
	} else {
		b = append(b, ko.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ko.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ko.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'ko_KP'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ko *ko_KP) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ko.periodsAbbreviated[0]...)
	} else {
		b = append(b, ko.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0xec, 0x8b, 0x9c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xeb, 0xb6, 0x84, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xec, 0xb4, 0x88, 0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'ko_KP'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ko *ko_KP) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ko.periodsAbbreviated[0]...)
	} else {
		b = append(b, ko.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0xec, 0x8b, 0x9c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xeb, 0xb6, 0x84, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xec, 0xb4, 0x88, 0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ko.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
