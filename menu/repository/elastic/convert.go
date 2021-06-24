package elastic

import (
	"fmt"
	domain "github.com/gnnchya/PosCoffee/menu/domain"
	"strconv"
	"strings"
)
func InToStruct(r map[string]interface{}) []domain.CreateStruct{
	var temp domain.CreateStruct
	var result []domain.CreateStruct
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		s := hit.(map[string]interface{})["_source"]
		temp.ID = fmt.Sprintf("%v", s.(map[string]interface{})["_id"])
		temp.Name = fmt.Sprintf("%v", s.(map[string]interface{})["name"])
		temp.Category= strings.Split(fmt.Sprintf("%v", s.(map[string]interface{})["category"]),",")
		temp.Ingredient = strings.Split(fmt.Sprintf("%v", s.(map[string]interface{})["ingredient"]),",")
		temp.Price = int64(s.(map[string]interface{})["price"].(float64))
		temp.Available,_ = strconv.ParseBool(fmt.Sprintf("%v", s.(map[string]interface{})["available"]))
		result = append(result, temp)
	}
	return result
}
