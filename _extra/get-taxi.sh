#!/bin/bash

set -x
set -e

cd /tmp
curl -LO  https://storage.googleapis.com/353solutions/c/data/taxi.tar
tar xf taxi.tar