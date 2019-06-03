// Package xobj 是为了创建一个统一的对象存储接口，以便切换存储服务之后可以不改业务代码
// 只需要修改配置内容和创建部分变量即可
package xobj

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	cos "github.com/tencentyun/cos-go-sdk-v5"
)

// cosClient 封装的 cos client
type cosClient struct {
	config Config
	client *cos.Client
}

// New 新建 cos 客户端
func newCosClient(bucket string, config Config) Client {
	u, _ := url.Parse(fmt.Sprintf("http://%s-%s.cos.%s.myqcloud.com",
		bucket, config.AppID, config.Region))
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.SecretID,
			SecretKey: config.SecretKey,
		},
	})
	return &cosClient{
		config: config,
		client: c,
	}
}

// Get 获取 cos 对象
func (c *cosClient) Get(key string) ([]byte, error) {
	resp, err := c.client.Object.Get(context.Background(), key, nil)
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return bs, nil
}

// Put 写文件
func (c *cosClient) Put(key string, f io.Reader) error {
	opt := &cos.ObjectPutOptions{}
	_, err := c.client.Object.Put(context.Background(), key, f, opt)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除文件
func (c *cosClient) Delete(key string) error {
	_, err := c.client.Object.Delete(context.Background(), key)
	if err != nil {
		return err
	}
	return nil
}