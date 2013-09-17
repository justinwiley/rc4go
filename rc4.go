package main
import "fmt"

// func p(m variant) {
//   fmt.Println(m)
// }

// for i from 0 to 255
//     S[i] := i
// endfor
// j := 0
// for i from 0 to 255
//     j := (j + S[i] + key[i mod keylength]) mod 256
//     swap values of S[i] and S[j]
// endfor

func ksa(keystring string) {
  var key = []byte(keystring)
  var s [256]byte
  var j, newj, newi byte

  for i := 0; i < 255; i++ {
    s[i] = byte(i)
  }
  for i := 0; i < 255; i++ {
    j = (j + s[byte(i)] + key[i % len(key)]) % 255
    fmt.Println("i: ", i)
    fmt.Println("j: ", j)
    newi++
    newj++
    // newi = s[j]
    // newj = s[i]
    // s[i] = newj
    // s[j] = newi
  }
}

func main() {
  ksa("coconuts")
}

