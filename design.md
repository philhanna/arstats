# Architecture: Ports and Adapters (Hexagonal)

This project was restructured from MVC to a Ports and Adapters (Hexagonal) architecture.

## Directory Structure

```
internal/
  domain/
    statistics.go       — pure business logic (Statistics type + calculations)
    names.go            — domain helpers (ToDisplayName, ToSectionName)
  ports/
    repository.go       — GameRepository interface (the secondary port)
  application/
    service.go          — use cases: ListGames, MostRecentGame, GetStatistics
  adapters/
    inifile/
      parser.go         — INI file parsing (pure bytes → section map)
      repository.go     — IniRepository implements GameRepository
    cli/
      presenter.go      — CLI output (List, PrintStatistics, ErrorMessage)
cmd/
  arstats.go            — composition root: wires adapters → port → service
```

## Layer Responsibilities

### Domain (`internal/domain/`)
Pure business logic with zero infrastructure dependencies. Contains the `Statistics`
type and all its calculations, plus the name-conversion helpers `ToDisplayName` and
`ToSectionName`. Nothing in this layer performs I/O or knows about file formats.

### Ports (`internal/ports/`)
Interface definitions that decouple the application from infrastructure. The
`GameRepository` interface is the **secondary (driven) port** — it expresses what
data the application needs without specifying where that data comes from.

### Application (`internal/application/`)
Use-case orchestration via `GameService`. Calls the port and domain functions;
performs no I/O and has no knowledge of file formats or output formatting.

### Adapters (`internal/adapters/`)
Concrete implementations that connect the application to the outside world.

- **`inifile/`** — secondary adapter. Reads and parses the Aisleriot INI config
  file and implements `GameRepository`. Swap this out to support a different
  backing store without touching any other layer.
- **`cli/`** — primary adapter. Formats and prints output to the terminal.
  Receives pre-fetched domain values; never reaches into the data store itself.

### Composition Root (`cmd/arstats.go`)
The only place where concrete types are wired together. Creates the `IniRepository`,
passes it to `NewGameService`, parses CLI flags, calls service methods, and delegates
output to the CLI presenter.

## Dependency Graph

```
domain ← ports ← application ← adapters ← cmd
```

No cycles. The domain has no inward dependencies. Each outer layer may depend on
inner layers but not vice versa.

## Key Design Decisions

- **`GameRepository` interface** decouples the application from the INI file format.
  A different data source (database, JSON file, etc.) can be supported by implementing
  the interface — no other layer needs to change.

- **CLI presenter receives pre-fetched values.** `PrintStatistics` takes a
  `*domain.Statistics` directly rather than a data provider. This makes the presenter
  a pure formatter, testable without any file I/O.

- **INI parsing is split from file I/O.** `parser.go` operates on `[]byte` and has
  no file system access, making it independently unit-testable.

- **`cmd/arstats.go` is the composition root only.** No business logic lives there;
  it wires adapters to ports and delegates everything else.
