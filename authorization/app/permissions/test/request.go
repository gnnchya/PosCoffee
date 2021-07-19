package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gnnchya/PosCoffee/authorize/service/permissions/permissionsin"
)

func (suite *PackageTestSuite) makeCreateReq(input *permissionsin.CreateInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return nil, nil, err
	}

	req, err = http.NewRequest("POST", "/permissions", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeCreateReqInvalidJSON() (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes := []byte("{{{}}}")

	req, err = http.NewRequest("POST", "/permissions", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeDeleteReq(input *permissionsin.DeleteInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	req, err = http.NewRequest("DELETE", fmt.Sprintf("/permissions/%s", input.ID), nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeListReq() (req *http.Request, w *httptest.ResponseRecorder, err error) {
	req, err = http.NewRequest("GET", "/permissions?page=1&per_page=10", nil)
	if err != nil {
		return nil, nil, err
	}
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeReadReq(input *permissionsin.ReadInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	req, err = http.NewRequest("GET", fmt.Sprintf("/permissions/%s", input.ID), nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeUpdateReq(input *permissionsin.UpdateInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return nil, nil, err
	}

	req, err = http.NewRequest("PUT", fmt.Sprintf("/permissions/%s", input.ID), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeUpdateReqInvalidJSON(input *permissionsin.UpdateInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes := []byte("{{{}}}")

	req, err = http.NewRequest("PUT", fmt.Sprintf("/permissions/%s", input.ID), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder(), nil
}
