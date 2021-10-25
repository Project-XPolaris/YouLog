rm -rf ./pack-output
mkdir ./pack-output
go build youlog.go
cp ./youlog ./pack-output
cp ./pack/* ./pack-output