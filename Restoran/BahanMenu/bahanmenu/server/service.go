package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "menu.bluebird.id"
	OnAdd     Status = 1
)

type BahanMenu struct {
	IDBahan   int32
	NamaBahan string
	Quantity  int32

	Status    int32
	CreatedBy string
	CreatedOn string
	UpdateBy  string
	UpdateOn  string
}
type BahanMenus []BahanMenu

type ReadWriter interface {
	AddBahanMenu(BahanMenu) error
	ReadBahanMenuByNamaBahanMenu(string) (BahanMenu, error)
	ReadBahanMenu() (BahanMenus, error)
	UpdateBahanMenu(BahanMenu) error
}

type BahanMenuService interface {
	AddBahanMenuService(context.Context, BahanMenu) error
	ReadBahanMenuByNamaBahanMenuService(context.Context, string) (BahanMenu, error)
	ReadBahanMenuService(context.Context) (BahanMenus, error)
	UpdateBahanMenuService(context.Context, BahanMenu) error
}
