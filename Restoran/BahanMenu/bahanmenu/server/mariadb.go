package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addBahanMenu = `insert into tbBahanMenu(namabahan,quantity,createdBy,status,createdOn
		)values (?,?,?,?,?)`
	selectBahanMenu = `select IDBahan,namabahan,quantity,createdBy,status 
		from tbBahanMenu`
	updateBahanMenu = `update tbBahanMenu set namaBahan=?,quantity=?,status=?,UpdateBy=?,
		UpdateOn=? where IDBahan=?`
	selectNamaBahanMenu = `select IDBahan,namaBahan,quantity,createdBy,status 
		from tbBahanMenu where namaBahan=?`
)

type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}
func (rw *dbReadWriter) AddBahanMenu(bahanmenu BahanMenu) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addBahanMenu, bahanmenu.NamaBahan, bahanmenu.Quantity, bahanmenu.CreatedBy, OnAdd, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadBahanMenu() (BahanMenus, error) {
	fmt.Println("read all berhasil")
	bahanmenu := BahanMenus{}
	rows, _ := rw.db.Query(selectBahanMenu)
	defer rows.Close()
	for rows.Next() {
		var m BahanMenu
		err := rows.Scan(&m.IDBahan, &m.NamaBahan, &m.Quantity, &m.CreatedBy, &m.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return bahanmenu, err
		}
		bahanmenu = append(bahanmenu, m)
	}
	return bahanmenu, nil
}

func (rw *dbReadWriter) UpdateBahanMenu(m BahanMenu) error {
	fmt.Println("update berhasil")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateBahanMenu, m.NamaBahan, m.Quantity, m.Status, m.UpdateBy, time.Now(), m.IDBahan)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadBahanMenuByNamaBahanMenu(namaBahan string) (BahanMenu, error) {
	m := BahanMenu{NamaBahan: namaBahan}
	err := rw.db.QueryRow(selectNamaBahanMenu, namaBahan).Scan(&m.IDBahan, &m.NamaBahan, &m.Quantity, &m.CreatedBy, &m.Status)

	if err != nil {
		return BahanMenu{}, err
	}

	return m, nil
}
