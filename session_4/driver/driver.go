package main

import "fmt"

type DBDriver struct {
}

type Row struct{}

func (d *DBDriver) Execute(sql string, vars []any) []Row {
	msg := d.Prepare(sql, vars)
	_ = msg
	// d.conn.send(msg)
	var out []byte
	// out := db.conn.recv()
	return d.Parse(out)
}

func (d *DBDriver) Prepare(sql string, vars []any) string {
	fmt.Println("DBDriver")
	return ""
}

func (d *DBDriver) Parse(data []byte) []Row {
	return nil
}

type PGDriver struct {
	DBDriver
}

func (d *PGDriver) Prepare(sql string, vars []any) string {
	fmt.Println("PGDriver")
	return ""
}

func (d *PGDriver) Parse(data []byte) []Row {
	return nil
}

func main() {
	drv := PGDriver{}
	drv.Execute("...", nil) // Will print DBDriver, not PGDriver
}
