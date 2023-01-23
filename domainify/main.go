// Program usówa niedopuszczalne dla nazwy witryny znaki, zastępuje spacje minusami
// a następnie dodaje do utworzonej nazyw główną domenę (.com, .net, .io, .pl)

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var tlds = []string{"com", "net", "io", "pl"}

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  s := bufio.NewScanner(os.Stdin)
  for s.Scan() {
    text := strings.ToLower(s.Text())
    var newText []rune
    for _, r := range text {
      if unicode.IsSpace(r) {
        r = '-'
      }
      if !strings.ContainsRune(allowedChars, r) {
        continue
      }
      newText = append(newText, r)
    }
    fmt.Println(string(newText) + "." +
      tlds[rand.Intn(len(tlds))])
  }
}
