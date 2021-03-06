package main

import (
	"encoding/json"
	"io/ioutil"
)

type RemoteConfig struct {
	RemoteCryptoMethod string `json:"remoteCryptoMethod"`
	RemotePassword     string `json:"remotePassword"`
	RemoteServer       string `json:"remoteServer"`
}

type Config struct {
	PprofAddr           string         `json:"pprof"`
	HTTPProxyAddr       string         `json:"httpProxyAddr"`
	SOCKS4Addr          string         `json:"socks4Addr"`
	SOCKS5Addr          string         `json:"socks5Addr"`
	LocalCryptoMethod   string         `json:"localCryptoMethod"`
	LocalCryptoPassword string         `json:"localPassword"`
	DNSCacheTimeout     int            `json:dnsCacheTimeout`
	RemoteConfigs       []RemoteConfig `json:"remotes"`
}

func (c *Config) String() string {
	data, _ := json.Marshal(c)
	return string(data)
}

func LoadConfig(s string) (*Config, error) {
	data, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err = json.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
