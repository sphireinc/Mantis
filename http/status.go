package http

type Status struct {
	Continue                              int64 // 100 RFC 7231, 6.2.1
	SwitchingProtocols                    int64 // 101 - RFC 7231, 6.2.2
	Processing                            int64 // 102 - RFC 2518, 10.1
	EarlyHints                            int64 // 103 - RFC 8297
	OK                                    int64 // 200 - RFC 7231, 6.3.1
	Created                               int64 // 201 - RFC 7231, 6.3.2
	Accepted                              int64 // 202 - RFC 7231, 6.3.3
	NonAuthoritativeInfo                  int64 // 203 - RFC 7231, 6.3.4
	NoContent                             int64 // 204 - RFC 7231, 6.3.5
	ResetContent                          int64 // 205 - RFC 7231, 6.3.6
	PartialContent                        int64 // 206 - RFC 7233, 4.1
	MultiStatus                           int64 // 207 - RFC 4918, 11.1
	AlreadyReported                       int64 // 208 - RFC 5842, 7.1
	IMUsed                                int64 // 226 - RFC 3229, 10.4.1
	MultipleChoices                       int64 // 300 - RFC 7231, 6.4.1
	MovedPermanently                      int64 // 301 - RFC 7231, 6.4.2
	Found                                 int64 // 302 - RFC 7231, 6.4.3
	SeeOther                              int64 // 303 - RFC 7231, 6.4.4
	NotModified                           int64 // 304 - RFC 7232, 4.1
	UseProxy                              int64 // 305 - RFC 7231, 6.4.5
	SwitchProxy                           int64 // 306 - RFC 7231, 6.4.6 (Unused)
	TemporaryRedirect                     int64 // 307 - RFC 7231, 6.4.7
	PermanentRedirect                     int64 // 308 - RFC 7538, 3
	BadRequest                            int64 // 400 - RFC 7231, 6.5.1
	Unauthorized                          int64 // 401 - RFC 7235, 3.1
	PaymentRequired                       int64 // 402 - RFC 7231, 6.5.2
	Forbidden                             int64 // 403 - RFC 7231, 6.5.3
	NotFound                              int64 // 404 - RFC 7231, 6.5.4
	MethodNotAllowed                      int64 // 405 - RFC 7231, 6.5.5
	NotAcceptable                         int64 // 406 - RFC 7231, 6.5.6
	ProxyAuthRequired                     int64 // 407 - RFC 7235, 3.2
	RequestTimeout                        int64 // 408 - RFC 7231, 6.5.7
	Conflict                              int64 // 409 - RFC 7231, 6.5.8
	Gone                                  int64 // 410 - RFC 7231, 6.5.9
	LengthRequired                        int64 // 411 - RFC 7231, 6.5.10
	PreconditionFailed                    int64 // 412 - RFC 7232, 4.2
	RequestEntityTooLarge                 int64 // 413 - RFC 7231, 6.5.11
	RequestURITooLong                     int64 // 414 - RFC 7231, 6.5.12
	UnsupportedMediaType                  int64 // 415 - RFC 7231, 6.5.13
	RequestedRangeNotSatisfiable          int64 // 416 - RFC 7233, 4.4
	ExpectationFailed                     int64 // 417 - RFC 7231, 6.5.14
	Teapot                                int64 // 418 - RFC 7168, 2.3.3
	MisdirectedRequest                    int64 // 421 - RFC 7540, 9.1.2
	UnprocessableEntity                   int64 // 422 - RFC 4918, 11.2
	Locked                                int64 // 423 - RFC 4918, 11.3
	FailedDependency                      int64 // 424 - RFC 4918, 11.4
	TooEarly                              int64 // 425 - RFC 8470, 5.2.
	UpgradeRequired                       int64 // 426 - RFC 7231, 6.5.15
	PreconditionRequired                  int64 // 428 - RFC 6585, 3
	TooManyRequests                       int64 // 429 - RFC 6585, 4
	RequestHeaderFieldsTooLarge           int64 // 431 - RFC 6585, 5
	UnavailableForLegalReasons            int64 // 451 - RFC 7725, 3
	InternalServerError                   int64 // 500 // RFC 7231, 6.6.1
	NotImplemented                        int64 // 501 // RFC 7231, 6.6.2
	BadGateway                            int64 // 502 // RFC 7231, 6.6.3
	ServiceUnavailable                    int64 // 503 // RFC 7231, 6.6.4
	GatewayTimeout                        int64 // 504 // RFC 7231, 6.6.5
	HTTPVersionNotSupported               int64 // 505 // RFC 7231, 6.6.6
	VariantAlsoNegotiates                 int64 // 506 // RFC 2295, 8.1
	InsufficientStorage                   int64 // 507 // RFC 4918, 11.5
	LoopDetected                          int64 // 508 // RFC 5842, 7.2
	NotExtended                           int64 // 510 // RFC 2774, 7
	NetworkAuthenticationRequired         int64 // 511 // RFC 6585, 6
	PageExpired                           int64 // 419 - Laravel Framework
	MethodFailure                         int64 // 420 - Spring Framework
	BlockedByWindowsParentalControls      int64 // 450 - Microsoft
	InvalidTokenEsri                      int64 // 498 - ArcGIS
	TokenRequiredEsri                     int64 // 499 - ArcGIS
	BandwidthLimitExceeded                int64 // 509 - Apache Web Server / cPanel
	SiteIsOverloaded                      int64 // 529 - Qualys SSLLabs
	SiteIsFrozen                          int64 // 530 - Pantheon
	NetworkReadTimeout                    int64 // 598 - Proxies
	NetworkConnectTimeoutError            int64 // 599 - Proxies
	LoginTimeout                          int64 // 440 - IIS
	RetryWith                             int64 // 449 - IIS
	Redirect                              int64 // 451 - IIS
	NoResponse                            int64 // 444 - Nginx
	RequestHeaderTooLarge                 int64 // 494 - Nginx
	SSLCertificateError                   int64 // 495 - Nginx
	HTTPRequestSentToHTTPSPort            int64 // 497 - Nginx
	ClientClosedRequest                   int64 // 499 - Nginx
	WebServerReturnedAnUnknownError       int64 // 520 - Cloudflare
	WebServerIsDown                       int64 // 521 - Cloudflare
	ConnectionTimedOut                    int64 // 522 - Cloudflare
	OriginIsUnreachable                   int64 // 523 - Cloudflare
	ATimeoutOccurred                      int64 // 524 - Cloudflare
	SSLHandshakeFailed                    int64 // 525 - Cloudflare
	InvalidSSLCertificate                 int64 // 526 - Cloudflare
	RailgunError                          int64 // 527 - Cloudflare
	AWSClientClosedConnectionWithLB       int64 // 460 - AWS Elastic Load Balancer
	AWSLBReceivedXForwardForHeaderMT30    int64 // 463 - AWS Elastic Load Balancer
	AWSUnauthorized                       int64 // 561 - AWS Elastic Load Balancer
	CachingResponseIsStale                int64 // 110 - Caching
	CachingRevalidationFailed             int64 // 111 - Caching
	CachingDisconnectedOperation          int64 // 112 - Caching
	CachingHeuristicExpiration            int64 // 113 - Caching
	CachingMiscellaneousWarning           int64 // 199 - Caching
	CachingTransformationApplied          int64 // 214 - Caching
	CachingMiscellaneousPersistentWarning int64 // 299
}

func (s *Status) Fill() {
	s.Continue = 100                           // RFC 7231, 6.2.1
	s.SwitchingProtocols = 101                 // RFC 7231, 6.2.2
	s.Processing = 102                         // RFC 2518, 10.1
	s.EarlyHints = 103                         // RFC 8297
	s.OK = 200                                 // RFC 7231, 6.3.1
	s.Created = 201                            // RFC 7231, 6.3.2
	s.Accepted = 202                           // RFC 7231, 6.3.3
	s.NonAuthoritativeInfo = 203               // RFC 7231, 6.3.4
	s.NoContent = 204                          // RFC 7231, 6.3.5
	s.ResetContent = 205                       // RFC 7231, 6.3.6
	s.PartialContent = 206                     // RFC 7233, 4.1
	s.MultiStatus = 207                        // RFC 4918, 11.1
	s.AlreadyReported = 208                    // RFC 5842, 7.1
	s.IMUsed = 226                             // RFC 3229, 10.4.1
	s.MultipleChoices = 300                    // RFC 7231, 6.4.1
	s.MovedPermanently = 301                   // RFC 7231, 6.4.2
	s.Found = 302                              // RFC 7231, 6.4.3
	s.SeeOther = 303                           // RFC 7231, 6.4.4
	s.NotModified = 304                        // RFC 7232, 4.1
	s.UseProxy = 305                           // RFC 7231, 6.4.5
	s.SwitchProxy = 306                        // RFC 7231, 6.4.6 (Unused)
	s.TemporaryRedirect = 307                  // RFC 7231, 6.4.7
	s.PermanentRedirect = 308                  // RFC 7538, 3
	s.BadRequest = 400                         // RFC 7231, 6.5.1
	s.Unauthorized = 401                       // RFC 7235, 3.1
	s.PaymentRequired = 402                    // RFC 7231, 6.5.2
	s.Forbidden = 403                          // RFC 7231, 6.5.3
	s.NotFound = 404                           // RFC 7231, 6.5.4
	s.MethodNotAllowed = 405                   // RFC 7231, 6.5.5
	s.NotAcceptable = 406                      // RFC 7231, 6.5.6
	s.ProxyAuthRequired = 407                  // RFC 7235, 3.2
	s.RequestTimeout = 408                     // RFC 7231, 6.5.7
	s.Conflict = 409                           // RFC 7231, 6.5.8
	s.Gone = 410                               // RFC 7231, 6.5.9
	s.LengthRequired = 411                     // RFC 7231, 6.5.10
	s.PreconditionFailed = 412                 // RFC 7232, 4.2
	s.RequestEntityTooLarge = 413              // RFC 7231, 6.5.11
	s.RequestURITooLong = 414                  // RFC 7231, 6.5.12
	s.UnsupportedMediaType = 415               // RFC 7231, 6.5.13
	s.RequestedRangeNotSatisfiable = 416       // RFC 7233, 4.4
	s.ExpectationFailed = 417                  // RFC 7231, 6.5.14
	s.Teapot = 418                             // RFC 7168, 2.3.3
	s.MisdirectedRequest = 421                 // RFC 7540, 9.1.2
	s.UnprocessableEntity = 422                // RFC 4918, 11.2
	s.Locked = 423                             // RFC 4918, 11.3
	s.FailedDependency = 424                   // RFC 4918, 11.4
	s.TooEarly = 425                           // RFC 8470, 5.2.
	s.UpgradeRequired = 426                    // RFC 7231, 6.5.15
	s.PreconditionRequired = 428               // RFC 6585, 3
	s.TooManyRequests = 429                    // RFC 6585, 4
	s.RequestHeaderFieldsTooLarge = 431        // RFC 6585, 5
	s.UnavailableForLegalReasons = 451         // RFC 7725, 3
	s.InternalServerError = 500                /// RFC 7231, 6.6.1
	s.NotImplemented = 501                     /// RFC 7231, 6.6.2
	s.BadGateway = 502                         /// RFC 7231, 6.6.3
	s.ServiceUnavailable = 503                 /// RFC 7231, 6.6.4
	s.GatewayTimeout = 504                     /// RFC 7231, 6.6.5
	s.HTTPVersionNotSupported = 505            /// RFC 7231, 6.6.6
	s.VariantAlsoNegotiates = 506              /// RFC 2295, 8.1
	s.InsufficientStorage = 507                /// RFC 4918, 11.5
	s.LoopDetected = 508                       /// RFC 5842, 7.2
	s.NotExtended = 510                        /// RFC 2774, 7
	s.NetworkAuthenticationRequired = 511      /// RFC 6585, 6
	s.PageExpired = 419                        // Laravel Framework
	s.MethodFailure = 420                      // Spring Framework
	s.BlockedByWindowsParentalControls = 450   // Microsoft
	s.InvalidTokenEsri = 498                   // ArcGIS
	s.TokenRequiredEsri = 499                  // ArcGIS
	s.BandwidthLimitExceeded = 509             // Apache Web Server / cPanel
	s.SiteIsOverloaded = 529                   // Qualys SSLLabs
	s.SiteIsFrozen = 530                       // Pantheon
	s.NetworkReadTimeout = 598                 // Proxies
	s.NetworkConnectTimeoutError = 599         // Proxies
	s.LoginTimeout = 440                       // IIS
	s.RetryWith = 449                          // IIS
	s.Redirect = 451                           // IIS
	s.NoResponse = 444                         // Nginx
	s.RequestHeaderTooLarge = 494              // Nginx
	s.SSLCertificateError = 495                // Nginx
	s.HTTPRequestSentToHTTPSPort = 497         // Nginx
	s.ClientClosedRequest = 499                // Nginx
	s.WebServerReturnedAnUnknownError = 520    // Cloudflare
	s.WebServerIsDown = 521                    // Cloudflare
	s.ConnectionTimedOut = 522                 // Cloudflare
	s.OriginIsUnreachable = 523                // Cloudflare
	s.ATimeoutOccurred = 524                   // Cloudflare
	s.SSLHandshakeFailed = 525                 // Cloudflare
	s.InvalidSSLCertificate = 526              // Cloudflare
	s.RailgunError = 527                       // Cloudflare
	s.AWSClientClosedConnectionWithLB = 460    // AWS Elastic Load Balancer
	s.AWSLBReceivedXForwardForHeaderMT30 = 463 // AWS Elastic Load Balancer
	s.AWSUnauthorized = 561                    // AWS Elastic Load Balancer
	s.CachingResponseIsStale = 110             // Caching
	s.CachingRevalidationFailed = 111          // Caching
	s.CachingDisconnectedOperation = 112       // Caching
	s.CachingHeuristicExpiration = 113         // Caching
	s.CachingMiscellaneousWarning = 199        // Caching
	s.CachingTransformationApplied = 214       // Caching
	s.CachingMiscellaneousPersistentWarning = 299
}

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
