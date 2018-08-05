package helper

// Check if currency is a valid currency
func ValidateCurrency(userCurrency string) bool {
	validCurrency := []string{"USD", "CAD", "EUR", "AED", "AFN", "ALL", "AMD", "ANG", "AOA",
		"ARS", "AUD", "AZN", "BAM", "BDT", "BGN", "BHD", "BIF", "BND", "BMD", "BOB", "BRL", "BTN", "BWP",
		"BYR", "BZD", "CDF", "CHF", "CLP", "CNY", "COP", "CRC", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD",
		"EEK", "EGP", "ERN", "ETB", "FKP", "GBP", "GEL", "GHS", "GIP", "GNF", "GTQ", "HKD", "HNL", "HRK",
		"HUF", "IDR", "ILS", "INR", "IQD", "IRR", "ISK", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF",
		"KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LTL", "LVL", "LYD", "MAD", "MDL", "MGA",
		"MKD", "MMK", "MOP", "MUR", "MWK", "MXN", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD",
		"OMR", "PAB", "PEN", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD",
		"SDG", "SEK", "SGD", "SLL", "SOS", "SSP", "STD", "STN", "SYP", "SZL", "THB", "TJS", "TND", "TOP",
		"TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "UYU", "UZS", "VEF", "VND", "VUV", "XAF", "XCD", "XOF",
		"XPF", "YER", "ZAR", "ZMK",
	}
	for _, currency := range validCurrency {
		if userCurrency == currency {
			return true
		}
	}
	return false
}
