package w

//go:noinline
func For[T any](getList func() []T, render func(T, int) *Ele) *Ele {
	point := HPoint()
	key, attrKey := RandomAttr()
	last := 0
	CreateEffect(func() {
		now := getList()
		if last == 0 {
			eles := Document.QuerySelectorAll(attrKey)
			for i := 0; i < len(eles); i++ {
				eles[i].Remove()
			}
			for i, v := range now {
				e := render(v, i).Attr(key, 1)
				point.InsertAdjacentElement("beforebegin", e)
			}
		} else {
			lenNow := len(now)
			// 清空
			if last < lenNow {
				// 新增新的
				for i := last; i < lenNow; i++ {
					e := render(now[i], i).Attr(key, 1)
					point.InsertAdjacentElement("beforebegin", e)
				}
			} else if last > lenNow {
				// 更新删除多余的
				eles := Document.QuerySelectorAll(attrKey)
				for i := lenNow; i < len(eles); i++ {
					eles[i].Remove()
				}
			}
		}
		last = len(now)
	})
	return point
}
