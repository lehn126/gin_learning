package config

type ConnPool struct {
	Enable          bool   `yaml:"enable" json:"enable"`
	MaxIdleConns    int    `yaml:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns" json:"maxOpenConns"`
	ConnMaxLifetime string `yaml:"connMaxLifetime" json:"connMaxLifetime"`
}

type SqliteDB struct {
	Enable   bool     `yaml:"enable" json:"enable"`
	File     string   `yaml:"file" json:"file"`
	ConnPool ConnPool `yaml:"connPool" json:"connPool"`
}

type MysqlDB struct {
	Enable   bool     `yaml:"enable" json:"enable"`
	Host     string   `yaml:"file" json:"file"`
	Port     string   `yaml:"port" json:"port"`
	User     string   `yaml:"user" json:"user"`
	Password string   `yaml:"password" json:"password"`
	Database string   `yaml:"database" json:"database"`
	Charset  string   `yaml:"charset" json:"charset"`
	ConnPool ConnPool `yaml:"connPool" json:"connPool"`
}

type Database struct {
	Sqlite map[string]SqliteDB `yaml:"sqlite" json:"sqlite"`
	Mysql  map[string]MysqlDB  `yaml:"mysql" json:"mysql"`
}

type DB struct {
	DB Database `yaml:"db" json:"db"`
}

var DB_CONFIG DB = DB{
	DB: Database{
		Sqlite: make(map[string]SqliteDB),
		Mysql:  make(map[string]MysqlDB),
	},
}
