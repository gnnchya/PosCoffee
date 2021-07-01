package report

func InitTemp()[]string{
	var temp []string
	temp = append(temp, "ID")
	temp = append(temp, "Cart ID")
	temp = append(temp, "Customer ID")
	temp = append(temp, "Finished Status")
	temp = append(temp, "Total Cost")
	temp = append(temp, "Total Price")
	temp = append(temp, "Payment Method")
	temp = append(temp, "Type of Order")
	temp = append(temp, "Destination Coordinates")
	temp = append(temp, "Time")
	temp = append(temp, "Menu ID")
	temp = append(temp, "Menu Name")
	temp = append(temp, "Menu Price")
	temp = append(temp, "Menu Amount")
	temp = append(temp, "Menu Option")
	return temp
}

func InitStock()[]string{
	var temp []string
	temp = append(temp,"ID")
	temp = append(temp,	"Item Name")
	temp = append(temp, "Category")
	temp = append(temp, "Amount")
	temp = append(temp, "Unit")
	temp = append(temp, "Cost per Unit")
	temp = append(temp, "Expired Date")
	temp = append(temp, "Import Date")
	temp = append(temp, "Supplier")
	temp = append(temp, "Total Cost")
	temp = append(temp, "Total Amount")
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