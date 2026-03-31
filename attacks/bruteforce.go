package attacks

import "auditor/core"

func BruteForceAttack(charset string, length int, targetHash string, hashType string) string {
	var jobs []core.Job

	var generate func(string, int)
	generate = func(prefix string, remaining int) {
		if remaining == 0 {
			jobs = append(jobs, core.Job{
				Word:       prefix,
				TargetHash: targetHash,
				HashType:   hashType,
			})
			return
		}

		for _, c := range charset {
			generate(prefix+string(c), remaining-1)
		}
	}

	generate("", length)

	return core.RunEngine(jobs, 4)
}