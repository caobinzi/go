#export GOPATH=`pwd`
rm -rf gen/*
genny -in validation/success.go -out gen/UserSuccess.go -pkg data gen "Generic=Data.User"
genny -in validation/success.go -out gen/IntSuccess.go -pkg data gen "Generic=int"
genny -in validation/success.go -out gen/StringSuccess.go -pkg data gen "Generic=string"
rm mytest
go build -o mytest
./mytest

