package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
	"unsafe"
)

type (
	GetResourceListRequest struct {
		Action            string
		Backend           string
		ResourceType      int      // 资源类型(必填)
		RegionId          int      // 地域ID(必填)
		ZoneId            int      // 可用区ID
		TopOrganizationId int      // 公司ID
		OrganizationId    int      // 项目ID
		ResourceIds       []string // 资源短ID数组
		Offset            int      // 偏移量
		Limit             int      // 请求量
		GenerateType      int      // 传1时ResourceId为长ID, 否则为短ID
		CommonResource    int      // 传1返回公共资源
		Status            int      // 资源状态
		UpdateTime        int      // 更新时间
		VPCId             string
		SubnetId          string
		BusinessId        string
		request_uuid      string
	}
	GetResourceListResponse struct {
		TotalCount    int      // 总数量
		ResourceIds   []string // 短ID数组
		ResourceUUIds []string // 长ID数组
		ZoneIds       []int    // 可用区ID数组
		RegionIds     []int
		Id2UUId       map[string]string       // 短ID对应长ID
		UUId2Id       map[string]string       // 长ID对应短ID
		UUId2Info     map[string]ResourceInfo // 长ID对应资源信息
	}
	ResourceInfo struct {
		Id                string // 短ID
		ResourceId        string // 长ID
		RegionId          int    // 地域ID
		ZoneId            int    // 可用区ID
		ResourceType      int    // 资源类型
		TopOrganizationId int    // 公司ID
		OrganizationId    int    // 项目ID
		Updated           int    // 更新时间
		Created           int    // 创建时间
		Status            int    // 资源状态(1正常2删除10冻结)
		VPCId             string // VPC ID
		SubnetId          string // 子网ID
		BusinessId        string // 业务组ID
	}
	CreateVPNTunnelRequest struct {
		Action       string
		Backend      string
		RegionId     int    `key:"RegionId" required:"true"`     // 地域id 来于api-gateway
		VPNGatewayId string `key:"VPNGatewayId" required:"true"` // VPNGatewayId
		request_uuid string
	}
	ListSubnetRequest struct {
		Action       string
		Backend      string
		RegionId     int    `key:"RegionId" required:"true"`     // 地域id 来于api-gateway
		SubnetworkId string `key:"SubnetworkId" required:"true"` // VPNGatewayId
		VNetId       string `key:"VNetId" required:"true"`
		AccountId    int    `key:"AccountId" required:"true"`
		request_uuid string
	}
	APIResponse struct {
		Action  string `json:"Action"`  // 返回名字
		RetCode int    `json:"RetCode"` // 状态码 0:成功
		Message string `json:"Message"` // 错误消息
	}
	IGetResourceListResponse struct {
		APIResponse
		TotalCount int
		Infos      []ResourceInfo
	}
)

func main() {
	//req := new(GetResourceListRequest)
	//req.Action = "IGetResourceList"
	//req.Backend = "UResource"
	//req.RegionId = 666888
	//req.TopOrganizationId = 200000115
	//req.ResourceType = 145
	//req.ResourceIds = []string{"vpngw-tbjara"}
	//req.OrganizationId = 200000489
	//req.Limit = 10000000
	//req.Offset = 0

	req := new(ListSubnetRequest)
	req.Action = "IListSubnet"
	req.Backend = "UVPCFEGO"
	req.SubnetworkId = "subnet-ah0f24"
	req.VNetId = "uvnet-5p1i44"
	req.AccountId = 200000489
	req.RegionId = 666888
	req.request_uuid = "12e684e9-036b-41e6-9152-d9e75b7b4c33"
	body, err := postrequest(req)
	if err != nil {
		fmt.Println(err)
	}
	//params := map[string]interface{}{
	//	"Backend":      req.Backend,
	//	"Action":       req.Action,
	//	"VPNGatewayId": req.VPNGatewayId,
	//	"RegionId":     req.RegionId,
	//}
	res := new(interface{})
	//err := PrivateAPIRequest(params, res)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(body, res)
	//if err != nil {
	//	fmt.Println(err)
	//}
	fmt.Println(*res)

	vpcResource, err := GetResourceList(&GetResourceListRequest{
		ResourceType: 126,
		//RegionId:          req.RegionId,   //允许不通region下vpc用udpn连接
		TopOrganizationId: 50120027,
		//OrganizationId:    req.OrganizationId,  //允许同一region下的不同vpc
		//ResourceIds: req.VPCIds,
	})
	fmt.Println(vpcResource.ResourceIds)
	fmt.Println(vpcResource.RegionIds)
}
func sendRequest(backend string, action string, req interface{}) ([]byte, error) {
	baseURL := "http://internal.api.pre.ucloudadmin.com"
	url := fmt.Sprintf("%s/?Backend=%s&Action=%s", baseURL, backend, action)
	fmt.Println(url)
	body, _ := json.Marshal(req)
	var req2 GetResourceListRequest
	json.Unmarshal(body, &req2)
	fmt.Println(req2)
	//VPNGatewayCommonLibs.LOG( action).WithFields(log.Fields{"request": string(body), "backend": backend, "url": url}).Debug("http requesting")
	resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	bodyResp, err := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	//VPNGatewayCommonLibs.LOG( action).WithFields(log.Fields{"request": string(body), "response": string(bodyResp), "backend": backend}).Debug("http request finish")
	return bodyResp, nil
}
func postrequest(req interface{}) ([]byte, error) {
	//URL := "http://internal.api.pre.ucloudadmin.com"
	URL := "http://192.168.153.95:5002"
	body, _ := json.Marshal(req)
	request, err := http.NewRequest("POST", URL, bytes.NewReader(body))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	client.Timeout = 10 * time.Second
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println("header is", resp.Header)
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
	return respBytes, nil
}

func PrivateAPIRequest(data map[string]interface{}, result interface{}) error {

	if _, ok := data["request_uuid"]; !ok {
		data["request_uuid"] = "tangshenzheeng"
	}

	//VPNGatewayCommonLibs.LOG( "PrivateAPIRequest").WithFields(log.Fields{"request": data, "backend": data["Backend"], "action": data["Action"]}).Debug("request")

	err := HttpRequest("http://internal.api.ucloud.cn", data, result)
	//err := HttpRequest("http://internal.api.pre.ucloudadmin.com", data, result)

	//VPNGatewayCommonLibs.LOG( "PrivateAPIRequest").WithFields(log.Fields{"response": result}).Debug("response")

	return err
}

func HttpRequest(url string, data map[string]interface{}, result interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	//VPNGatewayCommonLibs.LOG( "HttpRequest").WithFields(log.Fields{"url": url, "request": data}).Debug("request")
	res, err := SendHTTPPostRequest(url, "application/json", bytes.NewBuffer(body), 10)
	//VPNGatewayCommonLibs.LOG( "HttpRequest").WithFields(log.Fields{"response": res}).Debug("response")
	if err != nil {
		return err
	}
	err = json.Unmarshal(res, result)
	if err != nil {
		return err
	}
	return nil
}
func SendHTTPPostRequest(uri string, body_type string, body io.Reader, timeOut uint32) (res []byte, err error) {
	return sendHttpPostRequest(uri, body_type, body, timeOut)
}
func sendHttpPostRequest(url_path string, body_type string, body io.Reader, timeOut uint32) (res []byte, err error) {
	client := newTimeoutHTTPClient(time.Duration(timeOut) * time.Second)
	result, err := client.Post(url_path, body_type, body)
	if err != nil {
		return
	}
	defer result.Body.Close()
	res, err = ioutil.ReadAll(result.Body)
	return
}
func newTimeoutHTTPClient(timeOut time.Duration) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Dial: dialHTTPTimeout(timeOut),
		},
	}
}
func dialHTTPTimeout(timeOut time.Duration) func(net, addr string) (net.Conn, error) {
	return func(network, addr string) (c net.Conn, err error) {
		c, err = net.DialTimeout(network, addr, timeOut)
		if err != nil {
			return
		}
		if timeOut > 0 {
			c.SetDeadline(time.Now().Add(timeOut))
		}
		return
	}
}
func GetResourceList(req *GetResourceListRequest) (*GetResourceListResponse, error) {
	// 拼装请求体
	params := map[string]interface{}{
		"Backend":      "UResource",
		"Action":       "IGetResourceList",
		"ResourceType": req.ResourceType,
		"RegionId":     req.RegionId,
	}
	if req.ZoneId != 0 {
		params["ZoneId"] = req.ZoneId
	}
	if req.TopOrganizationId != 0 {
		params["TopOrganizationId"] = req.TopOrganizationId
	}
	if req.OrganizationId != 0 {
		params["OrganizationId"] = req.OrganizationId
	}
	if len(req.ResourceIds) > 0 {
		params["ResourceId"] = req.ResourceIds
	}
	if req.Offset != 0 {
		params["Offset"] = req.Offset
	} else {
		params["Offset"] = 0
	}
	if req.Limit != 0 {
		params["Limit"] = req.Limit
	} else {
		params["Limit"] = 10000000
	}
	if req.GenerateType != 0 {
		params["GenerateType"] = req.GenerateType
	}
	if req.CommonResource != 0 {
		params["CommonResource"] = req.CommonResource
	}
	if req.Status != 0 {
		params["Status"] = req.Status
	}
	if req.UpdateTime != 0 {
		params["UpdateTime"] = req.UpdateTime
	}
	if req.VPCId != "" {
		params["VPCId"] = req.VPCId
	}
	if req.SubnetId != "" {
		params["SubnetId"] = req.SubnetId
	}
	if req.BusinessId != "" {
		params["BusinessId"] = req.BusinessId
	}

	// 发送请求
	_res := new(IGetResourceListResponse)
	err := PrivateAPIRequest(params, _res)
	if err != nil {
		return nil, err
	}
	if _res.RetCode != 0 {
		return nil, errors.New(_res.Message)
	}

	// 处理返回
	res := &GetResourceListResponse{
		TotalCount:    _res.TotalCount,
		ResourceIds:   make([]string, 0),
		ResourceUUIds: make([]string, 0),
		ZoneIds:       make([]int, 0),
		RegionIds:     make([]int, 0),
		Id2UUId:       make(map[string]string, _res.TotalCount),
		UUId2Id:       make(map[string]string, _res.TotalCount),
		UUId2Info:     make(map[string]ResourceInfo, _res.TotalCount),
	}
	for _, info := range _res.Infos {
		res.ResourceIds = append(res.ResourceIds, info.Id)
		res.ResourceUUIds = append(res.ResourceUUIds, info.ResourceId)
		res.ZoneIds = append(res.ZoneIds, info.ZoneId)
		res.RegionIds = append(res.RegionIds, info.RegionId)
		res.Id2UUId[info.Id] = info.ResourceId
		res.UUId2Id[info.ResourceId] = info.Id
		res.UUId2Info[info.ResourceId] = info
	}

	return res, nil
}
