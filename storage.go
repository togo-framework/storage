// Package storage is togo's default filesystem storage provider. Blank-import
// (or `togo install togo-framework/storage`) to register it with the kernel.
package storage

import (
	"os"
	"path/filepath"

	"github.com/togo-framework/togo"
	tstorage "github.com/togo-framework/togo/storage"
)

func init() {
	togo.RegisterProviderFunc("storage", togo.PriorityService, func(k *togo.Kernel) error {
		k.Storage = NewFS(k.Config.StorageDir)
		return nil
	})
}

type fsStore struct{ root string }

// NewFS returns a filesystem-backed store rooted at root.
func NewFS(root string) tstorage.Storage { return &fsStore{root: root} }

func (s *fsStore) full(p string) string { return filepath.Join(s.root, filepath.Clean("/"+p)) }

func (s *fsStore) Put(path string, data []byte) error {
	full := s.full(path)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		return err
	}
	return os.WriteFile(full, data, 0o644)
}

func (s *fsStore) Get(path string) ([]byte, error) { return os.ReadFile(s.full(path)) }
func (s *fsStore) Delete(path string) error        { return os.Remove(s.full(path)) }
func (s *fsStore) Path(path string) string         { return s.full(path) }
