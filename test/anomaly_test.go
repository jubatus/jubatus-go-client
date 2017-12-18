package jubatus_client

import (
	"fmt"
	"testing"

	util "./util"
	anomaly "github.com/jubatus/jubatus-go-client/lib/anomaly"
	common "github.com/jubatus/jubatus-go-client/lib/common"
)

func GetClient(c *util.JubatusProcess) (*anomaly.AnomalyClient, error) {
	return anomaly.NewAnomalyClient("127.0.0.1:"+fmt.Sprintf("%d", c.Port), "tmp")
}

func _bootJubatus(t *testing.T, name string, config string) *util.JubatusProcess {
	jubatus, err := util.NewJubatusProcess("jubaanomaly", "config/anomaly.json")
	if err != nil {
		fmt.Print(err)
		t.Errorf("cannot boot jubatus")
		t.FailNow()
	}
	return jubatus
}

func _getClient(t *testing.T, jubatus *util.JubatusProcess) *anomaly.AnomalyClient {
	cli, err := anomaly.NewAnomalyClient(fmt.Sprintf("localhost:%d", jubatus.Port), "as")
	if err != nil {
		fmt.Print(err)
		t.Errorf("cannot connect to jubatus")
		t.FailNow()
	}
	return cli
}

func TestAdd(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaanomaly", "config/anomaly.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	d := common.NewDatum()
	cli.Add(d)
}

func TestUpdate(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaanomaly", "config/anomaly.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	d := common.NewDatum()
	result := cli.Add(d)
	d.AddNum("val", 3.1)
	cli.Update(result.ID, d)
}

func TestOverwrite(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaanomaly", "config/anomaly.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	d := common.NewDatum()
	result := cli.Add(d)
	d.AddNum("val", 3.1)
	cli.Overwrite(result.ID, d)
}

func TestClearRow(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaanomaly", "config/anomaly.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	d := common.NewDatum()

	result := cli.Add(d)
	if cli.ClearRow(result.ID) != true {
		t.Errorf("got invalid clear_row response")
		t.FailNow()
	}
}

func TestCalcScore(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaanomaly", "config/anomaly.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	d := common.NewDatum()
	d.AddNum("val", 1.1)
	cli.Add(d)
	d.AddNum("val", 3.1)
	cli.CalcScore(d)
}

func TestGetAllRows(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaanomaly", "config/anomaly.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	d := common.NewDatum()
	d.AddNum("val1", 1.1)
	cli.Add(d)
	d.AddNum("val2", 3.1)
	cli.Add(d)

	if len(cli.GetAllRows()) != 2 {
		t.Errorf("got invalid get_all_rows response")
		t.FailNow()
	}
}
