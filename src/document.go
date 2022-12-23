package document

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Doc struct {
	id           string
	Name         string
	Path         string
	creationTime time.Time
	sync.RWMutex
}

var ErrIsADirectory error = errors.New("path is a directory")
var ErrDoesNotExist error = errors.New("no file found in path")

// New returns a *Doc instance.
func New(path string) (*Doc, error) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, err
	}
	if stat.IsDir() {
		return nil, ErrIsADirectory
	}
	var d = Doc{
		id:           uuid.NewString(),
		Name:         stat.Name(),
		Path:         path,
		creationTime: time.Now(),
	}
	return &d, nil
}

// Base64 returns the base64 encoding of d content.
func (d *Doc) Base64() string {
	return base64.RawStdEncoding.EncodeToString(d.Content())
}

// Sha256 returns the hashed content of d.
func (d *Doc) Sha256() string {
	h := sha256.New()
	h.Write(d.Content())
	return hex.EncodeToString(h.Sum(nil))
}

// AbsoluteDir returns d path from root.
func (d *Doc) AbsoluteDir() string {
	abs, err := filepath.Abs(d.Path)
	if err != nil {
		panic(err)
	}
	return abs
}

// Dir returns d path from working dir.
func (d *Doc) Dir() string {
	return filepath.Dir(d.Path)
}

// Size returns the length of d content.
func (d *Doc) Size() int {
	return len(d.Content())
}

// Content returns the content of d.
func (d *Doc) Content() []byte {
	d.RLock()
	defer d.RUnlock()
	content, err := os.ReadFile(d.Path)
	if err != nil {
		panic(err)
	}
	return content
}
