<?php

$removedVowels = preg_replace('/a|e|i|o|u/', '', $argv[1]);

print($removedVowels);
