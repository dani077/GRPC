package endpoint

import (
	"context"
	"fmt"

	sv "git.bluebird.id/Restoran/Menu/menu/server"
)

func (me MenuEndpoint) AddMenuService(ctx context.Context, menu sv.Menu) error {
	_, err := me.AddMenuEndpoint(ctx, menu)
	return err
}

func (me MenuEndpoint) ReadMenuByNamaMenuService(ctx context.Context, namamenu string) (sv.Menu, error) {
	req := sv.Menu{NamaMenu: namamenu}
	fmt.Println(req)
	resp, err := me.ReadMenuByNamaMenuEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Menu)
	return cus, err
}

func (me MenuEndpoint) ReadMenuService(ctx context.Context) (sv.Menus, error) {
	resp, err := me.ReadMenuEndpoint(ctx, nil)
	//fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Menus), err
}

func (me MenuEndpoint) UpdateMenuService(ctx context.Context, cus sv.Menu) error {
	_, err := me.UpdateMenuEndpoint(ctx, cus)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (me MenuEndpoint) ReadMenuByKeteranganMenuService(ctx context.Context, keterangan string) (sv.Menus, error) {
	req := sv.Menu{Keterangan: keterangan}
	resp, err := me.ReadMenuByKeteranganMenuEndpoint(ctx, req)
	//fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Menus), err
}
