package filechecksum

import (
   "crypto/md5"
   "encoding/hex"
   "fmt"
   "io"
   "os"
)

type MD5 struct {
}

func(md *MD5) FileMD5(filename string) (string, error) {
   file, err := os.Open(filename)
   if err != nil {
      return "", err
   }
   md5 := md5.New()
   if _, err = io.Copy(md5, file); err != nil {
      return "", err
   }
   md5Str := hex.EncodeToString(md5.Sum(nil))
   return md5Str, nil
}


func NewMD5CheckSum()(*MD5) {
   return &MD5 {
   }
}
