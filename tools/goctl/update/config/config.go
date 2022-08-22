package config

import "github.com/Holyson/test-go-zero-cors/core/logx"

// Config defines a service configure for goctl update
type Config struct {
	logx.LogConf
	ListenOn string
	FileDir  string
	FilePath string
}
