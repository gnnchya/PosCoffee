package elastic

import (
	"bytes"
	"encoding/json"
	"github.com/gnnchya/PosCoffee/menu/domain"
)

func BuildUpdateRequest(t *domain.UpdateStruct) (buf bytes.Buffer, err error) {
	query := map[string]interface{}{
		"doc": map[string]interface{}{
			"name" : t.Name,
			"category" : t.Category,
			"ingredient" : t.Ingredient,
			"price" : t.Price,
			"available" : t.Available,
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return buf , err
	}
	return buf , err
}

func BuildCheckIndexRequest(index string)  (buf bytes.Buffer, err error) {
	query := map[string]interface{}{
		"query" : map[string]interface{}{
			"match": map[string]interface{}{
				"_index": index,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return buf , err
	}
	return buf , err
}

func BuildCheckIDRequest(id string)  (buf bytes.Buffer, err error) {
	query := map[string]interface{}{
		"query" : map[string]interface{}{
			"match": map[string]interface{}{
				"_id": id,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return buf , err
	}
	return buf , err
}

func BuildCheckNameRequest(name string)  (buf bytes.Buffer, err error) {
	query := map[string]interface{}{
		"query" : map[string]interface{}{
			"match": map[string]interface{}{
				"name": name,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return buf , err
	}
	return buf , err
}
