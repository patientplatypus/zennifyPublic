#!/bin/bash

#from https://www.hellorust.com/demos/add/index.html

rm -rf ./add.big.wasm
rm -rf ./add.wasm
rm -rf ../add.wasm

rustup update
rustup target add wasm32-unknown-unknown --toolchain nightly
rustc +nightly --target wasm32-unknown-unknown -O --crate-type=cdylib add.rs -o add.big.wasm

cargo install --git https://github.com/alexcrichton/wasm-gc 
wasm-gc add.big.wasm add.wasm  

while [ ! add.wasm ]
do
  sleep 2
done

mv add.wasm ..