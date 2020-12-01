package query

// CRUD query const
const (
	// Read
	IngridientAll = `select * from k_ingridient`
	// Create
	IngridientCreate = `insert into k_ingridient (
						created,
						updated, 
						uuid,
						name,
						weight
					) values ($1,$2,$3,$4,$5);`
	// Update
	IngridientUpdate = `
				update k_ingridient
					set updated = $1,
						name    = $2,
						weight  = $3
					where uuid = $4
				`
	// Delete
	IngridientDelete = `delete from k_ingridient where uuid = $1`
	// FindByUUID
	IngridientFindByUUID = `select * from k_ingridient where uuid = $1`
)
