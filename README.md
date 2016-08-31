## locales
<img align="right" src="https://raw.githubusercontent.com/go-playground/locales/master/logo.png">
![Project status](https://img.shields.io/badge/version-0.9.0-green.svg)
[![Build Status](https://semaphoreci.com/api/v1/joeybloggs/locales/branches/master/badge.svg)](https://semaphoreci.com/joeybloggs/locales)
[![Coverage Status](https://coveralls.io/repos/github/go-playground/locales/badge.svg?branch=master)](https://coveralls.io/github/go-playground/locales?branch=master)
[![GoDoc](https://godoc.org/github.com/go-playground/locales?status.svg)](https://godoc.org/github.com/go-playground/locales)
![License](https://img.shields.io/dub/l/vibe-d.svg)
[![Gitter](https://badges.gitter.im/go-playground/locales.svg)](https://gitter.im/go-playground/locales?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

Locales is a set of locales generated from the [Unicode CLDR Project](http://cldr.unicode.org/) which can be used independently or within
an i18n package; these were built for use with, but not exclusive to, [Universal Translator](https://github.com/go-playground/universal-translator).

Features
--------
- [x] Rules generated from the latest [CLDR](http://cldr.unicode.org/index/downloads) data, v29
- [x] Contains Cardinal, Ordinal and Range Plural Rules
- [x] Contains Month, Weekday and Timezone translations built in
- [x] Contains Date & Time formatting functions
- [x] Contains Number, Currency, Accounting and Percent formatting functions
- [x] Supports the "Gregorian" calendar only ( my time isn't unlimited, had to draw the line somewhere )

Full Tests
--------------------
I could sure use your help adding tests for every locale, it is a huge undertaking and I just don't have the free time to do it all at the moment;
any help would be **greatly appreciated!!!!** please see [issue](https://github.com/go-playground/locales/issues/1) for details.

Installation
-----------

Use go get 

```go
go get github.com/go-playground/locales
```  

NOTES
--------
You'll notice most return types are []byte, this is because most of the time the results will be concatenated with a larger body
of text and can avoid some allocations if already appending to a byte array, otherwise just cast as string.

Usage
-------
```go
package main

import (
	"fmt"
	"time"

	"github.com/go-playground/locales/currency"
	"github.com/go-playground/locales/en_CA"
)

const (
	dateTimeString = "Jan 2, 2006 at 3:04:05pm"
)

func main() {

	loc, _ := time.LoadLocation("America/Toronto")
	datetime := time.Date(2016, 02, 03, 9, 0, 1, 0, loc)

	l := en_CA.New()

	// Dates
	fmt.Println(string(l.FmtDateFull(datetime)))
	fmt.Println(string(l.FmtDateLong(datetime)))
	fmt.Println(string(l.FmtDateMedium(datetime)))
	fmt.Println(string(l.FmtDateShort(datetime)))

	// Times
	fmt.Println(string(l.FmtTimeFull(datetime)))
	fmt.Println(string(l.FmtTimeLong(datetime)))
	fmt.Println(string(l.FmtTimeMedium(datetime)))
	fmt.Println(string(l.FmtTimeShort(datetime)))

	// Months Wide
	fmt.Println(string(l.MonthWide(time.January)))
	fmt.Println(string(l.MonthWide(time.February)))
	fmt.Println(string(l.MonthWide(time.March)))
	// ...

	// Months Abbreviated
	fmt.Println(string(l.MonthAbbreviated(time.January)))
	fmt.Println(string(l.MonthAbbreviated(time.February)))
	fmt.Println(string(l.MonthAbbreviated(time.March)))
	// ...

	// Months Narrow
	fmt.Println(string(l.MonthNarrow(time.January)))
	fmt.Println(string(l.MonthNarrow(time.February)))
	fmt.Println(string(l.MonthNarrow(time.March)))
	// ...

	// Weekdays Wide
	fmt.Println(string(l.WeekdayWide(time.Sunday)))
	fmt.Println(string(l.WeekdayWide(time.Monday)))
	fmt.Println(string(l.WeekdayWide(time.Tuesday)))
	// ...

	// Weekdays Abbreviated
	fmt.Println(string(l.WeekdayAbbreviated(time.Sunday)))
	fmt.Println(string(l.WeekdayAbbreviated(time.Monday)))
	fmt.Println(string(l.WeekdayAbbreviated(time.Tuesday)))
	// ...

	// Weekdays Short
	fmt.Println(string(l.WeekdayShort(time.Sunday)))
	fmt.Println(string(l.WeekdayShort(time.Monday)))
	fmt.Println(string(l.WeekdayShort(time.Tuesday)))
	// ...

	// Weekdays Narrow
	fmt.Println(string(l.WeekdayNarrow(time.Sunday)))
	fmt.Println(string(l.WeekdayNarrow(time.Monday)))
	fmt.Println(string(l.WeekdayNarrow(time.Tuesday)))
	// ...

	var f64 float64

	f64 = -10356.4523

	// Number
	fmt.Println(string(l.FmtNumber(f64, 2)))

	// Currency
	fmt.Println(string(l.FmtCurrency(f64, 2, currency.CAD)))
	fmt.Println(string(l.FmtCurrency(f64, 2, currency.USD)))

	// Accounting
	fmt.Println(string(l.FmtAccounting(f64, 2, currency.CAD)))
	fmt.Println(string(l.FmtAccounting(f64, 2, currency.USD)))

	f64 = 78.12

	// Percent
	fmt.Println(string(l.FmtPercent(f64, 0)))

	// Plural Rules for locale, so you know what rules you must cover
	fmt.Println(l.PluralsCardinal())
	fmt.Println(l.PluralsOrdinal())

	// Cardinal Plural Rules
	fmt.Println(l.CardinalPluralRule(1, 0))
	fmt.Println(l.CardinalPluralRule(1.0, 0))
	fmt.Println(l.CardinalPluralRule(1.0, 1))
	fmt.Println(l.CardinalPluralRule(3, 0))

	// Ordinal Plural Rules
	fmt.Println(l.OrdinalPluralRule(21, 0)) // 21st
	fmt.Println(l.OrdinalPluralRule(22, 0)) // 22nd
	fmt.Println(l.OrdinalPluralRule(33, 0)) // 33rd
	fmt.Println(l.OrdinalPluralRule(34, 0)) // 34th

	// Range Plural Rules
	fmt.Println(l.RangePluralRule(1, 0, 1, 0)) // 1-1
	fmt.Println(l.RangePluralRule(1, 0, 2, 0)) // 1-2
	fmt.Println(l.RangePluralRule(5, 0, 8, 0)) // 5-8
}
```

NOTES:
-------
These rules were generated from the [Unicode CLDR Project](http://cldr.unicode.org/), if you encounter any issues
I strongly encourage contributing to the CLDR project to get the locale information corrected and the next time 
these locales are regenerated the fix will come with.

I do however realize that time constraints are often important and so there are two options:

1. Create your own locale, copy, paste and modify, and ensure it complies with the `Translator` interface.
2. Add an exception in the locale generation code directly and once regenerated, fix will be in place.

Please to not make fixes inside the locale files, they WILL get overwritten when the locales are regenerated.

License
------
Distributed under MIT License, please see license file in code for more details.
