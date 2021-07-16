package implement

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/gnnchya/PosCoffee/authen/service/authentication/authenticationin"
	"github.com/gnnchya/PosCoffee/authen/service/authentication/out"
	"github.com/modern-go/reflect2"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (impl *implementation) GenerateToken(input *authenticationin.LoginInput) (token *out.Token, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return nil, err
	}

	userID, err := impl.login(input.Username, input.Password)
	if err != nil {
		return nil, err
	}
	fmt.Println("userId", userID)
	//hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	return nil, err
	//}
	token, err = impl.GetToken(userID, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (impl *implementation) login(username, password string) (userID string, err error) {
	users := &domain.UserStruct{}
	filters := impl.filter.MakeFilterUserName(username)
	err = impl.usersRepo.Read(context.Background(), filters, users)
	if err != nil {
		return "", err
	}

	if !users.Verify{return "", fmt.Errorf("error : account has not been verified yet")}
	if users.DeletedAt != 0{return "", fmt.Errorf("error : account has been deleted")}
	if users.Password != password {return "", err}
	return users.UID, nil
}

func (impl *implementation) GetToken(userID, username, password string) (token *out.Token, err error) {
	if reflect2.IsNil(impl.config.ClientId) || reflect2.IsNil(impl.config.ClientSecret) {
		return nil, err
	}

	formData := url.Values{
		"uid":       {userID},
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
	//
	//if resp.Status != "200 OK" {
	//	return nil, util.Unauthorized(errors.New("error"))
	//}

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
