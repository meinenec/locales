package lt

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type lt struct {
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

// New returns a new instance of translator for the 'lt' locale
func New() locales.Translator {
	return &lt{
		locale:                 "lt",
		pluralsCardinal:        []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{5, 6, 2, 4},
		decimal:                []byte{0x2c},
		group:                  []byte{0xc2, 0xa0},
		minus:                  []byte{0xe2, 0x88, 0x92},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46, 0x20}, {0x47, 0x42, 0x50}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0x59}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x58, 0x41, 0x46}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x58, 0x4f, 0x46}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		percentSuffix:          []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x73, 0x61, 0x75, 0x73, 0x2e}, {0x76, 0x61, 0x73, 0x2e}, {0x6b, 0x6f, 0x76, 0x2e}, {0x62, 0x61, 0x6c, 0x2e}, {0x67, 0x65, 0x67, 0x2e}, {0x62, 0x69, 0x72, 0xc5, 0xbe, 0x2e}, {0x6c, 0x69, 0x65, 0x70, 0x2e}, {0x72, 0x75, 0x67, 0x70, 0x2e}, {0x72, 0x75, 0x67, 0x73, 0x2e}, {0x73, 0x70, 0x61, 0x6c, 0x2e}, {0x6c, 0x61, 0x70, 0x6b, 0x72, 0x2e}, {0x67, 0x72, 0x75, 0x6f, 0x64, 0x2e}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x53}, {0x56}, {0x4b}, {0x42}, {0x47}, {0x42}, {0x4c}, {0x52}, {0x52}, {0x53}, {0x4c}, {0x47}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x73, 0x61, 0x75, 0x73, 0x69, 0x6f}, {0x76, 0x61, 0x73, 0x61, 0x72, 0x69, 0x6f}, {0x6b, 0x6f, 0x76, 0x6f}, {0x62, 0x61, 0x6c, 0x61, 0x6e, 0x64, 0xc5, 0xbe, 0x69, 0x6f}, {0x67, 0x65, 0x67, 0x75, 0xc5, 0xbe, 0xc4, 0x97, 0x73}, {0x62, 0x69, 0x72, 0xc5, 0xbe, 0x65, 0x6c, 0x69, 0x6f}, {0x6c, 0x69, 0x65, 0x70, 0x6f, 0x73}, {0x72, 0x75, 0x67, 0x70, 0x6a, 0xc5, 0xab, 0xc4, 0x8d, 0x69, 0x6f}, {0x72, 0x75, 0x67, 0x73, 0xc4, 0x97, 0x6a, 0x6f}, {0x73, 0x70, 0x61, 0x6c, 0x69, 0x6f}, {0x6c, 0x61, 0x70, 0x6b, 0x72, 0x69, 0xc4, 0x8d, 0x69, 0x6f}, {0x67, 0x72, 0x75, 0x6f, 0x64, 0xc5, 0xbe, 0x69, 0x6f}},
		daysAbbreviated:        [][]uint8{{0x73, 0x6b}, {0x70, 0x72}, {0x61, 0x6e}, {0x74, 0x72}, {0x6b, 0x74}, {0x70, 0x6e}, {0xc5, 0xa1, 0x74}},
		daysNarrow:             [][]uint8{{0x53}, {0x50}, {0x41}, {0x54}, {0x4b}, {0x50}, {0xc5, 0xa0}},
		daysShort:              [][]uint8{{0x53, 0x6b}, {0x50, 0x72}, {0x41, 0x6e}, {0x54, 0x72}, {0x4b, 0x74}, {0x50, 0x6e}, {0xc5, 0xa0, 0x74}},
		daysWide:               [][]uint8{{0x73, 0x65, 0x6b, 0x6d, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x69, 0x73}, {0x70, 0x69, 0x72, 0x6d, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x69, 0x73}, {0x61, 0x6e, 0x74, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x69, 0x73}, {0x74, 0x72, 0x65, 0xc4, 0x8d, 0x69, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x69, 0x73}, {0x6b, 0x65, 0x74, 0x76, 0x69, 0x72, 0x74, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x69, 0x73}, {0x70, 0x65, 0x6e, 0x6b, 0x74, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x69, 0x73}, {0xc5, 0xa1, 0x65, 0xc5, 0xa1, 0x74, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x69, 0x73}},
		periodsAbbreviated:     [][]uint8{{0x70, 0x72, 0x69, 0x65, 0xc5, 0xa1, 0x70, 0x69, 0x65, 0x74}, {0x70, 0x6f, 0x70, 0x69, 0x65, 0x74}},
		periodsNarrow:          [][]uint8{{0x70, 0x72, 0x2e, 0x20, 0x70, 0x2e}, {0x70, 0x6f, 0x70, 0x2e}},
		periodsWide:            [][]uint8{{0x70, 0x72, 0x69, 0x65, 0xc5, 0xa1, 0x70, 0x69, 0x65, 0x74}, {0x70, 0x6f, 0x70, 0x69, 0x65, 0x74}},
		erasAbbreviated:        [][]uint8{{0x70, 0x72, 0x2e, 0x20, 0x4b, 0x72, 0x2e}, {0x70, 0x6f, 0x20, 0x4b, 0x72, 0x2e}},
		erasNarrow:             [][]uint8{{0x70, 0x72, 0x2e, 0x20, 0x4b, 0x72, 0x2e}, {0x70, 0x6f, 0x20, 0x4b, 0x72, 0x2e}},
		erasWide:               [][]uint8{{0x70, 0x72, 0x69, 0x65, 0xc5, 0xa1, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0xc5, 0xb3}, {0x70, 0x6f, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x61, 0x75, 0x73}},
		timezones:              map[string][]uint8{"WIB": {0x56, 0x61, 0x6b, 0x61, 0x72, 0xc5, 0xb3, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "AEST": {0x52, 0x79, 0x74, 0xc5, 0xb3, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "OESZ": {0x52, 0x79, 0x74, 0xc5, 0xb3, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "WESZ": {0x56, 0x61, 0x6b, 0x61, 0x72, 0xc5, 0xb3, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "MST": {0xc5, 0xa0, 0x69, 0x61, 0x75, 0x72, 0xc4, 0x97, 0x73, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x6f, 0x73, 0x20, 0x6b, 0x61, 0x6c, 0x6e, 0xc5, 0xb3, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "ART": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x6f, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "UYT": {0x55, 0x72, 0x75, 0x67, 0x76, 0x61, 0x6a, 0x61, 0x75, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "UYST": {0x55, 0x72, 0x75, 0x67, 0x76, 0x61, 0x6a, 0x61, 0x75, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "MYT": {0x4d, 0x61, 0x6c, 0x61, 0x69, 0x7a, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "WAT": {0x56, 0x61, 0x6b, 0x61, 0x72, 0xc5, 0xb3, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x6f, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "TMT": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0xc4, 0x97, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x6f, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "CDT": {0xc5, 0xa0, 0x69, 0x61, 0x75, 0x72, 0xc4, 0x97, 0x73, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x6f, 0x73, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "MESZ": {0x56, 0x69, 0x64, 0x75, 0x72, 0x69, 0x6f, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "ChST": {0xc4, 0x8c, 0x61, 0x6d, 0x6f, 0x72, 0x6f, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "LHDT": {0x4c, 0x6f, 0x72, 0x64, 0x6f, 0x20, 0x48, 0x61, 0x75, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "HKST": {0x48, 0x6f, 0x6e, 0x6b, 0x6f, 0x6e, 0x67, 0x6f, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "∅∅∅": {0x42, 0x72, 0x61, 0x7a, 0x69, 0x6c, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "SAST": {0x50, 0x69, 0x65, 0x74, 0xc5, 0xb3, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "GMT": {0x47, 0x72, 0x69, 0x6e, 0x76, 0x69, 0xc4, 0x8d, 0x6f, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "CHADT": {0xc4, 0x8c, 0x61, 0x74, 0x61, 0x6d, 0x6f, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "OEZ": {0x52, 0x79, 0x74, 0xc5, 0xb3, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "SRT": {0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x6f, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "GYT": {0x47, 0x61, 0x6a, 0x61, 0x6e, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "AEDT": {0x52, 0x79, 0x74, 0xc5, 0xb3, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "CLT": {0xc4, 0x8c, 0x69, 0x6c, 0xc4, 0x97, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "LHST": {0x4c, 0x6f, 0x72, 0x64, 0x6f, 0x20, 0x48, 0x61, 0x75, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "MDT": {0xc5, 0xa0, 0x69, 0x61, 0x75, 0x72, 0xc4, 0x97, 0x73, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x6f, 0x73, 0x20, 0x6b, 0x61, 0x6c, 0x6e, 0xc5, 0xb3, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "AST": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "CLST": {0xc4, 0x8c, 0x69, 0x6c, 0xc4, 0x97, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "EST": {0xc5, 0xa0, 0x69, 0x61, 0x75, 0x72, 0xc4, 0x97, 0x73, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x6f, 0x73, 0x20, 0x72, 0x79, 0x74, 0xc5, 0xb3, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "HNT": {0x4e, 0x69, 0x75, 0x66, 0x61, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x6e, 0x64, 0x6f, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "WARST": {0x56, 0x61, 0x6b, 0x61, 0x72, 0xc5, 0xb3, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x6f, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "HADT": {0x48, 0x61, 0x76, 0x61, 0x6a, 0xc5, 0xb3, 0xe2, 0x80, 0x93, 0x41, 0x6c, 0x65, 0x75, 0x74, 0xc5, 0xb3, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "BOT": {0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "ACWDT": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x6e, 0xc4, 0x97, 0x73, 0x20, 0x76, 0x61, 0x6b, 0x61, 0x72, 0xc5, 0xb3, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "PDT": {0xc5, 0xa0, 0x69, 0x61, 0x75, 0x72, 0xc4, 0x97, 0x73, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x6f, 0x73, 0x20, 0x52, 0x61, 0x6d, 0x69, 0x6f, 0x6a, 0x6f, 0x20, 0x76, 0x61, 0x6e, 0x64, 0x65, 0x6e, 0x79, 0x6e, 0x6f, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "WIT": {0x52, 0x79, 0x74, 0xc5, 0xb3, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "HAT": {0x4e, 0x69, 0x75, 0x66, 0x61, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x6e, 0x64, 0x6f, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "ACDT": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x6e, 0xc4, 0x97, 0x73, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "ADT": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "WART": {0x56, 0x61, 0x6b, 0x61, 0x72, 0xc5, 0xb3, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x6f, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "WAST": {0x56, 0x61, 0x6b, 0x61, 0x72, 0xc5, 0xb3, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x6f, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "PST": {0xc5, 0xa0, 0x69, 0x61, 0x75, 0x72, 0xc4, 0x97, 0x73, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x6f, 0x73, 0x20, 0x52, 0x61, 0x6d, 0x69, 0x6f, 0x6a, 0x6f, 0x20, 0x76, 0x61, 0x6e, 0x64, 0x65, 0x6e, 0x79, 0x6e, 0x6f, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "JDT": {0x4a, 0x61, 0x70, 0x6f, 0x6e, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "NZDT": {0x4e, 0x61, 0x75, 0x6a, 0x6f, 0x73, 0x69, 0x6f, 0x73, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "HKT": {0x48, 0x6f, 0x6e, 0x6b, 0x6f, 0x6e, 0x67, 0x6f, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "WEZ": {0x56, 0x61, 0x6b, 0x61, 0x72, 0xc5, 0xb3, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "AWST": {0x56, 0x61, 0x6b, 0x61, 0x72, 0xc5, 0xb3, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "BT": {0x42, 0x75, 0x74, 0x61, 0x6e, 0x6f, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "CHAST": {0xc4, 0x8c, 0x61, 0x74, 0x61, 0x6d, 0x6f, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "WITA": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x6e, 0xc4, 0x97, 0x73, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "MEZ": {0x56, 0x69, 0x64, 0x75, 0x72, 0x69, 0x6f, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "JST": {0x4a, 0x61, 0x70, 0x6f, 0x6e, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "GFT": {0x50, 0x72, 0x61, 0x6e, 0x63, 0xc5, 0xab, 0x7a, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0x47, 0x76, 0x69, 0x61, 0x6e, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "CAT": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x6e, 0xc4, 0x97, 0x73, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "CST": {0xc5, 0xa0, 0x69, 0x61, 0x75, 0x72, 0xc4, 0x97, 0x73, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x6f, 0x73, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "AKDT": {0x41, 0x6c, 0x69, 0x61, 0x73, 0x6b, 0x6f, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "ACWST": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x6e, 0xc4, 0x97, 0x73, 0x20, 0x76, 0x61, 0x6b, 0x61, 0x72, 0xc5, 0xb3, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "AWDT": {0x56, 0x61, 0x6b, 0x61, 0x72, 0xc5, 0xb3, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "COST": {0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "HAST": {0x48, 0x61, 0x76, 0x61, 0x6a, 0xc5, 0xb3, 0xe2, 0x80, 0x93, 0x41, 0x6c, 0x65, 0x75, 0x74, 0xc5, 0xb3, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "SGT": {0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0xc5, 0xab, 0x72, 0x6f, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "AKST": {0x41, 0x6c, 0x69, 0x61, 0x73, 0x6b, 0x6f, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "TMST": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0xc4, 0x97, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x6f, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "EDT": {0xc5, 0xa0, 0x69, 0x61, 0x75, 0x72, 0xc4, 0x97, 0x73, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x6f, 0x73, 0x20, 0x72, 0x79, 0x74, 0xc5, 0xb3, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "VET": {0x56, 0x65, 0x6e, 0x65, 0x73, 0x75, 0x65, 0x6c, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "ECT": {0x45, 0x6b, 0x76, 0x61, 0x64, 0x6f, 0x72, 0x6f, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "ACST": {0x43, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x6e, 0xc4, 0x97, 0x73, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "COT": {0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "EAT": {0x52, 0x79, 0x74, 0xc5, 0xb3, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "ARST": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x6f, 0x73, 0x20, 0x76, 0x61, 0x73, 0x61, 0x72, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "NZST": {0x4e, 0x61, 0x75, 0x6a, 0x6f, 0x73, 0x69, 0x6f, 0x73, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0xc5, 0xbe, 0x69, 0x65, 0x6d, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}, "IST": {0x49, 0x6e, 0x64, 0x69, 0x6a, 0x6f, 0x73, 0x20, 0x6c, 0x61, 0x69, 0x6b, 0x61, 0x73}},
	}
}

// Locale returns the current translators string locale
func (lt *lt) Locale() string {
	return lt.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'lt'
func (lt *lt) PluralsCardinal() []locales.PluralRule {
	return lt.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'lt'
func (lt *lt) PluralsOrdinal() []locales.PluralRule {
	return lt.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'lt'
func (lt *lt) PluralsRange() []locales.PluralRule {
	return lt.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'lt'
func (lt *lt) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	f := locales.F(n, v)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)

	if nMod10 == 1 && nMod100 < 11 && nMod100 > 19 {
		return locales.PluralRuleOne
	} else if nMod10 >= 2 && nMod10 <= 9 && nMod100 < 11 && nMod100 > 19 {
		return locales.PluralRuleFew
	} else if f != 0 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'lt'
func (lt *lt) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'lt'
func (lt *lt) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := lt.CardinalPluralRule(num1, v1)
	end := lt.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (lt *lt) MonthAbbreviated(month time.Month) []byte {
	return lt.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (lt *lt) MonthsAbbreviated() [][]byte {
	return lt.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (lt *lt) MonthNarrow(month time.Month) []byte {
	return lt.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (lt *lt) MonthsNarrow() [][]byte {
	return lt.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (lt *lt) MonthWide(month time.Month) []byte {
	return lt.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (lt *lt) MonthsWide() [][]byte {
	return lt.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (lt *lt) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return lt.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (lt *lt) WeekdaysAbbreviated() [][]byte {
	return lt.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (lt *lt) WeekdayNarrow(weekday time.Weekday) []byte {
	return lt.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (lt *lt) WeekdaysNarrow() [][]byte {
	return lt.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (lt *lt) WeekdayShort(weekday time.Weekday) []byte {
	return lt.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (lt *lt) WeekdaysShort() [][]byte {
	return lt.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (lt *lt) WeekdayWide(weekday time.Weekday) []byte {
	return lt.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (lt *lt) WeekdaysWide() [][]byte {
	return lt.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'lt' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lt *lt) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(lt.decimal) + len(lt.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lt.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(lt.group) - 1; j >= 0; j-- {
					b = append(b, lt.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(lt.minus) - 1; j >= 0; j-- {
			b = append(b, lt.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'lt' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (lt *lt) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(lt.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lt.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(lt.minus) - 1; j >= 0; j-- {
			b = append(b, lt.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, lt.percentSuffix...)

	b = append(b, lt.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'lt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lt *lt) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lt.currencies[currency]
	l := len(s) + len(lt.decimal) + len(lt.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lt.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(lt.group) - 1; j >= 0; j-- {
					b = append(b, lt.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(lt.minus) - 1; j >= 0; j-- {
			b = append(b, lt.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, lt.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, lt.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'lt'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lt *lt) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lt.currencies[currency]
	l := len(s) + len(lt.decimal) + len(lt.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lt.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(lt.group) - 1; j >= 0; j-- {
					b = append(b, lt.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(lt.minus) - 1; j >= 0; j-- {
			b = append(b, lt.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, lt.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, lt.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, lt.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'lt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lt *lt) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'lt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lt *lt) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'lt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lt *lt) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x6d, 0x27, 0x2e, 0x20}...)
	b = append(b, lt.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x27, 0x2e}...)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'lt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lt *lt) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x6d, 0x27, 0x2e, 0x20}...)
	b = append(b, lt.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x27, 0x2e, 0x2c, 0x20}...)
	b = append(b, lt.daysWide[t.Weekday()]...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'lt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lt *lt) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'lt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lt *lt) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'lt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lt *lt) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'lt'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lt *lt) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lt.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lt.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := lt.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
