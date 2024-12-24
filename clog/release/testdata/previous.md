# Changelog

All notable changes to this project will be documented in this file.

The format is based on Keep a Changelog, and this project adheres to Semantic Versioning.

## [v1.0.0](https://github.com/prysmaticlabs/prysm/compare/v5.1.2...v5.1.3) - 2024-12-04

### Added

- Add error counter for SSE endpoint. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14681)
- Add error count prom metric. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14670)
- Improve connection/disconnection logging. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14665)
- Better attestation packing for Electra. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14534)
- Add proto for `DataColumnIdentifier`, `DataColumnSidecar`, `DataColumnSidecarsByRangeRequest` and `MetadataV2`. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14649)
- validator REST: attestation v2. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14633)
- feat(issue-12348): add validator index label to validator_statuses me…. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14473)
- validator REST API: block v2 and Electra support. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14623)
- Add `POST /eth/v2/beacon/pool/attestations endpoint`. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14621)
- Add `/eth/v2/validator/aggregate_attestation`. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14481)
- Benchmark process slots. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14616)
- Allow Protobuf State To Be Created Without Copying. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14613)
- Simplify EjectedValidatorIndices. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14588)
- Rollback Block During Processing. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14554)
- Add `GET /eth/v2/beacon/pool/attestations` endpoint. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14560)
- Use engine api `get-blobs` for block subscriber. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14513)
- Update the monitor package to Electra. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14562)
- Add `/eth/v2/validator/aggregate_and_proofs`. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14490)
- Execution API Electra: requests as a sidecar. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14492)

### Changed

- Add metadata fields to getBlobSidecars. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14677)
- http response handling improvements. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14673)
- Use slot to determine fork version. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14653)
- Check if validator exists when applying pending deposit. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14666)
- Update light-client consensus types. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14652)
- Add missing Eth-Consensus-Version headers. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14647)
- Update light client protobufs. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14650)
- defer payload attribute computation. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14644)
- reorganizing p2p and backfill service registration for consistency. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14640)
- Rename remaining "deposit receipt" to "deposit request". [[PR]](https://github.com/prysmaticlabs/prysm/pull/14629)
- Optimize Message ID Computation. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14591)
- Return early blob constructor if not deneb. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14605)
- Fix various small things in state-native code. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14604)
- prevent panic by returning on connection error. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14602)
- Blocks after capella are execution. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14614)
- Build Proto State Object Once. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14612)
- Change the signature of ProcessPayload. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14610)
- Use ROBlock earlier in the pipeline. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14609)
- Fix order & add slashed=false to new validator instances. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14595)
- Electra: exclude empty requests in requests list. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14580)
- Simplify ExitedValidatorIndices. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14587)
- Use Read Only Head State When Computing Active Indices. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14592)
- Update beacon-chain pgo profile. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14589)
- Add missing version headers. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14566)
- Use ROBlock in block processing pipeline. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14571)
- Update CHANGELOG.md to reflect v5.1.2 hotfix release. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14547)
- Use read only validator for processing. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14558)
- Rollback on errors from forkchoice insertion. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14556)
- rollback on SaveState error. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14555)
- Update correlation penalty for EIP-7251. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14456)

### Removed

- Remove outdated spectest exclusions for EIP-6110. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14630)
- Remove validator count log. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14600)

### Fixed

- Fix Deadline Again During Rollback. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14686)
- `listenForNewNodes` and `FindPeersWithSubnet`: Stop using `ReadNodes` and use iterator instead. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14669)
- chore: fix 404 status URL. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14675)
- Remove kzg proof check for blob reconstructor. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14671)
- Diverse log improvements, comment additions and small refactors. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14658)
- Fix eventstream electra atts. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14655)
- adding nil checks on attestation interface. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14638)
- move get data after nil check for attestations. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14642)
- Validator REST api: adding in check for empty keys changed. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14637)
- Electra: unskipping merkle spec tests:. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14635)
- Use GetBlockAttestationV2 at handler. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14624)
- Rollback Block With Context Deadline. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14608)
- adding in a check to make sure duplicates are now allowed. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14601)
- eip7251: Bugfix and more withdrawal tests. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14578)
- small improvements to logs. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14405)
- keymanager API: bug fixes and inconsistencies. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14586)
- fix --backfill-oldest-slot flag handling. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14584)
- Fix length check between kzg commitments and exist. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14581)
- Docker fix: Update bazel-lib to latest version. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14579)
- Safe StreamEvents write loop. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14557)
- recover from panics when writing the event stream. [[PR]](https://github.com/prysmaticlabs/prysm/pull/14545)
## [v5.1.2](https://github.com/prysmaticlabs/prysm/compare/v5.1.1...v5.1.2) - 2024-10-16 

This is a hotfix release with one change. 

Prysm v5.1.1 contains an updated implementation of the beacon api streaming events endpoint. This
new implementation contains a bug that can cause a panic in certain conditions. The issue is
difficult to reproduce reliably and we are still trying to determine the root cause, but in the
meantime we are issuing a patch that recovers from the panic to prevent the node from crashing.

This only impacts the v5.1.1 release beacon api event stream endpoints. This endpoint is used by the
prysm REST mode validator (a feature which requires the validator to be configured to use the beacon
api intead of prysm's stock grpc endpoints) or accessory software that connects to the events api,
like https://github.com/ethpandaops/ethereum-metrics-exporter

### Fixed 

- Recover from panics when writing the event stream [#14545](https://github.com/prysmaticlabs/prysm/pull/14545)

## [v5.1.1](https://github.com/prysmaticlabs/prysm/compare/v5.1.0...v5.1.1) - 2024-10-15

This release has a number of features and improvements. Most notably, the feature flag 
`--enable-experimental-state` has been flipped to "opt out" via `--disable-experimental-state`. 
The experimental state management design has shown significant improvements in memory usage at
runtime. Updates to libp2p's gossipsub have some bandwidith stability improvements with support for
IDONTWANT control messages. 

The gRPC gateway has been deprecated from Prysm in this release. If you need JSON data, consider the
standardized beacon-APIs. 

Updating to this release is recommended at your convenience.

### Added

- Aggregate and proof committee validation for Electra.
- More tests for electra field generation.
- Light client support: Implement `ComputeFieldRootsForBlockBody`.
- Light client support: Add light client database changes.
- Light client support: Implement capella and deneb changes.
- Light client support: Implement `BlockToLightClientHeader` function.
- Light client support: Consensus types.
- GetBeaconStateV2: add Electra case.
- Implement [consensus-specs/3875](https://github.com/ethereum/consensus-specs/pull/3875).
- Tests to ensure sepolia config matches the official upstream yaml.
- `engine_newPayloadV4`,`engine_getPayloadV4` used for electra payload communication with execution client.  [pr](https://github.com/prysmaticlabs/prysm/pull/14492)
- HTTP endpoint for PublishBlobs.
- GetBlockV2, GetBlindedBlock, ProduceBlockV2, ProduceBlockV3: add Electra case.
- Add Electra support and tests for light client functions.
- fastssz version bump (better error messages).
- SSE implementation that sheds stuck clients. [pr](https://github.com/prysmaticlabs/prysm/pull/14413)
- Added GetPoolAttesterSlashingsV2 endpoint.
- Use engine API get-blobs for block subscriber to reduce block import latency and potentially reduce bandwidth.

### Changed

- Electra: Updated interop genesis generator to support Electra.
- `getLocalPayload` has been refactored to enable work in ePBS branch.
- `TestNodeServer_GetPeer` and `TestNodeServer_ListPeers` test flakes resolved by iterating the whole peer list to find
  a match rather than taking the first peer in the map.
- Passing spectests v1.5.0-alpha.4 and v1.5.0-alpha.5.
- Beacon chain now asserts that the external builder block uses the expected gas limit.
- Electra: Add electra objects to beacon API.
- Electra: Updated block publishing beacon APIs to support Electra.
- "Submitted builder validator registration settings for custom builders" log message moved to debug level.
- config: Genesis validator root is now hardcoded in params.BeaconConfig()
- `grpc-gateway-host` is renamed to http-host. The old name can still be used as an alias.
- `grpc-gateway-port` is renamed to http-port. The old name can still be used as an alias.
- `grpc-gateway-corsdomain` is renamed to http-cors-domain. The old name can still be used as an alias.
- `api-timeout` is changed from int flag to duration flag, default value updated.
- Light client support: abstracted out the light client headers with different versions.
- `ApplyToEveryValidator` has been changed to prevent misuse bugs, it takes a closure that takes a `ReadOnlyValidator` and returns a raw pointer to a `Validator`. 
- Removed gorilla mux library and replaced it with net/http updates in go 1.22.
- Clean up `ProposeBlock` for validator client to reduce cognitive scoring and enable further changes.
- Updated k8s-io/client-go to v0.30.4 and k8s-io/apimachinery to v0.30.4
- Migrated tracing library from opencensus to opentelemetry for both the beacon node and validator.
- Refactored light client code to make it more readable and make future PRs easier.
- Update light client helper functions to reference `dev` branch of CL specs
- Updated Libp2p Dependencies to allow prysm to use gossipsub v1.2 .
- Updated Sepolia bootnodes.
- Make committee aware packing the default by deprecating `--enable-committee-aware-packing`.
- Moved `ConvertKzgCommitmentToVersionedHash` to the `primitives` package.
- Updated correlation penalty for EIP-7251. 

### Deprecated
- `--disable-grpc-gateway` flag is deprecated due to grpc gateway removal.
- `--enable-experimental-state` flag is deprecated. This feature is now on by default. Opt-out with `--disable-experimental-state`.

### Removed

- Removed gRPC Gateway.
- Removed unused blobs bundle cache.
- Removed consolidation signing domain from params. The Electra design changed such that EL handles consolidation signature verification.
- Remove engine_getPayloadBodiesBy{Hash|Range}V2

### Fixed

- Fixed early release of read lock in BeaconState.getValidatorIndex.
- Electra: resolve inconsistencies with validator committee index validation.
- Electra: build blocks with blobs.
- E2E: fixed gas limit at genesis
- Light client support: use LightClientHeader instead of BeaconBlockHeader.
- validator registration log changed to debug, and the frequency of validator registration calls are reduced
- Core: Fix process effective balance update to safe copy validator for Electra.
- `== nil` checks before calling `IsNil()` on interfaces to prevent panics.
- Core: Fixed slash processing causing extra hashing.
- Core: Fixed extra allocations when processing slashings.
- remove unneeded container in blob sidecar ssz response
- Light client support: create finalized header based on finalizedBlock's version, not attestedBlock.
- Light client support: fix light client attested header execution fields' wrong version bug.
- Testing: added custom matcher for better push settings testing.
- Registered `GetDepositSnapshot` Beacon API endpoint.
- Fix rolling back of a block due to a context deadline.

### Security

No notable security updates.

## [v5.1.0](https://github.com/prysmaticlabs/prysm/compare/v5.0.4...v5.1.0) - 2024-08-20

This release contains 171 new changes and many of these are related to Electra! Along side the Electra changes, there
are nearly 100 changes related to bug fixes, feature additions, and other improvements to Prysm. Updating to this
release is recommended at your convenience.

⚠️ Deprecation Notice: Removal of gRPC Gateway and Gateway Flag Renaming ⚠️

In an upcoming release, we will be deprecating the gRPC gateway and renaming several associated flags. This change will
result in the removal of access to several internal APIs via REST, though the gRPC endpoints will remain unaffected. We
strongly encourage systems to transition to using the beacon API endpoints moving forward. Please refer to PR for more
details.

### Added

- Electra work
- Fork-specific consensus-types interfaces
- Fuzz ssz roundtrip marshalling, cloner fuzzing
- Add support for multiple beacon nodes in the REST API
- Add middleware for Content-Type and Accept headers
- Add debug logs for proposer settings
- Add tracing to beacon api package
- Add support for persistent validator keys when using remote signer. --validators-external-signer-public-keys and
  --validators-external-signer-key-file See the docs page for more info.
- Add AggregateKeyFromIndices to beacon state to reduce memory usage when processing attestations
- Add GetIndividualVotes endpoint
- Implement is_better_update for light client
- HTTP endpoint for GetValidatorParticipation
- HTTP endpoint for GetChainHead
- HTTP endpoint for GetValidatorActiveSetChanges
- Check locally for min-bid and min-bid-difference

### Changed

- Refactored slasher operations to their logical order
- Refactored Gwei and Wei types from math to primitives package.
- Unwrap payload bid from ExecutionData
- Change ZeroWei to a func to avoid shared ptr
- Updated go-libp2p to v0.35.2 and go-libp2p-pubsub to v0.11.0
- Use genesis block root in epoch 1 for attester duties
- Cleanup validator client code
- Old attestations log moved to debug. "Attestation is too old to broadcast, discarding it"
- Modify ProcessEpoch not to return the state as a returned value
- Updated go-bitfield to latest release
- Use go ticker instead of timer
- process_registry_updates no longer makes a full copy of the validator set
- Validator client processes sync committee roll separately
- Use vote pointers in forkchoice to reduce memory churn
- Avoid Cloning When Creating a New Gossip Message
- Proposer filters invalid attestation signatures
- Validator now pushes proposer settings every slot
- Get all beacon committees at once
- Committee-aware attestation packing

### Deprecated

- `--enable-debug-rpc-endpoints` is deprecated and debug rpc points are on by default.

### Removed

- Removed fork specific getter functions (i.e. PbCapellaBlock, PbDenebBlock, etc)

### Fixed

- Fixed debug log "upgraded stake to $fork" to only log on upgrades instead of every state transition
- Fixed nil block panic in API
- Fixed mockgen script
- Do not fail to build block when block value is unknown
- Fix prysmctl TUI when more than 20 validators were listed
- Revert peer backoff changes from. This was causing some sync committee performance issues.
- Increased attestation seen cache expiration to two epochs
- Fixed slasher db disk usage leak
- fix: Multiple network flags should prevent the BN to start
- Correctly handle empty payload from GetValidatorPerformance requests
- Fix Event stream with carriage return support
- Fix panic on empty block result in REST API
- engine_getPayloadBodiesByRangeV1 - fix, adding hexutil encoding on request parameters
- Use sync committee period instead of epoch in `createLightClientUpdate`

### Security

- Go version updated to 1.22
