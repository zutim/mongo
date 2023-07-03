package mongo

type Conf struct {
	Dsn  string
	Pool int
}

func NewMongoClient(conf *Conf) *MongoClient {
	return NewMongo(conf.Dsn, conf.Pool)
}
