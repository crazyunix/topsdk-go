package request

type TaobaoTbkDgVegasSendStatusRequest struct {
	/*
	   渠道管理id     */
	RelationId *string `json:"relation_id,omitempty" required:"false" `
	/*
	   会员运营id     */
	SpecialId *string `json:"special_id,omitempty" required:"false" `
	/*
	   加密后的值，需用MD5加密，32位小写     */
	DeviceValue *string `json:"device_value,omitempty" required:"false" `
	/*
	   入参类型(该模式下返回的结果为模糊匹配结果，和实际情况可能存在误差)： 1. IMEI 2. IDFA 3. OAID 4. MOBILE     */
	DeviceType *string `json:"device_type,omitempty" required:"false" `
	/*
	   已废弃，请勿传入     */
	ThorBizCode *string `json:"thor_biz_code,omitempty" required:"false" `
	/*
	   媒体pid     */
	Pid *string `json:"pid,omitempty" required:"false" `
	/*
	   查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，5-幸运赢免单，不传时默认查询超级红包数据     */
	ActivityCategory *int64 `json:"activity_category,omitempty" required:"false" `
}

func (s *TaobaoTbkDgVegasSendStatusRequest) SetRelationId(v string) *TaobaoTbkDgVegasSendStatusRequest {
	s.RelationId = &v
	return s
}
func (s *TaobaoTbkDgVegasSendStatusRequest) SetSpecialId(v string) *TaobaoTbkDgVegasSendStatusRequest {
	s.SpecialId = &v
	return s
}
func (s *TaobaoTbkDgVegasSendStatusRequest) SetDeviceValue(v string) *TaobaoTbkDgVegasSendStatusRequest {
	s.DeviceValue = &v
	return s
}
func (s *TaobaoTbkDgVegasSendStatusRequest) SetDeviceType(v string) *TaobaoTbkDgVegasSendStatusRequest {
	s.DeviceType = &v
	return s
}
func (s *TaobaoTbkDgVegasSendStatusRequest) SetThorBizCode(v string) *TaobaoTbkDgVegasSendStatusRequest {
	s.ThorBizCode = &v
	return s
}
func (s *TaobaoTbkDgVegasSendStatusRequest) SetPid(v string) *TaobaoTbkDgVegasSendStatusRequest {
	s.Pid = &v
	return s
}
func (s *TaobaoTbkDgVegasSendStatusRequest) SetActivityCategory(v int64) *TaobaoTbkDgVegasSendStatusRequest {
	s.ActivityCategory = &v
	return s
}

func (req *TaobaoTbkDgVegasSendStatusRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RelationId != nil {
		paramMap["relation_id"] = *req.RelationId
	}
	if req.SpecialId != nil {
		paramMap["special_id"] = *req.SpecialId
	}
	if req.DeviceValue != nil {
		paramMap["device_value"] = *req.DeviceValue
	}
	if req.DeviceType != nil {
		paramMap["device_type"] = *req.DeviceType
	}
	if req.ThorBizCode != nil {
		paramMap["thor_biz_code"] = *req.ThorBizCode
	}
	if req.Pid != nil {
		paramMap["pid"] = *req.Pid
	}
	if req.ActivityCategory != nil {
		paramMap["activity_category"] = *req.ActivityCategory
	}
	return paramMap
}

func (req *TaobaoTbkDgVegasSendStatusRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
