// Package dasgo provides Type Definitions for the Discord API.
package dasgo

import "time"

// Rate Limit Headers
// https://discord.com/developers/docs/topics/rate-limits#header-format-rate-limit-header-examples
const (
	FlagRateLimitHeaderDate       = "Date"
	FlagRateLimitHeaderLimit      = "X-RateLimit-Limit"
	FlagRateLimitHeaderRemaining  = "X-RateLimit-Remaining"
	FlagRateLimitHeaderReset      = "X-RateLimit-Reset"
	FlagRateLimitHeaderResetAfter = "X-RateLimit-Reset-After"
	FlagRateLimitHeaderBucket     = "X-RateLimit-Bucket"
	FlagRateLimitHeaderGlobal     = "X-RateLimit-Global"
	FlagRateLimitHeaderScope      = "X-RateLimit-Scope"
	FlagRateLimitHeaderRetryAfter = "Retry-After"
)

// Rate Limit Header
// https://discord.com/developers/docs/topics/rate-limits#header-format
type RateLimitHeader struct {
	Limit      int     `http:"X-RateLimit-Limit,omitempty"`
	Remaining  int     `http:"X-RateLimit-Remaining,omitempty"`
	Reset      float64 `http:"X-RateLimit-Reset,omitempty"`
	ResetAfter float64 `http:"X-RateLimit-Reset-After,omitempty"`
	Bucket     string  `http:"X-RateLimit-Bucket,omitempty"`
	Global     bool    `http:"X-RateLimit-Global,omitempty"`
	Scope      string  `http:"X-RateLimit-Scope,omitempty"`
}

// Rate Limit Scope Values
// https://discord.com/developers/docs/topics/rate-limits#header-format-rate-limit-header-examples
const (
	RateLimitScopeValueUser   = "user"
	RateLimitScopeValueGlobal = "global"
	RateLimitScopeValueShared = "shared"
)

// Rate Limit Response Structure
// https://discord.com/developers/docs/topics/rate-limits#exceeding-a-rate-limit-rate-limit-response-structure
type RateLimitResponse struct {
	Message    string  `json:"message"`
	RetryAfter float64 `json:"retry_after"`
	Global     bool    `json:"global"`
}

// Global Rate Limits
// https://discord.com/developers/docs/topics/rate-limits#global-rate-limit
const (
	// Global Rate Limit (Requests): 50 requests per second.
	FlagGlobalRateLimitRequest = 50

	// Global Rate Limit (Gateway): 120 commands per minute.
	FlagGlobalRateLimitGateway         = 120
	FlagGlobalRateLimitGatewayInterval = time.Minute

	// Global Rate Limit (Identify Command): Get Gateway Bot `max_concurrency + 1` per 5 Seconds.
	FlagGlobalRateLimitIdentifyInterval = time.Second * 5

	// Global Rate Limit (Identify Command): 1000 per day.
	FlagGlobalRateLimitIdentifyDaily         = 1000
	FlagGlobalRateLimitIdentifyDailyInterval = time.Hour * 24
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
