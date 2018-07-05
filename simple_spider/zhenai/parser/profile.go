package parser

import (
	"learngo/simple_spider/engine"
	"learngo/simple_spider/model"
	"log"
	"regexp"
	"strconv"
)

var GrenderRe = regexp.MustCompile(
	`<span class="label">性别：</span><span field="">([^<]+)</span>`)
var AgeRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>([0-9]+)岁</td>`)
var HeightRe = regexp.MustCompile(
	`<span class="label">身高：</span><span field="">([0-9]+)CM</span>`)
var WeightRe = regexp.MustCompile(
	`<span class="label">体重：</span><span field="">([0-9]+)KG</span>`)
var IncomeRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([^<]+)</td>`)
var MarriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
var EducationRe = regexp.MustCompile(
	` <td><span class="label">学历：</span>([^<]+)</td>`)
var OccupationRe = regexp.MustCompile(
	`<span class="label">职业：</span><span field="">([^<]+)</span>`)
var hoKouRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var XinzuoRe = regexp.MustCompile(
	`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var HouseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var CarRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var IdRe = regexp.MustCompile(
	`http://album.zhenai.com/u/(\d+)`)

func ParserProfile(content []byte, url, name string) engine.ParserResult {
	grender := extractString(content, GrenderRe)
	age, err := strconv.Atoi(extractString(content, AgeRe))
	if err != nil {
		log.Printf("Parse user age error, %s", err)
	}

	height, err := strconv.Atoi(extractString(content, HeightRe))
	if err != nil {
		log.Printf("Parse user height error: %s", err)
	}
	weight, err := strconv.Atoi(extractString(content, WeightRe))
	if err != nil {
		log.Printf("Parse user weight error: %s", err)
	}
	income := extractString(content, IncomeRe)
	marriage := extractString(content, MarriageRe)
	education := extractString(content, EducationRe)
	occupation := extractString(content, OccupationRe)
	hokou := extractString(content, hoKouRe)
	xinzuo := extractString(content, XinzuoRe)
	house := extractString(content, HouseRe)
	car := extractString(content, CarRe)
	id := extractString([]byte(url), IdRe)

	var profile model.Profile
	profile.Name = name
	profile.Gender = grender
	profile.Age = age
	profile.Height = height
	profile.Weight = weight
	profile.Income = income
	profile.Marriage = marriage
	profile.Education = education
	profile.Occupation = occupation
	profile.HoKou = hokou
	profile.Xinzuo = xinzuo
	profile.House = house
	profile.Car = car

	var item model.Item
	item.Url = url
	item.Id = id
	item.Type = "zhenai"
	item.Payload = profile

	result := engine.ParserResult{}

	result.Items = append(result.Items, item)
	return result

}

func extractString(content []byte, re *regexp.Regexp) string {
	ret := ""
	match := re.FindSubmatch(content)
	//fmt.Println(match)
	if len(match) == 2 {
		ret = string(match[1])
	} else {
		ret = "0"
	}
	return ret
}
