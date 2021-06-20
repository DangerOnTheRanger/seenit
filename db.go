package seenit


type Database interface {
	GetBucket(name string) (Bucket, error)
}

type Bucket interface {
	Has(key string) (bool, error)
	Get(key string) (string, error)
	Put(key string, val string) error
}
