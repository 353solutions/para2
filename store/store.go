package store

type Item struct {
	SKU   string
	Name  string
	Price int // In ¢
}

type Store struct {
	db *DB
}

func (s *Store) GetItem(sku string) (Item, error) {
	return s.db.Get(sku)
}

func (s *Store) Update(i Item) error {
	return s.db.Update(i)
}
