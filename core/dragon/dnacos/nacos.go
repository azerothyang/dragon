package dnacos

import (
	"dragon/core/dragon/conf"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"strconv"
)

// https://github.com/nacos-group/nacos-sdk-go
var NamingClient naming_client.INamingClient

func Init() {
	clientConfig := constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./log/nacos/log",
		CacheDir:            "./log/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: conf.Conf.Nacos.Ip,
			Port:   conf.Conf.Nacos.Port,
		},
	}

	// 创建服务发现客户端
	var err error
	NamingClient, err = clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		log.Println("创建服务发现客户端失败", err)
	}

	// 创建动态配置客户端
	//configClient, err := clients.CreateConfigClient(map[string]interface{}{
	//	"serverConfigs": serverConfigs,
	//	"clientConfig":  clientConfig,
	//})

	// Register instance：RegisterInstance
	port, _ := strconv.Atoi(conf.Conf.Server.K8s.Port)
	success, err := NamingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          conf.Conf.Server.K8s.Ip,
		Port:        uint64(port),
		ServiceName: conf.Conf.Server.Servicename,
		ClusterName: conf.Conf.Nacos.Clustername,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		GroupName:   conf.Conf.Nacos.Groupname,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": conf.Conf.Nacos.Idc},
	})
	if !success {
		log.Fatalln("nacos服务注册失败", err)
	}
	log.Println("nacos服务注册成功：", conf.Conf.Server.K8s.Ip+":"+conf.Conf.Server.K8s.Port)
}

func DeregisterInstance() {
	port, _ := strconv.Atoi(conf.Conf.Server.K8s.Port)
	NamingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          conf.Conf.Server.K8s.Ip,
		Port:        uint64(port),
		Cluster:     conf.Conf.Nacos.Clustername,
		ServiceName: conf.Conf.Server.Servicename,
		GroupName:   conf.Conf.Nacos.Groupname,
		Ephemeral:   true,
	})
}

// SelectOneHealthyInstance
func SelectOneHealthyInstance(serviceName string, groupName string, clusterNames []string) (instanceAddr string, instance *model.Instance, err error) {
	instance, err = NamingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: serviceName,
		GroupName:   groupName,    // default value is DEFAULT_GROUP
		Clusters:    clusterNames, // default value is DEFAULT
	})
	if instance == nil || err != nil {
		return
	}

	port := strconv.FormatInt(int64(instance.Port), 10)
	instanceAddr = instance.Ip + ":" + port
	return
}
