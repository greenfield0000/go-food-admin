package crudquery

const MenuPropetyCreate = "insert into k_menu_property(uuid, created, updated, userid, startdate, finishdate) values ($1,$2,$3,$4,$5,$6) returning id"
