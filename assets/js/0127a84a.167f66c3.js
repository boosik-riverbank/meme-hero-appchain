"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[9657],{3114:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>c,contentTitle:()=>r,default:()=>l,frontMatter:()=>s,metadata:()=>i,toc:()=>p});var o=t(5893),a=t(1151);const s={sidebar_position:4},r="x/ccv/democracy",i={id:"build/modules/democracy",title:"x/ccv/democracy",description:"The democracy modules comprise x/ccv/democracy/staking, x/ccv/democracy/distribution and x/ccv/democracy/governance with overrides and extensions required for normal operation when participating in ICS.",source:"@site/versioned_docs/version-v5.2.0/build/modules/04-democracy.md",sourceDirName:"build/modules",slug:"/build/modules/democracy",permalink:"/interchain-security/v5.2.0/build/modules/democracy",draft:!1,unlisted:!1,tags:[],version:"v5.2.0",sidebarPosition:4,frontMatter:{sidebar_position:4},sidebar:"tutorialSidebar",previous:{title:"x/ccv/consumer",permalink:"/interchain-security/v5.2.0/build/modules/consumer"},next:{title:"Developing an ICS consumer chain",permalink:"/interchain-security/v5.2.0/consumer-development/app-integration"}},c={},p=[{value:"Staking",id:"staking",level:2},{value:"Implications for consumer chains",id:"implications-for-consumer-chains",level:3},{value:"Governators (aka. Governors)",id:"governators-aka-governors",level:4},{value:"Tokenomics",id:"tokenomics",level:4},{value:"Integration",id:"integration",level:3},{value:"1. confirm that no modules are returning validator updates",id:"1-confirm-that-no-modules-are-returning-validator-updates",level:4},{value:"2. wire the module in app.go",id:"2-wire-the-module-in-appgo",level:4},{value:"Governance",id:"governance",level:2},{value:"Integration",id:"integration-1",level:3},{value:"Distribution",id:"distribution",level:2},{value:"How it works",id:"how-it-works",level:3},{value:"Integration",id:"integration-2",level:3}];function d(e){const n={a:"a",admonition:"admonition",code:"code",em:"em",h1:"h1",h2:"h2",h3:"h3",h4:"h4",li:"li",p:"p",pre:"pre",strong:"strong",ul:"ul",...(0,a.a)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(n.h1,{id:"xccvdemocracy",children:"x/ccv/democracy"}),"\n",(0,o.jsxs)(n.p,{children:["The democracy modules comprise ",(0,o.jsx)(n.code,{children:"x/ccv/democracy/staking"}),", ",(0,o.jsx)(n.code,{children:"x/ccv/democracy/distribution"})," and ",(0,o.jsx)(n.code,{children:"x/ccv/democracy/governance"})," with overrides and extensions required for normal operation when participating in ICS."]}),"\n",(0,o.jsx)(n.p,{children:"The modules are plug-and-play and only require small wiring changes to be enabled."}),"\n",(0,o.jsxs)(n.p,{children:["For a full integration check the ",(0,o.jsx)(n.code,{children:"consumer-democracy"})," ",(0,o.jsx)(n.a,{href:"https://github.com/cosmos/interchain-security/blob/main/app/consumer-democracy/app.go",children:"example app"}),"."]}),"\n",(0,o.jsx)(n.h2,{id:"staking",children:"Staking"}),"\n",(0,o.jsxs)(n.p,{children:["The ",(0,o.jsx)(n.code,{children:"x/ccv/democracy/staking"})," module allows the cosmos-sdk ",(0,o.jsx)(n.code,{children:"x/staking"})," module to be used alongside the interchain security ",(0,o.jsx)(n.code,{children:"consumer"})," module."]}),"\n",(0,o.jsxs)(n.p,{children:["The module uses overrides that allow the full ",(0,o.jsx)(n.code,{children:"x/staking"})," functionality with one notable difference - the staking module will no longer be used to provide the validator set to the consensus engine."]}),"\n",(0,o.jsx)(n.h3,{id:"implications-for-consumer-chains",children:"Implications for consumer chains"}),"\n",(0,o.jsxs)(n.p,{children:["The ",(0,o.jsx)(n.code,{children:"x/ccv/democracy/staking"})," allows consumer chains to ",(0,o.jsx)(n.strong,{children:(0,o.jsx)(n.em,{children:"separate governance from block production"})}),".\nThe validator set coming from the provider chain does not need to participate in governance - they only provide infrastructure (create blocks and maintain consensus)."]}),"\n",(0,o.jsx)(n.h4,{id:"governators-aka-governors",children:"Governators (aka. Governors)"}),"\n",(0,o.jsxs)(n.p,{children:["Validators registered with the ",(0,o.jsx)(n.code,{children:"x/ccv/democracy/staking"})," module become ",(0,o.jsx)(n.strong,{children:"Governators"}),".\nUnlike validators, governators are not required to run any chain infrastructure since they are not signing any blocks.\nHowever, governators retain a subset of the validator properties:"]}),"\n",(0,o.jsxs)(n.ul,{children:["\n",(0,o.jsxs)(n.li,{children:["new governators can be created (via ",(0,o.jsx)(n.code,{children:"MsgCreateValidator"}),")"]}),"\n",(0,o.jsx)(n.li,{children:"governators can accept delegations"}),"\n",(0,o.jsx)(n.li,{children:"governators can vote on governance proposals (with their self stake and delegations)"}),"\n",(0,o.jsxs)(n.li,{children:["governators earn block rewards -- the block rewards kept on the consumer (see the ",(0,o.jsx)(n.a,{href:"../build/modules/03-consumer.md#consumerredistributionfraction",children:"ConsumerRedistributionFraction param"}),") are distributed to all governators and their delegators."]}),"\n"]}),"\n",(0,o.jsx)(n.p,{children:"With these changes, governators can become community advocates that can specialize in chain governance and they get rewarded for their participation the same way the validators do.\nAdditionally, governators can choose to provide additional infrastructure such as RPC/API access points, archive nodes, indexers and similar software."}),"\n",(0,o.jsx)(n.h4,{id:"tokenomics",children:"Tokenomics"}),"\n",(0,o.jsx)(n.p,{children:"The consumer chain's token will remain a governance token. The token's parameters (inflation, max supply, burn rate) are completely under the control of the consumer chain."}),"\n",(0,o.jsx)(n.h3,{id:"integration",children:"Integration"}),"\n",(0,o.jsxs)(n.p,{children:["The ",(0,o.jsx)(n.code,{children:"x/ccv/democracy/staking"})," module provides these ",(0,o.jsx)(n.code,{children:"x/staking"})," overrides:"]}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-golang",children:"\n// InitGenesis delegates the InitGenesis call to the underlying x/staking module,\n// however, it returns no validator updates as validators are tracked via the\n// consumer chain's x/cvv/consumer module and so this module is not responsible for returning the initial validator set.\nfunc (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {\n    var genesisState types.GenesisState\n\n    cdc.MustUnmarshalJSON(data, &genesisState)\n    _ = am.keeper.InitGenesis(ctx, &genesisState)  // run staking InitGenesis\n\n    return []abci.ValidatorUpdate{}  // do not return validator updates\n}\n\n// EndBlock delegates the EndBlock call to the underlying x/staking module.\n// However, no validator updates are returned as validators are tracked via the\n// consumer chain's x/cvv/consumer module.\nfunc (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {\n    _ = am.keeper.BlockValidatorUpdates(ctx)  // perform staking BlockValidatorUpdates\n    return []abci.ValidatorUpdate{}  // do not return validator updates\n}\n"})}),"\n",(0,o.jsxs)(n.p,{children:["To integrate the ",(0,o.jsx)(n.code,{children:"democracy/staking"})," follow this guide:"]}),"\n",(0,o.jsx)(n.h4,{id:"1-confirm-that-no-modules-are-returning-validator-updates",children:"1. confirm that no modules are returning validator updates"}),"\n",(0,o.jsx)(n.admonition,{type:"warning",children:(0,o.jsxs)(n.p,{children:["Only the ",(0,o.jsx)(n.code,{children:"x/ccv/consumer"})," module should be returning validator updates."]})}),"\n",(0,o.jsx)(n.p,{children:"If some of your modules are returning validator updates please disable them while maintaining your business logic:"}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-diff",children:"func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {\n    var genesisState types.GenesisState\n\n    cdc.MustUnmarshalJSON(data, &genesisState)\n-\treturn am.keeper.InitGenesis(ctx, &genesisState)\n+ \t_ = am.keeper.InitGenesis(ctx, &genesisState)  // run InitGenesis but drop the result\n+\treturn []abci.ValidatorUpdate{}  // return empty validator updates\n}\n\n\nfunc (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {\n-\treturn am.keeper.BlockValidatorUpdates(ctx)\n+ \t_ = am.keeper.BlockValidatorUpdates(ctx)  // perform staking BlockValidatorUpdates\n+\treturn []abci.ValidatorUpdate{}  // return empty validator updates\n}\n"})}),"\n",(0,o.jsx)(n.h4,{id:"2-wire-the-module-in-appgo",children:"2. wire the module in app.go"}),"\n",(0,o.jsxs)(n.p,{children:["You ",(0,o.jsx)(n.strong,{children:"do not need to remove"})," the cosmos-sdk ",(0,o.jsx)(n.code,{children:"StakingKeeper"})," from your wiring."]}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-diff",children:'import (\n    ...\n+   ccvstaking "github.com/cosmos/interchain-security/v4/x/ccv/democracy/staking"\n)\n\nvar (\n    // replace the staking.AppModuleBasic\n    ModuleBasics = module.NewBasicManager(\n        auth.AppModuleBasic{},\n        genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),\n        bank.AppModuleBasic{},\n        capability.AppModuleBasic{},\n-\t\tsdkstaking.AppModuleBasic{},\n+\t\tccvstaking.AppModuleBasic{},  // replace sdkstaking\n        ...\n    )\n)\n\n\nfunc NewApp(...) {\n    ...\n\n    // use sdk StakingKeepeer\n    app.StakingKeeper = stakingkeeper.NewKeeper(\n        appCodec,\n        keys[stakingtypes.StoreKey],\n        app.AccountKeeper,\n        app.BankKeeper,\n        authtypes.NewModuleAddress(govtypes.ModuleName).String(),\n    )\n\n    app.MintKeeper = mintkeeper.NewKeeper(\n        appCodec,\n        keys[minttypes.StoreKey],\n        app.StakingKeeper,\n        app.AccountKeeper,\n        app.BankKeeper,\n        authtypes.FeeCollectorName,\n        authtypes.NewModuleAddress(govtypes.ModuleName).String(),\n    )\n\n    // no changes required for the distribution keeper\n    app.DistrKeeper = distrkeeper.NewKeeper(\n        appCodec,\n        keys[distrtypes.StoreKey],\n        app.AccountKeeper,\n        app.BankKeeper,\n        app.StakingKeeper,  // keep StakingKeeper!\n        authtypes.FeeCollectorName,\n        authtypes.NewModuleAddress(govtypes.ModuleName).String(),\n    )\n\n+   // pre-initialize ConsumerKeeper to satsfy ibckeeper.NewKeeper\n+\tapp.ConsumerKeeper = consumerkeeper.NewNonZeroKeeper(\n+\t\tappCodec,\n+\t\tkeys[consumertypes.StoreKey],\n+\t\tapp.GetSubspace(consumertypes.ModuleName),\n+\t)\n+\n+\tapp.IBCKeeper = ibckeeper.NewKeeper(\n+\t\tappCodec,\n+\t\tkeys[ibchost.StoreKey],\n+\t\tapp.GetSubspace(ibchost.ModuleName),\n+\t\t&app.ConsumerKeeper,\n+\t\tapp.UpgradeKeeper,\n+\t\tscopedIBCKeeper,\n+\t)\n+\n+\t// Create CCV consumer and modules\n+\tapp.ConsumerKeeper = consumerkeeper.NewKeeper(\n+\t\tappCodec,\n+\t\tkeys[consumertypes.StoreKey],\n+\t\tapp.GetSubspace(consumertypes.ModuleName),\n+\t\tscopedIBCConsumerKeeper,\n+\t\tapp.IBCKeeper.ChannelKeeper,\n+\t\t&app.IBCKeeper.PortKeeper,\n+\t\tapp.IBCKeeper.ConnectionKeeper,\n+\t\tapp.IBCKeeper.ClientKeeper,\n+\t\tapp.SlashingKeeper,\n+\t\tapp.BankKeeper,\n+\t\tapp.AccountKeeper,\n+\t\t&app.TransferKeeper,\n+\t\tapp.IBCKeeper,\n+\t\tauthtypes.FeeCollectorName,\n+\t)\n+\n+\t// Setting the standalone staking keeper is only needed for standalone to consumer changeover chains\n+  \t// New chains using the democracy/staking do not need to set this\n+\tapp.ConsumerKeeper.SetStandaloneStakingKeeper(app.StakingKeeper)\n\n\n\n    // change the slashing keeper dependency\n    app.SlashingKeeper = slashingkeeper.NewKeeper(\n        appCodec,\n        legacyAmino,\n        keys[slashingtypes.StoreKey],\n-\t\tapp.StakingKeeper,\n+\t\t&app.ConsumerKeeper,  // ConsumerKeeper implements StakingKeeper interface\n        authtypes.NewModuleAddress(govtypes.ModuleName).String(),\n    )\n\n    // register slashing module StakingHooks to the consumer keeper\n+\tapp.ConsumerKeeper = *app.ConsumerKeeper.SetHooks(app.SlashingKeeper.Hooks())\n+\tconsumerModule := consumer.NewAppModule(app.ConsumerKeeper, app.GetSubspace(consumertypes.ModuleName))\n\n        // register the module with module manager\n    // replace the x/staking module\n    app.MM = module.NewManager(\n        ...\n-\t\tsdkstaking.NewAppModule(appCodec, &app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),\n+\t\tccvstaking.NewAppModule(appCodec, *app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),\n        ...\n    )\n}\n'})}),"\n",(0,o.jsx)(n.h2,{id:"governance",children:"Governance"}),"\n",(0,o.jsxs)(n.p,{children:["The ",(0,o.jsx)(n.code,{children:"x/ccv/democracy/governance"})," module extends the ",(0,o.jsx)(n.code,{children:"x/governance"})," module with the functionality to filter proposals.\nThe module uses ",(0,o.jsx)(n.code,{children:"AnteHandler"})," to limit the types of proposals that can be executed.\nAs a result, consumer chains can limit the types of governance proposals that can be executed on chain to avoid inadvertent changes to the ICS protocol that could affect security properties."]}),"\n",(0,o.jsx)(n.h3,{id:"integration-1",children:"Integration"}),"\n",(0,o.jsxs)(n.p,{children:["Add new ",(0,o.jsx)(n.code,{children:"AnteHandler"})," to your ",(0,o.jsx)(n.code,{children:"app"}),"."]}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-go",children:'\n// app/ante/forbidden_proposals.go\npackage ante\n\nimport (\n    "fmt"\n\n    sdk "github.com/cosmos/cosmos-sdk/types"\n    govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"\n    govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"\n    ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"\n\n    "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"\n    "github.com/cosmos/cosmos-sdk/x/params/types/proposal"\n)\n\ntype ForbiddenProposalsDecorator struct {\n    isLegacyProposalWhitelisted func(govv1beta1.Content) bool\n    isModuleWhiteList           func(string) bool\n}\n\nfunc NewForbiddenProposalsDecorator(\n    whiteListFn func(govv1beta1.Content) bool,\n    isModuleWhiteList func(string) bool,\n) ForbiddenProposalsDecorator {\n    return ForbiddenProposalsDecorator{\n        isLegacyProposalWhitelisted: whiteListFn,\n        isModuleWhiteList:           isModuleWhiteList,\n    }\n}\n\nfunc (decorator ForbiddenProposalsDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {\n    currHeight := ctx.BlockHeight()\n\n    for _, msg := range tx.GetMsgs() {\n        // if the message is MsgSubmitProposal, check if proposal is whitelisted\n        submitProposalMgs, ok := msg.(*govv1.MsgSubmitProposal)\n        if !ok {\n            continue\n        }\n\n        messages := submitProposalMgs.GetMessages()\n        for _, message := range messages {\n            if sdkMsg, isLegacyProposal := message.GetCachedValue().(*govv1.MsgExecLegacyContent); isLegacyProposal {\n                // legacy gov proposal content\n                content, err := govv1.LegacyContentFromMessage(sdkMsg)\n                if err != nil {\n                    return ctx, fmt.Errorf("tx contains invalid LegacyContent")\n                }\n                if !decorator.isLegacyProposalWhitelisted(content) {\n                    return ctx, fmt.Errorf("tx contains unsupported proposal message types at height %d", currHeight)\n                }\n                continue\n            }\n            // not legacy gov proposal content and not whitelisted\n            if !decorator.isModuleWhiteList(message.TypeUrl) {\n                return ctx, fmt.Errorf("tx contains unsupported proposal message types at height %d", currHeight)\n            }\n        }\n    }\n\n    return next(ctx, tx, simulate)\n}\n\nfunc IsProposalWhitelisted(content v1beta1.Content) bool {\n    switch c := content.(type) {\n    case *proposal.ParameterChangeProposal:\n        return isLegacyParamChangeWhitelisted(c.Changes)\n\n    default:\n        return false\n    }\n}\n\nfunc isLegacyParamChangeWhitelisted(paramChanges []proposal.ParamChange) bool {\n    for _, paramChange := range paramChanges {\n        _, found := LegacyWhitelistedParams[legacyParamChangeKey{Subspace: paramChange.Subspace, Key: paramChange.Key}]\n        if !found {\n            return false\n        }\n    }\n    return true\n}\n\ntype legacyParamChangeKey struct {\n    Subspace, Key string\n}\n\n// Legacy params can be whitelisted\nvar LegacyWhitelistedParams = map[legacyParamChangeKey]struct{}{\n    {Subspace: ibctransfertypes.ModuleName, Key: "SendEnabled"}:    {},\n    {Subspace: ibctransfertypes.ModuleName, Key: "ReceiveEnabled"}: {},\n}\n\n// New proposal types can be whitelisted\nvar WhiteListModule = map[string]struct{}{\n    "/cosmos.gov.v1.MsgUpdateParams":               {},\n    "/cosmos.bank.v1beta1.MsgUpdateParams":         {},\n    "/cosmos.staking.v1beta1.MsgUpdateParams":      {},\n    "/cosmos.distribution.v1beta1.MsgUpdateParams": {},\n    "/cosmos.mint.v1beta1.MsgUpdateParams":         {},\n}\n\nfunc IsModuleWhiteList(typeUrl string) bool {\n    _, found := WhiteListModule[typeUrl]\n    return found\n}\n'})}),"\n",(0,o.jsxs)(n.p,{children:["Add the ",(0,o.jsx)(n.code,{children:"AnteHandler"})," to the list of supported antehandlers:"]}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-diff",children:'// app/ante_handler.go\npackage app\n\nimport (\n    ...\n\n+\tdemocracyante "github.com/cosmos/interchain-security/v4/app/consumer-democracy/ante"\n+\tconsumerante "github.com/cosmos/interchain-security/v4/app/consumer/ante"\n+\ticsconsumerkeeper "github.com/cosmos/interchain-security/v4/x/ccv/consumer/keeper"\n)\n\ntype HandlerOptions struct {\n    ante.HandlerOptions\n\n    IBCKeeper      *ibckeeper.Keeper\n+\tConsumerKeeper ibcconsumerkeeper.Keeper\n}\n\nfunc NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {\n    ....\n\n    anteDecorators := []sdk.AnteDecorator{\n        ...\n+\t\tconsumerante.NewMsgFilterDecorator(options.ConsumerKeeper),\n+\t\tconsumerante.NewDisabledModulesDecorator("/cosmos.evidence", "/cosmos.slashing"),\n+\t\tdemocracyante.NewForbiddenProposalsDecorator(IsProposalWhitelisted, IsModuleWhiteList),\n        ...\n    }\n\n    return sdk.ChainAnteDecorators(anteDecorators...), nil\n}\n'})}),"\n",(0,o.jsxs)(n.p,{children:["Wire the module in ",(0,o.jsx)(n.code,{children:"app.go"}),"."]}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-diff",children:'// app/app.go\npackage app\nimport (\n    ...\n    sdkgov "github.com/cosmos/cosmos-sdk/x/gov"\n    govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"\n    govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"\n    govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"\n\n+\tccvgov "github.com/cosmos/interchain-security/v4/x/ccv/democracy/governance"\n)\n\nvar (\n\n    // use sdk governance module\n    ModuleBasics = module.NewBasicManager(\n        ...\n        sdkgov.NewAppModuleBasic(\n            []govclient.ProposalHandler{\n                paramsclient.ProposalHandler,\n                upgradeclient.LegacyProposalHandler,\n                upgradeclient.LegacyCancelProposalHandler,\n            },\n        ),\n    )\n)\n\nfunc NewApp(...) {\n    // retain sdk gov router and keeper registrations\n    sdkgovRouter := govv1beta1.NewRouter()\n    sdkgovRouter.\n        AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).\n        AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).\n        AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(&app.UpgradeKeeper))\n    govConfig := govtypes.DefaultConfig()\n\n    app.GovKeeper = *govkeeper.NewKeeper(\n        appCodec,\n        keys[govtypes.StoreKey],\n        app.AccountKeeper,\n        app.BankKeeper,\n        app.StakingKeeper,\n        app.MsgServiceRouter(),\n        govConfig,\n        authtypes.NewModuleAddress(govtypes.ModuleName).String(),\n    )\n\n    app.GovKeeper.SetLegacyRouter(sdkgovRouter)\n\n\n    // register the module with module manager\n    // replace the x/gov module\n    app.MM = module.NewManager(\n-\t\tsdkgov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper, IsProposalWhitelisted, app.GetSubspace(govtypes.ModuleName), IsModuleWhiteList),\n+\t\tccvgov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper, IsProposalWhitelisted, app.GetSubspace(govtypes.ModuleName), IsModuleWhiteList),\n        ...\n    )\n}\n'})}),"\n",(0,o.jsx)(n.h2,{id:"distribution",children:"Distribution"}),"\n",(0,o.jsxs)(n.p,{children:["The ",(0,o.jsx)(n.code,{children:"x/ccv/democracy/distribution"})," module allows the consumer chain to send rewards to the provider chain while retaining the logic of the ",(0,o.jsx)(n.code,{children:"x/distribution"})," module for internal reward distribution to governators and their delegators."]}),"\n",(0,o.jsx)(n.h3,{id:"how-it-works",children:"How it works"}),"\n",(0,o.jsxs)(n.p,{children:["First, a percentage of the block rewards is sent to the provider chain, where is distributed only to opted-in validators and their delegators.\nSecond, the remaining rewards get distributed to the consumer chain's governators and their delegators.\nThe percentage that is sent to the provider chain corresponds to ",(0,o.jsx)(n.code,{children:"1 - ConsumerRedistributionFraction"}),".\nFor example, ",(0,o.jsx)(n.code,{children:'ConsumerRedistributionFraction = "0.75"'})," means that the consumer chain retains 75% of the rewards, while 25% gets sent to the provider chain"]}),"\n",(0,o.jsx)(n.h3,{id:"integration-2",children:"Integration"}),"\n",(0,o.jsxs)(n.p,{children:["Change the wiring in ",(0,o.jsx)(n.code,{children:"app.go"})]}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-diff",children:'import (\n    ...\n    distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"\n    distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"\n    sdkdistr "github.com/cosmos/cosmos-sdk/x/distribution"\n\n+   ccvdistr "github.com/cosmos/interchain-security/v4/x/ccv/democracy/distribution"\n)\n\nvar (\n    // replace sdk distribution AppModuleBasic\n    ModuleBasics = module.NewBasicManager(\n        auth.AppModuleBasic{},\n        genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),\n        bank.AppModuleBasic{},\n        capability.AppModuleBasic{},\n        ccvstaking.AppModuleBasic{}, // make sure you first swap the staking keeper\n        mint.AppModuleBasic{},\n-\t\tsdkdistr.AppModuleBasic{},\n+\t\tccvdistr.AppModuleBasic{},\n    )\n)\n\nfunc NewApp(...) {\n    ....\n\n    app.DistrKeeper = distrkeeper.NewKeeper(\n        appCodec,\n        keys[distrtypes.StoreKey],\n        app.AccountKeeper,\n        app.BankKeeper,\n        app.StakingKeeper,  // connect to sdk StakingKeeper\n        consumertypes.ConsumerRedistributeName,\n        authtypes.NewModuleAddress(govtypes.ModuleName).String(),\n    )\n\n    // register with the module manager\n    app.MM = module.NewManager(\n        ...\n-\t\tsdkdistr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, *app.StakingKeeper, authtypes.FeeCollectorName,     app.GetSubspace(distrtypes.ModuleName)),\n\n+\t\tccvdistr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, *app.StakingKeeper, authtypes.FeeCollectorName, app.GetSubspace(distrtypes.ModuleName)),\n        ccvstaking.NewAppModule(appCodec, *app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),\n        ...\n    )\n}\n'})})]})}function l(e={}){const{wrapper:n}={...(0,a.a)(),...e.components};return n?(0,o.jsx)(n,{...e,children:(0,o.jsx)(d,{...e})}):d(e)}},1151:(e,n,t)=>{t.d(n,{Z:()=>i,a:()=>r});var o=t(7294);const a={},s=o.createContext(a);function r(e){const n=o.useContext(s);return o.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function i(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(a):e.components||a:r(e.components),o.createElement(s.Provider,{value:n},e.children)}}}]);