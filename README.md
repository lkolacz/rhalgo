# Rolling Hash Algorithm

A rolling hash diffing algorithm. 
Comparing original and an updated version of file that return "delta" (a description) which can be used to upgrade an original version of the file into the new file. 
The description contains the information about each "change" per line (so the chunk is dynamic, dedicated to each line):
1. At index 0 - information if filename have been changed.
2. At n index you can find detected chunk changes, removals, additions
3. Unit tests functions describ and present real live example/the operation.
You can generate signature for a given delta via set env var `EQLDefaultHexPrivateKey` or use default one.

# Run unit tests

```
leszek.kolacz@Mac-Leszek ~ % cd path/to/project/rhalgo/src
leszek.kolacz@Mac-Leszek src % go test . -v
2023/01/18 13:45:52 You are using default privet key! Please set env var EQLDefaultHexPrivateKey!
=== RUN   TestRDiffHash
=== RUN   TestRDiffHash/rdiff_-_text_changed_in_liness
=== RUN   TestRDiffHash/rdiff_-_line_removed_from_file
=== RUN   TestRDiffHash/rdiff_-_line_added_to_file
=== RUN   TestRDiffHash/rdiff_-_mix_of_changes
=== RUN   TestRDiffHash/rdiff_-_mix_2_of_changes_(more_removes_at_the_end)
--- PASS: TestRDiffHash (0.00s)
    --- PASS: TestRDiffHash/rdiff_-_text_changed_in_liness (0.00s)
    --- PASS: TestRDiffHash/rdiff_-_line_removed_from_file (0.00s)
    --- PASS: TestRDiffHash/rdiff_-_line_added_to_file (0.00s)
    --- PASS: TestRDiffHash/rdiff_-_mix_of_changes (0.00s)
    --- PASS: TestRDiffHash/rdiff_-_mix_2_of_changes_(more_removes_at_the_end) (0.00s)
=== RUN   TestSignatureHash
=== RUN   TestSignatureHash/Signature_-_make_a_signature_on_delta_and_verify_it
--- PASS: TestSignatureHash (0.00s)
    --- PASS: TestSignatureHash/Signature_-_make_a_signature_on_delta_and_verify_it (0.00s)
PASS
ok      rdiff   0.325s
```
