"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[5265],{3653:(e,i,n)=>{n.r(i),n.d(i,{assets:()=>c,contentTitle:()=>r,default:()=>l,frontMatter:()=>a,metadata:()=>s,toc:()=>h});var o=n(5893),t=n(1151);const a={sidebar_position:6},r="Partial Set Security",s={id:"validators/partial-set-security-for-validators",title:"Partial Set Security",description:"The Partial Set Security (PSS) feature allows consumer chains to join as Opt-In or Top N.",source:"@site/versioned_docs/version-v5.2.0/validators/partial-set-security-for-validators.md",sourceDirName:"validators",slug:"/validators/partial-set-security-for-validators",permalink:"/interchain-security/v5.2.0/validators/partial-set-security-for-validators",draft:!1,unlisted:!1,tags:[],version:"v5.2.0",sidebarPosition:6,frontMatter:{sidebar_position:6},sidebar:"tutorialSidebar",previous:{title:"Joining Stride",permalink:"/interchain-security/v5.2.0/validators/joining-stride"},next:{title:"Inactive Validators Integration Guide",permalink:"/interchain-security/v5.2.0/integrators/integrating_inactive_validators"}},c={},h=[{value:"Messages",id:"messages",level:2},{value:"How to opt in to a consumer chain?",id:"how-to-opt-in-to-a-consumer-chain",level:3},{value:"How to opt out from a consumer chain?",id:"how-to-opt-out-from-a-consumer-chain",level:3},{value:"How to set specific per consumer chain commission rate?",id:"how-to-set-specific-per-consumer-chain-commission-rate",level:3},{value:"Queries",id:"queries",level:2},{value:"Which chains does a validator have to validate?",id:"which-chains-does-a-validator-have-to-validate",level:3},{value:"How do you know how much voting power you need to have to be in the top N for a chain?",id:"how-do-you-know-how-much-voting-power-you-need-to-have-to-be-in-the-top-n-for-a-chain",level:3},{value:"How to retrieve all the opted-in validators on a consumer chain?",id:"how-to-retrieve-all-the-opted-in-validators-on-a-consumer-chain",level:3},{value:"How to retrieve all the consumer validators on a consumer chain?",id:"how-to-retrieve-all-the-consumer-validators-on-a-consumer-chain",level:3},{value:"How can we see the commission rate a validator has set on a consumer chain?",id:"how-can-we-see-the-commission-rate-a-validator-has-set-on-a-consumer-chain",level:3}];function d(e){const i={a:"a",admonition:"admonition",code:"code",em:"em",h1:"h1",h2:"h2",h3:"h3",li:"li",p:"p",pre:"pre",strong:"strong",ul:"ul",...(0,t.a)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(i.h1,{id:"partial-set-security",children:"Partial Set Security"}),"\n",(0,o.jsxs)(i.p,{children:["The ",(0,o.jsx)(i.a,{href:"/interchain-security/v5.2.0/features/partial-set-security",children:"Partial Set Security (PSS) feature"})," allows consumer chains to join as Opt-In or Top N.\nHere, we show how a validator can opt in, opt out, or set a custom commission rate on a consumer chain, as well\nas useful queries that a validator can use to figure out which chains it has to validate, etc."]}),"\n",(0,o.jsx)(i.h2,{id:"messages",children:"Messages"}),"\n",(0,o.jsx)(i.h3,{id:"how-to-opt-in-to-a-consumer-chain",children:"How to opt in to a consumer chain?"}),"\n",(0,o.jsx)(i.admonition,{type:"warning",children:(0,o.jsx)(i.p,{children:"A validator is automatically opted in to a Top N chain if the validator belongs to the top N% of the validators on the provider chain."})}),"\n",(0,o.jsx)(i.p,{children:"In a Top N chain, a validator that does not belong to the top N% of the validators on the provider can still choose\nto opt in to a consumer chain. In other words, validators can opt in, in both Opt-In and Top N chains."}),"\n",(0,o.jsx)(i.p,{children:"A validator can opt in to a consumer chain by issuing the following message:"}),"\n",(0,o.jsx)(i.pre,{children:(0,o.jsx)(i.code,{className:"language-bash",children:"interchain-security-pd tx provider opt-in <consumer-id> <optional consumer-pub-key>\n"})}),"\n",(0,o.jsx)(i.p,{children:"where"}),"\n",(0,o.jsxs)(i.ul,{children:["\n",(0,o.jsxs)(i.li,{children:[(0,o.jsx)(i.code,{children:"consumer-id"})," is the consumer id identifier of the consumer chain the validator wants to opt in to;"]}),"\n",(0,o.jsxs)(i.li,{children:[(0,o.jsx)(i.code,{children:"consumer-pub-key"})," corresponds to the public key the validator wants to use on the consumer chain, and it has the\nfollowing format ",(0,o.jsx)(i.code,{children:'{"@type":"/cosmos.crypto.ed25519.PubKey","key":"<key>"}'}),"."]}),"\n"]}),"\n",(0,o.jsx)(i.p,{children:"A validator can opt in to any active consumer chain, so a validator can opt in to a chain even before it launches.\nA validator can use the following command to retrieve the currently existing consumer chains:"}),"\n",(0,o.jsx)(i.pre,{children:(0,o.jsx)(i.code,{className:"language-bash",children:"interchain-security-pd query provider list-consumer-chains\n"})}),"\n",(0,o.jsxs)(i.p,{children:["By setting the ",(0,o.jsx)(i.code,{children:"consumer-pub-key"}),", a validator can both opt in to a chain and assign a\npublic key on a consumer chain. Note that a validator can always ",(0,o.jsx)(i.a,{href:"/interchain-security/v5.2.0/features/key-assignment",children:"assign"}),"\na new consumer key at a later stage. The key-assignment ",(0,o.jsx)(i.a,{href:"/interchain-security/v5.2.0/features/key-assignment#rules",children:"rules"}),"\nstill apply when setting ",(0,o.jsx)(i.code,{children:"consumer-pub-key"})," when opting in."]}),"\n",(0,o.jsx)(i.admonition,{type:"warning",children:(0,o.jsxs)(i.p,{children:["Validators are strongly recommended to assign a separate key for each consumer chain\nand ",(0,o.jsx)(i.strong,{children:"not"})," reuse the provider key across consumer chains for security reasons."]})}),"\n",(0,o.jsx)(i.p,{children:"Note that a validator is only eligible for consumer rewards from a consumer chain if the validator is opted into that chain."}),"\n",(0,o.jsx)(i.h3,{id:"how-to-opt-out-from-a-consumer-chain",children:"How to opt out from a consumer chain?"}),"\n",(0,o.jsx)(i.p,{children:"A validator can opt out from a consumer by issuing the following message:"}),"\n",(0,o.jsx)(i.pre,{children:(0,o.jsx)(i.code,{className:"language-bash",children:"interchain-security-pd tx provider opt-out <consumer-chain-id>\n"})}),"\n",(0,o.jsx)(i.p,{children:"where"}),"\n",(0,o.jsxs)(i.ul,{children:["\n",(0,o.jsxs)(i.li,{children:[(0,o.jsx)(i.code,{children:"consumer-id"})," is the consumer identifier of the consumer chain."]}),"\n"]}),"\n",(0,o.jsx)(i.p,{children:"The opting out mechanism has the following rules:"}),"\n",(0,o.jsxs)(i.ul,{children:["\n",(0,o.jsx)(i.li,{children:"A validator cannot opt out from a Top N chain if it belongs to the top N% validators of the provider."}),"\n",(0,o.jsxs)(i.li,{children:["If a validator moves from the Top N to outside of the top N% of the validators on the provider, it will ",(0,o.jsx)(i.strong,{children:"not"}),"\nbe automatically opted-out. The validator has to manually opt out."]}),"\n",(0,o.jsxs)(i.li,{children:["A validator should stop its node on a consumer chain ",(0,o.jsx)(i.strong,{children:"only"})," after opting out and confirming through the ",(0,o.jsx)(i.code,{children:"has-to-validate"}),"\nquery (see ",(0,o.jsx)(i.a,{href:"/interchain-security/v5.2.0/validators/partial-set-security-for-validators#which-chains-does-a-validator-have-to-validate",children:"below"}),") that it does\nnot have to validate the consumer chain any longer. Otherwise, the validator risks getting jailed for downtime."]}),"\n"]}),"\n",(0,o.jsx)(i.admonition,{type:"warning",children:(0,o.jsxs)(i.p,{children:["If all validators opt out from an Opt-In chain, the chain will halt with a consensus failure upon receiving the ",(0,o.jsx)(i.code,{children:"VSCPacket"})," with an empty validator set."]})}),"\n",(0,o.jsx)(i.h3,{id:"how-to-set-specific-per-consumer-chain-commission-rate",children:"How to set specific per consumer chain commission rate?"}),"\n",(0,o.jsx)(i.p,{children:"A validator can choose to set a different commission rate on each of the consumer chains.\nThis can be done with the following command:"}),"\n",(0,o.jsx)(i.pre,{children:(0,o.jsx)(i.code,{className:"language-bash",children:"interchain-security-pd tx provider set-consumer-commission-rate <consumer-id> <commission-rate>\n"})}),"\n",(0,o.jsx)(i.p,{children:"where"}),"\n",(0,o.jsxs)(i.ul,{children:["\n",(0,o.jsxs)(i.li,{children:[(0,o.jsx)(i.code,{children:"consumer-id"})," is the consumer identifier of the consumer chain;"]}),"\n",(0,o.jsxs)(i.li,{children:[(0,o.jsx)(i.code,{children:"comission-rate"})," decimal in ",(0,o.jsx)(i.code,{children:"[minRate, 1]"})," where ",(0,o.jsx)(i.code,{children:"minRate"})," corresponds to the minimum commission rate set on the\nprovider chain (see ",(0,o.jsx)(i.code,{children:"min_commission_rate"})," in ",(0,o.jsx)(i.code,{children:"interchain-security-pd query staking params"}),")."]}),"\n"]}),"\n",(0,o.jsx)(i.p,{children:"If a validator does not set a commission rate on a consumer chain, the commission rate defaults to their commission rate on the provider chain."}),"\n",(0,o.jsx)(i.p,{children:"Validators can set their commission rate even for consumer chains that they are not currently opted in on, and the commission rate will be applied when they opt in. This is particularly useful for Top N chains, where validators might be opted in automatically,\nso validators can set the commission rate in advance."}),"\n",(0,o.jsxs)(i.p,{children:["If a validator opts out and then back in, this will ",(0,o.jsx)(i.em,{children:"not"})," reset their commission rate back to the default. Instead, their\nset commission rate still applies."]}),"\n",(0,o.jsx)(i.h2,{id:"queries",children:"Queries"}),"\n",(0,o.jsx)(i.p,{children:"PSS introduces a number of queries to assist validators in determining which consumer chains they have to validate, their commission rate per chain, etc."}),"\n",(0,o.jsx)(i.h3,{id:"which-chains-does-a-validator-have-to-validate",children:"Which chains does a validator have to validate?"}),"\n",(0,o.jsx)(i.p,{children:"Naturally, a validator is aware of the Opt-In chains it has to validate because in order to validate an Opt-In chain,\na validator has to manually opt in to the chain. This is not the case for Top N chains where a validator might be required\nto validate such a chain without explicitly opting in if it belongs to the top N% of the validators on the provider."}),"\n",(0,o.jsx)(i.p,{children:"We introduce the following query:"}),"\n",(0,o.jsx)(i.pre,{children:(0,o.jsx)(i.code,{className:"language-bash",children:"interchain-security-pd query provider has-to-validate <provider-validator-address>\n"})}),"\n",(0,o.jsxs)(i.p,{children:["that can be used by validator with ",(0,o.jsx)(i.code,{children:"provider-validator-address"})," address to retrieve the list of chains that it has to validate."]}),"\n",(0,o.jsx)(i.admonition,{type:"warning",children:(0,o.jsxs)(i.p,{children:["For a validator, the list of chains returned by ",(0,o.jsx)(i.code,{children:"has-to-validate"})," is the list of chains the validator ",(0,o.jsx)(i.strong,{children:"should"})," be validating to avoid\ngetting jailed for downtime."]})}),"\n",(0,o.jsx)(i.h3,{id:"how-do-you-know-how-much-voting-power-you-need-to-have-to-be-in-the-top-n-for-a-chain",children:"How do you know how much voting power you need to have to be in the top N for a chain?"}),"\n",(0,o.jsxs)(i.p,{children:["This can be seen as part of the ",(0,o.jsx)(i.code,{children:"list-consumer-chains"})," query:"]}),"\n",(0,o.jsx)(i.pre,{children:(0,o.jsx)(i.code,{className:"language-bash",children:"interchain-security-pd query provider list-consumer-chains\n"})}),"\n",(0,o.jsxs)(i.p,{children:["where the ",(0,o.jsx)(i.code,{children:"min_power_in_top_N"})," field shows the minimum voting power required to be\nautomatically opted in to the chain."]}),"\n",(0,o.jsxs)(i.p,{children:["Note that ",(0,o.jsx)(i.code,{children:"list-consumer-chains"})," shows the minimal voting power ",(0,o.jsx)(i.em,{children:"right now"}),", but\nthe automatic opt-in happens only when epochs end on the provider.\nIn consequence, a validators power might be large enough to be automatically opted in\nduring an epoch, but if their power is sufficiently decreased before the epoch ends,\nthey will not be opted in automatically."]}),"\n",(0,o.jsx)(i.h3,{id:"how-to-retrieve-all-the-opted-in-validators-on-a-consumer-chain",children:"How to retrieve all the opted-in validators on a consumer chain?"}),"\n",(0,o.jsx)(i.p,{children:"With the following query:"}),"\n",(0,o.jsx)(i.pre,{children:(0,o.jsx)(i.code,{className:"language-bash",children:"interchain-security-pd query provider consumer-opted-in-validators <consumer-id>\n"})}),"\n",(0,o.jsxs)(i.p,{children:["we can see all the opted-in validators on ",(0,o.jsx)(i.code,{children:"consumer-id"})," that were manually or automatically opted in."]}),"\n",(0,o.jsx)(i.h3,{id:"how-to-retrieve-all-the-consumer-validators-on-a-consumer-chain",children:"How to retrieve all the consumer validators on a consumer chain?"}),"\n",(0,o.jsx)(i.p,{children:"With the following query:"}),"\n",(0,o.jsx)(i.pre,{children:(0,o.jsx)(i.code,{className:"language-bash",children:"interchain-security-pd query provider consumer-validators <consumer-id>\n"})}),"\n",(0,o.jsxs)(i.p,{children:["we can see all the ",(0,o.jsx)(i.em,{children:"consumer validators"})," (i.e., validator set) of ",(0,o.jsx)(i.code,{children:"consumer-id"}),". The consumer validators are the\nones that are currently (or in the future, see warning) validating the consumer chain. A ",(0,o.jsx)(i.em,{children:"consumer validator"})," is an opted-in\nvalidator but not vice versa. For example, an opted-in validator ",(0,o.jsx)(i.code,{children:"V"})," might not be a consumer validator because ",(0,o.jsx)(i.code,{children:"V"})," is\ndenylisted or because ",(0,o.jsx)(i.code,{children:"V"})," is removed due to a validator-set cap."]}),"\n",(0,o.jsxs)(i.p,{children:["Note that the returned consumer validators from this query do not necessarily correspond to the validator set that is\nvalidating the consumer chain at this exact moment. This is because the ",(0,o.jsx)(i.code,{children:"VSCPacket"})," sent to a consumer chain might be\ndelayed and hence this query might return the validator set that the consumer chain would have at some future\npoint in time."]}),"\n",(0,o.jsx)(i.h3,{id:"how-can-we-see-the-commission-rate-a-validator-has-set-on-a-consumer-chain",children:"How can we see the commission rate a validator has set on a consumer chain?"}),"\n",(0,o.jsx)(i.p,{children:"Using the following query:"}),"\n",(0,o.jsx)(i.pre,{children:(0,o.jsx)(i.code,{className:"language-bash",children:"interchain-security-pd query provider validator-consumer-commission-rate <consumer-id> <provider-validator-address>\n"})}),"\n",(0,o.jsxs)(i.p,{children:["we retrieve the commission rate set by validator with ",(0,o.jsx)(i.code,{children:"provider-validator-address"})," address on ",(0,o.jsx)(i.code,{children:"consumer-id"}),"."]})]})}function l(e={}){const{wrapper:i}={...(0,t.a)(),...e.components};return i?(0,o.jsx)(i,{...e,children:(0,o.jsx)(d,{...e})}):d(e)}},1151:(e,i,n)=>{n.d(i,{Z:()=>s,a:()=>r});var o=n(7294);const t={},a=o.createContext(t);function r(e){const i=o.useContext(a);return o.useMemo((function(){return"function"==typeof e?e(i):{...i,...e}}),[i,e])}function s(e){let i;return i=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:r(e.components),o.createElement(a.Provider,{value:i},e.children)}}}]);