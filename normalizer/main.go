package main

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "anomy"
	password = "secret"
	dbname   = "normalizer"
)

var numbers = []string{
	"1234567890",
	"123 456 7891",
	"(123) 456 7892",
	"(123) 456-7893",
	"123-456-7894",
	"1234567892",
	"(123)456-7892",
}

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	must(err)
	createPhoneNumberTable(db)

	// for _, number := range numbers {
	// 	id, err := insertPhone(db, number)
	// 	must(err)
	// 	fmt.Printf("id=%d\n", id)
	// }

	phones, err := allPhones(db)
	must(err)
	for _, phone := range phones {
		fmt.Printf("Working on ... %+v\n", phone)
		number := normalize(phone.number)
		if number != phone.number {
			fmt.Println("Updating or removing ...", number)
			existing, err := findPhone(db, number)
			must(err)
			if existing != nil {
				must(deletePhone(db, phone.id))
			} else {
				phone.number = number
				must(updatePhone(db, phone))
			}
		} else {
			fmt.Println("No changes required!")
		}
	}

	defer db.Close()
	must(db.Ping())
}

type phone struct {
	id     int
	number string
}

func allPhones(db *sql.DB) ([]phone, error) {
	rows, err := db.Query("select id,number from phone_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []phone

	for rows.Next() {
		var p phone
		if err := rows.Scan(&p.id, &p.number); err != nil {
			return nil, err
		}
		ret = append(ret, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}

func updatePhone(db *sql.DB, p phone) error {
	q := `update phone_numbers set number=$2 where id=$1`
	_, err := db.Exec(q, p.id, p.number)
	return err
}

func deletePhone(db *sql.DB, id int) error {
	q := `delete from phone_numbers where id=$1`
	_, err := db.Exec(q, id)
	return err
}

func findPhone(db *sql.DB, number string) (*phone, error) {
	var p phone
	row := db.QueryRow("select * from phone_numbers where number=$1", number)
	err := row.Scan(&p.id, &p.number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &p, nil
}

func getPhone(db *sql.DB, id int) (string, error) {
	var number string
	row := db.QueryRow("select * from phone_numbers where id=$1", id)
	err := row.Scan(&id, &number)
	if err != nil {
		return "", err
	}

	return number, nil
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	var id int
	q := `insert into phone_numbers(number) values($1) returning id`
	err := db.QueryRow(q, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func createPhoneNumberTable(db *sql.DB) {
	q := `create table if not exists phone_numbers (
		id 		 serial primary key,
  	number VARCHAR(255) not null
	)`
	_, err := db.Exec(q)
	must(err)
}

func normalize(phone string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(phone, "")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// func normalize(phone string) string {
// 	var buf bytes.Buffer

// 	for _, ch := range phone {
// 		if ch >= '0' && ch <= '9' {
// 			buf.WriteRune(ch)
// 		}
// 	}

// 	return buf.String()
// }
