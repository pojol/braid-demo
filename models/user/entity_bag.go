package user

import (
	commproto "braid-demo/models/commproto"
	fmt "fmt"
)

func (b *EntityBagModule) produce(item *commproto.Item) {
	_, ok := b.Bag[item.DictID]

	// check if it is a auto-recovery item

	if !ok {
		b.Bag[item.DictID] = &commproto.ItemList{
			Items: []*commproto.Item{item},
		}
	} else {
		// check if it can be stacked
		// check if it overflows

	}
}

func (b *EntityBagModule) consume(dict bool, item *commproto.Item, num int32) (commproto.Item, error) {
	items, ok := b.Bag[item.DictID]
	refreshItem := commproto.Item{}

	if !ok {
		return refreshItem, fmt.Errorf("not found")
	}

	var rmv bool // remove tag
	var idx int
	var flag bool

	for k, v := range items.Items {
		equal := false

		if dict {
			equal = v.DictID == item.DictID
		} else {
			equal = v.ID == item.ID
		}

		if equal {
			if v.Num < num {
				return refreshItem, fmt.Errorf("not enough")
			}

			items.Items[k].Num -= item.Num
			if items.Items[k].Num == 0 /* not resource */ {
				rmv = true
				idx = k
			}

			flag = true
			break
		}
	}

	if !flag {
		return refreshItem, fmt.Errorf("not found")
	}

	if rmv {
		items.Items = append(items.Items[:idx], items.Items[idx+1:]...)
		if len(items.Items) == 0 {
			delete(b.Bag, item.DictID)
		}
	}

	return refreshItem, nil
}

func (b *EntityBagModule) enough(dict bool, item *commproto.Item) bool {
	items, ok := b.Bag[item.DictID]
	if !ok {
		return false
	}

	for _, v := range items.Items {
		equal := false

		if dict {
			equal = v.DictID == item.DictID
		} else {
			equal = v.ID == item.ID
		}

		if equal {

			b._checkRecoverIimeItem()
			b._checkTimeoutItem()

			if v.Num >= item.Num {
				return true
			}
			break
		}
	}

	return false
}

// _checkRecoverIimeItem - check auto-recovery item
func (b *EntityBagModule) _checkRecoverIimeItem() {

}

// _checkTimeoutItem - check timeout item
func (b *EntityBagModule) _checkTimeoutItem() {

}

// EnoughItemWithInsID - check if the item is enough with instance id (unique id), note: need to pass dictionary id
func (b *EntityBagModule) EnoughItemWithInsID(id string, dictid, num int32) bool {
	return b.enough(false, &commproto.Item{ID: id, DictID: dictid, Num: num})
}

// EnoughItem - check if the item is enough with dictionary id
func (b *EntityBagModule) EnoughItem(id, num int32) bool {
	return b.enough(true, &commproto.Item{DictID: id, Num: num})
}

// EnoughItems - check if the items are enough with dictionary id
func (b *EntityBagModule) EnoughItems(items []*commproto.Item) bool {
	for _, v := range items {
		if !b.enough(true, v) {
			return false
		}
	}

	return true
}

func (b *EntityBagModule) ProduceItem(item *commproto.Item, num uint32, reason, detail string) {

}

// ProduceItems - produce items
//
//	items: items to produce
//	reason: produce reason
//	detail: produce detail
func (b *EntityBagModule) ProduceItems(items []*commproto.Item, reason, detail string) {

}

// ExistItem - get item num with dictionary id
func (b *EntityBagModule) GetItemNum(id int32) int64 {
	var num int64

	if _, ok := b.Bag[id]; ok {

		for _, v := range b.Bag[id].Items {
			num += int64(v.Num)
		}

	}

	return num
}

// ConsumeItem - consume item (must check enough before consume)
func (b *EntityBagModule) ConsumeItem(id, num int32, reason, detail string) []*commproto.Item {
	refresh := []*commproto.Item{}

	ritem, err := b.consume(true, &commproto.Item{DictID: id, Num: num}, num)
	if err != nil {
		fmt.Errorf("consume item %v num %v reason %v error: %w", id, num, reason, err)
	}

	refresh = append(refresh, &commproto.Item{
		ID:     ritem.ID,
		DictID: ritem.DictID,
		Num:    ritem.Num,
	})

	// log record

	return refresh
}

// ConsumeItems - consume items (must check enough before consume)
func (b *EntityBagModule) ConsumeItems(items []*commproto.Item, reason, detail string) []*commproto.Item {

	refresh := []*commproto.Item{}

	for _, v := range items {
		ritem, err := b.consume(true, v, v.Num)
		if err != nil {
			fmt.Printf("consume item %v num %v reason %v error: %v", v.DictID, v.Num, reason, err)
		}

		refresh = append(refresh, &commproto.Item{
			ID:     ritem.ID,
			DictID: ritem.DictID,
			Num:    ritem.Num,
		})
	}

	// log record

	return refresh
}
