package catalog

import (
	"ArticleApp/repo"
	"os"
	"strings"
)

//CreateRepo creates an article repository based on env variable
func CreateRepo() Repository {
	repoType := os.Getenv("REPOTYPE")
	var aRepo Repository

	switch strings.ToUpper(repoType) {
	case "REDIS":
		aRepo = &repo.RedisArticleRepo{}
	default:
		aRepo = &repo.InMemoryArticleRepo{}
	}

	return (aRepo)
}
