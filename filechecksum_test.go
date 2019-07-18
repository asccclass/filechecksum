package filechecksum

import (
   "testing"
)

func TestFileMD5(t *testing.T) {
   file := "test.txt"
   wanted := "77d59390e1a44feaf14b9093df8021db"

   m := NewMD5CheckSum()
   checksum, _ := m.FileMD5(file)

   if wanted != checksum {
      t.Errorf("Got %s but want %s.", checksum, file)
   }
}
