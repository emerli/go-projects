package models

import (
	"encoding/json"
	"log"
	"net/http"
)

type BaseResp struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	ErrorCode string      `json:"errorCode,omitempty"`
}

func NewBaseResponseWOData() *BaseResp {
	return &BaseResp{
		Success: true,
		Message: "",
		Data:    nil,
	}
}

func NewBaseResponse(value interface{}) *BaseResp {
	return &BaseResp{
		Success: true,
		Message: "",
		Data:    value,
	}
}

func NewBaseResponseError(err error) *BaseResp {
	return &BaseResp{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}
}

func (r BaseResp) WriteTo(w http.ResponseWriter, status int) {
	data, _ := json.Marshal(r)

	w.WriteHeader(status)
	_, err := w.Write(data)

	if err != nil {
		log.Println("si è verificato un errore :", err.Error())
	}

}
