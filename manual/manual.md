# IOTServer

인핸드플러스 iot 서버이며 인핸드워치에 마춰 개발된 서버입니다.  
다른 용도로 사용이 불가능합니다.  
현재 아래와 같은 이슈 사항을 개선하려고 개발된 프로그램입니다.  
본 프로그램은 Golang으로 개발되었으며 사용된 프레임워크는 Gin 입니다.

> #이슈 사항

1. <image src ="./img/0kb.png">
2. 동영상 사이트는 약 2mb 미만 재생은 30초 화면은 검정색 혹은 파란 등 색상으로 보임
3. 복약 기록은 있으나 동영상 없는 현상
4. 인터넷 환경이 안좋으면 동영상 전송시간 오래 걸리거나 재생안되는 문제

> #필요한 부가서비스

- 데이터 베이스
- gRPC
- Azure Storage

> #사용된 패캐지

- github.com/Azure/azure-storage-blob-go
- github.com/gin-gonic/gin
- github.com/joho/godotenv
- iot/grpcx

  1. github.com/joho/godotenv
  2. go.mongodb.org/mongo-driver
  3. google.golang.org/grpc
  4. google.golang.org/protobuf

> #아키텍쳐

- 전체 구조도
  <image src="./img/iot.jpg">

> #설치

```golang
go mod iot_server init
go get mod tidy
```

> #환경설정

.env 파일 만들고 .env 파일에 맞게 작성하면됩니다.

```env
HOST=gRPC 서버주소
PORT=gRPC 포트
AZURE_STORAGE_ACCOUNT=애저스트로지key
```

> #실행

```go
go run main.go
```

> #테스트

테스트 위한 사전 준비가 필요합니다.

- 재생가능한 동영상 하나
- 워치 guid 하나
- 동영상 분할 전 md5 정보
- 동영상 분할후 분할된 동영상파일 조각 md5 정보로된 파일 들

테스트 진행

1. json 형식 데이터 생성  
   {
   "guid": "21IHPA02720A",
   "fileName": "100101_005332_DDA142164623.mp4",
   "fullFileMD5": "8a4ae3c9b472b09a944d4ede02f885e3",
   "subFile": ["a376ceeae4a2b0bee79aba4d6c100080",
   "d31119f325db61d41358722652b66c81",
   "2c0134ae4f33525c367ab70d7849603e",
   "7440ee1217367722141360fbb2c4d196",
   "7e524e73d435eb227a6f19de62caaa44",
   "e50ca16a818567a47f6fc1e4318f9f7d",
   "611e2056f0703ec455aea39fd76b46de",
   "e04ab7b88242b409fe592fea9bb96257"],
   }
2. http://서버주소/fileinfo 전송
3. http://서버주소/upload 파일 업로드  
   분할된 파일 하나씩 반복으로 전부보냄
4. 지정된 스토리지 데이터 존재여부확인 재생여부확인
