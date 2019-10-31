package datamgr

import (
	"base/log"
	"encoding/json"
	"io/ioutil"
	"sync"
)

type DataMgr struct {
	propData map[uint32]map[string]uint32
}


var (
	mgr     *DataMgr
	mgrOnce sync.Once
)

func GetMe() *DataMgr {
	if mgr == nil {
		mgrOnce.Do(func() {
			mgr = &DataMgr{

			}
			mgr.Init()
		})
	}
	return mgr
}

func (this *DataMgr) Init() {
	this.initPropData()
}

func (this *DataMgr) initPropData() {
	this.propData = make(map[uint32]map[string]uint32)

	fileData, err := ioutil.ReadFile("D:\\GoPath\\src\\jsondata\\prop.json")
	if err != nil {
		log.Error.Println("initPropData err:", err)
		return
	}

	err = json.Unmarshal(fileData, &this.propData)
	if err != nil {
		log.Error.Println("initPropData err:", err)
		return
	}
	log.Info.Println("initPropData success")
}

func (this *DataMgr) GetPropData(u uint32) map[string]uint32 {
	if data, ok := this.propData[u]; ok {
		return data
	}
	log.Error.Println("GetPropData wrong u:", u)
	return nil
}