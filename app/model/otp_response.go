package model

type OTPResponseData struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
	Result struct {
		Token   string `json:"token"`
		RefCode string `json:"ref_code"`
	} `json:"result"`
}

type OTPVaildateResponseData struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
	Result struct {
		Status bool `json:"status"`
	} `json:"result"`
}

// {
// 	"code": "000",
// 	"detail": "OK.",
// 	"result": {
// 	  "status": true
// 	}
//   }

//   {
// 	"code": "1006",
// 	"detail": "Token is invalid.",
// 	"result": null
//   }

// {
// 	"code": "000",
// 	"detail": "OK.",
// 	"result": {
// 	  "status": false
// 	}
//   }
