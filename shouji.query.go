package jisuapi

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type ShoujiQueryResponse struct {
	Status int64  `json:"status"` // 状态码
	Msg    string `json:"msg"`
	Result struct {
		Province string `json:"province"` // 省
		City     string `json:"city"`     // 市
		Company  string `json:"company"`  // 运营商
		Cardtype string `json:"cardtype"` // 卡类型
	} `json:"result"`
}

type ShoujiQueryResult struct {
	Result ShoujiQueryResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newShoujiQueryResult(result ShoujiQueryResponse, body []byte, http gorequest.Response) *ShoujiQueryResult {
	return &ShoujiQueryResult{Result: result, Body: body, Http: http}
}

// ShoujiQuery 手机号码归属地
// https://www.jisuapi.com/api/shouji/
func (c *Client) ShoujiQuery(ctx context.Context, shouji string, appkey string, notMustParams ...gorequest.Params) (*ShoujiQueryResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("shouji", shouji) // 手机号
	// 请求
	request, err := c.request(ctx, apiUrl+"/shouji/query?appkey="+appkey, params, http.MethodGet)
	if err != nil {
		return newShoujiQueryResult(ShoujiQueryResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ShoujiQueryResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newShoujiQueryResult(response, request.ResponseBody, request), err
}
