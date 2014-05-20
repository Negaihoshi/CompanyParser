package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	//"strings"
)

func main() {
	file, err := os.Open("test.json") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	/*
		location := [23]string{"Taipei_City", "Keelung_City", "New_Taipei_City", "Lienchiang_County", "Yilan_County",
			"Diaoyutai", "Hsinchu_City", "Hsinchu_County", "Taoyuan_County", "Miaoli_County",
			"Taichung_City", "Changhua_County", "Nantou_County", "Chiayi_City", "Chiayi_County",
			"Yunlin_County", "Tainan_City", "Kaohsiung_City", "Penghu_County", "Kinmen_County",
			"Pingtung_County", "Taitung_County", "Hualien_County"}
	*/
	//dataStruct := map[string]string{}

	/*
		taipeiCityFile, err := os.Create("data/Taipei_City.json")
		keelungCityFile, err := os.Create("data/Keelung_City.json")
		newTaipeiCityFile, err := os.Create("data/New_Taipei_City.json")
		lienchiangCountyFile, err := os.Create("data/Lienchiang_County.json")
		yilanCountyFile, err := os.Create("data/Yilan_County.json")
		diaoyutaiFile, err := os.Create("data/Diaoyutai.json")
		hsinchuCityFile, err := os.Create("data/Hsinchu_City.json")
		hsinchuCountyFile, err := os.Create("data/Hsinchu_County.json")
		taoyuanCountyFile, err := os.Create("data/Taoyuan_County.json")
		miaoliCountyFile, err := os.Create("data/Miaoli_County.json")
		taichungCityFile, err := os.Create("data/Taichung_City.json")
		changhuaCountyFile, err := os.Create("data/Changhua_County.json")
		nantouCountyFile, err := os.Create("data/Nantou_County.json")
		chiayiCityFile, err := os.Create("data/Chiayi_City.json")
		chiayiCountyFile, err := os.Create("data/Chiayi_County.json")
		yunlinCountyFile, err := os.Create("data/Yunlin_County.json")
		tainanCityFile, err := os.Create("data/Tainan_City.json")
		kaohsiungCityFile, err := os.Create("data/Kaohsiung_City.json")
		penghuCountyFile, err := os.Create("data/Penghu_County.json")
		kinmenCountyFile, err := os.Create("data/Kinmen_County.json")
		pingtungCountyFile, err := os.Create("data/Pingtung_County.json")
		taitungCountyFile, err := os.Create("data/Taitung_County.json")
		hualienCountyFile, err := os.Create("data/Hualien_County.json")
	*/

	//br := bufio.NewReader(file)
	br := bufio.NewScanner(file)

	/*
		writeTaipeiCity := bufio.NewWriter(taipeiCityFile)
		writeKeelungCity := bufio.NewWriter(keelungCityFile)
		writeNewTaipeiCity := bufio.NewWriter(newTaipeiCityFile)
		writeLienchiangCounty := bufio.NewWriter(lienchiangCountyFile)
		writeYilanCounty := bufio.NewWriter(yilanCountyFile)
		writeDiaoyutai := bufio.NewWriter(diaoyutaiFile)
		writeHsinchuCity := bufio.NewWriter(hsinchuCityFile)
		writeHsinchuCounty := bufio.NewWriter(hsinchuCountyFile)
		writeTaoyuanCounty := bufio.NewWriter(taoyuanCountyFile)
		writeMiaoliCounty := bufio.NewWriter(miaoliCountyFile)
		writeTaichungCity := bufio.NewWriter(taichungCityFile)
		writeChanghuaCounty := bufio.NewWriter(changhuaCountyFile)
		writeNantouCounty := bufio.NewWriter(nantouCountyFile)
		writeChiayiCity := bufio.NewWriter(chiayiCityFile)
		writeChiayiCounty := bufio.NewWriter(chiayiCountyFile)
		writeYunlinCounty := bufio.NewWriter(yunlinCountyFile)
		writeTainanCity := bufio.NewWriter(tainanCityFile)
		writeKaohsiungCity := bufio.NewWriter(kaohsiungCityFile)
		writePenghuCounty := bufio.NewWriter(penghuCountyFile)
		writeKinmenCounty := bufio.NewWriter(kinmenCountyFile)
		writePingtungCounty := bufio.NewWriter(pingtungCountyFile)
		writeTaitungCounty := bufio.NewWriter(taitungCountyFile)
		writeHualienCounty := bufio.NewWriter(hualienCountyFile)
	*/

	for br.Scan() {

		if err != nil {
			break
		}

		line := br.Text()

		locationreg, regexpErr := regexp.Compile(`(^\d+).+地":"([^\s]+)","登`)
		if regexpErr != nil {
			log.Fatal(regexpErr)
		}
		zonereg, zoneregRegexpErr := regexp.Compile(`(^\S\S\S).+`)
		if zoneregRegexpErr != nil {
			log.Fatal(zoneregRegexpErr)
		}

		result := locationreg.FindStringSubmatch(line)
		for num, v := range result {
			if num/2 == 1 {
				fmt.Println(v)
				result2 := zonereg.FindStringSubmatch(v)
				for num2, v2 := range result2 {
					if num2/1 == 1 {
						fmt.Printf("%s\n", v2)
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
