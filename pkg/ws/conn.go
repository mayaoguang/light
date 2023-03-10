package ws

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

// GetConnById 获取指定连接.
func (c *SocketConn) GetConnById(connId ConnId) (*SocketConn, error) {
	sockets.clientMu.RLock()
	defer sockets.clientMu.RUnlock()
	if v, ok := sockets.Clients[connId]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("connId not exist")
}

// Close 关闭连接.
func (c *SocketConn) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if err := c.Conn.Close(); err != nil {
		return
	}
	close(c.closeCh)
	c.delUser()
	// 退出群组
	for groupId := range c.Groups {
		NewGroup(groupId).Exit(c.ConnId)
	}
	sockets.clientMu.Lock()
	defer sockets.clientMu.Unlock()
	delete(sockets.Clients, c.ConnId)
	close(c.sendCh)
	return
}

// JoinGroup 加入组.
func (c *SocketConn) JoinGroup(groupId GroupId) error {
	if !new(SocketGroup).Exist(groupId) {
		return fmt.Errorf("groupId not exist")
	}
	NewGroup(groupId).Join(c.ConnId)
	c.mu.Lock()
	c.Groups[groupId] = struct{}{}
	c.mu.Unlock()
	return nil
}

// ExitGroup 退出组.
func (c *SocketConn) ExitGroup(groupId GroupId) {
	NewGroup(groupId).Exit(c.ConnId)
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.Groups, groupId)
	return
}

// SendMsg 发送消息.
func (c *SocketConn) SendMsg(msg interface{}) {
	b, _ := json.Marshal(msg)
	c.sendCh <- b
}

// production 生产.
func (c *SocketConn) production() {
	for {
		select {
		case data := <-c.sendCh:
			if err := c.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
			}
		case <-c.closeCh:
			return
		}
	}
}

// consumer 消费.
func (c *SocketConn) consumer(handle func(*SocketConn, []byte)) {
	for {
		// 1分钟收不到心跳自动断开
		_ = c.Conn.SetReadDeadline(time.Now().Add(time.Minute))
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			// 预防客户端未发送消息主动断开连接
			select {
			case <-c.closeCh:
				return
			case <-time.After(1 * time.Second):
				c.Close()
				continue
			}
		}
		_ = c.Conn.SetReadDeadline(time.Time{})
		handle(c, data)
	}
}
