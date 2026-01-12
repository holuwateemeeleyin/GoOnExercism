package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	//panic("Please implement the Units() function")
    return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	//panic("Please implement the NewBill() function")
    return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	//panic("Please implement the AddItem() function")
    unitValue, ok := units[unit]
	if !ok {
		return false
	}

	bill[item] += unitValue
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	//panic("Please implement the RemoveItem() function")
    unitValue, ok := units[unit]
	if !ok {
		return false
	}

	currentQty, exists := bill[item]
	if !exists || currentQty < unitValue {
		return false
	}

	if currentQty == unitValue {
		delete(bill, item)
	} else {
		bill[item] -= unitValue
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	//panic("Please implement the GetItem() function")
    qty, ok := bill[item]
	return qty, ok
}
