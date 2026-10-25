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
| v8 | tag `v8.19.7`, commit `a5e31f9079718ed356f5cfe1959c95b3b33c4be4` | [`conformity/stash-root-v8.19.7`](https://github.com/stroppy-io/go-elasticsearch/tree/conformity/stash-root-v8.19.7) | [`d8b1abc1f5fbc3ba2f12ef0e55c4a8626ebf58ac`](https://github.com/stroppy-io/go-elasticsearch/commit/d8b1abc1f5fbc3ba2f12ef0e55c4a8626ebf58ac) |
| v9 | tag `v9.5.2`, commit `3057ff1608d6fb9b88efe1f550dfd953b720c0d0` | [`conformity/stash-root-v9.5.2`](https://github.com/stroppy-io/go-elasticsearch/tree/conformity/stash-root-v9.5.2) | [`71a37637d8bddc7067a11a03593ee75ad1f2bedb`](https://github.com/stroppy-io/go-elasticsearch/commit/71a37637d8bddc7067a11a03593ee75ad1f2bedb) |

This branch (`stroppy`) holds only this README and the licence files. The
code is on the fork branches above, each of which is its upstream tag plus
the commits listed below and nothing else.

## Our changes

Only the test generator (`internal/build/cmd/generate/commands/gentests`)
changes. The client, its API and the YAML tests are unchanged.

| Commit (v8 / v9) | Files | What | Why |
|---|---|---|---|
| `d8b1abc1` / `71a37637` | `gentests/model.go`, `gentests/expand_test.go` (new) | A path that starts with a stashed value (`$profile.uid`) is read inside that value: `getKey(stash["$profile"], "uid")`. A stashed value later in a path (`nodes.$master.modules`) still names a key. | The generator read the stashed value as a key of the response (`getKey(getKey(mapi, fmt.Sprintf("%v", stash["$profile"])), "uid")`), which finds nothing, so the generated `TestXpack_SecurityUserProfile` (`security/50_user_profile.yml`) failed on every server, Elasticsearch included. No other generated test changes. The bug is also on upstream `main`. |

To see the full diff against upstream:
`git diff v8.19.7 conformity/stash-root-v8.19.7`.

## Rebases

A new upstream release gets a new branch, `conformity/<change>-<tag>`, with
the same commits replayed on the new tag. Old branches are kept, so every
pinned commit stays reachable.

| Date | Base | Fork branch | Fork commit | Notes |
|---|---|---|---|---|
| 2026-10-25 | `v8.19.7` (`a5e31f90`) | `conformity/stash-root-v8.19.7` | `d8b1abc1` | First fork. |
| 2026-10-25 | `v9.5.2` (`3057ff16`) | `conformity/stash-root-v9.5.2` | `71a37637` | Same commit on the v9 generator. |
