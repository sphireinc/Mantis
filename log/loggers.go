package log

// expand our string slice into a string separated by <>
func expand(additional []string) string {
	data := ""
	for _, datum := range additional {
		data = data + "<>" + datum
	}
	return data
}

// Info writes an info level log message
func (l *Log) Info(msg string, additional ...string) {
	l.writer(INFO, msg+" "+expand(additional))
}

// Warn writes a warning level log message
func (l *Log) Warn(msg string, additional ...string) {
	l.writer(WARN, msg+" "+expand(additional))
}

// Error writes an error level log message
func (l *Log) Error(msg string, additional ...string) {
	l.writer(ERROR, msg+" "+expand(additional))
}

// Fatal writes a fatal level log message, then os.Exits
func (l *Log) Fatal(msg string, additional ...string) {
	l.writer(FATAL, msg+" "+expand(additional))
}

// Panic writes a panic level log message then panics
func (l *Log) Panic(msg string, additional ...string) {
	l.writer(PANIC, msg+" "+expand(additional))
}
