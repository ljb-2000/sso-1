package config

const (
	// MySQL 配置
	MAX_POOL_SIZE        = 20
	DATABASE_DRIVER      = "mysql"
	DB_HOST              = "mysql.config:3306"
	DB_USER              = "root"
	DB_PASSWORD          = "root"
	DB_NAME              = "config"
	DB_CONNECTION_SUFFIX = "?parseTime=true"
	DELAY_MILLISECONDS   = 5000
)