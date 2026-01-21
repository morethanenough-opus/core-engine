package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

)

func GenerateRandomString(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}
	b := make([]byte, length)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func CalculateSHA256Hash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

func CalculateSHA512Hash(data string) string {
	hash := sha512.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

func GenerateRandomInteger(min, max int) (int, error) {
	if min >= max {
		return 0, errors.New("min must be less than max")
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()) + min, nil
}

func GetCurrentWorkingDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return dir, nil
}

func GetFileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

func DownloadFile(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func IsFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func GetFileExtension(path string) string {
	ext := filepath.Ext(path)
	if len(ext) > 0 {
		return strings.ToLower(ext[1:])
	}
	return ""
}

func GetFileBaseName(path string) string {
	return filepath.Base(path)
}

func RemoveFile(path string) error {
	return os.Remove(path)
}

func LogError(message string, err error) {
	log.Printf("%s: %v", message, err)
}

func CheckTimeout(duration time.Duration, timeoutChan chan bool) bool {
	select {
	case <-time.After(duration):
		return true
	case <-timeoutChan:
		return false
	}
}

func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

func ParseTime(timeStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, timeStr)
}