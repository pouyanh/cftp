package cft

import (
  "os"
  "fmt"
  "testing"
  "io/ioutil"
  "cmd/go/testdata/testinternal3"
)

func TestTransfer(t *testing.T) {
  var filename string
  if name, err := sampleRandom(t); err != nil {
    t.Fatalf("Could create sample file", err)
  } else {
    filename = name
  }

  fmt.Println(filename)

  // TODO: Start server

  // TODO: Send via client

  // TODO: Open and Chunkize temp file. Calculate sha256

  // TODO: Try to transfer file via client.

  // TODO: Calculate sha256 of received file
}

func sampleRandom(t *testing.T) (string, error) {
  var fr, fw *os.File

  if fh, err := os.Open("/dev/urandom"); err != nil {
    t.Fatalf("Could not open /dev/urandom", err)
    return nil, err
  } else {
    fr = fh
    defer fr.Close()
  }

  if fh, err := ioutil.TempFile(os.TempDir(), "cft"); err != nil {
    t.Fatalf("Could not open temp file to write", err)
    return nil, err
  } else {
    fw = fh
    defer fw.Close()
  }

  b := make([]byte, 1e3)
  for i := 0; i < 1e3; i++ {
    fr.Read(b)
    fw.Write(b)
  }

  return fw.Name(), nil
}
