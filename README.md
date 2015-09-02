## What?

It demands that a given list of environment variables be set. Optionally, you can specify defaults for them if they're not listed.

## How?
```go
import (
  "github.com/davidbanham/required_env"
)
required_env.Ensure(map[string]string{"SOMETHING_IMPORTANT": "", "PORT": "3000"})
```

Put the above in demo.go and:

`go run demo.go`
Will blow up with the error: "You must specify the env variable SOMETHING_IMPORTANT"

`SOMETHING_IMPORTANT=foo go run demo.go`
Will get you a happily running process in which os.Getenv("PORT") is equal to "3000"

`PORT=80 SOMETHING_IMPORTANT=foo go run demo.go`
Will get you a happily running process in which os.Getenv("PORT") is equal to "80"

## Why?
Sometimes when running unfamiliar code it's not obvious what environment variables the author expected to have set. This makes it clearer.

Sometimes when writing code, you want to allow a default option to be overriden by an environment variable. It's tedious to write that code over and over again, though.
