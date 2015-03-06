package genstr

import (
    "io"
    "crypto/rand"

)

var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")
var SimpleChars = []byte("ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz23456789-")
var Digits = []byte("0123456789")

func Standard(length int) string {
    return rand_char(length, StdChars)
}

func Simple(length int) string {
    return rand_char(length, SimpleChars)
}

func Number(length int) string {
    return rand_char(length, Digits)
}

func rand_char(length int, chars []byte) string {
    new_pword := make([]byte, length)
    random_data := make([]byte, length+(length/4)) // storage for random bytes.
    clen := byte(len(chars))
    maxrb := byte(256 - (256 % len(chars)))
    i := 0
    for {
        if _, err := io.ReadFull(rand.Reader, random_data); err != nil {
            panic(err)
        }
        for _, c := range random_data {
            if c >= maxrb {
                continue
            }
            new_pword[i] = chars[c%clen]
            i++
            if i == length {
                return string(new_pword)
            }
        }
    }
    panic("unreachable")
}

