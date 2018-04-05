package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "menu.bluebird.id"
	OnAdd     Status = 1
)

type KategoriMenu struct {
	IDKategori   int32
	KategoriMenu string
	Deskripsi    string
	Status       int32
	CreatedBy    string
	CreatedOn    string
	UpdateBy     string
	UpdateOn     string
}
type KategoriMenus []KategoriMenu

type ReadWriter interface {
	AddKategoriMenu(KategoriMenu) error
	ReadKategoriMenuByKategoriMenu(string) (KategoriMenu, error)
	ReadKategoriMenu() (KategoriMenus, error)
	UpdateKategoriMenu(KategoriMenu) error
}

type KategoriMenuService interface {
	AddKategoriMenuService(context.Context, KategoriMenu) error
	ReadKategoriMenuByKategoriMenuService(context.Context, string) (KategoriMenu, error)
	ReadKategoriMenuService(context.Context) (KategoriMenus, error)
	UpdateKategoriMenuService(context.Context, KategoriMenu) error
}
