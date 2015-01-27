package jubatus_client

import (
	"fmt"
	"testing"

	common "../lib/common"
	recommender "../lib/recommender"
	util "./util"
)

func GetClient(c *util.JubatusProcess) (*recommender.RecommenderClient, error) {
	return recommender.NewRecommenderClient("127.0.0.1:"+fmt.Sprintf("%d", c.Port), "tmp")
}

func _bootJubatus(t *testing.T, name string, config string) *util.JubatusProcess {
	jubatus, err := util.NewJubatusProcess("jubarecommender", "config/recommender.json")
	if err != nil {
		fmt.Print(err)
		t.Errorf("cannot boot jubatus")
		t.FailNow()
	}
	return jubatus
}

func _getClient(t *testing.T, jubatus *util.JubatusProcess) *recommender.RecommenderClient {
	cli, err := recommender.NewRecommenderClient(fmt.Sprintf("localhost:%d", jubatus.Port), "as")
	if err != nil {
		fmt.Print(err)
		t.Errorf("cannot connect to jubatus")
		t.FailNow()
	}
	return cli
}

func TestUpdateRow(t *testing.T) {
	jubatus := _bootJubatus(t, "jubarecommender", "config/recommender.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	rows_before := cli.GetAllRows()
	if 0 < len(rows_before) {
		t.Errorf("some row already exist before update")
		t.FailNow()
	}

	d := common.NewDatum()
	d.AddString("a", "b")
	cli.UpdateRow("a", d)

	rows_after := cli.GetAllRows()
	if len(rows_after) != 1 {
		t.Errorf("some row already exist before update")
		t.FailNow()
	}
}

func TestClearRow(t *testing.T) {
	jubatus := _bootJubatus(t, "jubarecommender", "config/recommender.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	d := common.NewDatum()
	d.AddString("a", "b")
	cli.UpdateRow("a", d)

	cli.ClearRow("a")

	rows_after := cli.GetAllRows()
	if 0 < len(rows_after) {
		t.Errorf("some row exist after clear row")
		t.FailNow()
	}
}

func TestClear(t *testing.T) {
	jubatus := _bootJubatus(t, "jubarecommender", "config/recommender.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)

	for i := 0; i < 100; i++ {
		d := common.NewDatum()
		d.AddString(fmt.Sprintf("a%d", i), "b")
		cli.UpdateRow(fmt.Sprintf("a%d", i), d)
	}
	rows_before := cli.GetAllRows()
	if len(rows_before) != 100 {
		t.Errorf("some row may be corrupted")
		t.FailNow()
	}

	cli.Clear()

	rows_after := cli.GetAllRows()
	if len(rows_after) != 0 {
		t.Errorf("failed to delete all rows")
		t.FailNow()
	}
}

func TestCompleteRowFromID(t *testing.T) {
	jubatus := _bootJubatus(t, "jubarecommender", "config/recommender.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)

	for i := 0; i < 100; i++ {
		d := common.NewDatum()
		d.AddNum(fmt.Sprintf("a%d", i), float64(i))
		d.AddNum(fmt.Sprintf("a%d", i+1), float64(i+1))
		cli.UpdateRow(fmt.Sprintf("a%d", i), d)
	}

	result := cli.CompleteRowFromId("a20")
	if 100 != len(result.NumValues()) {
		t.Errorf("invalid result length")
		t.FailNow()
	}
}

func TestCompleteRowFromDatum(t *testing.T) {
	jubatus := _bootJubatus(t, "jubarecommender", "config/recommender.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)

	for i := 0; i < 100; i++ {
		d := common.NewDatum()
		d.AddNum(fmt.Sprintf("a%d", i), float64(i))
		d.AddNum(fmt.Sprintf("a%d", i+1), float64(i+1))
		cli.UpdateRow(fmt.Sprintf("a%d", i), d)
	}

	result := cli.CompleteRowFromId("a20")
	if 100 != len(result.NumValues()) {
		t.Errorf("invalid result length")
		t.FailNow()
	}
}
