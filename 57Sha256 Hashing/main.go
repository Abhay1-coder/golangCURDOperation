/**
SHA-256 (Secure Hash Algorithm 256-bit) is a type of mathematical function that takes an input (like a piece of text or data) and processes it in a way that produces a fixed-size string of characters, typically in hexadecimal format. This string of characters is called a "hash."

Think of a hash like a unique fingerprint for a piece of data. It's a way to represent the data in a condensed and standardized form. No matter how large or small the input data is, the SHA-256 algorithm will always produce a 64-character hexadecimal hash.

Here's a simple explanation of SHA-256 hashes:

1. **Consistent Output**: If you feed the same input data into the SHA-256 algorithm, it will always produce the same hash. It's like taking a photograph of the data, and that photograph (the hash) will look the same every time you take it.

2. **Unique for Different Data**: Even a small change in the input data will result in a significantly different hash. So, two slightly different pieces of data will have completely different SHA-256 hashes.

3. **Fixed Size**: Regardless of the size or length of the input data, the SHA-256 hash will always have a fixed length of 64 characters.

4. **One-Way Function**: It's very easy to compute the hash from data, but it's computationally infeasible to reverse the process and obtain the original data from the hash. This property makes it suitable for storing passwords securely.

5. **Used for Data Integrity**: SHA-256 hashes are commonly used to verify the integrity of data. For example, when you download a file from the internet, you might also find its corresponding SHA-256 hash. After downloading the file, you can compute its hash and compare it to the provided hash to check if the file was tampered with during the download.

In essence, SHA-256 is a tool that helps ensure data hasn't been altered or corrupted, and it's widely used in cryptography, security, and data verification applications.
**/

// [_SHA256 hashes_](https://en.wikipedia.org/wiki/SHA-2) are
// frequently used to compute short identities for binary
// or text blobs. For example, TLS/SSL certificates use SHA256
// to compute a certificate's signature. Here's how to compute
// SHA256 hashes in Go.

package main

// Go implements several hash functions in various
// `crypto/*` packages.
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// Here we start with a new hash.
	h := sha256.New() //has is initilized

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	h.Write([]byte(s)) //s is feed to the hasing object

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil) //this is comverting into hash code. nil indicated that no other things to be appended here

	fmt.Println(s)
	fmt.Printf("%x\n", bs) //%x print in hexadecimal number
}
