package config

import "time"

var (
	AccessSignKey  = []byte("ssecca")
	RefreshSignKey = []byte("hserfer")
)

const (
	AccessTokenExpireTime  = time.Minute * 120
	RefreshTokenExpireTime = time.Minute * 240
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)
