package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/gnnchya/PosCoffee/menu/domain"

)

func (repo *Repository)query(ctx context.Context,buf bytes.Buffer) (map[string]interface{}, error){
	es := repo.Client
	var r  map[string]interface{}
	res, err := es.Search(
		es.Search.WithContext(ctx),
		es.Search.WithIndex(repo.Index),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	return r, err
}

func (repo *Repository)SearchCategory(keyword string,ctx context.Context)([]domain.InsertQ, error){
	q, err := repo.query(ctx,buildCategoryRequest(keyword))
	result := InToStruct(q)
	return result, err
}

func (repo *Repository)SearchIngredient(keyword string,ctx context.Context)([]domain.InsertQ, error){
	q, err := repo.query(ctx,buildIngredientRequest(keyword))
	result := InToStruct(q)
	return result, err
}

func (repo *Repository)SearchMenu(keyword string,ctx context.Context)([]domain.InsertQ, error){
	q, err := repo.query(ctx,buildMenuRequest(keyword))
	result := InToStruct(q)
	return result, err
}

func (repo *Repository)Read(id string,ctx context.Context)([]domain.InsertQ, error){
	q, err := repo.query(ctx,buildViewRequest(id))
	result := InToStruct(q)
	return result, err
}

func (repo *Repository)ReadAll(page int, size int,ctx context.Context)([]domain.InsertQ, error){
	q, err := repo.query(ctx,buildViewAllRequest(page, size,repo.Index))
	result := InToStruct(q)
	return result, err
}