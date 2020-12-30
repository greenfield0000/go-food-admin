package customquery

const MenuPropertyCheckDateCollision = `
							select count(*)
							from k_menu_property kmp
							where $1 between kmp.startdate and kmp.finishdate or $2 between kmp.startdate and kmp.finishdate;
	`
