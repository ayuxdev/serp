package config

func DefaultCfg() *Config {
	return &Config{
		ApiKey: "your-api-key",
		Engines: map[string]EngineConfig{
			"google": {
				Enabled:               true,
				ResultsPerPageAllowed: 100,
				EngineName:            "google",
			},
			"duckduckgo": {
				Enabled:               true,
				ResultsPerPageAllowed: 50,
				EngineName:            "duckduckgo",
			},
			"bing": {
				Enabled:               true,
				ResultsPerPageAllowed: 50,
				EngineName:            "bing",
			},
			"yahoo": {
				Enabled:               false,
				ResultsPerPageAllowed: 10,
				EngineName:            "yahoo",
			},
		},
	}
}
