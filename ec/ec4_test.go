package domains

import (
	"testing"
)

// given a slice of strings, return a comma separated string containing all of
// the unique domains assocated with either HTTP or HTTPS URLS.
//
// the domains in the result list should be squashed to lower case
//
// if the domain starts with www., w3., or web., that part of it should be
// stripped from the domain.
//
// For example, if the slice contains:
//
//  hi, i typically like to search <a href='https://www.google.com/'>http://google.com/!</a>
//  for information but sometimes i enjoy using http://www.duckduckgo.com/
//
//  sometimes it returns results from http://golang.org and https://reddit.com/
//
// the resulting slice should contain:
//
//  duckduckgo.com,google.com,golang.org,reddit.com

func TestExtractDomains3(t *testing.T) {
	tests := []struct {
		name string
		sl   []string
		ep   string
	}{
		{
			name: "Test1",
			sl:   []string{"http://www.aliababa.com"},
			ep:   "aliababa.com",
		},
		{
			name: "Test2",
			sl:   []string{"http://www.starterstudio.org"},
			ep:   "starterstudio.org",
		},
		{
			name: "Test3",
			sl: []string{
				"hi, i typically like to search <a href='https://www.google.com/'>http://google.com/!</a>",
				"for information but sometimes i enjoy using http://www.duckduckgo.com/",
				"",
				"sometimes it returns results from http://golang.org and https://reddit.com/",
			},
			ep: "google.com,duckduckgo.com,golang.org,reddit.com",
		},
		{
			name: "Test4",
			sl:   []string{"http://w3.roxhttp.com http://www.starterstudio.org https://www.http_rocks.com  http://test.com"},
			ep:   "roxhttp.com,starterstudio.org,http_rocks.com,test.com",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := ExtractDomains(test.sl)
			if res != test.ep {
				t.Fatalf("%s expected (%v) returned : (%q)", test.name, test.ep, res)
			}
		})
	}
}

////func TestExtractDomains(t *testing.T) {
////	tests := []struct {
////		name     string
////		url      string
////		expected string
////	}{
////		{
////			name:     "Alibaba",
////			url:      "https://www.alibaba.com/",
//			expected: "1688.com,activities.alibaba.com,activity.alibaba.com,ads.alibaba.com,alibaba.com,alibabacloud.com,alibabagroup.com,aliexpress.com,alimama.com,alios.cn,aliqin.cn,app.alibaba.com,arabic.alibaba.com,autonavi.com,buyer.alibaba.com,dingtalk.com,dutch.alibaba.com,facebook.com,fliggy.com,french.alibaba.com,german.alibaba.com,global.alipay.com,hebrew.alibaba.com,hindi.alibaba.com,i.alibaba.com,i.alicdn.com,idinfo.zjaic.gov.cn,img.alicdn.com,importer.alibaba.com,indonesian.alibaba.com,inspection.alibaba.com,ipp.alibabagroup.com,italian.alibaba.com,japanese.alibaba.com,ju.taobao.com,korean.alibaba.com,linkedin.com,logistics.alibaba.com,m.alibaba.com,onetouch.alibaba.com,portuguese.alibaba.com,rfq.alibaba.com,rule.alibaba.com,russian.alibaba.com,seller.alibaba.com,spanish.alibaba.com,supplier.alibaba.com,surveymonkey.com,taobao.com,taobao.lazada.sg,thai.alibaba.com,tmall.com,tradeassurance.alibaba.com,trademanager.alibaba.com,turkish.alibaba.com,twitter.com,ucweb.com,umeng.com,vietnamese.alibaba.com,wholesaler.alibaba.com,world.taobao.com,xiami.com,youtube.com",
//		}, {
//			name:     "Starter Studio",
//			url:      "https://www.starterstudio.org/",
//			expected: "",
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			resp, err := http.Get(test.url)
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			defer resp.Body.Close()
//			bytes, err := ioutil.ReadAll(resp.Body)
//
//			lines := strings.Split(string(bytes), "\n")
//
//			if res := ExtractDomains(lines); res != test.expected {
//				t.Fatalf("%s returned unexpected result: %q", test.name, res)
//			}
//		})
//	}
//}
