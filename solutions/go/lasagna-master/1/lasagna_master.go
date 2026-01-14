package lasagna

// PreparationTime calculates the total preparation time.
// Each layer takes timePerLayer minutes.
// If timePerLayer is 0, assume 2 minutes per layer.
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

// Quantities calculates the total amount of noodles (in grams)
// and sauce (in liters) needed based on the layers.
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		}
		if layer == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// AddSecretIngredient replaces the last ingredient in your list
// with the last ingredient from your friend's list.
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe scales the ingredient quantities for the given number of portions.
// The original recipe is for 2 portions.
func ScaleRecipe(recipe []float64, portions int) []float64 {
	factor := float64(portions) / 2.0
	scaled := make([]float64, len(recipe))

	for i, quantity := range recipe {
		scaled[i] = quantity * factor
	}

	return scaled
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
