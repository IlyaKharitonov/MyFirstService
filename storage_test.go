package main

import (
	"database/sql"
	"testing"
)

//подключение к базе
func strDB() (Storage, error) {
	dbtest, err := sql.Open("mysql", "root:1643@(0.0.0.0:3306)/usersdb")
	if err != nil {
		panic(err)
	}
	storage := Storage{
		database: dbtest,
	}
	return storage, err
}

func TestAdd(t *testing.T) {
	//
	storage, err := strDB()
	if err != nil {
		t.Fatal("Database connection error ", err)
	}
	//данные для  тестового добавления
	usertest := Data{
		Name: "Arnold",
		Age:  25,
	}
	//структура, в которую будем писать последние добавненные данные для сверки с оригиналом
	fromDB := Data{}
	//запускаем тестируемый метод
	err = storage.Add(usertest.Name, usertest.Age)
	if err != nil {
		t.Fatal("Database connection error ", err)
	}
	//достаем из базы последнюю добавленную строку и помещаем в fromDB
	rows, err := storage.database.Query("SELECT * FROM usersdb.users WHERE id=LAST_INSERT_ID();")
	if err != nil {
		t.Fatal("The request failed", err)
	}
	for rows.Next() {
		err = rows.Scan(&fromDB.ID, &fromDB.Name, &fromDB.Age)
		if err != nil {
			t.Fatal("Line scan error", err)
		}
	}
	//сравниваем ожидаемый результат
	if usertest.Name != fromDB.Name || usertest.Age != fromDB.Age {
		t.Errorf("Latest data does not match")
	}
}

func TestUpdate(t *testing.T) {

	storage, err := strDB()
	if err != nil {
		t.Fatal("Database connection error  ", err)
	}
	fromDB := Data{}
	//узнаем последний id  по которому будем менять запись
	id := storage.database.QueryRow("SELECT MAX(`id`) FROM users")
	err = id.Scan(&fromDB.ID)
	if err != nil {
		t.Fatal("Error scan result error SELECT MAX(`id`) FROM users ", err)
	}
	//данные для обновления
	userUpdate := Data{
		ID:   fromDB.ID,
		Name: "Borat",
		Age:  68,
	}
	//обновление данных в таблице
	err = storage.Update(userUpdate.ID, userUpdate.Name, userUpdate.Age)
	if err != nil {
		t.Fatal("Error executing the test request", err)
	}
	//получение этих последних обновленных данных
	rowAfter := storage.database.QueryRow("select * from usersdb.users where id=?", userUpdate.ID)
	err = rowAfter.Scan(&fromDB.ID, &fromDB.Name, &fromDB.Age)
	if err != nil {
		t.Fatal("Result scan error select * from usersdb.users where id=? ", err)
	}
	//сверяем измененную строку с данными для обновления
	if fromDB != userUpdate {
		t.Errorf("The data does not match")
	}
}

func TestCount(t *testing.T) {

	storage, err := strDB()
	if err != nil {
		t.Fatal("Database connection error", err)
	}
	fromDB := Data{}
	id := storage.database.QueryRow("select count(*) from usersdb.users")
	err = id.Scan(&fromDB.ID)
	if err != nil {
		t.Fatal("Result scan error select count(*) from usersdb.users ", err)
	}
	result, err := storage.Count()
	if err != nil {
		t.Fatal("Error storage.Count()", err)
	}
	if result != fromDB.ID {
		t.Errorf("The counters did not match")
	}
}

func TestGet(t *testing.T) {

	storage, err := strDB()
	if err != nil {
		t.Fatal("Database connection error", err)
	}
	id := storage.database.QueryRow("SELECT MAX(`id`) FROM users")
	fromDB := Data{}
	err = id.Scan(&fromDB.ID)
	if err != nil {
		t.Fatal("Result scan error SELECT MAX(`id`) FROM users ", err)
	}
	// ожидаемое значение по запросу
	testUser := Data{
		ID:   fromDB.ID,
		Name: "Borat",
		Age:  68,
	}
	user, err := storage.Get(fromDB.ID)
	if err != nil {
		t.Fatal("Error storage.Get", err)
	}
	if user != testUser {
		t.Errorf("The data did not match")
	}
}

func TestGetByName(t *testing.T) {

	//получим количество записей. тк все записи в таблице одинаковые name: "Borat" age: 68
	//то слайс данных будет иметь длинну равную значению count
	storage, err := strDB()
	if err != nil {
		t.Fatal("Database connection error", err)
	}
	testUser := Data{
		Name: "Khabib",
		Age:  45,
	}
	for i := 0; i < 3; i++ {
		storage.Add(testUser.Name, testUser.Age)
	}
	users, err := storage.GetByName(testUser.Name)
	if err != nil {
		t.Fatal("The request failed storage.GetByName", err)
	}
	if users[0].Name != testUser.Name {
		t.Errorf("Expected and actual did not match")
	}
}

func TestGetByAge(t *testing.T) {

	storage, err := strDB()
	if err != nil {
		t.Fatal("Database connection error", err)
	}
	testUser := Data{
		Name: "Abdul",
		Age:  45,
	}
	for i := 0; i < 3; i++ {
		storage.Add(testUser.Name, testUser.Age)
	}
	//ожидаемое значение из d.id структур
	users, err := storage.GetByAge(testUser.Age)
	if err != nil {
		t.Fatal("The request failed storage.GetByAge", err)
	}
	if users[0].Age != testUser.Age {
		t.Errorf("Expected and actual did not match")
	}
}
