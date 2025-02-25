package rpchttp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"social_todo_app_go/module/product/query"
	"time"
)

type rpcGetCategoryById struct {
	url string
}

func NewRpcGetCategoryById(url string) *rpcGetCategoryById {
	return &rpcGetCategoryById{url: url}
}

func (rpc *rpcGetCategoryById) FindWithIds(ctx context.Context, ids []uuid.UUID) ([]query.CategoryDTO, error) {
	method := "GET"
	payloadData := struct {
		Ids []uuid.UUID `json:"ids"`
	}{
		Ids: ids,
	}
	payloadByte, _ := json.Marshal(payloadData)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, rpc.url, bytes.NewReader(payloadByte))

	if err != nil {
		log.Println(err)
		return nil, errors.New("cannot get product category")
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, errors.New("cannot get product category")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("cannot get product category")
	}

	var responseData struct {
		Data []query.CategoryDTO `json:"data"`
	}
	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, errors.New("cannot get product category")
	}
	return responseData.Data, nil
}
