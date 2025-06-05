package store

import (
	"path"
	"testing"
)

func BenchmarkStore_Get(b *testing.B) {
	db, err := Open(path.Join(b.TempDir(), "store.db"))
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	i := Item{
		SKU:   "pen88",
		Name:  "Space Pen",
		Price: 301,
	}
	err = db.Update(i)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for b.Loop() {
		i2, err := db.Get(i.SKU)
		if err != nil {
			b.Fatal(err)
		}
		if i2.SKU != i.SKU {
			b.Fatal(i2)
		}
	}
}
