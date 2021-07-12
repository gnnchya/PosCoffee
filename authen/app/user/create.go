package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/app/view"
	"github.com/gnnchya/PosCoffee/authen/config"
	"github.com/gnnchya/PosCoffee/authen/domain"
	"github.com/gnnchya/PosCoffee/authen/service/user/userin"
	goxid "github.com/touchtechnologies-product/xid"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	AppJson = "application/json"
	AppFrom = "application/x-www-form-urlencoded"
)

func (ctrl *Controller) Create(c *gin.Context,role []string) {
	req, err := http.NewRequest("POST", config.Get().OauthServerURL, nil)
	if err != nil {
		return
	}
	req.Header.Set("Accept", AppJson)
	req.Header.Set("Content-Type", AppFrom)
	req.Header.Set("ClientID", config.Get().ClientId)
	req.Header.Set("ClientSecret", config.Get().ClientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	fmt.Println("response from oAuth", resp)
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.Status != "200 OK" {
		return
	}

	token := &domain.TokenView{}
	err = json.Unmarshal(body, token)
	if err != nil {
		return
	}

	input := &userin.CreateInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp2(c, 400, err)
		return
	}
	input.RoleID = role
	initID := goxid.New()
	input.ID = initID.Gen()
	initUID := goxid.New()
	input.UID = initUID.Gen()
	fmt.Println("Register", input)
	_, err = ctrl.service.Create(c, input)
	if err != nil {
		view.MakeErrResp2(c, 422, err)
		return
	}
	c.Set("Accept", AppJson)
	c.Set("Content-Type", AppFrom)
	c.Header("Authorization",token.AccessToken)
	c.Header("Refresh",token.RefreshToken)
	view.MakeSuccessResp(c, 200, "created")
}
