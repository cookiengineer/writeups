
server="http://jupiter.challenges.picoctf.org:6353";

curl "$server/"                 > index.html;
curl "$server/barbed_wire.jpeg" > barbed_wire.jpeg;
curl "$server/md5.js"           > md5.js;

