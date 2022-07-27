package astks

import (
	"errors"
	"github.com/yddeng/astk/pkg/util"
	"log"
	"os"
	"path"
)

type Store interface {
	Load(dataPath string) error
	Save() error
}

type storeBase struct {
	file     string
	filename string
}

var (
	stores   = map[storeName]Store{}
	needSave = map[storeName]bool{}
)

func loadStore(dataPath string) (err error) {
	for name, store := range stores {
		if err = store.Load(dataPath); err != nil {
			return errors.New(string(name) + " : " + err.Error())
		}
	}
	return
}

func saveStore(names ...storeName) {
	if len(names) == 0 {
		for name := range stores {
			needSave[name] = true
		}
	} else {
		for _, name := range names {
			needSave[name] = true
		}
	}
}

func doSave(final bool) {
	if final {
		for name, store := range stores {
			if err := store.Save(); err != nil {
				log.Printf("store %s save failed, %s\n", name, err)
			}
		}
	} else {
		for name := range needSave {
			if store, ok := stores[name]; ok {
				if err := store.Save(); err != nil {
					log.Printf("store %s save failed, %s\n", name, err)
				}
			}
		}
	}
	needSave = map[storeName]bool{}
}

type nodeMgrStore struct {
	storeBase
}

func (store *nodeMgrStore) Load(dataPath string) (err error) {
	store.filename = path.Join(dataPath, store.file)
	if err = util.DecodeJsonFromFile(&nodeMgr, store.filename); err != nil {
		if os.IsNotExist(err) {
			err = nil
			nodeMgr = &NodeMgr{
				Nodes: map[string]*Node{},
				Monitor: &Monitor{
					Cpu:           90,
					Mem:           90,
					Interval:      10,
					AlertInterval: 3600,
				},
			}
		}
		return
	}
	return
}

func (store *nodeMgrStore) Save() error {
	return util.EncodeJsonToFile(nodeMgr, store.filename)
}

type processMgrStore struct {
	storeBase
}

func (store *processMgrStore) Load(dataPath string) (err error) {
	store.filename = path.Join(dataPath, store.file)
	if err = util.DecodeJsonFromFile(&processMgr, store.filename); err != nil {
		if os.IsNotExist(err) {
			err = nil
			processMgr = &ProcessMgr{
				GenID:   0,
				Process: map[int]*Process{},
				Monitor: &Monitor{
					Cpu:           90,
					Mem:           90,
					Interval:      10,
					AlertInterval: 3600,
				},
				TagLabels: map[string]struct{}{},
				TagNodes:  map[string]struct{}{},
			}
			processMgr.refreshLabels()
		}
		return
	}
	processMgr.refreshLabels()
	return
}

func (store *processMgrStore) Save() error {
	return util.EncodeJsonToFile(processMgr, store.filename)
}

type cmdMgrStore struct {
	storeBase
}

func (store *cmdMgrStore) Load(dataPath string) (err error) {
	store.filename = path.Join(dataPath, store.file)
	if err = util.DecodeJsonFromFile(&cmdMgr, store.filename); err != nil {
		if os.IsNotExist(err) {
			err = nil
			cmdMgr = &CmdMgr{
				CmdMap:  map[int]*Cmd{},
				CmdLogs: map[int][]*CmdLog{},
			}
		}
		return
	}
	return
}

func (store *cmdMgrStore) Save() error {
	return util.EncodeJsonToFile(cmdMgr, store.filename)
}

type storeName string

const (
	snNodeMgr    storeName = "node_mgr"
	snCmdMgr     storeName = "cmd_mgr"
	snProcessMgr storeName = "process_mgr"
)

var (
	nodeMgr    *NodeMgr
	cmdMgr     *CmdMgr
	processMgr *ProcessMgr
)

func init() {
	stores[snNodeMgr] = &nodeMgrStore{storeBase{
		file: "node_mgr.json",
	}}
	stores[snCmdMgr] = &cmdMgrStore{storeBase{
		file: "cmd_mgr.json",
	}}
	stores[snProcessMgr] = &processMgrStore{storeBase{
		file: "process_mgr.json",
	}}
}
