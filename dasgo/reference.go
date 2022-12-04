package dasgo

import "time"

// Version
// https://discord.com/developers/docs/reference#api-versioning
const (
	VersionDiscordAPI          = "10"
	VersionDiscordVoiceGateway = "?v=4"
)

// Timestamp Format
// https://discord.com/developers/docs/reference#iso8601-datetime
const (
	TimestampFormatISO8601 = time.RFC3339
)

// Image Formats
// https://discord.com/developers/docs/reference#image-formatting-image-formats
const (
	ImageFormatJPEG   = "JPEG"
	ImageFormatPNG    = "PNG"
	ImageFormatWebP   = "WebP"
	ImageFormatGIF    = "GIF"
	ImageFormatLottie = "Lottie"
)

// CDN Endpoint Exceptions
// https://discord.com/developers/docs/reference#image-formatting-cdn-endpoints
const (
	CDNEndpointAnimatedHashPrefix       = "a_"
	CDNEndpointUserDiscriminatorDivisor = 5
)

// Locales
// https://discord.com/developers/docs/reference#locales
const (
	FlagLocalesDanish              = "da"
	FlagLocalesGerman              = "de"
	FlagLocalesEnglishUK           = "en-GB"
	FlagLocalesEnglishUS           = "en-US"
	FlagLocalesSpanish             = "es-ES"
	FlagLocalesFrench              = "fr"
	FlagLocalesCroatian            = "hr"
	FlagLocalesItalian             = "it"
	FlagLocalesLithuanian          = "lt"
	FlagLocalesHungarian           = "hu"
	FlagLocalesDutch               = "nl"
	FlagLocalesNorwegian           = "no"
	FlagLocalesPolish              = "pl"
	FlagLocalesPortugueseBrazilian = "pt-BR"
	FlagLocalesRomanian            = "ro"
	FlagLocalesFinnish             = "fi"
	FlagLocalesSwedish             = "sv-SE"
	FlagLocalesVietnamese          = "vi"
	FlagLocalesTurkish             = "tr"
	FlagLocalesCzech               = "cs"
	FlagLocalesGreek               = "el"
	FlagLocalesBulgarian           = "bg"
	FlagLocalesRussian             = "ru"
	FlagLocalesUkrainian           = "uk"
	FlagLocalesHindi               = "hi"
	FlagLocalesThai                = "th"
	FlagLocalesChineseChina        = "zh-CN"
	FlagLocalesJapanese            = "ja"
	FlagLocalesChineseTaiwan       = "zh-TW"
	FlagLocalesKorean              = "ko"
)
