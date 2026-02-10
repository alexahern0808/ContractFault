---
name: Bug report
about: A classification or report that looks wrong
title: "[bug] "
labels: ""
assignees: ""
---

**What happened?**

A clear description of the wrong behavior: the rule code you got vs the rule
code you expected.

**Reproducer**

Attach (or inline) the two contract versions and the consumer manifests. Small
files only — trim them to the endpoints/types involved.

```text
contractfault -old old.json -new new.json -consumers "consumers/*.json" -format json
```

**Version**

The git revision (or release tag) you ran.

**Expected vs actual**

- Expected magnitude / verdict / affected consumers:
- Actual output (trimmed):

---

*Note: classification disputes usually turn out to be documented behavior —
check the rule table in `docs/CONTRACT.md#33-rule-table` first. If the table
and the code disagree, the code is wrong; if you disagree with the table, open
a `feature_request` arguing the rule, not a bug.*
