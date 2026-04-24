package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	favoriteCards := []int {2,6,9}

    return favoriteCards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	lengthSlice := len(slice)

    if index > lengthSlice-1 || index < 0{
        return -1
    }
	value := slice[index]
    
    return value
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	lengthSlice := len(slice)
    
    newSlice := slice[:]

    if index > lengthSlice-1 || index < 0 || index == -1 {
        newSlice = append(slice,value)
    }else{
         newSlice[index] = value
    }

    return newSlice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	

    fullSlice := slice[:]
    if len(values) == 0{
    	return slice
    } else {
        fullSlice = append(values, slice...)
    }

    return fullSlice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	lengthSlice := len(slice)

    var fullSlice []int

    if index > lengthSlice-1 || index < 0{
        fullSlice = slice[:]
    }else{
    
        fullSlice = append(slice[:index],slice[index+1:]...)
    }


	return fullSlice















    
}
