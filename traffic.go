package musdk

import (
	"time"
)

func (c *Client) SaveTrafficLog(l UserTrafficLog) {
	c.userTFmu.Lock()
	defer c.userTFmu.Unlock()

	v, ok := c.userTraffic[l.UserId]
	if !ok {
		c.userTraffic[l.UserId] = l
		return
	}
	v.U += l.U
	v.D += l.D
	c.userTraffic[l.UserId] = v
	return
}

func (c *Client) submitTrafficLog() error {
	c.userTFmu.Lock()
	defer c.userTFmu.Unlock()

	if len(c.userTraffic) == 0 {
		c.logger.Infof("not traffic log,skip")
		return nil
	}

	logs := make([]UserTrafficLog, len(c.userTraffic))
	i := 0
	for _, v := range c.userTraffic {
		logs[i] = v
		i++
	}

	err := c.UpdateTraffic(logs)
	if err != nil {
		return err
	}
	c.logger.Infof("post traffic log len %d", len(logs))
	c.userTraffic = make(map[int64]UserTrafficLog)
	return nil
}

func (c *Client) UpdateTrafficDaemon() {
	for {
		c.submitTrafficLog()
		time.Sleep(time.Minute)
	}
}
