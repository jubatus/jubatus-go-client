package jubatus_client

import (
	"fmt"
	"testing"

	util "./util"
	burst "github.com/jubatus/jubatus-go-client/lib/burst"
)

func GetClient(c *util.JubatusProcess) (*burst.BurstClient, error) {
	return burst.NewBurstClient("127.0.0.1:"+fmt.Sprintf("%d", c.Port), "tmp")
}

func _bootJubatus(t *testing.T, name string, config string) *util.JubatusProcess {
	jubatus, err := util.NewJubatusProcess("jubaburst", "config/burst.json")
	if err != nil {
		fmt.Print(err)
		t.Errorf("cannot boot jubatus")
		t.FailNow()
	}
	return jubatus
}

func _getClient(t *testing.T, jubatus *util.JubatusProcess) *burst.BurstClient {
	cli, err := burst.NewBurstClient(fmt.Sprintf("localhost:%d", jubatus.Port), "as")
	if err != nil {
		fmt.Print(err)
		t.Errorf("cannot connect to jubatus")
		t.FailNow()
	}
	return cli
}

func makeDocuments(documents *[]burst.Document, pos float64, burstCount int, nonburstCount int) {
        document := burst.Document {
        	Pos:    pos,
        	Text:   "Jubatus",
        }
        for i:=0; i<burstCount; i++ {
                *documents = append(*documents, document)
        }
        document = burst.Document {
        	Pos:    pos,
        	Text:   "ユバタス",
        }
        for i:=0; i<nonburstCount; i++ {
                *documents = append(*documents, document)
        }
}

func TestAddKeyword(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaburst", "config/burst.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
        keyword := burst.KeywordWithParams {
        	Keyword:        "Jubatus",
        	ScalingParam:   1.001,
        	Gamma:          0.1,
        }
	if !(cli.AddKeyword(keyword)) {
		t.Errorf("got invalid add_keyword response")
		t.FailNow()
        }
}

func TestGetAllKeywords(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaburst", "config/burst.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)

        keyword := burst.KeywordWithParams {
                Keyword: "Jubatus",
                ScalingParam: 1.003,
                Gamma: 0.3,
        }
        cli.AddKeyword(keyword)

        keyword = burst.KeywordWithParams {
                Keyword: "ユバタス",
                ScalingParam: 1.002,
                Gamma: 0.2,
        }
        cli.AddKeyword(keyword)

        keyword = burst.KeywordWithParams {
                Keyword: "ゆばたす",
                ScalingParam: 1.001,
                Gamma: 0.1,
        }
        cli.AddKeyword(keyword)

        keywords := cli.GetAllKeywords()
        if (len(keywords) != 3) {
		t.Errorf("got invalid add_all_keywords response")
		t.FailNow()
        }
}

func TestRemoveKeyword(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaburst", "config/burst.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)

        keyword := burst.KeywordWithParams {
                Keyword: "Jubatus",
                ScalingParam: 1.003,
                Gamma: 0.3,
        }
        cli.AddKeyword(keyword)
        if !(cli.RemoveKeyword("Jubatus")) {
 		t.Errorf("got invalid remove_keyword response")
		t.FailNow()
        }
}

func TestRemoveAllKeywords(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaburst", "config/burst.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)

        keyword := burst.KeywordWithParams {
                Keyword: "Jubatus",
                ScalingParam: 1.003,
                Gamma: 0.3,
        }
        cli.AddKeyword(keyword)

        keyword = burst.KeywordWithParams {
                Keyword: "ユバタス",
                ScalingParam: 1.002,
                Gamma: 0.2,
        }
        cli.AddKeyword(keyword)

        keyword = burst.KeywordWithParams {
                Keyword: "ゆばたす",
                ScalingParam: 1.001,
                Gamma: 0.1,
        }
        cli.AddKeyword(keyword)

        if !(cli.RemoveAllKeywords()) {
		t.Errorf("got invalid remove_all_keywords response")
		t.FailNow()
        }
}

func TestAddDocuments(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaburst", "config/burst.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)

        keyword := burst.KeywordWithParams {
                Keyword: "Jubatus",
                ScalingParam: 1.003,
                Gamma: 0.3,
        }
        cli.AddKeyword(keyword)

        keyword = burst.KeywordWithParams {
                Keyword: "ユバタス",
                ScalingParam: 1.002,
                Gamma: 0.2,
        }
        cli.AddKeyword(keyword)

        documents := []burst.Document{}
        makeDocuments(&documents, 1.0, 5, 30)
        makeDocuments(&documents, 11.0, 15, 50)
        makeDocuments(&documents, 21.0, 500, 10)
        makeDocuments(&documents, 31.0, 2000, 10)
        makeDocuments(&documents, 41.0, 22222, 40)
        makeDocuments(&documents, 51.0, 10, 10)
        makeDocuments(&documents, 61.0, 5, 25)

        if (cli.AddDocuments(documents) != int32(len(documents))) {
		t.Errorf("got invalid add_documents response")
		t.FailNow()
        }
}

func TestGetResult(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaburst", "config/burst.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)

        keyword := burst.KeywordWithParams {
                Keyword: "Jubatus",
                ScalingParam: 1.003,
                Gamma: 0.3,
        }
        cli.AddKeyword(keyword)

        keyword = burst.KeywordWithParams {
                Keyword: "ユバタス",
                ScalingParam: 1.002,
                Gamma: 0.2,
        }
        cli.AddKeyword(keyword)

        documents := []burst.Document{}
        makeDocuments(&documents, 1.0, 5, 30)
        makeDocuments(&documents, 11.0, 15, 50)
        makeDocuments(&documents, 21.0, 500, 10)
        makeDocuments(&documents, 31.0, 2000, 10)
        makeDocuments(&documents, 41.0, 22222, 40)
        makeDocuments(&documents, 51.0, 10, 10)
        makeDocuments(&documents, 61.0, 5, 25)

        cli.AddDocuments(documents)
        errorResult := burst.Window{}
        errorResult.StartPos = -1

        if (cli.GetResult("Jubatus").StartPos == -1) {
		t.Errorf("got invalid get_result response")
		t.FailNow()
        }
}

func TestGetResultAt(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaburst", "config/burst.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)

        keyword := burst.KeywordWithParams {
                Keyword: "Jubatus",
                ScalingParam: 1.003,
                Gamma: 0.3,
        }
        cli.AddKeyword(keyword)

        keyword = burst.KeywordWithParams {
                Keyword: "ユバタス",
                ScalingParam: 1.002,
                Gamma: 0.2,
        }
        cli.AddKeyword(keyword)

        documents := []burst.Document{}
        makeDocuments(&documents, 1.0, 5, 30)
        makeDocuments(&documents, 11.0, 15, 50)
        makeDocuments(&documents, 21.0, 500, 10)
        makeDocuments(&documents, 31.0, 2000, 10)
        makeDocuments(&documents, 41.0, 22222, 40)
        makeDocuments(&documents, 51.0, 10, 10)
        makeDocuments(&documents, 61.0, 5, 25)

        cli.AddDocuments(documents)
        errorResult := burst.Window{}
        errorResult.StartPos = -1

        if (cli.GetResultAt("Jubatus", 41.0).StartPos == -1) {
		t.Errorf("got invalid get_result_at response")
		t.FailNow()
        }
}

func TestGetAllBurstedResults(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaburst", "config/burst.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)

        keyword := burst.KeywordWithParams {
                Keyword: "Jubatus",
                ScalingParam: 1.003,
                Gamma: 0.3,
        }
        cli.AddKeyword(keyword)

        keyword = burst.KeywordWithParams {
                Keyword: "ユバタス",
                ScalingParam: 1.002,
                Gamma: 0.2,
        }
        cli.AddKeyword(keyword)

        documents := []burst.Document{}

        makeDocuments(&documents, 1.0, 5, 30)
        makeDocuments(&documents, 11.0, 15, 50)
        makeDocuments(&documents, 21.0, 500, 10)
        makeDocuments(&documents, 31.0, 2000, 10)
        makeDocuments(&documents, 41.0, 200, 40)
        makeDocuments(&documents, 51.0, 10, 10)
        makeDocuments(&documents, 61.0, 22225, 25)
        makeDocuments(&documents, 71.0, 22222, 10)
        makeDocuments(&documents, 81.0, 22222, 15)
        makeDocuments(&documents, 91.0, 222, 3)

        cli.AddDocuments(documents)

        if (len(cli.GetAllBurstedResults()) == 0) {
		t.Errorf("got invalid get_all_bursted_results response")
		t.FailNow()
        }
}

func TestGetAllBurstedResultsAt(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaburst", "config/burst.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)

        keyword := burst.KeywordWithParams {
                Keyword: "Jubatus",
                ScalingParam: 1.003,
                Gamma: 0.3,
        }
        cli.AddKeyword(keyword)

        keyword = burst.KeywordWithParams {
                Keyword: "ユバタス",
                ScalingParam: 1.002,
                Gamma: 0.2,
        }
        cli.AddKeyword(keyword)

        documents := []burst.Document{}

        makeDocuments(&documents, 1.0, 5, 30)
        makeDocuments(&documents, 11.0, 15, 50)
        makeDocuments(&documents, 21.0, 500, 10)
        makeDocuments(&documents, 31.0, 2000, 10)
        makeDocuments(&documents, 41.0, 200, 40)
        makeDocuments(&documents, 51.0, 10, 10)
        makeDocuments(&documents, 61.0, 22225, 25)
        makeDocuments(&documents, 71.0, 22222, 10)
        makeDocuments(&documents, 81.0, 22222, 15)
        makeDocuments(&documents, 91.0, 222, 3)

        cli.AddDocuments(documents)

        if (len(cli.GetAllBurstedResultsAt(71.0)) == 0) {
		t.Errorf("got invalid get_all_bursted_results response")
		t.FailNow()
        }
}
