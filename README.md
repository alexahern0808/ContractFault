<!-- contractfault README — a seismic fault laboratory -->

<div align="center">

<img src="docs/assets/faultline-hero.svg" alt="ContractFault faultline hero" width="100%" />

# ContractFault

[![ci](https://github.com/michaeldelali/ContractFault/actions/workflows/ci.yml/badge.svg)](https://github.com/michaeldelali/ContractFault/actions/workflows/ci.yml)
![license](https://img.shields.io/badge/license-MIT-blue)
![go](https://img.shields.io/badge/go-1.24%2B-00ADD8)
![node](https://img.shields.io/badge/node-20%2B-339933)

**An API contract consumer-impact analyzer that treats every version bump like a seismic event.**

*Ingest two contract versions. Feel where the ground moves. See which downstream towns get shaken.*

</div>

---

## Why this exists — the seismology of APIs

Most "API diff" tools hand you a wall of text and let you guess what matters.
That is like handing a geologist a list of every rock that moved and asking them
to intuit whether the village downhill is in danger. **contractfault is not a
text differ.** It is a seismograph for API contracts.

The mental model is deliberately geological:

- A **contract** is a section of crust — a versioned map of endpoints and types.
- A **change** between two versions is a **tremor** along a fault line.
- Each tremor has a **category** (does the ground crack, shift, or just settle?)
  and a **severity** (how much energy was released?).
- **Consumers** are the towns built along the fault. A tremor only matters if a
  town sits on top of it.
- The aggregate energy, felt through the towns, is the **magnitude** — a single
  0–10 number with a plain-language **verdict**: `stable`, `tremor`, `shaken`,
  or `rupture`.

That framing is not decoration. It is the actual data model. contractfault
refuses to tell you "47 things changed." It tells you *"`checkout-web` (high
criticality) will rupture because `Order.couponCode` — which it reads — was
removed, and `cancelOrder` — which it calls — changed from POST to DELETE."*
