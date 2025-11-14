package botmiscfunctions

import (
	"math/rand"
	"slices"
)

type PropertiesOfElement struct {
	NameDesc  string
	TitleDesc string
}

// Берёт рандомные значения из категории
func RandomizeSlice(sliceCategory []PropertiesOfElement, displayRange int) []PropertiesOfElement {
	newSlice := []PropertiesOfElement{}
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
