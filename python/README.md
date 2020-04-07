# trunner

runner 上传参数的接口封装

## 使用

`pip install trunner>=1.0.2`

- tbase
    - [x] train: `python3 -m trunner.tbase --alg ddpg --codes 000001.SZ --seed 0`
    - [x] eval: `python3 -m trunner.tbase --eval --alg ddpg --codes 000001.SZ --seed 0`
    - [x] infer: `python3 -m trunner.tbase --infer --alg ddpg --codes 000001.SZ --seed 0`
    - 详细参数: [tbase args](https://github.com/tradingAI/tbase/blob/fbbab069d3594a1d11e3cc1b80ad14adb98a7d86/tbase/common/cmd_util.py#L112)
