package test

import (
	"fmt"
	"iot/grpc/client"
	"testing"
)

func TestSetFileInfo(t *testing.T) {
	args := []string{
		"a376ceeae4a2b0bee79aba4d6c100080",
		"d31119f325db61d41358722652b66c81",
		"2c0134ae4f33525c367ab70d7849603e",
		"7440ee1217367722141360fbb2c4d196",
		"7e524e73d435eb227a6f19de62caaa44",
		"e50ca16a818567a47f6fc1e4318f9f7d",
		"611e2056f0703ec455aea39fd76b46de",
		"e04ab7b88242b409fe592fea9bb96257",
	}
	result := client.SetFileInfo("21IHPA0000A", "100101_005332_DDA142164623.mp4", "ad4458a45e68ca64e2301ebe975f927f", args)
	fmt.Println(result)
}

func TestGetFileInfo(t *testing.T) {
	result := client.GetFileInfo("a376ceeae4a2b0bee79aba4d6c100080")
	println(result.SubFile)
}

func TestPutFileInfo(t *testing.T) {
	result := client.PutFileInfo("a376ceeae4a2b0bee79aba4d6c100080")
	fmt.Println(result)
}

func TestSubFileCount(t *testing.T) {
	result := client.SubFileCount("ad4458a45e68ca64e2301ebe975f927f")
	fmt.Println(result)
}

func TestDeleteFileInfo(t *testing.T) {
	result := client.DeleteFileInfo("ad4458a45e68ca64e2301ebe975f927f")
	fmt.Println(result)
}
