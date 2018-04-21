package main

import (
	"io/ioutil"
	"regexp"
	"os/exec"
	"fmt"
	"strings"
	"path/filepath"
	"strconv"
	"time"
)

func ParseAss(subtitlePath string) (s Subtitle, err error) {
	if tmpBuf, err1 := ioutil.ReadFile(subtitlePath); err != nil {
		err = err1
		return
	} else {
		tmpAssContent := string(tmpBuf)
		if reg, err1 := regexp.Compile("{{\\s*[\u4e00-\u9fa5_a-zA-Z0-9]+\\s*}}"); err != nil {
			err = err1
			return
		} else {
			results := reg.FindAllString(tmpAssContent, -1)
			if results == nil {
				s.Words=[]string{subtitlePath}
				return
			} else {
				s.Len = len(results)
				s.Words = results
				return
			}
		}
	}
}

func GenerateMovie(movieFile string, subtitlePath string, words []string) (string, error) {
	dstSubtitleFile, err := CopyFile(subtitlePath)
	defer DeleteFile(dstSubtitleFile)
	if err != nil {
		return "", err
	}
	if tmpBuf, err := ioutil.ReadFile(dstSubtitleFile); err != nil {
		return "", err
	} else {
		tmpAssContent := string(tmpBuf)
		if reg, err := regexp.Compile("{{\\s*[\u4e00-\u9fa5_a-zA-Z0-9]+\\s*}}"); err != nil {
			return "", err
		} else {
			results := reg.FindAllString(tmpAssContent, -1)
			if results == nil {
				return "", nil
			} else {
				wordsLen := len(results)
				content := tmpAssContent
				for i := 0; i < wordsLen; i++ {
					content = strings.Replace(content, results[i], words[i], 1)
				}
				ioutil.WriteFile(dstSubtitleFile, []byte(content), 0)
			}
		}
	}
	outputFile := filepath.Join(GetDirectory("."), strconv.FormatInt(time.Now().Unix(), 10)+".gif")
	err = MergeFilm(movieFile, dstSubtitleFile, outputFile)
	if err != nil {
		return "", err
	}
	return outputFile, nil

}

func MergeFilm(movieFile string, assFile string, outputFile string) error {
	assFile=filepath.Base(assFile)
	cmd := exec.Command("ffmpeg", "-i", movieFile,
		"-vf", fmt.Sprintf("ass=%s", assFile),
		"-y", outputFile)
	err := cmd.Run()
	return err

}
