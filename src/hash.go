package rdiff

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"path/filepath"
)

const (
	HashSize = 32
	Added    = "added"
	Changed  = "changed"
	Deleted  = "deleted"
	Rename   = "rename"
)

type Hash [HashSize]byte
type Alpha map[int]string
type Dict map[string]string
type Delta map[int]Dict

func Sum256Hash(b []byte) Hash {
	return sha256.Sum256(b)
}

func (hash Hash) String() string {
	return hex.EncodeToString(hash[:])
}

func contains(s []string, str string) bool {
	for _, val := range s {
		if val == str {
			return true
		}
	}

	return false
}

func Compute(filePathName string) Alpha {
	file, err := os.Open(filePathName)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	idx := 0
	fileName := filepath.Base(filePathName)
	fileHash := Sum256Hash([]byte(fileName))
	alpha := Alpha{idx: fileHash.String()}
	for scanner.Scan() {
		idx++
		text := scanner.Text()
		hash := Sum256Hash([]byte(text))
		alpha[idx] = hash.String()
	}

	return alpha
}

func Compare(filePathName string, alpha Alpha) Delta {
	file, err := os.Open(filePathName)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	idx := 0
	delta := Delta{}
	fileName := filepath.Base(filePathName)
	fileHash := Sum256Hash([]byte(fileName))
	if alpha[idx] != fileHash.String() {
		delta[idx] = Dict{
			"hashWas": alpha[idx],
			"hashIs":  fileHash.String(),
			"action":  Rename,
		}
	}
	newAlpha := Alpha{idx: fileHash.String()}
	newAlphaList := []string{}
	for scanner.Scan() {
		idx++
		text := scanner.Text()
		hash := Sum256Hash([]byte(text))
		newAlpha[idx] = hash.String()
		newAlphaList = append(newAlphaList, hash.String())
	}
	// decreasing because filename is in 0 index; skip it;
	alphaSize := len(alpha) - 1
	newAlphaSize := len(newAlpha) - 1
	idxAlpha := 0
	idxNewAlpha := 0
mainLoop:
	for {
		if idxAlpha < alphaSize {
			idxAlpha++
		}
		if idxNewAlpha < newAlphaSize {
			idxNewAlpha++
		}
		if idxNewAlpha >= newAlphaSize && idxAlpha >= alphaSize {
			break
		}
		if idxAlpha == alphaSize && idxNewAlpha > idxAlpha {
			for idxNewAlpha <= newAlphaSize {
				delta[idxNewAlpha] = Dict{
					"hashWas": "",
					"hashIs":  newAlpha[idxNewAlpha],
					"action":  Added,
				}
				idxNewAlpha++
			}
			continue mainLoop
		} else if idxNewAlpha == newAlphaSize && idxAlpha <= alphaSize {
			for idxAlpha <= alphaSize {
				if !contains(newAlphaList, alpha[idxAlpha]) {
					delta[idxAlpha] = Dict{
						"hashWas": alpha[idxAlpha],
						"action":  Deleted,
					}
				}
				idxAlpha++
			}
			continue mainLoop
		} else if alpha[idxAlpha] != newAlpha[idxNewAlpha] {
			// check if lines was removed
			for idxj := 1; idxj <= alphaSize-idxAlpha; idxj++ {
				if alpha[idxAlpha+idxj] == newAlpha[idxNewAlpha] {
					delta[idxAlpha] = Dict{
						"hashWas": alpha[idxAlpha],
						"action":  Deleted,
					}
					idxAlpha++
					continue mainLoop
				}
			}
			// check if any line was added
			for idxi := 1; idxi < newAlphaSize-idxNewAlpha; idxi++ {
				if alpha[idxAlpha] == newAlpha[idxNewAlpha+idxi] {
					value := Dict{
						"hashWas": alpha[idxAlpha],
						"hashIs":  newAlpha[idxNewAlpha],
						"action":  Added,
					}
					if alpha[idxAlpha] == "" {
						delta[idxNewAlpha] = value
					} else {
						delta[idxAlpha] = value
					}
					idxNewAlpha++
					continue mainLoop
				} else {
					break
				}
			}

			// if none of them then this line is a simple change
			delta[idxAlpha] = Dict{
				"hashWas": alpha[idxAlpha],
				"hashIs":  newAlpha[idxNewAlpha],
				"action":  Changed,
			}
		}

	}

	return delta
}
