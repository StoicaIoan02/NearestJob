package main

import "fmt"

func AdRezult(fin, fout string, start, final int, home string) error {
	s := Unload(fin)
	for i, lin := range s {
		if i < start || i >= final {
			continue
		}
		dist := TimeResponse(home, lin[den.Unitate_de_invatamant])
		s[i] = append(s[i], dist.Distance.Text, fmt.Sprintf("%d", dist.Distance.Value), dist.Duration.Text, fmt.Sprintf("%d", dist.Duration.Value), dist.Status)
	}

	Upload(fout, s)
	return nil
}
