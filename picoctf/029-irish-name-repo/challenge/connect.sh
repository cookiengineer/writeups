
server="http://jupiter.challenges.picoctf.org:50009";

curl "$server/"             > index.html;
curl "$server/login.html"   > login.html;
curl "$server/support.html" > support.html;

mkdir "W3.CSS Template_files";
curl "$server/W3.CSS%20Template_files/w3.css"  > "W3.CSS Template_files/w3.css";
curl "$server/W3.CSS%20Template_files/css.css" > "W3.CSS Template_files/css.css";

mkdir "irish";
curl "$server/irish/Aidan_Gillen.jpg"          > ./irish/Aidan_Gillen.jpg;
curl "$server/irish/Aidan_Higgins_pic2.jpg"    > ./irish/Aidan_Higgins_pic2.jpg;
curl "$server/irish/Alison_Doody.jpg"          > ./irish/Alison_Doody.jpg;
curl "$server/irish/Dylan_Moran_Melbourne.jpg" > ./irish/Dylan_Moran_Melbourne.jpg;
curl "$server/irish/TommyT.jpeg"               > ./irish/TommyT.jpeg;
curl "$server/irish/Brendan_Gleeson.jpg"       > ./irish/Brendan_Gleeson.jpg;

