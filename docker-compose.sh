set -e
cd "$(dirname "$0")"
echo `pwd`

cd ..
rm -rf proto
git clone https://github.com/tradingAI/proto.git
echo `pwd`

cd runner

bash proto.sh

go test ./...
