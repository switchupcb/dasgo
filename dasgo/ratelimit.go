// Package dasgo provides Type Definitions for the Discord API.
package dasgo

// Rate Limit Headers
// https://discord.com/developers/docs/topics/rate-limits#header-format-rate-limit-header-examples
const (
	FlagRateLimitHeaderLimit      = "X-RateLimit-Limit"
	FlagRateLimitHeaderRemaining  = "X-RateLimit-Remaining"
	FlagRateLimitHeaderReset      = "X-RateLimit-Reset"
	FlagRateLimitHeaderResetAfter = "X-RateLimit-Reset-After"
	FlagRateLimitHeaderBucket     = "X-RateLimit-Bucket"
	FlagRateLimitHeaderGlobal     = "X-RateLimit-Global"
	FlagRateLimitHeaderScope      = "X-RateLimit-Scope"
)

// Rate Limit Header
// https://discord.com/developers/docs/topics/rate-limits#header-format
type RateLimitHeader struct {
	Limit      int     `http:"X-RateLimit-Limit,omitempty"`
	Remaining  int     `http:"X-RateLimit-Remaining,omitempty"`
	Reset      int64   `http:"X-RateLimit-Reset,omitempty"`
	ResetAfter float64 `http:"X-RateLimit-Reset-After,omitempty"`
	Bucket     string  `http:"X-RateLimit-Bucket,omitempty"`
	Global     bool    `http:"X-RateLimit-Global,omitempty"`
	Scope      string  `http:"X-RateLimit-Scope,omitempty"`
}

// Rate Limit Response Structure
// https://discord.com/developers/docs/topics/rate-limits#exceeding-a-rate-limit-rate-limit-response-structure
type RateLimitResponse struct {
	Message    string  `json:"message"`
	RetryAfter float64 `json:"retry_after"`
	Global     bool    `json:"global"`
}

// Global Rate Limit
// https://discord.com/developers/docs/topics/rate-limits#global-rate-limit
const (
	// 50 requests per second.
	FlagGlobalRequestRateLimit = 50
)

// Invalid Request Limit (CloudFlare Bans)
// https://discord.com/developers/docs/topics/rate-limits#invalid-request-limit-aka-cloudflare-bans
const (
	// 10,000 requests per 10 minutes.
	FlagInvalidRequestRateLimit = 10000
)

var (
	InvalidRateLimitRequests = map[int]string{
		FlagHTTPResponseCodeUNAUTHORIZED:    HTTPResponseCodes[FlagHTTPResponseCodeUNAUTHORIZED],
		FlagHTTPResponseCodeFORBIDDEN:       HTTPResponseCodes[FlagHTTPResponseCodeFORBIDDEN],
		FlagHTTPResponseCodeTOOMANYREQUESTS: HTTPResponseCodes[FlagHTTPResponseCodeTOOMANYREQUESTS],
	}
)
