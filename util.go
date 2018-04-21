package main

import (
	"os"
	"path/filepath"
	"log"
	"strings"
	"io"
	"time"
	"strconv"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetDirectory(file string) string {
	dir, err := filepath.Abs(filepath.Dir(file))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func ChangePathFormat(file string) string{
	return strings.Replace(file, "\\", "/", -1)
}

func DeleteFile(file string) error {
	err := os.Remove(file)
	return err
}

func CopyFile(src string) (string, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()
	dst := filepath.Join(GetDirectory("."), strconv.FormatInt(time.Now().Unix(), 10)+".ass")
	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		DeleteFile(dst)
		return "", err
	}
	return dst, nil
}
