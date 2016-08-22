package se_FI

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type se_FI struct {
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

// New returns a new instance of translator for the 'se_FI' locale
func New() locales.Translator {
	return &se_FI{
		locale:                 "se_FI",
		pluralsCardinal:        []locales.PluralRule{2, 3, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                []byte{0x2c},
		group:                  []byte{0xc2, 0xa0},
		minus:                  []byte{0xe2, 0x88, 0x92},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44, 0x20}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e, 0x20}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c, 0x20}, {0x41, 0x4d, 0x44, 0x20}, {0x41, 0x4e, 0x47, 0x20}, {0x41, 0x4f, 0x41, 0x20}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53, 0x20}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44, 0x20}, {0x41, 0x57, 0x47, 0x20}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e, 0x20}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d, 0x20}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44, 0x20}, {0x42, 0x44, 0x54, 0x20}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e, 0x20}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44, 0x20}, {0x42, 0x49, 0x46, 0x20}, {0x42, 0x4d, 0x44, 0x20}, {0x42, 0x4e, 0x44, 0x20}, {0x42, 0x4f, 0x42, 0x20}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x42, 0x52, 0x4c, 0x20}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44, 0x20}, {0x42, 0x54, 0x4e, 0x20}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50, 0x20}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52, 0x20}, {0x42, 0x5a, 0x44, 0x20}, {0x43, 0x41, 0x44, 0x20}, {0x43, 0x44, 0x46, 0x20}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46, 0x20}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50, 0x20}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0x59, 0x20}, {0x43, 0x4f, 0x50, 0x20}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43, 0x20}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43, 0x20}, {0x43, 0x55, 0x50, 0x20}, {0x43, 0x56, 0x45, 0x20}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b, 0x20}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46, 0x20}, {0x44, 0x4b, 0x4b, 0x20}, {0x44, 0x4f, 0x50, 0x20}, {0x44, 0x5a, 0x44, 0x20}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50, 0x20}, {0x45, 0x52, 0x4e, 0x20}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42, 0x20}, {0x45, 0x55, 0x52, 0x20}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44, 0x20}, {0x46, 0x4b, 0x50, 0x20}, {0x46, 0x52, 0x46, 0x20}, {0x47, 0x42, 0x50, 0x20}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c, 0x20}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53, 0x20}, {0x47, 0x49, 0x50, 0x20}, {0x47, 0x4d, 0x44, 0x20}, {0x47, 0x4e, 0x46, 0x20}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51, 0x20}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44, 0x20}, {0x48, 0x4b, 0x44, 0x20}, {0x48, 0x4e, 0x4c, 0x20}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b, 0x20}, {0x48, 0x54, 0x47, 0x20}, {0x48, 0x55, 0x46, 0x20}, {0x49, 0x44, 0x52, 0x20}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0x49, 0x4c, 0x53, 0x20}, {0x49, 0x4e, 0x52, 0x20}, {0x49, 0x51, 0x44, 0x20}, {0x49, 0x52, 0x52, 0x20}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b, 0x20}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44, 0x20}, {0x4a, 0x4f, 0x44, 0x20}, {0x4a, 0x50, 0x59, 0x20}, {0x4b, 0x45, 0x53, 0x20}, {0x4b, 0x47, 0x53, 0x20}, {0x4b, 0x48, 0x52, 0x20}, {0x4b, 0x4d, 0x46, 0x20}, {0x4b, 0x50, 0x57, 0x20}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0x4b, 0x52, 0x57, 0x20}, {0x4b, 0x57, 0x44, 0x20}, {0x4b, 0x59, 0x44, 0x20}, {0x4b, 0x5a, 0x54, 0x20}, {0x4c, 0x41, 0x4b, 0x20}, {0x4c, 0x42, 0x50, 0x20}, {0x4c, 0x4b, 0x52, 0x20}, {0x4c, 0x52, 0x44, 0x20}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c, 0x20}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c, 0x20}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44, 0x20}, {0x4d, 0x41, 0x44, 0x20}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c, 0x20}, {0x4d, 0x47, 0x41, 0x20}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44, 0x20}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b, 0x20}, {0x4d, 0x4e, 0x54, 0x20}, {0x4d, 0x4f, 0x50, 0x20}, {0x4d, 0x52, 0x4f, 0x20}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52, 0x20}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52, 0x20}, {0x4d, 0x57, 0x4b, 0x20}, {0x4d, 0x58, 0x4e, 0x20}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52, 0x20}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e, 0x20}, {0x4e, 0x41, 0x44, 0x20}, {0x4e, 0x47, 0x4e, 0x20}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f, 0x20}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b, 0x20}, {0x4e, 0x50, 0x52, 0x20}, {0x4e, 0x5a, 0x44, 0x20}, {0x4f, 0x4d, 0x52, 0x20}, {0x50, 0x41, 0x42, 0x20}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e, 0x20}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b, 0x20}, {0x50, 0x48, 0x50, 0x20}, {0x50, 0x4b, 0x52, 0x20}, {0x50, 0x4c, 0x4e, 0x20}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47, 0x20}, {0x51, 0x41, 0x52, 0x20}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e, 0x20}, {0x52, 0x53, 0x44, 0x20}, {0x52, 0x55, 0x42, 0x20}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46, 0x20}, {0x53, 0x41, 0x52, 0x20}, {0x53, 0x42, 0x44, 0x20}, {0x53, 0x43, 0x52, 0x20}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47, 0x20}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b, 0x20}, {0x53, 0x47, 0x44, 0x20}, {0x53, 0x48, 0x50, 0x20}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c, 0x20}, {0x53, 0x4f, 0x53, 0x20}, {0x53, 0x52, 0x44, 0x20}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50, 0x20}, {0x53, 0x54, 0x44, 0x20}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50, 0x20}, {0x53, 0x5a, 0x4c, 0x20}, {0x54, 0x48, 0x42, 0x20}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53, 0x20}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54, 0x20}, {0x54, 0x4e, 0x44, 0x20}, {0x54, 0x4f, 0x50, 0x20}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59, 0x20}, {0x54, 0x54, 0x44, 0x20}, {0x54, 0x57, 0x44, 0x20}, {0x54, 0x5a, 0x53, 0x20}, {0x55, 0x41, 0x48, 0x20}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58, 0x20}, {0x55, 0x53, 0x44, 0x20}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55, 0x20}, {0x55, 0x5a, 0x53, 0x20}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46, 0x20}, {0x56, 0x4e, 0x44, 0x20}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56, 0x20}, {0x57, 0x53, 0x54, 0x20}, {0x58, 0x41, 0x46, 0x20}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x58, 0x43, 0x44, 0x20}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x58, 0x4f, 0x46, 0x20}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46, 0x20}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52, 0x20}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52, 0x20}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57, 0x20}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		percentSuffix:          []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x6f, 0xc4, 0x91, 0xc4, 0x91, 0x61, 0x6a, 0x61, 0x67, 0x65}, {0x67, 0x75, 0x6f, 0x76, 0x76, 0x61}, {0x6e, 0x6a, 0x75, 0x6b, 0xc4, 0x8d, 0x61}, {0x63, 0x75, 0x6f, 0xc5, 0x8b, 0x6f}, {0x6d, 0x69, 0x65, 0x73, 0x73, 0x65}, {0x67, 0x65, 0x61, 0x73, 0x73, 0x65}, {0x73, 0x75, 0x6f, 0x69, 0x64, 0x6e, 0x65}, {0x62, 0x6f, 0x72, 0x67, 0x65}, {0xc4, 0x8d, 0x61, 0x6b, 0xc4, 0x8d, 0x61}, {0x67, 0x6f, 0x6c, 0x67, 0x67, 0x6f, 0x74}, {0x73, 0x6b, 0xc3, 0xa1, 0x62, 0x6d, 0x61}, {0x6a, 0x75, 0x6f, 0x76, 0x6c, 0x61}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x4f}, {0x47}, {0x4e}, {0x43}, {0x4d}, {0x47}, {0x53}, {0x42}, {0xc4, 0x8c}, {0x47}, {0x53}, {0x4a}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x6f, 0xc4, 0x91, 0xc4, 0x91, 0x61, 0x6a, 0x61, 0x67, 0x65, 0x6d, 0xc3, 0xa1, 0x6e, 0x6e, 0x75}, {0x67, 0x75, 0x6f, 0x76, 0x76, 0x61, 0x6d, 0xc3, 0xa1, 0x6e, 0x6e, 0x75}, {0x6e, 0x6a, 0x75, 0x6b, 0xc4, 0x8d, 0x61, 0x6d, 0xc3, 0xa1, 0x6e, 0x6e, 0x75}, {0x63, 0x75, 0x6f, 0xc5, 0x8b, 0x6f, 0x6d, 0xc3, 0xa1, 0x6e, 0x6e, 0x75}, {0x6d, 0x69, 0x65, 0x73, 0x73, 0x65, 0x6d, 0xc3, 0xa1, 0x6e, 0x6e, 0x75}, {0x67, 0x65, 0x61, 0x73, 0x73, 0x65, 0x6d, 0xc3, 0xa1, 0x6e, 0x6e, 0x75}, {0x73, 0x75, 0x6f, 0x69, 0x64, 0x6e, 0x65, 0x6d, 0xc3, 0xa1, 0x6e, 0x6e, 0x75}, {0x62, 0x6f, 0x72, 0x67, 0x65, 0x6d, 0xc3, 0xa1, 0x6e, 0x6e, 0x75}, {0xc4, 0x8d, 0x61, 0x6b, 0xc4, 0x8d, 0x61, 0x6d, 0xc3, 0xa1, 0x6e, 0x6e, 0x75}, {0x67, 0x6f, 0x6c, 0x67, 0x67, 0x6f, 0x74, 0x6d, 0xc3, 0xa1, 0x6e, 0x6e, 0x75}, {0x73, 0x6b, 0xc3, 0xa1, 0x62, 0x6d, 0x61, 0x6d, 0xc3, 0xa1, 0x6e, 0x6e, 0x75}, {0x6a, 0x75, 0x6f, 0x76, 0x6c, 0x61, 0x6d, 0xc3, 0xa1, 0x6e, 0x6e, 0x75}},
		daysAbbreviated:        [][]uint8{{0x73, 0x6f, 0x74, 0x6e}, {0x76, 0x75, 0x6f, 0x73}, {0x6d, 0x61, 0xc5, 0x8b}, {0x67, 0x61, 0x73, 0x6b}, {0x64, 0x75, 0x6f, 0x72}, {0x62, 0x65, 0x61, 0x72}, {0x6c, 0xc3, 0xa1, 0x76}},
		daysNarrow:             [][]uint8{{0x53}, {0x4d}, {0x44}, {0x47}, {0x44}, {0x42}, {0x4c}},
		daysWide:               [][]uint8{{0x61, 0x65, 0x6a, 0x6c, 0x65, 0x67, 0x65}, {0x6d, 0xc3, 0xa5, 0x61, 0x6e, 0x74, 0x61}, {0x64, 0xc3, 0xa4, 0x6a, 0x73, 0x74, 0x61}, {0x67, 0x61, 0x73, 0x6b, 0x65, 0x76, 0x61, 0x68, 0x6b, 0x6f, 0x65}, {0x64, 0xc3, 0xa5, 0x61, 0x72, 0x73, 0x74, 0x61}, {0x62, 0x65, 0x61, 0x72, 0x6a, 0x61, 0x64, 0x61, 0x68, 0x6b, 0x65}, {0x6c, 0x61, 0x61, 0x76, 0x61, 0x64, 0x61, 0x68, 0x6b, 0x65}},
		periodsAbbreviated:     [][]uint8{{0x69, 0x2e, 0x62, 0x2e}, {0x65, 0x2e, 0x62, 0x2e}},
		periodsWide:            [][]uint8{{0x69, 0xc4, 0x91, 0x69, 0x74, 0x62, 0x65, 0x61, 0x69, 0x76, 0x65, 0x74}, {0x65, 0x61, 0x68, 0x6b, 0x65, 0x74, 0x62, 0x65, 0x61, 0x69, 0x76, 0x65, 0x74}},
		erasAbbreviated:        [][]uint8{{0x6f, 0x2e, 0x4b, 0x72, 0x2e}, {0x6d, 0x2e, 0x4b, 0x72, 0x2e}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x6f, 0x76, 0x64, 0x61, 0x6c, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x74, 0x75, 0x73, 0x61}, {0x6d, 0x61, 0xc5, 0x8b, 0xc5, 0x8b, 0x65, 0x6c, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x74, 0x75, 0x73, 0x61}},
		timezones:              map[string][]uint8{"SRT": {0x53, 0x52, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "WAT": {0x57, 0x41, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "HAST": {0x48, 0x41, 0x53, 0x54}, "EAT": {0x45, 0x41, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "WIB": {0x57, 0x49, 0x42}, "CST": {0x43, 0x53, 0x54}, "ADT": {0x41, 0x44, 0x54}, "VET": {0x56, 0x45, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "JST": {0x4a, 0x53, 0x54}, "CAT": {0x43, 0x41, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "WESZ": {0x6f, 0x61, 0x72, 0x6a, 0x65, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x68, 0x70, 0xc3, 0xa1, 0x20, 0x67, 0x65, 0x61, 0x73, 0x73, 0x69, 0xc3, 0xa1, 0x69, 0x67, 0x69}, "UYT": {0x55, 0x59, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "EDT": {0x45, 0x44, 0x54}, "HAT": {0x48, 0x41, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "BT": {0x42, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "GFT": {0x47, 0x46, 0x54}, "EST": {0x45, 0x53, 0x54}, "WEZ": {0x6f, 0x61, 0x72, 0x6a, 0x65, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x68, 0x70, 0xc3, 0xa1, 0x20, 0x64, 0xc3, 0xa1, 0x62, 0xc3, 0xa1, 0x6c, 0x61, 0xc5, 0xa1, 0xc3, 0xa1, 0x69, 0x67, 0x69}, "IST": {0x49, 0x53, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "GMT": {0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68, 0x20, 0x67, 0x61, 0x73, 0x6b, 0x6b, 0x61, 0x20, 0xc3, 0xa1, 0x69, 0x67, 0x69}, "WIT": {0x57, 0x49, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "MESZ": {0x67, 0x61, 0x73, 0x6b, 0x61, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x68, 0x70, 0xc3, 0xa1, 0x20, 0x67, 0x65, 0x61, 0x73, 0x73, 0x69, 0xc3, 0xa1, 0x69, 0x67, 0x69}, "OESZ": {0x6e, 0x75, 0x6f, 0x72, 0x74, 0x69, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x68, 0x70, 0xc3, 0xa1, 0x20, 0x67, 0x65, 0x61, 0x73, 0x73, 0x69, 0xc3, 0xa1, 0x69, 0x67, 0x69}, "PST": {0x50, 0x53, 0x54}, "AST": {0x41, 0x53, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "SGT": {0x53, 0x47, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "GYT": {0x47, 0x59, 0x54}, "COT": {0x43, 0x4f, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "ECT": {0x45, 0x43, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "CDT": {0x43, 0x44, 0x54}, "MST": {0x4d, 0x53, 0x54}, "PDT": {0x50, 0x44, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "MEZ": {0x67, 0x61, 0x73, 0x6b, 0x61, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x68, 0x70, 0xc3, 0xa1, 0x20, 0x64, 0xc3, 0xa1, 0x62, 0xc3, 0xa1, 0x6c, 0x61, 0xc5, 0xa1, 0xc3, 0xa1, 0x69, 0x67, 0x69}, "OEZ": {0x6e, 0x75, 0x6f, 0x72, 0x74, 0x69, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x68, 0x70, 0xc3, 0xa1, 0x20, 0x64, 0xc3, 0xa1, 0x62, 0xc3, 0xa1, 0x6c, 0x61, 0xc5, 0xa1, 0xc3, 0xa1, 0x69, 0x67, 0x69}, "ART": {0x41, 0x52, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}},
	}
}

// Locale returns the current translators string locale
func (se *se_FI) Locale() string {
	return se.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'se_FI'
func (se *se_FI) PluralsCardinal() []locales.PluralRule {
	return se.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'se_FI'
func (se *se_FI) PluralsOrdinal() []locales.PluralRule {
	return se.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'se_FI'
func (se *se_FI) PluralsRange() []locales.PluralRule {
	return se.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'se_FI'
func (se *se_FI) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'se_FI'
func (se *se_FI) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'se_FI'
func (se *se_FI) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (se *se_FI) MonthAbbreviated(month time.Month) []byte {
	return se.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (se *se_FI) MonthsAbbreviated() [][]byte {
	return se.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (se *se_FI) MonthNarrow(month time.Month) []byte {
	return se.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (se *se_FI) MonthsNarrow() [][]byte {
	return se.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (se *se_FI) MonthWide(month time.Month) []byte {
	return se.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (se *se_FI) MonthsWide() [][]byte {
	return se.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (se *se_FI) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return se.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (se *se_FI) WeekdaysAbbreviated() [][]byte {
	return se.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (se *se_FI) WeekdayNarrow(weekday time.Weekday) []byte {
	return se.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (se *se_FI) WeekdaysNarrow() [][]byte {
	return se.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (se *se_FI) WeekdayShort(weekday time.Weekday) []byte {
	return se.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (se *se_FI) WeekdaysShort() [][]byte {
	return se.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (se *se_FI) WeekdayWide(weekday time.Weekday) []byte {
	return se.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (se *se_FI) WeekdaysWide() [][]byte {
	return se.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'se_FI' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (se *se_FI) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(se.decimal) + len(se.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, se.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(se.group) - 1; j >= 0; j-- {
					b = append(b, se.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(se.minus) - 1; j >= 0; j-- {
			b = append(b, se.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'se_FI' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (se *se_FI) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(se.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, se.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(se.minus) - 1; j >= 0; j-- {
			b = append(b, se.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, se.percentSuffix...)

	b = append(b, se.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'se_FI'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (se *se_FI) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := se.currencies[currency]
	l := len(s) + len(se.decimal) + len(se.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, se.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(se.group) - 1; j >= 0; j-- {
					b = append(b, se.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(se.minus) - 1; j >= 0; j-- {
			b = append(b, se.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, se.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, se.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'se_FI'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (se *se_FI) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := se.currencies[currency]
	l := len(s) + len(se.decimal) + len(se.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, se.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(se.group) - 1; j >= 0; j-- {
					b = append(b, se.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(se.minus) - 1; j >= 0; j-- {
			b = append(b, se.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, se.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, se.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, se.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'se_FI'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (se *se_FI) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'se_FI'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (se *se_FI) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'se_FI'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (se *se_FI) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'se_FI'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (se *se_FI) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'se_FI'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (se *se_FI) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'se_FI'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (se *se_FI) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'se_FI'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (se *se_FI) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'se_FI'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (se *se_FI) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}
