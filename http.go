package helper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// httpDownload downloads a file from a url to a local file.
func httpDownload(url, dest string) error {
	log.Printf("downloading %s to %s", url, dest)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}
	defer resp.Body.Close()
	out, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to copy file: %v", err)
	}
	return nil
}
