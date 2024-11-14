package util

// Constraints for all supported currencies
const (
	NGN = "NGN"
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case NGN, USD, EUR, CAD:
		return true
	}
	return false
}