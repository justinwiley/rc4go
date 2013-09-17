    package main
    import "fmt"

    func ksa(keystring string) ([256]byte) {
      var key = []byte(keystring)

      var s [256]byte
      var j byte

      for i := 0; i <= 255; i++ {
        s[i] = byte(i)
      }
      for i := 0; i <= 255; i++ {
        j = (j + s[byte(i)] + key[i % len(key)]) % 255
        s[j], s[i] = s[i], s[j]
      }
      return s
    }

    func rc4(bufferstring string, key [256]byte) ([]byte) {
      buffer := []byte(bufferstring)
      res := make([]byte, len(buffer))
      var x, y, k, r byte = 0, 0, 0, 0

      for i, b := range(buffer) {
        x = (x + 1) % 255
        y = (y + key[x] % 255)
        y, x = x, y
        k = key[(key[x] + key[y]) % 255]
        r = b ^ k

        fmt.Printf("i: %d, b: %c, x: %#X, y: %#X, k: %#X, r: %#X  \n", i, b, x, y, k, r)
        res[i] = r
      }
      return res
    }

    func main() {
      keystring := "Secret"
      plaintext := "Attack at dawn"
      key := ksa(keystring)

      encrypted := rc4(plaintext, key)

      unencrypted := rc4(string(encrypted), key)

      fmt.Println("Encrypted")
      fmt.Printf("hex: %#X\n", string(encrypted))
      fmt.Printf("string: %s\n", string(encrypted))


      fmt.Println("Un-Encrypted")
      fmt.Printf("hex: %#X\n", string(unencrypted))
      fmt.Printf("string: %s\n", string(unencrypted))
    }

