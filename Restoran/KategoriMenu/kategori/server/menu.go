package server

import (
	"context"
)

type kategorimenu struct {
	writer ReadWriter
}

func NewKategoriMenu(writer ReadWriter) KategoriMenuService {
	return &kategorimenu{writer: writer}
}

func (m *kategorimenu) AddKategoriMenuService(ctx context.Context, kategorimenu KategoriMenu) error {

	err := m.writer.AddKategoriMenu(kategorimenu)
	if err != nil {
		return err
	}

	return nil
}
func (m *kategorimenu) ReadKategoriMenuService(ctx context.Context) (KategoriMenus, error) {
	men, err := m.writer.ReadKategoriMenu()
	if err != nil {
		return men, err
	}
	return men, nil
}

func (m *kategorimenu) UpdateKategoriMenuService(ctx context.Context, men KategoriMenu) error {
	err := m.writer.UpdateKategoriMenu(men)
	if err != nil {
		return err
	}
	return nil
}
func (m *kategorimenu) ReadKategoriMenuByKategoriMenuService(ctx context.Context, kategorimenu string) (KategoriMenu, error) {
	men, err := m.writer.ReadKategoriMenuByKategoriMenu(kategorimenu)
	if err != nil {
		return men, err
	}
	return men, nil
}
