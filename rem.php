#!/usr/bin/env php
<?php
$base = 16;
if (isset($argv[1]) && is_numeric($argv[1])) {
    echo $argv[1] . ' = ' . $argv[1] / $base . "\n";
} else {
    for ($i = 10; $i <= 100; $i++) {
        echo $i . ' = ' . $i / $base . "\n";
    }
}
