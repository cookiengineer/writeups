
server="http://mercury.picoctf.net:5080";


curl "$server/" > index.html;
curl "$server/mycss.css" > mycss.css;
curl "$server/myjs.js" > myjs.js;
curl "$server/robots.txt" > robots.txt;
curl "$server/.htaccess" > .htaccess;
curl "$server/.DS_Store" > .DS_Store;

