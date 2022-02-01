package main

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Cover struct {
	Id         int
	Name       string
	CreateTime string
	UpdateTime string
}

// SQLX version
func GetCoversX() ([]Cover, error) {
	query := "select id, name from cover"
	var covers []Cover

	err := db.Select(&covers, query)
	if err != nil {
		return nil, err
	}
	return covers, nil
}

func GetCoverX(id int) (*Cover, error) {
	query := "SELECT id, name FROM cover WHERE id=?"
	var cover Cover
	err := db.Get(&cover, query, id)
	if err != nil {
		return nil, err
	}
	return &cover, nil
}

func GetCovers() ([]Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select id, create_time, update_time, name from cover"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var covers []Cover
	for rows.Next() {
		currentCover := Cover{
			Id:         0,
			Name:       "",
			CreateTime: "",
			UpdateTime: "",
		}
		err := rows.Scan(&currentCover.Id, &currentCover.CreateTime, &currentCover.UpdateTime, &currentCover.Name)
		if err != nil {
			return nil, err
		}
		covers = append(covers, currentCover)
	}
	return covers, nil
}

func GetCover(id int) (*Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	// MySQL
	query := "SELECT id, name FROM cover where id=?"
	row := db.QueryRow(query, id)
	var cover Cover
	err = row.Scan(&cover.Id, &cover.Name)
	if err != nil {
		return nil, err
	}
	return &cover, nil
}

func AddCover(cover Cover) error {
	insertCommand := "INSERT INTO cover (id, name, create_time, update_time) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(insertCommand, cover.Id, cover.Name, cover.CreateTime, cover.UpdateTime)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("Can not insert command(AddCover)")
	}
	return nil
}

func UpdateCover(cover Cover) error {
	updateCommand := "UPDATE cover SET name=? WHERE id=?"
	result, err := db.Exec(updateCommand, cover.Name, cover.Id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("Can not update command(UpdateCover)")
	}
	return nil
}

func DeleteCover(id int) error {
	deleteCommand := "DELETE FROM cover WHERE id=?"
	result, err := db.Exec(deleteCommand, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("Can not delete command(DeleteCover)")
	}
	return nil
}

func main() {
	var err error
	db, err = sqlx.Open("mysql", "Arima_kishou0:My_secrete_passw0rd@tcp(localhost:3306)/techcoach")
	if err != nil {
		panic(err)
	}

	// newCover := Cover{8, "Bond8", "2022-01-23 17:33:45.987", "2022-01-23 17:33:45.987"}
	// err = AddCover(newCover)
	// if err != nil {
	// 	panic(err)
	// }

	// covers, err := GetCoversX()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, cover := range covers {
	// 	fmt.Println(cover)
	// }
	// updateCover := Cover{Id: 8, Name: "John"}
	// err = UpdateCover(updateCover)
	// if err != nil {
	// 	panic(err)
	// }

	// err = DeleteCover(9)
	// if err != nil {
	// 	panic(err)
	// }

	cover, err := GetCoverX(8)
	if err != nil {
		panic(err)
	}
	fmt.Println(cover)
}
