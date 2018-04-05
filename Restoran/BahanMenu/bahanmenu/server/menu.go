package server

import (
	"context"
)

type bahanmenu struct {
	writer ReadWriter
}

func NewBahanMenu(writer ReadWriter) BahanMenuService {
	return &bahanmenu{writer: writer}
}

func (m *bahanmenu) AddBahanMenuService(ctx context.Context, bahanmenu BahanMenu) error {

	err := m.writer.AddBahanMenu(bahanmenu)
	if err != nil {
		return err
	}

	return nil
}
func (m *bahanmenu) ReadBahanMenuService(ctx context.Context) (BahanMenus, error) {
	men, err := m.writer.ReadBahanMenu()
	if err != nil {
		return men, err
	}
	return men, nil
}

func (m *bahanmenu) UpdateBahanMenuService(ctx context.Context, men BahanMenu) error {
	err := m.writer.UpdateBahanMenu(men)
	if err != nil {
		return err
	}
	return nil
}
func (m *bahanmenu) ReadBahanMenuByNamaBahanMenuService(ctx context.Context, namabahanmenu string) (BahanMenu, error) {
	men, err := m.writer.ReadBahanMenuByNamaBahanMenu(namabahanmenu)
	if err != nil {
		return men, err
	}
	return men, nil
}
