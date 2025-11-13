package translations

type Translation struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
    Tag         string `json:"tag" gorm:"uniqueIndex:uniq_tag_lang;size:50;not null"`
    Lang        string `json:"lang" gorm:"uniqueIndex:uniq_tag_lang;size:5;not null"`
	Translation string `json:"translation" gorm:"type:text;not null"`
}