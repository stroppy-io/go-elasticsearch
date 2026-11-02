# stroppy-io/go-elasticsearch

A fork of [elastic/go-elasticsearch](https://github.com/elastic/go-elasticsearch),
kept for one purpose: generating and running the client's REST API tests
(from the YAML tests in
[elastic/elasticsearch-clients-tests](https://github.com/elastic/elasticsearch-clients-tests))
against Elasticsearch-compatible servers in
[stroppy](https://github.com/stroppy-io/stroppy)'s conformity suites. It is not
a separate client: use upstream for anything else.

| | |
|---|---|
| Upstream | https://github.com/elastic/go-elasticsearch |
| Licence | Apache-2.0, as upstream ([LICENSE](LICENSE), [NOTICE](NOTICE)) |

| Major | Base | Fork branch | Fork commit |
|---|---|---|---|
| v7 | tag `v7.17.10`, commit `559ea470985d31b06a20840257f8370b4c575c30` | [`conformity/product-skip-v7.17.10`](https://github.com/stroppy-io/go-elasticsearch/tree/conformity/product-skip-v7.17.10) | [`d6959e77cfdd5f3814305b28de5d9884979caef2`](https://github.com/stroppy-io/go-elasticsearch/commit/d6959e77cfdd5f3814305b28de5d9884979caef2) |
| v8 | tag `v8.19.7`, commit `a5e31f9079718ed356f5cfe1959c95b3b33c4be4` | [`conformity/product-skip-v8.19.7`](https://github.com/stroppy-io/go-elasticsearch/tree/conformity/product-skip-v8.19.7) | [`826259c944068ff16580254d983e68e1c6b23e95`](https://github.com/stroppy-io/go-elasticsearch/commit/826259c944068ff16580254d983e68e1c6b23e95) |
| v9 | tag `v9.5.2`, commit `3057ff1608d6fb9b88efe1f550dfd953b720c0d0` | [`conformity/product-skip-v9.5.2`](https://github.com/stroppy-io/go-elasticsearch/tree/conformity/product-skip-v9.5.2) | [`c70b3fdf74b6921d6bfd4374f41830c433d9d2ac`](https://github.com/stroppy-io/go-elasticsearch/commit/c70b3fdf74b6921d6bfd4374f41830c433d9d2ac) |

This branch (`stroppy`) holds only this README and the licence files. The
code is on the fork branches above, each of which is its upstream tag plus
the commits listed below and nothing else.

## Our changes

The test generator (`internal/build/cmd/generate/commands/gentests`) changes,
and the client gains one option the generated tests set. The client's API, its
behaviour without that option and the YAML tests are unchanged: a run that does
not set `ELASTICSEARCH_SKIP_PRODUCT_CHECK` builds and measures what upstream
does. Each change is one commit on top of the base tag.

| Commit (v8 / v9; v7) | Files | What | Why |
|---|---|---|---|
| `d8b1abc1` / `71a37637` | `gentests/model.go`, `gentests/expand_test.go` (new) | A path that starts with a stashed value (`$profile.uid`) is read inside that value: `getKey(stash["$profile"], "uid")`. A stashed value later in a path (`nodes.$master.modules`) still names a key. | The generator read the stashed value as a key of the response (`getKey(getKey(mapi, fmt.Sprintf("%v", stash["$profile"])), "uid")`), which finds nothing, so the generated `TestXpack_SecurityUserProfile` (`security/50_user_profile.yml`) failed on every server, Elasticsearch included. No other generated test changes. The bug is also on upstream `main`. |
| `de3c4bfa` / `ce1fe131` | `gentests/generator.go` | The x-pack tests' common setup creates `x_pack_rest_user` in the native realm only when it cannot authenticate as that user. | Elastic's test clusters define `x_pack_rest_user`, the superuser some x-pack tests authenticate as, in the file realm. The setup always created it in the native realm too, one native user more than a test that lists native users counts (`query_user` in `security/50_user_profile.yml` expects the two it creates). On a cluster without the file-realm user the setup does what upstream's does. This changes the setup code of every generated x-pack test. |
| `6b566def` / `dc85658f`; v7: `929d1271` | `gentests/generator.go`, `gentests/model.go`, `gentests/params_test.go` (new) | Request fields, header entries and the debug comment above each request are written in sorted order (`Action.ParamNames()`). | The generator walked the parameters and headers as Go maps, so two generations of the same YAML differed in field order, and with it which source line each call is on. Three generations are now byte-identical. No request a test sends changes. v7 carries only this commit: its tests (Elasticsearch 7.17's own `rest-api-spec` tests) need neither of the other two. |
| `826259c9` / `c70b3fdf`; v7: `d6959e77` | `elasticsearch.go`, `gentests/generator.go`, `elasticsearch_internal_test.go` | `Config.SkipProductCheck`, off by default, turns the client's product check off; the generated tests set it when `ELASTICSEARCH_SKIP_PRODUCT_CHECK` is set. | A conforming client refuses a server whose `2xx` responses lack `X-Elastic-Product: Elasticsearch` (7.17 also reads the version and tagline). That refusal is measured once as a finding of its own, but it also stops the YAML tests, the richest API test set, from measuring anything on an Elasticsearch-compatible server that does not send the header. The option measures that server's API behaviour in a separately labelled run, with the refusal still counted where it is measured. |

With the first two commits, `TestXpack_SecurityUserProfile` passes on Elasticsearch
8.19.23 with security on and `x_pack_rest_user` in the file realm.

To see the full diff against upstream:
`git diff v8.19.7 conformity/product-skip-v8.19.7`.

## Rebases

A new upstream release gets a new branch, `conformity/<change>-<tag>`, with
the same commits replayed on the new tag. Old branches are kept, so every
pinned commit stays reachable.

| Date | Base | Fork branch | Fork commit | Notes |
|---|---|---|---|---|
| 2026-10-25 | `v8.19.7` (`a5e31f90`) | `conformity/stash-root-v8.19.7` | `de3c4bfa` | First fork: the first two commits. |
| 2026-10-25 | `v9.5.2` (`3057ff16`) | `conformity/stash-root-v9.5.2` | `ce1fe131` | The same two commits on the v9 generator. |
| 2026-10-25 | `v8.19.7` (`a5e31f90`) | `conformity/stash-root-v8.19.7` | `6b566def` | Added the sorted field order; no rebase. |
| 2026-10-25 | `v9.5.2` (`3057ff16`) | `conformity/stash-root-v9.5.2` | `dc85658f` | Added the sorted field order; no rebase. |
| 2026-10-25 | `v7.17.10` (`559ea470`) | `conformity/field-order-v7.17.10` | `929d1271` | The sorted field order on the v7 generator. |
| 2026-11-02 | `v8.19.7` (`a5e31f90`) | `conformity/product-skip-v8.19.7` | `826259c9` | The product-check option, on the three existing commits. |
| 2026-11-02 | `v9.5.2` (`3057ff16`) | `conformity/product-skip-v9.5.2` | `c70b3fdf` | The same option on the v9 branch. |
| 2026-11-02 | `v7.17.10` (`559ea470`) | `conformity/product-skip-v7.17.10` | `d6959e77` | The same option on the v7 branch, which carries the sorted field order only. |
