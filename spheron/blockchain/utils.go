package blockchain

import (
	"encoding/json"
	"fmt"

	"github.com/akash-network/provider/spheron/blockchain/gen/OrderMatching"
	"github.com/akash-network/provider/spheron/entities"
	"github.com/ethereum/go-ethereum/common"
)

func stringSliceToAddressSlice(stringSlice []string) []common.Address {
	addressSlice := make([]common.Address, len(stringSlice))
	for i, str := range stringSlice {
		addressSlice[i] = common.HexToAddress(str) // Assuming common.Address has a constructor that accepts a string
	}
	return addressSlice
}

func getOrderSpec(spec entities.DeploymentSpec) OrderMatching.OrderMatchingSpecs {

	resourcesString, err := json.Marshal(spec.Resources)
	if err != nil {
		fmt.Println("Error marshalling object:", err)
		return OrderMatching.OrderMatchingSpecs{}
	}

	return OrderMatching.OrderMatchingSpecs{
		Resources: string(resourcesString),
		PlacementRequirement: OrderMatching.OrderMatchingPlacementRequirement{
			ProviderWallets: StringsToAddresses(spec.PlacementsRequirement.ProviderWallets),
			Attributes:      spec.PlacementsRequirement.Attributes.ToStringSlice(),
		},
	}
}

// StringToAddress converts a string to common.Address
func StringToAddress(str string) common.Address {
	return common.HexToAddress(str)
}

// AddressToString converts a common.Address to string
func AddressToString(address common.Address) string {
	return address.Hex()
}

// StringsToAddresses converts a slice of strings to a slice of common.Address
func StringsToAddresses(strs []string) []common.Address {
	addresses := make([]common.Address, len(strs))
	for i, str := range strs {
		addresses[i] = StringToAddress(str)
	}
	return addresses
}

// AddressesToStrings converts a slice of common.Address to a slice of strings
func AddressesToStrings(addresses []common.Address) []string {
	strs := make([]string, len(addresses))
	for i, address := range addresses {
		strs[i] = AddressToString(address)
	}
	return strs
}