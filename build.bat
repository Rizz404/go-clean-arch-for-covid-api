@echo off
echo Building with CGO_ENABLED=1...
set CGO_ENABLED=1
go build -o ./tmp/main.exe ./app
