package elastic

import (
	"fmt"
	domain "github.com/gnnchya/PosCoffee/menu/domain"
	"strconv"
	"strings"
)
func InToStruct(r map[string]interface{}) []domain.InsertQ{
	var temp domain.InsertQ
	var result []domain.InsertQ
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		s := hit.(map[string]interface{})["_source"]
		temp.Name = fmt.Sprintf("%v", s.(map[string]interface{})["name"])
		temp.Category= fmt.Sprintf("%v", s.(map[string]interface{})["actual_name"])
		temp.Ingredient = strings.Split(fmt.Sprintf("%v", s.(map[string]interface{})["super_power"]),",")
		temp.Price = int64(s.(map[string]interface{})["birth_date"].(float64))
		temp.Available,_ = strconv.ParseBool(fmt.Sprintf("%v", s.(map[string]interface{})["alive"]))
		result = append(result, temp)
	}
	return result
}
