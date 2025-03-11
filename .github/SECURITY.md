# Security Policy

## Supported versions

| Version  | Supported          |
|----------|--------------------|
| 1.0.x    | :white_check_mark: |
| 0.8.x    | :white_check_mark: |
| < 0.8    | :x:                |

## Reporting a vulnerability

ContractFault is an offline analysis tool: it reads JSON contract and manifest
files you point it at and writes reports. It makes no network calls, spawns no
processes, and executes nothing from the documents it parses.

If you still find a security-relevant issue — for example a parser crash that
could be weaponized in CI, path traversal via `-out`, or a manifest field that
escapes its validation — please report it privately:

- Email: `michaeldelali@users.noreply.github.com` (preferred, via GitHub contact)
- Or open a GitHub security advisory on this repository.

Please include the offending input files and the command line used. We aim to
