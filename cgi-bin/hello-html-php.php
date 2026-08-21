<?php
header("Cache-Control: no-cache");
header("Content-Type: text/html");

$time = mktime();
$ip = getenv("REMOTE_ADDR");
?>

<!DOCTYPE html>
<html>
<head>
<title>Hello CGI World</title>
</head>

<body>
    <h1 align=center>Hello HTML World</h1><hr/>
    <p>Hello World</p><br>
    <p>This page was generated with the PHP programming language</p><br>
    <p>This program was generated at: <?php echo $time; ?></p><br>
    <p>Your current IP Address is: <?php echo $ip; ?></p><br>
</body>
</html>

