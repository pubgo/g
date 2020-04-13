package tests

import (
	"fmt"
	"github.com/JesusIslam/tldr"
	"github.com/go-ego/gse"
	"github.com/pubgo/g/xerror"
	"github.com/yanyiwu/gojieba"
	"strings"
	"testing"
)

func TestGse(t *testing.T) {
	text := `使用LexRank的golang文本摘要程序
tldr是一个golang软件包，可使用lexrank算法自动总结文本。
词汇排名，称量和排名有两个主要步骤。tldr包括两个加权和两个排序算法，分别是Jaccard系数和汉明距离，然后分别是PageRank和中心度。默认设置使用汉明距离和页面等级。`
	var seg gse.Segmenter
	xerror.Panic(seg.LoadDict())

	// use DAG and HMM
	hmm := seg.Cut(text, true)
	fmt.Printf("cut use hmm: %#v \n\n", hmm)
	//
	cut := seg.Cut(text)
	fmt.Printf("cut: %#v \n\n", cut)

	hmm = seg.CutSearch(text, true)
	fmt.Printf("cut search use hmm: %#v \n\n", hmm)
	//
	cut = seg.CutSearch(text)
	fmt.Printf("cut search: %#v \n\n", cut)

	cut = seg.CutAll(text)
	fmt.Printf("cut all: %#v \n\n", cut)

	fmt.Println(seg.Slice([]byte(text)))
}

func TestTldr(t *testing.T) {
	//	https://github.com/JesusIslam/tldr

	var seg gse.Segmenter
	xerror.Panic(seg.LoadDict())

	intoSentences := 1
	bag := tldr.New()
	bag.SetWordTokenizer(func(s string) []string {
		fmt.Printf("%#v \n\n", seg.Cut(s))
		return seg.Cut(s)
	})

	fmt.Println(bag.OriginalSentences)

	result := xerror.PanicErr(bag.Summarize(`
先说一下自动文摘的方法。自动文摘（Automatic Summarization）的方法主要有两种：Extraction和Abstraction。其中Extraction是抽取式自动文摘方法，通过提取文档中已存在的关键词，句子形成摘要；Abstraction是生成式自动文摘方法，通过建立抽象的语意表示，使用自然语言生成技术，形成摘要。由于生成式自动摘要方法需要复杂的自然语言理解和生成技术支持，应用领域受限。所以本人学习的也是抽取式的自动文摘方法。

　　目前主要方法有：

基于统计：统计词频，位置等信息，计算句子权值，再简选取权值高的句子作为文摘，特点：简单易用，但对词句的使用大多仅停留在表面信息。
基于图模型：构建拓扑结构图，对词句进行排序。例如，TextRank/LexRank
基于潜在语义：使用主题模型，挖掘词句隐藏信息。例如，采用LDA，HMM
基于整数规划：将文摘问题转为整数线性规划，求全局最优解。
`, intoSentences))
	fmt.Println(result)
}

func TestTldr1(t *testing.T) {

	var text = `
清晨6点30分的大连，街上的行人很少，路灯也还泛着黄晕。王霞出了家门准备搭乘最早的班车赶往公司。途经星海公园东门时，她恍惚看到有人躺在路边，走近才发现，一位白头老翁倒地昏迷并且不停地抽搐……

上班途中路遇老人昏迷

她选择施救

事情发生在2019年12月18日。早晨6点30分，天还没亮，王霞同往常一样，快步走出家门准备搭乘最早的班车赶往公司。她途经星海公园东门准备过马路时，依稀看到路边躺着个老人。走近一看，她发现老人的整个身子都在抽搐，口里吐着白沫……

看到这一幕，王霞赶紧施救。

王霞：我俯下身喊他，看他是不是还有意识，想知道他身上有没有急救药。但是当时老人的手疯狂地抓我的衣服，没做任何回答。

王霞说，这个突发状况让她有些措手不及，但她决定要留下来帮助老人。

图片

半小时，8通急救电话

架起生命线

王霞赶紧拨打120急救电话，咨询急救人员应该如何处理。在急救人员的电话指导下，王霞移动老人双腿，让老人尽量保持身体平衡。

为了确认老人的情况变化，之后120又打来电话，让王霞观察老人并做详细汇报。“不要睡着，120马上就到。不要往后仰，累了就靠在我身上……”为避免老人的呕吐物影响呼吸，王霞用巾纸擦拭着老人的嘴角，并让老人靠在自己的怀里。

时间在这一刻仿佛凝固了。在王霞的照顾下，慢慢地老人的脸色不再青紫，嘴里的白沫也在减少。半个小时里，王霞接听、拨打了8次120，报告老人的最新情况，方便医生判断老人的病情，为老人架起了一条生命线。

在她的感染下

不少人围过来帮忙

当天气温极低，见老人没穿棉袄，王霞又将自己的羽绒服给老人披上。在她的带动下，过来了不少好心人，有路人提供了手套，还有一位冬泳大姐拿出了小毛毯。

王霞不停地跟老人说话，安慰他、鼓励他。身上仅穿着羊毛衫的王霞早已被冻透，腿打哆嗦，牙齿不停地打颤。7点整，医务人员抵达现场，用急救车载走了老人，王霞这才松了口气。

做了好事却遭到“批评”

“你脱掉的不是羽绒服

是你的命”

当天由于上班迟到，王霞不得不将自己的遭遇告知了单位领导和同事。就在大家佩服王霞的做法时，她却遭到了一个人的批评，这人就是她的主治医生。原来在2017年，王霞就患上了重疾，还做过大手术，身体虚弱敏感，免疫力极低，此后一直处在恢复期。

图片

王霞的遭遇被主治医生知道后，对方在微信里给予她这样严厉的“警告”：“你脱掉的不是羽绒服，是你的命！”

爱与善良

是一种天性

记者了解到，王霞是丰田工机（大连）有限公司的员工。当天因为救人错过了班车，王霞只能自己花钱打车上班，虽然迟到了，但在早会上公司全员都为她鼓掌。

图片

“其实现在想想也很后怕，担心帮不了老人太多。”王霞说，当时没有时间顾及太多，如今最牵挂的还是老人的病情。“可能爱与善良就是一种天性。”她说。

一件羽绒服

温暖了整个冬天
`

	text = strings.ReplaceAll(text, "。", ".")
	text = strings.ReplaceAll(text, "；", ".")
	text = strings.ReplaceAll(text, "！", ".")

	x := gojieba.NewJieba()
	defer x.Free()

	fmt.Println(x.Cut(text, true))
	fmt.Println(x.Tag(text))

	fmt.Println(x.Extract(text, 5))

	//fmt.Println(x.ExtractWithWeight(text, -1))

	var ss = make(map[string]int)
	for _, d1 := range x.Tokenize(text, gojieba.SearchMode, true) {
		ss[d1.Str] += 1
	}

	intoSentences := 3
	bag := tldr.New()
	bag.SetWordTokenizer(func(s string) []string {
		return x.Cut(s, true)
	})
	bag.SetDictionary(ss)

	fmt.Print("\n\n\n\n")

	result := xerror.PanicErr(bag.Summarize(text, intoSentences)).([]string)
	for _, fd := range result {
		fmt.Println(fd)
	}
}

// https://github.com/arjunmahishi/text-summary
// https://github.com/aquilax/go-summary
// https://github.com/aschlapsi/go-summary/blob/master/main.go
// https://github.com/deckarep/golang-set
