package main

import (
	"encoding/base64"
	"fmt"

	"github.com/sircoon4/abi-test/vm/libplanet"
	"github.com/sircoon4/bencodex-go"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

func main() {
	serializedPayloadHackAndSlash := "ZDE6UzcwOjBEAiBREB4k6V+8VAXxIe9I1s565xeRyEhTcvft0QtkQavdRAIgF/7gAwEK1hCXBI3vStWHsInR56Pjbpgr77ZGNfYJWJ8xOmFsZHU3OnR5cGVfaWR1MTY6aGFja19hbmRfc2xhc2gyMnU2OnZhbHVlc2R1MTI6YXBTdG9uZUNvdW50dTE6MHUxMzphdmF0YXJBZGRyZXNzMjA63+Olmd3JqvwM5CMsow1BKkN5eE91ODpjb3N0dW1lc2xldTEwOmVxdWlwbWVudHNsMTY6k0hvMSQP3ESqxHe5FvP/ZDE2OnESdTcOEDhHjXX+a8k/kgMxNjputrpDman3Q5V424Y2EJutMTY6vGXnkjiTjk6uEDh/1UnKlTE2OoCmFqSy2FdGlMeyRVDTd0sxNjqAhR6pZsrMRoS+9euOX7TnMTY6yetsy/AeaEm0XJX2GX1KBWV1NTpmb29kc2xldTI6aWQxNjpVMrOqojpcTo+szFsaRu5tdTE6cmxsdTE6MHU1OjMwMDAxZWV1NzpzdGFnZUlkdTM6MjEwdTE0OnRvdGFsUGxheUNvdW50dTE6MXU3OndvcmxkSWR1MTo1ZWVlMTpnMzI6RYIlDQ2jOwZ3moR10oPV3SEMaDubmZ100D+sT1j6a84xOmxpNGUxOm1sZHUxMzpkZWNpbWFsUGxhY2VzMToSdTc6bWludGVyc251Njp0aWNrZXJ1NDpNZWFkZWkxMDAwMDAwMDAwMDAwMDAwMDAwZWUxOm5pMzExNmUxOnA2NToE3wbqFQmA1Y+s/9KKFfFC0scchb6fM0FbN2zIUcbvextqIZ+RGm9zXti+cBRBMhv+/ghXR6k08lAxOUH2iJvMYzE6czIwOmtPF4trn3OiAUK7/onUOjgBB4aDMTp0dTI3OjIwMjQtMDctMDVUMDk6MjM6MjAuNTE3Njc2WjE6dWxlZQ=="
	data, err := base64.StdEncoding.DecodeString(serializedPayloadHackAndSlash)
	if err != nil {
		fmt.Println(err)
		return
	}
	value, err := bencodex.Decode(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	dict, ok := value.(*bencodextype.Dictionary)
	if !ok {
		fmt.Println("not a dictionary")
		return
	}
	abi, err := libplanet.ConvertToTransactionEthAbi(dict)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(abi)
}
