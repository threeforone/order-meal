package global

var (
	Conf   *Config
	OrderT *OrderTime
)

type Config struct {
	App   `mapstructure:"app"`
	Mysql `mapstructure:"mysql"`
	JWT   `mapstructure:"jwt"`
}

type App struct {
	Port int
}

type Mysql struct {
	Dsn          string `mapstructure:"dsn"`
	MaxOpenConns int    `mapstructure:"MaxOpenConns"`
	MaxIdleConns int    `mapstructure:"MaxIdleConns"`
}

type OrderTime struct {
	StartTime string `mapstructure:"start-time" json:"startTime"`
	EndTime   string `mapstructure:"end-time" json:"endTime"`
}

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" `
	ExpiresTime string `mapstructure:"expires-time"`
	BufferTime  string `mapstructure:"buffer-time"`
	Issuer      string `mapstructure:"issuer"`
}
