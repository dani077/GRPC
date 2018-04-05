package endpoint

import (
	"context"
	"fmt"

	sv "git.bluebird.id/Restoran/BahanMenu/bahanmenu/server"
)

func (me BahanMenuEndpoint) AddBahanMenuService(ctx context.Context, bahanmenu sv.BahanMenu) error {
	fmt.Println("masuk sini juga")
	_, err := me.AddBahanMenuEndpoint(ctx, bahanmenu)
	return err
}

func (me BahanMenuEndpoint) ReadBahanMenuByNamaBahanMenuService(ctx context.Context, namabahanmenu string) (sv.BahanMenu, error) {
	req := sv.BahanMenu{NamaBahan: namabahanmenu}
	fmt.Println(req)
	resp, err := me.ReadBahanMenuByNamaBahanMenuEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.BahanMenu)
	return cus, err
}

func (me BahanMenuEndpoint) ReadBahanMenuService(ctx context.Context) (sv.BahanMenus, error) {
	resp, err := me.ReadBahanMenuEndpoint(ctx, nil)
	//fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.BahanMenus), err
}

func (me BahanMenuEndpoint) UpdateBahanMenuService(ctx context.Context, cus sv.BahanMenu) error {
	_, err := me.UpdateBahanMenuEndpoint(ctx, cus)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
