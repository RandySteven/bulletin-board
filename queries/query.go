package queries

type (
	GoQuery        string
	TableMigration string
	DropTable      string
)

func (d DropTable) ToString() string {
	return string(d)
}

func (t TableMigration) ToString() string {
	return string(t)
}

func (q GoQuery) ToString() string {
	return string(q)
}
