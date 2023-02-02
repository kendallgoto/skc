# SKC Lang Server

## Prerequisites
Relies on [antlr.jar](https://www.antlr.org/download.html) to be installed at /antlr.jar

## Usage
Build the application via Docker and connect in your [browser](http://localhost:8080):
```bash
docker-compose up
```

## Running Natively
Go CPython bindings are pinned to Python 3.7.10 - if you have Python3.7.10 headers installed, you can build natively (only works on linux AFAIK):
```bash
make && ./bin/skcserver
```
If you need to downgrade Python 3.7.10 manually:
```bash
sudo add-apt-repository ppa:deadsnakes/ppa
sudo apt-get update
sudo apt-get install python3.7 python3.7-dev
# Then - adjust the pkg-config config to point to the correct version
# you can find the relevant .pc location with pkg-config --variable pc_path pkg-config for your system
cd /usr/lib/x86_64-linux-gnu/pkgconfig
sudo rm -f python3.pc # will be a symlink to whatever python version you have installed by default
sudo ln -s python-3.7.pc python3.pc
pkg-config python3 --cflags # should yield /usr/include/python3.7m
```
