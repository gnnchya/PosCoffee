package elastic

import (
	"bytes"
	"encoding/json"
	"log"
)

func buildMenuRequest(page int, size int,keyword string) bytes.Buffer {
	var buf bytes.Buffer
	from := (page-1)*size
	query := map[string]interface{}{
		"from": from,
		"size": size,
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query" : "*"+keyword+"*",
				"fields" : []interface{}{
					"name",
				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
	log.Fatalf("Error encoding query: %s", err)
	}
	return buf
}

func buildViewAllRequest(page int, size int, index string) bytes.Buffer{
	var buf bytes.Buffer
	from := (page-1)*size
	query := map[string]interface{}{
		"from": from,
		"size": size,
		"query" : map[string]interface{}{
			"match": map[string]interface{}{
				"_index": index,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	return buf
}

func buildViewRequest(id string) bytes.Buffer{
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query" : map[string]interface{}{
			"match": map[string]interface{}{
				"_id": id,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	return buf
}

func buildCategoryRequest(page int, size int,keyword string) bytes.Buffer {
	var buf bytes.Buffer
	from := (page-1)*size
	query := map[string]interface{}{
		"from": from,
		"size": size,
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query" : "*"+keyword+"*",
				"fields" : []interface{}{
					"category",
				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	return buf
}

func buildIngredientRequest(page int, size int,keyword string) bytes.Buffer {
	var buf bytes.Buffer
	from := (page-1)*size
	query := map[string]interface{}{
		"from": from,
		"size": size,
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query" : "*"+keyword+"*",
				"fields" : []interface{}{
					"ingredient.item_name",
				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	return buf
}