package stopwords

var ms = []string{
	"ada",
	"adakah",
	"adakan",
	"adalah",
	"adanya",
	"adapun",
	"agak",
	"agar",
	"akan",
	"aku",
	"akulah",
	"akupun",
	"al",
	"alangkah",
	"allah",
	"amat",
	"antara",
	"antaramu",
	"antaranya",
	"apa",
	"apa-apa",
	"apabila",
	"apakah",
	"apapun",
	"atas",
	"atasmu",
	"atasnya",
	"atau",
	"ataukah",
	"ataupun",
	"bagaimana",
	"bagaimanakah",
	"bagi",
	"bagimu",
	"baginya",
	"bahawa",
	"bahawasanya",
	"bahkan",
	"bahwa",
	"banyak",
	"banyaknya",
	"barangsiapa",
	"bawah",
	"beberapa",
	"begitu",
	"begitupun",
	"belaka",
	"belum",
	"belumkah",
	"berada",
	"berapa",
	"berikan",
	"beriman",
	"berkenaan",
	"berupa",
	"beserta",
	"biarpun",
	"bila",
	"bilakah",
	"bilamana",
	"bisa",
	"boleh",
	"bukan",
	"bukankah",
	"bukanlah",
	"dahulu",
	"dalam",
	"dalamnya",
	"dan",
	"dapat",
	"dapati",
	"dapatkah",
	"dapatlah",
	"dari",
	"daripada",
	"daripadaku",
	"daripadamu",
	"daripadanya",
	"demi",
	"demikian",
	"demikianlah",
	"dengan",
	"dengannya",
	"di",
	"dia",
	"dialah",
	"didapat",
	"didapati",
	"dimanakah",
	"engkau",
	"engkaukah",
	"engkaulah",
	"engkaupun",
	"hai",
	"hampir",
	"hampir-hampir",
	"hanya",
	"hanyalah",
	"hendak",
	"hendaklah",
	"hingga",
	"ia",
	"iaitu",
	"ialah",
	"ianya",
	"inginkah",
	"ini",
	"inikah",
	"inilah",
	"itu",
	"itukah",
	"itulah",
	"jadi",
	"jangan",
	"janganlah",
	"jika",
	"jikalau",
	"jua",
	"juapun",
	"juga",
	"kalau",
	"kami",
	"kamikah",
	"kamipun",
	"kamu",
	"kamukah",
	"kamupun",
	"katakan",
	"ke",
	"kecuali",
	"kelak",
	"kembali",
	"kemudian",
	"kepada",
	"kepadaku",
	"kepadakulah",
	"kepadamu",
	"kepadanya",
	"kepadanyalah",
	"kerana",
	"kerananya",
	"kesan",
	"ketika",
	"kini",
	"kita",
	"ku",
	"kurang",
	"lagi",
	"lain",
	"lalu",
	"lamanya",
	"langsung",
	"lebih",
	"maha",
	"mahu",
	"mahukah",
	"mahupun",
	"maka",
	"malah",
	"mana",
	"manakah",
	"manapun",
	"masih",
	"masing",
	"masing-masing",
	"melainkan",
	"memang",
	"mempunyai",
	"mendapat",
	"mendapati",
	"mendapatkan",
	"mengadakan",
	"mengapa",
	"mengapakah",
	"mengenai",
	"menjadi",
	"menyebabkan",
	"menyebabkannya",
	"mereka",
	"merekalah",
	"merekapun",
	"meskipun",
	"mu",
	"nescaya",
	"niscaya",
	"nya",
	"olah",
	"oleh",
	"orang",
	"pada",
	"padahal",
	"padamu",
	"padanya",
	"paling",
	"para",
	"pasti",
	"patut",
	"patutkah",
	"per",
	"pergilah",
	"perkara",
	"perkaranya",
	"perlu",
	"pernah",
	"pertama",
	"pula",
	"pun",
	"sahaja",
	"saja",
	"saling",
	"sama",
	"sama-sama",
	"samakah",
	"sambil",
	"sampai",
	"sana",
	"sangat",
	"sangatlah",
	"saya",
	"se",
	"seandainya",
	"sebab",
	"sebagai",
	"sebagaimana",
	"sebanyak",
	"sebelum",
	"sebelummu",
	"sebelumnya",
	"sebenarnya",
	"secara",
	"sedang",
	"sedangkan",
	"sedikit",
	"sedikitpun",
	"segala",
	"sehingga",
	"sejak",
	"sekalian",
	"sekalipun",
	"sekarang",
	"sekitar",
	"selain",
	"selalu",
	"selama",
	"selama-lamanya",
	"seluruh",
	"seluruhnya",
	"sementara",
	"semua",
	"semuanya",
	"semula",
	"senantiasa",
	"sendiri",
	"sentiasa",
	"seolah",
	"seolah-olah",
	"seorangpun",
	"separuh",
	"sepatutnya",
	"seperti",
	"seraya",
	"sering",
	"serta",
	"seseorang",
	"sesiapa",
	"sesuatu",
	"sesudah",
	"sesudahnya",
	"sesungguhnya",
	"sesungguhnyakah",
	"setelah",
	"setiap",
	"siapa",
	"siapakah",
	"sini",
	"situ",
	"situlah",
	"suatu",
	"sudah",
	"sudahkah",
	"sungguh",
	"sungguhpun",
	"supaya",
	"tadinya",
	"tahukah",
	"tak",
	"tanpa",
	"tanya",
	"tanyakanlah",
	"tapi",
	"telah",
	"tentang",
	"tentu",
	"terdapat",
	"terhadap",
	"terhadapmu",
	"termasuk",
	"terpaksa",
	"tertentu",
	"tetapi",
	"tiada",
	"tiadakah",
	"tiadalah",
	"tiap",
	"tiap-tiap",
	"tidak",
	"tidakkah",
	"tidaklah",
	"turut",
	"untuk",
	"untukmu",
	"wahai",
	"walau",
	"walaupun",
	"ya",
	"yaini",
	"yaitu",
	"yakni",
	"yang",
}