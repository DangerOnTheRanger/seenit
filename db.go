package seenit


type Database interface {
	GetBucket(name string) (Bucket, error)
}

type Bucket interface {
	Has(key string) (bool, error)
	Get(key string) (string, error)
	Put(key string, val string) error
}


type MockDatabase struct {
	Buckets map[string]MockBucket 
}

func (d *MockDatabase) GetBucket(name string) (Bucket, error) {
	if _, ok := d.Buckets[name]; !ok {
		d.Buckets[name] = MockBucket{Entries: make(map[string]string, 0)}
	}
	bucket := d.Buckets[name]
	return &bucket, nil
}

type MockBucket struct {
	Entries map[string]string
}

func (b *MockBucket) Has(key string) (bool, error) {
	_, ok := b.Entries[key]
	return ok, nil
}

func (b *MockBucket) Get(key string) (string, error) {
	val := b.Entries[key]
	return val, nil
}

func (b *MockBucket) Put(key string, val string) error {
	b.Entries[key] = val
	return nil
}
