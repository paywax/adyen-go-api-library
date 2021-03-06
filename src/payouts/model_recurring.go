/*
 * Adyen Payout API
 *
 * A set of API endpoints that allow you to store payout details, confirm, or decline a payout.  For more information, refer to [Online payouts](https://docs.adyen.com/checkout/online-payouts).
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payouts

import (
	"time"
)

// Recurring struct for Recurring
type Recurring struct {
	// The type of recurring contract to be used. Possible values: * `ONECLICK` – Payments details can be used to initiate a one-click payment, where the shopper enters the [card security code (CVC/CVV)](https://docs.adyen.com/payments-fundamentals/payment-glossary#card-security-code-cvc-cvv-cid). * `RECURRING` – Payments details can be used without the card security code to initiate [card-not-present transactions](https://docs.adyen.com/payments-fundamentals/payment-glossary#card-not-present-cnp). * `ONECLICK,RECURRING` – Payments details can be used regardless of whether the shopper is on your site or not. * `PAYOUT` – Payments details can be used to [make a payout](https://docs.adyen.com/checkout/online-payouts).
	Contract string `json:"contract,omitempty"`
	// A descriptive name for this detail.
	RecurringDetailName string `json:"recurringDetailName,omitempty"`
	// Date after which no further authorisations shall be performed. Only for 3D Secure 2.
	RecurringExpiry *time.Time `json:"recurringExpiry,omitempty"`
	// Minimum number of days between authorisations. Only for 3D Secure 2.
	RecurringFrequency string `json:"recurringFrequency,omitempty"`
	// The name of the token service.
	TokenService string `json:"tokenService,omitempty"`
}
