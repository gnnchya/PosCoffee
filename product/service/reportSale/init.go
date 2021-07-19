package reportSale

func initSale()[]string{
	var temp []string
	temp = append(temp, "ID")
	temp = append(temp, "Name")
	temp = append(temp, "Price")
	temp = append(temp, "Sold Amount")
	temp = append(temp, "Total Income")
	return temp
}
