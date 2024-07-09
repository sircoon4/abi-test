package vm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

func Run(input []byte) ([]byte, error) {
	action, err := extractActionFromSerializedPayload(input)
	if err != nil {
		return nil, err
	}

	actionType, ok := action.Get("type_id").(string)
	if !ok {
		return nil, fmt.Errorf("error while getting type_id")
	}
	actionValues, ok := action.Get("values").(*bencodextype.Dictionary)
	if !ok {
		return nil, fmt.Errorf("error while getting values")
	}

	var abi []byte
	switch actionType {
	case "hack_and_slash22":
		abi, err = convertToHackAndSlashEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "grinding2":
		abi, err = convertToGrindingEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "combination_equipment17":
		abi, err = convertToCombinationEquipmentEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "rapid_combination10":
		abi, err = convertToRapidCombinationEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "hack_and_slash_sweep10":
		abi, err = convertToHackAndSlashSweepEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "transfer_asset5":
		abi, err = convertToTransferAssetEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "claim_items":
		abi, err = convertToClaimItemsEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported action type: %s", actionType)
	}

	return common.CopyBytes(abi), nil
}

func RunSimple(input []byte) ([]byte, error) {
	abi, err := convertToSimpleEthAbi()
	if err != nil {
		return nil, err
	}
	return common.CopyBytes(abi), nil
}
