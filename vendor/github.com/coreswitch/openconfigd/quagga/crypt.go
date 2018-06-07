// Copyright 2018 OpenConfigd Project.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package quagga

import (
    "math/rand"
    "sync"
    "time"
    "unsafe"
)

/*
#include <crypt.h>
#include <stdlib.h>
#include <unistd.h>

#cgo LDFLAGS: -lcrypt
*/
import "C"

var (
    letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    passLetters = append(letters, []rune("!#$%&'()*+,-./:;<=>?@`{|}")...)
    initialized = false
)


func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func generateHash(key, salt string) string {
    var mutex = &sync.Mutex{}

    cKey := C.CString(key)
    defer C.free(unsafe.Pointer(cKey))
    cSalt := C.CString(salt)
    defer C.free(unsafe.Pointer(cSalt))
    mutex.Lock()
    cHash, _ := C.crypt(cKey, cSalt)
    mutex.Unlock()
    return C.GoString(cHash)
}

func Crypt(passwd string) string {
    return generateHash(passwd, randSeq(5))
}

func Verify(passwd, hash string) bool {
    if generateHash(passwd, hash) == hash {
        return true
    }
    return false
}

func GeneratePasswd() string {
    if initialized == false {
        rand.Seed(time.Now().UnixNano())
        initialized = true
    }
    b := make([]rune, 8)
    for i := range b {
        b[i] = passLetters[rand.Intn(len(passLetters))]
    }
    return string(b)
}
