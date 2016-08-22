package ca

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ca struct {
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

// New returns a new instance of translator for the 'ca' locale
func New() locales.Translator {
	return &ca{
		locale:                 "ca",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 3, 4, 6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x24}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58, 0x20}, {0xc2, 0xa5}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0xe2, 0x82, 0xa7}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0xc2, 0xa3}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x24}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52, 0x20}, {0xe2, 0x82, 0xaa}, {0xe2, 0x82, 0xb9}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0xc2, 0xa5}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0xe2, 0x82, 0xa9}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x24}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0xe0, 0xb8, 0xbf}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x4e, 0x54, 0x24}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x44}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0xe2, 0x82, 0xab}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44}, {0x43, 0x46, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0xc2, 0xa0, 0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x67, 0x65, 0x6e, 0x2e}, {0x66, 0x65, 0x62, 0x72, 0x2e}, {0x6d, 0x61, 0x72, 0xc3, 0xa7}, {0x61, 0x62, 0x72, 0x2e}, {0x6d, 0x61, 0x69, 0x67}, {0x6a, 0x75, 0x6e, 0x79}, {0x6a, 0x75, 0x6c, 0x2e}, {0x61, 0x67, 0x2e}, {0x73, 0x65, 0x74, 0x2e}, {0x6f, 0x63, 0x74, 0x2e}, {0x6e, 0x6f, 0x76, 0x2e}, {0x64, 0x65, 0x73, 0x2e}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x47, 0x4e}, {0x46, 0x42}, {0x4d, 0xc3, 0x87}, {0x41, 0x42}, {0x4d, 0x47}, {0x4a, 0x4e}, {0x4a, 0x4c}, {0x41, 0x47}, {0x53, 0x54}, {0x4f, 0x43}, {0x4e, 0x56}, {0x44, 0x53}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x64, 0x65, 0x20, 0x67, 0x65, 0x6e, 0x65, 0x72}, {0x64, 0x65, 0x20, 0x66, 0x65, 0x62, 0x72, 0x65, 0x72}, {0x64, 0x65, 0x20, 0x6d, 0x61, 0x72, 0xc3, 0xa7}, {0x64, 0xe2, 0x80, 0x99, 0x61, 0x62, 0x72, 0x69, 0x6c}, {0x64, 0x65, 0x20, 0x6d, 0x61, 0x69, 0x67}, {0x64, 0x65, 0x20, 0x6a, 0x75, 0x6e, 0x79}, {0x64, 0x65, 0x20, 0x6a, 0x75, 0x6c, 0x69, 0x6f, 0x6c}, {0x64, 0xe2, 0x80, 0x99, 0x61, 0x67, 0x6f, 0x73, 0x74}, {0x64, 0x65, 0x20, 0x73, 0x65, 0x74, 0x65, 0x6d, 0x62, 0x72, 0x65}, {0x64, 0xe2, 0x80, 0x99, 0x6f, 0x63, 0x74, 0x75, 0x62, 0x72, 0x65}, {0x64, 0x65, 0x20, 0x6e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x72, 0x65}, {0x64, 0x65, 0x20, 0x64, 0x65, 0x73, 0x65, 0x6d, 0x62, 0x72, 0x65}},
		daysAbbreviated:        [][]uint8{{0x64, 0x67, 0x2e}, {0x64, 0x6c, 0x2e}, {0x64, 0x74, 0x2e}, {0x64, 0x63, 0x2e}, {0x64, 0x6a, 0x2e}, {0x64, 0x76, 0x2e}, {0x64, 0x73, 0x2e}},
		daysNarrow:             [][]uint8{{0x64, 0x67}, {0x64, 0x6c}, {0x64, 0x74}, {0x64, 0x63}, {0x64, 0x6a}, {0x64, 0x76}, {0x64, 0x73}},
		daysShort:              [][]uint8{{0x64, 0x67, 0x2e}, {0x64, 0x6c, 0x2e}, {0x64, 0x74, 0x2e}, {0x64, 0x63, 0x2e}, {0x64, 0x6a, 0x2e}, {0x64, 0x76, 0x2e}, {0x64, 0x73, 0x2e}},
		daysWide:               [][]uint8{{0x64, 0x69, 0x75, 0x6d, 0x65, 0x6e, 0x67, 0x65}, {0x64, 0x69, 0x6c, 0x6c, 0x75, 0x6e, 0x73}, {0x64, 0x69, 0x6d, 0x61, 0x72, 0x74, 0x73}, {0x64, 0x69, 0x6d, 0x65, 0x63, 0x72, 0x65, 0x73}, {0x64, 0x69, 0x6a, 0x6f, 0x75, 0x73}, {0x64, 0x69, 0x76, 0x65, 0x6e, 0x64, 0x72, 0x65, 0x73}, {0x64, 0x69, 0x73, 0x73, 0x61, 0x62, 0x74, 0x65}},
		periodsAbbreviated:     [][]uint8{{0x61, 0x2e, 0x20, 0x6d, 0x2e}, {0x70, 0x2e, 0x20, 0x6d, 0x2e}},
		periodsNarrow:          [][]uint8{{0x61, 0x2e, 0x20, 0x6d, 0x2e}, {0x70, 0x2e, 0x20, 0x6d, 0x2e}},
		periodsWide:            [][]uint8{{0x61, 0x2e, 0x20, 0x6d, 0x2e}, {0x70, 0x2e, 0x20, 0x6d, 0x2e}},
		erasAbbreviated:        [][]uint8{{0x61, 0x43}, {0x64, 0x43}},
		erasNarrow:             [][]uint8{{0x61, 0x43}, {0x64, 0x43}},
		erasWide:               [][]uint8{{0x61, 0x62, 0x61, 0x6e, 0x73, 0x20, 0x64, 0x65, 0x20, 0x43, 0x72, 0x69, 0x73, 0x74}, {0x64, 0x65, 0x73, 0x70, 0x72, 0xc3, 0xa9, 0x73, 0x20, 0x64, 0x65, 0x20, 0x43, 0x72, 0x69, 0x73, 0x74}},
		timezones:              map[string][]uint8{"WIB": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x6f, 0x65, 0x73, 0x74, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa8, 0x73, 0x69, 0x61}, "HADT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x65, 0x73}, "EAT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0xc3, 0x80, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "ARST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "WEZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x4f, 0x65, 0x73, 0x74, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61}, "WAST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0xc3, 0x80, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "∅∅∅": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x42, 0x72, 0x61, 0x73, 0xc3, 0xad, 0x6c, 0x69, 0x61}, "WITA": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa8, 0x73, 0x69, 0x61}, "CHAST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "SAST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x73, 0x75, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0xc3, 0x80, 0x66, 0x72, 0x69, 0x63, 0x61}, "TMST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "SGT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72}, "UYST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69}, "ACWST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x2d, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "ACDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "ECT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x71, 0x75, 0x61, 0x64, 0x6f, 0x72}, "MYT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x4d, 0x61, 0x6c, 0xc3, 0xa0, 0x69, 0x73, 0x69, 0x61}, "ADT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x74, 0x6c, 0xc3, 0xa0, 0x6e, 0x74, 0x69, 0x63}, "OESZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x73, 0x74, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61}, "GMT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4d, 0x65, 0x72, 0x69, 0x64, 0x69, 0xc3, 0xa0, 0x20, 0x64, 0x65, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68}, "CLT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x58, 0x69, 0x6c, 0x65}, "EDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6d, 0xc3, 0xa8, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4e, 0x6f, 0x72, 0x64}, "PST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x50, 0x61, 0x63, 0xc3, 0xad, 0x66, 0x69, 0x63}, "HAST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x65, 0x73}, "MST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6d, 0x75, 0x6e, 0x74, 0x61, 0x6e, 0x79, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6d, 0xc3, 0xa8, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4e, 0x6f, 0x72, 0x64}, "TMT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "LHST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "HKST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "CDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6d, 0xc3, 0xa8, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4e, 0x6f, 0x72, 0x64}, "MESZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61}, "UYT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x69}, "WARST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x6f, 0x65, 0x73, 0x74, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "COST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x43, 0x6f, 0x6c, 0xc3, 0xb2, 0x6d, 0x62, 0x69, 0x61}, "WART": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x6f, 0x65, 0x73, 0x74, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "AWDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "JST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4a, 0x61, 0x70, 0xc3, 0xb3}, "OEZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x45, 0x73, 0x74, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61}, "ART": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "GFT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x6c, 0x61, 0x20, 0x47, 0x75, 0x61, 0x69, 0x61, 0x6e, 0x61, 0x20, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x61}, "LHDT": {0x48, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "AWST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "AST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x41, 0x74, 0x6c, 0xc3, 0xa0, 0x6e, 0x74, 0x69, 0x63}, "AKDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "NZDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x4e, 0x6f, 0x76, 0x61, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x61}, "ChST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "IST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0xc3, 0x8d, 0x6e, 0x64, 0x69, 0x61}, "BOT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x42, 0x6f, 0x6c, 0xc3, 0xad, 0x76, 0x69, 0x61}, "AEDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "PDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x50, 0x61, 0x63, 0xc3, 0xad, 0x66, 0x69, 0x63}, "EST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6d, 0xc3, 0xa8, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4e, 0x6f, 0x72, 0x64}, "SRT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d}, "ACST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "WIT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0xc3, 0xa8, 0x73, 0x69, 0x61}, "JDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4a, 0x61, 0x70, 0xc3, 0xb3}, "NZST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x4e, 0x6f, 0x76, 0x61, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x61}, "HAT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x54, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x6f, 0x76, 0x61}, "HKT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "CLST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x58, 0x69, 0x6c, 0x65}, "COT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x43, 0x6f, 0x6c, 0xc3, 0xb2, 0x6d, 0x62, 0x69, 0x61}, "CHADT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "ACWDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x2d, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "WAT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0xc3, 0x80, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "AEST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0xc3, 0xa0, 0x6c, 0x69, 0x61, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "CST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6d, 0xc3, 0xa8, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4e, 0x6f, 0x72, 0x64}, "MEZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61}, "VET": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x6e, 0x65, 0xc3, 0xa7, 0x75, 0x65, 0x6c, 0x61}, "CAT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0xc3, 0x80, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "MDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6d, 0x75, 0x6e, 0x74, 0x61, 0x6e, 0x79, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6d, 0xc3, 0xa8, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x4e, 0x6f, 0x72, 0x64}, "AKST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "HNT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa0, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x65, 0x20, 0x54, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x6f, 0x76, 0x61}, "WESZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x65, 0x73, 0x74, 0x69, 0x75, 0x20, 0x64, 0x65, 0x20, 0x6c, 0xe2, 0x80, 0x99, 0x4f, 0x65, 0x73, 0x74, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61}, "GYT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61}, "BT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x42, 0x68, 0x75, 0x74, 0x61, 0x6e}},
	}
}

// Locale returns the current translators string locale
func (ca *ca) Locale() string {
	return ca.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ca'
func (ca *ca) PluralsCardinal() []locales.PluralRule {
	return ca.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ca'
func (ca *ca) PluralsOrdinal() []locales.PluralRule {
	return ca.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ca'
func (ca *ca) PluralsRange() []locales.PluralRule {
	return ca.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ca'
func (ca *ca) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ca'
func (ca *ca) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 || n == 3 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n == 4 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ca'
func (ca *ca) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ca *ca) MonthAbbreviated(month time.Month) []byte {
	return ca.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ca *ca) MonthsAbbreviated() [][]byte {
	return ca.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ca *ca) MonthNarrow(month time.Month) []byte {
	return ca.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ca *ca) MonthsNarrow() [][]byte {
	return ca.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ca *ca) MonthWide(month time.Month) []byte {
	return ca.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ca *ca) MonthsWide() [][]byte {
	return ca.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ca *ca) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return ca.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ca *ca) WeekdaysAbbreviated() [][]byte {
	return ca.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ca *ca) WeekdayNarrow(weekday time.Weekday) []byte {
	return ca.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ca *ca) WeekdaysNarrow() [][]byte {
	return ca.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ca *ca) WeekdayShort(weekday time.Weekday) []byte {
	return ca.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ca *ca) WeekdaysShort() [][]byte {
	return ca.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ca *ca) WeekdayWide(weekday time.Weekday) []byte {
	return ca.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ca *ca) WeekdaysWide() [][]byte {
	return ca.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ca' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ca.decimal) + len(ca.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ca.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ca.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ca' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ca *ca) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ca.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ca.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ca.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ca'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ca.currencies[currency]
	l := len(s) + len(ca.decimal) + len(ca.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ca.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ca.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ca.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ca.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ca'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ca.currencies[currency]
	l := len(s) + len(ca.decimal) + len(ca.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ca.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(ca.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ca.currencyNegativePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ca.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ca.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ca.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'ca'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'ca'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ca.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'ca'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ca.monthsWide[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'ca'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, ca.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ca.monthsWide[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'ca'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'ca'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'ca'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'ca'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ca *ca) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ca.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
