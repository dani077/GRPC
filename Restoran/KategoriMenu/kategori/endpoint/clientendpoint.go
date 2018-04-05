package endpoint

import (
	"context"
	"fmt"

	sv "git.bluebird.id/Restoran/KategoriMenu/kategori/server"
)

func (me KategoriMenuEndpoint) AddKategoriMenuService(ctx context.Context, menu sv.KategoriMenu) error {
	_, err := me.AddKategoriMenuEndpoint(ctx, menu)
	return err
}

func (me KategoriMenuEndpoint) ReadKategoriMenuByKategoriMenuService(ctx context.Context, kategorimenu string) (sv.KategoriMenu, error) {
	req := sv.KategoriMenu{KategoriMenu: kategorimenu}
	//fmt.Println(req)
	resp, err := me.ReadKategoriMenuByKategoriMenuEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.KategoriMenu)
	return cus, err
}

func (me KategoriMenuEndpoint) ReadKategoriMenuService(ctx context.Context) (sv.KategoriMenus, error) {
	resp, err := me.ReadKategoriMenuEndpoint(ctx, nil)
	//fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.KategoriMenus), err
}

func (me KategoriMenuEndpoint) UpdateKategoriMenuService(ctx context.Context, cus sv.KategoriMenu) error {
	_, err := me.UpdateKategoriMenuEndpoint(ctx, cus)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
