package http

import (
	"encoding/json"
)

type ResponseJsonError struct {
	Error string `json:"error,omitempty"`
}

type ResponseJsonOk struct {
	Data string `json:"data,omitempty"`
}

type ResponseCodes struct {
	code        int16
	description string
}

func (r *ResponseCodes) String() string {
	marshaledStruct, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
}

// GetHTTPResponseCode returns the description of a numeric HTTP code
func GetHTTPResponseCode(code int16) ResponseCodes {
	codes := map[int16]ResponseCodes{
		001: {001, "Unknown"},
		100: {100, "Continue"},
		101: {101, "Switching Protocols"},
		102: {102, "Processing"},
		200: {200, "OK"},
		201: {201, "Created"},
		202: {202, "Accepted"},
		203: {203, "Non-Authoritative Information"},
		204: {204, "No Content"},
		205: {205, "Reset Content"},
		206: {206, "Partial Content"},
		207: {207, "Multi-Status"},
		300: {300, "Multiple Choices"},
		301: {301, "Moved Permanently"},
		302: {302, "Moved Temporarily"},
		303: {303, "See Other"},
		304: {304, "Not Modified"},
		305: {305, "Use Proxy"},
		307: {307, "Temporary Redirect"},
		400: {400, "Bad Request"},
		401: {401, "Unauthorized"},
		402: {402, "Payment Required"},
		403: {403, "Forbidden"},
		404: {404, "Not Found"},
		405: {405, "Method Not Allowed"},
		406: {406, "Not Acceptable"},
		407: {407, "Proxy Authentication Required"},
		408: {408, "Request Time-out"},
		409: {409, "Conflict"},
		410: {410, "Gone"},
		411: {411, "Length Required"},
		412: {412, "Precondition Failed"},
		413: {413, "Request Entity Too Large"},
		414: {414, "Request-URI Too Large"},
		415: {415, "Unsupported Media Type"},
		416: {416, "Requested Range Not Satisfiable"},
		417: {417, "Expectation Failed"},
		418: {418, "I\"m a teapot"},
		422: {422, "Unprocessable Entity"},
		423: {423, "Locked"},
		424: {424, "Failed Dependency"},
		425: {425, "Unordered Collection"},
		426: {426, "Upgrade Required"},
		428: {428, "Precondition Required"},
		429: {429, "Too Many Requests"},
		431: {431, "Request Header Fields Too Large"},
		451: {451, "Unavailable For Legal Reasons"},
		500: {500, "Internal Server Error"},
		501: {501, "Not Implemented"},
		502: {502, "Bad Gateway"},
		503: {503, "Service Unavailable"},
		504: {504, "Gateway Time-out"},
		505: {505, "HTTP Version Not Supported"},
		506: {506, "Variant Also Negotiates"},
		507: {507, "Insufficient Storage"},
		509: {509, "Bandwidth Limit Exceeded"},
		510: {510, "Not Extended"},
		511: {511, "Network Authentication Required"},
		718: {718, "The Bronx"},
	}

	if val, ok := codes[code]; ok {
		return val
	}

	return codes[001]
}
