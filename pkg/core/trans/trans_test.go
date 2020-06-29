package trans

import (
	"github.com/chcp/bsn-sdk-go/pkg/core/entity/msp"
	"github.com/chcp/bsn-sdk-go/pkg/util/keystore"
	"github.com/chcp/bsn-sdk-go/pkg/util/userstore"
	"encoding/base64"
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"

	pb "github.com/chcp/bsn-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
)

func TestGetRequestData(t *testing.T) {

	ks, _ := keystore.NewFileBasedKeyStore(nil, "./test/msp/keystore", false)
	us := userstore.NewUserStore("./test/msp")

	user := &msp.UserData{
		UserName: "sdktest",
		AppCode:  "app0006202004071529586812466",
	}

	us.Load(user)

	keystore.LoadKey(user, ks)

	var args [][]byte
	args = append(args, []byte("{\"baseKey\":\"test20200409\",\"baseValue\":\"this is string \"}"))

	request := &TransRequest{
		ChannelId:    "app0006202004071529586812466",
		ChaincodeId:  "cc_app0006202004071529586812466_00",
		Fcn:          "set",
		Args:         args,
		TransientMap: make(map[string][]byte),
	}

	data, txId, err := CreateRequest(user, request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Data:", data)
	fmt.Println("TxId:", txId)

}

func TestParseRequest(t *testing.T) {
	//data := "Cq0KCsMJCpgBCAMaDAiaw7H0BRCorfLqAiIcYXBwMDAwNjIwMjAwNDA3MTUyOTU4NjgxMjQ2NipANTgyZGFhMzg1NmQzMzBlN2UwMzUxYWEyNDZiZjFkNGY0YmE5YjBlYWJiYmExMGRjYmEzZmQyNGZhZDVmNjhlYjomEiQSImNjX2FwcDAwMDYyMDIwMDQwNzE1Mjk1ODY4MTI0NjZfMDASpQgKiAgKCkhvbmd6YW9NU1AS+QctLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS0KTUlJQ3dqQ0NBbWlnQXdJQkFnSVVFUTZXNFJaWXp3RjFQK09sT0o3NURPMytXbm93Q2dZSUtvWkl6ajBFQXdJdwpVakVMTUFrR0ExVUVCaE1DUTA0eEVEQU9CZ05WQkFnVEIwSmxhV3BwYm1jeEREQUtCZ05WQkFvVEEwSlRUakVQCk1BMEdBMVVFQ3hNR1kyeHBaVzUwTVJJd0VBWURWUVFERXdsaWMyNXliMjkwWTJFd0lCY05NakF3TkRBM01UQXoKTWpBd1doZ1BNakEzTlRBek1qVXhNRE0zTURCYU1Hd3hPekFOQmdOVkJBc1RCbU5zYVdWdWREQU9CZ05WQkFzVApCMmh2Ym1kNllXOHdEZ1lEVlFRTEV3ZGljMjVuWVhSbE1Bb0dBMVVFQ3hNRFkyOXRNUzB3S3dZRFZRUUREQ1J6ClpHdDBaWE4wUUdGd2NEQXdNRFl5TURJd01EUXdOekUxTWprMU9EWTRNVEkwTmpZd1dUQVRCZ2NxaGtqT1BRSUIKQmdncWhrak9QUU1CQndOQ0FBUmpMRVJiZjhycmFwYzUxQ1pKRjBpcFE1V3NENFd6TUNpcGhQdDNGT2tZVkJwawpKU2xnak44a0MwVTEzcnI3eUhJMks5Mkxwa1ZycCtFVGNVN2xmQkFIbzRIL01JSDhNQTRHQTFVZER3RUIvd1FFCkF3SUhnREFNQmdOVkhSTUJBZjhFQWpBQU1CMEdBMVVkRGdRV0JCUUJQdzU1NEVnckN3U2NnbC9TUFNJWWFRNGwKS3pBZkJnTlZIU01FR0RBV2dCUUNmUFhrZWlWYXlCa2FHQ1ZhSzcvY0ZEdTJSRENCbXdZSUtnTUVCUVlIQ0FFRQpnWTU3SW1GMGRISnpJanA3SW1obUxrRm1abWxzYVdGMGFXOXVJam9pYUc5dVozcGhieTVpYzI1bllYUmxMbU52CmJTSXNJbWhtTGtWdWNtOXNiRzFsYm5SSlJDSTZJbk5rYTNSbGMzUkFZWEJ3TURBd05qSXdNakF3TkRBM01UVXkKT1RVNE5qZ3hNalEyTmlJc0ltaG1MbFI1Y0dVaU9pSmpiR2xsYm5RaUxDSnliMnhsSWpvaVkyeHBaVzUwSW4xOQpNQW9HQ0NxR1NNNDlCQU1DQTBnQU1FVUNJUURLREI2Vm8ycWExem1qZGdEZGpJY1hvMHQvZzhVWStOa2xoR2pDClhLb3pyZ0lnS2JJaFNKQ0w5NHJ1T3NPU21MRTAwYnp5R3NMZlNDeGJSQko3ZDl2SGR1MD0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQoSGPoAtTXSeLzpmBV6PzlihJLZoVwzWN627BJlCmMKYQgBEiQSImNjX2FwcDAwMDYyMDIwMDQwNzE1Mjk1ODY4MTI0NjZfMDAaNwoDc2V0CjB7ImJhc2VLZXkiOiJ0ZXN0IiwiYmFzZVZhbHVlIjoidGhpcyBpcyBzdHJpbmcgIn0SRzBFAiEAwQEZy/3VW5M+/eIn0Srn9Uu8MRf3elDdcgTozzPSdJgCIHB1YdPOh6rsI5jgcGvnH5TNFkqty3TWl4S2TbXfL6Vw"

	data :="CsMKCtIJCpcBCAMaCwijkIb1BRCAw4VmIhxhcHAwMDAxMjAyMDA0MTYxMDIwMTUyOTE4NDUxKkBiNTY4NzAwYWQ0ZDJmMDNjOTM5OTdkMWRiMjVmNzU2ZDhmMTAyYTY2OGJlN2RlYTUwNjU3MzEyMjRjMGNiZGMyOiYSJBIiY2NfYXBwMDAwMTIwMjAwNDE2MTAxNzE0MTIzMzkyMF8wMBK1CAqYCAoLT3JnYk5vZGVNU1ASiAgtLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS1NSUlDMlRDQ0FuK2dBd0lCQWdJVUlsR0JUOG8rRWRIalF3U0QyK3gycEI1aE5WWXdDZ1lJS29aSXpqMEVBd0l3VGpFTE1Ba0dBMVVFQmhNQ1EwNHhFREFPQmdOVkJBZ1RCMEpsYVdwcGJtY3hEREFLQmdOVkJBb1RBMEpUVGpFUE1BMEdBMVVFQ3hNR1kyeHBaVzUwTVE0d0RBWURWUVFERXdWaWMyNWpZVEFnRncweU1EQTBNak14TVRReU1EQmFHQTh5TVRBd01ETXlNVEV4TURRd01Gb3dmekVMTUFrR0ExVUVCaE1DUTA0eFBEQU5CZ05WQkFzVEJtTnNhV1Z1ZERBUEJnTlZCQXNUQ0c5eVoySnViMlJsTUE0R0ExVUVDeE1IWW5OdVltRnpaVEFLQmdOVkJBc1RBMk52YlRFeU1EQUdBMVVFQXd3cGRHVnpkREF3TVRBeE1EQXlRR0Z3Y0RBd01ERXlNREl3TURReE5qRXdNakF4TlRJNU1UZzBOVEV3V1RBVEJnY3Foa2pPUFFJQkJnZ3Foa2pPUFFNQkJ3TkNBQVN5VlhlT2tVeWVBSlBKUGp5TVFObHBSR0lObVp4OFZLRit6aGpkR0JPT1oxaGpQL0N4UzR2ME5vd1AybTBsMnFiN3p2RjFiVWFiNCtkM3VUQWF1RXE3bzRJQkJqQ0NBUUl3RGdZRFZSMFBBUUgvQkFRREFnZUFNQXdHQTFVZEV3RUIvd1FDTUFBd0hRWURWUjBPQkJZRUZDbUpEZFVlR3FhK2N4YTk0dE8zL2NDT0MvNFhNQjhHQTFVZEl3UVlNQmFBRkFjSTRIK2tJczh2bjk0WllZcGtyZCs1bGRNS01JR2hCZ2dxQXdRRkJnY0lBUVNCbEhzaVlYUjBjbk1pT25zaWFHWXVRV1ptYVd4cFlYUnBiMjRpT2lKdmNtZGlibTlrWlM1aWMyNWlZWE5sTG1OdmJTSXNJbWhtTGtWdWNtOXNiRzFsYm5SSlJDSTZJblJsYzNRd01ERXdNVEF3TWtCaGNIQXdNREF4TWpBeU1EQTBNVFl4TURJd01UVXlPVEU0TkRVeElpd2lhR1l1Vkhsd1pTSTZJbU5zYVdWdWRDSXNJbkp2YkdVaU9pSmpiR2xsYm5RaWZYMHdDZ1lJS29aSXpqMEVBd0lEU0FBd1JRSWhBTnkwa1BOM0x2UC9YK0dRenUzckwyVDNpVTljNC8vUlcwa25iL2tEbDFEWkFpQng5MXQvT0JsajQrWitjeHFaL2dMUzhlNUxlNzZQb1p2SEJuVi9kQ1hoNXc9PS0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0SGGC0ILs4UdnUesuTPb5wOZv2yS2jOvAdTxJsCmoKaAgEEiQSImNjX2FwcDAwMDEyMDIwMDQxNjEwMTcxNDEyMzM5MjBfMDAaPgoDc2V0Cjd7ImJhc2VLZXkiOiJ0ZXN0MjAyMDA0OCIsImJhc2VWYWx1ZSI6InRoaXMgaXMgc3RyaW5nICJ9EkcwRQIhAKEhHIxE0wB6OwuGaA/ltPYyUu8ZGjH34vrRnHkfVV3qAiBeRX/SB0iyLwwCqrML4tnafu0R5bifIPenwj7uPFh7Rg=="
	ParseRequest(data)

}

func TestParseRequest2(t *testing.T) {

	chaincodeId := "abc"

	chain := &pb.ChaincodeID{Name: chaincodeId}

	bttes, _ := proto.Marshal(chain)

	fmt.Println(base64.StdEncoding.EncodeToString(bttes))

}
