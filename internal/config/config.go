package config

import (
	"time"

	"github.com/alecthomas/kong"
)

type Config struct {
	Debug bool `help:"Application mode" env:"DEBUG" required:"true"`

	Host        string `help:"Port to listen on"       env:"HOST"          default:":4000"`
	StatidDir   string `help:"Path to static assets"   env:"STATIC_DIR"    default:"ui/static"`
	DatabaseURL string `help:"Database connection URL" env:"DATABASE_URL"                       required:"true"`

	ClientOrigin string `env:"CLIENT_ORIGIN" required:"true"`

	// Redis
	RedisHost string `help:"Redis connection host" env:"REDIS_HOST" required:"true"`
	RedisPort string `help:"Redis port"            env:"REDIS_PORT" required:"true"`

	//MINIO
	// MinioPublicHost string `env:"MINIO_PUBLIC_HOST" required:"true"`
	MinioSecure    bool   `env:"MINIO_SECURE"      required:"true"`
	MinioHost      string `env:"MINIO_HOST"        required:"true"`
	MinioOrigin    string `env:"MINIO_ORIGIN"      required:"true"`
	MinioAccessKey string `env:"MINIO_ACCESS_KEY"  required:"true"`
	MinioSecretKey string `env:"MINIO_SECRET_KEY"  required:"true"`
	AppBucket      string `env:"MINIO_BUCKET"      required:"true"`

	ResetPasswordTokenExpiredIn time.Duration `env:"RESET_PASSWORD_TOKEN_EXPIRED_IN" required:"true"`

	// Token
	AccessTokenPrivateKey  string        `env:"ACCESS_TOKEN_PRIVATE_KEY"  required:"true"`
	AccessTokenPublicKey   string        `env:"ACCESS_TOKEN_PUBLIC_KEY"   required:"true"`
	RefreshTokenPrivateKey string        `env:"REFRESH_TOKEN_PRIVATE_KEY" required:"true"`
	RefreshTokenPublicKey  string        `env:"REFRESH_TOKEN_PUBLIC_KEY"  required:"true"`
	AccessTokenExpiresIn   time.Duration `env:"ACCESS_TOKEN_EXPIRED_IN"   required:"true"`
	RefreshTokenExpiresIn  time.Duration `env:"REFRESH_TOKEN_EXPIRED_IN"  required:"true"`
	AccessTokenMaxAge      int           `env:"ACCESS_TOKEN_MAXAGE"       required:"true"`
	RefreshTokenMaxAge     int           `env:"REFRESH_TOKEN_MAXAGE"      required:"true"`

	// Cookie
	CookieSecure bool   `env:"COOKIE_SECURE" required:"true"`
	CookieDomain string `env:"COOKIE_DOMAIN" required:"true"`

	//SMTP
	EmailFrom string `env:"EMAIL_FROM" required:"true"`
	SMTPHost  string `env:"SMTP_HOST"  required:"true"`
	SMTPPass  string `env:"SMTP_PASS"  required:"true"`
	SMTPPort  int    `env:"SMTP_PORT"  required:"true"`
	SMTPUser  string `env:"SMTP_USER"  required:"true"`
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
