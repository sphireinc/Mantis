package http

// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
const (
	StatusContinue           = 100 // RFC 7231, 6.2.1
	StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
	StatusProcessing         = 102 // RFC 2518, 10.1
	StatusEarlyHints         = 103 // RFC 8297

	StatusOK                   = 200 // RFC 7231, 6.3.1
	StatusCreated              = 201 // RFC 7231, 6.3.2
	StatusAccepted             = 202 // RFC 7231, 6.3.3
	StatusNonAuthoritativeInfo = 203 // RFC 7231, 6.3.4
	StatusNoContent            = 204 // RFC 7231, 6.3.5
	StatusResetContent         = 205 // RFC 7231, 6.3.6
	StatusPartialContent       = 206 // RFC 7233, 4.1
	StatusMultiStatus          = 207 // RFC 4918, 11.1
	StatusAlreadyReported      = 208 // RFC 5842, 7.1
	StatusIMUsed               = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices   = 300 // RFC 7231, 6.4.1
	StatusMovedPermanently  = 301 // RFC 7231, 6.4.2
	StatusFound             = 302 // RFC 7231, 6.4.3
	StatusSeeOther          = 303 // RFC 7231, 6.4.4
	StatusNotModified       = 304 // RFC 7232, 4.1
	StatusUseProxy          = 305 // RFC 7231, 6.4.5
	StatusSwitchProxy       = 306 // RFC 7231, 6.4.6 (Unused)
	StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7
	StatusPermanentRedirect = 308 // RFC 7538, 3

	StatusBadRequest                   = 400 // RFC 7231, 6.5.1
	StatusUnauthorized                 = 401 // RFC 7235, 3.1
	StatusPaymentRequired              = 402 // RFC 7231, 6.5.2
	StatusForbidden                    = 403 // RFC 7231, 6.5.3
	StatusNotFound                     = 404 // RFC 7231, 6.5.4
	StatusMethodNotAllowed             = 405 // RFC 7231, 6.5.5
	StatusNotAcceptable                = 406 // RFC 7231, 6.5.6
	StatusProxyAuthRequired            = 407 // RFC 7235, 3.2
	StatusRequestTimeout               = 408 // RFC 7231, 6.5.7
	StatusConflict                     = 409 // RFC 7231, 6.5.8
	StatusGone                         = 410 // RFC 7231, 6.5.9
	StatusLengthRequired               = 411 // RFC 7231, 6.5.10
	StatusPreconditionFailed           = 412 // RFC 7232, 4.2
	StatusRequestEntityTooLarge        = 413 // RFC 7231, 6.5.11
	StatusRequestURITooLong            = 414 // RFC 7231, 6.5.12
	StatusUnsupportedMediaType         = 415 // RFC 7231, 6.5.13
	StatusRequestedRangeNotSatisfiable = 416 // RFC 7233, 4.4
	StatusExpectationFailed            = 417 // RFC 7231, 6.5.14
	StatusTeapot                       = 418 // RFC 7168, 2.3.3
	StatusMisdirectedRequest           = 421 // RFC 7540, 9.1.2
	StatusUnprocessableEntity          = 422 // RFC 4918, 11.2
	StatusLocked                       = 423 // RFC 4918, 11.3
	StatusFailedDependency             = 424 // RFC 4918, 11.4
	StatusTooEarly                     = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              = 426 // RFC 7231, 6.5.15
	StatusPreconditionRequired         = 428 // RFC 6585, 3
	StatusTooManyRequests              = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   = 451 // RFC 7725, 3

	StatusInternalServerError           = 500 // RFC 7231, 6.6.1
	StatusNotImplemented                = 501 // RFC 7231, 6.6.2
	StatusBadGateway                    = 502 // RFC 7231, 6.6.3
	StatusServiceUnavailable            = 503 // RFC 7231, 6.6.4
	StatusGatewayTimeout                = 504 // RFC 7231, 6.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 7231, 6.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6

	StatusPageExpired                      = 419 // Laravel Framework
	StatusMethodFailure                    = 420 // Spring Framework
	StatusBlockedByWindowsParentalControls = 450 // Microsoft
	StatusInvalidTokenEsri                 = 498 // ArcGIS
	StatusTokenRequiredEsri                = 499 // ArcGIS
	StatusBandwidthLimitExceeded           = 509 // Apache Web Server / cPanel
	StatusSiteIsOverloaded                 = 529 // Qualys SSLLabs
	StatusSiteIsFrozen                     = 530 // Pantheon
	StatusNetworkReadTimeout               = 598 // Proxies
	StatusNetworkConnectTimeoutError       = 599 // Proxies
	StatusLoginTimeout                     = 440 // IIS
	StatusRetryWith                        = 449 // IIS
	StatusRedirect                         = 451 // IIS
	StatusNoResponse                       = 444 // Nginx
	StatusRequestHeaderTooLarge            = 494 // Nginx
	StatusSSLCertificateError              = 495 // Nginx
	StatusHTTPRequestSentToHTTPSPort       = 497 // Nginx
	StatusClientClosedRequest              = 499 // Nginx

	StatusWebServerReturnedAnUnknownError = 520 // Cloudflare
	StatusWebServerIsDown                 = 521 // Cloudflare
	StatusConnectionTimedOut              = 522 // Cloudflare
	StatusOriginIsUnreachable             = 523 // Cloudflare
	StatusATimeoutOccurred                = 524 // Cloudflare
	StatusSSLHandshakeFailed              = 525 // Cloudflare
	StatusInvalidSSLCertificate           = 526 // Cloudflare
	StatusRailgunError                    = 527 // Cloudflare

	StatusAWSClientClosedConnectionWithLB    = 460 // AWS Elastic Load Balancer
	StatusAWSLBReceivedXForwardForHeaderMT30 = 463 // AWS Elastic Load Balancer
	StatusAWSUnauthorized                    = 561 // AWS Elastic Load Balancer

	StatusCachingResponseIsStale                = 110 // Caching
	StatusCachingRevalidationFailed             = 111 // Caching
	StatusCachingDisconnectedOperation          = 112 // Caching
	StatusCachingHeuristicExpiration            = 113 // Caching
	StatusCachingMiscellaneousWarning           = 199 // Caching
	StatusCachingTransformationApplied          = 214 // Caching
	StatusCachingMiscellaneousPersistentWarning = 299 // Caching

)

var statusText = map[int]string{
	StatusContinue:           "Continue",
	StatusSwitchingProtocols: "Switching Protocols",
	StatusProcessing:         "Processing",
	StatusEarlyHints:         "Early Hints",

	StatusOK:                   "OK",
	StatusCreated:              "Created",
	StatusAccepted:             "Accepted",
	StatusNonAuthoritativeInfo: "Non-Authoritative Information",
	StatusNoContent:            "No Content",
	StatusResetContent:         "Reset Content",
	StatusPartialContent:       "Partial Content",
	StatusMultiStatus:          "Multi-Status",
	StatusAlreadyReported:      "Already Reported",
	StatusIMUsed:               "IM Used",

	StatusMultipleChoices:   "Multiple Choices",
	StatusMovedPermanently:  "Moved Permanently",
	StatusFound:             "Found",
	StatusSeeOther:          "See Other",
	StatusNotModified:       "Not Modified",
	StatusUseProxy:          "Use Proxy",
	StatusSwitchProxy:       "Switch Proxy",
	StatusTemporaryRedirect: "Temporary Redirect",
	StatusPermanentRedirect: "Permanent Redirect",

	StatusBadRequest:                   "Bad Request",
	StatusUnauthorized:                 "Unauthorized",
	StatusPaymentRequired:              "Payment Required",
	StatusForbidden:                    "Forbidden",
	StatusNotFound:                     "Not Found",
	StatusMethodNotAllowed:             "Method Not Allowed",
	StatusNotAcceptable:                "Not Acceptable",
	StatusProxyAuthRequired:            "Proxy Authentication Required",
	StatusRequestTimeout:               "Request Timeout",
	StatusConflict:                     "Conflict",
	StatusGone:                         "Gone",
	StatusLengthRequired:               "Length Required",
	StatusPreconditionFailed:           "Precondition Failed",
	StatusRequestEntityTooLarge:        "Request Entity Too Large",
	StatusRequestURITooLong:            "Request URI Too Long",
	StatusUnsupportedMediaType:         "Unsupported Media Type",
	StatusRequestedRangeNotSatisfiable: "Requested Range Not Satisfiable",
	StatusExpectationFailed:            "Expectation Failed",
	StatusTeapot:                       "I'm a teapot",
	StatusMisdirectedRequest:           "Misdirected Request",
	StatusUnprocessableEntity:          "Unprocessable Entity",
	StatusLocked:                       "Locked",
	StatusFailedDependency:             "Failed Dependency",
	StatusTooEarly:                     "Too Early",
	StatusUpgradeRequired:              "Upgrade Required",
	StatusPreconditionRequired:         "Precondition Required",
	StatusTooManyRequests:              "Too Many Requests",
	StatusRequestHeaderFieldsTooLarge:  "Request Header Fields Too Large",
	StatusUnavailableForLegalReasons:   "Unavailable For Legal Reasons",

	StatusInternalServerError:           "Internal Server Error",
	StatusNotImplemented:                "Not Implemented",
	StatusBadGateway:                    "Bad Gateway",
	StatusServiceUnavailable:            "Service Unavailable",
	StatusGatewayTimeout:                "Gateway Timeout",
	StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
	StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	StatusInsufficientStorage:           "Insufficient Storage",
	StatusLoopDetected:                  "Loop Detected",
	StatusNotExtended:                   "Not Extended",
	StatusNetworkAuthenticationRequired: "Network Authentication Required",

	StatusPageExpired:                      "Page Expired",
	StatusMethodFailure:                    "Method Failure",
	StatusBlockedByWindowsParentalControls: "Blocked By Windows Parental Controls",
	StatusInvalidTokenEsri:                 "Invalid Token Esri",
	StatusBandwidthLimitExceeded:           "Bandwidth Limit Exceeded",
	StatusSiteIsOverloaded:                 "Site Is Overloaded",
	StatusSiteIsFrozen:                     "Site Is Frozen",
	StatusNetworkReadTimeout:               "Network Read Timeout",
	StatusNetworkConnectTimeoutError:       "NetworkConnectTimeoutError",
	StatusLoginTimeout:                     "LoginTimeout",
	StatusRetryWith:                        "RetryWith",
	StatusNoResponse:                       "NoResponse",
	StatusRequestHeaderTooLarge:            "RequestHeaderTooLarge ",
	StatusSSLCertificateError:              "SSLCertificateError",
	StatusHTTPRequestSentToHTTPSPort:       "HTTPRequestSentToHTTPSPort",
	StatusClientClosedRequest:              "ClientClosedRequest",

	StatusWebServerReturnedAnUnknownError: "Web Server Returned An Unknown Error",
	StatusWebServerIsDown:                 "Web Server Is Down",
	StatusConnectionTimedOut:              "Connection Timed Out",
	StatusOriginIsUnreachable:             "Origin Is Unreachable",
	StatusATimeoutOccurred:                "A Timeout Occurred",
	StatusSSLHandshakeFailed:              "SSL Handshake Failed",
	StatusInvalidSSLCertificate:           "Invalid SSL Certificate",
	StatusRailgunError:                    "Railgun Error",

	StatusAWSClientClosedConnectionWithLB:    "AWS Client Closed Connection With LB",
	StatusAWSLBReceivedXForwardForHeaderMT30: "AWS LB Received-X-Forward For Header More Than 30 IP",
	StatusAWSUnauthorized:                    "AWS Unauthorized",

	StatusCachingResponseIsStale:                "Caching Response Is Stale",
	StatusCachingRevalidationFailed:             "Caching Revalidation Failed",
	StatusCachingDisconnectedOperation:          "Caching Disconnected Operation",
	StatusCachingHeuristicExpiration:            "Caching Heuristic Expiration",
	StatusCachingMiscellaneousWarning:           "Caching Miscellaneous Warning",
	StatusCachingTransformationApplied:          "Caching Transformation Applied",
	StatusCachingMiscellaneousPersistentWarning: "Caching Miscellaneous Persistent Warning",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
