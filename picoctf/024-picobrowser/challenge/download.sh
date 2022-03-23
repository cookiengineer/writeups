
server="http://jupiter.challenges.picoctf.org:50522";

curl "$server/"              > index.html;
curl "$server/flag"          > flag.html;
curl "$server/unimplemented" > unimplemented.html;

