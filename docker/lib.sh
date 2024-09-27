#!/bin/sh

git clone https://github.com/confluentinc/librdkafka.git
cd librdkafka
./configure
make
make install
CC="musl-gcc -static" ./configure --exec-prefix=/musl && make