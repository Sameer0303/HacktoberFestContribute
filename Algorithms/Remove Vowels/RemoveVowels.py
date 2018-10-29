#!/usr/bin/env python3

import sys 
import re

removedVowels = re.sub('a|e|i|o|u', '', sys.argv[1])
print(removedVowels)
