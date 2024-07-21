package csvreader

type Reader interface {
	Read() (record []string, err error)
}
