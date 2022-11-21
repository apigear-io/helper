package helper

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ExtractTarGz extracts a tar.gz file to a destination directory.
func ExtractTarGz(src, dest string) error {
	r, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer r.Close()
	gzipReader, err := gzip.NewReader(r)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %v", err)
	}
	defer gzipReader.Close()
	tr := tar.NewReader(gzipReader)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error reading tar: %v", err)

		}
		if hdr.Typeflag == tar.TypeReg {
			tf, err := os.Create(fmt.Sprintf("%s/%s", dest, hdr.Name))
			if err != nil {
				return fmt.Errorf("failed to create file: %v", err)
			}
			defer tf.Close()
			_, err = io.Copy(tf, tr)
			if err != nil {
				return fmt.Errorf("failed to copy file: %v", err)
			}
		}
	}
	return nil
}

// CreateTarGz creates a tar.gz file from a source directory.
func CreateTarGz(src, dst string) error {
	f, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer f.Close()
	gw := gzip.NewWriter(f)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}
		header.Name = path
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		if !info.IsDir() {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			if _, err := io.Copy(tw, f); err != nil {
				return err
			}
		}
		return nil
	})
}
