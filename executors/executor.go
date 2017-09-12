//
// Copyright (c) 2015 The heketi Authors
//
// This file is licensed to you under your choice of the GNU Lesser
// General Public License, version 3 or any later version (LGPLv3 or
// later), or the GNU General Public License, version 2 (GPLv2), in all
// cases as published by the Free Software Foundation.
//

package executors

import "encoding/xml"

type Executor interface {
	GlusterdCheck(host string) error
	PeerProbe(exec_host, newnode string) error
	PeerDetach(exec_host, detachnode string) error
	DeviceSetup(host, device, vgid string) (*DeviceInfo, error)
	GetDeviceInfo(host, device, vgid string) (*DeviceInfo, error)
	DeviceTeardown(host, device, vgid string) error
	BrickCreate(host string, brick *BrickRequest) (*BrickInfo, error)
	BrickDestroy(host string, brick *BrickRequest) error
	BrickDestroyCheck(host string, brick *BrickRequest) error
	VolumeCreate(host string, volume *VolumeRequest) (*Volume, error)
	VolumeDestroy(host string, volume string) error
	VolumeDestroyCheck(host, volume string) error
	VolumeExpand(host string, volume *VolumeRequest) (*Volume, error)
	VolumeReplaceBrick(host string, volume string, oldBrick *BrickInfo, newBrick *BrickInfo) error
	VolumeInfo(host string, volume string) (*Volume, error)
	GeoReplicationCreate(host, volume string, geoRep *GeoReplicationRequest) error
	GeoReplicationConfig(host, volume string, geoRep *GeoReplicationRequest) error
	GeoReplicationAction(host, volume, action string, geoRep *GeoReplicationRequest) error
	GeoReplicationVolumeStatus(host, volume string) (*GeoReplicationStatus, error)
	GeoReplicationStatus(host string) (*GeoReplicationStatus, error)
	HealInfo(host string, volume string) (*HealInfo, error)
	SetLogLevel(level string)
}

type GeoReplicationStatus struct {
	Volume []GeoReplicationVolume `xml:"volume"`
}
type GeoReplicationVolume struct {
	VolumeName string                 `xml:"name"`
	Sessions   GeoReplicationSessions `xml:"sessions"`
}

type GeoReplicationSessions struct {
	SessionList []GeoReplicationSession `xml:"session"`
}

type GeoReplicationSession struct {
	SessionSlave string               `xml:"session_slave"`
	Pairs        []GeoReplicationPair `xml:"pair"`
}
type GeoReplicationPair struct {
	MasterNode               string `xml:"master_node"`
	MasterBrick              string `xml:"master_brick"`
	SlaveUser                string `xml:"slave_user"`
	Slave                    string `xml:"slave"`
	SlaveNode                string `xml:"slave_node"`
	Status                   string `xml:"status"`
	CrawlStatus              string `xml:"crawl_status"`
	Entry                    string `xml:"entry"`
	Data                     string `xml:"data"`
	Meta                     string `xml:"meta"`
	Failures                 string `xml:"failures"`
	CheckpointCompleted      string `xml:"checkpoint_completed"`
	MasterNodeUUID           string `xml:"master_node_uuid"`
	LastSynced               string `xml:"last_string"`
	CheckpointTime           string `xml:"checkpoint_time"`
	CheckpointCompletionTime string `xml:"checkpoint_completion_time"`
}

type GeoReplicationRequest struct {
	ActionParams map[string]string
	SlaveHost    string
	SlaveVolume  string
	SlaveSSHPort int
}

// Enumerate durability types
type DurabilityType int

const (
	DurabilityNone DurabilityType = iota
	DurabilityReplica
	DurabilityDispersion
)

// Returns the size of the device
type DeviceInfo struct {
	// Size in KB
	Size       uint64
	ExtentSize uint64
}

// Brick description
type BrickRequest struct {
	VgId             string
	Name             string
	TpSize           uint64
	Size             uint64
	PoolMetadataSize uint64
	Gid              int64
}

// Returns information about the location of the brick
type BrickInfo struct {
	Path string
	Host string
}

type VolumeRequest struct {
	Bricks               []BrickInfo
	Name                 string
	Type                 DurabilityType
	GlusterVolumeOptions []string

	// Dispersion
	Data       int
	Redundancy int

	// Replica
	Replica int
}

type Brick struct {
	UUID      string `xml:"uuid,attr"`
	Name      string `xml:"name"`
	HostUUID  string `xml:"hostUuid"`
	IsArbiter int    `xml:"isArbiter"`
}

type Bricks struct {
	XMLName   xml.Name `xml:"bricks"`
	BrickList []Brick  `xml:"brick"`
}

type BrickHealStatus struct {
	HostUUID        string `xml:"hostUuid,attr"`
	Name            string `xml:"name"`
	Status          string `xml:"status"`
	NumberOfEntries string `xml:"numberOfEntries"`
}

type Option struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type Options struct {
	XMLName    xml.Name `xml:"options"`
	OptionList []Option `xml:"option"`
}

type Volume struct {
	XMLName         xml.Name `xml:"volume"`
	VolumeName      string   `xml:"name"`
	ID              string   `xml:"id"`
	Status          int      `xml:"status"`
	StatusStr       string   `xml:"statusStr"`
	BrickCount      int      `xml:"brickCount"`
	DistCount       int      `xml:"distCount"`
	StripeCount     int      `xml:"stripeCount"`
	ReplicaCount    int      `xml:"replicaCount"`
	ArbiterCount    int      `xml:"arbiterCount"`
	DisperseCount   int      `xml:"disperseCount"`
	RedundancyCount int      `xml:"redundancyCount"`
	Type            int      `xml:"type"`
	TypeStr         string   `xml:"typeStr"`
	Transport       int      `xml:"transport"`
	Bricks          Bricks
	OptCount        int `xml:"optCount"`
	Options         Options
}

type Volumes struct {
	XMLName    xml.Name `xml:"volumes"`
	Count      int      `xml:"count"`
	VolumeList []Volume `xml:"volume"`
}

type VolInfo struct {
	XMLName xml.Name `xml:"volInfo"`
	Volumes Volumes  `xml:"volumes"`
}

type HealInfoBricks struct {
	BrickList []BrickHealStatus `xml:"brick"`
}

type HealInfo struct {
	XMLName xml.Name       `xml:"healInfo"`
	Bricks  HealInfoBricks `xml:"bricks"`
}
