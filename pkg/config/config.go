package config

type Config struct {
	ApiKey                string               `yaml:"api_key"`                 // your private serpAPI key
	ResultsPerPageLimits  ResultsPerPageLimits `yaml:"results_per_page_limits"` // max number of results per page allowed by each engine
	DefaultEnabledEngines EnabledEngines       `yaml:"default_enabled_engines"` // default engines to use if none are specified in the request
}

type ResultsPerPageLimits struct {
	Google     int `yaml:"google"`
	DuckDuckGo int `yaml:"duckduckgo"`
	Bing       int `yaml:"bing"`
	Yahoo      int `yaml:"yahoo"`
}

type EnabledEngines struct {
	Google     bool `yaml:"google"`
	DuckDuckGo bool `yaml:"duckduckgo"`
	Bing       bool `yaml:"bing"`
	Yahoo      bool `yaml:"yahoo"`
}
