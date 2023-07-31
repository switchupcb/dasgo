package dasgo

import "encoding/json"

// Voice Payload Structure
// https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection
type VoicePayload struct {
	Op   int             `json:"op"`
	Data json.RawMessage `json:"d"`
}

// Voice Server Opcode (Event) Names
// https://discord.com/developers/docs/topics/opcodes-and-status-codes#voice-voice-opcodes
const (
	FlagVoiceOpcodeNameIdentify           = "IDENTIFY"
	FlagVoiceOpcodeNameSelectProtocol     = "SELECT_PROTOCOL"
	FlagVoiceOpcodeNameReady              = "READY"
	FlagVoiceOpcodeNameHeartbeat          = "HEARTBEAT"
	FlagVoiceOpcodeNameSessionDescription = "SESSION_DESCRIPTION"
	FlagVoiceOpcodeNameSpeaking           = "SPEAKING"
	FlagVoiceOpcodeNameResume             = "RESUME"
	FlagVoiceOpcodeNameHello              = "HELLO"
	FlagVoiceOpcodeNameResumed            = "RESUMED"
	FlagVoiceOpcodeNameClientDisconnect   = "CLIENT_DISCONNECT"
)

// Voice Server SendEvent Names
// https://discord.com/developers/docs/topics/opcodes-and-status-codes#voice-voice-opcodes
const (
	FlagVoiceSendEventNameIdentify       = "Identify"
	FlagVoiceSendEventNameSelectProtocol = "SelectProtocol "
	FlagVoiceSendEventNameHeartbeat      = "Heartbeat"
	FlagVoiceSendEventNameSpeaking       = "Speaking"
	FlagVoiceSendEventNameResume         = "Resume"
)

// Voice Identify Structure
// https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection-example-voice-identify-payload
type VoiceIdentify struct {
	ServerID  Snowflake `json:"server_id"`
	UserID    Snowflake `json:"user_id"`
	SessionID Snowflake `json:"session_id"`
	Token     string    `json:"token"`
}

// Select Protocol Structure
// https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-udp-connection-example-select-protocol-payload
type SelectProtocol struct {
	Protocol string             `json:"protocol"`
	Data     SelectProtocolData `json:"data"`
}

// Select Protocol Data Structure
// https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-udp-connection-example-select-protocol-payload
type SelectProtocolData struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Mode    string `json:"mode"`
}

// Voice Ready Structure
// https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection-example-voice-ready-payload
type VoiceReady struct {
	SSRC  int      `json:"ssrc"`
	IP    string   `json:"ip"`
	Port  int      `json:"port"`
	Modes []string `json:"modes"`
}

// Voice Heartbeat Structure
// https://discord.com/developers/docs/topics/voice-connections#heartbeating-example-heartbeat-payload
type VoiceHeartbeat struct {
	Data int64 `json:"d"`
}

// Session Description Structure
// https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-udp-connection-example-session-description-payload
type SessionDescription struct {
	Mode      string `json:"mode"`
	SecretKey []int  `json:"secret_key"`
}

// Speaking Structure
// https://discord.com/developers/docs/topics/voice-connections#speaking-example-speaking-payload
type Speaking struct {
	Speaking BitFlag `json:"speaking"`
	Delay    int     `json:"delay"`
	SSRC     int     `json:"ssrc"`
}

// Speaking Flags
// https://discord.com/developers/docs/topics/voice-connections#speaking
const (
	FlagSpeakingMicrophone BitFlag = 1 << 0
	FlagSpeakingSoundshare BitFlag = 1 << 1
	FlagSpeakingPriority   BitFlag = 1 << 2
)

// Voice HeartbeatACK Structure
// https://discord.com/developers/docs/topics/voice-connections#heartbeating-example-heartbeat-ack-payload
type VoiceHeartbeatACK struct {
	Data int64 `json:"d"`
}

// Voice Resume Structure
// https://discord.com/developers/docs/topics/voice-connections#resuming-voice-connection-example-resume-connection-payload
type VoiceResume struct {
	ServerID  Snowflake `json:"server_id"`
	SessionID Snowflake `json:"session_id"`
	Token     string    `json:"token"`
}

// Voice Hello Structure
// https://discord.com/developers/docs/topics/voice-connections#heartbeating-example-hello-payload-since-v3
type VoiceHello struct {
	HeartbeatInterval int64 `json:"heartbeat_interval"`
}

// Voice Resumed Structure
// https://discord.com/developers/docs/topics/voice-connections#resuming-voice-connection-example-resumed-payload
type VoiceResumed struct{}

// Client Disconnect Structure
type ClientDisconnect struct{}

// Voice Connection Encryption Modes
// https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-udp-connection-encryption-modes
const (
	FlagVoiceEncryptionModeNormal = "xsalsa20_poly1305"
	FlagVoiceEncryptionModeSuffix = "xsalsa20_poly1305_suffix"
	FlagVoiceEncryptionModeLite   = "xsalsa20_poly1305_lite"
)
