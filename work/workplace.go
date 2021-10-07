package work

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/softkr/grpcx/client"
	"io"
	"io/ioutil"
	"iot/log"
	"mime/multipart"
	"os"
	"path/filepath"
)

func FileNameCheck(length int, subFileName string) bool {
	if length > 1 {
		// db 존재여부
		if FileExist(subFileName) == false {
			// 파일존재 여부
			client.PutFileInfo(subFileName)
		}
		return true
	} else {

		return false
	}
}

func SaveFile(file *multipart.FileHeader, c *gin.Context) bool {
	err := c.SaveUploadedFile(file, "./files/"+file.Filename)
	if err != nil {
		log.Error.Println(err)
		return false
	} else {
		return true
	}

}

// GetMD5 md5 파일 검증
func GetMD5(fileName string) bool {
	url := filepath.Join("files", fileName)
	f, err := os.Open(url)
	if err != nil {
		log.Error.Println(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Error.Println(err)
		}
	}(f)
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Error.Println(err)
	}
	//fmt.sPrintf("%x", h.Sum(nil))
	md5 := fmt.Sprintf("%x", h.Sum(nil))
	if md5 == fileName {
		return true
	} else {
		os.Remove(url)
		return false
	}
}

// FileExist 존재여부
func FileExist(file string) bool {
	path := filepath.Join("files", file)
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

//Verification
func Verification(VideoMd5 string, fileName string) bool {
	url := filepath.Join("files", fileName)
	f, err := os.Open(url)
	if err != nil {
		log.Error.Println(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Error.Println(err)
		}
	}(f)
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Error.Println(err)
	}
	//fmt.sPrintf("%x", h.Sum(nil))
	md5 := fmt.Sprintf("%x", h.Sum(nil))
	if md5 == VideoMd5 {
		return true
	} else {
		return false
	}
}

//통합
func Integrated(subFiles []string, fileFullName string, VideoMd5 string, c *gin.Context) bool {
	for _, subFileName := range subFiles {
		fii, err := os.OpenFile("./files/"+fileFullName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
		if err != nil {
			log.Error.Println(err)
		}
		f, err := os.OpenFile("./files/"+subFileName, os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Error.Println(err)
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			log.Error.Println(err)
		}
		fii.Write(b)
		f.Close()
	}
	return Verification(VideoMd5, fileFullName)
}
