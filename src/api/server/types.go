package server

type Configuration struct {
	AppName string  `mapstructure:"AppName"`
	Server  *Server `mapstructure:"Server"`
	Mongo   *Mongo  `mapstructure:"Mongo"`
	Sql     *Sql    `mapstructure:"Sql"`
	Kafka   *Kafka  `mapstructure:"Kafka"`
}

type Server struct {
	Port   string   `mapstructure:"Port"`
	Origin []string `mapstructure:"Origin"`
	Header []string `mapstructure:"Header"`
}

type Mongo struct {
	ConnectionString string       `mapstructure:"ConnectionString"`
	Database         string       `mapstructure:"Database"`
	Collections      *Collections `mapstructure:"Collections"`
}

type Collections struct {
	User string `mapstructure:"User"`
}

type Sql struct {
	ConnectionString string `mapstructure:"ConnectionString"`
	DatabaseName     string `mapstructure:"DatabaseName"`
}

type Kafka struct {
	Broker string  `mapstructure:"Broker"`
	Group  string  `mapstructure:"Group"`
	Topics *Topics `mapstructure:"Topics"`
}

type Topics struct {
	CreatedUser string `mapstructure:"CreatedUser"`
	UpdatedUser string `mapstructure:"UpdatedUser"`
	DeletedUser string `mapstructure:"DeletedUser"`
}
