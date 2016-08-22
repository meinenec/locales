package vo_001

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type vo_001 struct {
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

// New returns a new instance of translator for the 'vo_001' locale
func New() locales.Translator {
	return &vo_001{
		locale:            "vo_001",
		pluralsCardinal:   []locales.PluralRule{2, 6},
		pluralsOrdinal:    nil,
		pluralsRange:      nil,
		decimal:           []byte{},
		group:             []byte{},
		minus:             []byte{},
		percent:           []byte{},
		perMille:          []byte{},
		timeSeparator:     []byte{0x3a},
		currencies:        [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44, 0x20}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e, 0x20}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c, 0x20}, {0x41, 0x4d, 0x44, 0x20}, {0x41, 0x4e, 0x47, 0x20}, {0x41, 0x4f, 0x41, 0x20}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53, 0x20}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44, 0x20}, {0x41, 0x57, 0x47, 0x20}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e, 0x20}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d, 0x20}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44, 0x20}, {0x42, 0x44, 0x54, 0x20}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e, 0x20}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44, 0x20}, {0x42, 0x49, 0x46, 0x20}, {0x42, 0x4d, 0x44, 0x20}, {0x42, 0x4e, 0x44, 0x20}, {0x42, 0x4f, 0x42, 0x20}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x42, 0x52, 0x4c, 0x20}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44, 0x20}, {0x42, 0x54, 0x4e, 0x20}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50, 0x20}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52, 0x20}, {0x42, 0x5a, 0x44, 0x20}, {0x43, 0x41, 0x44, 0x20}, {0x43, 0x44, 0x46, 0x20}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46, 0x20}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50, 0x20}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0x59, 0x20}, {0x43, 0x4f, 0x50, 0x20}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43, 0x20}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43, 0x20}, {0x43, 0x55, 0x50, 0x20}, {0x43, 0x56, 0x45, 0x20}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b, 0x20}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46, 0x20}, {0x44, 0x4b, 0x4b, 0x20}, {0x44, 0x4f, 0x50, 0x20}, {0x44, 0x5a, 0x44, 0x20}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50, 0x20}, {0x45, 0x52, 0x4e, 0x20}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42, 0x20}, {0x45, 0x55, 0x52, 0x20}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44, 0x20}, {0x46, 0x4b, 0x50, 0x20}, {0x46, 0x52, 0x46, 0x20}, {0x47, 0x42, 0x50, 0x20}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c, 0x20}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53, 0x20}, {0x47, 0x49, 0x50, 0x20}, {0x47, 0x4d, 0x44, 0x20}, {0x47, 0x4e, 0x46, 0x20}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51, 0x20}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44, 0x20}, {0x48, 0x4b, 0x44, 0x20}, {0x48, 0x4e, 0x4c, 0x20}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b, 0x20}, {0x48, 0x54, 0x47, 0x20}, {0x48, 0x55, 0x46, 0x20}, {0x49, 0x44, 0x52, 0x20}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0x49, 0x4c, 0x53, 0x20}, {0x49, 0x4e, 0x52, 0x20}, {0x49, 0x51, 0x44, 0x20}, {0x49, 0x52, 0x52, 0x20}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b, 0x20}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44, 0x20}, {0x4a, 0x4f, 0x44, 0x20}, {0x4a, 0x50, 0x59, 0x20}, {0x4b, 0x45, 0x53, 0x20}, {0x4b, 0x47, 0x53, 0x20}, {0x4b, 0x48, 0x52, 0x20}, {0x4b, 0x4d, 0x46, 0x20}, {0x4b, 0x50, 0x57, 0x20}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0x4b, 0x52, 0x57, 0x20}, {0x4b, 0x57, 0x44, 0x20}, {0x4b, 0x59, 0x44, 0x20}, {0x4b, 0x5a, 0x54, 0x20}, {0x4c, 0x41, 0x4b, 0x20}, {0x4c, 0x42, 0x50, 0x20}, {0x4c, 0x4b, 0x52, 0x20}, {0x4c, 0x52, 0x44, 0x20}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c, 0x20}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c, 0x20}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44, 0x20}, {0x4d, 0x41, 0x44, 0x20}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c, 0x20}, {0x4d, 0x47, 0x41, 0x20}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44, 0x20}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b, 0x20}, {0x4d, 0x4e, 0x54, 0x20}, {0x4d, 0x4f, 0x50, 0x20}, {0x4d, 0x52, 0x4f, 0x20}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52, 0x20}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52, 0x20}, {0x4d, 0x57, 0x4b, 0x20}, {0x4d, 0x58, 0x4e, 0x20}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52, 0x20}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e, 0x20}, {0x4e, 0x41, 0x44, 0x20}, {0x4e, 0x47, 0x4e, 0x20}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f, 0x20}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b, 0x20}, {0x4e, 0x50, 0x52, 0x20}, {0x4e, 0x5a, 0x44, 0x20}, {0x4f, 0x4d, 0x52, 0x20}, {0x50, 0x41, 0x42, 0x20}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e, 0x20}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b, 0x20}, {0x50, 0x48, 0x50, 0x20}, {0x50, 0x4b, 0x52, 0x20}, {0x50, 0x4c, 0x4e, 0x20}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47, 0x20}, {0x51, 0x41, 0x52, 0x20}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e, 0x20}, {0x52, 0x53, 0x44, 0x20}, {0x52, 0x55, 0x42, 0x20}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46, 0x20}, {0x53, 0x41, 0x52, 0x20}, {0x53, 0x42, 0x44, 0x20}, {0x53, 0x43, 0x52, 0x20}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47, 0x20}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b, 0x20}, {0x53, 0x47, 0x44, 0x20}, {0x53, 0x48, 0x50, 0x20}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c, 0x20}, {0x53, 0x4f, 0x53, 0x20}, {0x53, 0x52, 0x44, 0x20}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50, 0x20}, {0x53, 0x54, 0x44, 0x20}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50, 0x20}, {0x53, 0x5a, 0x4c, 0x20}, {0x54, 0x48, 0x42, 0x20}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53, 0x20}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54, 0x20}, {0x54, 0x4e, 0x44, 0x20}, {0x54, 0x4f, 0x50, 0x20}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59, 0x20}, {0x54, 0x54, 0x44, 0x20}, {0x54, 0x57, 0x44, 0x20}, {0x54, 0x5a, 0x53, 0x20}, {0x55, 0x41, 0x48, 0x20}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58, 0x20}, {0x55, 0x53, 0x44, 0x20}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55, 0x20}, {0x55, 0x5a, 0x53, 0x20}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46, 0x20}, {0x56, 0x4e, 0x44, 0x20}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56, 0x20}, {0x57, 0x53, 0x54, 0x20}, {0x58, 0x41, 0x46, 0x20}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x58, 0x43, 0x44, 0x20}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x58, 0x4f, 0x46, 0x20}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46, 0x20}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52, 0x20}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52, 0x20}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57, 0x20}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		monthsAbbreviated: [][]uint8{[]uint8(nil), {0x79, 0x61, 0x6e}, {0x66, 0x65, 0x62}, {0x6d, 0xc3, 0xa4, 0x7a}, {0x70, 0x72, 0x6c}, {0x6d, 0x61, 0x79}, {0x79, 0x75, 0x6e}, {0x79, 0x75, 0x6c}, {0x67, 0x73, 0x74}, {0x73, 0x65, 0x74}, {0x74, 0x6f, 0x6e}, {0x6e, 0x6f, 0x76}, {0x64, 0x65, 0x6b}},
		monthsNarrow:      [][]uint8{[]uint8(nil), {0x59}, {0x46}, {0x4d}, {0x50}, {0x4d}, {0x59}, {0x59}, {0x47}, {0x53}, {0x54}, {0x4e}, {0x44}},
		monthsWide:        [][]uint8{[]uint8(nil), {0x79, 0x61, 0x6e, 0x75, 0x6c}, {0x66, 0x65, 0x62, 0x75, 0x6c}, {0x6d, 0xc3, 0xa4, 0x7a, 0x75, 0x6c}, {0x70, 0x72, 0x69, 0x6c, 0x75, 0x6c}, {0x6d, 0x61, 0x79, 0x75, 0x6c}, {0x79, 0x75, 0x6e, 0x75, 0x6c}, {0x79, 0x75, 0x6c, 0x75, 0x6c}, {0x67, 0x75, 0x73, 0x74, 0x75, 0x6c}, {0x73, 0x65, 0x74, 0x75, 0x6c}, {0x74, 0x6f, 0x62, 0x75, 0x6c}, {0x6e, 0x6f, 0x76, 0x75, 0x6c}, {0x64, 0x65, 0x6b, 0x75, 0x6c}},
		daysAbbreviated:   [][]uint8{{0x73, 0x75, 0x2e}, {0x6d, 0x75, 0x2e}, {0x74, 0x75, 0x2e}, {0x76, 0x65, 0x2e}, {0x64, 0xc3, 0xb6, 0x2e}, {0x66, 0x72, 0x2e}, {0x7a, 0xc3, 0xa4, 0x2e}},
		daysNarrow:        [][]uint8{{0x53}, {0x4d}, {0x54}, {0x56}, {0x44}, {0x46}, {0x5a}},
		daysWide:          [][]uint8{{0x73, 0x75, 0x64, 0x65, 0x6c}, {0x6d, 0x75, 0x64, 0x65, 0x6c}, {0x74, 0x75, 0x64, 0x65, 0x6c}, {0x76, 0x65, 0x64, 0x65, 0x6c}, {0x64, 0xc3, 0xb6, 0x64, 0x65, 0x6c}, {0x66, 0x72, 0x69, 0x64, 0x65, 0x6c}, {0x7a, 0xc3, 0xa4, 0x64, 0x65, 0x6c}},
		erasAbbreviated:   [][]uint8{{0x62, 0x2e, 0x20, 0x74, 0x2e, 0x20, 0x6b, 0x72, 0x2e}, {0x70, 0x2e, 0x20, 0x74, 0x2e, 0x20, 0x6b, 0x72, 0x2e}},
		erasNarrow:        [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:          [][]uint8{{0x62, 0x2e, 0x20, 0x74, 0x2e, 0x20, 0x6b, 0x72, 0x2e}, {0x70, 0x2e, 0x20, 0x74, 0x2e, 0x20, 0x6b, 0x72, 0x2e}},
		timezones:         map[string][]uint8{"SGT": {0x53, 0x47, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "ECT": {0x45, 0x43, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "PDT": {0x50, 0x44, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "MEZ": {0x4d, 0x45, 0x5a}, "ARST": {0x41, 0x52, 0x53, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "GFT": {0x47, 0x46, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "GMT": {0x47, 0x4d, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "ADT": {0x41, 0x44, 0x54}, "EST": {0x45, 0x53, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "CDT": {0x43, 0x44, 0x54}, "IST": {0x49, 0x53, 0x54}, "ART": {0x41, 0x52, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "GYT": {0x47, 0x59, 0x54}, "VET": {0x56, 0x45, 0x54}, "COT": {0x43, 0x4f, 0x54}, "AST": {0x41, 0x53, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "UYT": {0x55, 0x59, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "HNT": {0x48, 0x4e, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "CST": {0x43, 0x53, 0x54}, "PST": {0x50, 0x53, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "HAT": {0x48, 0x41, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "SRT": {0x53, 0x52, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "WAT": {0x57, 0x41, 0x54}, "MST": {0x4d, 0x53, 0x54}, "EAT": {0x45, 0x41, 0x54}, "JST": {0x4a, 0x53, 0x54}, "WIT": {0x57, 0x49, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "CAT": {0x43, 0x41, 0x54}, "BT": {0x42, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "EDT": {0x45, 0x44, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "WAST": {0x57, 0x41, 0x53, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "WIB": {0x57, 0x49, 0x42}, "BOT": {0x42, 0x4f, 0x54}},
	}
}

// Locale returns the current translators string locale
func (vo *vo_001) Locale() string {
	return vo.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'vo_001'
func (vo *vo_001) PluralsCardinal() []locales.PluralRule {
	return vo.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'vo_001'
func (vo *vo_001) PluralsOrdinal() []locales.PluralRule {
	return vo.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'vo_001'
func (vo *vo_001) PluralsRange() []locales.PluralRule {
	return vo.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'vo_001'
func (vo *vo_001) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'vo_001'
func (vo *vo_001) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'vo_001'
func (vo *vo_001) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (vo *vo_001) MonthAbbreviated(month time.Month) []byte {
	return vo.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (vo *vo_001) MonthsAbbreviated() [][]byte {
	return vo.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (vo *vo_001) MonthNarrow(month time.Month) []byte {
	return vo.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (vo *vo_001) MonthsNarrow() [][]byte {
	return vo.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (vo *vo_001) MonthWide(month time.Month) []byte {
	return vo.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (vo *vo_001) MonthsWide() [][]byte {
	return vo.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (vo *vo_001) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return vo.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (vo *vo_001) WeekdaysAbbreviated() [][]byte {
	return vo.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (vo *vo_001) WeekdayNarrow(weekday time.Weekday) []byte {
	return vo.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (vo *vo_001) WeekdaysNarrow() [][]byte {
	return vo.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (vo *vo_001) WeekdayShort(weekday time.Weekday) []byte {
	return vo.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (vo *vo_001) WeekdaysShort() [][]byte {
	return vo.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (vo *vo_001) WeekdayWide(weekday time.Weekday) []byte {
	return vo.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (vo *vo_001) WeekdaysWide() [][]byte {
	return vo.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'vo_001' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vo *vo_001) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'vo_001' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (vo *vo_001) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'vo_001'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vo *vo_001) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := vo.currencies[currency]
	return append(append([]byte{}, symbol...), s...)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'vo_001'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vo *vo_001) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := vo.currencies[currency]
	return append(append([]byte{}, symbol...), s...)
}

// FmtDateShort returns the short date representation of 't' for 'vo_001'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vo *vo_001) FmtDateShort(t time.Time) []byte {

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

// FmtDateMedium returns the medium date representation of 't' for 'vo_001'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vo *vo_001) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vo.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'vo_001'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vo *vo_001) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vo.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'vo_001'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vo *vo_001) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vo.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x4d, 0x4d, 0x4d, 0x4d, 0x27, 0x61, 0x27}...)
	b = append(b, []byte{0x27, 0x64, 0x27, 0x2e, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x27, 0x64, 0x27, 0x2e, 0x20, 0x64, 0x27, 0x69, 0x64, 0x27}...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'vo_001'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vo *vo_001) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'vo_001'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vo *vo_001) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'vo_001'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vo *vo_001) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'vo_001'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (vo *vo_001) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := vo.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
