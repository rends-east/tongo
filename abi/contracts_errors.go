package abi

// Code autogenerated. DO NOT EDIT.

import ()

var contractErrors = map[ContractInterface]map[int32]string{
	BidaskPool: {83: "Insufficient gas for operation"},
	CoffeeCrossDex: {100: "Invalid payload to resent",
		101: "Dedust payload is empty",
		102: "ProxyTON payload is empty",
		200: "Not an admin",
	},
	CoffeeFactory: {250: "Unexpected",
		252: "Wrong sender",
		254: "Not an admin",
		255: "Wrong chain",
		256: "Wrong assets",
		257: "Same assets",
		258: "Wrong signature",
		267: "Invalid asset type",
	},
	CoffeeMevProtector: {100: "Empty tx to resend",
		101: "Dedust payload is empty",
		200: "Invalid jetton forward payload",
		201: "Invalid native payload",
		300: "Not an admin",
	},
	CoffeePool: {250: "Unexpected",
		252: "Wrong sender",
		255: "Wrong chain",
		256: "Wrong assets",
		258: "Wrong signature",
		262: "AMM failed",
		263: "Invalid AMM settings",
		264: "Invalid liquidity",
		265: "Invalid fee",
		266: "Invalid condition",
		267: "Invalid asset type",
		268: "Deadline exceeded",
		269: "Slippage tolerated",
		270: "Pool already exists",
		271: "Reserves ratio failed",
		272: "Already active",
		273: "Not active",
	},
	CoffeeStakingItem: {99: "Not initialized",
		100: "Already initialized",
		500: "Too early to unlock",
		501: "Too late to extend",
		502: "Zero points",
	},
	CoffeeStakingMaster: {81: "Unknown jetton wallet",
		99:  "Not initialized",
		100: "Already initialized",
		110: "Mismatching duration",
		111: "Zero reward rate",
		200: "Token not distributed",
	},
	CoffeeStakingVault: {80: "Invalid amount",
		82:  "Invalid dictionary value",
		99:  "Not initialized",
		100: "Already initialized",
		403: "Unauthorized access",
		503: "Wrong period",
		504: "Unexpected forward opcode",
	},
	CoffeeVault: {250: "Unexpected",
		252: "Wrong sender",
		253: "Not a factory",
		255: "Wrong chain",
		256: "Wrong assets",
		257: "Same assets",
		258: "Wrong signature",
		259: "Invalid balance",
		267: "Invalid asset type",
	},
	CoffeeVaultJetton: {260: "Invalid jetton master",
		272: "Already active",
		273: "Not active",
	},
	MultisigOrderV2: {104: "Unauthorized init",
		105:   "Already inited",
		106:   "Unauthorized sign",
		107:   "Already approved",
		111:   "Expired",
		112:   "Already executed",
		65535: "Unknown op",
	},
	MultisigV2: {100: "Not enough ton",
		101:  "Unauthorized execute",
		102:  "Signers outdated",
		103:  "Invalid dictionary sequence",
		108:  "Inconsistent data",
		109:  "Invalid threshold",
		110:  "Invalid signers",
		111:  "Expired",
		1007: "Unauthorized proposer",
		1008: "Invalid new order",
	},
	NftAuctionGetgemsV3: {403: "Unauthorized address",
		1000:  "Low bid",
		1001:  "Already initiated",
		1002:  "Invalid op, expect ownership_assigned",
		1003:  "Invalid op, expect ownership_assigned",
		1005:  "Already executed",
		1006:  "Already activated",
		1008:  "Not enough TONs",
		1009:  "Already has bid",
		1010:  "Auction in progress",
		65535: "Unknown operation",
	},
	NftAuctionGetgemsV4: {403: "Unauthorized address",
		1000:  "Low bid",
		1001:  "Already initiated",
		1002:  "Invalid op, expect ownership_assigned",
		1003:  "Invalid op, expect ownership_assigned",
		1005:  "Already executed",
		1006:  "Already activated",
		1008:  "Not enough TONs",
		1009:  "Already has bid",
		1010:  "Auction in progress",
		1014:  "Auction is in a broken state: auction duration is too long",
		1015:  "Auction has not been activated yet",
		1016:  "Auction is in a broken state",
		1017:  "Wrong currency",
		65535: "Unknown operation",
	},
	NftSaleGetgemsV2: {404: "Already sold or cancelled",
		450:   "Not enough TONs for sale",
		457:   "Not enough TONs for cancellation",
		458:   "Unauthorized cancellation",
		500:   "Unauthorized initiation",
		501:   "Invalid op, expected ownership_assigned",
		65535: "Unknown operation",
	},
	NftSaleGetgemsV3: {404: "Already sold or cancelled",
		405:   "Mode 32 not allowed",
		406:   "Can't execute command at this time",
		408:   "Inconsistent price and fees",
		409:   "Inconsistent price and fees",
		410:   "Inconsistent price and fees",
		450:   "Not enough TONs for sale",
		457:   "Not enough TONs for cancellation",
		458:   "Unauthorized cancellation",
		500:   "Unauthorized initiation",
		501:   "Invalid op, expected ownership_assigned",
		65535: "Unknown operation",
	},
	NftSaleGetgemsV4: {404: "Already sold or cancelled",
		405:   "Mode 32 not allowed",
		406:   "Can't execute command at this time",
		408:   "Inconsistent price and fees",
		409:   "Inconsistent price and fees",
		410:   "Inconsistent price and fees",
		450:   "Not enough TONs for sale",
		451:   "This sale is broken: royalty, fee, or profit is less than zero",
		457:   "Not enough TONs for cancellation",
		458:   "Unauthorized cancellation",
		459:   "This sale is only for a jetton",
		500:   "Unauthorized initiation",
		501:   "Invalid op, expected ownership_assigned",
		4501:  "Not enough funds for sale",
		65535: "Unknown operation",
	},
	SubscriptionV2: {460: "Not enough funds",
		461: "Too early request",
		462: "Too many requests",
		463: "Sender is not wallet",
		464: "Not initialized",
		465: "Initialized",
		466: "Sender is not owner",
		467: "Sender is not beneficiary",
		468: "Invalid charging date",
		469: "Invalid reduce payment amount",
		470: "Subscription is not activated",
		471: "Wallet address is withdraw address",
	},
	WalletV4R2: {33: "Invalid seqno",
		34: "Invalid subwallet ID",
		35: "Invalid signature",
		36: "Expired message",
		39: "Plugin not found",
		80: "Not enough ton",
	},
}

var defaultExitCodes = map[int32]string{
	0:   "Ok",
	1:   "Ok",
	2:   "Stack underflow",
	3:   "Stack overflow",
	4:   "Integer overflow or division by zero",
	5:   "Integer out of expected range",
	6:   "Invalid opcode",
	7:   "Type check error",
	8:   "Cell overflow",
	9:   "Cell underflow",
	10:  "Dictionary error",
	11:  "Unknown error",
	12:  "Impossible situation error",
	13:  "Out of gas error",
	-14: "Out of gas error",
}

func GetContractError(interfaces []ContractInterface, code int32) *string {
	for _, i := range interfaces {
		if errors, ok := contractErrors[i]; ok {
			if msg, ok := errors[code]; ok {
				return &msg
			}
		}
	}
	if msg, ok := defaultExitCodes[code]; ok {
		return &msg
	}
	return nil
}
