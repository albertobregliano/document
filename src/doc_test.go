package document_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/albertobregliano/document"
)

func TestDoc_Content(t *testing.T) {
	path := t.TempDir() + "/test.txt"
	err := os.WriteFile(path, []byte{4, 5, 6}, 0600)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{4, 5, 6}
	fileTest, err := document.New(path)
	if err != nil {
		panic(err)
	}
	tests := []struct {
		name string
		d    *document.Doc
		want []byte
	}{
		{"test1", fileTest, want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Content(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Doc.Content() = %v, want %v", got, tt.want)
			}
		})
	}
}
