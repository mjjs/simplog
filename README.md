# Simplog - The simple logging library
Simplog is a logging library for Golang. The project was done purely as a learning exercise and not designed to be used outside small pet projects, so running Simplog is not advisable in production environments.

The library was inspired by Golang's own [log/syslog](https://godoc.org/log/syslog), but it operates directly on files instead of communicating with syslog.
Simplog is intended to be a simple, small library so it won't bloat your executable size.

___

## Usage
The main part of Simplog is the `Logger` type, which represents a single logging instance. Each instance operates on its own, dedicated logfile. The simple way to write into the log is to call the logger's Write method, which writes a new log entry with the severity level which was chosen when creating the logger.

Writing single messages with other severity levels is possible using the corresponding logger methods.

```go
// Error handling left out for simplicity
logger, _ := simplog.New("my_log.log", simplog.INFO)

// Close the file handle opened by the logger instance
defer logger.Close();

// Writes an event to using the severity level the logger was constructed with
logger.Write("Some event happened!")

// You can override the severity level for single messages
logger.Crit("This is a critical event")
logger.Debug("A debug entry!")
```

## License
Simplog is licensed under the permissive [MIT license](LICENSE.md).
