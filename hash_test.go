package seenit

import (
	"os"
	"testing"
	"image/png"
)

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
