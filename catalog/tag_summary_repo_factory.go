package catalog

// import (
// 	"os"
// 	"strings"
// )

// //CreateTagRepo creates an tag repository based on env variable
// func CreateTagRepo() *TagSummaryRepo {
// 	repoType := os.Getenv("REPOTYPE")
// 	var tRepo *TagRepo

// 	switch strings.ToUpper(repoType) {
// 	case "REDIS":
// 		tRepo = &tagrepo.RedisTagRepo{}
// 	default:
// 		tRepo = &tageepo.InMemoryTagRepo{}
// 	}

// 	return (tRepo)
// }
