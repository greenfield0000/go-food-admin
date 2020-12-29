package crudquery

// CRUD query const
const (
	// Read
	DishAll = `select * from k_dish`
	// Create
	DishCreate = `insert into k_dish (
						created,
						updated, 
						uuid,
						cost,
						name,
						picture,
						weight
					) values ($1,$2,$3,$4,$5,$6,$7);`
	// Update
	DishUpdate = `
				update k_dish
					set updated = $1,
						cost    = $2,
						name    = $3,
						picture = $4,
						weight  = $5
					where uuid = $6
				`
	// Delete
	DishDelete = `delete from k_dish where uuid = $1`
	// FindByUUID
	DishFindByUUID = `select * from k_dish where uuid = $1`
)
