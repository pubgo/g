package compare

import (
	"errors"
	"github.com/DearMadMan/minhash"
	"github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
	"github.com/yanyiwu/gojieba"
	"go.uber.org/zap"
	"math"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

type CompareData struct {
	Url string `json:"url"`
	Src string `json:"src"`
	Des string `json:"des"`
}

type HashData struct {
	Content string `json:"content"`
}

var JieBa *gojieba.Jieba
var compareLock = sync.Mutex{}
var simLock = sync.Mutex{}

var mediaNameMap = map[string]string{
	"mp.weixin.qq.com":    "微信",
	"weixin.sogou.com":    "微信",
	"www.jianshu.com":     "简书",
	"www.zhihu.com":       "知乎",
	"zhuanlan.zhihu.com":  "知乎专栏",
	"www.toutiao.com":     "头条",
	"www.yidianzixun.com": "一点",
	"baijiahao.baidu.com": "百家",
	"www.sohu.com":        "搜狐",
	"m.sohu.com":          "搜狐",
	"xueqiu.com":          "雪球",
	"www.myzaker.com":     "zaker",
	"item.btime.com":      "北京时间",
	"kuaibao.qq.com":      "快报",
	"news.163.com":        "网易",
	"news.sina.com.cn":    "新浪",
	"news.qq.com":         "腾讯新闻",
}

func GetMediaName(desUrl string) string {
	mediaName := "网站"
	if strings.Contains(desUrl, "www.zhihu.com/question") {
		mediaName = "知乎问答"
	} else {
		parser, err := url.Parse(desUrl)
		if err != nil {
			zap.S().Infow("desUrl parser error -->", "error info", err.Error())
		} else {
			if name, ok := mediaNameMap[parser.Host]; ok && name != "" {
				mediaName = name
			}
		}
	}
	return mediaName
}

func GetSimCount(src, des string) int {
	simLock.Lock()
	similar := "" // 相似字数
	num := 0      // index
	step := 8     // 步长
	srcRune := []rune(src)
	srcLen := len(srcRune)
	for num < srcLen {
		if num+step >= srcLen {
			break
		}

		if strings.Contains(des, string(srcRune[num:num+step])) {
			for {
				if num+step+1 <= srcLen && strings.Contains(des, string(srcRune[num:num+step+1])) {
					step += 1
					continue
				} else {
					sim_temp := string(srcRune[num : num+step]) // 相似片段
					similar += sim_temp
					des = strings.Replace(des, sim_temp, "", 1)
					num += step
					step = 8
					break
				}
			}

		} else {
			num += 1
		}

	}
	simLock.Unlock()
	return len(strings.Split(regexp.MustCompile("[\u4e00-\u9fa5]").ReplaceAllString(similar, "a "), " ")) - 1
}

func GetScore(src, des string) float64 {
	compareLock.Lock()
	//desSet := mapset.NewSet(JieBa.Cut(des, true))
	srcSet := mapset.NewSet()
	for _, _i := range JieBa.Cut(src, true) {
		srcSet.Add(_i)
	}
	desSet := mapset.NewSet()
	for _, _i := range JieBa.Cut(des, true) {
		desSet.Add(_i)
	}
	compareLock.Unlock()
	return float64(len(srcSet.Intersect(desSet).ToSlice())) / float64(len(srcSet.Union(desSet).ToSlice()))
}

func Compare(c *gin.Context) {
	start := time.Now()
	compareData := CompareData{}
	if err := c.ShouldBindJSON(&compareData); err != nil {
		zap.S().Infow("参数解析失败 ==>", "error", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 3001, "message": err.Error()})
		return
	}

	// del space
	compareData.Src = strings.ReplaceAll(compareData.Src, " ", "")
	compareData.Des = strings.ReplaceAll(compareData.Des, " ", "")

	if compareData.Src == "" || compareData.Des == "" || compareData.Url == "" {
		c.JSON(http.StatusOK, gin.H{"code": 3001, "message": errors.New("参数不能为空！！！")})
		return
	}
	// score
	start1 := time.Now()
	compareScore := GetScore(compareData.Src, compareData.Des)
	zap.S().Infow("比对耗时", "waste_time", time.Since(start1).String())

	// sim count
	start2 := time.Now()
	simCount := GetSimCount(compareData.Src, compareData.Des)
	zap.S().Infow("计算相似字数耗时", "waste_time", time.Since(start2).String())

	// des count
	start3 := time.Now()
	desCount := len(strings.Split(regexp.MustCompile("[\u4e00-\u9fa5]").ReplaceAllString(compareData.Des, "a "), " ")) - 1
	zap.S().Infow("统计文章字数耗时", "waste_time", time.Since(start3).String())

	// media name
	mediaName := GetMediaName(compareData.Url)
	//zap.S().Infow("response result is ==>", "score", compareScore, "mediaName", mediaName, "desCount", desCount, "simCount", simCount)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": map[string]interface{}{
			"score":      compareScore,
			"sim_words":  simCount,
			"word_count": desCount,
			"media_name": mediaName,
		},
	})
	zap.S().Infow("任务处理耗时", "score", compareScore, "waste_time", time.Since(start).String(), "mediaName", mediaName, "desCount", desCount, "simCount", simCount)

	return

}

func MinHash(c *gin.Context) {
	hashData := HashData{}

	if err := c.ShouldBindJSON(&hashData); err != nil {
		zap.S().Infow("参数解析失败 ==>", "error", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 3001, "message": err.Error()})
		return
	}

	if hashData.Content == "" {
		c.JSON(http.StatusOK, gin.H{"code": 3001, "message": errors.New("参数不能为空！！！")})
		return
	}

	m := minhash.New(126)
	set := m.NewSet(JieBa.Cut(hashData.Content, true))

	b := 9
	r := 14
	lp := uint64(18446744073709551557)

	hashes := []uint64{}

	for i := 0; i < b; i++ {
		hs := uint64(0)
		for j := 0; j < r; j++ {
			current := uint64(set.Signatures[i*r+j])
			hs += current * uint64(math.Pow(7, float64(j)))
			hs %= lp
		}
		hashes = append(hashes, hs)
	}
}

func InitJieBa() {
	cur, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	DICT_DIR := filepath.Join(cur, "dict")
	DICT_PATH := filepath.Join(DICT_DIR, "jieba.dict.utf8")
	HMM_PATH := filepath.Join(DICT_DIR, "hmm_model.utf8")
	USER_DICT_PATH := filepath.Join(DICT_DIR, "user.dict.utf8")
	IDF_PATH := filepath.Join(DICT_DIR, "idf.utf8")
	STOP_WORDS_PATH := filepath.Join(DICT_DIR, "stop_words.utf8")

	//JieBa = gojieba.NewJieba(gojieba.DICT_PATH, gojieba.HMM_PATH, gojieba.USER_DICT_PATH, gojieba.IDF_PATH, gojieba.STOP_WORDS_PATH)
	JieBa = gojieba.NewJieba(DICT_PATH, HMM_PATH, USER_DICT_PATH, IDF_PATH, STOP_WORDS_PATH)
}

func main() {
	start0 := time.Now()
	InitJieBa()
	zap.S().Infow("初始化结巴分词耗时", "waste_time", time.Since(start0).String())

	defer JieBa.Free()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// compare
	r.POST("/similarities", Compare)

	if err := r.Run(":8080"); err != nil {
		panic(err.Error())
	}
}
