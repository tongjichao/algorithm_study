package union_find

type UF struct {
	Id    []int64
	Count int64
}

//初始化UF
func NewUF(n int64) *UF {
	uf := new(UF)
	uf.Count = n
	//	var arr [5]int64

	for i := int64(0); i < n; i++ {
		uf.Id = append(uf.Id, i)
	}

	return uf
}

/*
func (uf *UF) Count() int64 {
	return uf.Count
}
*/

func (uf *UF) Connected(p, q int64) bool {
	return uf.FindQuick(p) == uf.FindQuick(q)
}

//三种find方式
// quick-find

func (uf *UF) FindQuick(n int64) int64 {
	return uf.Id[n]
}

func (uf *UF) UnionFindQuick(p, q int64) {

	pId := uf.FindQuick(p)
	qId := uf.FindQuick(q)

	if pId == qId {
		return
	}

	for i := 0; i < len(uf.Id); i++ {
		if uf.Id[i] == pId {
			uf.Id[i] = qId
		}
	}

	uf.Count--
}
