package attacks

import (
	"auditor/core"
)

func DictionaryAttack(words []string, targetHash string, hashType string) string {
	var jobs []core.Job

	for _, word := range words {
		jobs = append(jobs, core.Job{
			Word:       word,
			TargetHash: targetHash,
			HashType:   hashType,
		})
	}

	return core.RunEngine(jobs, 4)
}