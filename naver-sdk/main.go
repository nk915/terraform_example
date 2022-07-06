package main

import (
	"log"

	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/services/server"
)

func main() {

	//apiKeys := ncloud.Keys()
	apiKeys := ncloud.APIKey{
		AccessKey: "",
		SecretKey: "",
	}

	client := server.NewAPIClient(server.NewConfiguration(&apiKeys))

	//	// Create server instance
	req := server.CreateServerInstancesRequest{
		//ServerImageProductCode:     ncloud.String("SPSW0LINUX000031"),
		ServerImageProductCode: ncloud.String("SPSW0LINUX000046"),
		//ServerProductCode:          ncloud.String("SPSVRSTAND000049"),
		ServerProductCode:          ncloud.String("SPSVRHICPU000001"),
		UserData:                   ncloud.String("#!/bin/sh\nyum -y install httpd"),
		IsProtectServerTermination: ncloud.Bool(false),
		ServerCreateCount:          ncloud.Int32(1),
	}

	if r, err := client.V2Api.CreateServerInstances(&req); err != nil {
		log.Println(err)
	} else {
		sList := r.ServerInstanceList
		log.Println(ncloud.StringValue(r.RequestId))
		log.Println(ncloud.StringValue(sList[0].ServerInstanceNo))
		log.Println(ncloud.StringValue(sList[0].ServerName))
	}

	// Delete server instance
	instanceList := []string{"instanceNum"}
	sreq := server.StopServerInstancesRequest{
		ServerInstanceNoList: ncloud.StringList(instanceList),
	}

	if r, err := client.V2Api.StopServerInstances(&sreq); err != nil {
		log.Print(err)
	} else {
		sList := r.ServerInstanceList
		log.Println(sList)
		log.Println(ncloud.String(*sList[0].ServerImageName))
		log.Println(ncloud.String(*sList[0].ServerInstanceNo))
		//log.Println(ncloud.StringInterfaceList(sList))
	}

	treq := server.TerminateServerInstancesRequest{
		ServerInstanceNoList: ncloud.StringList(instanceList),
	}
	if r, err := client.V2Api.TerminateServerInstances(&treq); err != nil {
		log.Print(err)
	} else {
		sList := r.ServerInstanceList
		log.Println(sList)
		log.Println(ncloud.String(*sList[0].ServerImageName))
		log.Println(ncloud.String(*sList[0].ServerInstanceNo))
		//log.Println(ncloud.StringInterfaceList(sList))
	}
}
