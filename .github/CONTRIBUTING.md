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
- **Every classification rule ships with a test.** New rule codes in the
  analyzer land together with table rows in `docs/CONTRACT.md` and a test in
  `internal/analyze`.

## Local workflow

```bash
make build    # Go CLI + TS viewer
make test     # go test ./... + viewer npm test
make vet      # go vet ./...
make ci       # exactly what CI runs
```

Please make sure `make ci` is green before opening a pull request.

## Pull requests

- One topic per PR. A PR that fixes a rule and adds a flag is two PRs.
- Conventional commit messages (`feat:`, `fix:`, `docs:`, `test:`, `chore:`).
- Update `CHANGELOG.md` under `[Unreleased]` as part of your change.
- Include the example you ran when the change affects output formatting — the
  expected output belongs in the PR description.

## Reporting issues

Open a bug report with the two contract versions and the consumer manifests
involved (trim them if needed — small reproducers are read, big ones are not).
For classification disputes, quote the rule code you expected versus the code
you got; that turns a debate into a bug.

## Code of conduct

The Contributor Covenant applies to everyone participating in this project —
see [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).
