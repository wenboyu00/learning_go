package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("city_list_test_file.html")
	if err != nil {
		panic(err)
	}
	//预期值
	const resultSize = 470
	var expectedCites = []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}
	var expectedUrls = []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	result := ParseCityList(contents)
	// 测试解析数量
	if len(result.Items) != resultSize {
		t.Errorf("result shoult have %d requests; but had %d",
			resultSize, len(result.Items))
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result shoult have %d requests; but had %d",
			resultSize, len(result.Requests))
	}
	// 测试解析城市
	for i, city := range expectedCites {
		if result.Items[i] != city {
			t.Errorf("expected city #%d: %s; but was %s",
				i, city, result.Items[i])
		}
	}
	// 测试解析url
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s",
				i, url, result.Requests[i].Url)
		}
	}
}
