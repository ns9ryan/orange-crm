package utils

import (
	"fmt"
	"github.com/duanhf2012/origin/v2/cluster"
	"go.uber.org/zap"
	"orange-crm/global"
	"os"
)

func SendGameClient(RpcFunc string, args interface{}) {
	// 如果是独立运行，则监听退出信号并优雅关闭
	if _, ok := os.LookupEnv("RUN_AS_STANDALONE"); ok {
		return
	}

	// 检测nodeid是否联上本机
	if cluster.GetCluster().IsNodeConnected(global.ToNodeId) == false || cluster.GetCluster().IsNodeRetire(global.ToNodeId) == true {
		global.LOG.Error("SendGameClient Error：", zap.String("nodeId", global.ToNodeId))
		return
	}

	serviceMethod := fmt.Sprintf("%s.%s", global.ToServiceName, RpcFunc)
	//在某些应用场景下不需要数据返回可以使用Go，它是不阻塞的,只需要填入输入参数
	err := global.CrmModule.GoNode(global.ToNodeId, serviceMethod, &args)
	if err != nil {
		global.LOG.Error("Go Origin error ：", zap.String("nodeId", global.ToNodeId), zap.Any("err", err))
	}

	global.LOG.Info("已发送Origin通知：", zap.String("nodeId", global.ToNodeId), zap.String("serviceMethod", serviceMethod), zap.Any("args", args))
}
