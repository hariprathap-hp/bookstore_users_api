package domain_user

import (
	"fmt"
	"strings"

	"github.com/hariprathap-hp/bookstore_users_api/src/data_resource/dbPostgres"
	"github.com/hariprathap-hp/bookstore_users_api/src/utils/errors"
)

const (
	insertintoDBQuery = "select id, first_name, last_name, email, created_at, status from users where id=$1"
	createDBQuery     = "insert into users (first_name, last_name, email, created_at, status, password) values ($1,$2,$3,$4,$5,$6) returning id"
	updateDBQuery     = "update users set first_name=$1, last_name=$2, email=$3 where id=$4"
	deleteDBQuey      = "delete from users where id=$1"
	searchDBQuery     = "select id, first_name, last_name, email, created_at, status from users where status=$1"
	loginQuery        = "select id, first_name, last_name, email, created_at, status from users where email=$1 AND password=$2 AND status=$3"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := dbPostgres.Client.Prepare(insertintoDBQuery)
	if err != nil {
		return errors.NewHTTPInternalServerError("db statement preparation failed")
	}
	defer stmt.Close()
	fmt.Println(user.Id)
	if getErr := stmt.QueryRow(user.Id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), "no rows in result set") {
			return errors.NewHTTPInternalServerError("no row with user_id found in database")
		}
		fmt.Println(getErr.Error())
		return errors.NewHTTPInternalServerError("executing query on database failed")
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := dbPostgres.Client.Prepare(createDBQuery)
	if err != nil {
		return errors.NewHTTPInternalServerError("db statement preparation failed")
	}
	defer stmt.Close()
	if saveErr := stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password).Scan(&user.Id); saveErr != nil {
		fmt.Println(saveErr)
		return errors.NewHTTPInternalServerError("creating a new user in database failed")
	}
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := dbPostgres.Client.Prepare(updateDBQuery)
	if err != nil {
		return errors.NewHTTPInternalServerError("db statement preparation failed")
	}
	defer stmt.Close()
	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if updateErr != nil {
		fmt.Println(updateErr)
		return errors.NewHTTPInternalServerError("user updation in db failed")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := dbPostgres.Client.Prepare(deleteDBQuey)
	if err != nil {
		return errors.NewHTTPInternalServerError("db statement preparation failed")
	}
	defer stmt.Close()
	_, delErr := stmt.Exec(user.Id)
	if delErr != nil {
		fmt.Println(delErr)
		return errors.NewHTTPInternalServerError("user deletion from db failed")
	}
	return nil
}

func (user *User) Search(status string) ([]User, *errors.RestErr) {
	stmt, err := dbPostgres.Client.Prepare(searchDBQuery)
	if err != nil {
		return nil, errors.NewHTTPInternalServerError("db statement preparation failed")
	}
	defer stmt.Close()
	fmt.Println(status)
	rows, searchErr := stmt.Query(status)
	if searchErr != nil {
		fmt.Println(searchErr)
		return nil, errors.NewHTTPInternalServerError("fetching users from database failed")
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		scanErr := rows.Scan(&user.Id, &user.FirstName,
			&user.LastName, &user.Email,
			&user.DateCreated, &user.Status)
		if scanErr != nil {
			return nil, errors.NewHTTPInternalServerError("failed during scanning result rows")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewHTTPNotFoundError("no users matching status")
	}
	return results, nil
}

func (user *User) Login() *errors.RestErr {
	stmt, err := dbPostgres.Client.Prepare(loginQuery)
	if err != nil {
		return errors.NewHTTPInternalServerError("db statement preparation failed")
	}
	defer stmt.Close()
	fmt.Println(user.Email)
	fmt.Println(user.Password)
	fmt.Println(user.Status)
	if getErr := stmt.QueryRow(user.Email, user.Password, StatusActive).Scan(&user.Id, &user.FirstName,
		&user.LastName, &user.Email,
		&user.DateCreated, &user.Status); getErr != nil {
		fmt.Println(getErr)
		return errors.NewHTTPInternalServerError("login by username and password failed")
	}
	fmt.Println(user)
	return nil
}
