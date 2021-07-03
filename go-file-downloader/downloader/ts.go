package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type fileInfo struct {
	protocol        string
	pureUrl         string
	format          string
	prefix          string
	fileNameOnly    string
	fileNamePrefix  string
	fixedFileLength int
	lastNum         int
	token           string
}

func Ts() {
	fileDownload()
}

func fileDownload() {
	url := "https://b01-kr-naver-vod.pstatic.net/navertv/c/read/v2/VOD_ALPHA/navertv_2021_05_31_1728/hls/0d415451-c1c7-11eb-a113-d21d68c12640-000007.ts?_lsu_sa_=68955cf621876c56bada35516315a2b80edd3d08690a3fdf35a789c997d03f3567236a8867752901d02c39e215352bdc988fe4a33d47e374f47e235e996f86687f0bff63a657db97c2199e88d7a33749222d586573ab0f6adfa397816bbbdf4a0e51b3aa198af4e98b64b8ef85f0eb9f63528cb1115870abff060da53f39afa5c0c1be453ce6d9ab67c527e5d8b12d3e"
	fileInfo := &fileInfo{}
	fileInfo.setStruct(url)
}

func (f *fileInfo) setStruct(url string) {

	f.getProtocol(url)
	f.getFormat()
	f.getPrefix()
	f.getFilenameOnly()
	f.getFileNamePrefix()

	if f.fixedFileLength != 0 {
		f.startIterateWithPrefix()
	} else {
		f.startIterate()
	}
}

func (f *fileInfo) getProtocol(url string) {
	arr := strings.Split(url, "://")
	f.protocol = arr[0]

	arr = strings.Split(arr[1], "?")
	f.pureUrl = arr[0]

	if len(arr) > 1 {
		f.token = arr[1]
	}
}

func (f *fileInfo) getFormat() {
	arr := strings.Split(f.pureUrl, "/")
	fileName := arr[len(arr)-1]

	arr = strings.Split(fileName, ".")
	f.format = "." + arr[len(arr)-1]
}

func (f *fileInfo) getPrefix() {
	arr := strings.Split(f.pureUrl, "/")

	dir := arr[:len(arr)-1]
	dirStr := strings.Join(dir, "/")

	f.prefix = dirStr

}

func (f *fileInfo) getFilenameOnly() {
	arr := strings.Split(f.pureUrl, "/")

	fileName := arr[len(arr)-1]
	arr = strings.Split(fileName, ".")
	f.fileNameOnly = arr[0]
}

func (f *fileInfo) getFileNamePrefix() {
	if strings.Contains(f.fileNameOnly, "-") {
		arr := strings.Split(f.fileNameOnly, "-")
		fileNamePrefixArr := arr[:len(arr)-1]
		fileNameLast := arr[len(arr)-1]
		f.fileNamePrefix = strings.Join(fileNamePrefixArr, "-") + "-"

		f.lastNum, _ = strconv.Atoi(fileNameLast)
		f.fixedFileLength = len(fileNameLast)
	} else {
		for i := 0; i < len(f.fileNameOnly); i++ {
			val := string(f.fileNameOnly[i])
			_, err := strconv.Atoi(val)
			if err == nil {
				f.fileNamePrefix = f.fileNameOnly[:i]
				f.lastNum, _ = strconv.Atoi(f.fileNameOnly[i:])
			}
		}
	}
}

func (f *fileInfo) startIterate() {
	var wait sync.WaitGroup
	wait.Add(f.lastNum)

	for i := f.lastNum; i > 0; i-- {
		numStr := strconv.Itoa(i)
		str := numStr
		result := fmt.Sprintf("%s://%s/%s%s%s?%s", f.protocol, f.prefix, f.fileNamePrefix, str, f.format, f.token)

		if len(f.token) == 0 {
			result = result[:len(result)-1]
		}

		func() {
			defer wait.Done()
			f.DownloadFile("repo/", result, numStr+".ts")
		}()
	}

	wait.Wait()

}

func (f *fileInfo) startIterateWithPrefix() {
	var wait sync.WaitGroup
	wait.Add(f.lastNum)

	for i := f.lastNum; i > 0; i-- {
		numStr := strconv.Itoa(i)
		length := len(numStr)
		zeroCount := f.fixedFileLength - length
		var zeroStr string
		for i := 0; i < zeroCount; i++ {
			zeroStr += "0"
		}
		str := zeroStr + numStr
		result := fmt.Sprintf("%s://%s/%s%s%s?%s", f.protocol, f.prefix, f.fileNamePrefix, str, f.format, f.token)

		if len(f.token) == 0 {
			result = result[:len(result)-1]
		}

		go func() {
			defer wait.Done()
			f.DownloadFile("repo/", result, numStr+".ts")
		}()
	}

	wait.Wait()
}

func (f *fileInfo) DownloadFile(path string, url string, filename string) error {
	filepath := path + filename
	// Get the data

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36")
	req.Header.Add("Referer", "")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer out.Close()

	// Write the body to file
	size, err := io.Copy(out, resp.Body)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Url is %s\n", url)
	fmt.Printf("Downloaded a file %s with size %d\n", filepath, size)

	return err
}
