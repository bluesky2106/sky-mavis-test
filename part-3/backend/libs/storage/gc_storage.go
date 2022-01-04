package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// GCStorageInterface : for testing
type GCStorageInterface interface {
	UploadFilePath(baseFolderPath, filePath string, isPublic bool) (string, error)
	UploadBinaryFile(baseFolderPath, fileName string, file multipart.File, isPublic bool) (string, error)
	MakePublic(object string) error
	DownloadFile(object string) ([]byte, error)
	ListByPrefix(prefix string) ([]string, error)
	Delete(object string) error
	SignURL(object string) (string, error)
}

// GCStorage : connect to gcloud
//
// [Credentials] file path
// [BucketName] bucket's name
type GCStorage struct {
	conf *Config
}

// SignedURLOptions : struct
type SignedURLOptions struct {
	GoogleAccessID string `json:"client_id"`
	PrivateKey     string `json:"private_key"`
}

// NewGCStorage : create google cloud storage struct
func NewGCStorage(credentials, bucketName, baseURL string) *GCStorage {
	conf := &Config{
		Credentials: credentials,
		BucketName:  bucketName,
		BaseURL:     baseURL,
	}
	return &GCStorage{
		conf: conf,
	}
}

func (gc *GCStorage) initSignURLOptions() (*SignedURLOptions, error) {
	file, err := ioutil.ReadFile(gc.conf.Credentials)
	if err != nil {
		return nil, err
	}
	var opts *SignedURLOptions
	if err := json.Unmarshal([]byte(file), &opts); err != nil {
		return nil, err
	}
	return opts, nil
}

func (gc *GCStorage) initClient() (*storage.Client, error) {
	ctx := context.Background()

	// Creates a client.
	opt := option.WithCredentialsFile(gc.conf.Credentials)
	return storage.NewClient(ctx, opt)
}

// UploadFilePath : upload file to google storage
//
// params: upload [filePath] to [baseFolderPath] and make it [isPublic]
func (gc *GCStorage) UploadFilePath(baseFolderPath, filePath string, isPublic bool) (string, error) {
	client, err := gc.initClient()
	if err != nil {
		return "", err
	}

	fullFilePath := fmt.Sprintf("%s/%s", baseFolderPath, filepath.Base(filePath))

	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	wc := client.Bucket(gc.conf.BucketName).Object(fullFilePath).NewWriter(context.Background())
	if isPublic {
		wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
		wc.CacheControl = "public, max-age=86400"
	}
	if _, err = io.Copy(wc, f); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	client.Close()
	return fmt.Sprintf("%s/%s", gc.conf.BucketName, fullFilePath), nil
}

// UploadBinaryFile : upload data to google storage
//
// params: upload [file] data with [fileName] to [baseFolderPath] and make it [isPublic]
func (gc *GCStorage) UploadBinaryFile(baseFolderPath, fileName string, file io.Reader, isPublic bool) (string, error) {
	client, err := gc.initClient()
	if err != nil {
		return "", err
	}

	fullFilePath := fmt.Sprintf("%s/%s", baseFolderPath, fileName)

	wc := client.Bucket(gc.conf.BucketName).Object(fullFilePath).NewWriter(context.Background())
	if isPublic {
		wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
		wc.CacheControl = "public, max-age=86400"
	}

	if _, err = io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	client.Close()
	return fmt.Sprintf("%s/%s", gc.conf.BucketName, fullFilePath), nil
}

// MakePublic : make object from private to public
func (gc *GCStorage) MakePublic(object string) error {
	client, err := gc.initClient()
	if err != nil {
		return err
	}

	ctx := context.Background()
	// [START public]
	acl := client.Bucket(gc.conf.BucketName).Object(object).ACL()
	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return err
	}
	// [END public]
	client.Close()
	return nil
}

// DownloadFile : download object
func (gc *GCStorage) DownloadFile(object string) ([]byte, error) {
	client, err := gc.initClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ctx := context.Background()
	// [START download_file]
	rc, err := client.Bucket(gc.conf.BucketName).Object(object).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	return data, nil
	// [END download_file]
}

// DownloadFileAndSave : download object
func (gc *GCStorage) DownloadFileAndSave(object string, saveTo string) error {
	client, err := gc.initClient()
	if err != nil {
		return err
	}
	defer client.Close()

	ctx := context.Background()
	// [START download_file]
	rc, err := client.Bucket(gc.conf.BucketName).Object(object).NewReader(ctx)
	if err != nil {
		return err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(saveTo, data, 0644)
	if err != nil {
		return err
	}

	return nil
	// [END download_file]
}

// ListByPrefix : list all objects in the bucket starting with prefix
func (gc *GCStorage) ListByPrefix(prefix string) (objects []string, err error) {
	client, err := gc.initClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// [START storage_list_files_with_prefix]
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	it := client.Bucket(gc.conf.BucketName).Objects(ctx, &storage.Query{
		Prefix:    prefix,
		Delimiter: "/",
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		objects = append(objects, attrs.Name)
	}
	// [END storage_list_files_with_prefix]
	return
}

// Delete : delete object in a bucket
func (gc *GCStorage) Delete(object string) error {
	client, err := gc.initClient()
	if err != nil {
		return err
	}
	defer client.Close()

	// [START delete_file]
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	o := client.Bucket(gc.conf.BucketName).Object(object)
	if err := o.Delete(ctx); err != nil {
		return err
	}
	// [END delete_file]
	return nil
}

// SignURL : get public URL from the private object
//
// params : [object]
// retrun: [url string], [err error]
func (gc *GCStorage) SignURL(object string, expireAt time.Time) (string, error) {
	opts, err := gc.initSignURLOptions()
	if err != nil {
		return "", err
	}
	return storage.SignedURL(gc.conf.BucketName, object, &storage.SignedURLOptions{
		GoogleAccessID: opts.GoogleAccessID,
		PrivateKey:     []byte(opts.PrivateKey),
		Method:         "GET",
		Expires:        expireAt,
	})
}
