package dish

const (
	Dish_create = `insert into k_dish (
						created,
						updated, 
						uuid,
						cost,
						name,
						picture,
						weight
					) values ($1,$2,$3,$4,$5,$6,$7);`
	Dish_All = `select * from k_dish`
)
