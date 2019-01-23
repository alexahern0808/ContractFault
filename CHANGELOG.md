# Changelog

All notable changes to ContractFault are documented here. The format follows
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/) and the project adheres
to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- (planned) monorepo mode: multi-service contracts in a single combined report.
- (planned) OpenAPI import shim for existing documents.

## [1.0.0] - 2026-08-09

### Added

- `-fail-on-behavioral` pipeline flag: treat behavioral-only shifts as a hard
  failure (exit 2) when the SLA demands it.
- `-quiet` mode printing only the one-line verdict summary.

### Changed

- Stabilized the report schema at `contractfault/v1` for 1.x.

## [0.8.0] - 2025-11-18

### Changed

- Determinism hardening: every collection sorted before emission; identical
  inputs now produce byte-identical JSON reports (verified in tests).
- Text renderer magnitude meter bounded and stable across terminals.

### Fixed

- Consumer manifests with unknown keys now fail loudly instead of silently
  disarming the blast-radius join.

## [0.7.0] - 2024-11-14

### Added

- TypeScript seismic viewer: colorized terminal impact map and a standalone
  animated SVG seismograph rendered from the JSON report.
- Viewer exit codes mirror the Go CLI (0/1/2) so it can double as a CI gate.

## [0.6.0] - 2023-09-21

### Added

- Seismic magnitude scoring on a compressed 0-10 scale with plain-language
  verdicts (`stable`, `tremor`, `shaken`, `rupture`).
- CI exit-code mapping: 0 stable, 1 shaken (behavioral), 2 rupture (breaking),
  3 usage/IO error.

## [0.5.0] - 2022-10-12

### Added
