package errorder

/**
Description:插入token至Tokens表

Parameter:
	tokens:将要插入Tokens表的Token数组

Return:
	错误
*/
//插入Token
func InsertErrOrder(errorder ErrorOrder) error {
	_, err := db.Insert(errorder)
	if err != nil {
		log.Errorw("ErrorOrder insert", "error", err)
	}
	return err
}
