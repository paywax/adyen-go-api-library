# PaymentSetupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | [**AdditionalData | CommonAdditionalData | 3DSecureAdditionalData | AirlineAdditionalData | CarRentalAdditionalData | Level23AdditionalData | LodgingAdditionalData | OpenInvoiceAdditionalData | RatepayAdditionalData | RetryAdditionalData | RiskAdditionalData | RiskStandaloneAdditionalData | TemporaryServices**]() | This field contains additional data, which may be required for a particular payment request.  The &#x60;additionalData&#x60; object consists of entries, each of which includes the key and value. | [optional] 
**AllowedPaymentMethods** | **[]string** | List of payments methods to be presented to the shopper. To refer to payment methods, use their &#x60;paymentMethod.type&#x60; from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: &#x60;\&quot;allowedPaymentMethods\&quot;:[\&quot;ideal\&quot;,\&quot;giropay\&quot;]&#x60; | [optional] 
**Amount** | [**Amount**](Amount.md) |  | 
**ApplicationInfo** | [**ApplicationInfo**](ApplicationInfo.md) |  | [optional] 
**BillingAddress** | [**Address**](Address.md) |  | [optional] 
**BlockedPaymentMethods** | **[]string** | List of payments methods to be hidden from the shopper. To refer to payment methods, use their &#x60;paymentMethod.type&#x60; from [Payment methods overview](https://docs.adyen.com/payment-methods).  Example: &#x60;\&quot;blockedPaymentMethods\&quot;:[\&quot;ideal\&quot;,\&quot;giropay\&quot;]&#x60; | [optional] 
**CaptureDelayHours** | **int32** | The delay between the authorisation and scheduled auto-capture, specified in hours. | [optional] 
**Channel** | **string** | The platform where a payment transaction takes place. This field is optional for filtering out payment methods that are only available on specific platforms. If this value is not set, then we will try to infer it from the &#x60;sdkVersion&#x60; or &#x60;token&#x60;.  Possible values: * iOS * Android * Web | [optional] 
**Company** | [**Company**](Company.md) |  | [optional] 
**Configuration** | [**Configuration**](Configuration.md) |  | [optional] 
**ConversionId** | **string** | Conversion ID that corresponds to the Id generated for tracking user payment journey. | [optional] 
**CountryCode** | **string** | The shopper country.  Format: [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) Example: NL or DE | 
**DateOfBirth** | [**time.Time**](time.Time.md) | The shopper&#39;s date of birth.  Format [ISO-8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DD | [optional] 
**DccQuote** | [**ForexQuote**](ForexQuote.md) |  | [optional] 
**DeliveryAddress** | [**Address**](Address.md) |  | [optional] 
**DeliveryDate** | [**time.Time**](time.Time.md) | The date and time the purchased goods should be delivered.  Format [ISO 8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DDThh:mm:ss.sssTZD  Example: 2017-07-17T13:42:40.428+01:00 | [optional] 
**EnableOneClick** | **bool** | When true and &#x60;shopperReference&#x60; is provided, the shopper will be asked if the payment details should be stored for future one-click payments. | [optional] 
**EnablePayOut** | **bool** | When true and &#x60;shopperReference&#x60; is provided, the payment details will be tokenized for payouts. | [optional] 
**EnableRealTimeUpdate** | **bool** | Choose if a specific transaction should use the Real-time Account Updater, regardless of other settings. | [optional] 
**EnableRecurring** | **bool** | When true and &#x60;shopperReference&#x60; is provided, the payment details will be tokenized for recurring payments. | [optional] 
**EntityType** | **string** | The type of the entity the payment is processed for. | [optional] 
**FraudOffset** | **int32** | An integer value that is added to the normal fraud score. The value can be either positive or negative. | [optional] 
**Installments** | [**Installments**](Installments.md) |  | [optional] 
**LineItems** | [**[]LineItem**](LineItem.md) | Price and product information about the purchased items, to be included on the invoice sent to the shopper. &gt; This field is required for Klarna, AfterPay, and RatePay. | [optional] 
**Mcc** | **string** | The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) (MCC) is a four-digit number, which relates to a particular market segment. This code reflects the predominant activity that is conducted by the merchant. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**MerchantOrderReference** | **string** | This reference allows linking multiple transactions to each other for reporting purposes (i.e. order auth-rate). The reference should be unique per billing cycle. The same merchant order reference should never be reused after the first authorised attempt. If used, this field should be supplied for all incoming authorisations. &gt; We strongly recommend you send the &#x60;merchantOrderReference&#x60; value to benefit from linking payment requests when authorisation retries take place. In addition, we recommend you provide &#x60;retry.orderAttemptNumber&#x60;, &#x60;retry.chainAttemptNumber&#x60;, and &#x60;retry.skipRetry&#x60; values in &#x60;PaymentRequest.additionalData&#x60;. | [optional] 
**Metadata** | [**map[string]interface{}**](.md) | Metadata consists of entries, each of which includes a key and a value. Limitations: Maximum 20 key-value pairs per request. When exceeding, the \&quot;177\&quot; error occurs: \&quot;Metadata size exceeds limit\&quot;. | [optional] 
**OrderReference** | **string** | When you are doing multiple partial (gift card) payments, this is the &#x60;pspReference&#x60; of the first payment. We use this to link the multiple payments to each other. As your own reference for linking multiple payments, use the &#x60;merchantOrderReference&#x60;instead. | [optional] 
**Origin** | **string** | Required for the Web integration.  Set this parameter to the origin URL of the page that you are loading the SDK from. | [optional] 
**RecurringExpiry** | **string** | Date after which no further authorisations shall be performed. Only for 3D Secure 2. | [optional] 
**RecurringFrequency** | **string** | Minimum number of days between authorisations. Only for 3D Secure 2. | [optional] 
**Reference** | **string** | The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\&quot;-\&quot;). Maximum length: 80 characters. | 
**ReturnUrl** | **string** | The URL to return to. | 
**RiskData** | [**RiskData**](RiskData.md) |  | [optional] 
**SdkVersion** | **string** | The version of the SDK you are using (for Web SDK integrations only). | [optional] 
**SessionValidity** | **string** | The maximum validity of the session. | [optional] 
**ShopperEmail** | **string** | The shopper&#39;s email address. We recommend that you provide this data, as it is used in velocity fraud checks. &gt; For 3D Secure 2 transactions, schemes require the &#x60;shopperEmail&#x60; for both &#x60;deviceChannel&#x60; **browser** and **app**. | [optional] 
**ShopperIP** | **string** | The shopper&#39;s IP address. In general, we recommend that you provide this data, as it is used in a number of risk checks (for instance, number of payment attempts or location-based checks). &gt; Required for 3D Secure 2 transactions. This field is also mandatory for some merchants depending on your business model. For more information, [contact Support](https://support.adyen.com/hc/en-us/requests/new). | [optional] 
**ShopperInteraction** | **string** | Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * &#x60;Ecommerce&#x60; - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * &#x60;ContAuth&#x60; - Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * &#x60;Moto&#x60; - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * &#x60;POS&#x60; - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal. | [optional] 
**ShopperLocale** | **string** | The combination of a language code and a country code to specify the language to be used in the payment. | [optional] 
**ShopperName** | [**Name**](Name.md) |  | [optional] 
**ShopperReference** | **string** | The shopper&#39;s reference to uniquely identify this shopper (e.g. user ID or account ID). &gt; This field is required for recurring payments. | [optional] 
**ShopperStatement** | **string** | The text to appear on the shopper&#39;s bank statement. | [optional] 
**SocialSecurityNumber** | **string** | The shopper&#39;s social security number. | [optional] 
**Splits** | [**[]Split**](Split.md) | The details of how the payment should be split when distributing a payment to a MarketPay Marketplace and its Accounts. | [optional] 
**StorePaymentMethod** | **bool** | When true and &#x60;shopperReference&#x60; is provided, the payment details will be stored. | [optional] 
**TelephoneNumber** | **string** | The shopper&#39;s telephone number. | [optional] 
**ThreeDSAuthenticationOnly** | **bool** | If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/checkout/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation. | [optional] 
**Token** | **string** | The token obtained when initializing the SDK.  &gt; This parameter is required for iOS and Android; not required for Web. | [optional] 
**TrustedShopper** | **bool** | Set to true if the payment should be routed to a trusted MID. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

