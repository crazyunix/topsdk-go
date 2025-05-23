package domain

type TaobaoTbkDgCpaActivityDetailResults struct {
	/*
	   奖励明细数据，KV结构。字段释义见文档：https://www.yuque.com/docs/share/7ecf8cf1-7f99-4633-a2ed-f9b6f8116af5?#     */
	FieldDetail *string `json:"field_detail,omitempty" `

	/*
	   明细类型，1：预估明细，2：结算明细     */
	CalcType *int64 `json:"calc_type,omitempty" `

	/*
	   明细记录主键id     */
	Id *int64 `json:"id,omitempty" `
}

func (s *TaobaoTbkDgCpaActivityDetailResults) SetFieldDetail(v string) *TaobaoTbkDgCpaActivityDetailResults {
	s.FieldDetail = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailResults) SetCalcType(v int64) *TaobaoTbkDgCpaActivityDetailResults {
	s.CalcType = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailResults) SetId(v int64) *TaobaoTbkDgCpaActivityDetailResults {
	s.Id = &v
	return s
}
