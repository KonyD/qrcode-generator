# qrcode-generator
qrcode generator CLI written in Go

**Flags:**
```bash
  -filename string
        Name of the png file. (default "qrcode.png")
  -size int
        Size of the png file. (default 256)
  -url string
        Url that will be converted to qrcode. (default "https://go.dev/")
```

**Example usage:**
```bash
go run main.go -url=https://go.dev/ -size=256 -filename=qrcode.png
```

**Building the project:**
```bash
go build main.go
```
