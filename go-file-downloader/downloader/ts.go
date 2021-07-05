package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

type fileInfo struct {
	url             string
	path            string
	repo            string
	protocol        string
	pureUrl         string
	format          string
	prefix          string
	fileNameOnly    string
	fileNamePrefix  string
	fixedFileLength int
	lastNum         int
	token           string
	limit           int
}

func Ts() {
	fileInfo := &fileInfo{repo: "repo", limit: 200}
	fileInfo.fileDownload()
}

func (f *fileInfo) fileDownload() {
	f.setStruct()
	f.startIterate()
}

func (f *fileInfo) setInput() {
	var url string
	fmt.Print("Please write url:")
	fmt.Scanln(&url)
	f.url = url
	var path string
	fmt.Print("Please write path:")
	fmt.Scanln(&path)
	f.path = path
}

func (f *fileInfo) setStruct() {
	f.setInput()
	f.setDefault()
	f.setFormat()
	f.setPrefix()
	f.setFilenameOnly()
	f.setFileNamePrefix()
	f.makeDirectory()
}

func (f *fileInfo) setDefault() {
	arr := strings.Split(f.url, "://")
	f.protocol = arr[0]

	arr = strings.Split(arr[1], "?")
	f.pureUrl = arr[0]

	if len(arr) > 1 {
		f.token = arr[1]
	}
}

func (f *fileInfo) setFormat() {
	arr := strings.Split(f.pureUrl, "/")
	fileName := arr[len(arr)-1]

	arr = strings.Split(fileName, ".")
	f.format = "." + arr[len(arr)-1]
}

func (f *fileInfo) setPrefix() {
	arr := strings.Split(f.pureUrl, "/")

	dir := arr[:len(arr)-1]
	dirStr := strings.Join(dir, "/")

	f.prefix = dirStr
}

func (f *fileInfo) setFilenameOnly() {
	arr := strings.Split(f.pureUrl, "/")

	fileName := arr[len(arr)-1]
	arr = strings.Split(fileName, ".")
	f.fileNameOnly = arr[0]
}

func (f *fileInfo) setFileNamePrefix() {
	if strings.Contains(f.fileNameOnly, "-") {
		f.setDashprefix()
	} else {
		f.setNormalprefix()
	}
}

func (f *fileInfo) setDashprefix() {
	arr := strings.Split(f.fileNameOnly, "-")
	fileNamePrefixArr := arr[:len(arr)-1]
	fileNameLast := arr[len(arr)-1]
	f.fileNamePrefix = strings.Join(fileNamePrefixArr, "-") + "-"

	f.lastNum, _ = strconv.Atoi(fileNameLast)
	f.fixedFileLength = len(fileNameLast)
}

func (f *fileInfo) setNormalprefix() {
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

func (f *fileInfo) startIterate() {
	var wg sync.WaitGroup
	wg.Add(f.lastNum)

	for j := 0; j <= f.lastNum/f.limit; j++ {

		thisLastNum := j*f.limit + f.limit
		thisCondition := thisLastNum - f.limit

		if thisLastNum > f.lastNum {
			thisLastNum = f.lastNum
			thisCondition = j * f.limit
		}

		fmt.Println("this Last Num is ", thisLastNum)
		fmt.Println("this Last Con is ", thisCondition)

		done2 := make(chan error)
		for i := thisLastNum; i > thisCondition && i > 0; i-- {
			count := i
			go func() {
				str := f.setDynamicStr(count)
				result := fmt.Sprintf("%s://%s/%s%s%s?%s", f.protocol, f.prefix, f.fileNamePrefix, str, f.format, f.token)

				if len(f.token) == 0 {
					result = result[:len(result)-1]
				}

				defer wg.Done()
				done2 <- DownloadFile(f.repo, f.path, result, str+".ts")
			}()
		}

		for i := 0; i < thisLastNum-thisCondition; i++ {
			<-done2
		}

	}

	wg.Wait()
}

func (f *fileInfo) setDynamicStr(num int) string {
	var str string
	numStr := strconv.Itoa(num)
	if f.fixedFileLength != 0 {
		length := len(numStr)
		zeroCount := f.fixedFileLength - length
		var zeroStr string
		for i := 0; i < zeroCount; i++ {
			zeroStr += "0"
		}
		str = zeroStr + numStr
	} else {
		str = numStr
	}

	return str
}

func (f *fileInfo) makeDirectory() {
	basicpath := f.repo
	if _, err := os.Stat(basicpath); os.IsNotExist(err) {
		err := os.Mkdir(basicpath, 0777)

		if err != nil {
			log.Fatal(err)
		}
	}

	filepath := fmt.Sprintf("%s/%s", f.repo, f.path)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		err := os.Mkdir(filepath, 0777)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func (f *fileInfo) startCmd() {
	_, err := exec.Command("/bin/sh", "ffmpeg.sh", fmt.Sprintf("%s/%s", f.repo, f.path)).Output()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Download success")
	}
}

func DownloadFile(repo string, path string, url string, filename string) error {
	client := &http.Client{}
	// set the data
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
	defer resp.Body.Close()

	// Create the file
	filepath := fmt.Sprintf("%s/%s/%s", repo, path, filename)
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

	fmt.Printf("Downloaded a file %s with size %d\n", filepath, size)

	return err
}
