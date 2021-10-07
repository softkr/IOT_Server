package api

import (
"bytes"
"encoding/json"
"fmt"
"io/ioutil"
"net/http"
)


type T struct {
	ReturnId          string `json:"return_id"`
	ReturnFilename    string `json:"return_filename"`
	ReturnProjectName string `json:"return_project_name"`
	ReturnMedication  bool `json:"return_medication"`
}

func Drug(drug *T) {
	pbytes, _ := json.Marshal(drug)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("https://medication.inhandplus.com/api/v1/drug", "application/json", buff)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

