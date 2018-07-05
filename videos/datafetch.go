package videos

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
	"log"
	"net/http"
	// "time"
)

const dbDest = "video-resource.db"

const TestXMLRss = `<?xml version="1.0" encoding="utf-8"?>
<rss version="5.1">
    <list page="1" pagecount="3368" pagesize="30" recordcount="101020">
        <video>
            <last>2017-09-09 10:44:16</last>
            <id>1</id>
            <tid>10</tid>
            <name>
                <![CDATA[无间道3 粤语]]>
            </name>
            <type>动作片</type>
            <pic>http://qicaicms.com/upload/vod/2017-11-16/15107736690.jpg</pic>
            <lang></lang>
            <area>内地</area>
            <year>2003</year>
            <state>0</state>
            <note>
                <![CDATA[高清]]>
            </note>
            <actor>
                <![CDATA[刘德华 梁朝伟 陈慧琳 黎明]]>
            </actor>
            <director>
                <![CDATA[刘伟强]]>
            </director>
            <dl>
                <dd flag="qiyi">
                    <![CDATA[]]>
                </dd>
            </dl>
            <des>
                <![CDATA[陈永仁被杀后，刘健明被内部调查，架空职权，无心工作，与妻子婚姻破裂，身心疲倦的刘健明，每日的生活就好像走向地狱一样，但他深信总有扭转局面的一日。同年，警队内有个年轻警司杨锦荣冒出，杨锦荣凭着功绩累累，正迈向警队最高层。令当时刘健明看到从前的自己。不久，二人成了好友，刘健明却发现杨锦荣神秘一面，一次2人展开追查一单案件时，刘健明发现被人跟踪，从内鬼得悉跟踪者竟是韩琛旧日拍档。]]>
            </des>
        </video>
        <video>
            <last>2017-09-09 10:44:17</last>
            <id>2</id>
            <tid>10</tid>
            <name>
                <![CDATA[让子弹飞（川话版）]]>
            </name>
            <type>动作片</type>
            <pic>http://qicaicms.com/upload/vod/2017-11-16/15107736510.jpg</pic>
            <lang></lang>
            <area>内地</area>
            <year>2010</year>
            <state>0</state>
            <note>
                <![CDATA[高清]]>
            </note>
            <actor>
                <![CDATA[姜文 葛优 周润发 刘嘉玲]]>
            </actor>
            <director>
                <![CDATA[姜文]]>
            </director>
            <dl>
                <dd flag="qiyi">
                    <![CDATA[]]>
                </dd>
            </dl>
            <des>
                <![CDATA[民国年间，花钱捐得县长的马邦德（葛优 饰）携妻（刘嘉玲 饰）及随从走马上任。途经南国某地，遭劫匪张麻子（姜文 饰）一伙伏击，随从尽死，只夫妻二人侥幸活命。马为保命，谎称自己是县长的汤师爷。为汤师爷许下的财富所动，张麻子摇身一变化身县长，带着手下赶赴鹅城上任。有道是天高皇帝远，鹅城地处偏僻，一方霸主黄四郎（周润发 饰）只手遮天，全然不将这个新来的县长放在眼里。张麻子痛打了黄的武教头（姜武 饰），黄则设计害死张的义子小六（张默 饰）。原本只想赚钱的马邦德，怎么也想不到竟会被卷入这场土匪和恶霸的角力之中。鹅城上空愁云密布，血雨腥风在所难免……]]>
            </des>
        </video>
    </list>
</rss>`

const TestXMLRssCatalog = `<?xml version="1.0" encoding="utf-8"?>
<rss version="5.1">
    <list page="1" pagecount="0" pagesize="30" recordcount="19246"></list>
    <class>
        <ty id="1">电影</ty>
        <ty id="2">连续剧</ty>
        <ty id="3">综艺</ty>
        <ty id="4">动漫</ty>
        <ty id="5">微电影</ty>
        <ty id="10">动作片</ty>
        <ty id="11">喜剧片</ty>
        <ty id="12">爱情片</ty>
        <ty id="13">科幻片</ty>
        <ty id="14">恐怖片</ty>
        <ty id="15">剧情片</ty>
        <ty id="16">战争片</ty>
        <ty id="17">纪录片</ty>
        <ty id="18">动画片</ty>
        <ty id="19">伦理片</ty>
        <ty id="20">国产剧</ty>
        <ty id="21">香港剧</ty>
        <ty id="22">台湾剧</ty>
        <ty id="23">日本剧</ty>
        <ty id="24">韩国剧</ty>
        <ty id="25">美国剧</ty>
        <ty id="26">英国剧</ty>
        <ty id="27">泰国剧</ty>
        <ty id="28">新加坡剧</ty>
        <ty id="29">其他剧</ty>
    </class>
</rss>
`

const apiSite1 = "http://zy.cmp4.cn/inc/api.php"
const apiSite2 = "http://qicaicms.com/inc/api.php"
const catalogSuffix = "?ac=list&t=1"
const videoSuffix = "?ac=videolist&pg=%v"

type Rss struct {
	gorm.Model
	// XMLName xml.Name  `xml:"rss"`
	Version string    `xml:"version,attr"`
	List    VideoList `xml:"list"`
}

type VideoList struct {
	// XMLName     xml.Name `xml:"list"`
	Page        int        `xml:"page,attr"`
	Pagecount   int        `xml:"pagecount,attr"`
	PageSize    int        `xml:"pagesize,attr"`
	Recordcount int        `xml:"recordcount,attr"`
	Videos      []VideoXml `xml:"video"`
}

type VideoXml struct {
	gorm.Model
	// XMLName  xml.Name `xml:"video"`
	Last     string   `xml:"last"`
	VID      int      `xml:"id"`
	Tid      int      `xml:"tdi"`
	Name     string   `xml:"name"`
	Type     string   `xml:"type"`
	Pic      string   `xml:"pic"`
	Lang     string   `xml:"lang"`
	Area     string   `xml:"area"`
	Year     string   `xml:"year"`
	State    string   `xml:"state"`
	Note     string   `xml:"note"`
	Actor    string   `xml:"actor"`
	Director string   `xml:"director"`
	DL       DataList `xml:"dl"`
	Desc     string   `xml:"des" json:"des"`
}

type Video struct {
	gorm.Model
	// XMLName  xml.Name `xml:"video"`
	Last     string `xml:"last"`
	VID      int    `xml:"id"`
	Tid      int    `xml:"tdi"`
	Name     string `xml:"name"`
	Type     string `xml:"type"`
	Pic      string `xml:"pic"`
	Lang     string `xml:"lang"`
	Area     string `xml:"area"`
	Year     string `xml:"year"`
	State    string `xml:"state"`
	Note     string `xml:"note"`
	Actor    string `xml:"actor,cdata"`
	Director string `xml:"director,cdata"`
	DL       string `xml:"dl"`
	Desc     string `xml:"des" json:"des"`
}

type DD struct {
	Flag   string
	Source string
}

type DataList struct {
	// XMLName xml.Name  `xml:"dl" json:"dl"`
	Data []DataRet `xml:"dd" json:"Data"`
}

type DataRet struct {
	// XMLName xml.Name `xml:"dd"`
	Flag    string `xml:"flag,attr" json:"Flag"`
	Content string `xml:",cdata" json:"Content"`
}

type RssCatalog struct {
	// XMLName xml.Name  `xml:"rss"`
	Version string      `xml:"version,attr"`
	List    CatalogList `xml:"list"`
	Class   TyList      `xml:"class"`
}

type TyList struct {
	// XMLName xml.Name `xml:"ty"`
	TyRet []Ty `xml:"ty"`
}

type Ty struct {
	gorm.Model
	TiD     int    `xml:"id,attr"`
	Content string `xml:",chardata"`
}

type CatalogList struct {
	// XMLName     xml.Name `xml:"list"`
	Page        int `xml:"page,attr"`
	Pagecount   int `xml:"pagecount,attr"`
	PageSize    int `xml:"pagesize,attr"`
	Recordcount int `xml:"recordcount,attr"`
}

func DBInit(dbFIle string) {
	db, err := gorm.Open("sqlite3", dbFIle)
	CheckErr(err)
	defer db.Close()
	db.AutoMigrate(&Video{})
	db.AutoMigrate(&Ty{})
}

func Fetch(num int, dbFile string) {
	DBInit(dbFile)
	db, err := gorm.Open("sqlite3", dbFile)
	CheckErr(err)
	defer db.Close()

	for i := 1; i < num; i++ {
		videoUrl := apiSite2 + fmt.Sprintf(videoSuffix, i)
		fmt.Println(videoUrl)
		store(db, videoUrl)
	}
}

func CatalogStore() {
	var err error
	DBInit(dbDest)
	db, err := gorm.Open("sqlite3", dbDest)
	CheckErr(err)
	defer db.Close()
	catalogUrl := apiSite2 + catalogSuffix
	fmt.Println(catalogUrl)
	catalog := GetCatalog(catalogUrl)
	for _, ty := range catalog.Class.TyRet {
		db.Create(&ty)
	}
}

func store(db *gorm.DB, url string) {
	var err error
	VideoRss := GetVideoList(url)
	var tmpDl []byte
	for _, v := range VideoRss.List.Videos {
		video := &Video{}
		video.Last = v.Last
		video.VID = v.VID
		video.Name = v.Name
		video.Type = v.Type
		video.Pic = v.Pic
		video.Lang = v.Lang
		video.Area = v.Area
		video.Year = v.Year
		video.State = v.State
		video.Note = v.Note
		video.Actor = v.Actor
		video.Director = v.Director
		tmpDl, err = json.Marshal(v.DL)
		CheckErr(err)
		video.DL = string(tmpDl)
		video.Desc = v.Desc
		// fmt.Println("video===>", v)
		db.Create(video)
	}
}

func GetCatalog(url string) *RssCatalog {
	var ret = RssCatalog{}
	resp, err := http.Get(url)
	CheckErr(err)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	err = xml.Unmarshal(data, &ret)
	CheckErr(err)
	return &ret
}

func GetVideoList(url string) *Rss {
	var ret = Rss{}
	resp, err := http.Get(url)
	CheckErr(err)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	err = xml.Unmarshal(data, &ret)
	CheckErr(err)
	return &ret
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
