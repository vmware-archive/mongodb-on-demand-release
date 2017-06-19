package adapter

import (
	"reflect"
	"text/template"
)

const (
	PlanStandalone       = "standalone"
	PlanShardedSet       = "sharded_set"
	PlanSingleReplicaSet = "single_replica_set"
)

var plans = map[string]*template.Template{}

func init() {
	funcs := template.FuncMap{
		"last": func(a interface{}, x int) bool {
			return reflect.ValueOf(a).Len()-1 == x
		},
		"isConfig": func(index int) bool {
			return index >= 9 && index < 12
		},
		"isInShard": func(index int) bool {
			return index < 12
		},
		"hasStorage": func(index int) bool {
			return index < 12
		},
		"processType": func(index int) string {
			if index > 11 && index < 15 {
				return "mongos"
			} else {
				return "mongod"
			}
		},
		"hasShardedCluster": func(index int) bool {
			return index > 11 && index < 15
		},
		"shardNumber": func(index int) int {
			return index % 3
		},
	}

	for k, s := range plansRaw {
		t := template.New(k).Funcs(funcs)
		t, err := t.Parse(s)
		if err != nil {
			panic(err)
		}

		t = t.Funcs(funcs)

		plans[k] = t
	}
}

var plansRaw = map[string]string{
	PlanStandalone: `{
    "options": {
        "downloadBase": "/var/lib/mongodb-mms-automation",
        "downloadBaseWindows": "C:\\mongodb-mms-automation"
    },
    "mongoDbVersions": [{
        "builds": [
                {
                    "bits": 64,
                    "flavor": "",
                    "gitVersion": "4249c1d2b5999ebbf1fdf3bc0e0e3b3ff5c0aaf2",
                    "maxOsVersion": "",
                    "minOsVersion": "",
                    "modules": [],
                    "platform": "osx",
                    "url": "https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-amazon-3.2.7.tgz",
                    "win2008plus": false,
                    "winVCRedistDll": "",
                    "winVCRedistOptions": [],
                    "winVCRedistUrl": "",
                    "winVCRedistVersion": ""
                }
            ],
        "name": "{{.Version}}"
    }],
    "backupVersions": [{
        "hostname": "{{index .Nodes 0}}",
        "logPath": "/var/vcap/sys/log/mongod_node/backup-agent.log",
        "logRotate": {
            "sizeThresholdMB": 1000,
            "timeThresholdHrs": 24
        }
    }],

    "monitoringVersions": [{
        "hostname": "{{index .Nodes 0}}",
        "logPath": "/var/vcap/sys/log/mongod_node/monitoring-agent.log",
        "logRotate": {
            "sizeThresholdMB": 1000,
            "timeThresholdHrs": 24
        }
    }],
    "processes": [{
        "args2_6": {
            "net": {
                "port": 28000
            },
            "storage": {
                "dbPath": "/var/vcap/store/mongodb-data"
            },
            "systemLog": {
                "destination": "file",
                "path": "/var/vcap/sys/log/mongod_node/mongodb.log"
            }
        },
        "hostname": "{{index .Nodes 0}}",
        "logRotate": {
            "sizeThresholdMB": 1000,
            "timeThresholdHrs": 24
        },
        "name": "{{index .Nodes 0}}",
        "processType": "mongod",
        "version": "{{.Version}}",
        "authSchemaVersion": 5
    }],
    "replicaSets": [],
    "roles": [],
    "sharding": [],

    "auth": {
        "autoUser": "mms-automation",
        "autoPwd": "{{.Password}}",
        "deploymentAuthMechanisms": [
            "SCRAM-SHA-1"
        ],
        "key": "{{.key}}",
        "keyfile": "/var/vcap/jobs/mongod_node/config/mongo_om.key",
        "disabled": false,
        "usersDeleted": [],
        "usersWanted": [
            {
                "db": "admin",
                "roles": [
                    {
                        "db": "admin",
                        "role": "clusterMonitor"
                    }
                ],
                "user": "mms-monitoring-agent",
                "initPwd": "{{.Password}}"
            },
            {
                "db": "admin",
                "roles": [
                    {
                        "db": "admin",
                        "role": "clusterAdmin"
                    },
                    {
                        "db": "admin",
                        "role": "readAnyDatabase"
                    },
                    {
                        "db": "admin",
                        "role": "userAdminAnyDatabase"
                    },
                    {
                        "db": "local",
                        "role": "readWrite"
                    },
                    {
                        "db": "admin",
                        "role": "readWrite"
                    }
                ],
                "user": "mms-backup-agent",
                "initPwd": "{{.Password}}"
            },
            {
               "db": "admin" ,
               "user": "admin" ,
               "roles": [
                 {
                     "db": "admin",
                     "role": "clusterAdmin"
                 },
                 {
                     "db": "admin",
                     "role": "readAnyDatabase"
                 },
                 {
                     "db": "admin",
                     "role": "userAdminAnyDatabase"
                 },
                 {
                     "db": "local",
                     "role": "readWrite"
                 },
                 {
                     "db": "admin",
                     "role": "readWrite"
                 }
               ],
               "initPwd": "{{.admin_password}}"
            }
        ],
        "autoAuthMechanism": "SCRAM-SHA-1"
    }
}`,

	PlanShardedSet: `{
   "options": {
        "downloadBase": "/var/lib/mongodb-mms-automation",
        "downloadBaseWindows": "C:\\mongodb-mms-automation"
    },
    "mongoDbVersions": [{
        "builds": [
                {
                    "bits": 64,
                    "flavor": "",
                    "gitVersion": "4249c1d2b5999ebbf1fdf3bc0e0e3b3ff5c0aaf2",
                    "maxOsVersion": "",
                    "minOsVersion": "",
                    "modules": [],
                    "platform": "osx",
                    "url": "https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-amazon-3.2.7.tgz",
                    "win2008plus": false,
                    "winVCRedistDll": "",
                    "winVCRedistOptions": [],
                    "winVCRedistUrl": "",
                    "winVCRedistVersion": ""
                }
            ],
        "name": "{{.Version}}"
    }],
    "backupVersions": [
    ],

    "monitoringVersions": [
    {
        "hostname": "{{index .Nodes 0}}",
        "logPath": "/var/vcap/sys/log/mongod_node/monitoring-agent.log",
        "logRotate": {
            "sizeThresholdMB": 1000,
            "timeThresholdHrs": 24
        }
    }
    ],
    "processes": [{{range $i, $node := .Nodes}}
      {{if $i}},{{end}}{
        "args2_6": {
            "net": {
                "port": 28000
            },
            {{if isInShard $i}}
            "replication": {
                "replSetName": "shard_{{shardNumber $i}}"
            },
            {{end}}
            {{if isConfig $i}}
            "sharding": {
                "clusterRole": "configsvr"
            },
            {{end}}
            {{if hasStorage $i}}
            "storage": {
                "dbPath": "/var/vcap/store/mongodb-data"
            },
            {{end}}
            "systemLog": {
                "destination": "file",
                "path": "/var/vcap/sys/log/mongod_node/mongodb.log"
            }
        },
        "hostname": "{{$node}}",
        "logRotate": {
            "sizeThresholdMB": 1000,
            "timeThresholdHrs": 24
        },
        "name": "{{$node}}",
        "processType": "{{processType $i}}",
        {{if hasShardedCluster $i}}
        "cluster": "sharded-cluster",
        {{end}}
        "version": "{{$.Version}}",
        "authSchemaVersion": 5
    }
    {{end}}
    ],
    "replicaSets": [{{range $i, $shard := .Shards}}
        {{if $i}},{{end}}{
           "_id": "shard_{{$.ID}}_{{$i}}",
           "members": [{{range $i, $node := $shard}}
               {{if $i}},{{end}}{
                   "_id": {{$i}},
                   "arbiterOnly": false,
                   "hidden": false,
                   "host": "{{$node}}",
                   "priority": 1,
                   "slaveDelay": 0,
                   "votes": 1
               }
               {{end}}
           ]
       }
       {{end}}
    ],
    "sharding": [
        {
                "shards": [
                  {{range $i, $shard := .Shards}}
                    {{if $i}},{{end}}{
                        "tags": [],
                        "_id": "shard_{{$.ID}}_{{$i}}",
                        "rs": "shard_{{$.ID}}_{{$i}}"
                    }
                  {{end}}
                ],
                "name": "sharded-cluster",
                "configServer": [],
                "configServerReplica": "shard_{{.ID}}_0",
                "collections": []
            }
    ],

    "auth":{
       "disabled":false,
       "autoPwd": "{{.Password}}",
       "autoUser":"mms-automation",
       "deploymentAuthMechanisms": [
           "MONGODB-CR"
       ],
       "key":"{{.Key}}",
       "keyfile":"/var/vcap/jobs/mongod_node/config/mongo_om.key",
       "usersWanted":[
          {
             "db":"admin",
             "initPwd":"{{.Password}}",
             "roles":[
                {
                   "db":"admin",
                   "role":"clusterMonitor"
                }
             ],
             "user":"mms-monitoring-agent"
          },
          {
             "db":"admin",
             "initPwd":"{{.Password}}",
             "roles":[
                {
                   "db":"admin",
                   "role":"clusterAdmin"
                },
                {
                   "db":"admin",
                   "role":"readAnyDatabase"
                },
                {
                   "db":"admin",
                   "role":"userAdminAnyDatabase"
                },
                {
                   "db":"local",
                   "role":"readWrite"
                },
                {
                   "db":"admin",
                   "role":"readWrite"
                }
             ],
             "user":"mms-backup-agent"
          }
       ],
       "usersDeleted":[],
       "autoAuthMechanism": "MONGODB-CR"
    }
}`,

	PlanSingleReplicaSet: `{
    "options": {
        "downloadBase": "/var/lib/mongodb-mms-automation",
        "downloadBaseWindows": "C:\\mongodb-mms-automation"
    },
    "mongoDbVersions": [{
        "builds": [
                {
                    "bits": 64,
                    "flavor": "",
                    "gitVersion": "4249c1d2b5999ebbf1fdf3bc0e0e3b3ff5c0aaf2",
                    "maxOsVersion": "",
                    "minOsVersion": "",
                    "modules": [],
                    "platform": "osx",
                    "url": "https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-amazon-3.2.7.tgz",
                    "win2008plus": false,
                    "winVCRedistDll": "",
                    "winVCRedistOptions": [],
                    "winVCRedistUrl": "",
                    "winVCRedistVersion": ""
                }
            ],
        "name": "{{.Version}}"
    }],
    "backupVersions": [{
        "hostname": "{{index .Nodes 0}}",
        "logPath": "/var/vcap/sys/log/mongod_node/backup-agent.log",
        "logRotate": {
            "sizeThresholdMB": 1000,
            "timeThresholdHrs": 24
        }
    }],

    "monitoringVersions": [{
        "hostname": "{{index .Nodes 0}}",
        "logPath": "/var/vcap/sys/log/mongod_node/monitoring-agent.log",
        "logRotate": {
            "sizeThresholdMB": 1000,
            "timeThresholdHrs": 24
        }
    }],
    "processes": [{{range $i, $node := .Nodes}}
      {{if $i}},{{end}}{
        "args2_6": {
            "net": {
                "port": 28000
            },
            "replication": {
                "replSetName": "pcf_repl"
            },
            "storage": {
                "dbPath": "/var/vcap/store/mongodb-data"
            },
            "systemLog": {
                "destination": "file",
                "path": "/var/vcap/sys/log/mongod_node/mongodb.log"
            }
        },
        "hostname": "{{$node}}",
        "logRotate": {
            "sizeThresholdMB": 1000,
            "timeThresholdHrs": 24
        },
        "name": "{{$node}}",
        "processType": "mongod",
        "version": "{{$.Version}}",
        "authSchemaVersion": 5
    }
    {{end}}
  ],
    "replicaSets": [{
        "_id": "pcf_repl",
        "members": [
          {{range $i, $node := .Nodes}}
          {{if $i}},{{end}}{
            "_id": {{$i}},
            "host": "{{$node}}"{{if last $.Nodes $i}},
            "arbiterOnly": true,
            "priority": 0
            {{end}}
          }
          {{end}}
        ]
    }],
    "roles": [],
    "sharding": [],

    "auth": {
        "autoUser": "mms-automation",
        "autoPwd": "{{.Password}}",
        "deploymentAuthMechanisms": [
            "SCRAM-SHA-1"
        ],
        "key": "{{.Key}}",
        "keyfile": "/var/vcap/jobs/mongod_node/config/mongo_om.key",
        "disabled": false,
        "usersDeleted": [],
        "usersWanted": [
            {
                "db": "admin",
                "roles": [
                    {
                        "db": "admin",
                        "role": "clusterMonitor"
                    }
                ],
                "user": "mms-monitoring-agent",
                "initPwd": "{{.Password}}"
            },
            {
                "db": "admin",
                "roles": [
                    {
                        "db": "admin",
                        "role": "clusterAdmin"
                    },
                    {
                        "db": "admin",
                        "role": "readAnyDatabase"
                    },
                    {
                        "db": "admin",
                        "role": "userAdminAnyDatabase"
                    },
                    {
                        "db": "local",
                        "role": "readWrite"
                    },
                    {
                        "db": "admin",
                        "role": "readWrite"
                    }
                ],
                "user": "mms-backup-agent",
                "initPwd": "{{.Password}}"
            },
            {
               "db": "admin" ,
               "user": "admin" ,
               "roles": [
                 {
                     "db": "admin",
                     "role": "clusterAdmin"
                 },
                 {
                     "db": "admin",
                     "role": "readAnyDatabase"
                 },
                 {
                     "db": "admin",
                     "role": "userAdminAnyDatabase"
                 },
                 {
                     "db": "local",
                     "role": "readWrite"
                 },
                 {
                     "db": "admin",
                     "role": "readWrite"
                 }
               ],
               "initPwd": "{{.AdminPassword}}"
            }
        ],
        "autoAuthMechanism": "SCRAM-SHA-1"
    }
}`,
}
