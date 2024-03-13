<?php

function xor_encrypt($in) {

	// $key = '<censored>';
	$key = base64_decode("MGw7JCQ5OC04PT8jOSpqdmkgJ25nbCorKCEkIzlscm5oKC4qLSgubjY=");
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

// This will output KNHL in a repeating manner, which is most likely our key
echo $data_enc;

?>
