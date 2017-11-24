package jubatus_client

import (
	"fmt"
	"testing"

	util "./util"
	bandit "github.com/jubatus/jubatus-go-client/lib/bandit"
)

func GetClient(c *util.JubatusProcess) (*bandit.BanditClient, error) {
	return bandit.NewBanditClient("127.0.0.1:"+fmt.Sprintf("%d", c.Port), "tmp")
}

func _bootJubatus(t *testing.T, name string, config string) *util.JubatusProcess {
	jubatus, err := util.NewJubatusProcess("jubabandit", "config/bandit.json")
	if err != nil {
		fmt.Print(err)
		t.Errorf("cannot boot jubatus")
		t.FailNow()
	}
	return jubatus
}

func _getClient(t *testing.T, jubatus *util.JubatusProcess) *bandit.BanditClient {
	cli, err := bandit.NewBanditClient(fmt.Sprintf("localhost:%d", jubatus.Port), "as")
	if err != nil {
		fmt.Print(err)
		t.Errorf("cannot connect to jubatus")
		t.FailNow()
	}
	return cli
}

func TestRegisterArm(t *testing.T) {
	jubatus := _bootJubatus(t, "jubabandit", "config/bandit.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	if cli.RegisterArm("1") != true {
		t.Errorf("failed register arm")
		t.FailNow()
	}
}

func TestDeleteArm(t *testing.T) {
	jubatus := _bootJubatus(t, "jubabandit", "config/bandit.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	cli.RegisterArm("1")
	if cli.DeleteArm("1") != true {
		t.Errorf("failed delete arm")
		t.FailNow()
	}
}

func TestSelectArm(t *testing.T) {
	jubatus := _bootJubatus(t, "jubabandit", "config/bandit.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	cli.SelectArm("bandit")
}

func TestRegisterReward(t *testing.T) {
	jubatus := _bootJubatus(t, "jubabandit", "config/bandit.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	cli.RegisterArm("1")
	if cli.RegisterReward("bandit", "1", 3.0) != true {
		t.Errorf("failded register reward")
		t.FailNow()
	}
}

func TestGetArmInfo(t *testing.T) {
	jubatus := _bootJubatus(t, "jubabandit", "config/bandit.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	cli.RegisterReward("bandit", "1", 1.1)
	cli.GetArmInfo("bandit")
}

func TestReset(t *testing.T) {
	jubatus := _bootJubatus(t, "jubabandit", "config/bandit.json")
	defer jubatus.Kill()
	cli := _getClient(t, jubatus)
	cli.RegisterReward("bandit", "1", 1.1)
	if cli.Reset("bandit") != true {
		t.Errorf("failded reset arm info")
		t.FailNow()
	}
}
