package config

import (
	"github.com/alecthomas/kong"
)

type Config struct {
	Host        string `help:"Port to listen on"        env:"HOST"          default:":4000"`
	StatidDir   string `help:"Path to static assets"    env:"STATIC_DIR"    default:"ui/static"`
	DatabaseURL string `help:"Database connection URL"  env:"DATABASE_URL"                             required:"true"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	parser, err := kong.New(cfg)
	if err != nil {
		return nil, err
	}

	// Parse command-line flags, environment variables, and config file
	_, err = parser.Parse(nil)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
