package customquery

const MenuAll = `select kmp.startdate, kmp.finishdate, kd.name as dish_name, kd.cost as dish_cost, kd.weight as dish_weight, ke.name as eat_name, kdc.name as dish_category_name
				  	from k_menu
				  	inner join k_dish kd on kd.id = k_menu.dish_id
				  	inner join k_dish_category kdc on kdc.id = kd.category_id
				  	inner join k_eattype ke on ke.id = k_menu.eat_type_id
				  	inner join k_menu_property kmp on kmp.id = k_menu.menu_property_id
				 order by kmp.startdate desc;`
