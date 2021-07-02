package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type fileInfo struct {
	lastUrl         string
	protocol        string
	pureUrl         string
	format          string
	prefix          string
	fileNameOnly    string
	fileNamePrefix  string
	fixedFileLength int
	lastNum         int
}

func Ts() {
	fileDownload()
}

func fileDownload() {
	// token := "_lsu_sa_=6705cef3b1d462e63cd6f55b69c55fb95e8a34d88d067fc83f772dc217e53225fe2fba9f61f5090700f533c285313b90e3cb7c94073af61f1fe93a53ca079792b61f166dfaedcd057a5405bba53397d0690c65db3124fbe15b1ebc19c41d056cb2097df5d273790e6c0972d3ced91d90edf7bfa1d1cbe5321cc3439b4a7bbced1dc12af7fbcee1125d0c0c2a1b918a2f"
	// DownloadFile("repo/sample.ts", fmt.Sprintf("%s?%s", startUrl, token))
	lastUrl := "https://b01-kr-naver-vod.pstatic.net/navertv/c/read/v2/VOD_ALPHA/navertv_2021_06_24_448/hls/86e70095-d49a-11eb-8ed1-5ebafcba569f-0000123.ts"
	fileInfo := &fileInfo{}
	fileInfo.setStruct(lastUrl)

	fmt.Println(fileInfo)
}

func (f *fileInfo) setStruct(lastUrl string) {

	f.lastUrl = lastUrl
	f.getProtocol()
	f.getFormat()
	f.getPrefix()
	f.getFilenameOnly()
	f.getFileNamePrefix()

}

func (f *fileInfo) getProtocol() {
	arr := strings.Split(f.lastUrl, "://")
	f.protocol = arr[0]
	f.pureUrl = arr[1]
}

func (f *fileInfo) getFormat() {
	arr := strings.Split(f.lastUrl, "/")
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
				break
			}
		}
	}
}

func (f *fileInfo) startIterate() {
}

func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
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

	fmt.Printf("Downloaded a file %s with size %d", filepath, size)

	return err
}
