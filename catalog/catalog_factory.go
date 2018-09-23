package catalog

//CreateCatalog creates a catalog and injects dependencies
func CreateCatalog() *ArticleCatalog {
	return &ArticleCatalog{
		repo: CreateRepo{},
	}
}
