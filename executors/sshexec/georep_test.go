package sshexec

import (
	"testing"

	"github.com/heketi/heketi/executors"
	"github.com/heketi/heketi/pkg/utils"
	"github.com/heketi/tests"
)

func TestGeoReplicationVolumeStatus(t *testing.T) {
	f := NewFakeSsh()
	defer tests.Patch(&sshNew,
		func(logger *utils.Logger, user string, file string) (Ssher, error) {
			return f, nil
		}).Restore()

	config := &SshConfig{
		PrivateKeyFile: "xkeyfile",
		User:           "xuser",
		CLICommandConfig: CLICommandConfig{
			Fstab: "/my/fstab",
		},
	}

	s, err := NewSshExecutor(config)
	tests.Assert(t, err == nil)
	tests.Assert(t, s != nil)

	// Mock ssh function
	f.FakeConnectAndExec = func(host string,
		commands []string,
		timeoutMinutes int,
		useSudo bool) ([]string, error) {

		tests.Assert(t, host == "host:22", host)
		tests.Assert(t, len(commands) == 1)
		tests.Assert(t, commands[0] == "gluster --mode=script volume geo-replication vol_1 status --xml", commands)

		resp := `
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		<cliOutput>
			<opRet>0</opRet>
			<opErrno>0</opErrno>
			<opErrstr/>
			<geoRep>
				<volume>
					<name>vol_1</name>
					<sessions>
						<session>
							<session_slave>b4927246-9c3e-4a84-bbc5-dd6219a3594a:ssh://1.2.3.4::vol_1:6ce8b0c2-1d7f-4e3d-b9be-a15794ff0a71</session_slave>
							<pair>
								<master_node>1.2.3.3</master_node>
								<master_brick>/var/lib/heketi/mounts/vg_a9f252c537af81890e0549698e1a9eb2/brick_e9fa4e176f71f57b9cf7e3f0e0e0b26d/brick</master_brick>
								<slave_user>root</slave_user>
								<slave>ssh://1.2.3.4::vol_1</slave>
								<slave_node>1.2.3.4</slave_node>
								<status>Passive</status>
								<crawl_status>N/A</crawl_status>
								<entry>N/A</entry>
								<data>N/A</data>
								<meta>N/A</meta>
								<failures>N/A</failures>
								<checkpoint_completed>N/A</checkpoint_completed>
								<master_node_uuid>e7615c4c-ec61-49e8-8346-4dc1a62c6923</master_node_uuid>
								<last_synced>N/A</last_synced>
								<checkpoint_time>N/A</checkpoint_time>
								<checkpoint_completion_time>N/A</checkpoint_completion_time>
							</pair>
							<pair>
								<master_node>1.2.3.4</master_node>
								<master_brick>/var/lib/heketi/mounts/vg_5ccd3561212255f850e72756affe1579/brick_ae7ae603e8168d74a219af61e53724a4/brick</master_brick>
								<slave_user>root</slave_user>
								<slave>ssh://1.2.3.4::vol_1</slave>
								<slave_node>1.2.3.4</slave_node>
								<status>Active</status>
								<crawl_status>Changelog Crawl</crawl_status>
								<entry>0</entry>
								<data>0</data>
								<meta>0</meta>
								<failures>0</failures>
								<checkpoint_completed>N/A</checkpoint_completed>
								<master_node_uuid>04a176b0-2528-45b9-9c8a-4b79b537d99a</master_node_uuid>
								<last_synced>2017-09-04 15:03:26</last_synced>
								<checkpoint_time>N/A</checkpoint_time>
								<checkpoint_completion_time>N/A</checkpoint_completion_time>
							</pair>
						</session>
					</sessions>
				</volume>
		  	</geoRep>
		</cliOutput>
		`
		return []string{resp}, nil
	}

	status, err := s.GeoReplicationVolumeStatus("host", "vol_1")
	tests.Assert(t, err == nil, err)
	tests.Assert(t, status != nil, status)
	tests.Assert(t, len(status.Volume) == 1, len(status.Volume))
	tests.Assert(t, len(status.Volume[0].Sessions.SessionList) == 1, len(status.Volume[0].Sessions.SessionList))
	tests.Assert(t, len(status.Volume[0].Sessions.SessionList[0].Pairs) == 2)
}

func TestGeoReplicationStatus(t *testing.T) {
	f := NewFakeSsh()
	defer tests.Patch(&sshNew,
		func(logger *utils.Logger, user string, file string) (Ssher, error) {
			return f, nil
		}).Restore()

	config := &SshConfig{
		PrivateKeyFile: "xkeyfile",
		User:           "xuser",
		CLICommandConfig: CLICommandConfig{
			Fstab: "/my/fstab",
		},
	}

	s, err := NewSshExecutor(config)
	tests.Assert(t, err == nil)
	tests.Assert(t, s != nil)

	// Mock ssh function
	f.FakeConnectAndExec = func(host string,
		commands []string,
		timeoutMinutes int,
		useSudo bool) ([]string, error) {

		tests.Assert(t, host == "host:22", host)
		tests.Assert(t, len(commands) == 1)
		tests.Assert(t, commands[0] == "gluster --mode=script volume geo-replication status --xml", commands)

		resp := `
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		<cliOutput>
			<opRet>0</opRet>
			<opErrno>0</opErrno>
			<opErrstr/>
			<geoRep>
				<volume>
					<name>vol_1</name>
					<sessions>
						<session>
							<session_slave>b4927246-9c3e-4a84-bbc5-dd6219a3594a:ssh://1.2.3.4::vol_1:6ce8b0c2-1d7f-4e3d-b9be-a15794ff0a71</session_slave>
							<pair>
								<master_node>1.2.3.3</master_node>
								<master_brick>/var/lib/heketi/mounts/vg_a9f252c537af81890e0549698e1a9eb2/brick_e9fa4e176f71f57b9cf7e3f0e0e0b26d/brick</master_brick>
								<slave_user>root</slave_user>
								<slave>ssh://1.2.3.4::vol_1</slave>
								<slave_node>1.2.3.4</slave_node>
								<status>Passive</status>
								<crawl_status>N/A</crawl_status>
								<entry>N/A</entry>
								<data>N/A</data>
								<meta>N/A</meta>
								<failures>N/A</failures>
								<checkpoint_completed>N/A</checkpoint_completed>
								<master_node_uuid>e7615c4c-ec61-49e8-8346-4dc1a62c6923</master_node_uuid>
								<last_synced>N/A</last_synced>
								<checkpoint_time>N/A</checkpoint_time>
								<checkpoint_completion_time>N/A</checkpoint_completion_time>
							</pair>
							<pair>
								<master_node>1.2.3.4</master_node>
								<master_brick>/var/lib/heketi/mounts/vg_5ccd3561212255f850e72756affe1579/brick_ae7ae603e8168d74a219af61e53724a4/brick</master_brick>
								<slave_user>root</slave_user>
								<slave>ssh://1.2.3.4::vol_1</slave>
								<slave_node>1.2.3.4</slave_node>
								<status>Active</status>
								<crawl_status>Changelog Crawl</crawl_status>
								<entry>0</entry>
								<data>0</data>
								<meta>0</meta>
								<failures>0</failures>
								<checkpoint_completed>N/A</checkpoint_completed>
								<master_node_uuid>04a176b0-2528-45b9-9c8a-4b79b537d99a</master_node_uuid>
								<last_synced>2017-09-04 15:03:26</last_synced>
								<checkpoint_time>N/A</checkpoint_time>
								<checkpoint_completion_time>N/A</checkpoint_completion_time>
							</pair>
						</session>
					</sessions>
				</volume>
				<volume>
					<name>vol_2</name>
					<sessions>
						<session>
							<session_slave>b4927246-9c3e-4a84-bbc5-dd6219a3594a:ssh://1.2.3.4::vol_2:6ce8b0c2-1d7f-4e3d-b9be-a15794ff0a71</session_slave>
							<pair>
								<master_node>1.2.3.4</master_node>
								<master_brick>/var/lib/heketi/mounts/vg_5ccd3561212255f850e72756affe1579/brick_ae7ae603e8168d74a219af61e53724a4/brick</master_brick>
								<slave_user>root</slave_user>
								<slave>ssh://1.2.3.4::vol_2</slave>
								<slave_node>1.2.3.4</slave_node>
								<status>Active</status>
								<crawl_status>Changelog Crawl</crawl_status>
								<entry>0</entry>
								<data>0</data>
								<meta>0</meta>
								<failures>0</failures>
								<checkpoint_completed>N/A</checkpoint_completed>
								<master_node_uuid>04a176b0-2528-45b9-9c8a-4b79b537d99a</master_node_uuid>
								<last_synced>2017-09-04 15:03:26</last_synced>
								<checkpoint_time>N/A</checkpoint_time>
								<checkpoint_completion_time>N/A</checkpoint_completion_time>
							</pair>
						</session>
					</sessions>
				</volume>
		  	</geoRep>
		</cliOutput>
		`
		return []string{resp}, nil
	}

	status, err := s.GeoReplicationStatus("host")
	tests.Assert(t, err == nil, err)
	tests.Assert(t, status != nil, status)
	tests.Assert(t, len(status.Volume) == 2, len(status.Volume))
	tests.Assert(t, len(status.Volume[0].Sessions.SessionList) == 1, len(status.Volume[0].Sessions.SessionList))
	tests.Assert(t, len(status.Volume[0].Sessions.SessionList[0].Pairs) == 2)
	tests.Assert(t, len(status.Volume[1].Sessions.SessionList[0].Pairs) == 1)
}

func TestGeoReplicationConfig(t *testing.T) {
	f := NewFakeSsh()
	defer tests.Patch(&sshNew,
		func(logger *utils.Logger, user string, file string) (Ssher, error) {
			return f, nil
		}).Restore()

	config := &SshConfig{
		PrivateKeyFile: "xkeyfile",
		User:           "xuser",
		CLICommandConfig: CLICommandConfig{
			Fstab: "/my/fstab",
		},
	}

	s, err := NewSshExecutor(config)
	tests.Assert(t, err == nil)
	tests.Assert(t, s != nil)

	// Mock ssh function
	f.FakeConnectAndExec = func(host string,
		commands []string,
		timeoutMinutes int,
		useSudo bool) ([]string, error) {

		tests.Assert(t, host == "host:22", host)
		tests.Assert(t, len(commands) == 5)
		want := []string{
			"gluster --mode=script volume geo-replication mastervolume slavehost::slavevolume config checkpoint now",
			"gluster --mode=script volume geo-replication mastervolume slavehost::slavevolume config use-tarssh true",
			"gluster --mode=script volume geo-replication mastervolume slavehost::slavevolume config ignore-deletes 1",
			"gluster --mode=script volume geo-replication mastervolume slavehost::slavevolume config sync-jobs 10",
			"gluster --mode=script volume geo-replication mastervolume slavehost::slavevolume config ssh_port 2222",
		}

		for _, w := range want {
			found := false
			for _, c := range commands {
				if c == w {
					found = true
					break
				}
			}
			tests.Assert(t, found, w)
		}
		return nil, nil

	}

	// Call function
	req := executors.GeoReplicationRequest{
		ActionParams: map[string]string{
			"checkpoint":     "now",
			"use-tarssh":     "true",
			"ignore-deletes": "true",
			"sync-jobs":      "10",
			"ssh-port":       "2222",
		},
		SlaveHost:   "slavehost",
		SlaveVolume: "slavevolume",
	}

	err = s.GeoReplicationConfig("host", "mastervolume", &req)
	tests.Assert(t, err == nil, err)
}

func TestGeoReplicationCreate(t *testing.T) {
	f := NewFakeSsh()
	defer tests.Patch(&sshNew,
		func(logger *utils.Logger, user string, file string) (Ssher, error) {
			return f, nil
		}).Restore()

	config := &SshConfig{
		PrivateKeyFile: "xkeyfile",
		User:           "xuser",
		CLICommandConfig: CLICommandConfig{
			Fstab: "/my/fstab",
		},
	}

	s, err := NewSshExecutor(config)
	tests.Assert(t, err == nil)
	tests.Assert(t, s != nil)

	// Mock ssh function
	f.FakeConnectAndExec = func(host string,
		commands []string,
		timeoutMinutes int,
		useSudo bool) ([]string, error) {

		tests.Assert(t, host == "host:22", host)
		tests.Assert(t, len(commands) == 1)
		tests.Assert(t, commands[0] == "gluster --mode=script volume geo-replication mastervolume slavehost::slavevolume create ssh-port 2222 no-verify", commands)

		return nil, nil

	}

	// Call function
	req := executors.GeoReplicationRequest{
		ActionParams: map[string]string{"option": "no-verify"},
		SlaveHost:    "slavehost",
		SlaveVolume:  "slavevolume",
		SlaveSSHPort: 2222,
	}

	err = s.GeoReplicationCreate("host", "mastervolume", &req)
	tests.Assert(t, err == nil, err)
}
