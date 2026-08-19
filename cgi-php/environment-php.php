<?php
header("Cache-Control: no-cache");
header("Content-Type: text/html");
?>
<!DOCTYPE html>
<html>
    <head><title>Environment Variables</title></head>
    <body><h1 align="center">Environment Variables</h1>
    <hr>
<?php
    $keys = array_keys($_SERVER);
    sort($keys);
    foreach ($keys as $variable) {
        echo "<b>" . htmlspecialchars($variable) . ":</b>" . htmlspecialchars($_SERVER[$variable]) . "<br>";
    }
?>
</body>
</html> 