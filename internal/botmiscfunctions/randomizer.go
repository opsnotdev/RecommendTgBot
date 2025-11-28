package botmiscfunctions

import (
	"math/rand"
	"recommendtgbot/internal/commonvariables"
	"slices"
)

// Берёт рандомные значения из категории
func RandomizeSlice(sliceCategory []commonvariables.PropertiesOfElement, displayRange int) []commonvariables.PropertiesOfElement {
	newSlice := []commonvariables.PropertiesOfElement{}
	if displayRange > len(sliceCategory) {
		displayRange = len(sliceCategory)
	}
	for len(newSlice) < displayRange {
		randomNumber := rand.Intn(len(sliceCategory))
		if !slices.Contains(newSlice, sliceCategory[randomNumber]) {
			newSlice = append(newSlice, sliceCategory[randomNumber])
		}
	}
	return newSlice
}
