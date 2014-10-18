package fileUploader

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/file"

	"github.com/golang/oauth2/google"
	"google.golang.org/cloud/storage"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	
	c := appengine.NewContext(r)

	bucketName, err := file.DefaultBucketName(c)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	
	config := google.NewAppEngineConfig(c, storage.ScopeFullControl)
	client := storage.New(appengine.AppID(c), config.NewTransport())
	b := client.BucketClient(bucketName)

	if r.Method == "POST" {
		upload(w, r, b)
	} else if r.Method == "GET" {
		view(w, r, b)
	}
}