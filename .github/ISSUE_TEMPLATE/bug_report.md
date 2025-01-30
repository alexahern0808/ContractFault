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
