package scripts

import (
  "fmt"
  "strings"
)

func RequestFormValue(obj string, names []string) {

  for _, name := range names {
    fmt.Printf("  %s.%s = r.FormValue(\"%s\")\n", obj, name, strings.ToLower(name))
  }
  
}
