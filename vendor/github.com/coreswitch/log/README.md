# log

log is a wrapper for [logrus](https://github.com/sirupsen/logrus/) Go logging package. To help debugging, log added
source code information and function name to logrus output.

#### Example

The simple example for enabling debug level for DEBUG and set output to stdout.

```go
package main

import (
  log "github.com/coreswitch/log"
)

func main() {
  log.SetLevel("debug")
  log.SetOutput("stdout")

  log.With("animal", "walrus").Info("A walrus appears")
}
```

The example set formatter to JSON.

```go
package main

import (
  log "github.com/coreswitch/log"
)

func main() {
  log.SetLevel("debug")
  log.SetJSONFormatter()

  log.With("animal", "zebra").Debug("A zebra appears")
}
```

Source field and function field is configurable.

```go
package main

import (
  log "github.com/coreswitch/log"
)

func main() {
  log.SourceField = false
  log.FuncField = false

  log.With("animal", "bird").Warn("A crow appears")
}
```
