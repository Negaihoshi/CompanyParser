package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"syscall"
	"time"
	//"strings"
)

//Parse District
func LoctionParser(location string, readline string, c chan int) string {

	if location == "臺北市" {
		taipeiCityFile, error := os.OpenFile("data/Taipei_City.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			taipeiCityFile, _ = os.Create("data/Taipei_City.json")
		}
		writeTaipeiCity := bufio.NewWriter(taipeiCityFile)
		fmt.Println(location)
		writeTaipeiCity.WriteString(readline + "\n")
		writeTaipeiCity.Flush()
		taipeiCityFile.Close()
	} else if location == "基隆市" {
		fmt.Println(location)
		keelungCityFile, error := os.OpenFile("data/Keelung_City.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			keelungCityFile, _ = os.Create("data/Keelung_City.json")
		}
		writeKeelungCity := bufio.NewWriter(keelungCityFile)
		writeKeelungCity.WriteString(readline + "\n")
		writeKeelungCity.Flush()
		keelungCityFile.Close()
	} else if location == "新北市" {
		newTaipeiCityFile, error := os.OpenFile("data/New_Taipei_City.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			newTaipeiCityFile, _ = os.Create("data/New_Taipei_City.json")
		}
		writeNewTaipeiCity := bufio.NewWriter(newTaipeiCityFile)
		fmt.Println(location)
		writeNewTaipeiCity.WriteString(readline + "\n")
		writeNewTaipeiCity.Flush()
		newTaipeiCityFile.Close()
	} else if location == "連江縣" {
		lienchiangCountyFile, error := os.OpenFile("data/Lienchiang_County.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			lienchiangCountyFile, _ = os.Create("data/Lienchiang_County.json")
		}
		writeLienchiangCounty := bufio.NewWriter(lienchiangCountyFile)
		fmt.Println(location)
		writeLienchiangCounty.WriteString(readline + "\n")
		writeLienchiangCounty.Flush()
		lienchiangCountyFile.Close()
	} else if location == "宜蘭縣" {
		yilanCountyFile, error := os.OpenFile("data/Yilan_County.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			yilanCountyFile, _ = os.Create("data/Yilan_County.json")
		}
		writeYilanCounty := bufio.NewWriter(yilanCountyFile)
		fmt.Println(location)
		writeYilanCounty.WriteString(readline + "\n")
		writeYilanCounty.Flush()
		yilanCountyFile.Close()
	} else if location == "釣魚台" {
		diaoyutaiFile, error := os.OpenFile("data/Diaoyutai.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			diaoyutaiFile, _ = os.Create("data/Diaoyutai.json")
		}
		writeDiaoyutai := bufio.NewWriter(diaoyutaiFile)
		fmt.Println(location)
		writeDiaoyutai.WriteString(readline + "\n")
		writeDiaoyutai.Flush()
		diaoyutaiFile.Close()
	} else if location == "新竹市" {
		hsinchuCityFile, error := os.OpenFile("data/Hsinchu_City.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			hsinchuCityFile, _ = os.Create("data/Hsinchu_City.json")
		}
		writeHsinchuCity := bufio.NewWriter(hsinchuCityFile)
		fmt.Println(location)
		writeHsinchuCity.WriteString(readline + "\n")
		writeHsinchuCity.Flush()
		hsinchuCityFile.Close()
	} else if location == "新竹縣" {
		hsinchuCountyFile, error := os.OpenFile("data/Hsinchu_County.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			hsinchuCountyFile, _ = os.Create("data/Hsinchu_County.json")
		}
		writeHsinchuCounty := bufio.NewWriter(hsinchuCountyFile)
		fmt.Println(location)
		writeHsinchuCounty.WriteString(readline + "\n")
		writeHsinchuCounty.Flush()
		hsinchuCountyFile.Close()
	} else if location == "桃園縣" {
		taoyuanCountyFile, error := os.OpenFile("data/Taoyuan_County.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			taoyuanCountyFile, _ = os.Create("data/Taoyuan_County.json")
		}
		writeTaoyuanCounty := bufio.NewWriter(taoyuanCountyFile)
		fmt.Println(location)
		writeTaoyuanCounty.WriteString(readline + "\n")
		writeTaoyuanCounty.Flush()
		taoyuanCountyFile.Close()
	} else if location == "苗栗縣" {
		miaoliCountyFile, error := os.OpenFile("data/Miaoli_County.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			miaoliCountyFile, _ = os.Create("data/Miaoli_County.json")
		}
		writeMiaoliCounty := bufio.NewWriter(miaoliCountyFile)
		fmt.Println(location)
		writeMiaoliCounty.WriteString(readline + "\n")
		writeMiaoliCounty.Flush()
		miaoliCountyFile.Close()
	} else if location == "臺中市" {
		taichungCityFile, error := os.OpenFile("data/Taichung_City.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			taichungCityFile, _ = os.Create("data/Taichung_City.json")
		}
		writeTaichungCity := bufio.NewWriter(taichungCityFile)
		fmt.Println(location)
		writeTaichungCity.WriteString(readline + "\n")
		writeTaichungCity.Flush()
		taichungCityFile.Close()
	} else if location == "彰化縣" {
		changhuaCountyFile, error := os.OpenFile("data/Changhua_County.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			changhuaCountyFile, _ = os.Create("data/Changhua_County.json")
		}
		writeChanghuaCounty := bufio.NewWriter(changhuaCountyFile)
		fmt.Println(location)
		writeChanghuaCounty.WriteString(readline + "\n")
		writeChanghuaCounty.Flush()
		changhuaCountyFile.Close()
	} else if location == "南投縣" {
		nantouCountyFile, error := os.OpenFile("data/Nantou_County.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			nantouCountyFile, _ = os.Create("data/Nantou_County.json")
		}
		writeNantouCounty := bufio.NewWriter(nantouCountyFile)
		fmt.Println(location)
		writeNantouCounty.WriteString(readline + "\n")
		writeNantouCounty.Flush()
		nantouCountyFile.Close()
	} else if location == "嘉義市" {
		chiayiCityFile, error := os.OpenFile("data/Chiayi_City.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			chiayiCityFile, _ = os.Create("data/Chiayi_City.json")
		}
		writeChiayiCity := bufio.NewWriter(chiayiCityFile)
		fmt.Println(location)
		writeChiayiCity.WriteString(readline + "\n")
		writeChiayiCity.Flush()
		chiayiCityFile.Close()
	} else if location == "嘉義縣" {
		chiayiCountyFile, error := os.OpenFile("data/Chiayi_County.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			chiayiCountyFile, _ = os.Create("data/Chiayi_County.json")
		}
		writeChiayiCounty := bufio.NewWriter(chiayiCountyFile)
		fmt.Println(location)
		writeChiayiCounty.WriteString(readline + "\n")
		writeChiayiCounty.Flush()
		chiayiCountyFile.Close()
	} else if location == "雲林縣" {
		yunlinCountyFile, error := os.OpenFile("data/Yunlin_County.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			yunlinCountyFile, _ = os.Create("data/Yunlin_County.json")
		}
		writeYunlinCounty := bufio.NewWriter(yunlinCountyFile)
		fmt.Println(location)
		writeYunlinCounty.WriteString(readline + "\n")
		writeYunlinCounty.Flush()
		yunlinCountyFile.Close()
	} else if location == "臺南市" {
		tainanCityFile, error := os.OpenFile("data/Tainan_City.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			tainanCityFile, _ = os.Create("data/Tainan_City.json")
		}
		writeTainanCity := bufio.NewWriter(tainanCityFile)
		fmt.Println(location)
		writeTainanCity.WriteString(readline + "\n")
		writeTainanCity.Flush()
		tainanCityFile.Close()
	} else if location == "高雄市" {
		kaohsiungCityFile, error := os.OpenFile("data/Keelung_City.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			kaohsiungCityFile, _ = os.Create("data/Kaohsiung_City.json")
		}
		writeKaohsiungCity := bufio.NewWriter(kaohsiungCityFile)
		fmt.Println(location)
		writeKaohsiungCity.WriteString(readline + "\n")
		writeKaohsiungCity.Flush()
		kaohsiungCityFile.Close()
	} else if location == "澎湖縣" {
		penghuCountyFile, error := os.OpenFile("data/Keelung_City.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			penghuCountyFile, _ = os.Create("data/Penghu_County.json")
		}
		writePenghuCounty := bufio.NewWriter(penghuCountyFile)
		fmt.Println(location)
		writePenghuCounty.WriteString(readline + "\n")
		writePenghuCounty.Flush()
		penghuCountyFile.Close()
	} else if location == "金門縣" {
		kinmenCountyFile, error := os.OpenFile("data/Keelung_City.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			kinmenCountyFile, _ = os.Create("data/Kinmen_County.json")
		}
		writeKinmenCounty := bufio.NewWriter(kinmenCountyFile)
		fmt.Println(location)
		writeKinmenCounty.WriteString(readline + "\n")
		writeKinmenCounty.Flush()
		kinmenCountyFile.Close()
	} else if location == "屏東縣" {
		pingtungCountyFile, error := os.OpenFile("data/Keelung_City.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			pingtungCountyFile, _ = os.Create("data/Pingtung_County.json")
		}
		writePingtungCounty := bufio.NewWriter(pingtungCountyFile)
		fmt.Println(location)
		writePingtungCounty.WriteString(readline + "\n")
		writePingtungCounty.Flush()
		pingtungCountyFile.Close()
	} else if location == "臺東縣" {
		taitungCountyFile, error := os.OpenFile("data/Taitung_County.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			taitungCountyFile, _ = os.Create("data/Taitung_County.json")
		}
		writeTaitungCounty := bufio.NewWriter(taitungCountyFile)
		fmt.Println(location)
		writeTaitungCounty.WriteString(readline + "\n")
		writeTaitungCounty.Flush()
		taitungCountyFile.Close()
	} else if location == "花蓮縣" {
		hualienCountyFile, error := os.OpenFile("data/Hualien_County.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			hualienCountyFile, _ = os.Create("data/Hualien_County.json")
		}
		writeHualienCounty := bufio.NewWriter(hualienCountyFile)
		fmt.Println(location)
		writeHualienCounty.WriteString(readline + "\n")
		writeHualienCounty.Flush()
		hualienCountyFile.Close()
	} else {
		fmt.Println("-----------------------------")
		fmt.Println("AllElse: ", location)
		elseFile, error := os.OpenFile("data/Else.json", syscall.O_RDWR|syscall.O_APPEND, 0660)
		if error != nil {
			elseFile, _ = os.Create("data/Else.json")
		}
		writeElse := bufio.NewWriter(elseFile)
		writeElse.WriteString(readline + "\n")
		writeElse.Flush()
		elseFile.Close()
	}

	return location
}

func main() {

	//開始時間
	start := time.Now()

	fileTitle := "0000000.json"
	for i := 0; i < 3; i++ {
		fileString := strconv.Itoa(i) + fileTitle

		file, err := os.Open(fileString) // For read access.
		if err != nil {
			log.Fatal(err)
		}

		CoreNumber := runtime.NumCPU()
		fmt.Printf("Use Core Number: %v\n", CoreNumber)
		runtime.GOMAXPROCS(CoreNumber)
		UserCore := make(chan int, CoreNumber)

		//br := bufio.NewReader(file)
		br := bufio.NewScanner(file)

		count := 0
		for br.Scan() {

			line := br.Text()

			//初始化 Regexp
			match, _ := regexp.MatchString("分公司", line)
			locationreg, regexpErr := regexp.Compile(`(^\d+).+地":"([^\s]+)","登`)
			if regexpErr != nil {
				log.Fatal(regexpErr)
			}
			zonereg, zoneregRegexpErr := regexp.Compile(`(^\S\S\S).+`)
			if zoneregRegexpErr != nil {
				log.Fatal(zoneregRegexpErr)
			}

			count++
			fmt.Println(count)
			result := locationreg.FindStringSubmatch(line)
			if match == false {
				locationreg, regexpErr = regexp.Compile(`(^\d+).+地":"([^\s]+)","登`)
			} else {
				locationreg, regexpErr = regexp.Compile(`(^\d+).+地":"([^\s]+)","分`)
			}
			result = locationreg.FindStringSubmatch(line)

			for num, v := range result {
				if num/2 == 1 {

					result2 := zonereg.FindStringSubmatch(v)
					for num2, v2 := range result2 {
						if num2/1 == 1 {
							//fmt.Printf("%s\n", v2)
							for i := 0; i < CoreNumber; i++ {
								go LoctionParser(v2, line, UserCore)
							}

						}
					}
					//fmt.Printf("%s\n", v)
				}
			}
			//fmt.Printf("String = ", result, "\n")
			//fmt.Printf("%v", result)
			//fmt.Printf("\n")
		}
		if err := br.Err(); err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	end := time.Now()

	//花費時間
	fmt.Printf("Spend Time: %vs\n", end.Sub(start).Seconds())
}
