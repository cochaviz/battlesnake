#!/usr/bin/env bash

$* &> log.tmp
less +gg/Turn log.tmp
cat log.tmp
rm log.tmp
