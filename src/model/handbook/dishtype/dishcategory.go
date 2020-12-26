package dishtype

type DishType int

const (
	// Салаты
	Salad DishType = iota + 1
	// Первые блюда
	FirstMeal
	// Вторые блюда
	SecondMeal
	// Гарниры
	SideDish
	// Хлеб
	Bread
	// Выпечка
	Bake
	// Кондитерские изделия
	Confectionery
	// Торты
	Cakes
)

var categoryMap map[string]categoryDescription

type categoryDescription struct {
	dishCategory DishType
	description  string
}

func (dc DishType) String() string {
	s := [...]string{"Salad", "FirstMeal", "SecondMeal", "SideDish", "Bread", "Bake", "Confectionery", "Cakes"}[dc]
	return s
}

// getSupportedDishCategory get all dish category (Map {key - category name, value - category description} )
func getSupportedDishCategory() map[string]categoryDescription {
	if categoryMap == nil {
		categoryMap = make(map[string]categoryDescription, 0)

		categories := []categoryDescription{
			{dishCategory: Salad, description: "Салаты"},
			{dishCategory: FirstMeal, description: "Первые блюда"},
			{dishCategory: SecondMeal, description: "Вторые блюда"},
			{dishCategory: SideDish, description: "Гарниры"},
			{dishCategory: Bread, description: "Хлеб"},
			{dishCategory: Bake, description: "Выпечка"},
			{dishCategory: Confectionery, description: "Кондитерские изделия"},
			{dishCategory: Cakes, description: "Торты"},
		}
		for _, category := range categories {
			categoryMap[category.description] = category
		}
	}

	return categoryMap
}

// GetCategoryIndexByName return category instance by name
func GetCategoryIndexByName(name string) (index int) {
	category := getSupportedDishCategory()
	value, ok := category[name]
	if !ok {
		return -1
	}
	return int(value.dishCategory)
}
