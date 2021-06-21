package elastic

import (
	"fmt"
	domain "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
	"strconv"
	"strings"
)
func InToStruct(r map[string]interface{}) []domain.InsertStruct{
	var temp domain.InsertStruct
	var result []domain.InsertStruct
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		s := hit.(map[string]interface{})["_source"]
		temp.Name = fmt.Sprintf("%v", s.(map[string]interface{})["name"])
		temp.ActualName = fmt.Sprintf("%v", s.(map[string]interface{})["actual_name"])
		temp.ActualLastName = fmt.Sprintf("%v", s.(map[string]interface{})["actual_lastname"])
		temp.Gender = fmt.Sprintf("%v", s.(map[string]interface{})["gender"])
		temp.BirthDate = int64(s.(map[string]interface{})["birth_date"].(float64))
		temp.Height,_ = strconv.Atoi(fmt.Sprintf("%v", s.(map[string]interface{})["height"]))
		temp.SuperPower = strings.Split(fmt.Sprintf("%v", s.(map[string]interface{})["super_power"]),",")
		temp.Alive,_ = strconv.ParseBool(fmt.Sprintf("%v", s.(map[string]interface{})["alive"]))
		temp.Universe = fmt.Sprintf("%v", s.(map[string]interface{})["universe"])
		temp.Movies = strings.Split(fmt.Sprintf("%v", s.(map[string]interface{})["movies"]),",")
		temp.Enemies = strings.Split(fmt.Sprintf("%v", s.(map[string]interface{})["enemies"]),",")
		temp.FamilyMember = strings.Split(fmt.Sprintf("%v", s.(map[string]interface{})["family_member"]),",")
		temp.About = fmt.Sprintf("%v", s.(map[string]interface{})["about"])
		result = append(result, temp)
	}
	return result
}
