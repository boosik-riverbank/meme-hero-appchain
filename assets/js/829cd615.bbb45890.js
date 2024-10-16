"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[1506],{3525:(e,n,r)=>{r.r(n),r.d(n,{assets:()=>d,contentTitle:()=>s,default:()=>h,frontMatter:()=>a,metadata:()=>o,toc:()=>l});var i=r(5893),t=r(1151);const a={sidebar_position:17,title:"Security aggregation"},s="ADR 016: Security aggregation",o={id:"adrs/adr-016-securityaggregation",title:"Security aggregation",description:"Changelog",source:"@site/versioned_docs/version-v4.5.0/adrs/adr-016-securityaggregation.md",sourceDirName:"adrs",slug:"/adrs/adr-016-securityaggregation",permalink:"/interchain-security/v4.5.0/adrs/adr-016-securityaggregation",draft:!1,unlisted:!1,tags:[],version:"v4.5.0",sidebarPosition:17,frontMatter:{sidebar_position:17,title:"Security aggregation"},sidebar:"tutorialSidebar",previous:{title:"Partial Set Security",permalink:"/interchain-security/v4.5.0/adrs/adr-015-partial-set-security"},next:{title:"ICS with Inactive Provider Validators",permalink:"/interchain-security/v4.5.0/adrs/adr-017-allowing-inactive-validators"}},d={},l=[{value:"Changelog",id:"changelog",level:2},{value:"Status",id:"status",level:2},{value:"Context",id:"context",level:2},{value:"Alternative Approaches",id:"alternative-approaches",level:2},{value:"Rewards",id:"rewards",level:3},{value:"Decision",id:"decision",level:2},{value:"Rewards will be sent back to external chains instead of paying rewards for external stakers on Cosmos chain",id:"rewards-will-be-sent-back-to-external-chains-instead-of-paying-rewards-for-external-stakers-on-cosmos-chain",level:3},{value:"Detailed Design",id:"detailed-design",level:2},{value:"Power Mixing",id:"power-mixing",level:3},{value:"Integration with <code>ICS provider</code>",id:"integration-with-ics-provider",level:4},{value:"Integration with <code>ICS consumer</code>",id:"integration-with-ics-consumer",level:4},{value:"Queries",id:"queries",level:3},{value:"Reward Handler",id:"reward-handler",level:3},{value:"Consequences",id:"consequences",level:2},{value:"Positive",id:"positive",level:3},{value:"Negative",id:"negative",level:3},{value:"Neutral",id:"neutral",level:3},{value:"Questions:",id:"questions",level:2},{value:"References",id:"references",level:2}];function c(e){const n={a:"a",code:"code",em:"em",h1:"h1",h2:"h2",h3:"h3",h4:"h4",li:"li",p:"p",pre:"pre",ul:"ul",...(0,t.a)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(n.h1,{id:"adr-016-security-aggregation",children:"ADR 016: Security aggregation"}),"\n",(0,i.jsx)(n.h2,{id:"changelog",children:"Changelog"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"2024-04-24: Initial draft of ADR"}),"\n"]}),"\n",(0,i.jsx)(n.h2,{id:"status",children:"Status"}),"\n",(0,i.jsx)(n.p,{children:"Proposed"}),"\n",(0,i.jsx)(n.h2,{id:"context",children:"Context"}),"\n",(0,i.jsx)(n.p,{children:"Security Aggregation enables staking of tokens from external sources such as Ethereum or Bitcoin to Cosmos blockchains. By integrating Security Aggregation, a Cosmos blockchain can be secured by both native tokens and external tokens (e.g. ETH, BTC)."}),"\n",(0,i.jsx)(n.p,{children:"Security Aggregation consists of the following parts:"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"A mechanism for delegating external tokens to Cosmos validators, such as Babylon or EigenLayer AVS contract."}),"\n",(0,i.jsx)(n.li,{children:"An oracle that tracks how much external stake has been delegated to each Cosmos validator and provides price feeds for external tokens."}),"\n",(0,i.jsx)(n.li,{children:"Power mixing:  a mechanism to combine external and native stake to derive the power of each validator."}),"\n",(0,i.jsx)(n.li,{children:"A reward distribution protocol that enables sending back rewards to the external source."}),"\n"]}),"\n",(0,i.jsx)(n.p,{children:"External staking information is received from an oracle together with price information of related stakes.\nThe CosmosLayer derives validator powers based on external and native staking information and initiates rewarding of external depositors."}),"\n",(0,i.jsxs)(n.p,{children:["This ADR describes the ",(0,i.jsx)(n.em,{children:"Cosmos modules"})," of the solution."]}),"\n",(0,i.jsx)(n.h2,{id:"alternative-approaches",children:"Alternative Approaches"}),"\n",(0,i.jsx)(n.h3,{id:"rewards",children:"Rewards"}),"\n",(0,i.jsx)(n.p,{children:"As an alternative to sending rewards back to the external chains, stakers could be rewarded on the Cosmos chain.\nThis would require a mapping of external addresses to addresses on Cosmos chain for each staker on external source.\nIn addition detailed external staking information such as staking addresses, amount of stakes per staker and validator, etc. have to be provided by the oracle."}),"\n",(0,i.jsx)(n.h2,{id:"decision",children:"Decision"}),"\n",(0,i.jsx)(n.h3,{id:"rewards-will-be-sent-back-to-external-chains-instead-of-paying-rewards-for-external-stakers-on-cosmos-chain",children:"Rewards will be sent back to external chains instead of paying rewards for external stakers on Cosmos chain"}),"\n",(0,i.jsx)(n.p,{children:"Rewards will be sent back to external chains instead of paying rewards for external stakers on Cosmos chain"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"due to amount of additional staking information to be sent and tracked by the oracle"}),"\n",(0,i.jsx)(n.li,{children:"due to the additional complexity of managing external and Cosmos addresses"}),"\n"]}),"\n",(0,i.jsx)(n.h2,{id:"detailed-design",children:"Detailed Design"}),"\n",(0,i.jsxs)(n.p,{children:["The ",(0,i.jsx)(n.code,{children:"Power Mixing"})," feature and ",(0,i.jsx)(n.code,{children:"Reward Distribution"})," protocol are an integral part of the Security Aggregation solution.\nThe ",(0,i.jsx)(n.code,{children:"Power Mixing"})," module provides the capability of deriving validator power based on stake originated from external sources such as Ethereum/Bitcoin and the native staking module.\nThe ",(0,i.jsx)(n.code,{children:"Reward Distribution"})," manages the process of sending rewards to external stakers."]}),"\n",(0,i.jsx)(n.h3,{id:"power-mixing",children:"Power Mixing"}),"\n",(0,i.jsxs)(n.p,{children:["Power Mixing provides the final validator powers based on staking information of the native chain and the external stakes. The information about external staking and related price feeds are received from an oracle.\nOnce the final validator powers are determined the result is submitted to the underlying CometBFT consensus layer by ",(0,i.jsx)(n.a,{href:"https://docs.cometbft.com/v0.38/spec/abci/abci++_app_requirements#updating-the-validator-set",children:"updating"})," the validator set."]}),"\n",(0,i.jsx)(n.p,{children:"Requirements:"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"validator updates are performed on each EndBlock"}),"\n",(0,i.jsx)(n.li,{children:"a validator's power is determined based on its native on-chain stakes and external stakes"}),"\n",(0,i.jsx)(n.li,{children:"price information of staked tokens is used to determine a validator\u2019s power, e.g. price ratio (price of native on-chain token / price of external stake)"}),"\n",(0,i.jsx)(n.li,{children:"price information of native/external tokens are received from an oracle"}),"\n",(0,i.jsx)(n.li,{children:"staking information from external sources received from the oracle"}),"\n",(0,i.jsxs)(n.li,{children:["native staking information are received from the ",(0,i.jsx)(n.code,{children:"Cosmos SDK Staking Module"})]}),"\n",(0,i.jsx)(n.li,{children:"set of validator stakes from oracle always have the current price, full set of validators, and current stakes"}),"\n"]}),"\n",(0,i.jsxs)(n.p,{children:["The ",(0,i.jsx)(n.code,{children:"Power Mixing"})," implementation"]}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:["queries current validators and their powers from ",(0,i.jsx)(n.a,{href:"https://github.com/cosmos/cosmos-sdk/blob/a6f3fbfbeb7ea94bda6369a7163a523e118a123c/x/staking/types/staking.pb.go#L415",children:"x/staking"}),"\nand from oracle (see below)."]}),"\n",(0,i.jsx)(n.li,{children:"calculates power updates by mixing power values of external and internal sources\nFollowing pseudocode snippet shows a possible implementation of how power mixing\nfeature works."}),"\n"]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-golang",children:"// PowerSource is an abstract entity providing validator powers which\n// are used by the mixer. This can be an oracle, staking module or an\n// IBC connected bridge.\ntype PowerSource interface {\n  GetValidatorUpdates() []abci.ValidatorUpdate\n}\n\n// MixPowers calculates power updates by mixing validator powers from different sources\nfunc (k *Keeper) MixPowers(source ...PowerSource) []abci.ValidatorUpdate {\n  var valUpdate []abci.ValidatorUpdate\n  for _, ps := range source {\n    // mix powers from two sets of validator updates an return set of validator updates\n    // with aggregated powers\n    valUpdate = mixPower(valUpdate, ps.GetValidatorUpdates())\n  }\n  return valUpdate\n}\n\nfunc (k *keeper) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {\n  // GetPowerSources (including local staking module)\n  registeredPowerSource := GetPowerSources()\n  return k.MixPowers(registeredPowerSource...)\n}\n"})}),"\n",(0,i.jsxs)(n.h4,{id:"integration-with-ics-provider",children:["Integration with ",(0,i.jsx)(n.code,{children:"ICS provider"})]}),"\n",(0,i.jsxs)(n.p,{children:["The provider module updates the validator set on CometBFT instead of the SDK staking module (x/staking). The provider implementation will intervene in this behavior and ensure that the validator updates are taken from the ",(0,i.jsx)(n.code,{children:"Power Mixing"})," feature."]}),"\n",(0,i.jsxs)(n.p,{children:["External power sources are managed by the provider module. Only registered power sources can provide input to the ",(0,i.jsx)(n.code,{children:"Power Mixing"})," feature.\nPower sources will be assigned a unique identifier which will be used by the oracle, provider module and the power mixing and rewarding feature."]}),"\n",(0,i.jsxs)(n.p,{children:["Updates with the next validator set are sent to consumer chains on each epoch (see ",(0,i.jsx)(n.code,{children:"EndBlockVSU()"}),").\nWhen collecting the validator updates for each consumer chain (see ",(0,i.jsx)(n.a,{href:"https://pkg.go.dev/github.com/cosmos/interchain-security/v4/x/ccv/provider/keeper#Keeper.QueueVSCPackets",children:(0,i.jsx)(n.code,{children:"QueueVSCPackets()"})}),"), the validator powers of the bonded validators will be updated with the validator powers from the external sources using the ",(0,i.jsx)(n.code,{children:"Power Mixing"})," module.\nThese updates are sent as part of the VSC packets to all registered consumer chains."]}),"\n",(0,i.jsxs)(n.h4,{id:"integration-with-ics-consumer",children:["Integration with ",(0,i.jsx)(n.code,{children:"ICS consumer"})]}),"\n",(0,i.jsx)(n.p,{children:"Consumer chains receive validator updates as part of VSC packets from the provider.\nThese packets contain validator powers which were already mixed with external staked powers."}),"\n",(0,i.jsx)(n.h3,{id:"queries",children:"Queries"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-protobuf",children:"// GetValidatorUpdates returns the power mixed validator results from the provided sources\nservice Query {\n  rpc GetValidatorUpdates(PowerMixedValUpdateRequest) PowerMixedValUpdateResponse {};\n}\n\n// PowerMixedValUpdateRequest contains the list of power sources on which the\n// power mixing should be based on\nmessage PowerMixedValUpdateRequest {\n  repeated PowerSource sources;\n}\n\n// PowerMixedValUpdateResponse returns the validator set with the updated powers\n// from the power mixing feature\nmessage PowerMixedValUpdateResponse {\n  repeated abci.ValidatorUpdate val_set\n}\n"})}),"\n",(0,i.jsx)(n.p,{children:"The following queries will be provided by the oracle"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-protobuf",children:'service Query {\n    rpc GetExtValidators(GetExtValidatorRequest) returns (ExtValidatorsResponse) {\n         option (google.api.http).get = "oracle/v1/get_validators";\n    };\n}\n\nmessage GetExtValidatorRequest {}\n\n// ExtValidatorsResponse is the response from GetExtValidators queries\nmessage ExtValidatorsResponse {\n  repeated ExtValPower powers;\n}\n\n// ExtValPower represents a validator with its staking and token information,\n// where:\n// `power_source_identifier` is the identifier of the registered power source\n// `validator_address` is the address of the validator\n// `stakes` is the total amount of stakes for a validator\n// `denom` is the source token of the stake e.g. ETH,BTC\n// `price_ratio` is the ratio of price of the external token to the price of the \'local\' token\nmessage ExtValPower {\n  string power_source_identifier;\n  string validator_address;\n  uint64 stakes;\n  string denom;\n  float  price_ratio;\n}\n\n// GetPrice returns a price feed for a given token\nservice Query {\n  rpc GetPrice(GetPriceRequest) returns (GetPriceResponse) {\n    option (google.api.http).get = "/oracle/v1/get_price";\n  };\n}\n'})}),"\n",(0,i.jsx)(n.p,{children:"For security reasons the amount of external stakes needs to be limited. Limitation of external staking could be driven by governance and is not subject of this version of the ADR."}),"\n",(0,i.jsx)(n.h3,{id:"reward-handler",children:"Reward Handler"}),"\n",(0,i.jsxs)(n.p,{children:["For native staked tokens the ",(0,i.jsx)(n.code,{children:"Distribution Module"})," of the Cosmos SDK is taking care of sending the rewards to stakers.\nFor stakes originated from external chains (Ethereum/Bitcoin) the ",(0,i.jsx)(n.code,{children:"Reward Handler"})," module sends rewards to EigenLayer/Babylon.\nThe transfer of rewards is done using a bridge between the Cosmos chain and the external provider chain."]}),"\n",(0,i.jsxs)(n.p,{children:["Note: currently there's no support paying rewards on EigenLayer (see ",(0,i.jsx)(n.a,{href:"https://www.coindesk.com/tech/2024/04/10/eigenlayer-cryptos-biggest-project-launch-this-year-is-still-missing-crucial-functionality/",children:"here"}),")"]}),"\n",(0,i.jsx)(n.h2,{id:"consequences",children:"Consequences"}),"\n",(0,i.jsx)(n.h3,{id:"positive",children:"Positive"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"Allow external depositors to stake their tokens to secure a Cosmos chain"}),"\n"]}),"\n",(0,i.jsx)(n.h3,{id:"negative",children:"Negative"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"Dependency to external sources e.g (price feeds) for validator power calculation"}),"\n",(0,i.jsx)(n.li,{children:"Security impact"}),"\n"]}),"\n",(0,i.jsx)(n.h3,{id:"neutral",children:"Neutral"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"Additional complexity for staking"}),"\n"]}),"\n",(0,i.jsx)(n.h2,{id:"questions",children:"Questions:"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:["Slashing: subject of this ADR? (Defined but ",(0,i.jsx)(n.a,{href:"https://www.coindesk.com/tech/2024/04/10/eigenlayer-cryptos-biggest-project-launch-this-year-is-still-missing-crucial-functionality/",children:"not activated"})," currently on EigenLayer)."]}),"\n"]}),"\n",(0,i.jsx)(n.h2,{id:"references",children:"References"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"https://docs.eigenlayer.xyz/",children:"EigenLayer"})}),"\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"https://babylonchain.io/",children:"Babylon"})}),"\n"]})]})}function h(e={}){const{wrapper:n}={...(0,t.a)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(c,{...e})}):c(e)}},1151:(e,n,r)=>{r.d(n,{Z:()=>o,a:()=>s});var i=r(7294);const t={},a=i.createContext(t);function s(e){const n=i.useContext(a);return i.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function o(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:s(e.components),i.createElement(a.Provider,{value:n},e.children)}}}]);