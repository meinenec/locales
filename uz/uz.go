package uz

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type uz struct {
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

// New returns a new instance of translator for the 'uz' locale
func New() locales.Translator {
	return &uz{
		locale:                 "uz",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                []byte{0xd9, 0xab},
		group:                  []byte{0xd9, 0xac},
		minus:                  []byte{0x2d},
		percent:                []byte{0xd9, 0xaa},
		perMille:               []byte{0xd8, 0x89},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x24}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x52, 0x24}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x24}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0xc2, 0xa5}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46, 0x20}, {0xc2, 0xa3}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x24}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0xe2, 0x82, 0xaa}, {0xe2, 0x82, 0xb9}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0xc2, 0xa5}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0xe2, 0x82, 0xa9}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x24}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x24}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x4e, 0x54, 0x24}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58}, {0x55, 0x53, 0x24}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55}, {0x73, 0x6f, 0xca, 0xbb, 0x6d}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46}, {0xe2, 0x82, 0xab}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x45, 0x43, 0x24}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44, 0x20}, {0x43, 0x46, 0x50, 0x46}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x79, 0x61, 0x6e}, {0x66, 0x65, 0x76}, {0x6d, 0x61, 0x72}, {0x61, 0x70, 0x72}, {0x6d, 0x61, 0x79}, {0x69, 0x79, 0x6e}, {0x69, 0x79, 0x6c}, {0x61, 0x76, 0x67}, {0x73, 0x65, 0x6e}, {0x6f, 0x6b, 0x74}, {0x6e, 0x6f, 0x79}, {0x64, 0x65, 0x6b}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x59}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x49}, {0x49}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x79, 0x61, 0x6e, 0x76, 0x61, 0x72}, {0x66, 0x65, 0x76, 0x72, 0x61, 0x6c}, {0x6d, 0x61, 0x72, 0x74}, {0x61, 0x70, 0x72, 0x65, 0x6c}, {0x6d, 0x61, 0x79}, {0x69, 0x79, 0x75, 0x6e}, {0x69, 0x79, 0x75, 0x6c}, {0x61, 0x76, 0x67, 0x75, 0x73, 0x74}, {0x53, 0x65, 0x6e, 0x74, 0x61, 0x62, 0x72}, {0x4f, 0x6b, 0x74, 0x61, 0x62, 0x72}, {0x6e, 0x6f, 0x79, 0x61, 0x62, 0x72}, {0x64, 0x65, 0x6b, 0x61, 0x62, 0x72}},
		daysAbbreviated:        [][]uint8{{0x59, 0x61}, {0x44, 0x75}, {0x53, 0x65}, {0x43, 0x68}, {0x50, 0x61}, {0x4a, 0x75}, {0x53, 0x68}},
		daysNarrow:             [][]uint8{{0x59}, {0x44}, {0x53}, {0x43}, {0x50}, {0x4a}, {0x53}},
		daysShort:              [][]uint8{{0x59, 0x61}, {0x44, 0x75}, {0x53, 0x65}, {0x43, 0x68}, {0x50, 0x61}, {0x4a, 0x75}, {0x53, 0x68}},
		daysWide:               [][]uint8{{0x79, 0x61, 0x6b, 0x73, 0x68, 0x61, 0x6e, 0x62, 0x61}, {0x64, 0x75, 0x73, 0x68, 0x61, 0x6e, 0x62, 0x61}, {0x73, 0x65, 0x73, 0x68, 0x61, 0x6e, 0x62, 0x61}, {0x63, 0x68, 0x6f, 0x72, 0x73, 0x68, 0x61, 0x6e, 0x62, 0x61}, {0x70, 0x61, 0x79, 0x73, 0x68, 0x61, 0x6e, 0x62, 0x61}, {0x6a, 0x75, 0x6d, 0x61}, {0x73, 0x68, 0x61, 0x6e, 0x62, 0x61}},
		periodsAbbreviated:     [][]uint8{{0x54, 0x4f}, {0x54, 0x4b}},
		periodsNarrow:          [][]uint8{{0x54, 0x4f}, {0x54, 0x4b}},
		periodsWide:            [][]uint8{{0x54, 0x4f}, {0x54, 0x4b}},
		erasAbbreviated:        [][]uint8{[]uint8(nil), []uint8(nil)},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{[]uint8(nil), []uint8(nil)},
		timezones:              map[string][]uint8{"MESZ": {0x4d, 0x61, 0x72, 0x6b, 0x61, 0x7a, 0x69, 0x79, 0x20, 0x59, 0x65, 0x76, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "AKDT": {0x41, 0x6c, 0x79, 0x61, 0x73, 0x6b, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "WART": {0x47, 0xca, 0xbb, 0x61, 0x72, 0x62, 0x69, 0x79, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "GYT": {0x47, 0x61, 0x79, 0x61, 0x6e, 0x61, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "HKT": {0x47, 0x6f, 0x6e, 0x6b, 0x6f, 0x6e, 0x67, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "ACDT": {0x4d, 0x61, 0x72, 0x6b, 0x61, 0x7a, 0x69, 0x79, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "BOT": {0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x79, 0x61, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "OESZ": {0x53, 0x68, 0x61, 0x72, 0x71, 0x69, 0x79, 0x20, 0x59, 0x65, 0x76, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "EDT": {0x53, 0x68, 0x61, 0x72, 0x71, 0x69, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "LHST": {0x4c, 0x6f, 0x72, 0x64, 0x2d, 0x58, 0x61, 0x75, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "AWDT": {0x47, 0xe2, 0x80, 0x98, 0x61, 0x72, 0x62, 0x69, 0x79, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "MST": {0x53, 0x68, 0x69, 0x6d, 0x6f, 0x6c, 0x69, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x74, 0x6f, 0x67, 0xca, 0xbb, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "SGT": {0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "ACWDT": {0x4d, 0x61, 0x72, 0x6b, 0x61, 0x7a, 0x69, 0x79, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x67, 0xe2, 0x80, 0x98, 0x61, 0x72, 0x62, 0x69, 0x79, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "IST": {0x48, 0x69, 0x6e, 0x64, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "AEST": {0x53, 0x68, 0x61, 0x72, 0x71, 0x69, 0x79, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "CHADT": {0x43, 0x68, 0x61, 0x74, 0x65, 0x6d, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "TMT": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0x61, 0x6e, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "HAST": {0x47, 0x61, 0x76, 0x61, 0x79, 0x69, 0x2d, 0x61, 0x6c, 0x65, 0x75, 0x74, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "JST": {0x59, 0x61, 0x70, 0x6f, 0x6e, 0x69, 0x79, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "WIT": {0x53, 0x68, 0x61, 0x72, 0x71, 0x69, 0x79, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x79, 0x61, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "WESZ": {0x47, 0xe2, 0x80, 0x98, 0x61, 0x72, 0x62, 0x69, 0x79, 0x20, 0x59, 0x65, 0x76, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "ACWST": {0x4d, 0x61, 0x72, 0x6b, 0x61, 0x7a, 0x69, 0x79, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x67, 0xe2, 0x80, 0x98, 0x61, 0x72, 0x62, 0x69, 0x79, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "MYT": {0x4d, 0x61, 0x6c, 0x61, 0x79, 0x7a, 0x69, 0x79, 0x61, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "GFT": {0x46, 0x72, 0x61, 0x6e, 0x73, 0x75, 0x7a, 0x20, 0x47, 0x76, 0x69, 0x61, 0x6e, 0x61, 0x73, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "NZST": {0x59, 0x61, 0x6e, 0x67, 0x69, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x79, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "ARST": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "HNT": {0x4e, 0x79, 0x75, 0x66, 0x61, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x6e, 0x64, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "MDT": {0x53, 0x68, 0x69, 0x6d, 0x6f, 0x6c, 0x69, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x74, 0x6f, 0x67, 0xca, 0xbb, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "WAST": {0x47, 0xca, 0xbb, 0x61, 0x72, 0x62, 0x69, 0x79, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "AWST": {0x47, 0xe2, 0x80, 0x98, 0x61, 0x72, 0x62, 0x69, 0x79, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "ECT": {0x45, 0x6b, 0x76, 0x61, 0x64, 0x6f, 0x72, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "EAT": {0x53, 0x68, 0x61, 0x72, 0x71, 0x69, 0x79, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "CDT": {0x4d, 0x61, 0x72, 0x6b, 0x61, 0x7a, 0x69, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "AST": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "EST": {0x53, 0x68, 0x61, 0x72, 0x71, 0x69, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "WARST": {0x47, 0xca, 0xbb, 0x61, 0x72, 0x62, 0x69, 0x79, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "WIB": {0x47, 0xca, 0xbb, 0x61, 0x72, 0x62, 0x69, 0x79, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x79, 0x61, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "GMT": {0x47, 0x72, 0x69, 0x6e, 0x76, 0x69, 0x63, 0x68, 0x20, 0x6f, 0xe2, 0x80, 0x98, 0x72, 0x74, 0x61, 0x63, 0x68, 0x61, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "CLT": {0x43, 0x68, 0x69, 0x6c, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "NZDT": {0x59, 0x61, 0x6e, 0x67, 0x69, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x79, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "LHDT": {0x4c, 0x6f, 0x72, 0x64, 0x2d, 0x58, 0x61, 0x75, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "VET": {0x56, 0x65, 0x6e, 0x65, 0x73, 0x75, 0x65, 0x6c, 0x61, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "SRT": {0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "ACST": {0x4d, 0x61, 0x72, 0x6b, 0x61, 0x7a, 0x69, 0x79, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "HADT": {0x47, 0x61, 0x76, 0x61, 0x79, 0x69, 0x2d, 0x61, 0x6c, 0x65, 0x75, 0x74, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "HAT": {0x4e, 0x79, 0x75, 0x66, 0x61, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x6e, 0x64, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "WEZ": {0x47, 0xe2, 0x80, 0x98, 0x61, 0x72, 0x62, 0x69, 0x79, 0x20, 0x59, 0x65, 0x76, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "PDT": {0x53, 0x68, 0x69, 0x6d, 0x6f, 0x6c, 0x69, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x54, 0x69, 0x6e, 0x63, 0x68, 0x20, 0x6f, 0x6b, 0x65, 0x61, 0x6e, 0x69, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "ART": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "ADT": {0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6b, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "AKST": {0x41, 0x6c, 0x79, 0x61, 0x73, 0x6b, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "HKST": {0x47, 0x6f, 0x6e, 0x6b, 0x6f, 0x6e, 0x67, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "BT": {0x42, 0x75, 0x74, 0x61, 0x6e, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "COT": {0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x79, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "TMST": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0x61, 0x6e, 0x69, 0x73, 0x74, 0x6f, 0x6e, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "CLST": {0x43, 0x68, 0x69, 0x6c, 0x69, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "JDT": {0x59, 0x61, 0x70, 0x6f, 0x6e, 0x69, 0x79, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "UYT": {0x55, 0x72, 0x75, 0x67, 0x76, 0x61, 0x79, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "UYST": {0x55, 0x72, 0x75, 0x67, 0x76, 0x61, 0x79, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "CAT": {0x4d, 0x61, 0x72, 0x6b, 0x61, 0x7a, 0x69, 0x79, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "WITA": {0x4d, 0x61, 0x72, 0x6b, 0x61, 0x7a, 0x69, 0x79, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x79, 0x61, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "CST": {0x4d, 0x61, 0x72, 0x6b, 0x61, 0x7a, 0x69, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "MEZ": {0x4d, 0x61, 0x72, 0x6b, 0x61, 0x7a, 0x69, 0x79, 0x20, 0x59, 0x65, 0x76, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "OEZ": {0x53, 0x68, 0x61, 0x72, 0x71, 0x69, 0x79, 0x20, 0x59, 0x65, 0x76, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "WAT": {0x47, 0xca, 0xbb, 0x61, 0x72, 0x62, 0x69, 0x79, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "SAST": {0x4a, 0x61, 0x6e, 0x75, 0x62, 0x69, 0x79, 0x20, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "ChST": {0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "AEDT": {0x53, 0x68, 0x61, 0x72, 0x71, 0x69, 0x79, 0x20, 0x41, 0x76, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x79, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "CHAST": {0x43, 0x68, 0x61, 0x74, 0x65, 0x6d, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "∅∅∅": {0x50, 0x65, 0x72, 0x75, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "PST": {0x53, 0x68, 0x69, 0x6d, 0x6f, 0x6c, 0x69, 0x79, 0x20, 0x41, 0x6d, 0x65, 0x72, 0x69, 0x6b, 0x61, 0x20, 0x54, 0x69, 0x6e, 0x63, 0x68, 0x20, 0x6f, 0x6b, 0x65, 0x61, 0x6e, 0x69, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x74, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}, "COST": {0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x79, 0x61, 0x20, 0x79, 0x6f, 0x7a, 0x67, 0x69, 0x20, 0x76, 0x61, 0x71, 0x74, 0x69}},
	}
}

// Locale returns the current translators string locale
func (uz *uz) Locale() string {
	return uz.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'uz'
func (uz *uz) PluralsCardinal() []locales.PluralRule {
	return uz.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'uz'
func (uz *uz) PluralsOrdinal() []locales.PluralRule {
	return uz.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'uz'
func (uz *uz) PluralsRange() []locales.PluralRule {
	return uz.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'uz'
func (uz *uz) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'uz'
func (uz *uz) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'uz'
func (uz *uz) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := uz.CardinalPluralRule(num1, v1)
	end := uz.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (uz *uz) MonthAbbreviated(month time.Month) []byte {
	return uz.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (uz *uz) MonthsAbbreviated() [][]byte {
	return uz.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (uz *uz) MonthNarrow(month time.Month) []byte {
	return uz.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (uz *uz) MonthsNarrow() [][]byte {
	return uz.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (uz *uz) MonthWide(month time.Month) []byte {
	return uz.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (uz *uz) MonthsWide() [][]byte {
	return uz.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (uz *uz) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return uz.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (uz *uz) WeekdaysAbbreviated() [][]byte {
	return uz.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (uz *uz) WeekdayNarrow(weekday time.Weekday) []byte {
	return uz.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (uz *uz) WeekdaysNarrow() [][]byte {
	return uz.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (uz *uz) WeekdayShort(weekday time.Weekday) []byte {
	return uz.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (uz *uz) WeekdaysShort() [][]byte {
	return uz.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (uz *uz) WeekdayWide(weekday time.Weekday) []byte {
	return uz.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (uz *uz) WeekdaysWide() [][]byte {
	return uz.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'uz' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (uz *uz) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(uz.decimal) + len(uz.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(uz.decimal) - 1; j >= 0; j-- {
				b = append(b, uz.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uz.group) - 1; j >= 0; j-- {
					b = append(b, uz.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'uz' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (uz *uz) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(uz.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(uz.decimal) - 1; j >= 0; j-- {
				b = append(b, uz.decimal[j])
			}

			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, uz.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'uz'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (uz *uz) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uz.currencies[currency]
	l := len(s) + len(uz.decimal) + len(uz.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(uz.decimal) - 1; j >= 0; j-- {
				b = append(b, uz.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uz.group) - 1; j >= 0; j-- {
					b = append(b, uz.group[j])
				}

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

	for j := len(uz.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, uz.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, uz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uz.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'uz'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (uz *uz) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uz.currencies[currency]
	l := len(s) + len(uz.decimal) + len(uz.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(uz.decimal) - 1; j >= 0; j-- {
				b = append(b, uz.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uz.group) - 1; j >= 0; j-- {
					b = append(b, uz.group[j])
				}

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

		for j := len(uz.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, uz.currencyNegativePrefix[j])
		}

		b = append(b, uz.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(uz.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, uz.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uz.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'uz'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (uz *uz) FmtDateShort(t time.Time) []byte {

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

// FmtDateMedium returns the medium date representation of 't' for 'uz'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (uz *uz) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = append(b, uz.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'uz'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (uz *uz) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = append(b, uz.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'uz'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (uz *uz) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, uz.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uz.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'uz'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (uz *uz) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'uz'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (uz *uz) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'uz'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (uz *uz) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	b = append(b, []byte{0x29}...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'uz'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (uz *uz) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := uz.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return b
}
