set -e
cd "$(dirname "$0")"
echo `pwd`

make proto

go test -v ./...
