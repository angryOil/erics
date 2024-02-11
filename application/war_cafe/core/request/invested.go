package request

import "context"

// 구조상 문제

/*
1. 여러번 실행되면 안된다: 악용 사례, <-- 시스템
2. 누군지 모른다: 실행 책임자 모른다 <-- auth 문제
*/

type InvestedRequest struct {
	Ctx                         context.Context
	ID                          int64
	WarCafeID                   int64
	CurrentNumberOfParticipants int
	AddNumberOfParticipants     int
}
