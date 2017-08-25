package glusterfs

import (
	"encoding/json"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
	"github.com/heketi/heketi/pkg/glusterfs/api"
	"github.com/heketi/utils"
)

// VolumeGeoReplicationStatus is the handler returning the geo-replication session
// status for a specific volume
func (a *App) VolumeGeoReplicationStatus(w http.ResponseWriter, r *http.Request) {
	logger.Debug("In VolumeGeoReplication")

	vars := mux.Vars(r)
	id := vars["id"]

	var volume *VolumeEntry
	var host string
	var err error

	err = a.db.View(func(tx *bolt.Tx) error {
		volume, err = NewVolumeEntryFromId(tx, id)
		if err == ErrNotFound {
			http.Error(w, "Volume Id not found", http.StatusNotFound)
			return err
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		cluster, err := NewClusterEntryFromId(tx, volume.Info.Cluster)
		if err == ErrNotFound {
			http.Error(w, "Cluster Id not found", http.StatusNotFound)
			return err
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		node, err := NewNodeEntryFromId(tx, cluster.Info.Nodes[0])
		if err == ErrNotFound {
			http.Error(w, "Node Id not found", http.StatusNotFound)
			return err
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		host = node.ManageHostName()

		return nil
	})
	if err != nil {
		return
	}

	resp, err := volume.NewGeoReplicationStatusResponse(a.executor, host)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		panic(err)
	}
}

// VolumeGeoReplicationDelete is the handler for deleting the geo-replication session
// for a specific volume
func (a *App) VolumeGeoReplicationDelete(w http.ResponseWriter, r *http.Request) {
}

// VolumeGeoReplication is the handler for managing a geo-replication session
func (a *App) VolumeGeoReplication(w http.ResponseWriter, r *http.Request) {
	logger.Debug("In VolumeGeoReplication")

	vars := mux.Vars(r)
	id := vars["id"]

	var volume *VolumeEntry
	var host string
	var err error

	var msg api.GeoReplicationRequest
	if err := utils.GetJsonFromRequest(r, &msg); err != nil {
		http.Error(w, "request unable to be parsed", http.StatusUnprocessableEntity)
		return
	}
	logger.Debug("Msg: %v", msg)

	err = a.db.View(func(tx *bolt.Tx) error {
		volume, err = NewVolumeEntryFromId(tx, id)
		if err == ErrNotFound {
			http.Error(w, "Volume Id not found", http.StatusNotFound)
			return err
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		cluster, err := NewClusterEntryFromId(tx, volume.Info.Cluster)
		if err == ErrNotFound {
			http.Error(w, "Cluster Id not found", http.StatusNotFound)
			return err
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		node, err := NewNodeEntryFromId(tx, cluster.Info.Nodes[0])
		if err == ErrNotFound {
			http.Error(w, "Node Id not found", http.StatusNotFound)
			return err
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		host = node.ManageHostName()

		return nil
	})
	if err != nil {
		return
	}

	// Perform GeoReplication action on volume in an asynchronous function
	a.asyncManager.AsyncHttpRedirectFunc(w, r, func() (string, error) {
		switch msg.Action {
		case api.GeoReplicationActionCreate:
			logger.Info("Creating geo-replication session for volume %s", volume.Info.Id)
			volume.GeoReplicationCreate(a.db, a.executor, host, msg)
		default:
			logger.LogError("Unsupported action %s", msg.Action)
		}
		return "/volumes/" + volume.Info.Id + "/georeplication", nil
	})

}
