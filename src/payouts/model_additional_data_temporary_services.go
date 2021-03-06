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

// AdditionalDataTemporaryServices struct for AdditionalDataTemporaryServices
type AdditionalDataTemporaryServices struct {
	// Customer code, if supplied by a customer. * Encoding: ASCII * maxLength: 25
	EnhancedSchemeDataCustomerReference string `json:"enhancedSchemeData.customerReference,omitempty"`
	// Total tax amount, in minor units. For example, 2000 means USD 20.00 * maxLength: 12
	EnhancedSchemeDataTotalTaxAmount string `json:"enhancedSchemeData.totalTaxAmount,omitempty"`
	// Name or ID associated with the individual working in a temporary capacity. * maxLength: 40
	EnhancedSchemeDataEmployeeName string `json:"enhancedSchemeData.employeeName,omitempty"`
	// Description of the job or task of the individual working in a temporary capacity. * maxLength: 40
	EnhancedSchemeDataJobDescription string `json:"enhancedSchemeData.jobDescription,omitempty"`
	// Date for the beginning of the pay period. * Format: ddMMyy * maxLength: 6
	EnhancedSchemeDataTempStartDate string `json:"enhancedSchemeData.tempStartDate,omitempty"`
	// Date of the end of the billing cycle. * Format: ddMMyy * maxLength: 6
	EnhancedSchemeDataTempWeekEnding string `json:"enhancedSchemeData.tempWeekEnding,omitempty"`
	// Name of the individual requesting temporary services. * maxLength: 40
	EnhancedSchemeDataRequestName string `json:"enhancedSchemeData.requestName,omitempty"`
	// Amount of time worked during a normal operation for the task or job. * maxLength: 7
	EnhancedSchemeDataRegularHoursWorked string `json:"enhancedSchemeData.regularHoursWorked,omitempty"`
	// Amount paid per regular hours worked, minor units. * maxLength: 7
	EnhancedSchemeDataRegularHoursRate string `json:"enhancedSchemeData.regularHoursRate,omitempty"`
}
