package loggers

type Logger interface {
	Printf(format string, v ...any)
	Println(v ...any)

	Fatal(v ...any)
	Fatalf(format string, v ...any)
}
