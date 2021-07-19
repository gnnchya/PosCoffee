package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gnnchya/PosCoffee/menu/domain"
	"github.com/gnnchya/PosCoffee/menu/service/user/userin"
	"io"
	"log"
	"strings"
)

func (repo *Repository) Create(ctx context.Context, title *userin.CreateInput) error{
	out, err := json.Marshal(title)
	if err != nil {
		return err
	}

	var b strings.Builder
	b.WriteString(string(out))
	req := esapi.IndexRequest{
		Index:      repo.Index,
		DocumentID: title.ID,
		Body:       strings.NewReader(b.String()),
		Refresh:    "true",
	}
	res, err := req.Do(ctx, repo.Client)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	fmt.Println("in create repo", title)
	return err
}

func (repo *Repository)Update(ctx context.Context, title *domain.UpdateStruct) error{
	if state ,err := repo.CheckExistID(ctx, title.ID); err != nil{
		return err
	} else if state == false{
		return fmt.Errorf("this ID does not exist")
	}
	buf, err := BuildUpdateRequest(title)
	if err != nil {
		return err
	}
	res, err := repo.Client.Update(
		repo.Index, title.ID, &buf,
		repo.Client.Update.WithContext(ctx),
		repo.Client.Update.WithPretty())
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return err
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	return err
}

func (repo *Repository)Delete(ctx context.Context, id string) error{
	if state ,err := repo.CheckExistID(ctx, id); err != nil{
		return err
	} else if state == false{
		return fmt.Errorf("this ID does not exist")
	}
	req := esapi.DeleteRequest{
		Index:      repo.Index,
		DocumentID: id,
		Refresh:    "true",
	}

	res, err := req.Do(ctx, repo.Client)
	fmt.Println("delete  :",  res)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	return err
}