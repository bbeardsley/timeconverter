#!/usr/bin/env bash

ls dist/timeconverter*.zip | xargs -n 1 sha256sum | sed 's#dist/##'
