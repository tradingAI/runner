set -e

cd /go/src/github.com/tradingAI
rm -rf proto
git clone https://github.com/tradingAI/proto.git
echo `pwd`

cd /go/src/github.com/tradingAI/runner

make proto
make clean
make run
