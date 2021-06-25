package totalprice

//type Menu struct{
//	ID         		string   `bson:"_id"`
//	Category       	[]string   `bson:"category"`
//	Name 			string   `bson:"name"`
//	Ingredient 		[]string `bson:"ingredient"`
//	Price      		int64    `bson:"price"`
//	Available 		bool	 `bson:"available"`
//	Amount 			int64    `bson:"amount"`
//	Option 			string   `bson:"option"`
//}
//
//type CreateStruct struct {
//	ID 				string  `bson:"_id"`
//	CustomerID 		string  `bson:"customer_id"`
//	Menu			[]Menu	`bson:"menu"`
//	TotalPrice		int64	`bson:"total_price"`
//}

//func CalculateTotalPrice(cart *userin.Input) (T int64){
//	for _,i := range cart.Menu{
//		T = T+(i.Amount*i.Price)
//	}
//	return T
//}
//
//func main(){
//	var test CreateStruct
//	test.ID = "new"
//	test.CustomerID = "new customer"
//	test.Menu = []Menu{{"1",[]string{"Coffee", "Iced", "Dairy-free"},  "Iced Americano", []string{"Coffee beans", "Water", "Ice", "Plastic cup"}, 5500, true, 3,"-"},
//	{"2",[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Small)", []string{"Coffee beans", "Water", "Small hot cup"}, 3500, true,2,"-"},
//	{"3",[]string{"Coffee", "Hot", "Dairy-free"},  "Hot Americano (Large)", []string{"Coffee beans", "Water", "Large hot cup"}, 4500, true,4,"-"},
//	{"4",[]string{"Coffee", "Iced"},  "Iced Espresso", []string{"Coffee beans", "Water", "Milk", "Ice", "Plastic cup"}, 5500, true,1,"-"}}
//
//
//	fmt.Println(CalculateTotalPrice(test).TotalPrice)
//}
