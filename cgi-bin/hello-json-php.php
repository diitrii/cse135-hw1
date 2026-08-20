<?php
header("Cache-Control: no-cache");
header("Content-type: application/json");

$date = time();
$address = $_SERVER['REMOTE_ADDR'];

$message = [
    'title' => 'Hello, PHP!',
    'heading' => 'Hello, PHP!',
    'message' => 'This page was generated with the Perl programming language',
    'time' => $date,
    'IP' => $address
];

echo json_encode($message);
?>