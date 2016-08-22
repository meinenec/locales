package lag

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type lag struct {
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

// New returns a new instance of translator for the 'lag' locale
func New() locales.Translator {
	return &lag{
		locale:                 "lag",
		pluralsCardinal:        []locales.PluralRule{1, 2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                []byte{},
		group:                  []byte{},
		minus:                  []byte{},
		percent:                []byte{},
		perMille:               []byte{},
		timeSeparator:          []byte{0x3a},
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44, 0x20}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e, 0x20}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c, 0x20}, {0x41, 0x4d, 0x44, 0x20}, {0x41, 0x4e, 0x47, 0x20}, {0x41, 0x4f, 0x41, 0x20}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53, 0x20}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44, 0x20}, {0x41, 0x57, 0x47, 0x20}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e, 0x20}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d, 0x20}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44, 0x20}, {0x42, 0x44, 0x54, 0x20}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e, 0x20}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44, 0x20}, {0x42, 0x49, 0x46, 0x20}, {0x42, 0x4d, 0x44, 0x20}, {0x42, 0x4e, 0x44, 0x20}, {0x42, 0x4f, 0x42, 0x20}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x42, 0x52, 0x4c, 0x20}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44, 0x20}, {0x42, 0x54, 0x4e, 0x20}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50, 0x20}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52, 0x20}, {0x42, 0x5a, 0x44, 0x20}, {0x43, 0x41, 0x44, 0x20}, {0x43, 0x44, 0x46, 0x20}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46, 0x20}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50, 0x20}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0x59, 0x20}, {0x43, 0x4f, 0x50, 0x20}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43, 0x20}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43, 0x20}, {0x43, 0x55, 0x50, 0x20}, {0x43, 0x56, 0x45, 0x20}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b, 0x20}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46, 0x20}, {0x44, 0x4b, 0x4b, 0x20}, {0x44, 0x4f, 0x50, 0x20}, {0x44, 0x5a, 0x44, 0x20}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50, 0x20}, {0x45, 0x52, 0x4e, 0x20}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42, 0x20}, {0x45, 0x55, 0x52, 0x20}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44, 0x20}, {0x46, 0x4b, 0x50, 0x20}, {0x46, 0x52, 0x46, 0x20}, {0x47, 0x42, 0x50, 0x20}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c, 0x20}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53, 0x20}, {0x47, 0x49, 0x50, 0x20}, {0x47, 0x4d, 0x44, 0x20}, {0x47, 0x4e, 0x46, 0x20}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51, 0x20}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44, 0x20}, {0x48, 0x4b, 0x44, 0x20}, {0x48, 0x4e, 0x4c, 0x20}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b, 0x20}, {0x48, 0x54, 0x47, 0x20}, {0x48, 0x55, 0x46, 0x20}, {0x49, 0x44, 0x52, 0x20}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0x49, 0x4c, 0x53, 0x20}, {0x49, 0x4e, 0x52, 0x20}, {0x49, 0x51, 0x44, 0x20}, {0x49, 0x52, 0x52, 0x20}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b, 0x20}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44, 0x20}, {0x4a, 0x4f, 0x44, 0x20}, {0x4a, 0x50, 0x59, 0x20}, {0x4b, 0x45, 0x53, 0x20}, {0x4b, 0x47, 0x53, 0x20}, {0x4b, 0x48, 0x52, 0x20}, {0x4b, 0x4d, 0x46, 0x20}, {0x4b, 0x50, 0x57, 0x20}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0x4b, 0x52, 0x57, 0x20}, {0x4b, 0x57, 0x44, 0x20}, {0x4b, 0x59, 0x44, 0x20}, {0x4b, 0x5a, 0x54, 0x20}, {0x4c, 0x41, 0x4b, 0x20}, {0x4c, 0x42, 0x50, 0x20}, {0x4c, 0x4b, 0x52, 0x20}, {0x4c, 0x52, 0x44, 0x20}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c, 0x20}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c, 0x20}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44, 0x20}, {0x4d, 0x41, 0x44, 0x20}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c, 0x20}, {0x4d, 0x47, 0x41, 0x20}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44, 0x20}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b, 0x20}, {0x4d, 0x4e, 0x54, 0x20}, {0x4d, 0x4f, 0x50, 0x20}, {0x4d, 0x52, 0x4f, 0x20}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52, 0x20}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52, 0x20}, {0x4d, 0x57, 0x4b, 0x20}, {0x4d, 0x58, 0x4e, 0x20}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52, 0x20}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e, 0x20}, {0x4e, 0x41, 0x44, 0x20}, {0x4e, 0x47, 0x4e, 0x20}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f, 0x20}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b, 0x20}, {0x4e, 0x50, 0x52, 0x20}, {0x4e, 0x5a, 0x44, 0x20}, {0x4f, 0x4d, 0x52, 0x20}, {0x50, 0x41, 0x42, 0x20}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e, 0x20}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b, 0x20}, {0x50, 0x48, 0x50, 0x20}, {0x50, 0x4b, 0x52, 0x20}, {0x50, 0x4c, 0x4e, 0x20}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47, 0x20}, {0x51, 0x41, 0x52, 0x20}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e, 0x20}, {0x52, 0x53, 0x44, 0x20}, {0x52, 0x55, 0x42, 0x20}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46, 0x20}, {0x53, 0x41, 0x52, 0x20}, {0x53, 0x42, 0x44, 0x20}, {0x53, 0x43, 0x52, 0x20}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47, 0x20}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b, 0x20}, {0x53, 0x47, 0x44, 0x20}, {0x53, 0x48, 0x50, 0x20}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c, 0x20}, {0x53, 0x4f, 0x53, 0x20}, {0x53, 0x52, 0x44, 0x20}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50, 0x20}, {0x53, 0x54, 0x44, 0x20}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50, 0x20}, {0x53, 0x5a, 0x4c, 0x20}, {0x54, 0x48, 0x42, 0x20}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53, 0x20}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54, 0x20}, {0x54, 0x4e, 0x44, 0x20}, {0x54, 0x4f, 0x50, 0x20}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59, 0x20}, {0x54, 0x54, 0x44, 0x20}, {0x54, 0x57, 0x44, 0x20}, {0x54, 0x53, 0x68}, {0x55, 0x41, 0x48, 0x20}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58, 0x20}, {0x55, 0x53, 0x44, 0x20}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55, 0x20}, {0x55, 0x5a, 0x53, 0x20}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46, 0x20}, {0x56, 0x4e, 0x44, 0x20}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56, 0x20}, {0x57, 0x53, 0x54, 0x20}, {0x58, 0x41, 0x46, 0x20}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x58, 0x43, 0x44, 0x20}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x58, 0x4f, 0x46, 0x20}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46, 0x20}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52, 0x20}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52, 0x20}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57, 0x20}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0x4b},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0x4b},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x46, 0xc3, 0xba, 0x6e, 0x67, 0x61, 0x74, 0xc9, 0xa8}, {0x4e, 0x61, 0x61, 0x6e, 0xc9, 0xa8}, {0x4b, 0x65, 0x65, 0x6e, 0x64, 0x61}, {0x49, 0x6b, 0xc3, 0xba, 0x6d, 0x69}, {0x49, 0x6e, 0x79, 0x61, 0x6d, 0x62, 0x61, 0x6c, 0x61}, {0x49, 0x64, 0x77, 0x61, 0x61, 0x74, 0x61}, {0x4d, 0xca, 0x89, 0xca, 0x89, 0x6e, 0x63, 0x68, 0xc9, 0xa8}, {0x56, 0xc9, 0xa8, 0xc9, 0xa8, 0x72, 0xc9, 0xa8}, {0x53, 0x61, 0x61, 0x74, 0xca, 0x89}, {0x49, 0x6e, 0x79, 0x69}, {0x53, 0x61, 0x61, 0x6e, 0x6f}, {0x53, 0x61, 0x73, 0x61, 0x74, 0xca, 0x89}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x46}, {0x4e}, {0x4b}, {0x49}, {0x49}, {0x49}, {0x4d}, {0x56}, {0x53}, {0x49}, {0x53}, {0x53}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x4b, 0xca, 0x89, 0x66, 0xc3, 0xba, 0x6e, 0x67, 0x61, 0x74, 0xc9, 0xa8}, {0x4b, 0xca, 0x89, 0x6e, 0x61, 0x61, 0x6e, 0xc9, 0xa8}, {0x4b, 0xca, 0x89, 0x6b, 0x65, 0x65, 0x6e, 0x64, 0x61}, {0x4b, 0x77, 0x69, 0x69, 0x6b, 0x75, 0x6d, 0x69}, {0x4b, 0x77, 0x69, 0x69, 0x6e, 0x79, 0x61, 0x6d, 0x62, 0xc3, 0xa1, 0x6c, 0x61}, {0x4b, 0x77, 0x69, 0x69, 0x64, 0x77, 0x61, 0x61, 0x74, 0x61}, {0x4b, 0xca, 0x89, 0x6d, 0xca, 0x89, 0xca, 0x89, 0x6e, 0x63, 0x68, 0xc9, 0xa8}, {0x4b, 0xca, 0x89, 0x76, 0xc9, 0xa8, 0xc9, 0xa8, 0x72, 0xc9, 0xa8}, {0x4b, 0xca, 0x89, 0x73, 0x61, 0x61, 0x74, 0xca, 0x89}, {0x4b, 0x77, 0x69, 0x69, 0x6e, 0x79, 0x69}, {0x4b, 0xca, 0x89, 0x73, 0x61, 0x61, 0x6e, 0x6f}, {0x4b, 0xca, 0x89, 0x73, 0x61, 0x73, 0x61, 0x74, 0xca, 0x89}},
		daysAbbreviated:        [][]uint8{{0x50, 0xc3, 0xad, 0x69, 0x6c, 0x69}, {0x54, 0xc3, 0xa1, 0x61, 0x74, 0x75}, {0xc3, 0x8d, 0x6e, 0x65}, {0x54, 0xc3, 0xa1, 0x61, 0x6e, 0x6f}, {0x41, 0x6c, 0x68}, {0x49, 0x6a, 0x6d}, {0x4d, 0xc3, 0xb3, 0x6f, 0x73, 0x69}},
		daysNarrow:             [][]uint8{{0x50}, {0x54}, {0x45}, {0x4f}, {0x41}, {0x49}, {0x4d}},
		daysWide:               [][]uint8{{0x4a, 0x75, 0x6d, 0x61, 0x70, 0xc3, 0xad, 0x69, 0x72, 0x69}, {0x4a, 0x75, 0x6d, 0x61, 0x74, 0xc3, 0xa1, 0x74, 0x75}, {0x4a, 0x75, 0x6d, 0x61, 0xc3, 0xad, 0x6e, 0x65}, {0x4a, 0x75, 0x6d, 0x61, 0x74, 0xc3, 0xa1, 0x61, 0x6e, 0x6f}, {0x41, 0x6c, 0x61, 0x6d, 0xc3, 0xad, 0x69, 0x73, 0x69}, {0x49, 0x6a, 0x75, 0x6d, 0xc3, 0xa1, 0x61}, {0x4a, 0x75, 0x6d, 0x61, 0x6d, 0xc3, 0xb3, 0x6f, 0x73, 0x69}},
		periodsAbbreviated:     [][]uint8{{0x54, 0x4f, 0x4f}, {0x4d, 0x55, 0x55}},
		periodsWide:            [][]uint8{{0x54, 0x4f, 0x4f}, {0x4d, 0x55, 0x55}},
		erasAbbreviated:        [][]uint8{{0x4b, 0x53, 0x41}, {0x4b, 0x41}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x4b, 0xc9, 0xa8, 0x72, 0xc9, 0xa8, 0x73, 0x69, 0x74, 0xca, 0x89, 0x20, 0x73, 0xc9, 0xa8, 0x20, 0x61, 0x6e, 0x61, 0x76, 0x79, 0x61, 0x61, 0x6c}, {0x4b, 0xc9, 0xa8, 0x72, 0xc9, 0xa8, 0x73, 0x69, 0x74, 0xca, 0x89, 0x20, 0x61, 0x6b, 0x61, 0x76, 0x79, 0x61, 0x61, 0x6c, 0x77, 0x65}},
		timezones:              map[string][]uint8{"CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "EAT": {0x45, 0x41, 0x54}, "WIT": {0x57, 0x49, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "COT": {0x43, 0x4f, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "ADT": {0x41, 0x44, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "WAT": {0x57, 0x41, 0x54}, "CST": {0x43, 0x53, 0x54}, "ART": {0x41, 0x52, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "EST": {0x45, 0x53, 0x54}, "PST": {0x50, 0x53, 0x54}, "PDT": {0x50, 0x44, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "VET": {0x56, 0x45, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "CDT": {0x43, 0x44, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "BT": {0x42, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "GFT": {0x47, 0x46, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "CAT": {0x43, 0x41, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "SRT": {0x53, 0x52, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "AST": {0x41, 0x53, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "GYT": {0x47, 0x59, 0x54}, "ECT": {0x45, 0x43, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "BOT": {0x42, 0x4f, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "HAT": {0x48, 0x41, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "MST": {0x4d, 0x53, 0x54}, "EDT": {0x45, 0x44, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "IST": {0x49, 0x53, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "SGT": {0x53, 0x47, 0x54}, "UYT": {0x55, 0x59, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "WIB": {0x57, 0x49, 0x42}},
	}
}

// Locale returns the current translators string locale
func (lag *lag) Locale() string {
	return lag.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'lag'
func (lag *lag) PluralsCardinal() []locales.PluralRule {
	return lag.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'lag'
func (lag *lag) PluralsOrdinal() []locales.PluralRule {
	return lag.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'lag'
func (lag *lag) PluralsRange() []locales.PluralRule {
	return lag.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'lag'
func (lag *lag) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if n == 0 {
		return locales.PluralRuleZero
	} else if (i == 0 || i == 1) && n != 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'lag'
func (lag *lag) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'lag'
func (lag *lag) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (lag *lag) MonthAbbreviated(month time.Month) []byte {
	return lag.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (lag *lag) MonthsAbbreviated() [][]byte {
	return lag.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (lag *lag) MonthNarrow(month time.Month) []byte {
	return lag.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (lag *lag) MonthsNarrow() [][]byte {
	return lag.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (lag *lag) MonthWide(month time.Month) []byte {
	return lag.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (lag *lag) MonthsWide() [][]byte {
	return lag.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (lag *lag) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return lag.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (lag *lag) WeekdaysAbbreviated() [][]byte {
	return lag.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (lag *lag) WeekdayNarrow(weekday time.Weekday) []byte {
	return lag.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (lag *lag) WeekdaysNarrow() [][]byte {
	return lag.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (lag *lag) WeekdayShort(weekday time.Weekday) []byte {
	return lag.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (lag *lag) WeekdaysShort() [][]byte {
	return lag.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (lag *lag) WeekdayWide(weekday time.Weekday) []byte {
	return lag.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (lag *lag) WeekdaysWide() [][]byte {
	return lag.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'lag' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lag *lag) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'lag' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (lag *lag) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'lag'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lag *lag) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lag.currencies[currency]
	l := len(s) + len(lag.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(lag.decimal) - 1; j >= 0; j-- {
				b = append(b, lag.decimal[j])
			}

			continue
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(lag.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, lag.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(lag.minus) - 1; j >= 0; j-- {
			b = append(b, lag.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, lag.currencyPositiveSuffix...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'lag'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lag *lag) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lag.currencies[currency]
	l := len(s) + len(lag.decimal)

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(lag.decimal) - 1; j >= 0; j-- {
				b = append(b, lag.decimal[j])
			}

			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(lag.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, lag.currencyNegativePrefix[j])
		}

		for j := len(lag.minus) - 1; j >= 0; j-- {
			b = append(b, lag.minus[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(lag.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, lag.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, lag.currencyNegativeSuffix...)
	} else {

		b = append(b, lag.currencyPositiveSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'lag'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lag *lag) FmtDateShort(t time.Time) []byte {

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
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'lag'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lag *lag) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, lag.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'lag'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lag *lag) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, lag.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'lag'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lag *lag) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, lag.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, lag.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'lag'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lag *lag) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lag.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'lag'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lag *lag) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lag.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lag.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'lag'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lag *lag) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lag.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lag.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'lag'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (lag *lag) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lag.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lag.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := lag.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
