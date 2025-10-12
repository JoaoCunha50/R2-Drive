package translations

type BaseTranslation struct {
	Tag string
	Translation string
}

type Translation struct {
	// TO DO
	BaseTranslation
	Language string
}