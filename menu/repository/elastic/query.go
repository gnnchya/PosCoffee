package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func (repo *Repository)SearchCategory(keyword string,ctx context.Context)([]domain.CreateStruct, error){
	page,size,err := repo.CheckPagination(ctx , buildCategoryRequest(1,10,keyword))
	var result []domain.CreateStruct
	for i := 1; i < page; i++{
		q, _ := repo.query(ctx,buildCategoryRequest(i,size,keyword))
		res := InToStruct(q)
		for _,x := range res{
			result = append(result,x)
		}
	}
	return result, err
}

func (repo *Repository)SearchIngredient(keyword string,ctx context.Context)([]domain.CreateStruct, error){
	page,size,err := repo.CheckPagination(ctx , buildIngredientRequest(1,10,keyword))
	var result []domain.CreateStruct
	for i := 1; i < page; i++{
		q, _ := repo.query(ctx,buildIngredientRequest(i,size,keyword))
		res := InToStruct(q)
		for _,x := range res{
			result = append(result,x)
		}
	}
	return result, err
}

func (repo *Repository)SearchMenu(keyword string,ctx context.Context)([]domain.CreateStruct, error){
	page,size,err := repo.CheckPagination(ctx , buildMenuRequest(1,10,keyword))
	var result []domain.CreateStruct
	for i := 1; i < page; i++{
		q, _ := repo.query(ctx,buildMenuRequest(i,size,keyword))
		fmt.Println("q")
		res := InToStruct(q)
		for _,x := range res{
			result = append(result,x)
		}
	}
	return result, err
}

func (repo *Repository)Read(id string,ctx context.Context)(result domain.CreateStruct, err error){
	if found , err := repo.CheckExistID(ctx, id); found == false{
		return result , fmt.Errorf("error : ID does not exist")
	} else if err != nil{
		return result , err
	}
	q, err := repo.query(ctx,buildViewRequest(id))
	fmt.Println(q)
	results := InToStruct(q)
	fmt.Println(results[0])
	return results[0], err
}

func (repo *Repository)ReadAll(page int, size int,ctx context.Context)([]domain.CreateStruct, error){
	q, err := repo.query(ctx,buildViewAllRequest(page, size,repo.Index))
	result := InToStruct(q)
	return result, err
}