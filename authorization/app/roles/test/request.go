package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gnnchya/PosCoffee/authorize/service/roles/rolesin"
)

const ROUTE = "roles"

func setHeaderContentTypeJson(req *http.Request) {
	req.Header.Set("content-type", "application/json")
}

func (suite *PackageTestSuite) makeCreateReq(input *rolesin.CreateInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return nil, nil, err
	}

	req, err = http.NewRequest("POST", fmt.Sprintf("/%s", ROUTE), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	setHeaderContentTypeJson(req)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeCreateReqInvalidJSON() (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes := []byte("{{{}}}")

	req, err = http.NewRequest("POST", fmt.Sprintf("/%s", ROUTE), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	setHeaderContentTypeJson(req)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeReadReq(input *rolesin.ReadInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	req, err = http.NewRequest("GET", fmt.Sprintf("/%s/%s", ROUTE, input.ID), nil)
	if err != nil {
		return nil, nil, err
	}

	setHeaderContentTypeJson(req)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeUpdateReq(input *rolesin.UpdateInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return nil, nil, err
	}

	req, err = http.NewRequest("PUT", fmt.Sprintf("/%s/%s", ROUTE, input.ID), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	setHeaderContentTypeJson(req)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeUpdateReqInvalidJSON(input *rolesin.UpdateInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes := []byte("{{{}}}")

	req, err = http.NewRequest("PUT", fmt.Sprintf("/%s/%s", ROUTE, input.ID), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	setHeaderContentTypeJson(req)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeDeleteReq(input *rolesin.DeleteInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	req, err = http.NewRequest("DELETE", fmt.Sprintf("/%s/%s", ROUTE, input.ID), nil)
	if err != nil {
		return nil, nil, err
	}

	setHeaderContentTypeJson(req)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeListReq() (req *http.Request, w *httptest.ResponseRecorder, err error) {
	req, err = http.NewRequest("GET", fmt.Sprintf("/%s?page=1&per_page=10", ROUTE), nil)
	if err != nil {
		return nil, nil, err
	}

	setHeaderContentTypeJson(req)
	return req, httptest.NewRecorder(), nil
}
