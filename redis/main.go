package main

import (
	"fmt"

	"encoding/json"
	"github.com/go-redis/redis"
)

type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type Ingredient struct{
	IngredientName    string   `bson:"ingredient_name" json:"ingredient-name"`
	Amount      		int64    `bson:"amount" json:"amount"`
}

type CreateStruct struct {
	ID         		string   `bson:"_id" json:"id"`
	Category       	[]string   `bson:"category" json:"category" validate:"required"`
	Name 			string   `bson:"name" json:"name" validate:"required"`
	Ingredient 		[]Ingredient `bson:"ingredient" json:"ingredient" validate:"required"`
	Price      		int64    `bson:"price" json:"price" validate:"required"`
	Available 		bool	 `bson:"available" json:"available" validate:"required"`
}

func main() {
	fmt.Println("Go Redis Tutorial")
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})


	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Printf("Cannot Ping: %v\n", err.Error())
	} else {
		fmt.Printf("Pong: %v\n", pong)
	}

	json, err := json.Marshal(CreateStruct{"ffkhjvkjg", []string{"Coffee", "Iced", "Dairy-free"},  "Iced Americano", []Ingredient{Ingredient{"Coffee beans", 004}, {"Water", 002}, {"Ice", 00025}, {"Plastic cup", 100}}, 5500, true})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get("id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}