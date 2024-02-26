package helpers


import (
    "io"
    "os"
    "bytes"
)

func ReadFromFile(fileName string) ([]byte, error) {
    inFile, err := os.Open(fileName)
    if err != nil {
        return nil, err
    }
    defer inFile.Close()

    var buf bytes.Buffer
    _, err = io.Copy(&buf, inFile)
    if err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}
