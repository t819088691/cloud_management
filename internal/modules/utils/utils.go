package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func RandString(length int64) string {
	sources := []byte("0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sourceLength := len(sources)
	var i int64 = 0
	for ; i < length; i++ {
		result = append(result, sources[r.Intn(sourceLength)])
	}
	return string(result)
}

func MD5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

func WorkDir() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	wd := filepath.Dir(execPath)
	if filepath.Base(wd) == "bin" {
		wd = filepath.Dir(wd)
	}
	return wd, nil
}

func FileExists(file string) bool {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	if os.IsPermission(err) {
		return false
	}
	return true
}
