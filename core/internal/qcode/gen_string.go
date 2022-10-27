// Code generated by "stringer -linecomment -type=QType,MType,SelType,SkipType,PagingType,AggregrateOp,ValType,ExpOp -output=./gen_string.go"; DO NOT EDIT.

package qcode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[QTUnknown-0]
	_ = x[QTQuery-1]
	_ = x[QTSubscription-2]
	_ = x[QTMutation-3]
	_ = x[QTInsert-4]
	_ = x[QTUpdate-5]
	_ = x[QTDelete-6]
	_ = x[QTUpsert-7]
}

const _QType_name = "UnknownQuerySubcriptionMutationInsertUpdateDeleteUpsert"

var _QType_index = [...]uint8{0, 7, 12, 23, 31, 37, 43, 49, 55}

func (i QType) String() string {
	if i < 0 || i >= QType(len(_QType_index)-1) {
		return "QType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _QType_name[_QType_index[i]:_QType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MTInsert-1]
	_ = x[MTUpdate-2]
	_ = x[MTUpsert-3]
	_ = x[MTDelete-4]
	_ = x[MTConnect-5]
	_ = x[MTDisconnect-6]
	_ = x[MTNone-7]
	_ = x[MTKeyword-8]
}

const _MType_name = "MTInsertMTUpdateMTUpsertMTDeleteMTConnectMTDisconnectMTNoneMTKeyword"

var _MType_index = [...]uint8{0, 8, 16, 24, 32, 41, 53, 59, 68}

func (i MType) String() string {
	i -= 1
	if i >= MType(len(_MType_index)-1) {
		return "MType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _MType_name[_MType_index[i]:_MType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SelTypeNone-0]
	_ = x[SelTypeUnion-1]
	_ = x[SelTypeMember-2]
}

const _SelType_name = "SelTypeNoneSelTypeUnionSelTypeMember"

var _SelType_index = [...]uint8{0, 11, 23, 36}

func (i SelType) String() string {
	if i < 0 || i >= SelType(len(_SelType_index)-1) {
		return "SelType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SelType_name[_SelType_index[i]:_SelType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SkipTypeNone-0]
	_ = x[SkipTypeUserNeeded-1]
	_ = x[SkipTypeBlocked-2]
	_ = x[SkipTypeRemote-3]
}

const _SkipType_name = "SkipTypeNoneSkipTypeUserNeededSkipTypeBlockedSkipTypeRemote"

var _SkipType_index = [...]uint8{0, 12, 30, 45, 59}

func (i SkipType) String() string {
	if i < 0 || i >= SkipType(len(_SkipType_index)-1) {
		return "SkipType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SkipType_name[_SkipType_index[i]:_SkipType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PTOffset-0]
	_ = x[PTForward-1]
	_ = x[PTBackward-2]
}

const _PagingType_name = "PTOffsetPTForwardPTBackward"

var _PagingType_index = [...]uint8{0, 8, 17, 27}

func (i PagingType) String() string {
	if i < 0 || i >= PagingType(len(_PagingType_index)-1) {
		return "PagingType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PagingType_name[_PagingType_index[i]:_PagingType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AgCount-1]
	_ = x[AgSum-2]
	_ = x[AgAvg-3]
	_ = x[AgMax-4]
	_ = x[AgMin-5]
}

const _AggregrateOp_name = "AgCountAgSumAgAvgAgMaxAgMin"

var _AggregrateOp_index = [...]uint8{0, 7, 12, 17, 22, 27}

func (i AggregrateOp) String() string {
	i -= 1
	if i < 0 || i >= AggregrateOp(len(_AggregrateOp_index)-1) {
		return "AggregrateOp(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _AggregrateOp_name[_AggregrateOp_index[i]:_AggregrateOp_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ValStr-1]
	_ = x[ValNum-2]
	_ = x[ValBool-3]
	_ = x[ValList-4]
	_ = x[ValVar-5]
	_ = x[ValNone-6]
}

const _ValType_name = "ValStrValNumValBoolValListValVarValNone"

var _ValType_index = [...]uint8{0, 6, 12, 19, 26, 32, 39}

func (i ValType) String() string {
	i -= 1
	if i < 0 || i >= ValType(len(_ValType_index)-1) {
		return "ValType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ValType_name[_ValType_index[i]:_ValType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OpNop-0]
	_ = x[OpAnd-1]
	_ = x[OpOr-2]
	_ = x[OpNot-3]
	_ = x[OpEquals-4]
	_ = x[OpNotEquals-5]
	_ = x[OpGreaterOrEquals-6]
	_ = x[OpLesserOrEquals-7]
	_ = x[OpGreaterThan-8]
	_ = x[OpLesserThan-9]
	_ = x[OpIn-10]
	_ = x[OpNotIn-11]
	_ = x[OpLike-12]
	_ = x[OpNotLike-13]
	_ = x[OpILike-14]
	_ = x[OpNotILike-15]
	_ = x[OpSimilar-16]
	_ = x[OpNotSimilar-17]
	_ = x[OpRegex-18]
	_ = x[OpNotRegex-19]
	_ = x[OpIRegex-20]
	_ = x[OpNotIRegex-21]
	_ = x[OpContains-22]
	_ = x[OpContainedIn-23]
	_ = x[OpHasInCommon-24]
	_ = x[OpHasKey-25]
	_ = x[OpHasKeyAny-26]
	_ = x[OpHasKeyAll-27]
	_ = x[OpIsNull-28]
	_ = x[OpIsNotNull-29]
	_ = x[OpTsQuery-30]
	_ = x[OpFalse-31]
	_ = x[OpNotDistinct-32]
	_ = x[OpDistinct-33]
	_ = x[OpEqualsTrue-34]
	_ = x[OpNotEqualsTrue-35]
	_ = x[OpSelectExists-36]
}

const _ExpOp_name = "OpNopOpAndOpOrOpNotOpEqualsOpNotEqualsOpGreaterOrEqualsOpLesserOrEqualsOpGreaterThanOpLesserThanOpInOpNotInOpLikeOpNotLikeOpILikeOpNotILikeOpSimilarOpNotSimilarOpRegexOpNotRegexOpIRegexOpNotIRegexOpContainsOpContainedInOpHasInCommonOpHasKeyOpHasKeyAnyOpHasKeyAllOpIsNullOpIsNotNullOpTsQueryOpFalseOpNotDistinctOpDistinctOpEqualsTrueOpNotEqualsTrueOpSelectExists"

var _ExpOp_index = [...]uint16{0, 5, 10, 14, 19, 27, 38, 55, 71, 84, 96, 100, 107, 113, 122, 129, 139, 148, 160, 167, 177, 185, 196, 206, 219, 232, 240, 251, 262, 270, 281, 290, 297, 310, 320, 332, 347, 361}

func (i ExpOp) String() string {
	if i < 0 || i >= ExpOp(len(_ExpOp_index)-1) {
		return "ExpOp(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ExpOp_name[_ExpOp_index[i]:_ExpOp_index[i+1]]
}
