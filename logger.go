package queue

type EmptyLogger struct{}

func (l *EmptyLogger) Print(v ...interface{})                 {}
func (l *EmptyLogger) Printf(format string, v ...interface{}) {}
func (l *EmptyLogger) Println(v ...interface{})               {}

func (l *EmptyLogger) Fatal(v ...interface{})                 {}
func (l *EmptyLogger) Fatalf(format string, v ...interface{}) {}
func (l *EmptyLogger) Fatalln(v ...interface{})               {}

func (l *EmptyLogger) Panic(v ...interface{})                 {}
func (l *EmptyLogger) Panicf(format string, v ...interface{}) {}
func (l *EmptyLogger) Panicln(v ...interface{})               {}
