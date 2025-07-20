# CustomHashTable & HashTableGroups

<p align="center">
  <img alt="Go Version" src="https://img.shields.io/badge/Go-1.22-blue?logo=go" />
  <img alt="Build Status" src="https://img.shields.io/badge/build-passing-brightgreen" />
  <img alt="License" src="https://img.shields.io/badge/License-MIT-success" />
  <img alt="Coverage" src="https://img.shields.io/badge/tests-100%25%20coverage-success" />
</p>

**CustomHashTable** is a lightweight educational implementation of a hash table in Go, designed to shed light on hashing internals, collision handling, and performance‑tuning techniques.

The project ships with two core structures:

| Structure               | Purpose | Highlights |
|-------------------------|---------|-----------|
| `HashTable[K, V]`       | A standalone, generic hash table | • Separate‑chaining for collisions<br>• Dynamic resizing<br>• Clean, well‑commented Go code |
| `HashTableGroups[G, T]` | A collection of hash tables, one per *group key* | • Nested / sharded / multi‑tenant data layouts<br>• Lazily creates sub‑tables<br>• Useful for region or tenant isolation |

---

## ✨ Features

- **Generics‑based API** (Go ≥1.18)
- **Separate‑chaining** collision resolution
- **Dynamic resizing** with configurable load factor
- **HashTableGroups** for nested or sharded data models
- **100 % unit‑test coverage** with micro‑benchmarks
- Minimal dependencies, idiomatic Go

---

## 📦 Installation

```bash
go get github.com/egoryolkin/customhashtable
```
---

## 🔬 Benchmarks

Run micro‑benchmarks to gauge performance under different workloads:

### For HashTableRMBasic

```bash
go test ./internal/hash_table -bench=. -benchtime=5s -benchmem
```

### For HashTableGroupRMBasic

```bash
go test ./internal/hash_table_group -bench=. -benchtime=5s -benchmem
```

---

## 🛠️ Roadmap

- Open‑addressing variant (linear & quadratic probing)
- Thread‑safe implementation using `sync.RWMutex`
- Iterator API
- Profiling & memory insight docs

---
