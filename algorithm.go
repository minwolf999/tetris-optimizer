package main

/*********************************************************
* This algorithm manage tetrise to optimise possibility
* Is working recursively.
* He check tetrominos && test all possibilities
* Optimize breaking when all tetrominos fit in the NewTab.
* After recursilving, this function DeleteLastForm
* For managing possibilities.
*
* After breaking in loop, Optimise, overwrite the tab
* If new form had writing
*********************************************************/

func Optimize(listTetris [][]string, newtab []string) {
	form := listTetris[0]
	write := false

	for i := range newtab {
		if len(newtab[i:]) < len(form) {
			DeleteLastForm(newtab, form)
			break
		}

		for y := range newtab[i] {
			if newtab[i][y] != '.' && form[0][0] != '.' {
				continue
			}

			for x := range form {
				if len(newtab[i+x]) < y+len(form[x]) {
					DeleteLastForm(newtab, form)
					break
				}

				if x != 0 && CountNbrPoint(newtab, []rune(GetLetter(form))[0]) == 0 {
					break
				}

				for j := range form[x] {
					if form[x][j] == '.' {
						continue
					}

					if form[x][j] != '.' && newtab[i+x][y+j] != '.' {
						DeleteLastForm(newtab, form)
						break
					}

					temp := newtab[i+x][:y+j] + string(form[x][j]) + newtab[i+x][y+j+1:]
					newtab[i+x] = temp

					write = true
				}
			}

			if write {
				write = false
				r := []rune(GetLetter(form))[0]
				if len(listTetris) == 1 && TabContain(newtab, GetLetter(form), CountNbrPoint(form, r)) {

					nbrPoint := CountNbrPoint(newtab, '.')

					if nbrPoint < res.NbrOfPoints {
						res.NbrOfPoints = nbrPoint
						res.result = nil
						res.result = append(res.result, newtab...)

						found = true
					}

				} else if TabContain(newtab, GetLetter(form), CountNbrPoint(form, r)) {
					Optimize(listTetris[1:], newtab)

					if found {
						return
					}
				}

				DeleteLastForm(newtab, form)
			}
		}
	}
}
