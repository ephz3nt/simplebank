package util

// Constants for all supported currencies.
const (
	USD = "USD"
	CAD = "CAD"
	EUR = "EUR"
)

// IsSupportCurrency returns true if the currency is supported
func IsSupportCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
