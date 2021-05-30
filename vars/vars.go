package vars

var (
	URL string
	FILE string
	NAME string
)

var Header = map[string]string{
	"User-Agent" : "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
	"Accept" : "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"Accept-Encoding" : "gzip, deflate",
	"Content-Type" : "application/x-www-form-urlencoded",
	"Authorization": "Basic YWRtaW46YWRtaW4=", //admin:admin
}

var (
	ThreadNum = 200
	//ThreadNum为goroutine个数
)