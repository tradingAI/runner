import codecs
import os

from setuptools import find_packages, setup


def read(fname):
    return codecs.open(os.path.join(
        os.path.dirname(__file__), fname)).read().strip()


with open("README.md", "r") as fh:
    long_description = fh.read()


def read_install_requires():
    reqs = []
    return reqs


setup(name='trunner',
      version=read('trunner/VERSION.txt'),
      description='runner于运行job的py package 封装',
      url='https://github.com/tradingAI/runner/tree/master/python',
      author='liuwen',
      author_email='liuwen.w@qq.com',
      long_description=long_description,
      long_description_content_type="text/markdown",
      packages=find_packages(),
      include_package_data=True,
      zip_safe=False,
      classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
        ],
      python_requires='>=3',
      install_requires=read_install_requires(),
      package_data={'': ['*.csv', '*.txt']},
      )
