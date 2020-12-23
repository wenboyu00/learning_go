package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"learning_go/web-tutorial/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleCompanyCorrect(t *testing.T) {
	testBody := model.Company{ID: 123, Name: "google", Country: "USA"}
	testJson, _ := json.Marshal(testBody)
	reader := bytes.NewReader(testJson)
	r := httptest.NewRequest(http.MethodPost, "/companies", reader)
	w := httptest.NewRecorder() // 捕获响应记录响应
	handleCompany(w, r)
	result, _ := ioutil.ReadAll(w.Result().Body)
	c := model.Company{}
	json.Unmarshal(result, &c)
	if c.ID != 123 {
		t.Error("Failed to handle company correctly")
	}
}
