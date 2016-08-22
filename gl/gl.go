package gl

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type gl struct {
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
	currencyPositiveSuffix []byte
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

// New returns a new instance of translator for the 'gl' locale
func New() locales.Translator {
	return &gl{
		locale:                 "gl",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6, 2},
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53, 0x20}, {0x24, 0x41}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x24, 0x52}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x24, 0x43, 0x41}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0xc2, 0xa5}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0xe2, 0x82, 0xa7}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46, 0x20}, {0xc2, 0xa3}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44}, {0x24, 0x48, 0x4b}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0xe2, 0x82, 0xaa}, {0xe2, 0x82, 0xb9}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0xc2, 0xa5, 0x4a, 0x50}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0xe2, 0x82, 0xa9}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x24, 0x4d, 0x58}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x24}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0xe0, 0xb8, 0xbf}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x24, 0x4e, 0x54}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58}, {0x24}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46}, {0xe2, 0x82, 0xab}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x45, 0x43, 0x24}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44, 0x20}, {0x43, 0x46, 0x50, 0x46}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x78, 0x61, 0x6e}, {0x66, 0x65, 0x62}, {0x6d, 0x61, 0x72}, {0x61, 0x62, 0x72}, {0x6d, 0x61, 0x69}, {0x78, 0x75, 0xc3, 0xb1}, {0x78, 0x75, 0x6c}, {0x61, 0x67, 0x6f}, {0x73, 0x65, 0x74}, {0x6f, 0x75, 0x74}, {0x6e, 0x6f, 0x76}, {0x64, 0x65, 0x63}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x58}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x58}, {0x58}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x78, 0x61, 0x6e, 0x65, 0x69, 0x72, 0x6f}, {0x66, 0x65, 0x62, 0x72, 0x65, 0x69, 0x72, 0x6f}, {0x6d, 0x61, 0x72, 0x7a, 0x6f}, {0x61, 0x62, 0x72, 0x69, 0x6c}, {0x6d, 0x61, 0x69, 0x6f}, {0x78, 0x75, 0xc3, 0xb1, 0x6f}, {0x78, 0x75, 0x6c, 0x6c, 0x6f}, {0x61, 0x67, 0x6f, 0x73, 0x74, 0x6f}, {0x73, 0x65, 0x74, 0x65, 0x6d, 0x62, 0x72, 0x6f}, {0x6f, 0x75, 0x74, 0x75, 0x62, 0x72, 0x6f}, {0x6e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x72, 0x6f}, {0x64, 0x65, 0x63, 0x65, 0x6d, 0x62, 0x72, 0x6f}},
		daysAbbreviated:        [][]uint8{{0x64, 0x6f, 0x6d}, {0x6c, 0x75, 0x6e, 0x73}, {0x6d, 0x61, 0x72}, {0x6d, 0xc3, 0xa9, 0x72}, {0x78, 0x6f, 0x76}, {0x76, 0x65, 0x6e}, {0x73, 0xc3, 0xa1, 0x62}},
		daysNarrow:             [][]uint8{{0x44}, {0x4c}, {0x4d}, {0x4d}, {0x58}, {0x56}, {0x53}},
		daysShort:              [][]uint8{{0x64, 0x6f, 0x6d}, {0x6c, 0x75, 0x6e, 0x73}, {0x6d, 0x74}, {0x6d, 0x63}, {0x78, 0x76}, {0x76, 0x65}, {0x73, 0xc3, 0xa1, 0x62}},
		daysWide:               [][]uint8{{0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x6f}, {0x6c, 0x75, 0x6e, 0x73}, {0x6d, 0x61, 0x72, 0x74, 0x65, 0x73}, {0x6d, 0xc3, 0xa9, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x73}, {0x78, 0x6f, 0x76, 0x65, 0x73}, {0x76, 0x65, 0x6e, 0x72, 0x65, 0x73}, {0x73, 0xc3, 0xa1, 0x62, 0x61, 0x64, 0x6f}},
		periodsAbbreviated:     [][]uint8{{0x61, 0x2e, 0x6d, 0x2e}, {0x70, 0x2e, 0x6d, 0x2e}},
		periodsNarrow:          [][]uint8{{0x61, 0x2e, 0x6d, 0x2e}, {0x70, 0x2e, 0x6d, 0x2e}},
		periodsWide:            [][]uint8{{0x61, 0x2e, 0x6d, 0x2e}, {0x70, 0x2e, 0x6d, 0x2e}},
		erasAbbreviated:        [][]uint8{{0x61, 0x2e, 0x43, 0x2e}, {0x64, 0x2e, 0x43, 0x2e}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x61, 0x6e, 0x74, 0x65, 0x73, 0x20, 0x64, 0x65, 0x20, 0x43, 0x72, 0x69, 0x73, 0x74, 0x6f}, {0x64, 0x65, 0x73, 0x70, 0x6f, 0x69, 0x73, 0x20, 0x64, 0x65, 0x20, 0x43, 0x72, 0x69, 0x73, 0x74, 0x6f}},
		timezones:              map[string][]uint8{"ACDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "MST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x61, 0x73, 0x20, 0x6d, 0x6f, 0x6e, 0x74, 0x61, 0xc3, 0xb1, 0x61, 0x73, 0x20, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x61, 0x73}, "ARST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x72, 0x78, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "HKT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "CHADT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "WITA": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "JDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x58, 0x61, 0x70, 0xc3, 0xb3, 0x6e}, "UYT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69}, "MDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x61, 0x73, 0x20, 0x6d, 0x6f, 0x6e, 0x74, 0x61, 0xc3, 0xb1, 0x61, 0x73, 0x20, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x61, 0x73}, "PST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x6f, 0x20, 0x50, 0x61, 0x63, 0xc3, 0xad, 0x66, 0x69, 0x63, 0x6f}, "ADT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x6f, 0x20, 0x41, 0x74, 0x6c, 0xc3, 0xa1, 0x6e, 0x74, 0x69, 0x63, 0x6f}, "ART": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x72, 0x78, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "AKDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x63, 0x61}, "WEZ": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "LHST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "WART": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x72, 0x78, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "WARST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x72, 0x78, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "CAT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "GFT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x47, 0xc3, 0xbc, 0x69, 0x61, 0x6e, 0x61, 0x20, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x61}, "EST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x41, 0x6d, 0xc3, 0xa9, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "MYT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x4d, 0x61, 0x6c, 0x61, 0x69, 0x73, 0x69, 0x61}, "GYT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x61, 0x20, 0x47, 0xc3, 0xbc, 0x69, 0x61, 0x6e, 0x61}, "WIB": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "CHAST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "SAST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x53, 0x75, 0x64, 0xc3, 0xa1, 0x66, 0x72, 0x69, 0x63, 0x61}, "IST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x61, 0x20, 0x49, 0x6e, 0x64, 0x69, 0x61}, "ACST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "COST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "PDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x6f, 0x20, 0x50, 0x61, 0x63, 0xc3, 0xad, 0x66, 0x69, 0x63, 0x6f}, "OESZ": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "GMT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x6f, 0x20, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x69, 0x61, 0x6e, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68}, "WIT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "EDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x6d, 0xc3, 0xa9, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "HAT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x54, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x6f, 0x76, 0x61}, "VET": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61}, "BT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x42, 0x75, 0x74, 0xc3, 0xa1, 0x6e}, "NZDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x4e, 0x6f, 0x76, 0x61, 0x20, 0x43, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x61}, "WESZ": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "ACWDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "AWST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "∅∅∅": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x6f, 0x20, 0x41, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x61, 0x73}, "EAT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "HADT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x6f}, "ACWST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "WAT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "AEDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "CDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x61, 0x20, 0x7a, 0x6f, 0x6e, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "UYST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69}, "AWDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "MESZ": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "OEZ": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "JST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x58, 0x61, 0x70, 0xc3, 0xb3, 0x6e}, "AKST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x63, 0x61}, "ChST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "WAST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "COT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "TMT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x54, 0x75, 0x72, 0x63, 0x6f, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0xc3, 0xa1, 0x6e}, "HNT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x54, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x6f, 0x76, 0x61}, "ECT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72}, "HAST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x6f}, "SGT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72}, "CLT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "MEZ": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "CLST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "NZST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x4e, 0x6f, 0x76, 0x61, 0x20, 0x43, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x61}, "LHDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "AEST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "CST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "TMST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x54, 0x75, 0x72, 0x63, 0x6f, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0xc3, 0xa1, 0x6e}, "AST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x6f, 0x20, 0x41, 0x74, 0x6c, 0xc3, 0xa1, 0x6e, 0x74, 0x69, 0x63, 0x6f}, "HKST": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x76, 0x65, 0x72, 0xc3, 0xa1, 0x6e, 0x20, 0x64, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "SRT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d}, "BOT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x6f, 0x20, 0x64, 0x65, 0x20, 0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61}},
	}
}

// Locale returns the current translators string locale
func (gl *gl) Locale() string {
	return gl.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'gl'
func (gl *gl) PluralsCardinal() []locales.PluralRule {
	return gl.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'gl'
func (gl *gl) PluralsOrdinal() []locales.PluralRule {
	return gl.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'gl'
func (gl *gl) PluralsRange() []locales.PluralRule {
	return gl.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'gl'
func (gl *gl) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'gl'
func (gl *gl) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'gl'
func (gl *gl) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := gl.CardinalPluralRule(num1, v1)
	end := gl.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (gl *gl) MonthAbbreviated(month time.Month) []byte {
	return gl.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (gl *gl) MonthsAbbreviated() [][]byte {
	return gl.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (gl *gl) MonthNarrow(month time.Month) []byte {
	return gl.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (gl *gl) MonthsNarrow() [][]byte {
	return gl.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (gl *gl) MonthWide(month time.Month) []byte {
	return gl.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (gl *gl) MonthsWide() [][]byte {
	return gl.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (gl *gl) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return gl.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (gl *gl) WeekdaysAbbreviated() [][]byte {
	return gl.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (gl *gl) WeekdayNarrow(weekday time.Weekday) []byte {
	return gl.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (gl *gl) WeekdaysNarrow() [][]byte {
	return gl.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (gl *gl) WeekdayShort(weekday time.Weekday) []byte {
	return gl.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (gl *gl) WeekdaysShort() [][]byte {
	return gl.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (gl *gl) WeekdayWide(weekday time.Weekday) []byte {
	return gl.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (gl *gl) WeekdaysWide() [][]byte {
	return gl.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'gl' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(gl.decimal) + len(gl.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, gl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, gl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'gl' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (gl *gl) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(gl.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gl.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, gl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, gl.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'gl'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := gl.currencies[currency]
	l := len(s) + len(gl.decimal) + len(gl.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, gl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, gl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, gl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, gl.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'gl'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := gl.currencies[currency]
	l := len(s) + len(gl.decimal) + len(gl.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, gl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, gl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, gl.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, gl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, gl.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, gl.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'gl'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl) FmtDateShort(t time.Time) []byte {

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

// FmtDateMedium returns the medium date representation of 't' for 'gl'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, gl.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'gl'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, gl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'gl'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, gl.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, gl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'gl'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'gl'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'gl'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'gl'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (gl *gl) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, gl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := gl.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
