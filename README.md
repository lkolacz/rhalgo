# Rolling Hash Algorithm

A rolling hash diffing algorithm. 
Comparing original and an updated version of file that return "delta" (a description) which can be used to upgrade an original version of the file into the new file. 
The description contains the information about each "change" per line (so the chunk is dynamic, dedicated to each line):
1. At index 0 - information if filename have been changed.
2. At n index you can find detected chunk changes, removals, additions
3. Unit tests functions describ and present real live example/the operation.

