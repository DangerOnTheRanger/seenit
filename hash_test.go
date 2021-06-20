package seenit

import (
	"os"
	"testing"
	"image/png"
)

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


func TestHaveSeen(t *testing.T) {
	db := MockDatabase{Buckets: make(map[string]MockBucket, 0)}
	hash := "thisisatesthash"
	community := "test-community"
	seen, err := HaveSeen(community, hash, &db)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if seen {
		t.Fatal("Hash incorrectly marked as seen")
	}
	_, hasBucket := db.Buckets[community]
	if !hasBucket {
		t.Fatalf("Bucket %q not added", community)
	}
	err = RecordHash(community, hash, &db)
	if err != nil {
		t.Fatal(err.Error())
	}
	seen, err = HaveSeen(community, hash, &db)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !seen {
		t.Fatal("Hash incorrectly marked as unseen")
	}
}

func TestHash(t *testing.T) {
	imageFile, err := os.Open("test_image.png")
	defer imageFile.Close()
	if err != nil {
		t.Fatal(err.Error())
	}
	image, err := png.Decode(imageFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	hash, err := ImageToHash(image)
	if err != nil {
		t.Fatal(err.Error())
	}
	expectedHash := "0080c0c0c0c08000"
	if hash != expectedHash {
		t.Fatalf("Expected hash %q, got %q", expectedHash, hash)
	}
}
