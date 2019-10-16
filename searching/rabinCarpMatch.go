package searching

import "fmt"

/*
Given a text txt[0..n-1] and a pattern pat[0..m-1], write a function search(char pat[], char txt[])
that prints all
occurrences of pat[] in txt[]. You may assume that n > m.
 */

/*
Input:  txt[] = "THIS IS A TEST TEXT"
        pat[] = "TEST"
Output: Pattern found at index 10

Input:  txt[] =  "AABAACAADAABAABA"
        pat[] =  "AABA"
Output: Pattern found at index 0
        Pattern found at index 9
        Pattern found at index 12
 */
/*
The Naive String Matching algorithm slides the pattern one by one.
After each slide, it one by one checks
characters at the current shift and if all characters match then prints the match.
O(n*m)
 */

func SimpleSearch(txt string, pattern string) {
	for i:=0;i<len(txt);i +=1 {
		// find substring from i to len(pattern)
		if i + len(pattern) > len(txt) {
			break
		}
		s := txt[i:len(pattern) + i]
		if s == pattern {
			fmt.Println(i)
		}
	}
}
/*
But unlike the Naive algorithm, Rabin Karp algorithm matches the hash value of the pattern with
the hash value of current substring of text, and if the hash values match then only it starts
matching individual characters. So Rabin Karp algorithm needs to calculate hash values for
following strings.

1) Pattern itself.
2) All the substrings of text of length m.
 */

/*
 Instead of checking two strings we hash the pattern and all the
substrings of text of length m and we compare their hashes. If they match we found a match.

We need an efficient string hashing ADT for it.

First compute hash function of pattern.
Next compute the hash of the first len(pattern) chars in txt.
Next we have for loop that check all the other chars. We throw away the first char from the current string
and append the next char from the txt in the current string. we check if the hash is the same as the pattern.
We also need to check char by char if they are equal.

O(n+m)
 */

const q = 15461
const d = 256

func RabinKarpSearch(txt string, pattern string) {
	m := len(pattern)
	n := len(txt)
	i, j := 0, 0
	patHash := 0 // hash value for pattern
	txtHash := 0 // hash value for txt
	h := 1

	// The value of h would be "pow(d, M-1)%q"
	// d is the number of characters in the input alphabet = 256
	for i = 0; i < m-1; i += 1 {
		h = (h * d) % q
	}

	// Calculate the hash value of pattern and first
	// window of text
	for i = 0; i < m; i += 1 {
		patHash = (d*patHash + int(pattern[i])) % q
		txtHash = (d*txtHash + int(txt[i])) % q
	}

	// Slide the pattern over text one by one
	for i = 0; i <= n-m; i += 1 {
		if patHash == txtHash {
			/* Check for characters one by one */
			for j = 0; j < m; j  +=1 {
				if txt[i+j] != pattern[j] {
					break
				}
			}
			// if p == t and pat[0...M-1] = txt[i, i+1, ...i+M-1]
			if j == m {
				fmt.Println("Pattern found at index ", i)
			}
		}
		// Calculate hash value for next window of text: Remove
		// leading digit, add trailing digit
		if i < n-m {
			txtHash = (d*(txtHash - int(txt[i])*h) + int(txt[i+m]))%q
			// We might get negative value of t, converting it
			// to positive
			if txtHash < 0 {
				txtHash = txtHash + q
			}
		}
	}
}
