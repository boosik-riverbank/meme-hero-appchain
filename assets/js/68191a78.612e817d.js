"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[4438],{266:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>a,contentTitle:()=>o,default:()=>h,frontMatter:()=>r,metadata:()=>d,toc:()=>c});var i=t(5893),s=t(1151);const r={sidebar_position:12,title:"Improving testing and increasing confidence"},o="ADR 011: Improving testing and increasing confidence",d={id:"adrs/adr-011-improving-test-confidence",title:"Improving testing and increasing confidence",description:"Changelog",source:"@site/docs/adrs/adr-011-improving-test-confidence.md",sourceDirName:"adrs",slug:"/adrs/adr-011-improving-test-confidence",permalink:"/interchain-security/adrs/adr-011-improving-test-confidence",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:12,frontMatter:{sidebar_position:12,title:"Improving testing and increasing confidence"},sidebar:"tutorialSidebar",previous:{title:"Standalone to Consumer Changeover",permalink:"/interchain-security/adrs/adr-010-standalone-changeover"},next:{title:"Separate Releasing",permalink:"/interchain-security/adrs/adr-012-separate-releasing"}},a={},c=[{value:"Changelog",id:"changelog",level:2},{value:"Status",id:"status",level:2},{value:"Context",id:"context",level:2},{value:"Current state of testing",id:"current-state-of-testing",level:4},{value:"Unit testing",id:"unit-testing",level:3},{value:"Integration testing",id:"integration-testing",level:3},{value:"End-to-end testing",id:"end-to-end-testing",level:3},{value:"Decision",id:"decision",level:2},{value:"1. Connect specifications to code and tooling",id:"1-connect-specifications-to-code-and-tooling",level:3},{value:"Decision context and hypothesis",id:"decision-context-and-hypothesis",level:4},{value:"Main benefit",id:"main-benefit",level:4},{value:"2. Improve e2e tooling",id:"2-improve-e2e-tooling",level:3},{value:"Matrix tests",id:"matrix-tests",level:4},{value:"Introducing e2e regression testing",id:"introducing-e2e-regression-testing",level:4},{value:"Introducing e2e CometMock tests",id:"introducing-e2e-cometmock-tests",level:4},{value:"3. Introduce innovative testing approaches",id:"3-introduce-innovative-testing-approaches",level:3},{value:"Model",id:"model",level:4},{value:"Driver",id:"driver",level:4},{value:"Harness",id:"harness",level:4},{value:"Consequences",id:"consequences",level:2},{value:"Positive",id:"positive",level:3},{value:"Negative",id:"negative",level:3},{value:"Neutral",id:"neutral",level:3},{value:"References",id:"references",level:2}];function l(e){const n={a:"a",blockquote:"blockquote",code:"code",h1:"h1",h2:"h2",h3:"h3",h4:"h4",img:"img",li:"li",ol:"ol",p:"p",strong:"strong",table:"table",tbody:"tbody",td:"td",th:"th",thead:"thead",tr:"tr",ul:"ul",...(0,s.a)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(n.h1,{id:"adr-011-improving-testing-and-increasing-confidence",children:"ADR 011: Improving testing and increasing confidence"}),"\n",(0,i.jsx)(n.h2,{id:"changelog",children:"Changelog"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"2023-08-11: Proposed, first draft of ADR."}),"\n"]}),"\n",(0,i.jsx)(n.h2,{id:"status",children:"Status"}),"\n",(0,i.jsx)(n.p,{children:"Proposed"}),"\n",(0,i.jsx)(n.h2,{id:"context",children:"Context"}),"\n",(0,i.jsx)(n.p,{children:"Testing, QA, and maintenance of interchain-security libraries is an ever-evolving area of software engineering we have to keep incrementally improving. The purpose of the QA process is to catch bugs as early as possible. In an ideal development workflow a bug should never reach production. A bug found in the specification stage is a lot cheaper to resolve than a bug discovered in production (or even in testnet). Ideally, all bugs should be found during the CI execution, and we hope that no bugs will ever even reach the testnet (although nothing can replace actual system stress test under load interacting with users)."}),"\n",(0,i.jsx)(n.p,{children:"During development and testnet operation the following types of bugs were the most commonly found:"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"improper iterator usage"}),"\n",(0,i.jsx)(n.li,{children:"unbounded array access/iteration"}),"\n",(0,i.jsx)(n.li,{children:"improper input handling and validation"}),"\n",(0,i.jsx)(n.li,{children:"improper cached context usage"}),"\n",(0,i.jsx)(n.li,{children:"non-determinism check (improper use of maps in go, relying on random values)"}),"\n",(0,i.jsx)(n.li,{children:"KV store management and/or how keys are defined"}),"\n",(0,i.jsx)(n.li,{children:"deserialization issues arising from consumer/provider versioning mismatch"}),"\n"]}),"\n",(0,i.jsx)(n.p,{children:"Such bugs can be discovered earlier with better tooling. Some of these bugs can induce increases in block times, chain halts, state corruption, or introduce an attack surface which is difficult to remove if other systems have started depending on that behavior."}),"\n",(0,i.jsx)(n.h4,{id:"current-state-of-testing",children:"Current state of testing"}),"\n",(0,i.jsx)(n.p,{children:"Our testing suites consist of multiple parts, each with their own trade-offs and benefits with regards to code coverage, complexity and confidence they provide."}),"\n",(0,i.jsx)(n.h3,{id:"unit-testing",children:"Unit testing"}),"\n",(0,i.jsxs)(n.p,{children:["Unit testing is employed mostly for testing single-module functionality. It is the first step in testing and often the most practical. While highly important, unit tests often ",(0,i.jsx)(n.strong,{children:"test a single piece of code"})," and don't test relationships between different moving parts, this makes them less valuable when dealing with multi-module interactions."]}),"\n",(0,i.jsx)(n.p,{children:"Unit tests often employ mocks to abstract parts of the system that are not under test. Mocks are not equivalent to actual models and should not be treated as such."}),"\n",(0,i.jsx)(n.p,{children:"Out of all the approaches used, unit testing has the most tools available and the coverage can simply be displayed as % of code lines tested. Although this is a very nice and very easy to understand metric, it does not speak about the quality of the test coverage."}),"\n",(0,i.jsx)(n.p,{children:"Since distributed systems testing is a lot more involved, unit tests are oftentimes not sufficient to cover complex interactions. Unit tests are still necessary and helpful, but in cases where unit tests are not helpful e2e or integration tests should be favored."}),"\n",(0,i.jsx)(n.h3,{id:"integration-testing",children:"Integration testing"}),"\n",(0,i.jsxs)(n.p,{children:["With integration testing we ",(0,i.jsx)(n.strong,{children:"test the multi-module interactions"})," while isolating them from the remainder of the system.\nIntegration tests can uncover bugs that are often missed by unit tests."]}),"\n",(0,i.jsxs)(n.p,{children:["It is very difficult to gauge the actual test coverage imparted by integration tests and the available tooling is limited.\nIn interchain-security we employ the ",(0,i.jsx)(n.code,{children:"ibc-go/testing"})," framework to test interactions in-memory."]}),"\n",(0,i.jsx)(n.p,{children:"At present, integration testing does not involve the consensus layer - it is only concerned with application level state and logic."}),"\n",(0,i.jsx)(n.h3,{id:"end-to-end-testing",children:"End-to-end testing"}),"\n",(0,i.jsx)(n.p,{children:"In our context end-to-end testing comprises of tests that use the actual application binaries in an isolated environment (e.g. docker container). During test execution the inputs are meant to simulate actual user interaction, either by submitting transactions/queries using the command line or using gRPC/REST APIs and checking for state changes after an action has been performed. With this testing strategy we also include the consensus layer in all of our runs. This is the closest we can get to testing user interactions without starting a full testnet."}),"\n",(0,i.jsx)(n.p,{children:"End-to-end testing strategies vary between different teams and projects and we strive to unify our approach to the best of our ability (at least for ICS and gaia)."}),"\n",(0,i.jsx)(n.p,{children:"The available tooling does not give us significant (or relevant) line of code coverage information since most of the tools are geared towards analyzing unit tests and simple code branch evaluation."}),"\n",(0,i.jsx)(n.p,{children:"We aim to adapt our best practices by learning from other similar systems and projects such as cosmos-sdk, ibc-go and CometBFT."}),"\n",(0,i.jsx)(n.h2,{id:"decision",children:"Decision"}),"\n",(0,i.jsx)(n.h3,{id:"1-connect-specifications-to-code-and-tooling",children:"1. Connect specifications to code and tooling"}),"\n",(0,i.jsx)(n.p,{children:"Oftentimes, specifications are disconnected from the development and QA processes. This gives rise to problems where the specification does not reflect the actual state of the system and vice-versa.\nUsually specifications are just text files that are rarely used and go unmaintained after a while, resulting in consistency issues and misleading instructions/expectations about system behavior."}),"\n",(0,i.jsx)(n.h4,{id:"decision-context-and-hypothesis",children:"Decision context and hypothesis"}),"\n",(0,i.jsx)(n.p,{children:"Specifications written in a dedicated and executable specification language are easier to maintain than the ones written entirely in text.\nAdditionally, we can create models based on the specification OR make the model equivalent to a specification."}),"\n",(0,i.jsxs)(n.p,{children:["Models do not care about the intricacies of implementation and neither do specifications. Since both models and specifications care about concisely and accurately describing a system (such as a finite state machine), we see a benefit of adding model based tools (such as ",(0,i.jsx)(n.a,{href:"https://github.com/informalsystems/quint",children:"quint"}),") to our testing and development workflows."]}),"\n",(0,i.jsx)(n.h4,{id:"main-benefit",children:"Main benefit"}),"\n",(0,i.jsx)(n.p,{children:"MBT tooling can be used to generate test traces that can be executed by multiple different testing setups."}),"\n",(0,i.jsx)(n.h3,{id:"2-improve-e2e-tooling",children:"2. Improve e2e tooling"}),"\n",(0,i.jsx)(n.h4,{id:"matrix-tests",children:"Matrix tests"}),"\n",(0,i.jsxs)(n.p,{children:["Instead of only running tests against current ",(0,i.jsx)(n.code,{children:"main"})," branch we should adopt an approach where we also:"]}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:[(0,i.jsx)(n.strong,{children:"run regression tests against different released software versions"})," (",(0,i.jsx)(n.code,{children:"ICS v1 vs v2 vs v3"}),")"]}),"\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.strong,{children:"run non-determinism tests to uncover issues quickly"})}),"\n"]}),"\n",(0,i.jsxs)(n.p,{children:["Matrix tests can be implemented using ",(0,i.jsx)(n.a,{href:"https://github.com/informalsystems/CometMock",children:"CometMock"})," and refactoring our current e2e CI setup."]}),"\n",(0,i.jsx)(n.h4,{id:"introducing-e2e-regression-testing",children:"Introducing e2e regression testing"}),"\n",(0,i.jsx)(n.p,{children:"This e2e test suite would execute using a cronjob in our CI (nightly, multiple times a day etc.)"}),"\n",(0,i.jsxs)(n.p,{children:["Briefly, the same set of traces is run against different ",(0,i.jsx)(n.strong,{children:"maintained"})," versions of the software and the ",(0,i.jsx)(n.code,{children:"main"})," branch.\nThis would allow us to discover potential issues during development instead of in a testnet scenarios."]}),"\n",(0,i.jsxs)(n.p,{children:["The most valuable issues that can be discovered in this way are ",(0,i.jsx)(n.strong,{children:"state breaking changes"}),", ",(0,i.jsx)(n.strong,{children:"regressions"})," and ",(0,i.jsx)(n.strong,{children:"version incompatibilities"}),"."]}),"\n",(0,i.jsxs)(n.p,{children:["The setup is illustrated by the image below.\n",(0,i.jsx)(n.img,{alt:"e2e matrix tests",src:t(2548).Z+"",width:"2170",height:"1624"})]}),"\n",(0,i.jsx)(n.p,{children:"This table explains which versions are tested against each other for the same set of test traces:"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"\u2705 marks a passing test"}),"\n",(0,i.jsx)(n.li,{children:"\u274c marks a failing test"}),"\n"]}),"\n",(0,i.jsxs)(n.table,{children:[(0,i.jsx)(n.thead,{children:(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"USES: ICS v1 PROVIDER"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"start chain"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"add key"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"delegate"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"undelegate"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"redelegate"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"downtime"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"equivocation"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"stop chain"})})]})}),(0,i.jsxs)(n.tbody,{children:[(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:(0,i.jsx)(n.strong,{children:"v1 consumer (sdk45,ibc4.3)"})}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"})]}),(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:(0,i.jsx)(n.strong,{children:"v2 consumer (sdk45, ibc4.4)"})}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"})]}),(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:(0,i.jsx)(n.strong,{children:"v3 consumer (sdk47, ibc7)"})}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"})]}),(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:(0,i.jsx)(n.strong,{children:"main consumer"})}),(0,i.jsx)(n.td,{children:"\u274c"}),(0,i.jsx)(n.td,{children:"\u274c"}),(0,i.jsx)(n.td,{children:"\u274c"}),(0,i.jsx)(n.td,{children:"\u274c"}),(0,i.jsx)(n.td,{children:"\u274c"}),(0,i.jsx)(n.td,{children:"\u274c"}),(0,i.jsx)(n.td,{children:"\u274c"}),(0,i.jsx)(n.td,{children:"\u274c"})]}),(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:(0,i.jsx)(n.strong,{children:"neutron"})}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u274c"})]}),(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:(0,i.jsx)(n.strong,{children:"stride"})}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u274c"})]})]})]}),"\n",(0,i.jsx)(n.h4,{id:"introducing-e2e-cometmock-tests",children:"Introducing e2e CometMock tests"}),"\n",(0,i.jsxs)(n.p,{children:["CometMock is a mock implementation of the ",(0,i.jsx)(n.a,{href:"https://github.com/cometbft/cometbft",children:"CometBFT"})," consensus engine. It supports most operations performed by CometBFT while also being lightweight and relatively easy to use."]}),"\n",(0,i.jsxs)(n.p,{children:['CometMock tests allow more nuanced control of test scenarios because CometMock can "fool" the blockchain app into thinking that a certain number of blocks had passed.\n',(0,i.jsx)(n.strong,{children:"This allows us to test very nuanced scenarios, difficult edge cases and long-running operations (such as unbonding operations)."})]}),"\n",(0,i.jsx)(n.p,{children:"Examples of tests made easier with CometMock are listed below:"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"regression tests"}),"\n",(0,i.jsx)(n.li,{children:"non-determinism tests"}),"\n",(0,i.jsx)(n.li,{children:"upgrade tests"}),"\n",(0,i.jsx)(n.li,{children:"state-breaking changes"}),"\n"]}),"\n",(0,i.jsxs)(n.p,{children:["With CometMock, the ",(0,i.jsx)(n.strong,{children:"matrix test"})," approach can also be used. The image below illustrates a CometMock setup that can be used to discover non-deterministic behavior and state-breaking changes.\n",(0,i.jsx)(n.img,{alt:"e2e matrix tests",src:t(6068).Z+"",width:"3714",height:"2082"})]}),"\n",(0,i.jsx)(n.p,{children:"This table explains which versions are tested against each other for the same set of test traces:"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"\u2705 marks a passing test"}),"\n",(0,i.jsx)(n.li,{children:"\u274c marks a failing test"}),"\n"]}),"\n",(0,i.jsxs)(n.table,{children:[(0,i.jsx)(n.thead,{children:(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"SCENARIO"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"start chain"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"add key"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"delegate"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"undelegate"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"redelegate"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"downtime"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"equivocation"})}),(0,i.jsx)(n.th,{children:(0,i.jsx)(n.strong,{children:"stop chain"})})]})}),(0,i.jsxs)(n.tbody,{children:[(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:(0,i.jsx)(n.strong,{children:"v3 provi + v3 consu"})}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"})]}),(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:(0,i.jsx)(n.strong,{children:"main provi + main consu"})}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"})]}),(0,i.jsxs)(n.tr,{children:[(0,i.jsx)(n.td,{children:(0,i.jsx)(n.strong,{children:"commit provi + commit consu"})}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u274c"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u274c"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u2705"}),(0,i.jsx)(n.td,{children:"\u274c"}),(0,i.jsx)(n.td,{children:"\u274c"})]})]})]}),"\n",(0,i.jsxs)(n.p,{children:["Briefly; multiple versions of the application are run against the same CometMock instance and any deviations in app behavior would result in ",(0,i.jsx)(n.code,{children:"app hash"})," errors (the apps would be in different states after performing the same set of actions)."]}),"\n",(0,i.jsx)(n.h3,{id:"3-introduce-innovative-testing-approaches",children:"3. Introduce innovative testing approaches"}),"\n",(0,i.jsx)(n.p,{children:"When discussing e2e testing, some very important patterns emerge - especially if test traces are used instead of ad-hoc tests written by hand."}),"\n",(0,i.jsx)(n.p,{children:"We see a unique opportunity to clearly identify concerns and modularize the testing architecture."}),"\n",(0,i.jsxs)(n.p,{children:["The e2e testing frameworks can be split into a ",(0,i.jsx)(n.strong,{children:"pipeline consisting of 3 parts: model, driver and harness"}),"."]}),"\n",(0,i.jsx)(n.h4,{id:"model",children:"Model"}),"\n",(0,i.jsx)(n.p,{children:"Model is the part of the system that can emulate the behavior of the system under test.\nIdeally, it is very close to the specification and is written in a specification language such as quint, TLA+ or similar.\nOne of the purposes of the model is that it can be used to generate test traces."}),"\n",(0,i.jsx)(n.h4,{id:"driver",children:"Driver"}),"\n",(0,i.jsx)(n.p,{children:"The purpose of the driver is to accept test traces (generated by the model or written by hand), process them and provide inputs to the next part of the pipeline."}),"\n",(0,i.jsx)(n.p,{children:"Basically, the driver sits between the model and the actual infrastructure on which the test traces are being executed on."}),"\n",(0,i.jsx)(n.h4,{id:"harness",children:"Harness"}),"\n",(0,i.jsx)(n.p,{children:"Harness is the infrastructure layer of the pipeline that accepts inputs from the driver."}),"\n",(0,i.jsx)(n.p,{children:"There can be multiple harnesses as long as they can perform four things:"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"bootstrap a test execution environment (local, docker, k8s\u2026)"}),"\n",(0,i.jsx)(n.li,{children:"accept inputs from drivers"}),"\n",(0,i.jsx)(n.li,{children:"perform the action specified by the driver"}),"\n",(0,i.jsx)(n.li,{children:"report results after performing actions"}),"\n"]}),"\n",(0,i.jsx)(n.h2,{id:"consequences",children:"Consequences"}),"\n",(0,i.jsx)(n.p,{children:"The procedure outlined in this ADR is not an all-or-nothing approach. Concepts introduced here do not rely on each other, so this ADR may only be applied partially without negative impact on test coverage and code confidence."}),"\n",(0,i.jsx)(n.h3,{id:"positive",children:"Positive"}),"\n",(0,i.jsxs)(n.ol,{children:["\n",(0,i.jsx)(n.li,{children:"introduction of maintainable MBT solutions"}),"\n"]}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:'improvement over the current "difftest" setup that relies on an opinionated typescript model and go driver'}),"\n"]}),"\n",(0,i.jsxs)(n.ol,{start:"2",children:["\n",(0,i.jsx)(n.li,{children:"increased code coverage and confidence"}),"\n"]}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"using CometMock allows us to run more tests in less time"}),"\n",(0,i.jsx)(n.li,{children:"adding matrix e2e tests allows us to quickly pinpoint differences between code versions"}),"\n"]}),"\n",(0,i.jsx)(n.h3,{id:"negative",children:"Negative"}),"\n",(0,i.jsx)(n.p,{children:"It might be easier to forgo the MBT tooling and instead focus on pure property based testing"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"https://github.com/cosmos/interchain-security/pull/667",children:"PBT proof of concept"})}),"\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"https://github.com/flyingmutant/rapid",children:"property based testing in go"})}),"\n"]}),"\n",(0,i.jsx)(n.p,{children:'The solutions are potentially expensive if we increase usage of the CI pipeline - this is fixed by running "expensive" tests using a cronjob, instead of running them on every commit.'}),"\n",(0,i.jsx)(n.h3,{id:"neutral",children:"Neutral"}),"\n",(0,i.jsx)(n.p,{children:"The process of changing development and testing process is not something that can be thought of and delivered quickly. Luckily, the changes can be rolled out incrementally without impacting existing workflows."}),"\n",(0,i.jsx)(n.h2,{id:"references",children:"References"}),"\n",(0,i.jsxs)(n.blockquote,{children:["\n",(0,i.jsx)(n.p,{children:"Are there any relevant PR comments, issues that led up to this, or articles referenced for why we made the given design choice? If so link them here!"}),"\n"]}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"https://github.com/cosmos/gaia/issues/2427",children:"https://github.com/cosmos/gaia/issues/2427"})}),"\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"https://github.com/cosmos/gaia/issues/2420",children:"https://github.com/cosmos/gaia/issues/2420"})}),"\n",(0,i.jsx)(n.li,{children:(0,i.jsx)(n.a,{href:"https://github.com/cosmos/ibc-go/tree/main/e2e",children:"ibc-go e2e tests"})}),"\n"]})]})}function h(e={}){const{wrapper:n}={...(0,s.a)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(l,{...e})}):l(e)}},6068:(e,n,t)=>{t.d(n,{Z:()=>i});const i=t.p+"assets/images/cometmock_matrix_test-714f36252aff9df4214823e3145d0ef5.png"},2548:(e,n,t)=>{t.d(n,{Z:()=>i});const i=t.p+"assets/images/matrix_e2e_tests-30681305077301daaf3097e1952b54bb.png"},1151:(e,n,t)=>{t.d(n,{Z:()=>d,a:()=>o});var i=t(7294);const s={},r=i.createContext(s);function o(e){const n=i.useContext(r);return i.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function d(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(s):e.components||s:o(e.components),i.createElement(r.Provider,{value:n},e.children)}}}]);