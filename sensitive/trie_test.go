package sensitive

import "testing"

func TestTrie_HasSensitive(t *testing.T) {
	SENSITIVE_WORD_TYPE_DESC := []string{
		"百刀网", "逼", "冰妹陪玩", "弹簧跳刀", "出售小号",
	}
	SensTREE := NewTrieByWords(SENSITIVE_WORD_TYPE_DESC)

	type TC struct {
		Word   string
		Expect bool
	}

	inputs := []TC{
		{Word: "百刀网", Expect: true},
		{Word: "人民", Expect: false},
		{Word: "冰妹陪玩", Expect: true},
		{Word: "弹簧跳刀", Expect: true},
		{Word: "冰妹陪玩", Expect: true},
		{Word: "大家好", Expect: false},
		{Word: "出售小号", Expect: true},
		{Word: "haha出售小号hhh1", Expect: true},
		{Word: "haha大家好hhh1", Expect: false},
		{Word: "大家好弹簧跳刀大家好", Expect: true},
		{Word: "逼", Expect: true},
		{Word: "妈", Expect: false},
	}

	for _, v := range inputs {
		got := SensTREE.HasSensitive(v.Word)
		expect := v.Expect

		if got != expect {

			t.Errorf("input %s,got %v,expect :%v", v.Word, got, expect)
		}
	}
}
