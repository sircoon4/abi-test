package vm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sircoon4/abi-test/vm/actions"
	"github.com/sircoon4/bencodex-go/bencodextype"
)

func Run(input []byte) ([]byte, error) {
	action, err := actions.ExtractActionFromSerializedPayload(input)
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
		abi, err = actions.ConvertToHackAndSlashEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "grinding2":
		abi, err = actions.ConvertToGrindingEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "combination_equipment17":
		abi, err = actions.ConvertToCombinationEquipmentEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "rapid_combination10":
		abi, err = actions.ConvertToRapidCombinationEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "hack_and_slash_sweep10":
		abi, err = actions.ConvertToHackAndSlashSweepEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "transfer_asset5":
		abi, err = actions.ConvertToTransferAssetEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "claim_items":
		abi, err = actions.ConvertToClaimItemsEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	case "daily_reward7":
		abi, err = actions.ConvertToDailyRewardEthAbi(actionValues)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported_%s", actionType)
	}

	return common.CopyBytes(abi), nil
}
