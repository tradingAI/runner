set -e
cd "$(dirname "$0")"

make proto
go test -v ./...
