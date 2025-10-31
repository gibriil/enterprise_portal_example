<?php
echo "<h1>Keycloak User Claims</h1>";
echo "<pre>";
foreach ($_SERVER as $key => $value) {
    if (strpos($key, 'OIDC_') === 0 || strpos($key, 'REMOTE_') === 0) {
        echo "$key = $value\n";
    }
}
echo "</pre>";

echo "<h2>Full PHP Info</h2>";
phpinfo();
?>