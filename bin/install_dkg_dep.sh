#!/bin/bash

if [ -e .dep/dkg ]; then
  exit 0
fi

if [ ! -d .dep/dkg ]; then
  mkdir -p .dep/dkg
  cd .dep/dkg
  git clone --depth 1 -b master git://github.com/tangerine-network/bls.git &
  git clone --depth 1 -b master git://github.com/tangerine-network/mcl.git &
  wait
  cd bls
  make test_go -j MCL_USE_OPENSSL=0
  cd ../../../
fi
cd vendor/github.com/tangerine-network && rm -rf bls mcl
ln -s ../../../.dep/dkg/* .
