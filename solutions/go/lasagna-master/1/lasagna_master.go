package lasagna

func PreparationTime(layers []string, time int) int {
	//panic("Please implement the PreparationTime function")
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// func Quantities(layers []string) (noodles int, sauce float64) {
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
	//return
}

func AddSecretIngredient(friendsList, myList []string) {
	lastItem := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = lastItem
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, nrOfPortions int) []float64 {
	scaledRecipe := []float64{}

	for _, amount := range amounts {
		scaledRecipe = append(scaledRecipe, amount*float64(nrOfPortions)/2)
	}
	return scaledRecipe
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
