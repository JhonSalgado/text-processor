package stopwords

var Stopwords map[string][]string = make(map[string][]string)

func init() {
	Stopwords["ar"] = ar
	Stopwords["bg"] = bg
	Stopwords["ca"] = ca
	Stopwords["cs"] = cs
	Stopwords["da"] = da
	Stopwords["de"] = de
	Stopwords["en"] = en
	Stopwords["es"] = es
	Stopwords["fi"] = fi
	Stopwords["fr"] = fr
	Stopwords["gu"] = gu
	Stopwords["he"] = he
	Stopwords["hi"] = hi
	Stopwords["hu"] = hu
	Stopwords["id"] = id
	Stopwords["it"] = it
	Stopwords["ms"] = ms
	Stopwords["nb"] = nb
	Stopwords["nl"] = nl
	Stopwords["pl"] = pl
	Stopwords["pt"] = pt
	Stopwords["ro"] = ro
	Stopwords["ru"] = ru
	Stopwords["sk"] = sk
	Stopwords["sv"] = sv
	Stopwords["tr"] = tr
	Stopwords["uk"] = uk
	Stopwords["vi"] = vi
}
