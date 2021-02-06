package main

import (
	"database/sql"
)

type Storage struct {
	database *sql.DB
}

type Storager interface {
	Add(name string, age int) (err error)
	Update(id int, newName string, newAge int) (err error)
	Count() (count int, err error)
	Get(id int) (user Data, err error)
	GetByName(name string) (users []Data, err error)
	GetByAge(age int) (users []Data, err error)
}

type Data struct {
	ID   int
	Name string
	Age  int
}

// Добавляет запись в базу
func (s Storage) Add(name string, age int) (err error) {
	_, err = s.database.Exec("insert into usersdb.users(name,age) values(?,?)", name, age)
	return
}

// обновляет данные записи по id
func (s Storage) Update(id int, newName string, newAge int) (err error) {
	_, err = s.database.Exec("update usersdb.users set name = ?, age = ? where id = ?", newName, newAge, id)
	return
}

// показать количество записей в хранилище
func (s Storage) Count() (count int, err error) {
	err = s.database.QueryRow("select count(*) from usersdb.users").Scan(&count)
	return
}

// Получить запись по id
func (s Storage) Get(id int) (user Data, err error) {
	err = s.database.QueryRow("select * from usersdb.users where id=?", id).Scan(&user.ID, &user.Name, &user.Age)
	return
}

// Получение записи по имени пользователя
func (s Storage) GetByName(name string) (users []Data, err error) {
	rows, err := s.database.Query("select * from usersdb.users where name=?", name)
	if err != nil {
		return nil, err
	}
	// users := []Data{}
	for rows.Next() {
		d := Data{}
		err = rows.Scan(&d.ID, &d.Name, &d.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, d)
	}
	return users, err
}

// Получение записей по возрасту
func (s Storage) GetByAge(age int) (users []Data, err error) {
	rows, err := s.database.Query("select * from usersdb.users where age=?", age)
	if err != nil {
		return nil, err
	}
	users = []Data{}
	for rows.Next() {
		d := Data{}
		err = rows.Scan(&d.ID, &d.Name, &d.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, d)
	}
	return users, err
}



