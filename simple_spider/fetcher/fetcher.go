package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"math/rand"
)

var userAgent = []string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:6.0) Gecko/20100101 Firefox/6.0",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Opera/9.80 (Windows NT 6.1; U; zh-cn) Presto/2.9.168 Version/11.50",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 2.0.50727; SLCC2; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.3; .NET4.0C; Tablet PC 2.0; .NET4.0E)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; InfoPath.3)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; GTB7.0)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1)",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; ) AppleWebKit/534.12 (KHTML, like Gecko) Maxthon/3.0 Safari/534.12",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.3; .NET4.0C; .NET4.0E)",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.3; .NET4.0C; .NET4.0E; SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.3 (KHTML, like Gecko) Chrome/6.0.472.33 Safari/534.3 SE 2.X MetaSr 1.0",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.3; .NET4.0C; .NET4.0E)",

}

func Determine(reader *bufio.Reader) encoding.Encoding {
	bytes, err := reader.Peek(1024)
	if err != nil {
		fmt.Println("data peek error: ", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	//fmt.Println(e, name, certain)
	return e
}

//var ratelimiter = time.Tick(100 * time.Millisecond)

func Fetcher(url string) ([]byte, error) {
	//<- ratelimiter
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent",
		userAgent[rand.Intn(len(userAgent))])
	//resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	httpclient := http.DefaultClient

	resp, err := httpclient.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error occured，status code %s", resp.StatusCode)
	}

	defer resp.Body.Close()

	bodyReader := bufio.NewReader(resp.Body)
	encod := Determine(bodyReader)
	//reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	reader := transform.NewReader(bodyReader, encod.NewDecoder())
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(content))
	return content, nil
}


