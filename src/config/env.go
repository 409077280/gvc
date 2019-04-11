package config

// 环境配置文件
// 可配置多个环境配置，进行切换

type Env struct {
	DEBUG             bool

	API_SERVER_PORT       string
	STATIC_SERVER_PORT    string
	REDIRECT_SERVER_PORT  string

	ALLOW_ORIGINS string

	DB_IP       string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_SSLMODEL string

	REDIS_IP          string
	REDIS_PORT        string
	REDIS_PASSWORD    string
	REDIS_DB          int
	REDIS_SESSION_DB  int
	REDIS_CACHE_DB    int
	APP_SECRET        string

	ACCESS_LOG      bool
	ACCESS_LOG_PATH string
	ERROR_LOG       bool
	ERROR_LOG_PATH  string
	INFO_LOG        bool
	INFO_LOG_PATH   string

	SQL_LOG bool

	TEMPLATE_PATH string // 静态文件相对路径
}

var env = Env{
	DEBUG: true,

	API_SERVER_PORT:       ":8000",
	STATIC_SERVER_PORT:    ":443",
	REDIRECT_SERVER_PORT:  ":80",

	ALLOW_ORIGINS : "https://localhost",

	DB_IP:       "47.100.197.243",
	DB_PORT:     "5432",
	DB_USERNAME: "postgres",
	DB_PASSWORD: "$qwf250251",
	DB_NAME:     "postgres",

	REDIS_IP:       "127.0.0.1",
	REDIS_PORT:     "6379",
	REDIS_PASSWORD: "",
	REDIS_DB:       0,

	REDIS_SESSION_DB: 1,
	REDIS_CACHE_DB:   2,

	ACCESS_LOG:      true,
	ACCESS_LOG_PATH: "logs/access.log",

	ERROR_LOG:      true,
	ERROR_LOG_PATH: "logs/error.log",

	INFO_LOG:      true,
	INFO_LOG_PATH: "logs/info.log",

	TEMPLATE_PATH: "./templates",

	//APP_SECRET: "YbskZqLNT6TEVLUA9HWdnHmZErypNJpL",
	APP_SECRET: "something-very-secret",
}

func GetEnv() *Env {
	return &env
}
