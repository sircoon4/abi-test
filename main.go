package main

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

func main() {
	serializedPayloadList := [][]byte{}
	serializedPayloadHackAndSlash := []byte("ZDE6UzcwOjBEAiBREB4k6V+8VAXxIe9I1s565xeRyEhTcvft0QtkQavdRAIgF/7gAwEK1hCXBI3vStWHsInR56Pjbpgr77ZGNfYJWJ8xOmFsZHU3OnR5cGVfaWR1MTY6aGFja19hbmRfc2xhc2gyMnU2OnZhbHVlc2R1MTI6YXBTdG9uZUNvdW50dTE6MHUxMzphdmF0YXJBZGRyZXNzMjA63+Olmd3JqvwM5CMsow1BKkN5eE91ODpjb3N0dW1lc2xldTEwOmVxdWlwbWVudHNsMTY6k0hvMSQP3ESqxHe5FvP/ZDE2OnESdTcOEDhHjXX+a8k/kgMxNjputrpDman3Q5V424Y2EJutMTY6vGXnkjiTjk6uEDh/1UnKlTE2OoCmFqSy2FdGlMeyRVDTd0sxNjqAhR6pZsrMRoS+9euOX7TnMTY6yetsy/AeaEm0XJX2GX1KBWV1NTpmb29kc2xldTI6aWQxNjpVMrOqojpcTo+szFsaRu5tdTE6cmxsdTE6MHU1OjMwMDAxZWV1NzpzdGFnZUlkdTM6MjEwdTE0OnRvdGFsUGxheUNvdW50dTE6MXU3OndvcmxkSWR1MTo1ZWVlMTpnMzI6RYIlDQ2jOwZ3moR10oPV3SEMaDubmZ100D+sT1j6a84xOmxpNGUxOm1sZHUxMzpkZWNpbWFsUGxhY2VzMToSdTc6bWludGVyc251Njp0aWNrZXJ1NDpNZWFkZWkxMDAwMDAwMDAwMDAwMDAwMDAwZWUxOm5pMzExNmUxOnA2NToE3wbqFQmA1Y+s/9KKFfFC0scchb6fM0FbN2zIUcbvextqIZ+RGm9zXti+cBRBMhv+/ghXR6k08lAxOUH2iJvMYzE6czIwOmtPF4trn3OiAUK7/onUOjgBB4aDMTp0dTI3OjIwMjQtMDctMDVUMDk6MjM6MjAuNTE3Njc2WjE6dWxlZQ==")
	serializedPayloadGrinding := []byte("ZDE6UzcwOjBEAiBNIJpUkfHyhnA3wCHRFt9iuyHqPbKPbEys9MF/RsG1OQIgTB1dU35a5Cr2odqIkNUE6hEZ/PIdrLt5fiidK3YrQsoxOmFsZHU3OnR5cGVfaWR1OTpncmluZGluZzJ1Njp2YWx1ZXNkdTE6YTIwOlE4aopjoX/XOXBM3WCk5oN/7lWIdTE6Y3R1MTplbDE2OutbnGA0koFAu+zs6J5VuAVldTI6aWQxNjpJtSFKgJLXQpKj9T0cXokPZWVlMTpnMzI6RYIlDQ2jOwZ3moR10oPV3SEMaDubmZ100D+sT1j6a84xOmxpMWUxOm1sZHUxMzpkZWNpbWFsUGxhY2VzMToSdTc6bWludGVyc251Njp0aWNrZXJ1NDpNZWFkZWkxMDAwMDAwMDAwMDAwMDAwMDAwZWUxOm5pMjc1ZTE6cDY1OgTcDJ8HG/Fe9jSl94B4LbBfLyUe2gJZmIVaLbjVMXNKh14btogWXuVmcKGhwxuQeDQHjkVD1NXeLyA+X4Jj5Y0vMTpzMjA6ga9iZHvpgtBiwvnBCbOlOdjJy8ExOnR1Mjc6MjAyNC0wNy0wNVQwOToyMzowNi4xNzYyODhaMTp1bGVl")
	serializedPayloadCombinationEquipment := []byte("ZDE6UzcwOjBEAiAAqENcATSUfl/izYJWShs0nTLf5XDqg+n11hTIlSZ7kAIgWzmye7xj2lTPL1bdO+FztKTgSPHYRXBHTFni5OGjDSMxOmFsZHU3OnR5cGVfaWR1MjM6Y29tYmluYXRpb25fZXF1aXBtZW50MTd1Njp2YWx1ZXNkdTE6YTIwOvsmvr466BLcc1rbsNdF7jN1jxN2dTE6aGZ1MTppdTM6MzczdTI6aWQxNjqwfyDSYCl6QqXK30uDXCoBdTE6cHR1MzpwaWRudTE6cnUxOjF1MTpzdTE6MWVlZTE6ZzMyOkWCJQ0NozsGd5qEddKD1d0hDGg7m5mddNA/rE9Y+mvOMTpsaTFlMTptbGR1MTM6ZGVjaW1hbFBsYWNlczE6EnU3Om1pbnRlcnNudTY6dGlja2VydTQ6TWVhZGVpMTAwMDAwMDAwMDAwMDAwMDAwMGVlMTpuaTQ0ZTE6cDY1OgTtl6zT5sIF42uyL9icLWTkFBfLykCjUDRNSPgVzoJ3P0UPXpBZci4nWTcs/pSwsxcfhgUULLgxb5jhWzvRNMUGMTpzMjA695+PB+EmhdKUge+5CMjRa/kIWmIxOnR1Mjc6MjAyNC0wNy0wNVQwOToyMzoxMy41ODcyMTVaMTp1bGVl")

	serializedPayloadList = append(serializedPayloadList, serializedPayloadHackAndSlash)
	serializedPayloadList = append(serializedPayloadList, serializedPayloadGrinding)
	serializedPayloadList = append(serializedPayloadList, serializedPayloadCombinationEquipment)

	action, err := extractActionFromSerializedPayload(serializedPayloadList[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	actionType, ok := action.Get("type_id").(string)
	if !ok {
		fmt.Println("error while getting type_id")
		return
	}
	actionValues, ok := action.Get("values").(*bencodextype.Dictionary)
	if !ok {
		fmt.Println("error while getting values")
		return
	}

	var abi []byte
	switch actionType {
	case "hack_and_slash22":
		abi, err = convertToHackAndSlashEthAbi(actionValues)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "grinding2":
		abi, err = convertToGrindingEthAbi(actionValues)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "combination_equipment17":
		abi, err = convertToCombinationEquipmentEthAbi(actionValues)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Println("not supported action type")
		return
	}
	fmt.Println(abi)
}

func extractActionFromSerializedPayload(serializedPayload []byte) (*bencodextype.Dictionary, error) {
	encoded, err := base64.StdEncoding.DecodeString(string(serializedPayload))
	if err != nil {
		return nil, err
	}
	decoded, err := bencodex.Decode(encoded)
	if err != nil {
		return nil, err
	}
	dict, ok := decoded.(*bencodextype.Dictionary)
	if !ok {
		return nil, fmt.Errorf("error while casting to dictionary")
	}
	action, ok := dict.Get([]byte{0x61}).([]any)[0].(*bencodextype.Dictionary)
	if !ok {
		return nil, fmt.Errorf("error while getting action")
	}
	fmt.Println(action)
	return action, nil
}

type HackAndSlash struct {
	Id             [16]byte       `abi:"id"`
	Costumes       [][16]byte     `abi:"costumes"`
	Equipments     [][16]byte     `abi:"equipments"`
	Foods          [][16]byte     `abi:"foods"`
	R              [][]int64      `abi:"r"`
	WorldId        int64          `abi:"worldId"`
	StageId        int64          `abi:"stageId"`
	StageBuffId    int64          `abi:"stageBuffId"`
	AvatarAddress  common.Address `abi:"avatarAddress"`
	TotalPlayCount int64          `abi:"totalPlayCount"`
	ApStoneCount   int64          `abi:"apStoneCount"`
}

func convertToHackAndSlashEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleHackAndSlash, _ = abi.NewType("tuple", "struct HackAndSlash", []abi.ArgumentMarshaling{
		{Name: "id", Type: "uint8[16]"},
		{Name: "costumes", Type: "uint8[16][]"},
		{Name: "equipments", Type: "uint8[16][]"},
		{Name: "foods", Type: "uint8[16][]"},
		{Name: "r", Type: "int64[][]"},
		{Name: "worldId", Type: "int64"},
		{Name: "stageId", Type: "int64"},
		{Name: "stageBuffId", Type: "int64"},
		{Name: "avatarAddress", Type: "address"},
		{Name: "totalPlayCount", Type: "int64"},
		{Name: "apStoneCount", Type: "int64"},
	})

	var arguments = abi.Arguments{
		abi.Argument{Name: "HackAndSlash", Type: TupleHackAndSlash, Indexed: false},
	}

	idValue, _ := actionValues.Get("id").([]byte)
	costumesList := [][16]byte{}
	for _, costume := range actionValues.Get("costumes").([]any) {
		costumeValue, _ := costume.([]byte)
		costumesList = append(costumesList, [16]byte(costumeValue))
	}
	equipmentsList := [][16]byte{}
	for _, equipment := range actionValues.Get("equipments").([]any) {
		equipmentValue, _ := equipment.([]byte)
		equipmentsList = append(equipmentsList, [16]byte(equipmentValue))
	}
	foodsList := [][16]byte{}
	for _, food := range actionValues.Get("foods").([]any) {
		foodValue, _ := food.([]byte)
		foodsList = append(foodsList, [16]byte(foodValue))
	}
	rList := [][]int64{}
	for _, r := range actionValues.Get("r").([]any) {
		rInfo := []int64{}
		rFirstValue, _ := strconv.Atoi(r.([]any)[0].(string))
		rSecondValue, _ := strconv.Atoi(r.([]any)[1].(string))
		rInfo = append(rInfo, int64(rFirstValue), int64(rSecondValue))
		rList = append(rList, rInfo)
	}
	worldIdValue, _ := strconv.Atoi(actionValues.Get("worldId").(string))
	stageIdValue, _ := strconv.Atoi(actionValues.Get("stageId").(string))
	stageBuffIdValue := -1
	if actionValues.Contains("stageBuffId") {
		if actionValues.Get("stageBuffId") != nil {
			stageBuffIdValue, _ = strconv.Atoi(actionValues.Get("stageBuffId").(string))
		}
	}
	avatarAddressValue := common.BytesToAddress(actionValues.Get("avatarAddress").([]byte))
	totalPlayCountValue, _ := strconv.Atoi(actionValues.Get("totalPlayCount").(string))
	apStoneCountValue, _ := strconv.Atoi(actionValues.Get("apStoneCount").(string))

	result, err := arguments.Pack(HackAndSlash{
		Id:             [16]byte(idValue),
		Costumes:       costumesList,
		Equipments:     equipmentsList,
		Foods:          foodsList,
		R:              rList,
		WorldId:        int64(worldIdValue),
		StageId:        int64(stageIdValue),
		StageBuffId:    int64(stageBuffIdValue),
		AvatarAddress:  avatarAddressValue,
		TotalPlayCount: int64(totalPlayCountValue),
		ApStoneCount:   int64(apStoneCountValue),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

type Grinding struct {
	Id [16]byte       `abi:"id"`
	A  common.Address `abi:"a"`
	E  [][16]byte     `abi:"e"`
	C  bool           `abi:"c"`
}

func convertToGrindingEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleGrinding, _ = abi.NewType("tuple", "struct Grinding", []abi.ArgumentMarshaling{
		{Name: "id", Type: "uint8[16]"},
		{Name: "a", Type: "address"},
		{Name: "e", Type: "uint8[16][]"},
		{Name: "c", Type: "bool"},
	})

	var arguments = abi.Arguments{
		abi.Argument{Name: "Grinding", Type: TupleGrinding, Indexed: false},
	}

	idValue, _ := actionValues.Get("id").([]byte)
	aValue := common.BytesToAddress(actionValues.Get("a").([]byte))
	eList := [][16]byte{}
	for _, e := range actionValues.Get("e").([]any) {
		eValue, _ := e.([]byte)
		eList = append(eList, [16]byte(eValue))
	}
	cValue, _ := actionValues.Get("c").(bool)

	result, err := arguments.Pack(Grinding{
		Id: [16]byte(idValue),
		A:  aValue,
		E:  eList,
		C:  cValue,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

type CombinationEquipment struct {
	Id  [16]byte       `abi:"id"`
	A   common.Address `abi:"a"`
	S   int64          `abi:"s"`
	R   int64          `abi:"r"`
	I   int64          `abi:"i"`
	P   bool           `abi:"p"`
	H   bool           `abi:"h"`
	Pid int64          `abi:"pid"`
}

func convertToCombinationEquipmentEthAbi(actionValues *bencodextype.Dictionary) ([]byte, error) {
	var TupleCombinationEquipment, _ = abi.NewType("tuple", "struct CombinationEquipment", []abi.ArgumentMarshaling{
		{Name: "id", Type: "uint8[16]"},
		{Name: "a", Type: "address"},
		{Name: "s", Type: "int64"},
		{Name: "r", Type: "int64"},
		{Name: "i", Type: "int64"},
		{Name: "p", Type: "bool"},
		{Name: "h", Type: "bool"},
		{Name: "pid", Type: "int64"},
	})

	var arguments = abi.Arguments{
		abi.Argument{Name: "CombinationEquipment", Type: TupleCombinationEquipment, Indexed: false},
	}

	idValue, _ := actionValues.Get("id").([]byte)
	aValue := common.BytesToAddress(actionValues.Get("a").([]byte))
	sValue, _ := strconv.Atoi(actionValues.Get("s").(string))
	rValue, _ := strconv.Atoi(actionValues.Get("r").(string))
	iValue := -1
	if actionValues.Get("i") != nil {
		iValue, _ = strconv.Atoi(actionValues.Get("i").(string))
	}
	pValue, _ := actionValues.Get("p").(bool)
	hValue, _ := actionValues.Get("h").(bool)
	pidValue := -1
	if actionValues.Get("pid") != nil {
		pidValue, _ = strconv.Atoi(actionValues.Get("pid").(string))
	}

	result, err := arguments.Pack(CombinationEquipment{
		Id:  [16]byte(idValue),
		A:   aValue,
		S:   int64(sValue),
		R:   int64(rValue),
		I:   int64(iValue),
		P:   pValue,
		H:   hValue,
		Pid: int64(pidValue),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
