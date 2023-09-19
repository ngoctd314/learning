package voting

import "fmt"

func TimeKey(id string) string {
	return fmt.Sprint("time:" + id)
}

func ScoreKey(id string) string {
	return fmt.Sprint("score:", id)
}
