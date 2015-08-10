// Author: zheng-ji.info

package monitor

const (
	STAT_PERSIS_WRITE      = iota //统计持久库写类型
	STAT_PERSIS_READ              //统计持久库读类型
	STAT_INNERCACHE_READ          //统计内部缓读类型
	STAT_CACHECLUSTER_READ        //统计 cachecluser 读类型

	//达到统计的阀值
	NUM_PER_STAT   = 100
	STAT_UNINT_NUM = 4 //统计
)
