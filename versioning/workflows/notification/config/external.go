package config

type ExternalConfig struct {
	TemporalHost      string `mapstructure:"TEMPORAL_HOST"`
	TemporalNamespace string `mapstructure:"TEMPORAL_NAMESPACE"`
}
