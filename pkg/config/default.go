package config

func DefaultCfg() *Config {
	return &Config{
		ApiKey: "your-api-key",
		ResultsPerPageLimits: ResultsPerPageLimits{
			Google:     100,
			DuckDuckGo: 50,
			Bing:       50,
			Yahoo:      10, // SerpAPI has discontinued max results per page for Yahoo and it defaults to 10
		},
		DefaultEnabledEngines: EnabledEngines{
			Google:     true,
			DuckDuckGo: true,
			Bing:       true,
			Yahoo:      false,
		},
	}
}
