package fileUploader

import (
	"fmt"
	"net/http"

	"strings"

	"io/ioutil"

	"google.golang.org/cloud/storage"
)

func view(w http.ResponseWriter, r *http.Request, b *storage.BucketClient) {
	file := strings.Trim(r.URL.Path, "/")
	rc, err := b.NewReader(file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	slurp, err := ioutil.ReadAll(rc)
	defer rc.Close()
	if err != nil {
		fmt.Fprintln(w, err)
	}

	stat, err := b.Stat(file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	w.Header().Set("Content-Type", stat.ContentType)
	w.Write(slurp)
}