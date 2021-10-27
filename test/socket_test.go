package test

import (
	"fmt"
	"iot/grpc/client"
	"testing"
)

// 워치 온라인
func TestWatchStatus(t *testing.T) {
	on := client.WatchState("21IHPA0000A", "121.134.241.239:55770", 1)
	fmt.Println(on)
	off := client.WatchState("21IHPA0000A", "121.134.241.239:55770", 2)
	fmt.Println(off)
}

// 워치 업데이트
func TestWatchUpdate(t *testing.T) {
	result := client.WatchUpdate("21IHPA0000A", 100, 2)
	fmt.Println(result)
}
