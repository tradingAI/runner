.PHONY: install update upload

install:
	pip3 install -e . --user

update:
	git pull origin master
	pip3 install -e . --user

upload:
	# 上传到 pypi.org 方便用户使用 pip 安装
	python3 setup.py sdist bdist_wheel
	python3 -m twine upload dist/*
	rm -rf build
	rm -rf dist
