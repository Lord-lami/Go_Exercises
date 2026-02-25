package lasagna

// TODO: define the 'PreparationTime()' function

// PreparationTime receives the slice of layers 
// and average preparation time
// calculates the total preparation time 
// of a lasagna.
func PreparationTime(layers []string, averagePreparationTime int) int {
	if averagePreparationTime == 0 {
		averagePreparationTime = 2
	}
	return len(layers) * averagePreparationTime
}

// TODO: define the 'Quantities()' function

// Quantites recieves a slice of layers
// and calcultes the quantity of 
// Noodles in grams and
// Sauce in liters needed.
func Quantities(layers []string) (weightOfNoodles int, volumeOfSauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			weightOfNoodles += 50
		} else if layer == "sauce" {
			volumeOfSauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function

// AddSecretIngredient takes the list of ingredients from your friend
// and your list of ingredients
// it changes the last element of your list ("?")
// to the last element of your friend's list
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function

// ScaleRecipe takes the slice of amounts needed to make two portions of lasagna 
// and the number of portions required 
// it returns a new slice of amounts adjusted for the required portions
func ScaleRecipe(basicAmounts []float64, requiredPortions int) (adjustedAmounts []float64) {
	for _, amount := range basicAmounts {
		adjustedAmounts = append(adjustedAmounts, (amount/2) * float64(requiredPortions))
	}
	return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
