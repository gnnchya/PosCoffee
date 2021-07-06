package implement

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/gnnchya/PosCoffee/authen/service/authentication/autenitcationin"
	"github.com/gnnchya/PosCoffee/authen/service/authentication/out"
	"github.com/gnnchya/PosCoffee/authen/service/util"
	"github.com/modern-go/reflect2"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (impl *implementation) GenerateToken(input *authenitcationin.LoginInput) (token *out.Token, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return nil, util.ValidationCreateErr(err)
	}

	userID, err := impl.login(input.Username, input.Password)
	if err != nil {
		return nil, err
	}
	token, err = impl.getToken(userID, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (impl *implementation) login(username, password string) (userID string, err error) {
	users := &domain.Users{}
	filters := MakeFilterEmailORMobileNumber(impl.filter, username)
	err = impl.usersRepo.Read(context.Background(), filters, users)
	if err != nil {
		return "", util.Unauthorized(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	if err != nil {
		return "", util.Unauthorized(err)
	}

	return users.ID, nil
}

func (impl *implementation) getToken(userID, username, password string) (token *out.Token, err error) {
	if reflect2.IsNil(impl.config.ClientId) || reflect2.IsNil(impl.config.ClientSecret) {
		return nil, err
	}

	formData := url.Values{
		"user_id":       {userID},
		"client_id":     {impl.config.ClientId},
		"client_secret": {impl.config.ClientSecret},
		"redirect_uri":  {impl.config.RedirectUri},
		"grant_type":    {impl.config.GrantType},
		"username":      {username},
		"password":      {password},
	}

	contentType := AppFrom
	application := AppJson
	data := strings.NewReader(formData.Encode())

	req, err := http.NewRequest("POST", impl.config.OauthServerURL, data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", application)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.Status != "200 OK" {
		return nil, util.Unauthorized(errors.New("error"))
	}

	resBody := &out.RespBody{}
	err = json.Unmarshal(body, resBody)
	if err != nil {
		return nil, err
	}
	token = &out.Token{
		AccessToken:  resBody.Data.AccessToken,
		RefreshToken: resBody.Data.RefreshToken,
		Expired:      resBody.Data.Expired,
	}

	return token, nil
}
