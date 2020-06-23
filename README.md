# log

```go

func New(log *logrus.Logger) log.LoggerProvider {
    log.SetFormatter(new(TextFormatter))
    return &Provider{log: log}
}

type Provider struct {
    log *logrus.Logger
}

func (fac *Provider) Prefix(prefix string) log.Logger {
    return fac.log.WithField("prefix", prefix)
}

```

```go
var logProvider = New(logrus.New())

var log = logProvider.Prefix("Main")

log.Print("hello logger")
```
