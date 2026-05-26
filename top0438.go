package main

func findAnagrams(s string, p string) (res []int) {
	var pc, wc [26]int
	for _, c := range p {
		pc[c-'a']++
	}
	for i := range s {
		wc[s[i]-'a']++
		if i >= len(p) {
			wc[s[i-len(p)]-'a']--
		}
		if i >= len(p)-1 && wc == pc {
			res = append(res, i-len(p)+1)
		}
	}
	return
}
