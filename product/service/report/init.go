package report

func InitTemp()[]string{
	var temp []string
	temp = append(temp, "ID")
	temp = append(temp, "Cart")
	temp = append(temp, "CustomerID")
	temp = append(temp, "Finished")
	temp = append(temp, "TotalCost")
	temp = append(temp, "TotalPrice")
	temp = append(temp, "PaymentMethod")
	temp = append(temp, "TypeOfOrder")
	temp = append(temp, "Destination.Coordinates")
	temp = append(temp, "Time")
	temp = append(temp, "MenuID")
	temp = append(temp, "MenuName")
	temp = append(temp, "MenuPrice")
	temp = append(temp, "MenuAmount")
	temp = append(temp, "MenuOption")
	return temp
}

func InitStock()[]string{
	var temp []string
	temp = append(temp,"ID")
	temp = append(temp,	"ItemName")
	temp = append(temp, "Category")
	temp = append(temp, "Amount")
	temp = append(temp, "Unit")
	temp = append(temp, "CostPerUnit")
	temp = append(temp, "EXPDate")
	temp = append(temp, "ImportDate")
	temp = append(temp, "Supplier")
	temp = append(temp, "TotalCost")
	temp = append(temp, "TotalAmount")
	temp = append(temp, "Status")
	return temp
}

func InitCal()[]string{
	var temp []string
	temp = append(temp, "Total Income")
	temp = append(temp, "Total Cost")
	temp = append(temp, "Total Profit")
	return temp
}