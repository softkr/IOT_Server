package main

import (
	"fmt"
	"iot/azure"
	"iot/grpc/client"
	"iot/log"
	"iot/socket"
	_ "iot/socket"
	"iot/work"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type WatchFile struct {
	Guid     string   `json:"guid"`
	FileName string   `json:"fileName"`
	VideoMd5 string   `json:"videoMd5"`
	SubFile  []string `json:"subFile"`
}

// 애저 스토리지 전송
func azureStrorage(Project, Guid, FileName string) {
	azure.Storage(Project, Guid, FileName)
}

// 서부파일 삭제
func subFileRemove(subFiles []string) {
	for _, fileName := range subFiles {
		err := os.Remove("./files/" + fileName)
		if err != nil {
			log.Error.Println(err)
		}
	}
}

func init() {
	go socket.Run()
}

func main() {
	router := gin.Default()
	// 폴더 생성
	_, err := os.Stat("./files")
	if err != nil {
		err := os.Mkdir("files", os.ModePerm)
		if err != nil {
			log.Error.Println(err)
			return
		}
	}

	router.POST("/fileinfo", func(c *gin.Context) {
		var data WatchFile
		err := c.BindJSON(&data)
		if err != nil {
			log.Error.Println(err)
			return
		}
		client.SetFileInfo(data.Guid, data.FileName, data.VideoMd5, data.SubFile)
		c.JSON(http.StatusCreated, gin.H{"status": 201})
	})

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		data := client.GetFileInfo(file.Filename)
		subFiles := strings.Split(data.SubFile, " ")

		if work.FileNameCheck(len(subFiles), file.Filename) == false {
			// 파일 존재여부
			c.JSON(http.StatusOK, gin.H{"status": 204, "messages": fmt.Sprintf("%v not found", file.Filename)})
		} else if work.SaveFile(file, c) == false {
			// 파일저장
			c.JSON(http.StatusOK, gin.H{"status": 502, "messages": "Save Failed"})
		} else if work.GetMD5(file.Filename) == false {
			// 파일 md5 검증
			c.JSON(http.StatusOK, gin.H{"status": 203, "messages": "invalid file"})
		} else {
			if client.SubFileCount(data.VideoMd5) == 0 {
				// 파일 수량 맞으면 통합 진행
				work.Integrated(subFiles, data.FileName, data.VideoMd5, c)
				subFileRemove(subFiles)
				client.DeleteFileInfo(data.VideoMd5)
				go func() {
					project := client.GetProject(data.Guid).Project
					azureStrorage(project, data.Guid, data.FileName)
				}()
			}
			c.JSON(http.StatusCreated, gin.H{"status": 201, "messages": "정상처리 되었습니다."})
		}
	})
	router.Run(":8090")
}
