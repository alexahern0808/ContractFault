# Contributing to ContractFault

First off, thank you for considering a contribution. ContractFault is built to
stay small, sharp and dependency-free, and that philosophy shapes every rule
below.

## Ground rules

- **Go side stays stdlib-only.** No third-party modules in `go.mod`. If a
  feature genuinely needs a dependency, open a discussion first — the bar is
  "trivial to vendor and audit".
- **Viewer stays TypeScript-first.** No frameworks; the viewer must keep
  building with plain `tsc` and running under Node 20+.
- **Determinism is not negotiable.** Any report-producing change must keep
  output byte-identical for identical inputs. Add or extend a determinism test
  alongside your change.
