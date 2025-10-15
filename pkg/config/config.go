package config

type Config struct {
	ApiKey  string                  `yaml:"api_key"`
	Engines map[string]EngineConfig `yaml:"engines"`
}

type EngineConfig struct {
	Enabled               bool   `yaml:"enabled"`
	ResultsPerPageAllowed int    `yaml:"results_per_page_allowed"` // max number of results per page allowed by the engine
	EngineName            string `yaml:"engine_name"`
}
