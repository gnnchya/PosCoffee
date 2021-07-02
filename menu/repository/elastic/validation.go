package elastic

import (
	"bytes"
	"context"
)

func (repo *Repository) CheckExistID(ctx context.Context, id string) (bool, error) {
	buf, err := BuildCheckIDRequest(id)
	if err != nil{
		return false, err
	}
	result, err := repo.query(ctx,buf)
	found := int((result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]).(float64))
	if found == 0 {
		return false, nil
	}
	return true, err
}

func (repo *Repository) CheckExistName(ctx context.Context, name string) (bool, error) {
	buf, err := BuildCheckNameRequest(name)
	if err != nil{
		return false, err
	}
	result, err := repo.query(ctx,buf)
	found := int((result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]).(float64))
	if found == 0 {
		return false, nil
	}
	return true, err
}

func (repo *Repository) CheckExistIndex(ctx context.Context, Index string) (bool, error) {
	buf, err := BuildCheckIndexRequest(Index)
	if err != nil{
		return false, err
	}
	result, err := repo.query(ctx,buf)
	found := int((result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]).(float64))
	if found == 0 {
		return false, nil
	}
	return true, err
}

func (repo *Repository) CheckPagination(ctx context.Context, buf bytes.Buffer) (page int,size int, err error) {
	result, err := repo.query(ctx,buf)
	value := int((result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]).(float64))
	if value > 10 {
		return value+9/10, 10, nil
	}
	return 0,10,nil
}
