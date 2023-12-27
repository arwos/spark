package tcp

import "time"

type Config struct {
	TCP TCPConfig `yaml:"tcp"`
}

type TCPConfig struct {
	Config  []TCPConfigItem `yaml:"config"`
	Timeout time.Duration   `yaml:"timeout,omitempty"`
}

type TCPConfigItem struct {
	Port  int                 `yaml:"port"`
	Certs []TCPConfigCertItem `yaml:"certs,omitempty"`
}

type TCPConfigCertItem struct {
	Public  string `yaml:"pub"`
	Private string `yaml:"priv"`
}
