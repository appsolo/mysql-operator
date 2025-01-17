/*
Copyright 2018 Pressinfra SRL

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package constants

import "github.com/blang/semver"

const (
	// MysqlPort is the default mysql port.
	MysqlPort = 3306

	// UserID should be selected according to different versions of mysql
	Mysql57UserID = 999
	Mysql8UserID  = 1001

	// OrcTopologyDir path where orc conf secret is mounted
	OrcTopologyDir = "/var/run/orc-topology"

	// SidecarServerPort represents the port on which http server will run
	SidecarServerPort = 8080
	// SidecarServerProbePath the probe path
	SidecarServerProbePath = "/health"

	// ExporterPort is the port that metrics will be exported
	ExporterPort = 9125

	// ExporterPath is the path on which metrics are expose
	ExporterPath = "/metrics"

	// OperatorDbName represent the database name that is used by operator to
	// manage the mysql cluster. This database contains a table with
	// initialization history and table managed by pt-heartbeat. Be aware that
	// when changing this value to update the orchestrator chart value for
	// SlaveLagQuery in charts/mysql-operator/values.yaml.
	OperatorDbName = "sys_operator"

	// OperatorStatusTableName represents the name of the table that contains information about MySQL status, like:
	// if mysql is configure by the operator, if PURGE_GTID is set or not, etc
	OperatorStatusTableName = "status"

	// ConfVolumeMountPath is the path where mysql configs will be mounted
	ConfVolumeMountPath = "/etc/mysql"

	// DataVolumeMountPath is the path to mysql data
	DataVolumeMountPath = "/var/lib/mysql"

	// TmpfsVolumeMountPath is the path for the tmpfs mount
	TmpfsVolumeMountPath = "/tmp"

	// ConfMapVolumeMountPath represents the temp config mount path in init containers
	ConfMapVolumeMountPath = "/mnt/conf"

	// ConfDPath is the path to extra mysql configs dir
	ConfDPath = "/etc/mysql/conf.d"

	// ConfClientPath represents the path to the client MySQL client configuration
	// it's important to have a different extension than .cnf to be ignore by MySQL include
	ConfClientPath = "/etc/mysql/client.conf"

	// ConfHeartBeatPath the path where to put the heartbeat.conf file
	// it's important to have a different extension than .cnf to be ignore by MySQL include
	ConfHeartBeatPath = "/etc/mysql/heartbeat.conf"

	// RcloneConfigFile represents the path to the file that contains rclone
	// configs. This path should be the same as defined in docker entrypoint
	// script from mysql-operator-sidecar/docker-entrypoint.sh. /tmp/rclone.conf
	RcloneConfigFile = "/tmp/rclone.conf"

	// ShPreStop used in mysql container, if the pod to be deleted is master, then preStop would do GracefulMasterTakeover
	// before mysql container is deleted.
	ShPreStop = "pre-shutdown-ha.sh"
)

var (
	// MySQLDefaultVersion is the version for mysql that should be used
	MySQLDefaultVersion = semver.MustParse("5.7.35")
	// MySQLTagsToSemVer maps simple version to semver versions
	MySQLTagsToSemVer = map[string]string{
		"5.7": "5.7.35",
		"8.0": "8.0.20",
	}
	// MysqlImageVersions is a map of supported mysql version and their image
	MysqlImageVersions = map[string]string{
		// percona:5.7.35 CentOS based image
		"5.7.35": "percona@sha256:caab4e854bd75040d07802bf1862bfef1d2b4db0acbc9c4aaf5c21c698fdd393",
		// percona:5.7.31-centos
		"5.7.31": "percona@sha256:68dad5e2efeb6893e2d7d116a1eae144f2c641c17d00e7869397395590c91651",
		// This version of mysql has a bug and doesn't work with the operator,
		// see: https://github.com/bitpoke/mysql-operator/issues/509
		"5.7.29": "percona@sha256:d801123bbfaf750924f993f5c59189d144a93feb928b8aef95e541dd61c62881",
		// Percona:5.7.26 CentOS based image
		"5.7.26": "percona@sha256:713c1817615b333b17d0fbd252b0ccc53c48a665d4cfcb42178167435a957322",
		// Percona:5.7.24 CentOS based image
		"5.7.24": "percona@sha256:b3b7fb177b416563c46fe012298e042ec1607cc0539ce6014146380b0d27b08c",
		// Percona:8.0.20-11 CentOS based image
		"8.0.20": "percona@sha256:6d4524eccd26af7bd7fb623c567159dfbd7f3d9a0e2f7bebd54af1e9ca9903dc",
	}
)
