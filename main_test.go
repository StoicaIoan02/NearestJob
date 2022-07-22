package main

// func Test2AdRezult(t *testing.T) {
// 	SchoolsInput := "result/ilfov_schools_home_time_sorted.txt"
// 	SchoolsOuput := "result/ilfov_schools_home_leu_time_sorted.txt"
// 	start := 0
// 	finish := 123
// 	// Upload(SchoolsInput, Table(URL2, start, finish))
// 	err := AdRezult(SchoolsInput, SchoolsOuput, start, finish, HOME) // Adauga 5 elemente la finalul fiecarei linii
// 	if err != nil {
// 		t.Errorf("got error: %v", err)
// 	}
// 	s := Unload(SchoolsOuput)
// 	got := s[0][den.Status]
// 	want := "OK"
// 	if got != want {
// 		t.Errorf("got %s, want %s", got, want)
// 	}
// }

// func TestSortFile(t *testing.T) {
// 	fileToSort := "ilfov_school_dist.txt"
// 	err := SortFile(fileToSort)
// 	if err != nil {
// 		t.Errorf("testSortFile error%v", err)
// 	}
// 	s := Unload(fileToSort)
// 	got := s[0][den.DurationValue]
// 	want := "5639"
// 	if got != want {
// 		t.Errorf("testSortFile: got %s, want %s", got, want)
// 	}
// }

// func TestAdRezult(t *testing.T) {
// 	SchoolsInput := "ilfov_aux_2.txt"
// 	SchoolsOuput := "ilfov_school_dist.txt"
// 	start := 0
// 	finish := 2
// 	Upload(SchoolsInput, Table(URL2, start, finish))
// 	err := AdRezult(SchoolsInput, SchoolsOuput, start, finish, HOME) // Adauga 5 elemente la finalul fiecarei linii
// 	if err != nil {
// 		t.Errorf("got error: %v", err)
// 	}
// 	s := Unload(SchoolsOuput)
// 	got := s[0][den.Status]
// 	want := "OK"
// 	if got != want {
// 		t.Errorf("got %s, want %s", got, want)
// 	}
// }

// func TestTime(t *testing.T) {
// 	got := TimeResponse("Str. Sfânta Agnes 21, Popești-Leordeni 077160, România", "Șoseaua Banatului 75, Chitila 077045, România")
// 	want := Rezult{}
// 	want.Distance.Text = "25,5, km"
// 	want.Distance.Value = 25505
// 	want.Duration.Text = "1 hour 32 mins"
// 	want.Duration.Value = 5497
// 	want.Status = "OK"
// 	//{Rezult.Distance{"25,5, km", 25505}, {"1 hour 32 mins", 5497}, "OK"}
// 	if got.Distance.Value-want.Distance.Value > 600 || got.Duration.Value-want.Duration.Value > 600 || got.Status != want.Status {
// 		t.Errorf("Time: got %v, want %v", got, want)
// 	}
// }

// func TestTable(t *testing.T) {
// 	got := Table(URL1, 0, 2)
// 	want := [][]string{{"1592", "BUFTEA (ORAŞ BUFTEA)",
// 		"ȘCOALA GIMNAZIALĂ NR. 2 BUFTEA", "GIMNAZIAL",
// 		"INFORMATICA SI TEHNOLOGIA INFORMATIEI SI A COMUNICATIILOR",
// 		"18 / 0", "", "4 ani"}, {
// 		"1574", "PANTELIMON (ORAŞ PANTELIMON)",
// 		"ŞCOALA GIMNAZIALA NR. 1 PANTELIMON", "GIMNAZIAL",
// 		"INFORMATICA SI TEHNOLOGIA INFORMATIEI SI A COMUNICATIILOR",
// 		"18 / 0", "", "4 ani",
// 	}}
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %v, want %v", got, want)
// 	}
// }

// func TestUpload(t *testing.T) {
// 	got := Upload("ilfov_aux.txt", Table(URL1, 0, 2))
// 	want := 1
// 	if got != want {
// 		t.Errorf("got %v, want %v", got, want)
// 	}

// 	t.Run("verificare scriere/citire fisier: ", func(t *testing.T) {
// 		start := 0
// 		nr_linii := 40
// 		URL := URL5
// 		file := "result/aux.txt"
// 		Upload(file, Table(URL, start, nr_linii))
// 		got := Unload(file)
// 		want := Table(URL, start, nr_linii)
// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})

// }
