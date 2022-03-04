#!/usr/bin/env bash

$* &> log.tmp
less +gg log.tmp
cat log.tmp
rm log.tmp
