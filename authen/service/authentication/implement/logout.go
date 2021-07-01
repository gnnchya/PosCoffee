package implement

import (
	"errors"
	"fmt"
	"net/http"
)

func (impl *implementation) Logout(token string) (err error) {
	_, err = impl.revokeToken(token)
	if err != nil {
		return err
	}
	return
}

func (impl *implementation) revokeToken(accessToken string) (token bool, err error) {
	contentType := AppJson
	application := AppJson
	auth := fmt.Sprintf("Bearer %s", accessToken)
	req, err := http.NewRequest("DELETE", impl.config.RevokeTokenURL, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Accept", application)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return false, errors.New("error delete token")

	}
	return true, nil
}
