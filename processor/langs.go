package processor

// path of the stopwords files folder
var basePath string = "../stopwords/"

// lang_files contains the location of the stopwords file for each language
var langFiles = map[string]string{
	"ar": "arabic.txt",
	"bg": "bulgarian.txt",
	"ca": "catalan.txt",
	"cs": "czech.txt",
	"da": "danish.txt",
	"nl": "dutch.txt",
	"en": "english.txt",
	"fi": "finnish.txt",
	"fr": "french.txt",
	"de": "german.txt",
	"gu": "gujarati.txt",
	"he": "hebrew.txt",
	"hi": "hindi.txt",
	"hu": "hungarian.txt",
	"id": "indonesian.txt",
	"ms": "malaysian.txt",
	"it": "italian.txt",
	"nb": "norwegian.txt",
	"pl": "polish.txt",
	"pt": "portuguese.txt",
	"ro": "romanian.txt",
	"ru": "russian.txt",
	"sk": "slovak.txt",
	"es": "spanish.txt",
	"sv": "swedish.txt",
	"tr": "turkish.txt",
	"uk": "ukrainian.txt",
	"vi": "vietnamese.txt",
}
