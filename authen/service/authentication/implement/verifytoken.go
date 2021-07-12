package implement

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/service/authentication/out"
	"io"
	"io/ioutil"
	"net/http"
)

func (impl *implementation) VerifyToken(accessToken string) (userID *string, err error) {
	contentType := AppJson
	application := AppJson
	auth := fmt.Sprintf("Bearer %s", accessToken)
	req, err := http.NewRequest("POST", impl.config.ValidateTokenURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", application)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	fmt.Println("response from oAuth", resp)
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.Status != "200 OK" {
		return nil, errors.New("error")
	}

	resBody := &out.RepsValidate{}
	err = json.Unmarshal(body, resBody)
	if err != nil {
		return nil, err
	}

	if resBody.Data.UserId == "" {
		return userID, nil
	}

	userId := &out.ValidateToken{
		UserId:  resBody.Data.UserId,
		Expired: resBody.Data.Expired,
	}

	userID = &userId.UserId

	return userID, nil
}
