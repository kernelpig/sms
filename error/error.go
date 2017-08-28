package error

// isv/isp错误码, 参考 https://help.aliyun.com/document_detail/55284.html?spm=5176.doc56189.6.552.fp3vqU
const (
	RamPermissionDeny         = "isp.RAM_PERMISSION_DENY"         //RAM权限DENY
	OutOfService              = "isv.OUT_OF_SERVICE"              //业务停机
	ProductUnSubscript        = "isv.PRODUCT_UN_SUBSCRIPT"        //未开通云通信产品的阿里云客户
	ProductUnsubscribe        = "isv.PRODUCT_UNSUBSCRIBE"         //产品未开通
	AccountNotExists          = "isv.ACCOUNT_NOT_EXISTS"          //账户不存在
	AccountAbnormal           = "isv.ACCOUNT_ABNORMAL"            //账户异常
	SmsTemplate_Illegal       = "isv.SMS_TEMPLATE_ILLEGAL"        //短信模板不合法
	SmsSignatureIllegal       = "isv.SMS_SIGNATURE_ILLEGAL"       //短信签名不合法
	InvalidParameters         = "isv.INVALID_PARAMETERS"          //参数异常
	SystemError               = "isp.SYSTEM_ERROR"                //系统错误
	MobileNumberIllegal       = "isv.MOBILE_NUMBER_ILLEGAL"       //非法手机号
	MobileCountOverLimit      = "isv.MOBILE_COUNT_OVER_LIMIT"     //手机号码数量超过限制
	TemplateMissingParameters = "isv.TEMPLATE_MISSING_PARAMETERS" //模板缺少变量
	BusinessLimitControl      = "isv.BUSINESS_LIMIT_CONTROL"      //业务限流
	InvalidJsonParam          = "isv.INVALID_JSON_PARAM"          //JSON参数不合法，只接受字符串值
	BlackKeyControlLimit      = "isv.BLACK_KEY_CONTROL_LIMIT"     //黑名单管控
	ParamLengthLimit          = "isv.PARAM_LENGTH_LIMIT"          //参数超出长度限制
	ParamNotSupportUrl        = "isv.PARAM_NOT_SUPPORT_URL"       //不支持URL
	AmountNotEnough           = "isv.AMOUNT_NOT_ENOUGH"           //账户余额不足
)
