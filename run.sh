#export GOPATH=`pwd`
rm -rf gen/*
genny -in validation/success.go -out data/UserSuccess.go -pkg data gen "Generic=User"
genny -in validation/success.go -out data/IntSuccess.go -pkg data gen "Generic=int"
genny -in validation/success.go -out data/StringSuccess.go -pkg data gen "Generic=string"
rm mytest
go build -o mytest
./mytest

