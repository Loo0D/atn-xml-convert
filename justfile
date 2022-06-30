release:
   mkdir release
   go build -o release/atn-xml-convert-linux-x86_64
   GOOS=darwin go build -o release/atn-xml-convert-macos-x86_64
   GOOS=windows go build -o release/atn-xml-convert-windows-x86_64.exe
   GOOS=darwin GOARCH=arm64 go build -o release/atn-xml-convert-macos-arm64
   upx -9 release/atn-xml-convert-linux-x86_64
   upx -9 release/atn-xml-convert-macos-x86_64
   upx -9 release/atn-xml-convert-windows-x86_64.exe
   upx -9 release/atn-xml-convert-macos-arm64

clean:
    rm -rf release/
    rm -f atn-xml-convert