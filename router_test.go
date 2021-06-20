package seenit

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"strconv"
	"testing"
)


func TestServeLanding(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }
	record := httptest.NewRecorder()
	handler := http.HandlerFunc(ServeLanding)
	handler.ServeHTTP(record, request)

	if status := record.Code; status != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, status)
    }

	outbuf := new(bytes.Buffer)
	err = RenderLanding(outbuf)
	if err != nil {
		t.Fatal(err)
	}
	if outbuf.String() != record.Body.String() {
		t.Errorf("Mismatch in expected body content - got: %s\n\nexpected: %s", outbuf.String(), record.Body.String())
	}
}


func TestServeUpload(t *testing.T) {
	query := url.Values{}
	query.Add("community", "Matrix")
	request, err := http.NewRequest("POST", "/upload", strings.NewReader(query.Encode()))
    if err != nil {
        t.Fatal(err)
    }
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    request.Header.Add("Content-Length", strconv.Itoa(len(query.Encode())))
	record := httptest.NewRecorder()
	handler := http.HandlerFunc(ServeUpload)
	handler.ServeHTTP(record, request)

	if status := record.Code; status != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, status)
    }

	outbuf := new(bytes.Buffer)
	err = RenderUpload("Matrix", outbuf)
	if err != nil {
		t.Fatal(err)
	}
	if outbuf.String() != record.Body.String() {
		t.Errorf("Mismatch in expected body content - got: %s\n\nexpected: %s", outbuf.String(), record.Body.String())
	}
}


func buildPostBody(t *testing.T, filename string, community string) (*bytes.Buffer, string) {
	file, err := os.Open(filename)

    if err != nil {
        t.Fatal(err)
    }
    defer file.Close()

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile("image", filepath.Base(file.Name()))
    if err != nil {
        t.Fatal(err)
    }
    io.Copy(part, file)
	cWriter, err := writer.CreateFormField("community")
	if err != nil {
		t.Fatal(err)
	}
	io.WriteString(cWriter, community)
    writer.Close()

	return body, writer.Boundary()
}

func TestServeResult(t *testing.T) {
	body, boundary := buildPostBody(t, "test_image.png", "Matrix")
	request, err := http.NewRequest("POST", "/result", body)
    if err != nil {
        t.Fatal(err)
    }
	request.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", boundary))
	request.Header.Add("Content-Length", strconv.Itoa(body.Len()))

	db := MockDatabase{Buckets: make(map[string]MockBucket, 0)}
	
	record := httptest.NewRecorder()
	handler := http.HandlerFunc(BindHandler(ServeResult, &db))
	handler.ServeHTTP(record, request)

	if status := record.Code; status != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, status)
    }

	outbuf := new(bytes.Buffer)
	err = RenderUnseen(outbuf)
	if err != nil {
		t.Fatal(err)
	}
	if outbuf.String() != record.Body.String() {
		t.Errorf("Mismatch in expected body content - got: %s\n\nexpected: %s", outbuf.String(), record.Body.String())
	}

	body, boundary = buildPostBody(t, "test_image.png", "Matrix")
	request, err = http.NewRequest("POST", "/result", body)
    if err != nil {
        t.Fatal(err)
    }
	request.Header.Add("Content-Type", fmt.Sprintf("multipart/form-data; boundary=%s", boundary))
	request.Header.Add("Content-Length", strconv.Itoa(body.Len()))
	record = httptest.NewRecorder()
	handler = http.HandlerFunc(BindHandler(ServeResult, &db))
	handler.ServeHTTP(record, request)

	if status := record.Code; status != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, status)
    }

	outbuf = new(bytes.Buffer)
	err = RenderSeen(outbuf)
	if err != nil {
		t.Fatal(err)
	}
	if outbuf.String() != record.Body.String() {
		t.Errorf("Mismatch in expected body content - got: %s\n\nexpected: %s", outbuf.String(), record.Body.String())
	}
}
