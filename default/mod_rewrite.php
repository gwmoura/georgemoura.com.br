<?php

$path = parse_url($_SERVER['REQUEST_URI'], PHP_URL_PATH);

if ($path == '/help') {
  echo 'This is some help text.';
  exit;
}

echo 'Welcome to the site!';