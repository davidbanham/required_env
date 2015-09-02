package required_env

import (
  "os"
  "log"
)

func Ensure(demands map[string]string) {
  for key, val := range demands {
    if os.Getenv(key) == "" {
      if val != "" {
        os.Setenv(key, val)
      } else {
        log.Fatal("You must specify the env variable ", key)
      }
    }
  }
}
