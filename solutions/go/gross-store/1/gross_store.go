package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	grossUnits := map[string] int{
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen" : 6,
        "dozen":12,
        "small_gross": 120,
        "gross":144,
        "great_gross":1728,
    }

    return grossUnits
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)

    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_,exists:=units[unit]
    if exists == false {
        return false
    }else{
        _,exists:=bill[item]
		if exists == false {
            bill[item] = units[unit]
        }else{
            bill[item] = bill[item]+units[unit]
        }
        return true
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    billQty,existsBill := bill[item]
	if existsBill == false{
        return false
    }
    
	unitQty,existsUnit := units[unit]
    if existsUnit == false {
        return false
    }
	newQty := billQty - unitQty
    
    if newQty < 0{
        return false
    }else if newQty == 0 {
		delete(bill,item)
        return true
    }else{
        bill[item] = newQty

        return true
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	billQty,existsBill :=  bill[item]

    if existsBill == false{
        return 0,false
    }
    return billQty,true
}
