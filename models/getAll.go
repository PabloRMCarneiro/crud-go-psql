package models

import (
	"crud/db"
)

func GetAll() (todos []Todo, err error){
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row, err  := conn.Query(`SELECT * FROM todos WHERE`)
	
	if err != nil {
		return
	}
	
	for row.Next() {
		var todo Todo

		err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		
		if err != nil {
			continue
		}
		
		todos = append(todos, todo)
	}
	return 
}