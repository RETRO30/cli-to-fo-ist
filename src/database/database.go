package database

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
)

type row struct {
	title   string
	checked bool
}

type DatabaseObj struct {
	rows []row
}

func NewRow(title string, checked bool) row {
	return row{title, checked}
}

func NewDatabaseObj() (DatabaseObj, error) {
	var EmptyDatabaseObj DatabaseObj
	file_name := "database.csv"
	is_creted := false
	file, err := os.Open(file_name)
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(file_name)
			is_creted = true
			if err != nil {
				return EmptyDatabaseObj, errors.New("Error creating file")
			}
		} else {
			return EmptyDatabaseObj, errors.New("Error opening file")
		}
	}
	defer file.Close()
	rows := []row{}

	if is_creted {
		file.WriteString("id,title,checked\n")
	} else {
		reader := csv.NewReader(file)
		titles, err := reader.Read()
		if err != nil {
			return EmptyDatabaseObj, errors.New("Error reading file")
		}
		if titles[0] != "id" || titles[1] != "title" || titles[2] != "checked" {
			return EmptyDatabaseObj, errors.New("Error reading file")
		}
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return EmptyDatabaseObj, errors.New("Error reading file")
			}
			title := record[1]
			checked, err := strconv.ParseBool(record[2])
			if err != nil {
				return EmptyDatabaseObj, errors.New("Error reading file")
			}
			rows = append(rows, NewRow(title, checked))
		}
	}
	return DatabaseObj{rows}, nil
}

func (db DatabaseObj) Add(title string) {
	db.rows = append(db.rows, NewRow(title, false))
}

func (db DatabaseObj) Delete(id int) {
	db.rows = append(db.rows[:id], db.rows[id+1:]...)
}

func (db DatabaseObj) Edit(id int, title string, checked bool) {
	db.rows[id].title = title
	db.rows[id].checked = checked
}

func (db DatabaseObj) Get(id int) row {
	return db.rows[id]
}

func (db DatabaseObj) GetAll() []row {
	return db.rows
}

func (db DatabaseObj) save() {
	file_name := "database.csv"
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("id,title,checked\n")
	for i, row := range db.rows {
		file.WriteString(strconv.Itoa(i) + "," + row.title + "," + strconv.FormatBool(row.checked) + "\n")
	}
}
