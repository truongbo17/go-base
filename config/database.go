package config

type DatabaseConnection struct {
	DatabaseRelation DatabaseRelation `mapstructure:",squash"`
	DatabaseMongo    DatabaseMongo    `mapstructure:",squash"`
}

const DefaultStringSizeMySql uint = 256

type DatabaseRelation struct {
	Username string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASS"`
	Host     string `mapstructure:"DB_HOST"`
	Port     uint   `mapstructure:"DB_PORT"`
	Database string `mapstructure:"DB_DATABASE"`
}

type DatabaseMongo struct {
	Uri string `mapstructure:"MONGO_URI"`
}
