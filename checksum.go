package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func CalcChecksum(src string) (string, error) {
	hash := sha256.New()
	f, err := os.Open(src)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func AppendChecksum(dst string, src string) error {
	f, err := os.OpenFile(dst, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	hash, err := CalcChecksum(src)
	if err != nil {
		return err
	}
	fmt.Fprintf(f, "%s %s\n", hash, src)
	return nil
}
