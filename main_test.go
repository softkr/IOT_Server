package main

import (
	"fmt"
	"iot/azure"
	"regexp"
	"strings"
	"testing"
)

func TestAzure(t *testing.T) {
	azure.Storage("test", "1222", "210816_205132_C9A73862A7EF.mp4")
}

func TestDate(t *testing.T) {
	date := "210816_205132_C9A73862A7EF.mp4"
	re := regexp.MustCompile(`(\d{2})(\d{2})(\d{2})`)
	date = re.ReplaceAllString(strings.Split(date, "_")[0], "20$1-$2-$3")
	fmt.Println(date, "++++")

}
