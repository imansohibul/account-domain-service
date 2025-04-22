package entity

// Currency represents the currency type
type Currency int16

// Currency is an enumeration of currency types
// The enumeration values are:
// 0 - Unspecified
// 1 - IDR (Indonesian Rupiah)
const (
	CurrencyUnspecified Currency = iota
	CurrencyIDR
)
