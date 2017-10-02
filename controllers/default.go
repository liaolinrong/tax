package controllers

import (
	"fmt"

	"strconv"

	"github.com/astaxie/beego"
)

type TaxParamAll struct {
	//个税起征点
	Tax_base float64

	//分级征税点
	Tax_grade1 float64
	Tax_grade2 float64
	Tax_grade3 float64
	Tax_grade4 float64
	Tax_grade5 float64
	Tax_grade6 float64
	Tax_grade7 float64

	//分级税率
	Tax_rate1 float64
	Tax_rate2 float64
	Tax_rate3 float64
	Tax_rate4 float64
	Tax_rate5 float64
	Tax_rate6 float64
	Tax_rate7 float64

	//分级速算扣除数
	Tax_sub_num1 float64
	Tax_sub_num2 float64
	Tax_sub_num3 float64
	Tax_sub_num4 float64
	Tax_sub_num5 float64
	Tax_sub_num6 float64
	Tax_sub_num7 float64
}

type TaxParamLocation struct {
	//上限
	Tax_yanglao_max   float64
	Tax_yiliao_max    float64
	Tax_shiye_max     float64
	Tax_gongshang_max float64
	Tax_shengyu_max   float64
	Tax_gjj_max       float64

	//下限
	Tax_yanglao_min   float64
	Tax_yiliao_min    float64
	Tax_shiye_min     float64
	Tax_gongshang_min float64
	Tax_shengyu_min   float64
	Tax_gjj_min       float64

	//缴纳比率
	Tax_yanglao_rate   float64
	Tax_yiliao_rate    float64
	Tax_shiye_rate     float64
	Tax_gongshang_rate float64
	Tax_shengyu_rate   float64
	Tax_gjj_rate       float64

	//上限 公司方面
	Tax_yanglao_max_gs   float64
	Tax_yiliao_max_gs    float64
	Tax_shiye_max_gs     float64
	Tax_gongshang_max_gs float64
	Tax_shengyu_max_gs   float64
	Tax_gjj_max_gs       float64

	//下限 公司方面
	Tax_yanglao_min_gs   float64
	Tax_yiliao_min_gs    float64
	Tax_shiye_min_gs     float64
	Tax_gongshang_min_gs float64
	Tax_shengyu_min_gs   float64
	Tax_gjj_min_gs       float64

	//缴纳比率 公司方面
	Tax_yanglao_rate_gs   float64
	Tax_yiliao_rate_gs    float64
	Tax_shiye_rate_gs     float64
	Tax_gongshang_rate_gs float64
	Tax_shengyu_rate_gs   float64
	Tax_gjj_rate_gs       float64
}

var tpa TaxParamAll
var LocalTaxParamMap map[string]TaxParamLocation

func init() {
	tpa = TaxParamAll{
		//起征点
		Tax_base: 3500,

		//分级征税点
		Tax_grade1: 0.0,
		Tax_grade2: 1500.0,
		Tax_grade3: 4500.0,
		Tax_grade4: 9000.0,
		Tax_grade5: 35000.0,
		Tax_grade6: 55000.0,
		Tax_grade7: 80000.0,

		//分级税率
		Tax_rate1: 0.03,
		Tax_rate2: 0.1,
		Tax_rate3: 0.2,
		Tax_rate4: 0.25,
		Tax_rate5: 0.3,
		Tax_rate6: 0.35,
		Tax_rate7: 0.45,

		//分级速算扣除数
		Tax_sub_num1: 0,
		Tax_sub_num2: 105,
		Tax_sub_num3: 555,
		Tax_sub_num4: 1005,
		Tax_sub_num5: 2755,
		Tax_sub_num6: 5505,
		Tax_sub_num7: 13505,
	}

	LocalTaxParamMap = map[string]TaxParamLocation{
		"hangzhou": TaxParamLocation{
			Tax_yanglao_max: 967.0, Tax_yiliao_max: 241.0, Tax_shiye_max: 60.0,
			Tax_gongshang_max: 0.0, Tax_shengyu_max: 0.0, Tax_gjj_max: 2076.0,

			Tax_yanglao_rate: 0.08, Tax_yiliao_rate: 0.02, Tax_shiye_rate: 0.005,
			Tax_gongshang_rate: 0.0, Tax_shengyu_rate: 0.0, Tax_gjj_rate: 0.12,

			Tax_yanglao_max_gs: 1693.0, Tax_yiliao_max_gs: 1390.0, Tax_shiye_max_gs: 181.0,
			Tax_gongshang_max_gs: 14.0, Tax_shengyu_max_gs: 145.0, Tax_gjj_max_gs: 2076.0,

			Tax_yanglao_rate_gs: 0.14, Tax_yiliao_rate_gs: 0.115, Tax_shiye_rate_gs: 0.015,
			Tax_gongshang_rate_gs: 0.0012, Tax_shengyu_rate_gs: 0.012, Tax_gjj_rate_gs: 0.12,

			Tax_yanglao_min: 193.0, Tax_yiliao_min: 48.0, Tax_shiye_min: 12.0,
			Tax_gongshang_min: 0.0, Tax_shengyu_min: 0.0, Tax_gjj_min: 176.0,

			Tax_yanglao_min_gs: 338.0, Tax_yiliao_min_gs: 278.0, Tax_shiye_min_gs: 36.0,
			Tax_gongshang_min_gs: 2.0, Tax_shengyu_min_gs: 29.0, Tax_gjj_min_gs: 176.0},

		"beijing": TaxParamLocation{
			Tax_yanglao_max: 1551.0, Tax_yiliao_max: 387.0, Tax_shiye_max: 38.0,
			Tax_gongshang_max: 0.0, Tax_shengyu_max: 0.0, Tax_gjj_max: 2327.0,

			Tax_yanglao_rate: 0.08, Tax_yiliao_rate: 0.02, Tax_shiye_rate: 0.002,
			Tax_gongshang_rate: 0.0, Tax_shengyu_rate: 0.0, Tax_gjj_rate: 0.12,

			Tax_yanglao_max_gs: 3877.0, Tax_yiliao_max_gs: 1938.0, Tax_shiye_max_gs: 193.0,
			Tax_gongshang_max_gs: 96.0, Tax_shengyu_max_gs: 155.0, Tax_gjj_max_gs: 2327.0,

			Tax_yanglao_rate_gs: 0.20, Tax_yiliao_rate_gs: 0.10, Tax_shiye_rate_gs: 0.01,
			Tax_gongshang_rate_gs: 0.005, Tax_shengyu_rate_gs: 0.008, Tax_gjj_rate_gs: 0.12,

			Tax_yanglao_min: 206.0, Tax_yiliao_min: 51.0, Tax_shiye_min: 5.0,
			Tax_gongshang_min: 0.0, Tax_shengyu_min: 0.0, Tax_gjj_min: 206.0,

			Tax_yanglao_min_gs: 517.0, Tax_yiliao_min_gs: 258.0, Tax_shiye_min_gs: 25.0,
			Tax_gongshang_min_gs: 12.0, Tax_shengyu_min_gs: 20.0, Tax_gjj_min_gs: 206.0},

		"shanghai": TaxParamLocation{
			Tax_yanglao_max: 1308.0, Tax_yiliao_max: 327.0, Tax_shiye_max: 81.0,
			Tax_gongshang_max: 0.0, Tax_shengyu_max: 0.0, Tax_gjj_max: 1145.0,

			Tax_yanglao_rate: 0.08, Tax_yiliao_rate: 0.02, Tax_shiye_rate: 0.005,
			Tax_gongshang_rate: 0.0, Tax_shengyu_rate: 0.0, Tax_gjj_rate: 0.07,

			Tax_yanglao_max_gs: 3434.0, Tax_yiliao_max_gs: 1798.0, Tax_shiye_max_gs: 245.0,
			Tax_gongshang_max_gs: 81.0, Tax_shengyu_max_gs: 163.0, Tax_gjj_max_gs: 1145.0,

			Tax_yanglao_rate_gs: 0.21, Tax_yiliao_rate_gs: 0.11, Tax_shiye_rate_gs: 0.015,
			Tax_gongshang_rate_gs: 0.005, Tax_shengyu_rate_gs: 0.01, Tax_gjj_rate_gs: 0.07,

			Tax_yanglao_min: 261.0, Tax_yiliao_min: 65.0, Tax_shiye_min: 16.0,
			Tax_gongshang_min: 0.0, Tax_shengyu_min: 0.0, Tax_gjj_min: 127.0,

			Tax_yanglao_min_gs: 686.0, Tax_yiliao_min_gs: 359.0, Tax_shiye_min_gs: 49.0,
			Tax_gongshang_min_gs: 16.0, Tax_shengyu_min_gs: 32.0, Tax_gjj_min_gs: 127.0},

		"guangzhou": TaxParamLocation{
			Tax_yanglao_max: 1484.0, Tax_yiliao_max: 371.0, Tax_shiye_max: 92.0,
			Tax_gongshang_max: 0.0, Tax_shengyu_max: 0.0, Tax_gjj_max: 6187.0,

			Tax_yanglao_rate: 0.08, Tax_yiliao_rate: 0.02, Tax_shiye_rate: 0.005,
			Tax_gongshang_rate: 0.0, Tax_shengyu_rate: 0.0, Tax_gjj_rate: 0.2,

			Tax_yanglao_max_gs: 2598.0, Tax_yiliao_max_gs: 1484.0, Tax_shiye_max_gs: 222.0,
			Tax_gongshang_max_gs: 92.0, Tax_shengyu_max_gs: 157.0, Tax_gjj_max_gs: 6187.0,

			Tax_yanglao_rate_gs: 0.14, Tax_yiliao_rate_gs: 0.08, Tax_shiye_rate_gs: 0.012,
			Tax_gongshang_rate_gs: 0.005, Tax_shengyu_rate_gs: 0.0085, Tax_gjj_rate_gs: 0.2,

			Tax_yanglao_min: 296.0, Tax_yiliao_min: 74.0, Tax_shiye_min: 18.0,
			Tax_gongshang_min: 0.0, Tax_shengyu_min: 0.0, Tax_gjj_min: 94.0,

			Tax_yanglao_min_gs: 519.0, Tax_yiliao_min_gs: 296.0, Tax_shiye_min_gs: 44.0,
			Tax_gongshang_min_gs: 18.0, Tax_shengyu_min_gs: 31.0, Tax_gjj_min_gs: 94.0},

		"shenzhen": TaxParamLocation{
			Tax_yanglao_max: 1452.0, Tax_yiliao_max: 363.0, Tax_shiye_max: 181.0,
			Tax_gongshang_max: 0.0, Tax_shengyu_max: 0.0, Tax_gjj_max: 6054.0,

			Tax_yanglao_rate: 0.08, Tax_yiliao_rate: 0.02, Tax_shiye_rate: 0.01,
			Tax_gongshang_rate: 0.0, Tax_shengyu_rate: 0.0, Tax_gjj_rate: 0.2,

			Tax_yanglao_max_gs: 2542.0, Tax_yiliao_max_gs: 1126.0, Tax_shiye_max_gs: 363.0,
			Tax_gongshang_max_gs: 72.0, Tax_shengyu_max_gs: 90.0, Tax_gjj_max_gs: 6054.0,

			Tax_yanglao_rate_gs: 0.14, Tax_yiliao_rate_gs: 0.062, Tax_shiye_rate_gs: 0.02,
			Tax_gongshang_rate_gs: 0.004, Tax_shengyu_rate_gs: 0.005, Tax_gjj_rate_gs: 0.2,

			Tax_yanglao_min: 162.0, Tax_yiliao_min: 40.0, Tax_shiye_min: 20.0,
			Tax_gongshang_min: 0.0, Tax_shengyu_min: 0.0, Tax_gjj_min: 90.0,

			Tax_yanglao_min_gs: 284.0, Tax_yiliao_min_gs: 125.0, Tax_shiye_min_gs: 40.0,
			Tax_gongshang_min_gs: 8.0, Tax_shengyu_min_gs: 10.0, Tax_gjj_min_gs: 90.0},

		"quzhou": TaxParamLocation{
			Tax_yanglao_max: 967.0, Tax_yiliao_max: 120.0, Tax_shiye_max: 120.0,
			Tax_gongshang_max: 0.0, Tax_shengyu_max: 0.0, Tax_gjj_max: 1451.0,

			Tax_yanglao_rate: 0.08, Tax_yiliao_rate: 0.01, Tax_shiye_rate: 0.01,
			Tax_gongshang_rate: 0.0, Tax_shengyu_rate: 0.0, Tax_gjj_rate: 0.12,

			Tax_yanglao_max_gs: 1693.0, Tax_yiliao_max_gs: 604.0, Tax_shiye_max_gs: 241.0,
			Tax_gongshang_max_gs: 72.0, Tax_shengyu_max_gs: 96.0, Tax_gjj_max_gs: 1451.0,

			Tax_yanglao_rate_gs: 0.14, Tax_yiliao_rate_gs: 0.05, Tax_shiye_rate_gs: 0.02,
			Tax_gongshang_rate_gs: 0.006, Tax_shengyu_rate_gs: 0.008, Tax_gjj_rate_gs: 0.12,

			Tax_yanglao_min: 193.0, Tax_yiliao_min: 24.0, Tax_shiye_min: 24.0,
			Tax_gongshang_min: 0.0, Tax_shengyu_min: 0.0, Tax_gjj_min: 67.0,

			Tax_yanglao_min_gs: 338.0, Tax_yiliao_min_gs: 120.0, Tax_shiye_min_gs: 48.0,
			Tax_gongshang_min_gs: 14.0, Tax_shengyu_min_gs: 19.0, Tax_gjj_min_gs: 67.0},
	}
	fmt.Println("init data, hangzhou yanglao max: ", LocalTaxParamMap["hangzhou"].Tax_yanglao_max)
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
	c.TplName = "wel.html"
}

func (this *MainController) After() {
	// Get form value.
	bfmoney := this.GetString("bfmoney")
	location := this.GetString("location")

	fmt.Println("beforemoney:", bfmoney)
	fmt.Println("location:", location)

	bfm, err := strconv.Atoi(bfmoney)
	if err != nil {
		//errString := fmt.Sprintf("出错了: %.2f\n", err)
		errString := fmt.Sprintf("出错了: 无法解析您的输入(%s)，请输入整数\n", bfmoney)
		this.Ctx.WriteString(errString)
		return
	}

	f_bfm := float64(bfm)

	fmt.Println("tpa.Tax_rate1: ", tpa.Tax_rate1)
	fmt.Println("LocalTaxParamMap[location].Tax_yanglao_rate: ", LocalTaxParamMap[location].Tax_yanglao_rate)

	yanglao := f_bfm * LocalTaxParamMap[location].Tax_yanglao_rate
	if yanglao > LocalTaxParamMap[location].Tax_yanglao_max {
		yanglao = LocalTaxParamMap[location].Tax_yanglao_max
	} else if yanglao < LocalTaxParamMap[location].Tax_yanglao_min {
		yanglao = LocalTaxParamMap[location].Tax_yanglao_min
	}

	yanglao_gs := f_bfm * LocalTaxParamMap[location].Tax_yanglao_rate_gs
	if yanglao_gs > LocalTaxParamMap[location].Tax_yanglao_max_gs {
		yanglao_gs = LocalTaxParamMap[location].Tax_yanglao_max_gs
	} else if yanglao_gs < LocalTaxParamMap[location].Tax_yanglao_min_gs {
		yanglao_gs = LocalTaxParamMap[location].Tax_yanglao_min_gs
	}

	yiliao := f_bfm * LocalTaxParamMap[location].Tax_yiliao_rate
	if yiliao > LocalTaxParamMap[location].Tax_yiliao_max {
		yiliao = LocalTaxParamMap[location].Tax_yiliao_max
	} else if yiliao < LocalTaxParamMap[location].Tax_yiliao_min {
		yiliao = LocalTaxParamMap[location].Tax_yiliao_min
	}

	yiliao_gs := f_bfm * LocalTaxParamMap[location].Tax_yiliao_rate_gs
	if yiliao_gs > LocalTaxParamMap[location].Tax_yiliao_max_gs {
		yiliao_gs = LocalTaxParamMap[location].Tax_yiliao_max_gs
	} else if yiliao_gs < LocalTaxParamMap[location].Tax_yiliao_min_gs {
		yiliao_gs = LocalTaxParamMap[location].Tax_yiliao_min_gs
	}

	shiye := f_bfm * LocalTaxParamMap[location].Tax_shiye_rate
	if shiye > LocalTaxParamMap[location].Tax_shiye_max {
		shiye = LocalTaxParamMap[location].Tax_shiye_max
	} else if shiye < LocalTaxParamMap[location].Tax_shiye_min {
		shiye = LocalTaxParamMap[location].Tax_shiye_min
	}

	shiye_gs := f_bfm * LocalTaxParamMap[location].Tax_shiye_rate_gs
	if shiye_gs > LocalTaxParamMap[location].Tax_shiye_max_gs {
		shiye_gs = LocalTaxParamMap[location].Tax_shiye_max_gs
	} else if shiye_gs < LocalTaxParamMap[location].Tax_shiye_min_gs {
		shiye_gs = LocalTaxParamMap[location].Tax_shiye_min_gs
	}

	gongshang := f_bfm * LocalTaxParamMap[location].Tax_gongshang_rate
	if gongshang > LocalTaxParamMap[location].Tax_gongshang_max {
		gongshang = LocalTaxParamMap[location].Tax_gongshang_max
	} else if gongshang < LocalTaxParamMap[location].Tax_gongshang_min {
		gongshang = LocalTaxParamMap[location].Tax_gongshang_min
	}

	gongshang_gs := f_bfm * LocalTaxParamMap[location].Tax_gongshang_rate_gs
	if gongshang_gs > LocalTaxParamMap[location].Tax_gongshang_max_gs {
		gongshang_gs = LocalTaxParamMap[location].Tax_gongshang_max_gs
	} else if gongshang_gs < LocalTaxParamMap[location].Tax_gongshang_min_gs {
		gongshang_gs = LocalTaxParamMap[location].Tax_gongshang_min_gs
	}

	shengyu := f_bfm * LocalTaxParamMap[location].Tax_shengyu_rate
	if shengyu > LocalTaxParamMap[location].Tax_shengyu_max {
		shengyu = LocalTaxParamMap[location].Tax_shengyu_max
	} else if shengyu < LocalTaxParamMap[location].Tax_shengyu_min {
		shengyu = LocalTaxParamMap[location].Tax_shengyu_min
	}

	shengyu_gs := f_bfm * LocalTaxParamMap[location].Tax_shengyu_rate_gs
	if shengyu_gs > LocalTaxParamMap[location].Tax_shengyu_max_gs {
		shengyu_gs = LocalTaxParamMap[location].Tax_shengyu_max_gs
	} else if shengyu_gs < LocalTaxParamMap[location].Tax_shengyu_min_gs {
		shengyu_gs = LocalTaxParamMap[location].Tax_shengyu_min_gs
	}

	gjj := f_bfm * LocalTaxParamMap[location].Tax_gjj_rate
	if gjj > LocalTaxParamMap[location].Tax_gjj_max {
		gjj = LocalTaxParamMap[location].Tax_gjj_max
	} else if gjj < LocalTaxParamMap[location].Tax_gjj_min {
		gjj = LocalTaxParamMap[location].Tax_gjj_min
	}

	gjj_gs := f_bfm * LocalTaxParamMap[location].Tax_gjj_rate_gs
	if gjj_gs > LocalTaxParamMap[location].Tax_gjj_max_gs {
		gjj_gs = LocalTaxParamMap[location].Tax_gjj_max_gs
	} else if gjj_gs < LocalTaxParamMap[location].Tax_gjj_min_gs {
		gjj_gs = LocalTaxParamMap[location].Tax_gjj_min_gs
	}

	sub51 := f_bfm - yanglao - yiliao - shiye - gongshang - shengyu - gjj

	//fmt.Println("sub51:", sub51)
	chao := sub51 - tpa.Tax_base

	//fmt.Println("chao:", chao)
	tax := 0.0
	if chao > tpa.Tax_grade1 && chao <= tpa.Tax_grade2 {
		tax = chao*tpa.Tax_rate1 - tpa.Tax_sub_num1
	} else if chao > tpa.Tax_grade2 && chao <= tpa.Tax_grade3 {
		tax = chao*tpa.Tax_rate2 - tpa.Tax_sub_num2
	} else if chao > tpa.Tax_grade3 && chao <= tpa.Tax_grade4 {
		tax = chao*tpa.Tax_rate3 - tpa.Tax_sub_num3
	} else if chao > tpa.Tax_grade4 && chao <= tpa.Tax_grade5 {
		tax = chao*tpa.Tax_rate4 - tpa.Tax_sub_num4
	} else if chao > tpa.Tax_grade5 && chao <= tpa.Tax_grade6 {
		tax = chao*tpa.Tax_rate5 - tpa.Tax_sub_num5
	} else if chao > tpa.Tax_grade6 && chao <= tpa.Tax_grade7 {
		tax = chao*tpa.Tax_rate6 - tpa.Tax_sub_num6
	} else if chao > tpa.Tax_grade7 {
		tax = chao*tpa.Tax_rate7 - tpa.Tax_sub_num7
	}

	fmt.Println("gongshang:", gongshang, "shengyu", shengyu)
	baoxiansum := yanglao + yiliao + shiye + gongshang + shengyu
	baoxiansum_gs := yanglao_gs + yiliao_gs + shiye_gs + gongshang_gs + shengyu_gs

	alljiaona := baoxiansum + gjj + tax
	alljiaona_gs := baoxiansum_gs + gjj_gs

	finalmoney := sub51 - tax
	// result := fmt.Sprintf("到手: %.2f\n\n个人所得税: %.2f\n\n养老保险个人(%.2f%%): %.2f\t养老保险公司(%.2f%%): %.2f\n医疗保险个人(%.2f%%): %.2f\t医疗保险公司(%.2f%%): %.2f\n失业保险个人(%.2f%%): %.2f\t失业保险公司(%.2f%%): %.2f\n工伤保险个人(%.2f%%): %.2f\t工伤保险公司(%.2f%%): %.2f\n生育保险个人(%.2f%%): %.2f\t生育保险公司(%.2f%%): %.2f\n公积金个人(%.2f%%): %.2f\t公积金公司(%.2f%%): %.2f\n\n个人缴纳总和: %.2f\t公司缴纳总和: %.2f",
	// 	finalmoney, tax,
	// 	LocalTaxParamMap[location].Tax_yanglao_rate*100, yanglao, LocalTaxParamMap[location].Tax_yanglao_rate_gs*100, yanglao_gs,
	// 	LocalTaxParamMap[location].Tax_yiliao_rate*100, yiliao, LocalTaxParamMap[location].Tax_yiliao_rate_gs*100, yiliao_gs,
	// 	LocalTaxParamMap[location].Tax_shiye_rate*100, shiye, LocalTaxParamMap[location].Tax_shiye_rate_gs*100, shiye_gs,
	// 	LocalTaxParamMap[location].Tax_gongshang_rate*100, gongshang, LocalTaxParamMap[location].Tax_gongshang_rate_gs*100, gongshang_gs,
	// 	LocalTaxParamMap[location].Tax_shengyu_rate*100, shengyu, LocalTaxParamMap[location].Tax_shengyu_rate_gs*100, shengyu_gs,
	// 	LocalTaxParamMap[location].Tax_gjj_rate*100, gjj, LocalTaxParamMap[location].Tax_gjj_rate_gs*100, gjj_gs,
	// 	alljiaona, alljiaona_gs)

	// this.Ctx.WriteString(result)

	this.Data["finalmoney"] = fmt.Sprintf("%.2f", finalmoney)
	this.Data["tax"] = fmt.Sprintf("%.2f", tax)

	this.Data["yanglao"] = fmt.Sprintf("%.2f", yanglao)
	this.Data["yiliao"] = fmt.Sprintf("%.2f", yiliao)
	this.Data["shiye"] = fmt.Sprintf("%.2f", shiye)
	this.Data["gongshang"] = fmt.Sprintf("%.2f", gongshang)
	this.Data["shengyu"] = fmt.Sprintf("%.2f", shengyu)
	this.Data["gjj"] = fmt.Sprintf("%.2f", gjj)

	this.Data["yanglao_gs"] = fmt.Sprintf("%.2f", yanglao_gs)
	this.Data["yiliao_gs"] = fmt.Sprintf("%.2f", yiliao_gs)
	this.Data["shiye_gs"] = fmt.Sprintf("%.2f", shiye_gs)
	this.Data["gongshang_gs"] = fmt.Sprintf("%.2f", gongshang_gs)
	this.Data["shengyu_gs"] = fmt.Sprintf("%.2f", shengyu_gs)
	this.Data["gjj_gs"] = fmt.Sprintf("%.2f", gjj_gs)

	this.Data["yanglao_rate"] = fmt.Sprintf("%.2f", LocalTaxParamMap[location].Tax_yanglao_rate*100)
	this.Data["yiliao_rate"] = fmt.Sprintf("%.2f", LocalTaxParamMap[location].Tax_yiliao_rate*100)
	this.Data["shiye_rate"] = fmt.Sprintf("%.2f", LocalTaxParamMap[location].Tax_shiye_rate*100)
	this.Data["gongshang_rate"] = fmt.Sprintf("%.2f", LocalTaxParamMap[location].Tax_gongshang_rate*100)
	this.Data["shengyu_rate"] = fmt.Sprintf("%.2f", LocalTaxParamMap[location].Tax_shengyu_rate*100)
	this.Data["gjj_rate"] = fmt.Sprintf("%.2f", LocalTaxParamMap[location].Tax_gjj_rate*100)

	this.Data["yanglao_rate_gs"] = fmt.Sprintf("%.2f", LocalTaxParamMap[location].Tax_yanglao_rate_gs*100)
	this.Data["yiliao_rate_gs"] = fmt.Sprintf("%.2f", LocalTaxParamMap[location].Tax_yiliao_rate_gs*100)
	this.Data["shiye_rate_gs"] = fmt.Sprintf("%.2f", LocalTaxParamMap[location].Tax_shiye_rate_gs*100)
	this.Data["gongshang_rate_gs"] = fmt.Sprintf("%.2f", LocalTaxParamMap[location].Tax_gongshang_rate_gs*100)
	this.Data["shengyu_rate_gs"] = fmt.Sprintf("%.2f", LocalTaxParamMap[location].Tax_shengyu_rate_gs*100)
	this.Data["gjj_rate_gs"] = fmt.Sprintf("%.2f", LocalTaxParamMap[location].Tax_gjj_rate_gs*100)

	this.Data["alljiaona"] = fmt.Sprintf("%.2f", alljiaona)
	this.Data["alljiaona_gs"] = fmt.Sprintf("%.2f", alljiaona_gs)

	this.TplName = "tax_after.html"
	return
}
