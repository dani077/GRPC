package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addKategoriMenu = `insert into tbKategoriMenu(kategori,deskripsi,createdBy,status,createdOn
		)values (?,?,?,?,?)`
	selectKategoriMenu = `select IDkategoriMenu,kategori,deskripsi,createdBy,status 
		from tbKategoriMenu`
	updateKategoriMenu = `update tbKategoriMenu set kategori=?,deskripsi=?,UpdateBy=?,
		UpdateOn=? where IDkategoriMenu=?`
	selectKategori = `select IDKategoriMenu,kategori,deskripsi,createdBy,status 
		from tbKategoriMenu where kategori=?`
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
func (rw *dbReadWriter) AddKategoriMenu(menu KategoriMenu) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addKategoriMenu, menu.KategoriMenu, menu.Deskripsi, menu.CreatedBy, OnAdd, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadKategoriMenu() (KategoriMenus, error) {
	fmt.Println("read all berhasil")
	menu := KategoriMenus{}
	rows, _ := rw.db.Query(selectKategoriMenu)
	defer rows.Close()
	for rows.Next() {
		var m KategoriMenu
		err := rows.Scan(&m.IDKategori, &m.KategoriMenu, &m.Deskripsi, &m.CreatedBy, &m.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return menu, err
		}
		menu = append(menu, m)
	}
	return menu, nil
}

func (rw *dbReadWriter) UpdateKategoriMenu(m KategoriMenu) error {
	fmt.Println("update berhasil")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateKategoriMenu, m.KategoriMenu, m.Deskripsi, m.UpdateBy, time.Now(), m.IDKategori)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadKategoriMenuByKategoriMenu(kategorimenu string) (KategoriMenu, error) {
	m := KategoriMenu{KategoriMenu: kategorimenu}
	err := rw.db.QueryRow(selectKategori, kategorimenu).Scan(&m.IDKategori, &m.KategoriMenu, &m.Deskripsi, &m.CreatedBy, &m.Status)

	if err != nil {
		return KategoriMenu{}, err
	}

	return m, nil
}
