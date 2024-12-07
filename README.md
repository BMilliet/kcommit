# kcommit
Git commit generator on Karma style


## kcommit custom configs
By default kcommit used the basic types from karma such as:

- feat: Adding a new feature.
- fix: Fixing a bug.
- refactor: Code change that does not change functionality or fix bugs.
- docs: Documentation update.
- style: Changes related to code style (spaces, missing semicolons, etc.) without affecting the logic.
- test: Adding or updating tests.
- chore: Updates that do not affect production code (build configuration, dependencies, etc.).

But it's possible to define your own types creating a file `.kcommitrc` at the root of your project following the format below:

```json
{
  "commitTypes": [
    {
      "type": "feat",
      "description": "Adds a new feature to the project."
    },
    {
      "type": "fix",
      "description": "Fixes a bug in the code."
    },
    {
      "type": "docs",
      "description": "Updates documentation only, without changing the code."
    },
    {
      "type": "style",
      "description": "Changes that do not affect functionality (e.g., formatting, whitespace)."
    },
    {
      "type": "refactor",
      "description": "Refactors code without changing existing functionality."
    },
    {
      "type": "test",
      "description": "Adds or updates automated tests."
    },
    {
      "type": "chore",
      "description": "Auxiliary tasks, such as dependency updates or configuration changes."
    },
    {
      "type": "perf",
      "description": "Performance improvements in the code."
    },
    {
      "type": "ci",
      "description": "Changes to the continuous integration configuration."
    },
    {
      "type": "build",
      "description": "Changes related to the build system or external dependencies."
    },
    {
      "type": "revert",
      "description": "Reverts a previous commit."
    }
  ]
}
```