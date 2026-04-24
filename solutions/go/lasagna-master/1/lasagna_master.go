package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutePerLayer int)int{
    if minutePerLayer == 0 {
        return len(layers) * 2
    }else{
        return len(layers) * minutePerLayer
    }
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string)(int,float64){
    var sauceNoodle int
    var sauce float64
    for i := 0 ; i < len(layers); i ++{
        if layers[i] == "noodles"{
            sauceNoodle += 50
        }else if layers[i] == "sauce"{
            sauce += 0.2
        }
    }

    return sauceNoodle,sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string,myList []string){
    myListLastIndex := len(myList)-1
    friendListLastIndex := len(friendList)-1
    myList[myListLastIndex] = friendList[friendListLastIndex ]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64,portions int) []float64{
	var temp float64

    newRecipe := make([]float64, len(quantities))
    copy(newRecipe, quantities)
    
    for i := 0 ; i < len(newRecipe); i++{
        temp = newRecipe[i] / 2
        newRecipe[i] = temp * float64(portions)
    }

    return newRecipe
}



// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
