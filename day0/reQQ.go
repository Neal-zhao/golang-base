package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	chanImageUrls chan string
	waitGroup sync.WaitGroup
	chanTask chan string
	//reImg = `https?://[^"]+?(\.((.jpg)|(png)|(.jpeg)|(gif)|(bmp)))`
	reQQEmail = `(\d+)@qq.com`
	reLinke = `href="(https?://[\s\S]+?)"`
	rePhone = `1[3456789]\d\s?\d{4}\s?\d{4}`
	reIdcard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
	reImg = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func GetFilenameFromUrl(url string) (filename string) {
	lastIndex := strings.LastIndex(url,"/")
	filename = url[lastIndex + 1:]
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	return
}
func CheckOK()  {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成了爬取任务\n",url)
		count++
		if count == 26 {
			close(chanImageUrls)
			break
		}
	}
}
func getImgs(url string) (urls []string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr,-1)
	fmt.Printf("共找到%d条结果\n",len(results))
	for _,result := range results {
		url := result[0]
		urls = append(urls,url)
	}
	return
}
func getImgUrls(url string)  {
	urls := getImgs(url)
	for i,url := range urls {
		fmt.Println(string(i)+"->"+url)
		chanImageUrls <- url
	}
	chanTask <- url
	waitGroup.Done()
}
func DownloadFile(url string,filename string) (ok bool)  {
	resp,err := http.Get(url)
	HandleError(err,"http.Get url error")
	defer resp.Body.Close()
	bytes,err := ioutil.ReadAll(resp.Body)
	HandleError(err,"ioutil.ReadAll err")
	filename = "E:/topgoer.com/src/github.com/student/3.0/img/" + filename
	err = ioutil.WriteFile(filename,bytes,0666)
	if err != nil {
		return false
	}
	return true
}
func DownloadImg()  {
	for url := range chanImageUrls{
		filename := GetFilenameFromUrl(url)
		ok := DownloadFile(url,filename)

		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
}
func main()  {
	pageStart := flag.Int("pageStart",1,"page Start")
	pageEnd := flag.Int("pageEnd",2,"page end")
	flag.Parse()
	chanImageUrls = make(chan string, 1000000)
	chanTask = make(chan string, 26)
	fmt.Println(*pageStart,*pageEnd)
	for i:=*pageStart; i<*pageEnd;i++ {
		waitGroup.Add(1)
		//go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/"+strconv.Itoa(i)+".html")
		fmt.Println("https://zztt41.com/index/index?id=1&page="+strconv.Itoa(i))
		go getImgUrls("https://zztt41.com/index/index?id=1&page="+strconv.Itoa(i))
	}

	waitGroup.Add(1)
	go CheckOK()
	for i:=0;i<5;i++ {
		waitGroup.Add(1)
		go DownloadImg()
	}
	waitGroup.Wait()
	/*GetEmail2()
	// 2.抽取的爬邮箱
	GetEmail2("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	// 3.爬链接
	GetLink("http://www.baidu.com/s?wd=%E8%B4%B4%E5%90%A7%20%E7%95%99%E4%B8%8B%E9%82%AE%E7%AE%B1&rsv_spt=1&rsv_iqid=0x98ace53400003985&issp=1&f=8&rsv_bp=1&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_dl=ib&rsv_sug2=0&inputT=5197&rsv_sug4=6345")
	// 4.爬手机号
	GetPhone("https://www.zhaohaowang.com/")
	// 5.爬身份证号
	GetIdCard("https://henan.qq.com/a/20171107/069413.htm")
	// 6.爬图片
	GetImg("http://image.baidu.com/search/index?tn=baiduimage&ps=1&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=%E7%BE%8E%E5%A5%B3")
	*/
}
func GetIdCard(url string)  {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reIdcard)
	results := re.FindAllStringSubmatch(pageStr,-1)
	for _,result := range results{
		fmt.Println(result)
	}
}
func GetLink(url string)  {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reLinke)
	results := re.FindAllStringSubmatch(pageStr,-1)
	for _,result := range results{
		fmt.Println(result[1])
	}
}
func GetPhone(url string)  {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(rePhone)
	results := re.FindAllStringSubmatch(pageStr,-1)
	for _,result := range results{
		fmt.Println(result)
	}
}
func GetImg(url string)  {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr,-1)
	for _,result := range results{
		fmt.Println(result[0])
	}
}
func GetPageStr(url string) (pageStr string)  {
	resp,err := http.Get(url)
	HandleError(err,"http.Get url err")
	defer resp.Body.Close()
	pageBytes,err := ioutil.ReadAll(resp.Body)
	HandleError(err,"ioutil.ReadAll error")
	pageStr = string(pageBytes)
	return pageStr
}
func GetEmail2(url string)  {
	//resp,err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	resp,err := http.Get(url)
	HandleError(err," http.Get url")
	defer resp.Body.Close()
	pageBytes ,err := ioutil.ReadAll(resp.Body)
	HandleError(err," ioutil.ReadAll")
	pageStr := string(pageBytes)
	re := regexp.MustCompile(reQQEmail)
	results := re.FindAllStringSubmatch(pageStr,-1)
	for _,result := range results{
		fmt.Println("email:",result[0])
		fmt.Println("qq:",result[1])
	}
}
func HandleError(err error, why string)  {
	if err != nil {
		fmt.Println(why,err)
	}
}