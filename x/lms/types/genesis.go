package types

func ValidateGenesis(data *GenesisState) error {
	return nil
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}
