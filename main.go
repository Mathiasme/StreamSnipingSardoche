package main

import (
	"encoding/csv"
	"github.com/KnutZuidema/golio"
	"log"
	"os"
  "strconv"
	"github.com/sirupsen/logrus"
)

func main() {
	//create Api Client for EUW
	client := golio.NewClient("API_KEY",
		golio.WithRegion("kr"),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))

	GameIDs := []string{
		"KR_6004453559",
		"KR_6004442770",
		"KR_6004441948",
		"KR_6004368921",
		"KR_6004366866",
		"KR_6004277089",
		"KR_6004361825",
		"KR_6004267073",
		"KR_6004251096",
		"KR_6004076637",
		"KR_6004068831",
		"KR_6004063734",
		"KR_6003937941",
		"KR_6003931452",
		"KR_6003807474",
		"KR_6003813562",
		"KR_6003678772",
		"KR_6003663769",
		"KR_6003487340",
		"KR_6003463913",
		"KR_6003258574",
		"KR_6002867290",
		"KR_6002826273",
		"KR_6002814834",
		"KR_6002840786",
		"KR_6002737749",
		"KR_6002773889",
		"KR_6002647210",
		"KR_6002581579",
		"KR_6002495608",
		"KR_6002476737",
		"KR_6002393912",
		"KR_6002345217",
		"KR_6002310948",
		"KR_6002205309",
		"KR_6002162840",
		"KR_6002170259",
		"KR_6001972255",
		"KR_6001728712",
		"KR_6001705868",
		"KR_6001740802",
		"KR_6001608834",
		"KR_6001606772",
		"KR_6001643108",
		"KR_6001579678",
		"KR_6001323449",
		"KR_6001249957",
		"KR_6001276974",
		"KR_6001233266",
		"KR_6001177621",
		"KR_6001151340",
		"KR_6001050456",
		"KR_6000973854",
		"KR_6000889796",
		"KR_6000812401",
		"KR_6000717910",
		"KR_6000714338",
		"KR_6000598960",
		"KR_6000673303",
		"KR_6000439341",
		"KR_6000057204",
		"KR_6000072720",
		"KR_5999927730",
		"KR_5999922519",
		"KR_5999868213",
		"KR_5999863817",
		"KR_5999401210",
		"KR_5999315059",
		"KR_5999228887",
		"KR_5999204653",
		"KR_5999136752",
		"KR_5999057870",
		"KR_5998942751",
		"KR_5998634756",
		"KR_5998539151",
		"KR_5998434043",
		"KR_5998357090",
		"KR_5998298774",
		"KR_5998222600",
		"KR_5998117786",
		"KR_5998130770",
		"KR_5998070157",
		"KR_5997945588",
		"KR_5997921559",
		"KR_5997857768",
		"KR_5997864439",
		"KR_5996901759",
		"KR_5996656972",
		"KR_5996702316",
		"KR_5996608661",
		"KR_5996554671",
		"KR_5995476888",
		"KR_5995452629",
		"KR_5995056239",
		"KR_5995090670",
		"KR_5994956400",
		"KR_5994619139",
		"KR_5994714646",
		"KR_5994229059",
		"KR_5994198690",
	}

	var PUUIDs []string
  SummonerName := make(map[string]string)
	for _, GameID := range GameIDs {
		MatchData, err := client.Riot.Match.Get(GameID)
		if err != nil {
			log.Println(err)
		}
		ListOfParticipants := MatchData.Info.Participants
		for _, Participant := range ListOfParticipants {
			//if Participant.PUUID != "2zD1_0crJPkEZNUvvcpAElNoy7UItjb7vPRb1F_QyslmaLo72PPDWNPfXUNTaNQ1e66Vn4XSZiw0Cg" {
				PUUIDs = append(PUUIDs, Participant.PUUID)
        SummonerName[Participant.PUUID] = Participant.SummonerName
		//	}
		}
		printCount(PUUIDs, SummonerName)
	}
}

func printCount(arr []string, SummonerNames map[string]string) {
	//Create a map of values for each element
	dict := make(map[string]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	//log.Println(dict)
	//ExportToCSV
	csvFile, err := os.Create("sardoche.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	Row := []string{
		"PUUID",
    "Summoner Name",
		"Nombre d'apparition(s) en partie",
	}
	csvwriter := csv.NewWriter(csvFile)
	err = csvwriter.Write(Row)
	if err != nil {
		log.Fatal(err)
	}
	csvwriter.Flush()

	for k, v := range dict {

		DataRow := []string{
			k,
      SummonerNames[k],
			IntToString(v),
		}

		_ = csvwriter.Write(DataRow)

	}
  csvwriter.Flush()
  csvFile.Close()
}

func IntToString(Entier int) string {
  return strconv.Itoa(Entier)
}

