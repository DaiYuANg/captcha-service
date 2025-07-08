package constant

type StoreType string

const (
	Memory   StoreType = "memory"
	Badger   StoreType = "badger"
	Redis    StoreType = "redis"
	Etcd     StoreType = "etcd"
	Sqlite   StoreType = "sqlite"
	Bbolt    StoreType = "bbolt"
	Valkey   StoreType = "valkey"
	Mysql    StoreType = "mysql"
	Postgres StoreType = "postgres"
)
