#!/usr/bin/env bash

set -e

version="0.3.1"
package="timeconverter"
bin="timeconverter"
repo="https://github.com/bbeardsley/timeconverter"

platforms=(
  "windows/amd64"
  "windows/386"
  "darwin/amd64"
  "darwin/386"
  "linux/amd64"
  "linux/386"
  "linux/arm"
)

rm -f dist/*

for platform in "${platforms[@]}"
do
  platform_split=(${platform//\// })
  GOOS=${platform_split[0]}
  GOARCH=${platform_split[1]}

  BIN="${bin}"
  if [ $GOOS = "windows" ]; then
    BIN="${bin}.exe"
  fi
  ZIP=$package'_v'$version'_'$GOOS'_'$GOARCH'.zip'

  echo "Building for GOOS=$GOOS GOARCH=$GOARCH"

  env GOOS=$GOOS GOARCH=$GOARCH go build -o dist/${BIN} &&
  zip -q dist/${ZIP} -j dist/${BIN} &&
  rm dist/${BIN}
done

DARWIN_AMD64=${package}_v${version}_darwin_amd64.zip
DARWIN_386=${package}_v${version}_darwin_386.zip

cat << EOF > dist/${package}.rb
# This file was generated by release.sh
require 'formula'
class TimeConverter < Formula
  homepage '${repo}'
  version '${version}'

  if Hardware::CPU.is_64_bit?
     url '${repo}/releases/download/v${version}/${DARWIN_AMD64}'
     sha '$( sha256sum dist/${DARWIN_AMD64} | awk '{ print $1 }' | xargs printf )'
  else
     url '${repo}/releases/download/v${version}/${DARWIN_386}'
     sha '$( sha256sum dist/${DARWIN_386} | awk '{ print $1 }' | xargs printf )'
  end

  def install
     bin.install '${BIN}'
  end
end
EOF
