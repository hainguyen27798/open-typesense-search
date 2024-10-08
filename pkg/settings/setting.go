package settings

type Config struct {
	Server         ServerSettings
	MongoDB        MongoDBSettings
	EmbeddingModel EmbeddingModelSettings
	QDrant         QDrantSettings
}

type ServerSettings struct {
	Port int
	Mode string
}

type MongoDBSettings struct {
	URL         string
	Username    string
	Password    string
	Database    string
	MaxPoolSize uint64
}

type EmbeddingModelSettings struct {
	Auth        string
	TextEMURL   string
	TextEMName  string
	ImageEMURL  string
	ImageEMName string
}

type QDrantSettings struct {
	Host string
	Port int
}
