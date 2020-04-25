set -e

cd /go/src/github.com/tradingAI/runner

cd ..
rm -rf proto
git clone https://github.com/tradingAI/proto.git
echo `pwd`

cd runner

make proto
make mockserv
