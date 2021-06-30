package reportStock

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
