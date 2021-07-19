package reportStock

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
