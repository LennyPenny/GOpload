package fileUploader

import (
	"fmt"
	"net/http"

	"math/rand"

	"path/filepath"

	"io/ioutil"

	"google.golang.org/cloud/storage"
)

var SECRETKEY = "f6ce4fa4b8c5a9e9d5a50d877b4461b3050e888a97edbc7a5da71638ec0f5c40" //change this to something secure so nobody except you can upload stuff

var lookUp = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
func randomFileName(orig string) (string) {
	ext := filepath.Ext(orig)

	leng := len(lookUp) - 1
	i1 := rand.Intn(leng)
	i2 := rand.Intn(leng)
	return lookUp[i1:i1+1] + lookUp[i2:i2+1] + ext
}

func upload(w http.ResponseWriter, r *http.Request, b *storage.BucketClient) {
	if r.FormValue("thisisscecret") != SECRETKEY { fmt.Fprintln(w, "eheh"); return }
	filee, header, err := r.FormFile("file")
	fileName := randomFileName(header.Filename)
	defer filee.Close()
	
	if err != nil { 
		fmt.Fprintln(w, err) 
		return 
	}

	cont, err := ioutil.ReadAll(filee)
	if err != nil { 
		fmt.Fprintln(w, err)
		return 
	}

	wc := b.NewWriter(fileName, &storage.Object {
		ContentType: header.Header.Get("Content-Type"),
	})

	if _, err := wc.Write(cont); err != nil {
		fmt.Fprintln(w, err) 
	}
	if err := wc.Close(); err != nil {
		fmt.Fprintln(w, err) 
	}

	_, err = wc.Object()
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintln(w, fileName)
}