package glusterfs

import (
	"github.com/boltdb/bolt"
	"github.com/heketi/heketi/executors"
	"github.com/heketi/heketi/pkg/glusterfs/api"
	"github.com/lpabon/godbc"
)

func (v *VolumeEntry) GeoReplicationCreate(db *bolt.DB,
	executor executors.Executor,
	host string,
	msg api.GeoReplicationRequest) error {

	logger.Debug("In GeoReplicationCreate")

	godbc.Require(db != nil)

	geoRep := &executors.GeoReplicationRequest{
		ActionParams: msg.ActionParams,
		SlaveVolume:  msg.SlaveVolume,
		SlaveHost:    msg.SlaveHost,
		SlaveSSHPort: msg.SlaveSSHPort,
	}

	if err := executor.GeoReplicationCreate(host, v.Info.Name, geoRep); err != nil {
		return err
	}

	return nil
}

func (v *VolumeEntry) NewGeoReplicationStatusResponse(executor executors.Executor,
	host string) (resp *api.GeoReplicationVolumeStatus, err error) {

	sessions, err := executor.GeoReplicationVolumeStatus(host, v.Info.Name)
	if err != nil {
		return nil, err
	}

	resp = &api.GeoReplicationVolumeStatus{
		Volume: api.GeoReplicationVolume{
			VolumeName: v.Info.Name,
			Sessions:   api.GeoReplicationSessions{},
		},
	}

	for _, session := range sessions {
		p := []api.GeoReplicationPair{}
		for _, pair := range session.Pairs {
			p = append(p, api.GeoReplicationPair{
				Status: pair.Status,
			})
		}

		resp.Volume.Sessions.SessionList = append(resp.Volume.Sessions.SessionList, api.GeoReplicationSession{
			SessionSlave: session.SessionSlave,
			Pairs:        p,
		})
	}
	return nil, nil
}
