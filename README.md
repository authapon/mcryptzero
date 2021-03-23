# mcryptZero
Encryption and Decryption with rotate bit randomize.

*This project is developed for personal used - You can use it by your own risk and purpose !!!*

**Example Usage**

```go
package main

import (
	"fmt"
	"github.com/authapon/mcryptzero"
)

func main() {
	key := []byte("Secret Key")
	data := []byte("This is data...")
	encrypt := mcryptzero.Encrypt(data, key)
	decrypt := mcryptzero.Decrypt(encrypt, key)
	
	fmt.Printf("Key = %s\nData = %s\nEncrypt = %v\nDecrypt = %s\n", key, data, encrypt, decrypt)
}
```
