<?php

function xor_encrypt($in) {

	// from output of 01_find_key.php
	$key = "KNHL";
	$text = $in;
	$outText = '';

	// Iterate through each character
	for($i=0;$i<strlen($text);$i++) {
		$outText .= $text[$i] ^ $key[$i % strlen($key)];
	}

	return $outText;

}

$data = array( "showpassword"=>"yes", "bgcolor"=>"#ffffff" );
$data_json = json_encode($data);
$data_enc = xor_encrypt($data_json);
$data_base64 = base64_encode($data_enc);

echo $data_base64;

?>
