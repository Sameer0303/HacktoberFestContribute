#!/usr/bin/env node

let removedVowels = process.argv[2].replace(/a|e|i|o|u/g, '');
console.log(removedVowels);
