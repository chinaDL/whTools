package ip

import (
	"fmt"
	"testing"
)

func TestIPRange_Do(t *testing.T) {
	ips := []string{
		"192.168.1.1/30",
		"192.168.1.1/24",
		"192.168.1.1/16",
		"192.168.1.1/8",
		"192.168.1.1/2",
		"192.168.1.1-192.168.255.255-",
		"192.168.0.1-192.168.255.255",
		"192.168.1.1-192.168.255.255",
		"192.168.1.1-255",
		"192.168.1.1-2.255",
		"192.168.1.*",
		"192.168.*.*",
		"192.*.*.*",
		"192.*.1.*",
		"192.168.1.1",
		"192.168.1.1,192.168.1.2",
	}
	for _, v := range ips {
		r, err := ParseIPStr(v)
		if err != nil {
			fmt.Printf("%s 错误: %s\n", v, err.Error())
			continue
		}
		r.Do(func(ipRange *IPRange) bool {
			fmt.Printf("%s 结果: 开始IP: %s 结束IP: %s 总数: %d\n",
				v, ipRange.Start2IP(), ipRange.End2IP(), ipRange.Count())
			return true
		})

		if r.Count() < 10 {
			fmt.Printf("    %+v\n", r.AllToStr())
		}
		if ok, _ := r.Include("192.8.88.*"); ok {
			fmt.Printf("    ----%s 包含 %s \n", v, "192.8.88.*")
		}

	}
}
