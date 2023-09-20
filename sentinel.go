package sentinel

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"log"
)

func InitSentinel() {
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("sentinel init error:%v", err)
	}
	LoadRuleDirect()
}

/**
// 定义限流规则的结构体类型 Rule
type Rule struct {
    // 规则唯一标识符，可选参数
    ID string `json:"id,omitempty"`

    // 表示该规则作用于哪个资源（Resource 名称）
    Resource string `json:"resource"`

    // 令牌发放策略，Sentinel 提供了多种策略，如固定等待时间、滑动窗口等算法
    TokenCalculateStrategy TokenCalculateStrategy `json:"tokenCalculateStrategy"`

    // 限流控制行为，包括丢弃、等待和匀速通过等方式
    ControlBehavior ControlBehavior `json:"controlBehavior"`

    // 阈值设置，当 StatIntervalInMs 设为 1000ms 时表示每秒允许通过的请求数
    Threshold float64 `json:"threshold"`

    // 关联策略，例如熔断器可以通过关联主备节点资源实现更灵活的限流设计
    RelationStrategy RelationStrategy `json:"relationStrategy"`

    // 被关联的资源名称，表示当前限流规则所依赖的资源（可选）
    RefResource string `json:"refResource"`

    // 最大排队等待时间，只有在 ControlBehavior 是 Throttling 时才生效
    // 表示请求在排队后最长等待时间
    MaxQueueingTimeMs uint32 `json:"maxQueueingTimeMs"`

    // 预热期间长度，即系统启动后在此期间内逐渐增加请求通过量直到最大值
    WarmUpPeriodSec uint32 `json:"warmUpPeriodSec"`

    // 预热冷却因子，即达到最大值后，多长时间内将允许的请求通过量逐渐减少至阈值
    WarmUpColdFactor uint32 `json:"warmUpColdFactor"`

    // 统计时间间隔，即滑动窗口大小，指定在多长时间内统计一次请求信息（可选）
    StatIntervalInMs uint32 `json:"statIntervalInMs"`

    // 自适应限流算法相关参数
    LowMemUsageThreshold int64 `json:"lowMemUsageThreshold"`       // 内存使用率低于该阈值时，使用预设阈值控制 QPS
    HighMemUsageThreshold int64 `json:"highMemUsageThreshold"`     // 内存使用率高于该阈值时，使用预设阈值控制 QPS
    MemLowWaterMarkBytes int64 `json:"memLowWaterMarkBytes"`       // 内存利用率低水位线，根据实际内存情况动态调整 QPS
    MemHighWaterMarkBytes int64 `json:"memHighWaterMarkBytes"`     // 内存利用率高水位线，同上
}
*/
func LoadRuleDirect() {
	flow.LoadRules([]*flow.Rule{
		{
			Resource:               "/order/get", // 资源名称: 订单
			TokenCalculateStrategy: flow.Direct,  // 固定数量的令牌
			ControlBehavior:        flow.Reject,  // 当请求超过阈值时，直接拒绝
			Threshold:              10.0,         // 阈值10
			StatIntervalInMs:       1000,         // 统计时间间隔，1000ms即每隔一秒统计一次请求信息
		},
	})
}
