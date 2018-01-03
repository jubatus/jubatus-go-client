package jubatus_client

import (
	"fmt"
	"reflect"
	"testing"

	util "./util"
	classifier "github.com/jubatus/jubatus-go-client/lib/classifier"
	common "github.com/jubatus/jubatus-go-client/lib/common"
)

func GetClient(c *util.JubatusProcess) (*classifier.ClassifierClient, error) {
	return classifier.NewClassifierClient("127.0.0.1:"+fmt.Sprintf("%d", c.Port), "tmp")
}

func _bootJubatus(t *testing.T, name string, config string) *util.JubatusProcess {
	jubatus, err := util.NewJubatusProcess("jubaclassifier", "config/classifier.json")
	if err != nil {
		fmt.Print(err)
		t.Errorf("cannot boot jubatus")
		t.FailNow()
	}
	return jubatus
}

func _getClient(t *testing.T, jubatus *util.JubatusProcess) *classifier.ClassifierClient {
	cli, err := classifier.NewClassifierClient(fmt.Sprintf("localhost:%d", jubatus.Port), "as")
	if err != nil {
		fmt.Print(err)
		t.Errorf("cannot connect to jubatus")
		t.FailNow()
	}
	return cli
}

func TestSetLabel(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaclassifier", "config/classifier.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	if cli.SetLabel("label") != true {
		t.Errorf("got invalid set_label response")
		t.FailNow()
	}
}

func TestGetLabels(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaclassifier", "config/classifier.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	cli.SetLabel("label")
	var label = map[string]uint64{"label": 0}
	if !(reflect.DeepEqual(cli.GetLabels(), label)) {
		t.Errorf("got invalid get_labels response")
		t.FailNow()
	}
}

func TestDeleteLabel(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaclassifier", "config/classifier.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	cli.SetLabel("label")
	if cli.DeleteLabel("label") != true {
		t.Errorf("got invalid delete_label response")
		t.FailNow()
	}
}

func TestTrain(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaclassifier", "config/classifier.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	d := common.NewDatum()
	d.AddString("skey1", "val1")
	d.AddString("skey2", "val2")
	d.AddNum("nkey1", 1.0)
	d.AddNum("nkey2", 2.0)
	if cli.Train([]classifier.LabeledDatum{classifier.LabeledDatum{"label", d}}) != 1 {
		t.Errorf("got invalid train response")
		t.FailNow()
	}
}

func TestClassify(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaclassifier", "config/classifier.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	d := common.NewDatum()
	d.AddString("skey1", "val1")
	d.AddString("skey2", "val2")
	d.AddNum("nkey1", 1.0)
	d.AddNum("nkey2", 2.0)
	cli.Classify([]common.Datum{d})
}

func TestSave(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaclassifier", "config/classifier.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	if len(cli.Save("classifier.save_test.model")) != 1 {
		t.Errorf("got invalid save response")
		t.FailNow()
	}
}

func TestLoad(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaclassifier", "config/classifier.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	if cli.Load("classifier.save_test.model") != true {
		t.Errorf("got invalid load response")
		t.FailNow()
	}
}

func TestClear(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaclassifier", "config/classifier.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	if cli.Clear() != true {
		t.Errorf("got invalid clear response")
		t.FailNow()
	}
}

func TestGetStatus(t *testing.T) {
	jubatus := _bootJubatus(t, "jubaclassifier", "config/classifier.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	cli.GetStatus()
}
