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
	protocol    string
	format      string
	prefix      string
	startNum    string
	lastNum     string
	numMaxLenth int
}

func Ts() {
	fileDownload()
}

func fileDownload() {
	// token := "_lsu_sa_=6705cef3b1d462e63cd6f55b69c55fb95e8a34d88d067fc83f772dc217e53225fe2fba9f61f5090700f533c285313b90e3cb7c94073af61f1fe93a53ca079792b61f166dfaedcd057a5405bba53397d0690c65db3124fbe15b1ebc19c41d056cb2097df5d273790e6c0972d3ced91d90edf7bfa1d1cbe5321cc3439b4a7bbced1dc12af7fbcee1125d0c0c2a1b918a2f"
	// DownloadFile("repo/sample.ts", fmt.Sprintf("%s?%s", startUrl, token))
	startUrl := "https://b01-kr-naver-vod.pstatic.net/navertv/c/read/v2/VOD_ALPHA/navertv_2021_06_24_448/hls/86e70095-d49a-11eb-8ed1-5ebafcba569f-0000001.ts"
	lastUrl := "https://b01-kr-naver-vod.pstatic.net/navertv/c/read/v2/VOD_ALPHA/navertv_2021_06_24_448/hls/86e70095-d49a-11eb-8ed1-5ebafcba569f-0000123.ts"
	fileInfo := &fileInfo{}
	fileInfo.setStruct(startUrl, lastUrl)
}

func (f *fileInfo) setStruct(startUrl string, lastUrl string) {

	f.protocol = getProtocol(startUrl)
	f.format = getFormat(startUrl)
	f.prefix = getPrefix(startUrl)

	fmt.Println(f.protocol)
	fmt.Println(f.format)
	fmt.Println(f.prefix)
}

func getProtocol(url string) string {
	arr := strings.Split(url, "://")
	return arr[0]
}

func getFormat(url string) string {
	arr := strings.Split(url, "/")
	fileName := arr[len(arr)-1]

	arr = strings.Split(fileName, ".")
	return "." + arr[len(arr)-1]
}

func getPrefix(url string) string {
	arr := strings.Split(url, "://")
	pureUrl := arr[len(arr)-1]
	arr = strings.Split(pureUrl, "/")

	dir := arr[:len(arr)-1]
	dirStr := strings.Join(dir, "/")

	fileName := arr[len(arr)-1]
	arr = strings.Split(fileName, ".")
	fileNameOnly := arr[0]

	var prefixName string
	if strings.Contains(fileNameOnly, "-") {
		arr = strings.Split(fileNameOnly, "-")
		arr = arr[:len(arr)-1]
		prefixName = strings.Join(arr, "-")
	} else {
		for i := 0; i < len(fileNameOnly); i++ {
			val := string(fileNameOnly[i])
			_, err := strconv.Atoi(val)
			if err == nil {
				prefixName = fileNameOnly[:i]
				break
			}
		}
	}

	return dirStr + "/" + prefixName

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
