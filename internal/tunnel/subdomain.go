package tunnel

import (
	"fmt"
	"strings"
)

var blockedSubdomains = []string{
	"paypal",
	"apple",
	"google",
	"microsoft",
	"facebook",
	"instagram",
	"amazon",
	"netflix",
	"spotify",
	"twitter",
	"linkedin",
	"github",
	"dropbox",
	"icloud",
	"chase",
	"wellsfargo",
	"bankofamerica",
	"citibank",
	"coinbase",
	"binance",
	"stripe",
	"venmo",
	"cashapp",
	"zelle",
	"metamask",
	"outlook",
	"hotmail",
	"yahoo",
	"whatsapp",
	"telegram",
	"signal",
	"discord",
	"slack",
	"zoom",
	"docusign",
	"adobe",
	"salesforce",
	"shopify",
	"ebay",
	"walmart",
	"usps",
	"fedex",
	"ups",
	"dhl",
}

func ValidateSubdomain(subdomain string) error {
	if subdomain == "" {
		return nil
	}

	lower := strings.ToLower(subdomain)
	normalized := strings.ReplaceAll(lower, "-", "")
	normalized = strings.ReplaceAll(normalized, ".", "")

	for _, brand := range blockedSubdomains {
		if normalized == brand || strings.Contains(normalized, brand) {
			return fmt.Errorf("subdomain %q is not allowed: resembles a protected brand name", subdomain)
		}
	}

	return nil
}
