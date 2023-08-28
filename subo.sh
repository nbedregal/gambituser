git add .
git commit -m "Último commit"
git push
export GOOS = linux
export GOARCH = amd64
go build main.go
rm main.zip
zip main.zip main