package seenit

import (
	"fmt"
	"golang.org/x/image/draw"
	_ "golang.org/x/image/webp"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

const (
	seenMarker = "seen"
)

func HaveSeen(community string, hash string, db Database) (bool, error) {
	bucket, err := db.GetBucket(community)
	if err != nil {
		return false, err
	}
	return bucket.Has(hash)
}

func RecordHash(community string, hash string, db Database) error {
	bucket, err := db.GetBucket(community)
	if err != nil {
		return err
	}
	return bucket.Put(hash, seenMarker)
}


func ImageToHash(img image.Image) (string, error) {
	coloredResized := image.NewRGBA(image.Rect(0, 0, 8, 8))
	draw.NearestNeighbor.Scale(coloredResized, coloredResized.Rect, img, img.Bounds(), draw.Over, nil)
	grayResized := image.NewGray(coloredResized.Bounds())
	for y := coloredResized.Bounds().Min.Y; y < coloredResized.Bounds().Max.Y; y++ {
		for x := coloredResized.Bounds().Min.X; x < coloredResized.Bounds().Max.X; x++ {
			grayResized.Set(x, y, coloredResized.At(x, y))
		}
	}

	var pixelSum uint64
	for y := grayResized.Bounds().Min.Y; y < grayResized.Bounds().Max.Y; y++ {
		for x := grayResized.Bounds().Min.X; x < grayResized.Bounds().Max.Y; x++ {
			pixel := grayResized.GrayAt(x, y)
			value := pixel.Y
			pixelSum += uint64(value)
		}
	}
	totalPixels := uint64(grayResized.Bounds().Dx() * grayResized.Bounds().Dy())
	avgValue := uint8(pixelSum / totalPixels)

	var imageHash uint64
	for y := grayResized.Bounds().Min.Y; y < grayResized.Bounds().Max.Y; y++ {
		for x := grayResized.Bounds().Min.X; x < grayResized.Bounds().Max.Y; x++ {
			pixel := grayResized.GrayAt(x, y)
			value := pixel.Y
			if value > avgValue {
				hashIndex := (y * 8) + x
				imageHash |= (1 << hashIndex)
			}
		}
	}
	return fmt.Sprintf("%016x", imageHash), nil
}
